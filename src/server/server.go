package server

import (
	"context"
	"log/slog"
	"net/netip"
	"sync/atomic"
	"time"
	"unsafe"

	"github.com/opennox/libs/common"
	noxlog "github.com/opennox/libs/log"
	"github.com/opennox/libs/noxnet/netxfer"
	"github.com/opennox/libs/platform"
	"github.com/opennox/libs/strman"
	ns4 "github.com/opennox/noxscript/ns/v4"

	"github.com/opennox/libs/console"

	noxflags "github.com/opennox/opennox/v1/common/flags"
	"github.com/opennox/opennox/v1/common/gsync"
	"github.com/opennox/opennox/v1/common/ntype"
	"github.com/opennox/opennox/v1/internal/netlist"
	"github.com/opennox/opennox/v1/internal/netstr"
)

var (
	serverLast uintptr // atomic
	servers    gsync.Map[uintptr, *Server]
)

func getServer(h uintptr) *Server {
	s, _ := servers.Load(h)
	return s
}

func New(log *slog.Logger, pr console.Printer, sm *strman.StringManager) *Server {
	if log == nil {
		log = noxlog.New("server").Logger
	} else {
		log = noxlog.WithSystem(log, "server")
	}
	s := &Server{
		Log:     log,
		Printer: pr, sm: sm,
		loopHooks: make(chan func()),
		port:      common.GamePort,
		NetList:   netlist.New(),
		UseNAT:    true,
	}
	s.handle = atomic.AddUintptr(&serverLast, 1)
	servers.Store(s.handle, s)

	s.Rand.init(nil)
	s.NetStr = netstr.NewStreams(s.Frame)
	s.NetStr.IsHost = func() bool {
		return noxflags.HasGame(noxflags.GameHost)
	}
	s.NetStr.IsFlag4 = func() bool {
		return noxflags.HasGame(noxflags.GameFlag4)
	}
	s.NetStr.GameFrame = s.Frame
	s.NetStr.KeyRand = s.Rand.Logic.IntClamp
	s.NetStr.PacketDropRand = s.Rand.Other.Int
	s.Types.init(s.Log)
	s.Objs.init(s.handle)
	s.Modif.init(sm)
	s.Players.init(s)
	s.Teams.init(sm, pr)
	s.Abils.init(s)
	s.Spells.init(s)
	s.Storage.init()
	s.NoxScriptVM.Init(s)
	s.MapSend.init(s)
	s.http.init()
	return s
}

type ObjectScriptID uint32

type Server struct {
	Log *slog.Logger
	console.Printer
	handle     uintptr
	sm         *strman.StringManager
	frame      uint32
	tickRate   uint32
	rateNext   time.Duration
	updateFunc func() bool
	tickHooks  tickHooks
	loopHooks  chan func()
	ExtServer  unsafe.Pointer // populated by the caller of New

	Rand         serverRandom
	Walls        serverWalls
	WPs          serverWaypoints
	Types        serverObjTypes
	Armor        serverArmor
	Weapons      serverWeapons
	Objs         serverObjects
	Modif        serverModifiers
	NPCs         serverNPCs
	Map          serverMap
	Doors        serverDoors
	MapGroups    ServerMapGroups
	Audio        serverAudio
	Activators   serverActivators
	Players      serverPlayers
	Teams        serverTeams
	Abils        serverAbilities
	Spells       serverSpells
	Balance      serverBalance
	AI           serverAI
	Storage      serverStorage
	NoxScriptVM  NoxScriptVM
	VMs          ScriptVMs
	ScriptEvents scriptEvents

	ShouldCallMapInit  bool
	ShouldCallMapEntry bool

	NetStr    *netstr.Streams
	NetList   *netlist.List
	NetXfer   netxfer.State[XferConn]
	onXferExt []func(ind ntype.PlayerInd, obj NetXferExt)
	lis       netstr.Server
	OwnIPStr  string
	OwnIP     netip.Addr
	UseNAT    bool
	Announce  bool
	MapSend   serverMapSend
	port      int
	http      httpService
	nat       natService

	updateFunc2 func() bool

	CurrentMapXxx    func() string
	CurrentMapYyy    func() string
	NetSendPacketXxx func(a1 int, buf []byte, a4, a5, a6 int) int // Nox_xxx_netSendPacket_4E5030
	ClientConn       func() *netstr.Client
	ObjectByNetCode  func(code int) *Object
	NetXferLocal     func(data netxfer.Data)
}

