// Package discover implements Nox server discovery.
package discover

import (
	"context"
	"errors"
	"log/slog"
	"net"
	"net/netip"
	"sync"
	"time"

	"github.com/opennox/lobby"
	"golang.org/x/sync/errgroup"

	noxflags "github.com/opennox/opennox/v1/common/flags"
)

const (
	discoverTimeout = 2 * time.Second
)

const (
	priorityStatic = -2
	priorityLAN    = -1
	priorityLobby  = 0
	priorityXWIS   = 1
)

type Server struct {
	lobby.Game
	Source   string
	IP       netip.Addr // only IPv4 for now
	Priority int        // lower values mean higher priority
	Ping     time.Duration
	NoPing   bool // server doesn't support UDP pings
}

type SearchFunc func(ctx context.Context, log *slog.Logger, out chan<- Server) error

var (
	backends = make(map[string]SearchFunc)
	fallback = make(map[string]SearchFunc)
)

// RegisterBackend register a new server discovery backend.
func RegisterBackend(name string, fnc SearchFunc) {
	if _, ok := backends[name]; ok {
		panic("already registered: " + name)
	}
	backends[name] = func(ctx context.Context, log *slog.Logger, out chan<- Server) error {
		log = log.With("src", name)
		return fnc(ctx, log, out)
	}
}

// RegisterFallback register a new fallback server discovery backend.
func RegisterFallback(name string, fnc SearchFunc) {
	if _, ok := fallback[name]; ok {
		panic("already registered: " + name)
	}
	fallback[name] = fnc
}

func isTimeout(err error) bool {
	return err != nil && (errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded))
}

func serversWith(ctx context.Context, log *slog.Logger, out chan<- Server, list map[string]SearchFunc) error {
	var (
		wg   sync.WaitGroup
		errc = make(chan error, 1)
	)
	for name, fnc := range list {
		wg.Add(1)
		name, fnc := name, fnc
		go func() {
			defer wg.Done()
			if err := fnc(ctx, log, out); err != nil && !isTimeout(err) {
				log.Warn("discover error", "src", name, "err", err)
				select {
				case errc <- err:
				default:
				}
			}
		}()
	}
	wg.Wait()
	select {
	case err := <-errc:
		return err
	default:
	}
	return nil
}

type ServerFunc func(s Server) error

// EachServer discovers Nox servers and calls fnc for each of them sequentially.
// It will not ping the servers additionally. See PingEachServer.
func EachServer(ctx context.Context, log *slog.Logger, fnc ServerFunc) error {
	return eachServer(ctx, log, nil, fnc)
}

// PingEachServer discovers Nox servers, pings them via UDP and calls fnc for each of them sequentially.
// If pc is provided, it will be used to make ping requests and receive responses.
func PingEachServer(ctx context.Context, log *slog.Logger, pc *net.UDPConn, fnc ServerFunc) error {
	if pc == nil {
		l, err := net.ListenUDP("udp4", &net.UDPAddr{IP: nil, Port: 0})
		if err != nil {
			return err
		}
		defer l.Close()
		pc = l
	}
	return eachServer(ctx, log, pc, fnc)
}

func mergeInfo(g1 lobby.Game, g2 *lobby.Game) lobby.Game {
	if g2 == nil {
		return g1
	}
	if g1.Name == "" {
		g1.Name = g2.Name
	}
	if g1.Map == "" {
		g1.Map = g2.Map
	}
	if g1.Mode == "" {
		g1.Mode = g2.Mode
	}
	g1.Players.Cur = g2.Players.Cur
	g1.Players.Max = g2.Players.Max
	return g1
}

func eachServer(rctx context.Context, log *slog.Logger, pc *net.UDPConn, fnc ServerFunc) error {
	log.Info("searching for servers")
	ctx1 := rctx
	dt1, dt2 := discoverTimeout, discoverTimeout
	if deadline, ok := rctx.Deadline(); ok {
		dt := time.Until(deadline)
		dt1, dt2 = dt/2, dt/2
	}
	ctx1, cancel1 := context.WithTimeout(rctx, dt1)
	defer cancel1()
	n, err1 := eachServerWith(ctx1, log, pc, fnc, backends)
	if err1 != nil {
		log.Warn("search", "err", err1)
	}
	if n != 0 {
		return err1
	}
	log.Warn("no servers, searching with fallback")
	ctx2, cancel2 := context.WithTimeout(rctx, dt2)
	defer cancel2()
	n, err2 := eachServerWith(ctx2, log, pc, fnc, fallback)
	if n != 0 {
		return err2
	}
	if err2 != nil {
		log.Error("fallback search", "err", err2)
		if err1 != nil {
			log.Error("search failed", "err", err1)
		}
	}
	return err1
}

