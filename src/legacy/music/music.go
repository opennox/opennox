package music

import (
	"path/filepath"
	"strings"
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/legacy/client/audio/ail"
	"github.com/noxworld-dev/opennox/v1/legacy/timer"
)

type MusicState struct {
	MusicIdx, Volume, Position, D uint32
}

var _ = [1]struct{}{}[16-unsafe.Sizeof(MusicState{})]

type MusicBlock struct {
	Inner  MusicState
	Loaded uint32
}

var _ = [1]struct{}{}[20-unsafe.Sizeof(MusicBlock{})]

type Module struct {
	dir string

	// private members
	currentPlaying      MusicState       // 5d4594_816092
	dword_5d4594_816356 uint32           // 5d4594_816356
	counterValue        int32            // 587000_93168
	lastUpdatedTick     uint64           // 5d4594_816380
	timer               timer.TimerGroup // 5d4594_816148
	block_idx           uint32           // 5d4594_816352
	blocks              [2]MusicBlock    // 5d4594_816108
	playingStream       ail.Stream       // 5d4594_816364
	dword_5d4594_816344 uint32           // 5d4594_816344
	block_5d4594_816060 MusicState       // 5d4594_816060

	// members which has external usages
	dword_5d4594_816368      *uint32            // 5d4594_816368
	dword_5d4594_816372      *uint32            // 5d4594_816372
	dword_5d4594_816376      *ail.Driver        // 5d4594_816376
	dword_587000_93156       *uint32            // 587000_93156
	dword_587000_93160       *uint32            // 587000_93160, used by dialog.go
	dword_5d4594_816340      *uint32            // 5d4594_816340
	counter_5d4594_816244    *timer.TimerGroup  // 5d4594_816244, used by dword_587000_93164
	ptr_counter_587000_81128 **timer.TimerGroup // 587000_81128
	dword_5d4594_816348      *uint32            // 5d4594_816348

	Sub_43F130    func() ail.Driver
	checkDialogs  func() bool
	Sub_413890    func() string
	PlatformTicks func() uint64
}

var audioTable = []string{
	"",
	"chap1", "chap2wiz", "chap2con", "chap2war", "chap3", "chap4", "chap5", "chap6", "chap7", "chap8",
	"chap9", "chapa", "chapb", "title", "town1", "town2", "town3", "sub1", "sub2", "sub3",
	"wander1", "wander2", "wander3", "credits", "shell", "action1", "action2", "action3", "wander4",
}

func NewModule(
	dir string,
	dword_5d4594_816368 *uint32,
	dword_5d4594_816372 *uint32,
	dword_5d4594_816376 *ail.Driver,
	dword_587000_93156 *uint32,
	dword_587000_93160 *uint32,
	dword_5d4594_816340 *uint32,
	counter_5d4594_816244 *timer.TimerGroup,
	ptr_counter_587000_81128 **timer.TimerGroup,
	dword_5d4594_816348 *uint32,
	Sub_43F130 func() ail.Driver,
	checkDialogs func() bool,
	Sub_413890 func() string,
	PlatformTicks func() uint64,
) *Module {
	m := &Module{
		dir:          dir,
		counterValue: -1,

		dword_5d4594_816368:      dword_5d4594_816368,
		dword_5d4594_816372:      dword_5d4594_816372,
		dword_5d4594_816376:      dword_5d4594_816376,
		dword_587000_93156:       dword_587000_93156,
		dword_587000_93160:       dword_587000_93160,
		dword_5d4594_816340:      dword_5d4594_816340,
		counter_5d4594_816244:    counter_5d4594_816244,
		ptr_counter_587000_81128: ptr_counter_587000_81128,
		dword_5d4594_816348:      dword_5d4594_816348,
		Sub_43F130:               Sub_43F130,
		checkDialogs:             checkDialogs,
		Sub_413890:               Sub_413890,
		PlatformTicks:            PlatformTicks,
	}
	return m
}

