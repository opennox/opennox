package dialog

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/noxworld-dev/opennox/v1/legacy/client/audio/ail"
)

type Dialog struct {
	isEnabled          *uint32     // dword_587000_122848 - Whether dialog component is enabled.
	volume             *uint32     // dword_5d4594_830868 - stores volume level of dialog at 0-100 range
	fileToRead         **byte      // dword_5d4594_830872
	currentPlayingFile **byte      // dword_5d4594_830972
	currentStream      *ail.Stream // dword_5d4594_831088 - Stores a dialog stream currently playing
	state              *uint32     // dword_5d4594_830864 - Internal state of dialog component
	someFlag           *uint32     // dword_5d4594_831084 - Internal state being toggled when loading failed?
	isFallbackMode     *uint32     // dword_587000_122856 - Initially true, becomes false after first dialog read success
	isInitialized      *uint32     // dword_5d4594_831076
	audioDriver        *ail.Driver // dword_5d4594_831092
}

func (d *Dialog) GoString() string {
	return fmt.Sprintf("Dialog{isEnabled: %v, volume: %v, fileToRead: %v, currentPlayingFile: %v, currentStream: %v, state: %v, someFlag: %v, isFallbackMode: %v, isInitialized: %v, audioDriver: %v}",
		*d.isEnabled, *d.volume, *d.fileToRead, *d.currentPlayingFile, *d.currentStream, *d.state, *d.someFlag, *d.isFallbackMode, *d.isInitialized, *d.audioDriver)
}

func Get_dword_587000_122848() int {
	return int(*Instance.isEnabled)
}

func Get_dword_5d4594_831088() ail.Stream {
	return ail.Stream(*Instance.currentStream)
}

func Set_dword_5d4594_831088(v ail.Stream) {
	*Instance.currentStream = v
}

func Get_dword_587000_122856() int {
	return int(*Instance.isFallbackMode)
}

func Set_dword_587000_122856(v int) {
	*Instance.isFallbackMode = uint32(v)
}

//export sub_44D930
func sub_44D930() int32 {
	if Get_dword_587000_122848() == 0 {
		return 0
	}
	if *Instance.fileToRead != nil || Get_dword_5d4594_831088() != 0 {
		return 1
	}
	return 0
}

func Sub_44D930() bool {
	return sub_44D930() != 0
}

func bool2int(b bool) int {
	if b {
		return 1
	}
	return 0
}

func Nox_xxx_WorkerHurt_44D810() int32 {
	if *Instance.isInitialized != 0 {
		return 1
	}
	*Instance.audioDriver = ail.Driver(Sub_43F130())
	*Instance.isEnabled = uint32(bool2int(*Instance.audioDriver != 0))
	ptr_830876 := Get_counter_5d4594_830876()
	ptr_830876.Init()
	(&ptr_830876.Timers[0]).SetParams(uint32(0x1F4), uint32(0x4000))
	*Instance.state = 0
	*Instance.currentPlayingFile = nil
	*Instance.fileToRead = nil
	Set_dword_5d4594_831080(0)
	*Instance.someFlag = 0
	*Instance.isInitialized = 1

	v, res := GetStringManager().GetVariantInFile("Con03B.scr:WorkerHurt", "C:\\NoxPost\\src\\client\\Audio\\AudDiag.c")
	if res && v.Str2 != "" {
		Nox_xxx_playDialogFile_44D900(InternCStr(v.Str2), 0)
	}
	return 1
}

func Sub_44D8F0() {
	*Instance.fileToRead = nil
	*Instance.currentPlayingFile = nil
}

func Sub_44D5C0(s ail.Stream, a2 int) {
	if s != 0 {
		v2 := (Get_counter_587000_81128_ptr().Timers[0].Current >> 16) *
			((Get_counter_5d4594_830876().Timers[0].Current >> 16) *
				((Get_counter_5d4594_830980().Timers[0].Current >> 16) *
					uint32(a2) /
					0x4000) /
				0x4000) /
			0x4000
		Get_counter_5d4594_830876().Timers[0].Flags &= 0xFFFFFFFD
		Get_counter_587000_122852_ptr().Timers[0].Flags &= 0xFFFFFFFD
		s.SetVolume(int(127 * v2 / 100))
	}
}

