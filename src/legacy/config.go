package legacy

/*
#include <stdint.h>
#include "GAME1.h"
#include "GAME1_3.h"
#include "GAME2.h"
#include "GAME2_2.h"
#include "GAME3_2.h"
#include "GAME5_2.h"
#include "common__system__settings.h"
extern void* dword_587000_127004;
extern unsigned int dword_5d4594_1193156;
extern unsigned int nox_client_translucentFrontWalls_805844;
extern unsigned int nox_client_highResFrontWalls_80820;
extern unsigned int nox_client_highResFloors_154952;
extern unsigned int nox_client_fadeObjects_80836;
extern unsigned int nox_client_renderBubbles_80844;
extern uint32_t nox_server_connectionType_3596;
extern uint32_t nox_server_kickQuestPlayerMinVotes_229992;
extern uint32_t nox_server_resetQuestMinVotes_229988;
extern uint32_t nox_server_sanctuaryHelp_54276;
extern uint32_t nox_server_sendMotd_108752;
*/
import "C"

var (
	WriteConfigLegacy func(name string)
)

//export nox_common_writecfgfile
func nox_common_writecfgfile(str *C.char) {
	WriteConfigLegacy(GoString(str))
}
