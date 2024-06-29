package legacy

/*
#include <stdint.h>

extern uint32_t dword_5d4594_816092;
extern uint32_t dword_5d4594_816364;
extern uint32_t dword_5d4594_816376;
extern uint32_t dword_5d4594_830864;
extern char* dword_5d4594_830872;
extern char* dword_5d4594_830972;
extern uint32_t dword_5d4594_831076;
extern uint32_t dword_5d4594_831084;
extern uint32_t dword_5d4594_831088;
extern uint32_t dword_5d4594_831092;
extern uint32_t dword_587000_122856;
extern uint32_t dword_587000_122848;
extern void* dword_587000_81128;
extern void* dword_587000_122852;
*/
import "C"
import (
	"fmt"
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"github.com/noxworld-dev/opennox/v1/legacy/client/audio/ail"
)

type Dialog struct {
	isEnabled          *uint32     // dword_587000_122848 - Whether dialog component is enabled.
	volume             *uint32     // dword_5d4594_830868 - stores volume level of dialog at 0-100 range
	fileToRead         **C.char    // dword_5d4594_830872
	currentPlayingFile **C.char    // dword_5d4594_830972
	currentStream      *ail.Stream // dword_5d4594_831088 - Stores a dialog stream currently playing
	state              *uint32     // dword_5d4594_830864 - Internal state of dialog component
	someFlag           *uint32     // dword_5d4594_831084 - Internal state being toggled when loading failed?
	isFallbackMode     *uint32     // dword_587000_122856 - Initially true, becomes false after first dialog read success
	isInitialized      *uint32     // dword_5d4594_831076
}

func (d *Dialog) GoString() string {
	return fmt.Sprintf("Dialog{isEnabled: %v, volume: %v, fileToRead: %v, currentPlayingFile: %v, currentStream: %v, state: %v, someFlag: %v, isFallbackMode: %v, isInitialized: %v}",
		*d.isEnabled, *d.volume, C.GoString(*d.fileToRead), C.GoString(*d.currentPlayingFile), *d.currentStream, *d.state, *d.someFlag, *d.isFallbackMode, *d.isInitialized)
}

func DialogInit() *Dialog {
	return &Dialog{
		isEnabled:          (*uint32)(&C.dword_587000_122848),
		volume:             memmap.PtrUint32(0x5D4594, 830868),
		fileToRead:         (**C.char)(&C.dword_5d4594_830872),
		currentPlayingFile: (**C.char)(&C.dword_5d4594_830972),
		currentStream:      (*ail.Stream)((unsafe.Pointer)(&C.dword_5d4594_831088)),
		state:              (*uint32)(&C.dword_5d4594_830864),
		someFlag:           (*uint32)(&C.dword_5d4594_831084),
		isFallbackMode:     (*uint32)(&C.dword_587000_122856),
		isInitialized:      (*uint32)(&C.dword_5d4594_831076),
	}
}

func Get_dword_587000_122848() int {
	return int(C.dword_587000_122848)
}

func Get_dword_5d4594_831088() ail.Stream {
	return ail.Stream(C.dword_5d4594_831088)
}

func Set_dword_5d4594_831088(v ail.Stream) {
	C.dword_5d4594_831088 = C.uint(v)
}

func Get_dword_587000_122856() int {
	return int(C.dword_587000_122856)
}

func Set_dword_587000_122856(v int) {
	C.dword_587000_122856 = C.uint(v)
}

//export sub_44D930
func sub_44D930() C.int {
	if Get_dword_587000_122848() == 0 {
		return 0
	}
	if C.dword_5d4594_830872 != nil || Get_dword_5d4594_831088() != 0 {
		return 1
	}
	return 0
}

func Sub_44D930() bool {
	return sub_44D930() != 0
}

func Nox_xxx_WorkerHurt_44D810() int32 {
	if C.dword_5d4594_831076 != 0 {
		return 1
	}
	C.dword_5d4594_831092 = C.uint32_t(sub_43F130())
	C.dword_587000_122848 = C.uint32_t(bool2int(C.dword_5d4594_831092 != 0))
	ptr_830876 := Get_counter_5d4594_830876()
	TimerGroupInit(ptr_830876)
	TimerSetParams(&ptr_830876.Timers[0], 0x1F4 /* 500 */, 0x4000)
	C.dword_5d4594_830864 = 0
	C.dword_5d4594_830972 = nil
	C.dword_5d4594_830872 = nil
	*memmap.PtrUint32(0x5D4594, 831080) = uint32(0) // This address looks unused
	C.dword_5d4594_831084 = 0
	C.dword_5d4594_831076 = 1

	v, res := GetServer().S().Strings().GetVariantInFile("Con03B.scr:WorkerHurt", "C:\\NoxPost\\src\\client\\Audio\\AudDiag.c")
	if res && v.Str2 != "" {
		Nox_xxx_playDialogFile_44D900(v.Str2, 0)
	}
	return 1
}

