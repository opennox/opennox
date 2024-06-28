package legacy

/*
#include <stdint.h>

extern uint32_t dword_5d4594_816092;
extern uint32_t dword_5d4594_816364;
extern uint32_t dword_5d4594_816376;
extern uint32_t dword_5d4594_830864;
extern uint32_t dword_5d4594_830872;
extern uint32_t dword_5d4594_830972;
extern uint32_t dword_5d4594_831076;
extern uint32_t dword_5d4594_831084;
extern uint32_t dword_5d4594_831088;
extern uint32_t dword_5d4594_831092;
extern uint32_t dword_587000_122856;
extern uint32_t dword_587000_122848;
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
		fileToRead:         (**C.char)((unsafe.Pointer)(&C.dword_5d4594_830872)),
		currentPlayingFile: (**C.char)((unsafe.Pointer)(&C.dword_5d4594_830972)),
		currentStream:      (*ail.Stream)((unsafe.Pointer)(&C.dword_5d4594_831088)),
		state:              (*uint32)(&C.dword_5d4594_830864),
		someFlag:           (*uint32)(&C.dword_5d4594_831084),
		isFallbackMode:     (*uint32)(&C.dword_587000_122856),
		isInitialized:      (*uint32)(&C.dword_5d4594_831076),
	}
}

func Get_dword_5d4594_830872() int {
	return int(C.dword_5d4594_830872)
}

func Set_dword_5d4594_830872(v int) {
	C.dword_5d4594_830872 = C.uint(v)
}

func Set_dword_5d4594_830972(v int) {
	C.dword_5d4594_830972 = C.uint(v)
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
	if Get_dword_5d4594_830872() != 0 || Get_dword_5d4594_831088() != 0 {
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
	TimerGroupInit(counter_5d4594_830876)
	TimerSetParams(&counter_5d4594_830876.Timers[0], 0x1F4 /* 500 */, 0x4000)
	C.dword_5d4594_830864 = 0
	C.dword_5d4594_830972 = 0
	C.dword_5d4594_830872 = 0
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
	Set_dword_5d4594_830872(0)
	Set_dword_5d4594_830972(0)
}

//export sub_44D8F0
func sub_44D8F0() { Sub_44D8F0() }

func Get_dword_587000_122852() unsafe.Pointer {
	return unsafe.Pointer(counter_5d4594_830980)
}

func Get_dword_587000_93164() unsafe.Pointer {
	return unsafe.Pointer(counter_5d4594_816244)
}

func Get_dword_587000_127004() unsafe.Pointer {
	return unsafe.Pointer(counter_5d4594_1045324)
}
