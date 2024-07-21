package dialog

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/noxworld-dev/opennox-lib/strman"

	"github.com/noxworld-dev/opennox/v1/legacy/client/audio/ail"
	"github.com/noxworld-dev/opennox/v1/legacy/timer"
)

func NewDialog(
	dir string,
	isEnabled *uint32, state *uint32, someFlag *uint32, isFallbackMode *uint32, isInitialized *uint32, audioDriver *ail.Driver,
	dword830860 *uint32,
	getStrings func() *strman.StringManager,
	timers1 *timer.TimerGroup, timers2 *timer.TimerGroup,
	timers3 *timer.TimerGroup, timers4 *timer.TimerGroup,
	fallbackDir func() string,

	sub_43F130 func() ail.Driver, // get audio device
	sub_43DBD0 func(), // looks like music? pair with 43DBE0
	sub_43DBE0 func(), // looks like music? pair with 43DBD0
	sub_43DC40 func() bool, // unknown

	set_dword_5d4594_831080 func(uint32),
	get_dword_587000_93160 func() uint32,
) *Dialog {
	return &Dialog{
		dir:            dir,
		isEnabled:      isEnabled,
		state:          state,
		someFlag:       someFlag,
		isFallbackMode: isFallbackMode,
		isInitialized:  isInitialized,
		audioDriver:    audioDriver,
		getStrings:     getStrings,
		timers1:        timers1,
		timers2:        timers2,
		timers3:        timers3,
		timers4:        timers4,
		fallbackDir:    fallbackDir,
		dword830860:    dword830860,

		newDriver:  sub_43F130,
		sub_43DBD0: sub_43DBD0,
		sub_43DBE0: sub_43DBE0,
		sub_43DC40: sub_43DC40,

		set_dword_5d4594_831080: set_dword_5d4594_831080,
		get_dword_587000_93160:  get_dword_587000_93160,
	}
}

type Dialog struct {
	dir                string
	isEnabled          *uint32           // dword_587000_122848 - Whether dialog component is enabled.
	volume             uint32            // dword_5d4594_830868 - stores volume level of dialog at 0-100 range
	fileToRead         string            // dword_5d4594_830872
	currentPlayingFile string            // dword_5d4594_830972
	currentStream      ail.Stream        // dword_5d4594_831088 - Stores a dialog stream currently playing
	state              *uint32           // dword_5d4594_830864 - Internal state of dialog component
	someFlag           *uint32           // dword_5d4594_831084 - Internal state being toggled when loading failed?
	isFallbackMode     *uint32           // dword_587000_122856 - Initially true, becomes false after first dialog read success
	isInitialized      *uint32           // dword_5d4594_831076
	audioDriver        *ail.Driver       // dword_5d4594_831092
	timers1            *timer.TimerGroup // 5d4594_830876
	timers2            *timer.TimerGroup // 5d4594_830980
	timers3            *timer.TimerGroup // 587000_81128
	timers4            *timer.TimerGroup // 587000_122852
	getStrings         func() *strman.StringManager
	fallbackDir        func() string // fallback audio file name provider
	dword830860        *uint32

	newDriver  func() ail.Driver // get audio device
	sub_43DBD0 func()            // looks like music? pair with 43DBE0
	sub_43DBE0 func()            // looks like music? pair with 43DBD0
	sub_43DC40 func() bool       // unknown

	set_dword_5d4594_831080 func(uint32)
	get_dword_587000_93160  func() uint32
}

func (d *Dialog) GoString() string {
	return fmt.Sprintf("Dialog{dir: %v, isEnabled: %v, volume: %v, fileToRead: %v, currentPlayingFile: %v, currentStream: %v, state: %v, someFlag: %v, isFallbackMode: %v, isInitialized: %v, audioDriver: %v}",
		d.dir, *d.isEnabled, d.volume, d.fileToRead, d.currentPlayingFile, d.currentStream, *d.state, *d.someFlag, *d.isFallbackMode, *d.isInitialized, *d.audioDriver)
}

func (d *Dialog) IsInitialized() bool {
	return *d.isInitialized != 0
}

func (d *Dialog) State() int {
	return int(*d.state)
}

func (d *Dialog) FileToRead() string {
	return d.fileToRead
}

func (d *Dialog) CurrentPlayingFile() string {
	return d.currentPlayingFile
}

func (d *Dialog) GetStream() ail.Stream {
	return d.currentStream
}

func (d *Dialog) SetStream(v ail.Stream) {
	d.currentStream = v
}

func (d *Dialog) getDriver() ail.Driver {
	return *d.audioDriver
}

func (d *Dialog) IsFallbackMode() int {
	return int(*d.isFallbackMode)
}

func (d *Dialog) Sub_44D930() bool {
	if *d.isEnabled == 0 {
		return false
	}
	if d.fileToRead != "" || d.GetStream() != 0 {
		return true
	}
	return false
}

