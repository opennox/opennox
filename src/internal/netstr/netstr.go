package netstr

import (
	"encoding"
	"encoding/binary"
	"errors"
	"fmt"
	"math/rand"
	"net"
	"net/netip"
	"sync"
	"time"
	"unsafe"

	"github.com/noxworld-dev/opennox-lib/common"
	"github.com/noxworld-dev/opennox-lib/env"
	"github.com/noxworld-dev/opennox-lib/log"
	"github.com/noxworld-dev/opennox-lib/noxnet"
	"github.com/noxworld-dev/opennox-lib/platform"

	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/ntype"
	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc"
)

const (
	maxStructs = 128
)

var (
	Log    = log.New("network")
	Global = NewStreams()
)

func NewStreams() *Streams {
	return &Streams{
		playerIDs: make(map[Index]struct{}),
		KeyRand: func(min, max int) int {
			return min + rand.Intn(max-min)
		},
		PacketDropRand: func(min, max int) int {
			return min + rand.Intn(max-min)
		},
		Xor:           true,
		MaxPacketLoss: 3,
	}
}

type Streams struct {
	once           sync.Once
	ticks1         uint64
	ticks2         uint64
	cntX           int
	initIndex      Index
	streams        [maxStructs]*stream
	streams2       [maxStructs]stream2
	timing         [maxStructs]timingStruct
	transfer       [maxStructs]uint32
	allocQueue     alloc.ClassT[queueItem]
	playerIDs      map[Index]struct{}
	GameFrame      func() uint32
	GetMaxPlayers  func() int
	OnDiscover     func(src, dst []byte) int
	OnExtPacket    func(pc net.PacketConn, buf1 []byte, from netip.AddrPort)
	KeyRand        func(min, max int) int
	PacketDropRand func(min, max int) int
	PacketDrop     int
	MaxPacketLoss  int
	Flag1          bool
	Xor            bool
	Debug          bool
}

func (g *Streams) Init() {
	g.once.Do(func() {
		for i := range g.streams {
			g.streams[i] = nil
		}
		for i := range g.streams2 {
			g.streams2[i] = stream2{}
		}
		g.allocQueue = alloc.NewClassT("GQueue", queueItem{}, 200)
	})
}

func (g *Streams) addTransferStats(n int, ind Index) {
	g.transfer[ind.i] += uint32(n)
}

func (g *Streams) getTransferStats(v Index) uint32 {
	val := g.transfer[v.i]
	g.transfer[v.i] = 0
	return val
}

func (v Index) TransferStats() uint32 {
	if !v.Valid() {
		return 0
	}
	return v.g.getTransferStats(v)
}

func (g *Streams) Update() {
	g.sub5524C0()
	g.MaybeSendAll()
}

type timingStruct struct {
	field0  uint32
	field4  uint32
	field8  [5]uint32
	field28 uint32
}

type queueItem struct {
	next   *queueItem
	ticks  uint64
	active bool
	size   uint32
	data   [1024]byte
}

type stream2 struct {
	active bool
	lost   uint8
	addr   netip.AddrPort
	cur    uint8 // index into arr
	arr    [10]int
	ticks  uint64
}

func (g *Streams) getFreeNetStructInd() (Index, bool) {
	for i := range g.streams {
		if g.streams[i] == nil {
			return Index{g, i}, true
		}
	}
	return Index{nil, -1}, false
}

func (g *Streams) getFreeNetStruct2Ind() int {
	for i := range g.streams2 {
		if !g.streams2[i].active {
			return i
		}
	}
	return -1
}