func Sub_44D8F0() {
	C.dword_5d4594_830872 = nil
	C.dword_5d4594_830972 = nil
}

//export sub_44D8F0
func sub_44D8F0() { Sub_44D8F0() }

func Get_counter_587000_81128_ptr() *TimerGroup {
	return (*TimerGroup)(C.dword_587000_81128)
}

func Get_counter_5d4594_830876() *TimerGroup {
	return (*TimerGroup)(memmap.PtrOff(0x5D4594, 830876))
}

func Get_counter_5d4594_830980() *TimerGroup {
	return (*TimerGroup)(memmap.PtrOff(0x5D4594, 830980))
}

func Get_counter_587000_122852_ptr() *TimerGroup {
	return (*TimerGroup)(C.dword_587000_122852)
}

//export sub_44D5C0
func sub_44D5C0(a1 int, a2 int) {
	Sub_44D5C0(ail.Stream(a1), a2)
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
	if C.dword_5d4594_831076 == 0 {
		return
	}
	timer_830876_ptr := Get_counter_5d4594_830876()
	timer_830980_ptr := Get_counter_5d4594_830980()
	audioStream2 := Get_dword_5d4594_831088()
	switch C.dword_5d4594_830864 {
	case 0:
		if C.dword_5d4594_830872 != nil && C.dword_587000_122848 != 0 {
			if sub_44D660(C.dword_5d4594_830872) == 0 {
				C.dword_5d4594_830872 = nil
			} else if C.dword_587000_122856 == 0 || *memmap.PtrUint32(0x587000, 93160) == 0 || C.dword_5d4594_831084 != 0 {
				C.dword_5d4594_830864 = 2
			} else {
				C.dword_5d4594_831084 = 1
				Sub_43DBD0()
				C.dword_5d4594_830864 = 1
			}
		} else if C.dword_5d4594_831084 != 0 {
			Sub_43DBE0()
			C.dword_5d4594_831084 = 0
		}
	case 1:
		if Sub_43DC40() == 0 {
			C.dword_5d4594_830864 = 2
		}
	case 2:
		TimerSetRaw(&timer_830876_ptr.Timers[0], 0x4000)
		if sub_44D7E0(int(memmap.Int32(0x5D4594, 830868))) == 0 {
			C.dword_5d4594_830864 = 0
			C.dword_5d4594_830872 = nil
		} else {
			C.dword_5d4594_830864 = 3
			C.dword_5d4594_830972 = C.dword_5d4594_830872
			*memmap.PtrUint32(0x5D4594, 830860) = memmap.Uint32(0x5D4594, 830868)
		}
	case 3:
		if C.dword_587000_122848 == 0 || C.dword_5d4594_830972 == nil || C.dword_5d4594_830872 != C.dword_5d4594_830972 || Get_dword_5d4594_831088() == 0 || Get_dword_5d4594_831088().Status() == 2 {
			C.dword_5d4594_830864 = 4
			TimerSetInterp(&timer_830876_ptr.Timers[0], 0)
		}
	case 4:
		if audioStream2 == 0 || audioStream2.Status() == 2 || (timer_830876_ptr.Timers[0].Current&0xFFFF0000) == 0 {
			sub_44D640()
			C.dword_5d4594_830864 = 0
			if C.dword_5d4594_831084 != 0 {
				Sub_43DBE0()
				C.dword_5d4594_831084 = 0
			}
			if C.dword_5d4594_830972 == C.dword_5d4594_830872 {
				C.dword_5d4594_830872 = nil
			}
		}
	}
	TimerGroupUpdate(timer_830980_ptr)
	TimerGroupUpdate(timer_830876_ptr)
	if audioStream2 != 0 && (int32(*(*uint8)(C.dword_587000_81128))&2 != 0 || timer_830876_ptr.Timers[0].Flags&2 != 0 || timer_830980_ptr.Timers[0].Flags&2 != 0) {
		Sub_44D5C0(audioStream2, int(memmap.Int32(0x5D4594, 830860)))
	}
}

func Sub_44D8C0() {
	if C.dword_5d4594_831076 != 0 {
		sub_44D8F0()
		sub_44D640()
		Sub_44D3A0()
		C.dword_5d4594_831076 = 0
	}
}