func (d *Dialog) Nox_xxx_WorkerHurt_44D810() int32 {
	if *d.isInitialized != 0 {
		return 1
	}
	*d.audioDriver = d.newDriver()
	if *d.audioDriver != 0 {
		*d.isEnabled = 1
	} else {
		*d.isEnabled = 0
	}
	d.timers1.Init()
	d.timers1.Timers[0].SetParams(uint32(0x1F4), uint32(0x4000))
	*d.state = 0
	d.currentPlayingFile = ""
	d.fileToRead = ""
	d.set_dword_5d4594_831080(0)
	*d.someFlag = 0
	*d.isInitialized = 1

	v, ok := d.getStrings().GetVariantInFile("Con03B.scr:WorkerHurt", "C:\\NoxPost\\src\\client\\Audio\\AudDiag.c")
	if ok && v.Str2 != "" {
		d.PlayFile(v.Str2, 0)
	}
	return 1
}

func (d *Dialog) Sub_44D8F0() {
	d.fileToRead = ""
	d.currentPlayingFile = ""
}

func (d *Dialog) Sub_44D5C0(s ail.Stream, a2 int) {
	if s != 0 {
		v2 := (d.timers3.Timers[0].Current >> 16) *
			((d.timers1.Timers[0].Current >> 16) *
				((d.timers2.Timers[0].Current >> 16) *
					uint32(a2) /
					0x4000) /
				0x4000) /
			0x4000
		d.timers1.Timers[0].Flags &= 0xFFFFFFFD
		d.timers4.Timers[0].Flags &= 0xFFFFFFFD
		s.SetVolume(int(127 * v2 / 100))
	}
}

func (d *Dialog) Sub_44D3A0() {
	if *d.isInitialized == 0 {
		return
	}
	audioStream2 := d.GetStream()
	switch *d.state {
	case 0:
		if d.fileToRead != "" && *d.isEnabled != 0 {
			if !d.playFile(d.fileToRead) {
				d.fileToRead = ""
			} else if *d.isFallbackMode == 0 || d.get_dword_587000_93160() == 0 || *d.someFlag != 0 {
				*d.state = 2
			} else {
				*d.someFlag = 1
				d.sub_43DBD0()
				*d.state = 1
			}
		} else if *d.someFlag != 0 {
			d.sub_43DBE0()
			*d.someFlag = 0
		}
	case 1:
		if d.sub_43DC40() == false {
			*d.state = 2
		}
	case 2:
		d.timers1.Timers[0].SetRaw(0x4000)
		if d.sub_44D7E0(int(d.volume)) == 0 {
			*d.state = 0
			d.fileToRead = ""
		} else {
			*d.state = 3
			d.currentPlayingFile = d.fileToRead
			*d.dword830860 = d.volume
		}
	case 3:
		if *d.isEnabled == 0 || d.currentPlayingFile == "" || d.fileToRead != d.currentPlayingFile || d.GetStream() == 0 || d.GetStream().Status() == 2 {
			*d.state = 4
			d.timers1.Timers[0].SetInterp(uint32(0))
		}
	case 4:
		if audioStream2 == 0 || audioStream2.Status() == 2 || (d.timers1.Timers[0].Current&0xFFFF0000) == 0 {
			d.sub_44D640()
			*d.state = 0
			if *d.someFlag != 0 {
				d.sub_43DBE0()
				*d.someFlag = 0
			}
			if d.currentPlayingFile == d.fileToRead {
				d.fileToRead = ""
			}
		}
	}
	d.timers2.Update()
	d.timers1.Update()
	if audioStream2 != 0 && (d.timers3.Timers[0].Flags&2 != 0 || d.timers1.Timers[0].Flags&2 != 0 || d.timers2.Timers[0].Flags&2 != 0) {
		d.Sub_44D5C0(audioStream2, int(*d.dword830860))
	}
}

func (d *Dialog) Sub_44D8C0() {
	if *d.isInitialized != 0 {
		d.Sub_44D8F0()
		d.sub_44D640()
		d.Sub_44D3A0()
		*d.isInitialized = 0
	}
}

func (d *Dialog) sub_44D640() {
	if s := d.GetStream(); s != 0 {
		s.Close()
		d.SetStream(0)
	}
}

func (d *Dialog) playFile(name string) bool {
	d.sub_44D640()
	*d.isFallbackMode = 0
	path := filepath.Join(d.dir, name)
	if !strings.Contains(path, ".") {
		path += ".wav"
	}
	s := d.getDriver().OpenStream(path)
	d.SetStream(s)
	if s != 0 {
		return true
	}
	if d.fallbackDir == nil {
		return true
	}
	fdir := d.fallbackDir()
	if fdir == "" {
		return true
	}
	*d.isFallbackMode = 1
	path2 := filepath.Join(fdir, path)
	s = d.getDriver().OpenStream(path2)
	d.SetStream(s)
	return s != 0
}

func (d *Dialog) PlayFile(file string, vol int) int {
	if *d.isEnabled == 0 || file == "" {
		return 1
	}
	if vol > 100 {
		vol = 100
	}
	d.fileToRead = file
	d.volume = uint32(vol)
	return 1
}

func (d *Dialog) sub_44D7E0(a1 int) int {
	s := d.GetStream()
	if s == 0 {
		return 0
	}
	d.Sub_44D5C0(s, a1)
	s.Start()
	return 1
}
