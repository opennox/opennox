package legacy

/*
#include "GAME1.h"
#include "GAME1_1.h"
#include "GAME1_3.h"
#include "GAME2.h"
#include "GAME2_1.h"
#include "GAME3_3.h"
#include "GAME4_2.h"
extern uint32_t dword_5d4594_1049844;
extern uint32_t dword_5d4594_1563096;
void nox_xxx_monstersAllBelongToHost_4DB6A0();
void nox_xxx_unitsNewAddToList_4DAC00();
int sub_41C280(void* a1);
int nox_xxx_parseFileInfoData_41C3B0(int a1);
int sub_41C780(int a1);
*/
import "C"
import (
	"errors"
	"unsafe"

	"github.com/opennox/opennox/v1/internal/cryptfile"
	"github.com/opennox/opennox/v1/server"
)

var (
	Nox_savegame_rm                      func(name string, rmDir bool) error
	Nox_client_countPlayerFiles04_4DC7D0 func() int
	Nox_xxx_gameGet_4DB1B0               func() bool
	Sub_4DCC90                           func() int
	Sub_4DB1C0                           func() unsafe.Pointer
	Sub_4DCBF0                           func(a1 int)
	Nox_xxx_serverIsClosing_446180       func() int
	Sub_4DCC10                           func(a1p *server.Object) int
	Sub_4DCFB0                           func(a1p *server.Object)
	Sub_4DD0B0                           func(a1p *server.Object)
	Nox_setSaveFileName_4DB130           func(s string)
)

//export nox_setSaveFileName_4DB130
func nox_setSaveFileName_4DB130(s *C.char) {
	Nox_setSaveFileName_4DB130(GoString(s))
}

//export nox_savegame_rm_4DBE10
func nox_savegame_rm_4DBE10(cname *C.char, rmDir int) {
	if cname == nil {
		return
	}
	saveName := GoString(cname)
	Nox_savegame_rm(saveName, rmDir != 0)
}

//export nox_client_countPlayerFiles04_4DC7D0
func nox_client_countPlayerFiles04_4DC7D0() int { return Nox_client_countPlayerFiles04_4DC7D0() }

//export nox_xxx_gameGet_4DB1B0
func nox_xxx_gameGet_4DB1B0() int { return bool2int(Nox_xxx_gameGet_4DB1B0()) }

//export sub_4DCC90
func sub_4DCC90() int { return Sub_4DCC90() }

//export sub_4DB1C0
func sub_4DB1C0() unsafe.Pointer { return Sub_4DB1C0() }

//export sub_4DCBF0
func sub_4DCBF0(a1 int) { Sub_4DCBF0(a1) }

//export nox_xxx_serverIsClosing_446180
func nox_xxx_serverIsClosing_446180() int { return Nox_xxx_serverIsClosing_446180() }

//export sub_4DCC10
func sub_4DCC10(a1p *nox_object_t) int { return Sub_4DCC10(asObjectS(a1p)) }

//export sub_4DCFB0
func sub_4DCFB0(a1p *nox_object_t) { Sub_4DCFB0(asObjectS(a1p)) }

//export sub_4DD0B0
func sub_4DD0B0(a1p *nox_object_t) { Sub_4DD0B0(asObjectS(a1p)) }
func Nox_xxx_destroyEveryChatMB_528D60() {
	C.nox_xxx_destroyEveryChatMB_528D60()
}
func Nox_xxx_quickBarClose_4606B0() {
	C.nox_xxx_quickBarClose_4606B0()
}
func Nox_xxx_monstersAllBelongToHost_4DB6A0() {
	C.nox_xxx_monstersAllBelongToHost_4DB6A0()
}
func Nox_xxx_mapSaveMap_51E010(a1 string, a2 int) bool {
	return C.nox_xxx_mapSaveMap_51E010(internCStr(a1), C.int(a2)) != 0
}

