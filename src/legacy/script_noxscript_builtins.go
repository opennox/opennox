package legacy

/*
#include "defs.h"
int nox_script_SetQuestInt_514BE0();
int nox_script_SetQuestFloat_514C10();
int nox_script_GetQuestInt_514C40();
int nox_script_GetQuestFloat_514C60();
int nox_script_ResetQuestStatus_514C90();
int nox_script_SetRoamFlag_515C40();
int nox_script_SetRoamFlagGroup_515CB0();
int nox_script_JournalDelete_515550();
int nox_script_JournalEdit_5155A0();
int nox_script_RetreatLevel_515DF0();
int nox_script_RetreatLevelGroup_515E50();
int nox_script_SetResumeLevel_515E80();
int nox_script_SetResumeLevelGroup_515EE0();
int nox_script_GiveExp_516190();
int nox_script_IsTalking_5166A0();
int nox_script_MakeFriendly_516720();
int nox_script_MakeEnemy_516760();
int nox_script_BecomePet_5167D0();
int nox_script_BecomeEnemy_516810();
int nox_script_builtin_516790();
int nox_script_builtin_516850();
int nox_script_OblivionGive_516890();
int nox_script_PlayerIsTrading_5166E0();
void nox_script_StartupScreen_516600_A();
int sub_512E80(wchar2_t* a1);
*/
import "C"
import (
	"unsafe"

	"github.com/opennox/noxscript/ns/asm"

	"github.com/opennox/opennox/v1/legacy/common/ccall"
	"github.com/opennox/opennox/v1/server/noxscript"
)

var (
	Nox_script_shouldReadMoreXxx     func(fi asm.Builtin) bool
	Nox_script_shouldReadEvenMoreXxx func(fi asm.Builtin) bool
)

//export nox_script_shouldReadMoreXxx
func nox_script_shouldReadMoreXxx(fi int) C.bool {
	return C.bool(Nox_script_shouldReadMoreXxx(asm.Builtin(fi)))
}

//export nox_script_shouldReadEvenMoreXxx
func nox_script_shouldReadEvenMoreXxx(fi int) C.bool {
	return C.bool(Nox_script_shouldReadEvenMoreXxx(asm.Builtin(fi)))
}

func wrapScriptC(fnc unsafe.Pointer) noxscript.Builtin {
	return func(_ noxscript.VM) int {
		return ccall.CallIntVoid(fnc)
	}
}

func CallScriptBuiltin(fi asm.Builtin) (int, bool) {
	if fi < 0 || int(fi) >= len(noxScriptBuiltins) {
		return 0, false
	}
	fnc := noxScriptBuiltins[fi]
	if fnc == nil {
		return 0, false
	}
	res := fnc(GetServer().NoxScriptC())
	return res, true
}

func Nox_script_StartupScreen_516600_A() {
	C.nox_script_StartupScreen_516600_A()
}

func Sub_512E80(str string) int {
	cstr, _ := CWString(str)
	return int(C.sub_512E80(cstr))
}

var noxScriptBuiltins = [asm.BuiltinGetScore + 1]noxscript.Builtin{
	asm.BuiltinSetQuestStatus:      wrapScriptC(C.nox_script_SetQuestInt_514BE0),
	asm.BuiltinSetQuestStatusFloat: wrapScriptC(C.nox_script_SetQuestFloat_514C10),
	asm.BuiltinGetQuestStatus:      wrapScriptC(C.nox_script_GetQuestInt_514C40),
	asm.BuiltinGetQuestStatusFloat: wrapScriptC(C.nox_script_GetQuestFloat_514C60),
	asm.BuiltinResetQuestStatus:    wrapScriptC(C.nox_script_ResetQuestStatus_514C90),
	asm.BuiltinSetRoamFlag:         wrapScriptC(C.nox_script_SetRoamFlag_515C40),
	asm.BuiltinGroupSetRoamFlag:    wrapScriptC(C.nox_script_SetRoamFlagGroup_515CB0),
	asm.BuiltinJournalDelete:       wrapScriptC(C.nox_script_JournalDelete_515550),
	asm.BuiltinJournalEdit:         wrapScriptC(C.nox_script_JournalEdit_5155A0),
	asm.BuiltinGiveXp:              wrapScriptC(C.nox_script_GiveExp_516190),
	asm.BuiltinIsTalking:           wrapScriptC(C.nox_script_IsTalking_5166A0),
	asm.BuiltinMakeFriendly:        wrapScriptC(C.nox_script_MakeFriendly_516720),
	asm.BuiltinMakeEnemy:           wrapScriptC(C.nox_script_MakeEnemy_516760),
	asm.BuiltinBecomePet:           wrapScriptC(C.nox_script_BecomePet_5167D0),
	asm.BuiltinBecomeEnemy:         wrapScriptC(C.nox_script_BecomeEnemy_516810),
	asm.BuiltinUnknownb8:           wrapScriptC(C.nox_script_builtin_516790),
	asm.BuiltinUnknownb9:           wrapScriptC(C.nox_script_builtin_516850),
	asm.BuiltinSetHalberd:          wrapScriptC(C.nox_script_OblivionGive_516890),
	asm.BuiltinIsTrading:           wrapScriptC(C.nox_script_PlayerIsTrading_5166E0),
}
