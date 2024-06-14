package netstr

import (
	"bytes"
	"encoding/binary"
	"math/rand"
	"net"
	"net/netip"
	"time"

	"github.com/noxworld-dev/opennox-lib/log"
	"github.com/noxworld-dev/opennox-lib/noxnet"
	"github.com/noxworld-dev/opennox-lib/platform"

	"github.com/noxworld-dev/opennox/v1/common/ntype"
	"github.com/noxworld-dev/opennox/v1/server/netlib"
)

const (
	maxStructs   = 128
	maskID       = byte(maxStructs - 1) // 0x7F
	reliableFlag = byte(maxStructs)     // 0x80
	anyID        = byte(0xff)
)

const (
	codeInit0          = noxnet.Op(0x00)
	codeNewStream1     = noxnet.Op(0x01)
	codeErr2           = noxnet.Op(0x02)
	code3              = noxnet.Op(0x03)
	codeAlreadyJoined4 = noxnet.Op(0x04)
	code5              = noxnet.Op(0x05)
	code6              = noxnet.Op(0x06)
	code7              = noxnet.Op(0x07)
	code8              = noxnet.Op(0x08)
	code9              = noxnet.Op(0x09)
	code10             = noxnet.Op(0x0A)
	code11             = noxnet.Op(0x0B)
	codeDiscover12     = noxnet.Op(0x0C)
	codeInfo13         = noxnet.Op(0x0D)
	codeJoin14         = noxnet.Op(0x0E)
	codeErr15          = noxnet.Op(0x0F)
	codePing16         = noxnet.Op(0x10)
	code17             = noxnet.Op(0x11)
	codePong18         = noxnet.Op(0x12)
	codeErrCode19      = noxnet.Op(0x13)
	codeACK20          = noxnet.Op(0x14)
	codeErr21          = noxnet.Op(0x15)
	code31             = noxnet.Op(0x1F)
	code32             = noxnet.Op(0x20)
	code33             = noxnet.Op(0x21)
	code34             = noxnet.Op(0x22)
	code35             = noxnet.Op(0x23)
	code36             = noxnet.Op(0x24)
	code37             = noxnet.Op(0x25)
	code38             = noxnet.Op(0x26)
)

var (
	Log = log.New("network")
)

func NewStreams(frame func() uint32) *Streams {
	s := &Streams{
		playerIDs: make(map[int]struct{}),
		KeyRand: func(min, max int) int {
			return min + rand.Intn(max-min)
		},
		PacketDropRand: func(min, max int) int {
			return min + rand.Intn(max-min)
		},
		IsHost: func() bool {
			return false
		},
		IsFlag4: func() bool {
			return false
		},
		GameFrame:     frame,
		Now:           platform.Ticks,
		Sleep:         platform.Sleep,
		Debug:         DebugSockets,
		Log:           Log,
		Xor:           true,
		MaxPacketLoss: 3,
	}
	s.reset()
	return s
}

type Streams struct {
	lastSendReliable time.Duration
	now              time.Duration
	cntX             int
	lis              Server
	streams          [maxStructs]*Conn
	streams2         [maxStructs]stream2
	timing           [maxStructs]timingStruct
	transfer         [maxStructs]uint32
	playerIDs        map[int]struct{}
	sendXorBuf       [4096]byte // TODO: remove this buffer?
	Log              *log.Logger
	Now              func() time.Duration
	Sleep            func(dt time.Duration)
	IsHost           func() bool
	IsFlag4          func() bool
	GameFrame        func() uint32
	GetMaxPlayers    func() int
	OnDiscover       func(src, dst []byte) int
	OnExtPacket      func(pc net.PacketConn, buf1 []byte, from netip.AddrPort)
	KeyRand          func(min, max int) int
	PacketDropRand   func(min, max int) int
	PacketDrop       int
	MaxPacketLoss    int
	Responded        bool
	Xor              bool
	Debug            bool
}

func (g *Streams) reset() {
	for i, ns := range g.streams {
		if ns != nil && ns.pc != nil {
			_ = ns.Close()
		}
		g.streams[i] = nil
	}
	for i := range g.streams2 {
		g.streams2[i] = stream2{}
	}
}

func (g *Streams) ConnByPlayerInd(pind ntype.PlayerInd) *Conn {
	return g.streams[int(pind)+1]
}

func (g *Streams) getForPacket(i int) handle {
	return handle{g, i}
}

func (g *Streams) GetTimingByInd1(ind ntype.PlayerInd) int {
	return int(g.timing[int(ind)+1].field28)
}