func Sub_41A590(cf *cryptfile.CryptFile, u *server.Object, pinfo *server.PlayerInfo) error {
	old := cryptfile.Global()
	cryptfile.SetGlobal(cf)
	defer cryptfile.SetGlobal(old)
	if C.sub_41A590(u.CObj(), pinfo.C()) == 0 {
		return errors.New("failed")
	}
	return nil
}

func Sub_41AA30(cf *cryptfile.CryptFile, u *server.Object, pinfo *server.PlayerInfo) error {
	old := cryptfile.Global()
	cryptfile.SetGlobal(cf)
	defer cryptfile.SetGlobal(old)
	if C.sub_41AA30(u.CObj(), pinfo.C()) == 0 {
		return errors.New("failed")
	}
	return nil
}

func Sub_41AC30(cf *cryptfile.CryptFile, u *server.Object, pinfo *server.PlayerInfo) error {
	old := cryptfile.Global()
	cryptfile.SetGlobal(cf)
	defer cryptfile.SetGlobal(old)
	if C.sub_41AC30(u.CObj(), pinfo.C()) == 0 {
		return errors.New("failed")
	}
	return nil
}

func Nox_xxx_guiFieldbook_41B420(cf *cryptfile.CryptFile, u *server.Object, pinfo *server.PlayerInfo) error {
	old := cryptfile.Global()
	cryptfile.SetGlobal(cf)
	defer cryptfile.SetGlobal(old)
	if C.nox_xxx_guiFieldbook_41B420(u.CObj(), pinfo.C()) == 0 {
		return errors.New("failed")
	}
	return nil
}

func Nox_xxx_guiSpellbook_41B660(cf *cryptfile.CryptFile, u *server.Object, pinfo *server.PlayerInfo) error {
	old := cryptfile.Global()
	cryptfile.SetGlobal(cf)
	defer cryptfile.SetGlobal(old)
	if C.nox_xxx_guiSpellbook_41B660(u.CObj(), pinfo.C()) == 0 {
		return errors.New("failed")
	}
	return nil
}

func Nox_xxx_guiEnchantment_41B9C0(cf *cryptfile.CryptFile, u *server.Object, pinfo *server.PlayerInfo) error {
	old := cryptfile.Global()
	cryptfile.SetGlobal(cf)
	defer cryptfile.SetGlobal(old)
	if C.nox_xxx_guiEnchantment_41B9C0(u.CObj(), pinfo.C()) == 0 {
		return errors.New("failed")
	}
	return nil
}

func Sub_41BEC0(cf *cryptfile.CryptFile, u *server.Object, pinfo *server.PlayerInfo) error {
	old := cryptfile.Global()
	cryptfile.SetGlobal(cf)
	defer cryptfile.SetGlobal(old)
	if C.sub_41BEC0(u.CObj(), pinfo.C()) == 0 {
		return errors.New("failed")
	}
	return nil
}

func Sub_41C080(cf *cryptfile.CryptFile, u *server.Object, pinfo *server.PlayerInfo) error {
	old := cryptfile.Global()
	cryptfile.SetGlobal(cf)
	defer cryptfile.SetGlobal(old)
	if C.sub_41C080(u.CObj(), pinfo.C()) == 0 {
		return errors.New("failed")
	}
	return nil
}

func Sub_41C280(cf *cryptfile.CryptFile) error {
	old := cryptfile.Global()
	cryptfile.SetGlobal(cf)
	defer cryptfile.SetGlobal(old)
	if C.sub_41C280(nil) == 0 {
		return errors.New("failed")
	}
	return nil
}

func Nox_xxx_parseFileInfoData_41C3B0(cf *cryptfile.CryptFile) error {
	old := cryptfile.Global()
	cryptfile.SetGlobal(cf)
	defer cryptfile.SetGlobal(old)
	if C.nox_xxx_parseFileInfoData_41C3B0(0) == 0 {
		return errors.New("failed")
	}
	return nil
}

func Sub_41C780(cf *cryptfile.CryptFile) error {
	old := cryptfile.Global()
	cryptfile.SetGlobal(cf)
	defer cryptfile.SetGlobal(old)
	if C.sub_41C780(0) == 0 {
		return errors.New("failed")
	}
	return nil
}
