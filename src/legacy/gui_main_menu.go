package legacy

/*
#include "GAME1_2.h"
#include "GAME1_3.h"
#include "GAME2.h"
#include "GAME2_1.h"
#include "GAME3.h"
#include "GAME3_2.h"
#include "GAME4.h"
#include "GAME4_1.h"
#include "GAME5_2.h"
#include "client__shell__noxworld.h"
#include "client__shell__selchar.h"
#include "client__shell__mainmenu.h"
extern unsigned int dword_5d4594_831220;
extern uint32_t dword_5d4594_831260;
int winMainMenuAnimOutStartFnc();
int winMainMenuAnimOutDoneFnc();
int nox_client_drawGeneralCallback_4A2200();
*/
import "C"
import (
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/client/gui"
)

var (
	WinMainMenuAnimOutStartFnc   func() int
	Sub_44E320                   func()
	WinMainMenuAnimOutDoneFnc    func() int
	Sub_4A24C0                   func(a1 int) int
	Nox_game_showMainMenu_4A1C00 func() int
	Sub_43BE40                   func(a1 int)
	Sub_43BE30                   func() int
	Sub_4A18E0                   func(a1 *gui.Window, a2, a3, a4 int) int
)

//export winMainMenuAnimOutStartFnc
func winMainMenuAnimOutStartFnc() int { return WinMainMenuAnimOutStartFnc() }

//export sub_44E320
func sub_44E320() { Sub_44E320() }

//export winMainMenuAnimOutDoneFnc
func winMainMenuAnimOutDoneFnc() int { return WinMainMenuAnimOutDoneFnc() }

//export sub_4A24C0
func sub_4A24C0(a1 int) int { return Sub_4A24C0(a1) }

//export nox_game_showMainMenu_4A1C00
func nox_game_showMainMenu_4A1C00() int { return Nox_game_showMainMenu_4A1C00() }

//export sub_43BE40
func sub_43BE40(a1 int) { Sub_43BE40(a1) }

//export sub_43BE30
func sub_43BE30() int { return Sub_43BE30() }

//export sub_4A18E0
func sub_4A18E0(a1 unsafe.Pointer, a2, a3, a4 int) int { return Sub_4A18E0(AsWindowP(a1), a2, a3, a4) }

//export nox_client_drawGeneralCallback_4A2200
func nox_client_drawGeneralCallback_4A2200() int {
	if err := GetClient().DrawGeneral(false); err != nil {
		guiLog.Println(err)
		return 0
	}
	return 1
}

func Get_nox_client_drawGeneralCallback_4A2200() unsafe.Pointer {
	return unsafe.Pointer(C.nox_client_drawGeneralCallback_4A2200)
}

func Get_nox_game_showGameSel_4379F0() unsafe.Pointer {
	return unsafe.Pointer(C.nox_game_showGameSel_4379F0)
}

func Sub_461440(v int) {
	C.sub_461440(C.int(v))
}

func Sub_4D7440(v int) {
	C.sub_4D7440(C.int(v))
}

func Nox_xxx_cliSetMinimapZoom_472520(v int) {
	C.nox_xxx_cliSetMinimapZoom_472520(C.int(v))
}
func Nox_xxx_monsterListFree_5174F0() {
	C.nox_xxx_monsterListFree_5174F0()
}
func Nox_client_countSaveFiles_4DC550() int {
	return int(C.nox_client_countSaveFiles_4DC550())
}
func Sub_4A7A60(a1 int) {
	C.sub_4A7A60(C.int(a1))
}
func Nox_client_lockScreenBriefing_450160(a1 int, a2 int, a3 int) {
	C.nox_client_lockScreenBriefing_450160(C.int(a1), C.int(a2), C.char(a3))
}
func Sub_43E8C0(a1 int) {
	C.sub_43E8C0(C.int(a1))
}
func Sub_43D9B0(a1 int, a2 int) {
	C.sub_43D9B0(C.int(a1), C.int(a2))
}
func Sub_4D6F80(a1 int) {
	C.sub_4D6F80(C.int(a1))
}
func Sub_4A7A70(a1 int) {
	C.sub_4A7A70(C.int(a1))
}

func Get_nox_game_showSelChar_4A4DB0() unsafe.Pointer {
	return C.nox_game_showSelChar_4A4DB0
}
func Get_nox_game_showSelClass_4A4840() unsafe.Pointer {
	return C.nox_game_showSelClass_4A4840
}
func Get_winMainMenuAnimOutStartFnc() unsafe.Pointer {
	return C.winMainMenuAnimOutStartFnc
}
func Get_winMainMenuAnimOutDoneFnc() unsafe.Pointer {
	return C.winMainMenuAnimOutDoneFnc
}
func Get_sub_4A22A0() unsafe.Pointer {
	return C.sub_4A22A0
}
