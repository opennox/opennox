package legacy

/*
#include <stdint.h>

extern uint32_t dword_5d4594_816092;
extern uint32_t dword_5d4594_816364;
extern uint32_t dword_5d4594_816376;
extern uint32_t dword_5d4594_830864;
extern uint32_t dword_5d4594_830868;
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
	"unsafe"

	"github.com/noxworld-dev/opennox-lib/strman"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"github.com/noxworld-dev/opennox/v1/legacy/client/audio/ail"
	"github.com/noxworld-dev/opennox/v1/legacy/dialog"
	"github.com/noxworld-dev/opennox/v1/legacy/timer"
)

func Get_counter_5d4594_830876() *timer.TimerGroup {
	return (*timer.TimerGroup)(memmap.PtrOff(0x5D4594, 830876))
}

func Get_counter_5d4594_830980() *timer.TimerGroup {
	return (*timer.TimerGroup)(memmap.PtrOff(0x5D4594, 830980))
}

func Set_dword_5d4594_831080(v uint32) {
	*memmap.PtrUint32(0x5D4594, 831080) = v
}

func Get_dword_5d4594_830860() uint32 {
	return memmap.Uint32(0x5D4594, 830860)
}

func Set_dword_5d4594_830860(v uint32) {
	*memmap.PtrUint32(0x5D4594, 830860) = v
}

func Get_counter_587000_81128_ptr() *timer.TimerGroup {
	return (*timer.TimerGroup)(C.dword_587000_81128)
}

func Get_counter_587000_122852_ptr() *timer.TimerGroup {
	return (*timer.TimerGroup)(C.dword_587000_122852)
}

func init() {
	dialog.Sub_43F130 = sub_43F130
	dialog.Sub_43DBD0 = Sub_43DBD0
	dialog.Sub_43DBE0 = Sub_43DBE0
	dialog.Sub_43DC40 = Sub_43DC40
	dialog.Sub_413890 = Sub_413890
	dialog.GetStringManager = func() *strman.StringManager {
		return GetServer().S().Strings()
	}
	dialog.Get_counter_5d4594_830876 = Get_counter_5d4594_830876
	dialog.Get_counter_5d4594_830980 = Get_counter_5d4594_830980
	dialog.Set_dword_5d4594_831080 = Set_dword_5d4594_831080
	dialog.Get_dword_587000_93160 = func() uint32 { return uint32(Get_dword_587000_93160()) }
	dialog.Get_dword_5d4594_830860 = Get_dword_5d4594_830860
	dialog.Set_dword_5d4594_830860 = Set_dword_5d4594_830860

	dialog.GoString = func(p *byte) string { return C.GoString((*C.char)(unsafe.Pointer(p))) }
	dialog.InternCStr = func(p string) *byte { return (*byte)(unsafe.Pointer(internCStr(p))) }
	dialog.Instance = dialog.NewDialog(
		(*uint32)(&C.dword_587000_122848), (*uint32)(&C.dword_5d4594_830868),
		(**byte)(unsafe.Pointer(&C.dword_5d4594_830872)), (**byte)(unsafe.Pointer(&C.dword_5d4594_830972)), (*ail.Stream)(unsafe.Pointer(&C.dword_5d4594_831088)),
		(*uint32)(&C.dword_5d4594_830864), (*uint32)(&C.dword_5d4594_831084), (*uint32)(&C.dword_587000_122856), (*uint32)(&C.dword_5d4594_831076), (*ail.Driver)(unsafe.Pointer(&C.dword_5d4594_831092)))
	dialog.Get_counter_587000_81128_ptr = Get_counter_587000_81128_ptr
	dialog.Get_counter_587000_122852_ptr = Get_counter_587000_122852_ptr
}

//export sub_44D930
func sub_44D930() int {
	return bool2int(dialog.Sub_44D930())
}

//export nox_xxx_playDialogFile_44D900
func nox_xxx_playDialogFile_44D900(a1p *byte, a2 int) int {
	return dialog.Nox_xxx_playDialogFile_44D900(a1p, a2)
}