func eachServerWith(ctx context.Context, log *slog.Logger, pc *net.UDPConn, fnc ServerFunc, blist map[string]SearchFunc) (int, error) {
	start := time.Now()
	out := make(chan Server, 10)
	errc := make(chan error, 1)

	go func() {
		defer close(out)
		if err := serversWith(ctx, log, out, blist); err != nil {
			errc <- err
		}
	}()

	final := out
	if pc != nil {
		final = make(chan Server, 10)
		pinger := NewPinger(log, pc)
		defer pinger.Close()
		wg, ctx := errgroup.WithContext(ctx)
		for i := 0; i < 3; i++ {
			wg.Go(func() error {
				for s := range out {
					if !s.NoPing {
						s.NoPing = true
						start := time.Now()
						g, err := pinger.Ping(ctx, netip.AddrPortFrom(s.IP, uint16(s.Port)))
						if err != nil {
							log.Warn("ping error", "src", s.Source, "err", err)
						} else {
							s.Ping = time.Since(start)
							s.Game = mergeInfo(s.Game, g)
						}
					}
					select {
					case final <- s:
					case <-ctx.Done():
						return nil
					}
				}
				return nil
			})
		}
		err := wg.Wait()
		close(final)
		if err != nil {
			return 0, err
		}
	}

	n := 0
	defer func() {
		log.Info("found servers", "n", n, "dt", time.Since(start))
	}()
	for s := range final {
		n++
		if err := fnc(s); err != nil {
			return n, err
		}
	}
	select {
	case err := <-errc:
		return n, err
	default:
	}
	return n, nil
}

func ListServers(ctx context.Context, log *slog.Logger) ([]Server, error) {
	l, err := net.ListenUDP("udp4", &net.UDPAddr{IP: nil, Port: 0})
	if err != nil {
		return nil, err
	}
	defer l.Close()
	return listServersWith(ctx, log, l)
}

func ListServersWith(ctx context.Context, log *slog.Logger, pc *net.UDPConn) ([]Server, error) {
	return listServersWith(ctx, log, pc)
}

func (a Server) key() serverKey {
	return serverKey{Addr: a.IP.String(), Port: a.Port}
}

type serverKey struct {
	Addr string
	Port int
}

func listServersWith(ctx context.Context, log *slog.Logger, pc *net.UDPConn) ([]Server, error) {
	if _, ok := ctx.Deadline(); !ok {
		var cancel func()
		ctx, cancel = context.WithTimeout(ctx, discoverTimeout)
		defer cancel()
	}
	var (
		start = time.Now()

		seenMu sync.RWMutex
		seen   = make(map[serverKey]*Server)

		pingIn   = make(chan Server)
		pingResp = make(map[serverKey]*Server)
		pingDone = make(chan struct{})
	)
	go func() {
		defer close(pingDone)
		ping := NewPinger(log, pc)
		defer ping.Close()
		var (
			remaining int
			pingOut   = make(chan *lobby.Game, 50)
		)
		procPing := func(g *lobby.Game) {
			remaining--
			key := serverKey{Addr: g.Address, Port: g.Port}
			seenMu.RLock()
			s := seen[key]
			seenMu.RUnlock()
			if s != nil {
				pingResp[key] = &Server{
					Game:   *g,
					Ping:   time.Since(start),
					NoPing: true,
				}
			} else {
				log.Warn("cannot find server", "addr", g.Address, "port", g.Port)
			}
		}
	pingLoop:
		for {
			select {
			case g := <-pingOut:
				procPing(g)
			case s, ok := <-pingIn:
				if !ok {
					break pingLoop
				}
				if err := ping.SendPing(pingOut, netip.AddrPortFrom(s.IP, uint16(s.Port))); err != nil {
					log.Warn("cannot ping server", "addr", s.IP, "port", s.Port, "err", err)
				} else {
					remaining++
				}
			}
		}
		if remaining <= 0 {
			return
		}
		ctx, cancel := context.WithTimeout(ctx, pingTimeout)
		defer cancel()
	loop:
		for remaining > 0 {
			select {
			case <-ctx.Done():
				log.Info("ping result", "timeout", remaining, "total", len(seen))
				break loop
			case g := <-pingOut:
				procPing(g)
			}
		}
	}()
	err := EachServer(ctx, log, func(s Server) error {
		key := s.key()

		seenMu.Lock()
		s2, ok := seen[key]
		if !ok {
			seen[key] = &s
		} else if s.Priority < s2.Priority {
			*s2 = s
		}
		seenMu.Unlock()
		if !ok {
			select {
			case <-ctx.Done():
			case pingIn <- s:
			}
		}
		return nil
	})
	if err != nil {
		log.Warn("list servers", "err", err)
	}
	close(pingIn)
	<-pingDone
	out := make([]Server, 0, len(seen))
	for _, v := range seen {
		if p, ok := pingResp[v.key()]; ok {
			v.Ping = p.Ping
			v.NoPing = v.NoPing || p.NoPing
			v.Game = mergeInfo(v.Game, &p.Game)
		}
		out = append(out, *v)
	}
	return out, err
}

func gameFlagsToMode(f noxflags.GameFlag) lobby.GameMode {
	switch {
	case f.Has(noxflags.GameModeKOTR):
		return lobby.ModeKOTR
	case f.Has(noxflags.GameModeCTF):
		return lobby.ModeCTF
	case f.Has(noxflags.GameModeFlagBall):
		return lobby.ModeFlagBall
	case f.Has(noxflags.GameModeChat):
		return lobby.ModeChat
	case f.Has(noxflags.GameModeArena):
		return lobby.ModeArena
	case f.Has(noxflags.GameModeElimination):
		return lobby.ModeElimination
	case f.Has(noxflags.GameModeQuest):
		return lobby.ModeQuest
	case f.Has(noxflags.GameModeCoop):
		return lobby.ModeCoop
	}
	return lobby.ModeCustom
}