func (g *Streams) ByPlayer(p ntype.Player) *Conn {
	return g.ConnByPlayerInd(p.PlayerIndex())
}

func (g *Streams) StreamByPlayerInd(pind ntype.PlayerInd) netlib.Stream {
	return g.ConnByPlayerInd(pind)
}

func (g *Streams) HostStream() netlib.Stream {
	return g.Host()
}

func (g *Streams) Host() *Conn {
	return g.streams[0]
}

func (g *Streams) Update() {
	g.maybeSendCode11()
	g.MaybeSendReliable()
}

type timingStruct struct {
	field0  uint32
	field4  uint32
	field8  [5]uint32
	field28 uint32
}

type stream2 struct {
	active  bool
	lost    uint8
	addr    netip.AddrPort
	cur     uint8 // index into samples
	samples [10]time.Duration
	ticks   time.Duration
}

func (nx *stream2) nextPingPacket(t time.Duration) []byte {
	nx.ticks = t

	var buf [8]byte
	buf[2] = byte(codePing16)
	buf[3] = nx.cur
	binary.LittleEndian.PutUint32(buf[4:], uint32(nx.ticks/time.Millisecond))
	return buf[:]
}

func (g *Streams) getFreeIndex() (int, bool) {
	for i, it := range g.streams {
		if it == nil {
			return i, true
		}
	}
	return -1, false
}

func (g *Streams) getFreeIndexWithAddr(addr netip.AddrPort) (int, bool) {
	for i, it := range g.streams {
		if it == nil {
			return i, true
		} else if it.addr == addr {
			return i, false
		}
	}
	return -1, false
}

func (g *Streams) getFreeNetStruct2Ind() int {
	for i := range g.streams2 {
		if !g.streams2[i].active {
			return i
		}
	}
	return -1
}

func (g *Streams) connByAddr(addr netip.AddrPort) *Conn {
	for i := range g.streams {
		ns := g.streams[i]
		if ns == nil {
			continue
		}
		if addr == ns.addr {
			return ns
		}
	}
	return nil
}

func (g *Streams) struct2IndByAddr(addr netip.AddrPort) int {
	for i := range g.streams2 {
		nx := &g.streams2[i]
		if addr == nx.addr {
			return i
		}
	}
	return -1
}

func (g *Streams) hasConnWithIndAndAddr(ind handle, addr netip.AddrPort) bool {
	for i := range g.streams {
		ns := g.streams[i]
		if ns == nil {
			continue
		}
		if addr == ns.addr && ind.i == i {
			return true
		}
	}
	return false
}

type Handler func(conn netlib.StreamID, buf []byte) int

type Options struct {
	Port       int
	Max        int
	BufferSize int
	// SendPoll polls additional data for sending.
	// Caller must copy the data to provided buffer and return its size.
	SendPoll  Handler
	OnReceive Handler
	OnJoin    JoinCheck
	Check17   Check17
}

func (g *Streams) newStream(id int, opt *Options) *Conn {
	ind, ok := g.getFreeIndex()
	if !ok {
		return nil
	}
	return g.newStreamAt(ind, id, opt)
}

func (g *Streams) newStreamAt(ind int, id int, opt *Options) *Conn {
	ns := &Conn{
		g:         g,
		ind:       ind,
		id:        id,
		max:       uint32(opt.Max),
		seqRecv:   -1,
		xorKey:    0,
		sendPoll:  opt.SendPoll,
		onReceive: opt.OnReceive,
		onJoin:    opt.OnJoin,
		check17:   opt.Check17,
	}
	ns.resetReliable()
	if dsz := opt.BufferSize; dsz > 0 {
		dsz -= dsz % 4
		opt.BufferSize = dsz
	} else {
		opt.BufferSize = 1024
	}
	ns.recv.Grow(2 + opt.BufferSize)

	ns.sendBuf = make([]byte, 2+opt.BufferSize)
	ns.sendWrite = 2
	*ns.data2hdr() = [2]byte{anyID, 0}
	g.streams[ind] = ns
	return ns
}

func (g *Streams) maybeSendCode11() {
	g.now = g.Now()
	for _, ns := range g.streams {
		if ns == nil {
			continue
		}
		if ns.field38 == 1 && ns.frame40+300 < g.GameFrame() {
			ns.SendCode11()
		}
	}
}