func (m *Module) Sub_43D2D0() {
	if *m.dword_5d4594_816340 == 0 {
		return
	}
	if m.playingStream == 0 {
		return
	}

	m.counter_5d4594_816244.Update()
	m.timer.Update()

	v1 := (*m.ptr_counter_587000_81128).Timers[0].Current >> 16
	if m.counterValue == int32(v1) {
		if (m.PlatformTicks() - m.lastUpdatedTick) > 0x32 {
			(*m.ptr_counter_587000_81128).Timers[0].Flags &= 0xFFFFFFFD
			(*m.ptr_counter_587000_81128).Timers[1].Flags &= 0xFFFFFFFD
			(*m.ptr_counter_587000_81128).Timers[2].Flags &= 0xFFFFFFFD
		}
	} else {
		m.lastUpdatedTick = m.PlatformTicks()
		m.counterValue = int32(v1)
	}

	if m.playingStream != 0 {
		if (*m.ptr_counter_587000_81128).Timers[0].Flags&2 != 0 || m.timer.Timers[0].Flags&2 != 0 || m.counter_5d4594_816244.Timers[0].Flags&2 != 0 {
			m.changeVolume(m.playingStream, m.currentPlaying.Volume)
		}
	}
}

func (m *Module) changeVolume(a1 ail.Stream, a2 uint32) {
	if a1 == 0 {
		return
	}
	v2 := ((*m.ptr_counter_587000_81128).Timers[0].Current >> 16) *
		((m.timer.Timers[0].Current >> 16) *
			((m.counter_5d4594_816244.Timers[0].Current >> 16) * a2 / 0x4000) / 0x4000) / 0x4000
	m.timer.Timers[0].Flags &= 0xFFFFFFFD
	m.counter_5d4594_816244.Timers[0].Flags &= 0xFFFFFFFD
	a1.SetVolume((int)(127*v2) / 100)
}

func (m *Module) pause() {
	if s := m.playingStream; s != 0 {
		s.Pause(true)
	}
}

func (m *Module) unpause() {
	if s := m.playingStream; s != 0 {
		s.Pause(false)
	}
}

// Return current music information
func (m *Module) GetCurrentBlock() MusicState {
	ret := m.blocks[m.block_idx].Inner
	if m.playingStream != 0 {
		if ret.MusicIdx == m.currentPlaying.MusicIdx {
			ret.Position = uint32(m.playingStream.Position())
		}
	}
	return ret
}

func (m *Module) Sub_43D650() {
	if s := m.playingStream; s != 0 {
		s.Close()
		m.playingStream = 0
	}
	m.currentPlaying.MusicIdx = 0
}

func (m *Module) startPlay(a1p *MusicState) bool {
	if *m.dword_5d4594_816376 == 0 {
		return false
	}

	ind := a1p.MusicIdx
	if ind <= 0 || ind >= uint32(len(audioTable)) {
		return false
	}
	aname := audioTable[ind]
	m.Sub_43D650()
	*m.dword_587000_93160 = 0
	path := filepath.Join(m.dir, aname)
	if !strings.Contains(path, ".") {
		path += ".wav"
	}
	s := m.dword_5d4594_816376.OpenStream(path)
	if s == 0 {
		if m.checkDialogs() {
			return false
		}
		v5 := m.Sub_413890()
		if v5 == "" {
			return false
		}
		*m.dword_587000_93160 = 1
		path2 := filepath.Join(v5, path)
		s = m.dword_5d4594_816376.OpenStream(path2)
		if s == 0 {
			return false
		}
	}
	s.SetPosition(int(a1p.Position))
	m.timer.Timers[0].SetRaw(0)
	m.timer.Timers[0].SetInterp(0x4000)
	m.changeVolume(s, a1p.Volume)
	s.Start()
	m.currentPlaying = *a1p
	a1p.Position = 0
	m.playingStream = s
	return true
}

func (m *Module) Init() {
	if *m.dword_5d4594_816340 != 0 {
		return
	}
	v1 := m.Sub_43F130()
	*m.dword_5d4594_816376 = v1
	if v1 != 0 {
		*m.dword_587000_93156 = 1
	} else {
		*m.dword_587000_93156 = 0
	}
	m.timer.Init()
	m.timer.Timers[0].SetParams(0x3E8, 0x4000)
	*m.dword_5d4594_816348 = 0
	m.currentPlaying.MusicIdx = 0
	m.block_idx = 0
	m.blocks[0].Inner.MusicIdx = 0
	m.dword_5d4594_816356 = 0
	*m.dword_5d4594_816372 = 0
	*m.dword_5d4594_816368 = 0
	m.playingStream = 0
	*m.dword_5d4594_816340 = 1
}