func (g *Streams) structByAddr(addr netip.AddrPort) *stream {
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

func (g *Streams) hasStructWithIndAndAddr(ind Index, addr netip.AddrPort) bool {
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

type Handler func(id Index, buf []byte, a3 unsafe.Pointer) int

type Options struct {
	Port      int
	Max       int
	DataSize  int
	Data3Size int
	Func1     Handler
	Func2     Handler
	Check14   Check14
	Check17   Check17
}

func (g *Streams) newStruct(opt *Options) *stream {
	ns := &stream{g: g}

	if opt.Data3Size > 0 {
		ns.data3, _ = alloc.Malloc(uintptr(opt.Data3Size))
	}
	if dsz := opt.DataSize; dsz > 0 {
		dsz -= dsz % 4
		opt.DataSize = dsz
	} else {
		opt.DataSize = 1024
	}
	data1, _ := alloc.Make([]byte{}, opt.DataSize+2)
	ns.data1 = data1
	ns.data1xxx = 0
	ns.data1yyy = 0

	data2, _ := alloc.Make([]byte{}, opt.DataSize+2)
	data2[0] = 0xff
	ns.data2 = data2
	ns.data2xxx = 2
	ns.data2yyy = 2

	ns.max = uint32(opt.Max)
	ns.func1 = opt.Func1
	ns.func2 = opt.Func2
	ns.ind28 = -1
	ns.xorKey = 0

	ns.check14 = opt.Check14
	ns.check17 = opt.Check17
	return ns
}

func (ns *stream) Close() error {
	if ns == nil {
		return nil
	}
	_ = ns.pc.Close()
	ns.pc = nil
	ns.freeData()
	return nil
}

func (v Index) Close() {
	if ns := v.get(); ns != nil {
		_ = ns.Close()
		v.g.streams[v.i] = nil
	}
}

func (g *Streams) Player(p ntype.Player) Index {
	return g.PlayerInd(p.PlayerIndex())
}

func (g *Streams) PlayerInd(pind ntype.PlayerInd) Index {
	return Index{g, int(pind) + 1}
}

func (g *Streams) IndexRaw(ind int) Index {
	return Index{g, ind}
}

func (g *Streams) First() Index {
	return Index{g, 0}
}

type Index struct {
	g *Streams
	i int // hiding it in a struct helps prevent direct casts
}

func (v Index) Raw() int {
	return v.i
}

func (v Index) Valid() bool {
	return v.g != nil && v.i >= 0 && v.i < maxStructs
}

func (v Index) IsFirst() bool {
	return v.i == 0
}

func (v Index) Player() ntype.PlayerInd {
	return ntype.PlayerInd(v.i - 1)
}

func (v Index) get() *stream {
	if !v.Valid() {
		return nil
	}
	return v.g.streams[v.i]
}

func (g *Streams) NewClient(narg *Options) (Index, error) {
	if narg == nil {
		return Index{nil, -2}, NewConnectFailErr(-2, errors.New("empty options"))
	}
	ind, ok := g.getFreeNetStructInd()
	if !ok {
		return Index{nil, -8}, NewConnectFailErr(-8, errors.New("no more slots for net structs"))
	}
	ns := g.newStruct(narg)
	g.streams[ind.i] = ns
	return ind, nil
}

var optionsBuffer [1024]byte

func (g *Streams) Dial(ind Index, host string, port int, cport int, opts encoding.BinaryMarshaler) error {
	if g.Debug {
		Log.Println("NET_CONNECT", ind, host, port)
	}
	ns := ind.get()
	if ns == nil {
		return NewConnectFailErr(-3, errors.New("no net struct"))
	}
	if len(host) == 0 {
		return NewConnectFailErr(-4, errors.New("empty hostname"))
	}
	if port < 1024 || port > 0x10000 {
		return NewConnectFailErr(-15, errors.New("invalid port"))
	}
	var ip netip.Addr
	if host[0] < '0' || host[0] > '9' {
		list, err := net.LookupIP(host)
		if err != nil || len(list) == 0 {
			log.Printf("error: cannot find ip for a host %q: %v", host, err)

			return NewConnectFailErr(-4, err)
		}
		ip, _ = netip.AddrFromSlice(list[0].To4())
	} else {
		var err error
		ip, err = netip.ParseAddr(host)
		if err != nil {
			log.Printf("error: cannot parse host %q: %v", host, err)
			return NewConnectFailErr(-4, err)
		}
	}
	addr := netip.AddrPortFrom(ip, uint16(port))
	ns.SetAddr(addr)

	for {
		sock, err := Listen(netip.AddrPortFrom(netip.IPv4Unspecified(), uint16(cport)))
		if err == nil {
			ns.pc = sock
			break
		} else if !ErrIsInUse(err) {
			return NewConnectFailErr(-1, err)
		}
		cport++
	}
	g.Flag1 = false
	var v12 [1]byte // TODO: sending zero, is that correct? if so, set explicitly here
	v11, err := g.Send(ind, v12[:], SendNoLock|SendFlagFlush)
	if err != nil {
		return fmt.Errorf("cannot send data: %w", err)
	}

	id, retries, flags, counter := ind, 60, 6, 0
	for {
		ns := id.get()
		if ns == nil {
			return NewConnectFailErr(-23, errors.New("no net struct"))
		}
		//counter = 0 // TODO: is this correct?
		counter++
		if counter > 20*retries {
			return NewConnectFailErr(-23, errors.New("timeout"))
		}
		g.ServeInitialPackets(id, flags|1)
		g.MaybeSendAll()
		f28 := int(ns.ind28)
		if f28 >= v11 {
			break
		}
		platform.Sleep(30 * time.Millisecond)
	}

	ns = id.get()
	if g.Flag1 && ns.ID().Valid() {
		vs := optionsBuffer[:]
		copy(vs, make([]byte, len(vs)))
		data, err := opts.MarshalBinary()
		if err != nil {
			return err
		}
		vs = vs[:3+len(data)]

		vs[0] = byte(noxnet.MSG_ACCEPTED)
		vs[1] = ns.Data2hdr()[1]
		vs[2] = 32
		if len(data) > 0 {
			copy(vs[3:], data[:153])
		}
		_, _ = g.Send(id, vs, SendNoLock|SendFlagFlush)
	}

	if !ns.ID().Valid() {
		return NewConnectFailErr(-1, errors.New("invalid net struct id"))
	}
	return nil
}

func (g *Streams) DialWait(ind Index, timeout time.Duration, send func(), check func() bool) error {
	deadline := platform.Ticks() + timeout
	if g.Debug {
		Log.Println("CONNECT_WAIT_LOOP_START", deadline)
	}
	for {
		now := platform.Ticks()
		if now >= deadline {
			return NewConnectFailErr(-19, errors.New("timeout 2"))
		}
		g.ServeInitialPackets(ind, 1)
		send()
		g.MaybeSendAll()
		if check() {
			break
		}
		platform.Sleep(5 * time.Millisecond)
	}
	return nil
}

type stream struct {
	g    *Streams
	pc   net.PacketConn
	addr netip.AddrPort
	id   Index

	data1    []byte
	data1xxx int
	data1yyy int

	data2    []byte
	data2xxx int
	data2yyy int

	max         uint32
	playerInd21 uint32
	ticks22     uint64
	ticks23     uint64
	field24     uint32
	cnt28       uint8
	ind28       int8
	queue       *queueItem
	data3       unsafe.Pointer
	func1       Handler
	func2       Handler
	xorKey      byte
	field38     byte
	field39     byte
	frame40     uint32
	check14     Check14
	check17     Check17
}

func (ns *stream) freeData() {
	if ns == nil {
		return
	}
	if ns.data3 != nil {
		alloc.Free(ns.data3)
	}
	alloc.FreeSlice(ns.data1)
	alloc.FreeSlice(ns.data2)
	*ns = stream{g: ns.g}
}

func (ns *stream) Addr() netip.AddrPort {
	if ns == nil {
		return netip.AddrPort{}
	}
	return ns.addr
}

func (ns *stream) SetAddr(addr netip.AddrPort) {
	ns.addr = addr
}

func (ns *stream) ID() Index {
	if ns == nil {
		return Index{nil, -1}
	}
	return ns.id
}

func (ns *stream) Data2hdr() *[2]byte {
	if ns == nil {
		var v [2]byte
		return &v
	}
	return (*[2]byte)(ns.data2[:2])
}

func (ns *stream) Datax2() []byte {
	if ns == nil {
		return nil
	}
	return ns.data2[:ns.data2xxx]
}

func (ns *stream) Data1xxx() []byte {
	if ns == nil {
		return nil
	}
	return ns.data1[ns.data1xxx:]
}

func (ns *stream) Data1yyy() []byte {
	if ns == nil {
		return nil
	}
	return ns.data1[ns.data1yyy:ns.data1xxx]
}

func (ns *stream) Data2xxx() []byte {
	if ns == nil {
		return nil
	}
	return ns.data2[ns.data2xxx:]
}

func (ns *stream) callFunc1(id Index, buf []byte, data3 unsafe.Pointer) int {
	if ns.func1 == nil {
		return 0
	}
	return ns.func1(id, buf, data3)
}

func (ns *stream) callFunc2(id Index, buf []byte, data3 unsafe.Pointer) int {
	if ns.func2 == nil {
		return 0
	}
	return ns.func2(id, buf, data3)
}

const (
	SendNoLock    = 0x1
	SendFlagFlush = 0x2
)

func (g *Streams) Send(id Index, buf []byte, flags int) (int, error) {
	ns := id.get()
	if ns == nil {
		return -3, errors.New("no net struct")
	}
	if len(buf) == 0 {
		return -2, errors.New("empty buffer")
	}
	var (
		idd    Index
		ei, si Index
	)
	if ns.id.i == -1 {
		ei = Index{g, maxStructs}
		si = Index{g, 0}
		idd = id
	} else {
		si = id
		ei = Index{g, id.i + 1}
		idd = ns.ID()
	}
	if flags&SendNoLock != 0 {
		n := len(buf)
		for i := si.i; i < ei.i; i++ {
			ns2 := Index{g, i}.get()
			if ns2 != nil && ns2.ID() == idd {
				v12, err := ns2.addToQueue(buf)
				if v12 == -1 {
					return -1, err
				}
				n = v12
				if flags&SendFlagFlush != 0 {
					ns2.maybeSendQueue(byte(v12), true)
				}
			}
		}
		return n, nil
	}
	n := len(buf)
	for i := si.i; i < ei.i; i++ {
		ns2 := Index{g, i}.get()
		if ns2 == nil {
			continue
		}
		if ns2.ID() != idd {
			continue
		}
		d2x := ns2.Data2xxx()
		if n+1 > len(d2x) {
			return -7, errors.New("buffer too short")
		}
		if flags&SendFlagFlush != 0 {
			copy(d2x[:2], ns2.Data2hdr()[:2])
			copy(d2x[2:2+n], buf)
			n2, err := ns2.WriteTo(d2x[:n+2], ns2.addr)
			if n2 == -1 {
				return -1, err
			}
			g.addTransferStats(n+2, Index{g, i})
			return n2, nil
		}
		copy(d2x[:n], buf)
		ns2.data2xxx += n
	}
	return n, nil
}

func (g *Streams) netCrypt(key byte, p []byte) {
	if len(p) == 0 || !g.Xor {
		return
	}
	for i := range p {
		p[i] ^= key
	}
}

func (ns *stream) addToQueue(buf []byte) (int, error) {
	if ns == nil {
		return -3, errors.New("no net struct")
	}
	if len(buf) > 1024-2 {
		return -1, errors.New("buffer too large")
	}
	if len(buf) == 0 {
		return -1, errors.New("empty buffer")
	}
	q := ns.g.allocQueue.NewObject()
	if q == nil {
		return -1, errors.New("cannot alloc gqueue")
	}
	q.next = ns.queue
	ns.queue = q

	q.active = true
	q.size = uint32(len(buf) + 2)
	q.data[0] = ns.Data2hdr()[0] | 0x80
	q.data[1] = ns.cnt28
	ns.cnt28++
	copy(q.data[2:], buf)
	return int(q.data[1]), nil
}

func (ns *stream) maybeSendQueue(hdrByte byte, checkHdr bool) int {
	if ns == nil {
		return -3
	}
	for it := ns.queue; it != nil; it = it.next {
		if checkHdr {
			if it.data[1] == hdrByte {
				sz := int(it.size)
				it.active = false
				it.ticks = ns.g.ticks2 + 2000
				if _, err := ns.WriteTo(it.data[:sz], ns.addr); err != nil {
					Log.Println(err)
					return 0
				}
			}
		} else {
			if it.active {
				sz := int(it.size)
				it.active = false
				it.ticks = ns.g.ticks2 + 2000
				if _, err := ns.WriteTo(it.data[:sz], ns.addr); err != nil {
					Log.Println(err)
					return 0
				}
			}
		}
	}
	return 0
}

func (ns *stream) setActiveInQueue(hdrByte byte, checkHdr bool) int {
	if ns == nil {
		return -3
	}
	for it := ns.queue; it != nil; it = it.next {
		if checkHdr {
			if it.data[1] == hdrByte {
				it.active = true
				continue
			}
		} else {
			if it.ticks <= ns.g.ticks2 {
				it.active = true
				continue
			}
		}
	}
	return 0
}

func (g *Streams) GetInitInd() Index {
	return g.initIndex
}

func (g *Streams) InitNew(narg *Options) (Index, error) {
	if narg == nil {
		return Index{nil, -2}, errors.New("empty options")
	}
	if narg.Max > maxStructs {
		return Index{nil, -2}, errors.New("max limit reached")
	}
	ind, ok := g.getFreeNetStructInd()
	if !ok {
		return Index{nil, -8}, errors.New("no more slots for net structs")
	}
	ns := g.newStruct(narg)
	g.streams[ind.i] = ns
	ns.Data2hdr()[0] = byte(ind.i)
	ns.id = Index{g, -1}

	if narg.Port < 1024 || narg.Port > 0x10000 {
		narg.Port = common.GamePort
	}

	for {
		sock, err := Listen(netip.AddrPortFrom(netip.IPv4Unspecified(), uint16(narg.Port)))
		if err == nil {
			ns.pc = sock
			break
		} else if !ErrIsInUse(err) {
			return Index{nil, -5}, err
		}
		narg.Port++
	}
	g.initIndex = ind
	return ind, nil
}

func (g *Streams) SendReadPacket(ind Index, flags byte) int {
	ns := ind.get()
	if ns == nil {
		return -3
	}
	find := ind
	var min, max Index
	if ns.id.i != -1 {
		min = ind
		max = Index{g, ind.i + 1}
		find = ns.id
	} else {
		min = Index{g, 0}
		max = Index{g, maxStructs}
	}
	for j := min.i; j < max.i; j++ {
		ns2 := g.streams[j]
		if ns2 == nil || ns2.id != find {
			continue
		}
		ns2.maybeSendQueue(0, false)
		if flags&1 == 0 {
			data2 := ns2.Data2xxx()
			n := ns2.callFunc1(Index{g, j}, data2, ns2.data3)
			if n > 0 && n <= len(data2) {
				ns2.data2xxx += n
			}
		}
		datax2 := ns2.Datax2()
		if len(datax2) > 2 {
			_, err := ns.WriteTo(datax2, ns2.addr)
			if err != nil {
				return -1
			}
			g.addTransferStats(len(datax2), Index{g, j})
			ns2.data2xxx = ns2.data2yyy
		}
	}
	return 0
}

func (ns *stream) maybeFreeQueue(a2 byte, a3 int) int {
	if ns == nil {
		return -3
	}
	var (
		next *queueItem
		prev *queueItem
	)
	for it := ns.queue; it.next != nil; it = next {
		next = it.next
		if a3 == 0 {
			if it.data[1] != a2 {
				prev = it
				continue
			}
		} else if a3 == 1 {
			if a2 < 0x20 || a2 > 0xE0 {
				if it.data[1] >= a2 {
					prev = it
					continue
				}
			} else {
				if it.data[1] >= a2 {
					prev = it
					continue
				}
			}
		} else if a3 != 2 {
			prev = it
			continue
		}
		if head := ns.queue; it == head {
			ns.queue = head.next
		}
		if prev != nil {
			prev.next = next
		}
		it.next = nil
		ns.g.allocQueue.FreeObjectFirst(it)
	}
	return 0
}

func (g *Streams) ReadPackets(ind Index) int {
	if !ind.Valid() {
		return -3
	}
	ns := ind.get()
	if ns == nil {
		return 0
	}
	v4 := ns.ID()
	var si, ei Index
	if v4.i == -1 {
		si, ei = Index{g, 0}, Index{g, maxStructs}
		v4 = ind
	} else {
		si, ei = ind, Index{g, ind.i + 1}
		ns2 := v4.get()
		if ns2 == nil || ns2.id.i != -1 {
			ns.freeData()
			g.streams[ind.i] = nil
			return 0
		}
	}
	for i := si.i; i < ei.i; i++ {
		ii := Index{g, i}
		ns2 := ii.get()
		if ns2 == nil || ns2.ID() != v4 {
			continue
		}
		g.SendReadPacket(ii, 1)
		var buf [1]byte
		buf[0] = 11
		_, _ = g.Send(ii, buf[:], 0)
		g.SendReadPacket(ii, 1)
		v4.get().playerInd21--
		ns.maybeFreeQueue(0, 2)
		ns2.freeData()
		g.streams[ii.i] = nil
	}
	return 0
}

func platformTicks() uint64 {
	if env.IsE2E() {
		return uint64(platform.Ticks())
	}
	return uint64(platform.Ticks() / time.Millisecond)
}

func (g *Streams) sub5524C0() {
	g.ticks2 = platformTicks()
	for i, ns := range g.streams {
		if ns != nil && ns.field38 == 1 {
			if ns.frame40+300 < g.GameFrame() {
				g.ReadPackets(Index{g, i})
			}
		}
	}
}

func (g *Streams) MaybeSendAll() {
	now := platformTicks()
	g.ticks2 = now
	if now-g.ticks1 <= 1000 {
		return
	}
	for _, ns := range g.streams {
		if ns != nil {
			ns.setActiveInQueue(0, false)
			ns.maybeSendQueue(0, false)
		}
	}
	g.ticks1 = now
}

func (g *Streams) GetTimingByInd1(ind int) int {
	return int(g.timing[1+ind].field28)
}

func (v Index) IP() netip.Addr {
	ns := v.get()
	if ns == nil {
		return netip.Addr{}
	}
	return ns.addr.Addr()
}

func (g *Streams) SendClose(ind Index) {
	if ind.get() == nil {
		return
	}
	g.SendReadPacket(ind, 0)
	var buf [1]byte
	buf[0] = 10
	_, _ = g.Send(ind, buf[:1], 0)
	g.SendReadPacket(ind, 0)
	ind.Close()
}

func (g *Streams) processStreamOp0(id Index, out []byte, pid Index, p1 byte, ns1 *stream, from netip.AddrPort) int {
	if noxflags.HasGame(noxflags.GameHost) && noxflags.HasGame(noxflags.GameFlag4) {
		return 0
	}
	out[0] = 0
	out[1] = p1
	if int(ns1.playerInd21) >= g.GetMaxPlayers()+(g.cntX-1) {
		out[2] = 2
		return 3
	}
	if pid.i != -1 {
		Log.Printf("pid must be set to -1 when joining: was %d (%s)\n", pid, from.String())
		// pid in the request must be -1 (0xff); fail if it's not
		out[2] = 2
		return 3
	}
	// now, find free net struct index and use it as pid
	for i, it := range g.streams {
		if it == nil {
			pid = Index{g, i}
			break
		}
		if from == it.addr {
			out[2] = 4 // already joined?
			return 3
		}
	}
	if pid.i == -1 {
		out[2] = 2
		return 3
	}
	ns2 := g.newStruct(&Options{
		Data3Size: 4,
		DataSize:  len(ns1.data2),
		Check14:   ns1.check14,
		Check17:   ns1.check17,
	})
	g.streams[pid.i] = ns2
	ns1.playerInd21++

	hdr := ns2.Data2hdr()
	hdr[0] = byte(id.i)
	if hdr[1] == p1 {
		hdr[1]++
	}
	ns2.id = id
	ns2.pc = ns1.pc
	ns2.func1 = ns1.func1
	ns2.func2 = ns1.func2

	g.timing[id.i] = timingStruct{field28: 1}
	key := byte(g.KeyRand(1, 255))
	if !g.Xor {
		key = 0
	}
	ns2.xorKey = 0 // send this packet without xor encoding

	ns2.SetAddr(from)

	out[0] = byte(noxnet.MSG_ACCEPTED)
	out[1] = p1
	out[2] = 1
	binary.LittleEndian.PutUint32(out[3:], uint32(pid.i))
	out[7] = key
	v67, _ := g.Send(pid, out[:8], SendNoLock|SendFlagFlush)

	ns2.xorKey = key
	ns2.field38 = 1
	ns2.field39 = byte(v67)
	ns2.frame40 = g.GameFrame()
	return 0
}

func (g *Streams) processStreamOp6(pid Index, out []byte, ns1 *stream, packetCur []byte) int {
	if len(packetCur) < 4 {
		return 0
	}
	v := binary.LittleEndian.Uint32(packetCur[:4])

	ns2 := pid.get()
	if ns2 == nil {
		return 0
	}
	out[0] = 37
	ns1.callFunc2(pid, out[:1], ns2.data3)

	g.timing[pid.i].field4 = v
	var v18 *[2]byte
	if ns1.id.i == -1 {
		out[0] = byte(ns2.id.i)
		v18 = ns2.Data2hdr()
	} else {
		out[0] = byte(ns1.id.i)
		v18 = ns1.Data2hdr()
	}
	out[1] = v18[0]
	out[2] = 8
	binary.LittleEndian.PutUint32(out[3:], v)
	return 7
}

func (g *Streams) processStreamOp7(pid Index, out []byte, ns1 *stream) int {
	ns4 := pid.get()
	if ns4 == nil {
		return 0
	}
	v31 := int(g.ticks2) - int(ns4.field24)
	v32 := -1
	if v31 >= 1 {
		v32 = 256000 / v31
	}
	out[0] = 35
	binary.LittleEndian.PutUint32(out[1:], uint32(v32))
	if ns1.id.i == -1 {
		ns1.callFunc2(pid, out[:5], ns4.data3)
	} else {
		ns1.callFunc2(pid, out[:5], ns1.data3)
	}
	return 0
}

func (g *Streams) processStreamOp8(pid Index, out []byte, ns1 *stream, packetCur []byte) int {
	ns5 := pid.get()
	if ns5 == nil && binary.LittleEndian.Uint32(packetCur) != uint32(ns5.ticks22) {
		return 0
	}
	ns5.field24 = uint32(int(g.ticks2) - int(ns5.ticks23))
	out[0] = 36
	binary.LittleEndian.PutUint32(out[1:], ns5.field24)
	var v19 unsafe.Pointer
	if ns1.id.i == -1 {
		v19 = ns5.data3
	} else {
		v19 = ns1.data3
	}
	ns1.callFunc2(pid, out[:5], v19)

	out[0] = ns1.Data2hdr()[0]
	out[1] = ns5.Data2hdr()[1]
	out[2] = 9
	binary.LittleEndian.PutUint32(out[3:], uint32(g.ticks2))
	return 7
}

func (g *Streams) processStreamOp9(pid Index, packetCur []byte) int {
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

type Check14 func(out []byte, packet []byte, a4a bool, add func(pid Index) bool) int

func (g *Streams) processStreamOp14(out []byte, packet []byte, ns1 *stream, p1 byte, from netip.AddrPort, check Check14) int {
	out[0] = 0
	out[1] = p1

	a4a := false
	if int(ns1.playerInd21) >= g.GetMaxPlayers()-1 {
		a4a = true
	}
	if n := check(out, packet, a4a, func(pid Index) bool {
		if _, ok := g.playerIDs[pid]; ok {
			return false
		}
		g.playerIDs[pid] = struct{}{}
		g.cntX++
		return true
	}); n != 0 {
		return n
	}

	ind2 := g.getFreeNetStruct2Ind()
	if ind2 < 0 {
		out[2] = byte(noxnet.MSG_SERVER_JOIN_OK) // OK
		return 3
	}
	nx := &g.streams2[ind2]
	nx.active = true
	nx.cur = 0
	nx.lost = 0
	nx.addr = from

	return copy(out, nx.makeTimePacket())
}

func (g *Streams) Sub552E70(ind Index) int {
	var buf [5]byte

	ns := ind.get()
	if ns == nil {
		return -3
	}
	var (
		min, max Index
		find     Index
	)
	if ns.id.i == -1 {
		min = Index{g, 0}
		max = Index{g, maxStructs}
		find = ind
	} else {
		min = ind
		max = Index{g, ind.i + 1}
		find = ns.id
	}
	buf[0] = 6
	for i := min.i; i < max.i; i++ {
		ns2 := g.streams[i]
		if ns2 != nil && ns2.id == find {
			ns2.ticks22 = g.ticks2
			ns2.ticks23 = ns2.ticks22
			binary.LittleEndian.PutUint32(buf[1:], uint32(ns2.ticks22))
			_, _ = g.Send(Index{g, i}, buf[:5], SendFlagFlush)
		}
	}
	return 0
}

func (g *Streams) readXxx(id1, id2 Index, a3 byte, buf []byte) bool {
	ns2 := id2.get()
	if ns2 == nil || ns2.field38 != 1 || ns2.field39 > a3 {
		return false
	}
	ns1 := id1.get()
	if int(ns1.playerInd21) > (g.GetMaxPlayers() - 1) {
		g.ReadPackets(id2)
		return true
	}
	if len(buf) >= 4 && buf[4] == 32 {
		ns2.field38 = 2
		ns2.field39 = 0xff
		ns2.frame40 = 0
		ns1.callFunc2(id2, buf[4:], ns2.data3)
	}
	return true
}

func (g *Streams) processStreamOp(id Index, packet []byte, out []byte, from netip.AddrPort) int {
	if len(packet) < 2 {
		return 0
	}
	pid := Index{g, int(int8(packet[0]))}
	p1 := packet[1]
	packetCur := packet[2:]

	ns1 := id.get()
	if ns1 == nil {
		return 0
	}

	pidb := pid // TODO: some of the functions assume it's different from pid, check what's wrong
	for len(packetCur) != 0 {
		op := packetCur[0]
		packetCur = packetCur[1:]
		if g.Debug {
			Log.Printf("processStreamOp: op=%d [%d]\n", op, len(packetCur))
		}
		switch noxnet.Op(op) {
		default:
			return 0
		case 0:
			return g.processStreamOp0(id, out, pidb, p1, ns1, from)
		case 1:
			if len(packetCur) < 5 {
				return 0
			}
			v11 := Index{g, int(binary.LittleEndian.Uint32(packetCur[:4]))}
			xor := packetCur[4]
			packetCur = packetCur[5:]

			ns1.id = v11
			ns1.data2[0] = byte(v11.i)
			ns1.xorKey = xor
			g.Flag1 = true
		case 2:
			ns1.id.i = -18
			g.Flag1 = true
		case 3: // ack?
			ns1.id.i = -12
			g.Flag1 = true
		case 4:
			ns1.id.i = -13
			g.Flag1 = true
		case 5:
			if len(packetCur) < 4 {
				return 0
			}
			v := binary.LittleEndian.Uint32(packetCur[:4])
			out[0] = ns1.xorKey
			out[1] = 0
			out[2] = 7
			binary.LittleEndian.PutUint32(out[3:], v)
			return 7
		case 6:
			return g.processStreamOp6(pid, out, ns1, packetCur)
		case 7:
			return g.processStreamOp7(pid, out, ns1)
		case 8:
			return g.processStreamOp8(pid, out, ns1, packetCur)
		case 9:
			return g.processStreamOp9(pid, packetCur)
		case 10:
			return g.processStreamOp10(id, pid, out, ns1)
		case 11:
			ns7 := pid.get()
			if ns7 == nil {
				return 0
			}
			out[0] = 33
			ns1.callFunc2(pid, out[:1], ns7.data3)
			id.Close()
			return 0
		case noxnet.MSG_SERVER_JOIN: // join game request?
			return g.processStreamOp14(out, packet, ns1, p1, from, ns1.check14)
		case 17:
			return g.processStreamOp17(out, packet, p1, from, ns1.check17)
		case 18:
			return g.processStreamOp18(out, packet, from)
		case noxnet.MSG_ACCEPTED:
			if len(packetCur) < 1 {
				return 0
			}
			v14 := packetCur[0]
			packetCur = packetCur[1:]

			ns8 := pidb.get()
			if ns8 == nil {
				return 0
			}
			Log.Printf("switch 31: 0x%x 0x%x\n", v14, ns8.ind28)
			if v14 != byte(ns8.ind28) {
				ns9 := pid.get()
				ns9.setActiveInQueue(v14, true)
				ns9.maybeFreeQueue(v14, 1)
				ns8.ind28 = int8(v14)
				out[0] = 38
				out[1] = v14
				ns1.callFunc2(pid, out[:2], ns8.data3)
			}
		}
	}
	return 0
}

func (g *Streams) recvRoot(pc net.PacketConn, buf []byte) (int, netip.AddrPort) {
	n, src, err := g.recvRaw(pc, buf)
	if err == nil {
		if ns := g.structByAddr(src); ns != nil && ns.xorKey != 0 {
			g.netCrypt(ns.xorKey, buf[:n])
		}
	}
	if noxflags.HasGame(noxflags.GameHost) {
		return n, src
	}
	if r := g.PacketDropRand(1, 99); r < g.PacketDrop {
		return 0, src
	}
	return n, src
}

func (g *Streams) recvRaw(pc net.PacketConn, buf []byte) (int, netip.AddrPort, error) {
	n, src, err := ReadFrom(pc, buf)
	if err != nil {
		return n, src, err
	}
	if n >= 2 && binary.LittleEndian.Uint16(buf[:2]) == 0xF13A { // extension packet code
		g.OnExtPacket(pc, buf, src)
		return 0, src, nil
	}
	return n, src, nil
}

func (g *Streams) processStreamOp10(id Index, pid Index, out []byte, ns1 *stream) int {
	if pid.i == -1 {
		return 0
	}
	ns6 := pid.get()
	if ns6 == nil || ns6.field38 == 1 {
		return 0
	}
	out[0] = 34
	ns1.callFunc2(pid, out[:1], ns6.data3)

	g.timing[id.i] = timingStruct{}

	if _, ok := g.playerIDs[pid]; ok {
		delete(g.playerIDs, pid)
		g.cntX--
	}

	g.ReadPackets(pid)
	return 0
}

type Check17 func(out []byte, packet []byte) int

func (g *Streams) processStreamOp17(out []byte, packet []byte, p1 byte, from netip.AddrPort, check Check17) int {
	out[0] = 0
	out[1] = p1
	if n := check(out, packet); n != 0 {
		return n
	}
	ind2 := g.getFreeNetStruct2Ind()
	if ind2 < 0 {
		out[2] = 20
		return 3
	}
	nx := &g.streams2[ind2]
	nx.active = true
	nx.cur = 0
	nx.lost = 0
	nx.addr = from

	return copy(out, nx.makeTimePacket())
}

func (g *Streams) processStreamOp18(out []byte, packet []byte, from netip.AddrPort) int {
	dt := int(platformTicks()) - int(binary.LittleEndian.Uint32(packet[4:]))
	ind := g.struct2IndByAddr(from)
	if ind < 0 {
		return 0
	}
	nx := &g.streams2[ind]
	if packet[3] != nx.cur {
		return 0
	}
	nx.arr[nx.cur] = dt
	nx.cur++
	if int(nx.cur) >= len(nx.arr) {
		return 0
	}
	return copy(out, nx.makeTimePacket())
}

func (g *Streams) ServeInitialPackets(id Index, flags int) int {
	ns := id.get()
	if ns == nil {
		return -3
	}
	ns2 := ns

	argp := 1
	var err error
	if flags&1 != 0 {
		argp, err = CanReadConn(ns.pc)
		if err != nil || argp == 0 {
			return -1
		}
	}
	buf, bfree := alloc.Make([]byte{}, 256)
	defer bfree()

	v26 := 1
	for {
		n, src := g.recvRoot(ns.pc, ns.Data1xxx())
		if n == -1 {
			return -1
		}
		if n < 3 {
			ns.data1xxx = 0
			ns.data1yyy = 0
			if flags&1 == 0 || flags&4 != 0 {
				return n
			}
			argp, err = CanReadConn(ns.pc)
			if err != nil {
				return -1
			} else if argp == 0 {
				return n
			}
			continue
		}
		ns.data1xxx += n
		hdr := ns.Data1yyy()[:3]
		id2 := int(hdr[0])
		id2i := Index{g, id2 & 127}
		v9 := hdr[1]
		op := noxnet.Op(hdr[2])
		if g.Debug {
			Log.Printf("servNetInitialPackets: op=%d (%s)\n", int(op), op.String())
		}
		if op == noxnet.MSG_SERVER_DISCOVER {
			n = g.OnDiscover(ns.Data1yyy(), buf)
			if n > 0 {
				n, _ = ns.WriteToRaw(buf[:n], src)
			}
			ns.data1xxx = 0
			ns.data1yyy = 0
			if flags&1 == 0 || flags&4 != 0 {
				return n
			}
			argp, err = CanReadConn(ns.pc)
			if err != nil {
				return -1
			} else if argp == 0 {
				return n
			}
			continue
		}
		if op >= noxnet.MSG_SERVER_JOIN && op <= noxnet.MSG_SERVER_JOIN_OK {
			v26 = 1
		} else {
			if id2 == 255 {
				if v26 != 1 {
					goto continueX
				}
			} else {
				v26 = 0
				if !g.hasStructWithIndAndAddr(id2i, src) {
					goto continueX
				}
				v26 = 1
			}
			if ns.id.i == -1 {
				ns2 = id2i.get()
			}
			if id2&maxStructs == 0 {
				if ns2 == nil {
					goto continueX
				}
				if v9 != byte(ns2.ind28) {
					ns9 := id2i.get()
					ns9.setActiveInQueue(v9, true)
					ns9.maybeFreeQueue(v9, 1)
					ns2.ind28 = int8(v9)
					v20 := 0
					if g.readXxx(id, id2i, v9, ns.Data1yyy()) {
						v20 = 0
					} else {
						v20 = 1
					}
					buf[0] = 38
					buf[1] = byte(ns2.ind28)
					ns.callFunc2(id2i, buf[:2], ns2.data3)
					if v20 == 0 {
						goto continueX
					}
				}
			} else if id2 == 255 {
				if op == 0 {
					data := ns.Data1yyy()
					n = g.processStreamOp(id, data, buf, src)
					if n > 0 {
						n, _ = ns.WriteTo(buf[:n], src)
					}
					goto continueX
				}
			} else {
				data := ns.Data1yyy()
				data[0] &= 127
				id2 = int(data[0])
				id2i = Index{g, id2}
				if ns2 == nil {
					goto continueX
				}
				hdr2 := ns2.Data2hdr()
				if hdr2[1] != v9 {
					goto continueX
				}
				hdr2[1]++
				if g.readXxx(id, id2i, v9, ns.Data1yyy()) {
					goto continueX
				}
			}
		}
		if op < 32 {
			data := ns.Data1yyy()
			n = g.processStreamOp(id, data, buf, src)
			if n > 0 {
				n, _ = ns.WriteTo(buf[:n], src)
			}
		} else {
			if ns2 != nil && flags&2 == 0 {
				data := ns.Data1yyy()[2:n]
				ns.callFunc2(id2i, data, ns2.data3)
			}
		}
	continueX:
		ns.data1xxx = 0
		ns.data1yyy = 0
		if flags&1 == 0 || flags&4 != 0 {
			return n
		}
		argp, err = CanReadConn(ns.pc)
		if err != nil {
			return -1
		} else if argp == 0 {
			return n
		}
	}
	// unreachable
}

func (nx *stream2) makeTimePacket() []byte {
	nx.ticks = platformTicks()

	var buf [8]byte
	buf[2] = 16
	buf[3] = nx.cur
	binary.LittleEndian.PutUint32(buf[4:], uint32(nx.ticks))
	return buf[:]
}

func (g *Streams) sendTime(ind2 int) (int, error) {
	ns := g.streams[g.initIndex.i]
	ns2 := &g.streams2[ind2]
	buf := ns2.makeTimePacket()
	return ns.WriteTo(buf, ns2.addr)
}

func (g *Streams) sendCode20(ind2 int) (int, error) {
	ns := g.streams[g.initIndex.i]
	var buf [3]byte
	buf[0] = 0
	buf[1] = 0
	buf[2] = 20

	nx := &g.streams2[ind2]
	nx.active = false

	return ns.WriteTo(buf[:3], nx.addr)
}

func (g *Streams) sendCode19(code byte, ind2 int) (int, error) {
	ns := g.streams[g.initIndex.i]
	var buf [4]byte
	buf[0] = 0
	buf[1] = 0
	buf[2] = 19
	buf[3] = code

	nx := &g.streams2[ind2]
	nx.active = false

	return ns.WriteTo(buf[:4], nx.addr)
}

func (g *Streams) ProcessStats(v105, v107 int) {
	start := platformTicks()
	for i := range g.streams2 {
		nx := &g.streams2[i]
		if !nx.active {
			continue
		}
		v2 := nx.cur
		if int(v2) >= len(nx.arr) {
			if int(nx.lost) > g.MaxPacketLoss {
				_, _ = g.sendCode19(1, i)
				_, _ = g.sendCode20(i)
				continue
			}
			cnt := 0
			sum := 0
			for i = 0; i < 10; i++ {
				if nx.arr[i] > 0 {
					cnt++
					sum += nx.arr[i]
				}
			}
			avg := sum / cnt
			if v105 != -1 && avg < v105 {
				_, _ = g.sendCode19(0, i)
			}
			if v107 != -1 && avg > v107 {
				_, _ = g.sendCode19(1, i)
			}
			_, _ = g.sendCode20(i)
		} else {
			if start-nx.ticks > 2000 {
				g.streams2[i].arr[v2] = -1
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

var queueIter *queueItem

func QueueFirst(ind Index) []byte {
	ns := ind.get()
	if ns == nil {
		return nil
	}
	it := ns.queue
	if it == nil {
		return nil
	}
	queueIter = it.next
	return it.data[2:it.size]
}

func QueueNext(ind Index) []byte {
	if queueIter == nil {
		return nil
	}
	ns := ind.get()
	if ns == nil {
		return nil
	}
	next := queueIter
	queueIter = next
	return next.data[2:next.size]
}

func (g *Streams) WaitServerResponse(ind1 Index, a2 int, a3 int, flags int) int {
	if g.Debug {
		Log.Printf("nox_xxx_cliWaitServerResponse_5525B0: %d, %d, %d, %d\n", ind1, a2, a3, flags)
	}
	ns := ind1.get()
	if ns == nil {
		return -3
	}

	if int(ns.ind28) >= a2 {
		return 0
	}
	for v6 := 0; v6 <= 20*a3; v6++ {
		platform.Sleep(50 * time.Millisecond)
		g.ServeInitialPackets(ind1, flags|1)
		g.MaybeSendAll()
		if int(ns.ind28) >= a2 {
			return 0
		}
		// FIXME(awesie)
		return 0
	}
	return -23
}

func SendSelfRaw(ind Index, buf []byte) int {
	return ind.get().SendSelfRaw(buf)
}

func (ns *stream) SendSelfRaw(buf []byte) int {
	if ns == nil {
		return 0
	}
	n, _ := ns.WriteToRaw(buf, ns.addr)
	return n
}

func (ns *stream) WriteToRaw(buf []byte, addr netip.AddrPort) (int, error) {
	if ns == nil || len(buf) == 0 {
		return 0, nil
	}
	return WriteTo(ns.pc, buf, addr)
}

var sendXorBuf [4096]byte // TODO: remove this global buffer

func (ns *stream) WriteTo(buf []byte, addr netip.AddrPort) (int, error) {
	if ns == nil || len(buf) == 0 {
		return 0, nil
	}
	ns2 := ns.g.structByAddr(addr)
	if ns2 == nil || ns2.xorKey == 0 {
		return ns.WriteToRaw(buf, addr)
	}
	dst := sendXorBuf[:len(buf)]
	ns.g.netCryptDst(ns2.xorKey, buf, dst)
	return ns.WriteToRaw(dst, ns2.addr)
}

func (g *Streams) netCryptDst(key byte, src, dst []byte) {
	if len(src) == 0 || len(dst) == 0 {
		return
	}
	if !g.Xor {
		copy(dst, src)
		return
	}
	for i := range src {
		dst[i] = key ^ src[i]
	}
}