func Sub_44D3A0() {
	if *Instance.isInitialized == 0 {
		return
	}
	timer_830876_ptr := Get_counter_5d4594_830876()
	timer_830980_ptr := Get_counter_5d4594_830980()
	audioStream2 := Get_dword_5d4594_831088()
	switch *Instance.state {
	case 0:
		if *Instance.fileToRead != nil && *Instance.isEnabled != 0 {
			if !sub_44D660(GoString(*Instance.fileToRead)) {
				*Instance.fileToRead = nil
			} else if *Instance.isFallbackMode == 0 || Get_dword_587000_93160() == 0 || *Instance.someFlag != 0 {
				*Instance.state = 2
			} else {
				*Instance.someFlag = 1
				Sub_43DBD0()
				*Instance.state = 1
			}
		} else if *Instance.someFlag != 0 {
			Sub_43DBE0()
			*Instance.someFlag = 0
		}
	case 1:
		if Sub_43DC40() == 0 {
			*Instance.state = 2
		}
	case 2:
		(&timer_830876_ptr.Timers[0]).SetRaw(0x4000)
		if sub_44D7E0(int(*Instance.volume)) == 0 {
			*Instance.state = 0
			*Instance.fileToRead = nil
		} else {
			*Instance.state = 3
			*Instance.currentPlayingFile = *Instance.fileToRead
			Set_dword_5d4594_830860(*Instance.volume)
		}
	case 3:
		if *Instance.isEnabled == 0 || *Instance.currentPlayingFile == nil || *Instance.fileToRead != *Instance.currentPlayingFile || Get_dword_5d4594_831088() == 0 || Get_dword_5d4594_831088().Status() == 2 {
			*Instance.state = 4
			(&timer_830876_ptr.Timers[0]).SetInterp(uint32(0))
		}
	case 4:
		if audioStream2 == 0 || audioStream2.Status() == 2 || (timer_830876_ptr.Timers[0].Current&0xFFFF0000) == 0 {
			sub_44D640()
			*Instance.state = 0
			if *Instance.someFlag != 0 {
				Sub_43DBE0()
				*Instance.someFlag = 0
			}
			if *Instance.currentPlayingFile == *Instance.fileToRead {
				*Instance.fileToRead = nil
			}
		}
	}
	timer_830980_ptr.Update()
	timer_830876_ptr.Update()
	if audioStream2 != 0 && (Get_counter_587000_81128_ptr().Timers[0].Flags&2 != 0 || timer_830876_ptr.Timers[0].Flags&2 != 0 || timer_830980_ptr.Timers[0].Flags&2 != 0) {
		Sub_44D5C0(audioStream2, int(Get_dword_5d4594_830860()))
	}
}

func Sub_44D8C0() {
	if *Instance.isInitialized != 0 {
		Sub_44D8F0()
		sub_44D640()
		Sub_44D3A0()
		*Instance.isInitialized = 0
	}
}

func sub_44D640() {
	if s := Get_dword_5d4594_831088(); s != 0 {
		s.Close()
		Set_dword_5d4594_831088(0)
	}
}

func sub_44D660(name string) bool {
	sub_44D640()
	path := filepath.Join("dialog", name)
	*Instance.isFallbackMode = 0
	if !strings.Contains(path, ".") {
		path += ".wav"
	}
	s := get_dword_5d4594_831092().OpenStream(path, 51200)
	Set_dword_5d4594_831088(s)
	if s != 0 {
		return true
	}
	v4 := Sub_413890()
	if v4 == "" {
		return s != 0
	}
	*Instance.isFallbackMode = 1
	path2 := filepath.Join(v4, path)
	s = get_dword_5d4594_831092().OpenStream(path2, 51200)
	Set_dword_5d4594_831088(s)
	return s != 0
}

func Nox_xxx_playDialogFile_44D900(a1 *byte, a2 int) int {
	var (
		v2 int
	)
	if !(*Instance.isEnabled != 0 && a1 != nil) {
		return 1
	}
	v2 = a2
	if a2 > 100 {
		v2 = 100
	}
	*Instance.fileToRead = a1
	*Instance.volume = uint32(v2)
	return 1
}

func sub_44D7E0(a1 int) int {
	s := Get_dword_5d4594_831088()
	if s == 0 {
		return 0
	}
	Sub_44D5C0(s, a1)
	s.Start()
	return 1
}

func get_dword_5d4594_831092() ail.Driver {
	return ail.Driver(*Instance.audioDriver)
}
