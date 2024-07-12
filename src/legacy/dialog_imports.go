package legacy

/*
#include <stdint.h>

extern uint32_t dword_5d4594_816092;
extern uint32_t dword_5d4594_816364;
extern uint32_t dword_5d4594_816376;
extern uint32_t dword_5d4594_830864;
extern uint32_t dword_5d4594_831076;
extern uint32_t dword_5d4594_831084;
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
	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc"
	"github.com/noxworld-dev/opennox/v1/legacy/dialog"
	"github.com/noxworld-dev/opennox/v1/legacy/timer"
)

func Set_dword_5d4594_831080(v uint32) {
	*memmap.PtrUint32(0x5D4594, 831080) = v
}

var (
	Dialogs *dialog.Dialog
)

func initDialog() {
	Dialogs = dialog.NewDialog(
		"dialog",
		(*uint32)(&C.dword_587000_122848),
		(*uint32)(&C.dword_5d4594_830864), (*uint32)(&C.dword_5d4594_831084), (*uint32)(&C.dword_587000_122856), (*uint32)(&C.dword_5d4594_831076), (*ail.Driver)(unsafe.Pointer(&C.dword_5d4594_831092)),
		memmap.PtrUint32(0x5D4594, 830860),
		func() *strman.StringManager {
			return GetServer().S().Strings()
		},
		memmap.PtrT[timer.TimerGroup](0x5D4594, 830876),
		memmap.PtrT[timer.TimerGroup](0x5D4594, 830980),
		(*timer.TimerGroup)(C.dword_587000_81128),
		(*timer.TimerGroup)(C.dword_587000_122852),
		Sub_413890,

		func() ail.Driver {
			return Sub_43F130()
		},
		Sub_43DBD0,
		Sub_43DBE0,
		Sub_43DC40,
		Set_dword_5d4594_831080,
		func() uint32 { return uint32(Get_dword_587000_93160()) },
	)
}

//export sub_44D930
func sub_44D930() int {
	return bool2int(Dialogs.Sub_44D930())
}

//export nox_xxx_playDialogFile_44D900
func nox_xxx_playDialogFile_44D900(a1p *byte, a2 int) int {
	return Dialogs.PlayFile(alloc.GoString(a1p), a2)
}