func (s *Server) Close() {
	servers.Delete(s.handle)
}

func (s *Server) Strings() *strman.StringManager {
	return s.sm
}

func (s *Server) Frame() uint32 {
	return atomic.LoadUint32(&s.frame)
}

func (s *Server) SetFrame(v uint32) {
	atomic.StoreUint32(&s.frame, v)
}

func (s *Server) FrameTS() time.Duration {
	return time.Second * time.Duration(s.Frame()) / time.Duration(s.TickRate())
}

func (s *Server) IncFrame() {
	atomic.AddUint32(&s.frame, 1)
}

func (s *Server) SetInitialFrame() {
	if noxflags.HasGame(noxflags.GameHost) {
		s.SetFrame(1)
	} else {
		s.SetFrame(0)
	}
}

func (s *Server) TickRate() uint32 {
	return atomic.LoadUint32(&s.tickRate)
}

func (s *Server) SetTickRate(v uint32) {
	atomic.StoreUint32(&s.tickRate, v)
}

func (s *Server) AsFrames(dt ns4.Duration) int {
	if dt.IsInfinite() {
		return 0 // TODO: check
	}
	frames, ok := dt.Frames()
	if ok {
		return frames
	}
	dur, _ := dt.Time()
	return s.DurToFrames(dur)
}

func (s *Server) SecToFrames(sec int) uint32 {
	return s.TickRate() * uint32(sec)
}

func (s *Server) SecToFramesF(sec float64) uint32 {
	return uint32(float64(s.TickRate()) * sec)
}

func (s *Server) DurToFrames(dur time.Duration) int {
	return int(float64(s.TickRate()) * dur.Seconds())
}

func (s *Server) SetUpdateFunc(fnc func() bool) {
	s.updateFunc = fnc
}

func (s *Server) SetUpdateFunc2(fnc func() bool) {
	s.updateFunc2 = fnc
}

func (s *Server) Update() bool {
	if s.updateFunc == nil {
		return true
	}
	if !s.updateFunc() {
		return false
	}
	return true
}

func (s *Server) Update2() bool {
	if s.updateFunc2 == nil {
		return true
	}
	if !s.updateFunc2() {
		return false
	}
	return true
}

func (s *Server) QueueInLoop(ctx context.Context, fnc func()) {
	select {
	case <-ctx.Done():
	case s.loopHooks <- fnc:
	}
}

func (s *Server) RunLoopHooks() {
	s.LoopSleep(time.Millisecond)
}

func (s *Server) LoopSleep(dt time.Duration) {
	if dt <= 0 {
		return
	}
	tm := time.NewTimer(dt)
	defer tm.Stop()
	for {
		select {
		case <-tm.C:
			return
		case fnc := <-s.loopHooks:
			fnc()
		}
	}
}

func (s *Server) SetRateLimit(fps int) {
	step := time.Duration(0)
	if fps != 0 {
		step = time.Second / time.Duration(fps)
	}
	s.rateNext = platform.Ticks() + step
}

func (s *Server) RateWait() {
	ticks := platform.Ticks()
	dt := s.rateNext - ticks
	if dt > 0 {
		s.LoopSleep(dt)
	}
}

func (s *Server) SetServerPort(port int) {
	if port <= 0 {
		port = common.GamePort
	}
	s.port = port
}

func (s *Server) ServerPort() int {
	if s.port <= 0 {
		return common.GamePort
	}
	return s.port
}

func (s *Server) HTTPPort() int {
	return InferHTTPPort(s.ServerPort())
}

func InferHTTPPort(port int) int {
	return port
}
