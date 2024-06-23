package legacy

/*
#include "GAME1.h"
#include "GAME1_1.h"
#include "GAME1_2.h"
#include "GAME1_3.h"
#include "GAME2.h"
#include "GAME2_1.h"
#include "GAME2_2.h"
#include "GAME2_3.h"
#include "GAME3.h"
#include "GAME3_1.h"
#include "GAME3_2.h"
#include "GAME4.h"
#include "client__shell__optsback.h"
#include "client__gui__servopts__guiserv.h"
#include "client__shell__noxworld.h"
#include "client__gui__guicon.h"
#include "client__gui__guibook.h"
#include "client__gui__guisave.h"
#include "client__shell__selchar.h"
#include "client__shell__mainmenu.h"
*/
import "C"
import "unsafe"

var (
	Sub_4A1BE0                             func(a1 int) int
	Sub_41E300                             func(a1 int) int
	Sub_4A50A0                             func() int
	Sub_4A50D0                             func() int
	Nox_client_resetScreenParticles_431510 func()
	Sub_46D6F0                             func() int
	Sub_413A00                             func(a1 int)
	Sub_44A400                             func()
	Nox_game_showSelChar_4A4DB0            func() int
	Nox_savegame_sub_46D580                func()
	Sub_450580                             func()
	Sub_4DB170                             func(a1 bool, a2 unsafe.Pointer, a3 int)
)

var _ = [1]struct{}{}[1277-unsafe.Offsetof(Nox_savegame_xxx{}.Stage)]
var _ = [1]struct{}{}[1280-unsafe.Sizeof(Nox_savegame_xxx{})] // FIXME: should be 1278

type Nox_savegame_xxx struct {
	Flags        uint32     // 0
	Path         [1024]byte // 4
	Field_1028   [128]byte  // 1028
	Map_name     [32]byte   // 1156
	Timestamp    SYSTEMTIME // 1188
	Field_1204   [20]byte   // 1204
	Player_name  [25]uint16 // 1224
	Player_class byte       // 1274
	Field_1275   byte       // 1275
	Field_1276   byte       // 1276
	Stage        byte       // 1277
}

var _ = [1]struct{}{}[52-unsafe.Sizeof(Nox_screenParticle{})]

type nox_screenParticle = C.nox_screenParticle
type Nox_screenParticle struct {
	Draw_fnc unsafe.Pointer      // 0, 0, func(unsafe.Pointer, *Nox_screenParticle) int
	Field_4  uint32              // 1, 4
	Field_8  uint32              // 2, 8
	Field_12 uint32              // 3, 12, color
	Field_16 uint32              // 4, 16, vx
	Field_20 uint32              // 5, 20, vy
	Field_24 uint32              // 6, 24, x
	Field_28 uint32              // 7, 28, y
	Field_32 uint32              // 8, 32
	Field_36 uint32              // 9, 36
	Field_40 [4]byte             // 10, 40
	Field_44 *Nox_screenParticle // 11, 44, next
	Field_48 *Nox_screenParticle // 12, 48, prev
}

//export sub_4A1BE0
func sub_4A1BE0(a1 int) int { return Sub_4A1BE0(a1) }

//export sub_41E300
func sub_41E300(a1 int) int { return Sub_41E300(a1) }

//export sub_4A50A0
func sub_4A50A0() int { return Sub_4A50A0() }

//export sub_4A50D0
func sub_4A50D0() int { return Sub_4A50D0() }

//export nox_client_resetScreenParticles_431510
func nox_client_resetScreenParticles_431510() { Nox_client_resetScreenParticles_431510() }

//export sub_46D6F0
func sub_46D6F0() int { return Sub_46D6F0() }

//export sub_413A00
func sub_413A00(a1 int) { Sub_413A00(a1) }

//export sub_44A400
func sub_44A400() { Sub_44A400() }

//export nox_game_showSelChar_4A4DB0
func nox_game_showSelChar_4A4DB0() int { return Nox_game_showSelChar_4A4DB0() }

//export nox_savegame_sub_46D580
func nox_savegame_sub_46D580() { Nox_savegame_sub_46D580() }

//export sub_450580
func sub_450580() { Sub_450580() }

//export sub_4DB170
func sub_4DB170(a1 int, a2 unsafe.Pointer, a3 int) { Sub_4DB170(a1 != 0, a2, a3) }

func Sub_41E470() {
	C.sub_41E470()
}
func Nox_xxx_reconStart_41E400() {
	C.nox_xxx_reconStart_41E400()
}
func Sub_467C10() int {
	return int(C.sub_467C10())
}
func Nox_xxx_guiServerOptionsTryHide_4574D0() int {
	return int(C.nox_xxx_guiServerOptionsTryHide_4574D0())
}
func Sub_446780() {
	C.sub_446780()
}
func Sub_460660() {
	C.sub_460660()
}
func Sub_467BB0() {
	C.sub_467BB0()
}
func Sub_4602F0() {
	C.sub_4602F0()
}
func Nox_xxx_playerAnimCheck_4372B0() int {
	return int(C.nox_xxx_playerAnimCheck_4372B0())
}
func Sub_41DA10(a1 int) {
	C.sub_41DA10(C.int(a1))
}
func Nox_xxx_bookHideMB_45ACA0(a1 int) int {
	return int(C.nox_xxx_bookHideMB_45ACA0(C.int(a1)))
}
func Sub_4AD9B0(a1 int) {
	C.sub_4AD9B0(C.int(a1))
}
func Nox_xxx_bookShowMB_45AD70(a1 int) {
	C.nox_xxx_bookShowMB_45AD70(C.int(a1))
}
func Sub_41A000(a1 string, a2 *Nox_savegame_xxx) int {
	cstr := CString(a1)
	defer StrFree(cstr)
	return int(C.sub_41A000(cstr, (*C.nox_savegame_xxx)(unsafe.Pointer(a2))))
}
func Sub_43D9E0(a1 unsafe.Pointer) {
	C.sub_43D9E0((*C.int4)(a1))
}

func Get_sub_4A18E0() unsafe.Pointer {
	return C.sub_4A18E0
}
func Get_sub_4A50A0() unsafe.Pointer {
	return C.sub_4A50A0
}
func Get_sub_4A50D0() unsafe.Pointer {
	return C.sub_4A50D0
}