func (g *Streams) processStreamOp9(pid handle, packetCur []byte) int {
	if len(packetCur) < 4 {
		return 0
	}
	v := binary.LittleEndian.Uint32(packetCur[:4])
	p := &g.timing[pid.i]
	dv := v - p.field4
	if dv <= 0 || dv >= 1000 {
		return 0
	}
	ind := p.field0
	p.field8[ind] = dv
	v26 := (ind + 1) % 5
	v27 := v26
	if v26 == 0 {
		for _, vv := range p.field8 {
			v26 += vv
		}
		p.field28 = v26 / uint32(len(p.field8))
	}
	p.field0 = v27
	return 0
}

func netCrypt(key byte, p []byte) {
	if len(p) == 0 {
		return
	}
	for i := range p {
		p[i] ^= key
	}
}

func (g *Streams) recvRoot(dst *bytes.Buffer, pc net.PacketConn) (int, netip.AddrPort) {
	const bufsize = 16384
	var buf [bufsize]byte // TODO: get rid of this buffer
	n, src, err := g.recvRaw(pc, buf[:])
	if err == nil && g.Xor {
		if ns := g.connByAddr(src); ns != nil && ns.xorKey != 0 {
			netCrypt(ns.xorKey, buf[:n])
		}
	}
	if !g.IsHost() && g.PacketDrop > 0 {
		if r := g.PacketDropRand(1, 99); r < g.PacketDrop {
			return 0, src
		}
	}
	dst.Write(buf[:n])
	return n, src
}

func (g *Streams) recvRaw(pc net.PacketConn, buf []byte) (int, netip.AddrPort, error) {
	n, src, err := readFrom(g.Debug, g.Log, pc, buf)
	if err != nil {
		return n, src, err
	}
	if n >= 2 && binary.LittleEndian.Uint16(buf[:2]) == 0xF13A { // extension packet code
		g.OnExtPacket(pc, buf, src)
		return 0, src, nil
	}
	return n, src, nil
}

func (g *Streams) processPong(out []byte, packet []byte, from netip.AddrPort) int {
	if len(packet) < 8 {
		return 0
	}
	token := binary.LittleEndian.Uint32(packet[4:])
	dt := g.Now() - time.Duration(token)*time.Millisecond
	ind := g.struct2IndByAddr(from)
	if ind < 0 {
		return 0
	}
	nx := &g.streams2[ind]
	if packet[3] != nx.cur {
		return 0
	}
	nx.samples[nx.cur] = dt
	nx.cur++
	if int(nx.cur) >= len(nx.samples) {
		return 0
	}
	// TODO: this overwrites first 2 bytes
	t := g.Now()
	return copy(out, nx.nextPingPacket(t))
}

func (g *Streams) sendTime(ind2 int) (int, error) {
	ns2 := &g.streams2[ind2]
	buf := ns2.nextPingPacket(g.Now())
	return g.lis.writeTo(buf, ns2.addr)
}

func (g *Streams) sendCode20(ind2 int) (int, error) {
	var buf [3]byte
	buf[0] = 0
	buf[1] = 0
	buf[2] = byte(codeACK20)

	nx := &g.streams2[ind2]
	nx.active = false

	return g.lis.writeTo(buf[:3], nx.addr)
}

func (g *Streams) sendCode19(code byte, ind2 int) (int, error) {
	var buf [4]byte
	buf[0] = 0
	buf[1] = 0
	buf[2] = byte(codeErrCode19)
	buf[3] = code

	nx := &g.streams2[ind2]
	nx.active = false

	return g.lis.writeTo(buf[:4], nx.addr)
}

func (g *Streams) ProcessStats(min, max time.Duration) {
	start := g.Now()
	for i := range g.streams2 {
		nx := &g.streams2[i]
		if !nx.active {
			continue
		}
		v2 := nx.cur
		if int(v2) >= len(nx.samples) {
			if int(nx.lost) > g.MaxPacketLoss {
				_, _ = g.sendCode19(1, i)
				_, _ = g.sendCode20(i)
				continue
			}
			cnt := 0
			var sum time.Duration
			for i = 0; i < 10; i++ {
				if nx.samples[i] > 0 {
					cnt++
					sum += nx.samples[i]
				}
			}
			avg := sum / time.Duration(cnt)
			if min != -1 && avg < min {
				_, _ = g.sendCode19(0, i)
			}
			if max != -1 && avg > max {
				_, _ = g.sendCode19(1, i)
			}
			_, _ = g.sendCode20(i)
		} else {
			if start-nx.ticks > 2*time.Second {
				g.streams2[i].samples[v2] = -1
				nx.lost++
				if int(nx.lost) <= g.MaxPacketLoss {
					nx.cur++
					_, _ = g.sendTime(i)
				} else {
					_, _ = g.sendCode19(1, i)
				}
			}
		}
	}
}