func (m *Module) Update() {
	current_block := m.blocks[m.block_idx]
	if *m.dword_5d4594_816340 == 0 {
		return
	}

	switch *m.dword_5d4594_816348 {
	case 0:
		if m.dword_5d4594_816356 != 0 && *m.dword_587000_93156 != 0 {
			*m.dword_5d4594_816348 = 3
			break
		}
		if current_block.Inner.MusicIdx == 0 {
			break
		}
		if *m.dword_587000_93156 != 0 {
			m.timer.Timers[0].SetRaw(0x4000)
			if m.startPlay(&current_block.Inner) {
				*m.dword_5d4594_816348 = 1
				current_block.Loaded = 1
			} else {
				current_block.Inner.MusicIdx = 0
			}
		}
	case 1:
		if *m.dword_587000_93156 != 0 && current_block.Inner.MusicIdx == m.currentPlaying.MusicIdx && m.playingStream != 0 && m.playingStream.Status() != 2 {
			if m.dword_5d4594_816356 != 0 {
				*m.dword_5d4594_816348 = 4
				m.timer.Timers[0].SetInterp(0)
			}
		} else {
			*m.dword_5d4594_816348 = 2
			m.timer.Timers[0].SetInterp(0)
		}
	case 2:
		if m.playingStream == 0 || m.playingStream.Status() == 2 || (m.timer.Timers[0].Current&0xFFFF0000) == 0 {
			m.Sub_43D650()
			*m.dword_5d4594_816348 = 0
		}
	case 3:
		if m.dword_5d4594_816356 != 0 && *m.dword_587000_93156 != 0 {
			break
		}
		if *m.dword_587000_93156 != 0 || current_block.Inner.MusicIdx != m.currentPlaying.MusicIdx || m.playingStream == 0 || m.playingStream.Status() == 2 {
			m.Sub_43D650()
			*m.dword_5d4594_816348 = 0
		} else {
			m.timer.Timers[0].SetInterp(0x4000)
			m.unpause()
			*m.dword_5d4594_816348 = 1
		}
	case 4:
		if m.playingStream != 0 && m.playingStream.Status() != 2 {
			if m.timer.Timers[0].Current&0xFFFF0000 != 0 {
				m.pause()
				*m.dword_5d4594_816348 = 3
			}
		} else {
			m.Sub_43D650()
			*m.dword_5d4594_816348 = 0
		}
	}
}

func (m *Module) SetNextMusic(a1 MusicState) {
	if a1.Volume > 100 {
		a1.Volume = 100
	}
	if m.dword_5d4594_816344 != 0 {
		m.block_5d4594_816060 = a1
	} else {
		m.block_idx = (m.block_idx + 1) % uint32(len(m.blocks))
		m.blocks[m.block_idx] = MusicBlock{Inner: a1, Loaded: 0}
	}
}

func (m *Module) Sub_43DBD0() {
	m.dword_5d4594_816356 += 1
	m.Update()
}

func (m *Module) Sub_43DBE0() {
	if m.dword_5d4594_816356 > 0 {
		m.dword_5d4594_816356 -= 1
	}
	m.Update()
}

func (m *Module) Sub_43DC40() bool {
	if *m.dword_5d4594_816340 == 0 {
		return false
	}
	if *m.dword_5d4594_816348 != 1 && *m.dword_5d4594_816348 != 4 && *m.dword_5d4594_816348 != 2 {
		return false
	}
	return true
}

func (m *Module) Sub_43DD70(a1, a2 uint32) {
	m.block_5d4594_816060 = m.GetCurrentBlock()
	m.SetNextMusic(MusicState{MusicIdx: a1, Volume: a2, Position: 0, D: 0})
	m.dword_5d4594_816344 = 1
}

func (m *Module) Sub_43DDA0() {
	m.dword_5d4594_816344 = 0
	m.SetNextMusic(m.block_5d4594_816060)
}
