package legacy

/*
#include "GAME1.h"
#include "GAME1_2.h"
#include "GAME1_3.h"
#include "GAME2.h"
#include "GAME2_2.h"
#include "GAME3_2.h"
#include "GAME5.h"
#include "common__net_list.h"
#include "client__system__parsecmd.h"
#include "common__object__armrlook.h"
#include "common__object__weaplook.h"
#include "common__log.h"

extern unsigned int dword_5d4594_2650652;
extern void* dword_587000_81128;
extern unsigned int dword_587000_93156;
extern unsigned int dword_5d4594_816340;
extern unsigned int dword_5d4594_816348;

extern nox_gui_animation* nox_wnd_xxx_1309740;
*/
import "C"

var (
	Nox_exit                                              func(exitCode int)
	Nox_xxx_gameGetScreenBoundaries_43BEB0_get_video_mode func(w, h, d *int)
	Sub_4AA9C0                                            func() int
)

//export nox_exit
func nox_exit(exitCode int) { Nox_exit(exitCode) }

//export nox_xxx_gameGetScreenBoundaries_43BEB0_get_video_mode
func nox_xxx_gameGetScreenBoundaries_43BEB0_get_video_mode(w, h, d *int) {
	Nox_xxx_gameGetScreenBoundaries_43BEB0_get_video_mode(w, h, d)
}

//export sub_4AA9C0
func sub_4AA9C0() int { return Sub_4AA9C0() }

func Nox_xxx_loadLook_415D50() {
	C.nox_xxx_loadLook_415D50()
}

func Nox_xxx_loadModifyers_4158C0() {
	C.nox_xxx_loadModifyers_4158C0()
}

func Sub_4D11A0() {
	C.sub_4D11A0()
}

func Sub_431370() int {
	return int(C.sub_431370())
}

func Nox_xxx_tileAlloc_410F60_init() int {
	return int(C.nox_xxx_tileAlloc_410F60_init())
}

func Nox_xxx_initSinCosTables_414C90() {
	C.nox_xxx_initSinCosTables_414C90()
}

func Nox_xxx_loadMapCycle_4D0A30() {
	C.nox_xxx_loadMapCycle_4D0A30()
}

func Nox_xxx_mapSelectFirst_4D0E00() {
	C.nox_xxx_mapSelectFirst_4D0E00()
}

func Sub_4134D0() {
	C.sub_4134D0()
}

func Sub_413920() {
	C.sub_413920()
}

func Sub_431380() {
	C.sub_431380()
}

func Nox_xxx_tileFree_410FC0_free() {
	C.nox_xxx_tileFree_410FC0_free()
}

func Sub_42EDC0() {
	C.sub_42EDC0()
}

func Sub_4D11D0() {
	C.sub_4D11D0()
}

func Sub_4D0DA0() {
	C.sub_4D0DA0()
}

func Nox_common_maplist_free_4D0970() {
	C.nox_common_maplist_free_4D0970()
}

func Sub_451970() {
	C.sub_451970()
}

func Sub_431270() {
	C.sub_431270()
}

func Sub_4875F0() {
	C.sub_4875F0()
}

func Sub_4870A0() {
	C.sub_4870A0()
}

func Sub_431290() {
	C.sub_431290()
}

func Sub_43D990() {
	C.sub_43D990()
}

func Nox_xxx_servSetPlrLimit_409F80(v int) {
	C.nox_xxx_servSetPlrLimit_409F80(C.int(v))
}

func Nox_xxx_guiChatShowHide_445730(v bool) {
	C.nox_xxx_guiChatShowHide_445730(C.int(bool2int(v)))
}
