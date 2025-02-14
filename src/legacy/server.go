package legacy

/*
#include "GAME1.h"
#include "GAME1_2.h"
#include "GAME1_1.h"
#include "GAME1_3.h"
#include "GAME2.h"
#include "GAME2_1.h"
#include "GAME2_2.h"
#include "GAME3_2.h"
#include "GAME3_3.h"
#include "GAME4.h"
#include "GAME4_1.h"
#include "GAME5.h"
#include "GAME5_2.h"
#include "client__gui__guiquit.h"
#include "common__system__team.h"
#include "server__system__server.h"
#include "server__script__script.h"
#include "server__script__activator.h"
#include "common__magic__speltree.h"
#include "common__net_list.h"
#include "common__crypt.h"
#include "common__log.h"

extern unsigned int dword_5d4594_2650652;
extern unsigned int dword_5d4594_2649712;
extern unsigned int dword_5d4594_1548524;
extern uint32_t dword_5d4594_1563096;
extern uint32_t dword_5d4594_528252;
extern uint32_t dword_5d4594_528260;
extern uint32_t dword_5d4594_2488604;

extern uint32_t nox_tile_def_cnt;
extern nox_tileDef_t nox_tile_defs_arr[176];

void nox_xxx_netlist_4DEB50();
void nox_xxx_updateUnits_51B100();
void nox_xxx_voteUptate_506F30();
void sub_4E4170();
void sub_4EC720();
unsigned int sub_50D890();
void nox_xxx_gameTick_4D2580_server_D();
int nox_xxx_netUpdateObjectSpecial_527E50(nox_object_t* a1p, nox_object_t* a2p);
void sub_4139C0();
int sub_4DCF20();
int sub_4E76C0();
bool sub_57B140();
nox_object_t* sub_537700();

void nox_xxx_updateUnits_51B100_D();
*/
import "C"
import (
	"image"
	"unsafe"

	"github.com/opennox/libs/object"
	"github.com/opennox/libs/spell"
	"github.com/opennox/libs/strman"
	"github.com/opennox/libs/types"

	noxflags "github.com/opennox/opennox/v1/common/flags"
	"github.com/opennox/opennox/v1/common/ntype"
	"github.com/opennox/opennox/v1/server"
)

var (
	Nox_mapToGameFlags func(v int) noxflags.GameFlag
	Sub_40A1A0         func() int
)

type Server interface {
	S() *server.Server
	SetUpdateFunc2(fnc func() bool)
	ServerPort() int
	ServStartCountdown(dt int, text strman.ID)
	SwitchMap(name string)
	GetFlag3592() bool
	Sub40A040settings(a1 int, a2 int)
	CreateObjectAt(obj, owner server.Obj, pt types.Pointf)
	NetUpdateRemotePlrAudioEvents(obj *server.Object, v2 unsafe.Pointer, v18 int8)
	Nox_xxx_mapDamageUnitsAround(pos types.Pointf, r1, r2 float32, dmg int, dtyp object.DamageType, who *server.Object, a7 server.Obj, damageWalls bool)
	Nox_xxx_mapReset5028E0()
	Nox_xxx_free503F40()
	ObjectDeleteLast(obj *server.Object)
	ObjectsAddPending()
	FinalizeDeletingObjects()
	TeamsResetYyy() int
	TeamsRemoveActive(hooks bool) int
	TeamRemove(t *server.Team, netUpd bool)
	DelayedDelete(obj *server.Object)
	Sub504720(a1, a2 uint32) int32
	ApplyForce(obj *server.Object, vec types.Pointf, force float64)
	PlayerSpell(u *server.Object)
	Nox_script_event_playerLeave(pl *server.Player)
	NoxScriptC() NoxScript
	Nox_xxx_spellAccept4FD400(spellID spell.ID, a2, obj3, obj4 *server.Object, sa *server.SpellAcceptArg, lvl int) bool
	Nox_xxx_generateRetreatPath_50CA00(path []types.Pointf, u *server.Object, a4 *types.Pointf) int
	Nox_xxx_creatureSetDetailedPath_50D220(obj *server.Object, a2 *types.Pointf)
	Sub_50CB20(a1 *server.Object, a2 *types.Pointf) *server.Waypoint
	Sub_50B810(obj *server.Object, pos *types.Pointf) bool
	Nox_xxx_mapDamageToWalls_534FC0(rect image.Rectangle, pos types.Pointf, rad float32, dmg int, dtyp object.DamageType, who *server.Object) bool
	Nox_xxx_damageToMap_534BC0(gx, gy int, dmg int, dtyp object.DamageType, who *server.Object) int
	Nox_xxx_wall_4DF1E0(a1 int)
}

var (
	GetServer func() Server
)

//export nox_getHostPlayerUnit
func nox_getHostPlayerUnit() *nox_object_t {
	return asObjectC(GetServer().S().Players.HostUnit())
}

//export nox_xxx_servStartCountdown_40A2A0
func nox_xxx_servStartCountdown_40A2A0(a1 int, a2 *C.char) {
	GetServer().ServStartCountdown(a1, strman.ID(GoString(a2)))
}

//export sub_40A040_settings
func sub_40A040_settings(a1 C.short, a2 C.uchar) {
	GetServer().Sub40A040settings(int(a1), int(a2))
}

//export nox_server_SetLastObjectScriptID
func nox_server_SetLastObjectScriptID(id C.uint) {
	GetServer().S().Objs.SetLastObjectScriptID(server.ObjectScriptID(id))
}

//export nox_server_LastObjectScriptID
func nox_server_LastObjectScriptID() C.uint {
	return C.uint(GetServer().S().Objs.LastObjectScriptID())
}

//export nox_server_NextObjectScriptID
func nox_server_NextObjectScriptID() C.uint {
	return C.uint(GetServer().S().Objs.NextObjectScriptID())
}

//export nox_xxx_servGetPort_40A430
func nox_xxx_servGetPort_40A430() int {
	return GetServer().ServerPort()
}

//export sub_40A300
func sub_40A300() int {
	return bool2int(GetServer().GetFlag3592())
}

//export nox_xxx_mapLoad_4D2450
func nox_xxx_mapLoad_4D2450(a1 *C.char) {
	GetServer().SwitchMap(GoString(a1))
}

//export nox_mapToGameFlags_4CFF50
func nox_mapToGameFlags_4CFF50(v int) int {
	return int(Nox_mapToGameFlags(v))
}

//export sub_40A1A0
func sub_40A1A0() int {
	return Sub_40A1A0()
}

//export gameFPS
func gameFPS() uint32 {
	return GetServer().S().TickRate()
}

//export gameFrame
func gameFrame() uint32 {
	return GetServer().S().Frame()
}
func Sub_409A70(a1 int) int {
	return int(C.sub_409A70(C.short(a1)))
}
func Nox_xxx_netInformTextMsg2_4DA180(a1 int, a2 unsafe.Pointer) {
	C.nox_xxx_netInformTextMsg2_4DA180(C.int(a1), (*C.uchar)(a2))
}
func Nox_xxx_netReportUnitHeight_4D9020(a1 ntype.PlayerInd, a2 *server.Object) {
	C.nox_xxx_netReportUnitHeight_4D9020(C.int(a1), asObjectC(a2))
}
func Sub_511100(a1 int) {
	C.sub_511100(C.int(a1))
}
func Nox_xxx_netUpdateRemotePlr_501CA0(a1 *server.Object) {
	C.nox_xxx_netUpdateRemotePlr_501CA0(asObjectC(a1))
}
func Nox_xxx_netSendObjects2Plr_519410(a1 *server.Object, a2 *server.Object) {
	C.nox_xxx_netSendObjects2Plr_519410(asObjectC(a1), asObjectC(a2))
}
func Sub_4D6770(a1 ntype.PlayerInd) {
	C.sub_4D6770(C.int(a1))
}
func Sub_4D6880(a1 int, a2 int) {
	C.sub_4D6880(C.int(a1), C.int(a2))
}
func Sub_4D60B0() {
	C.sub_4D60B0()
}
func Sub_4CFDF0(a1 int) {
	C.sub_4CFDF0(C.int(a1))
}
func Nox_xxx_playerMapTracksObj_4173D0(a1 int, a2 *server.Object) int {
	return int(C.nox_xxx_playerMapTracksObj_4173D0(C.int(a1), asObjectC(a2)))
}
func Sub_519710(a1 unsafe.Pointer) int {
	return int(C.sub_519710(a1))
}
func Nox_xxx_updateUnits_51B100_D() {
	C.nox_xxx_updateUnits_51B100_D()
}
func Nox_xxx_decay_511750() {
	C.nox_xxx_decay_511750()
}
func Nox_server_checkVictory_509A60() {
	C.nox_server_checkVictory_509A60()
}
func Nox_xxx_allocHitArray_5486D0() {
	C.nox_xxx_allocHitArray_5486D0()
}
func Nox_xxx_updateObjectsVelocity_5118A0(a1 float32) {
	C.nox_xxx_updateObjectsVelocity_5118A0(C.float(a1))
}
func Sub_548B60() {
	C.sub_548B60()
}
func Sub_537700() *server.Object {
	return asObjectS(C.sub_537700())
}
func Nox_xxx_collide_548740() {
	C.nox_xxx_collide_548740()
}
func Nox_xxx_updatePoison_4EE8F0(a1 *server.Object, a2 int) {
	C.nox_xxx_updatePoison_4EE8F0(asObjectC(a1), C.int(a2))
}
func Nox_xxx_updateUnitBuffs_4FF620(a1 *server.Object) {
	C.nox_xxx_updateUnitBuffs_4FF620(asObjectC(a1))
}
func Nox_xxx_reconAttempt_41E390() {
	C.nox_xxx_reconAttempt_41E390()
}
func Sub_5096F0() {
	C.sub_5096F0()
}
func Sub_416720() {
	C.sub_416720()
}
func Sub_40AA00() int {
	return int(C.sub_40AA00())
}
func Sub_40AA40() int {
	return int(C.sub_40AA40())
}
func Nox_xxx_countNonEliminatedPlayersInTeam_40A830(a1 *server.Team) int {
	return int(C.nox_xxx_countNonEliminatedPlayersInTeam_40A830((*nox_team_t)(a1.C())))
}
func Sub_40A770() int {
	return int(C.sub_40A770())
}
func Sub_40A6B0() int {
	return int(C.sub_40A6B0())
}
func Sub_40A6A0(a1 int) {
	C.sub_40A6A0(C.int(a1))
}
func Nox_xxx_netReportAllLatency_4D3050() {
	C.nox_xxx_netReportAllLatency_4D3050()
}
func Sub_4183C0() {
	C.sub_4183C0()
}
func Sub_4D7150() {
	C.sub_4D7150()
}
func Sub_4D71F0() {
	C.sub_4D71F0()
}
func Nox_server_checkWarpGate_4D7600() {
	C.nox_server_checkWarpGate_4D7600()
}
func Sub_4D7A80() {
	C.sub_4D7A80()
}
func Sub_4EC720() {
	C.sub_4EC720()
}
func Sub_50D890() {
	C.sub_50D890()
}
func Sub_4E4170() {
	C.sub_4E4170()
}
func Nox_xxx_voteUptate_506F30() {
	C.nox_xxx_voteUptate_506F30()
}
func Nox_xxx_net_4263C0() {
	C.nox_xxx_net_4263C0()
}
func Sub_4E76C0() {
	C.sub_4E76C0()
}
func Nox_xxx_protectData_56F5C0() {
	C.nox_xxx_protectData_56F5C0()
}
func Sub_57B140() bool {
	return bool(C.sub_57B140())
}
func Sub_57B0A0() {
	C.sub_57B0A0()
}
func Sub_4DF120(a1 unsafe.Pointer) {
	C.sub_4DF120(a1)
}
func Sub_4DF180(a1 unsafe.Pointer) {
	C.sub_4DF180(a1)
}
func Nox_xxx_netUpdateObjectSpecial_527E50(a1 *server.Object, a2 *server.Object) {
	C.nox_xxx_netUpdateObjectSpecial_527E50(asObjectC(a1), asObjectC(a2))
}
func Sub_4D15C0() {
	C.sub_4D15C0()
}
func Sub_4D7B40() {
	C.sub_4D7B40()
}
func Sub_41E4B0(a1 int) {
	C.sub_41E4B0(C.int(a1))
}
func Sub_56F1C0() {
	C.sub_56F1C0()
}
func Sub_4E4ED0() {
	C.sub_4E4ED0()
}
func Nox_motd_4463E0(a1 int) {
	C.nox_motd_4463E0(C.int(a1))
}
func Sub_4259C0() {
	C.sub_4259C0()
}
func Sub_518770() int {
	return int(C.sub_518770())
}
func Nox_xxx_allocItemRespawnArray_4ECA60() int {
	return int(C.nox_xxx_allocItemRespawnArray_4ECA60())
}
func Nox_xxx_registerShopClasses_50E2A0() int {
	return int(C.nox_xxx_registerShopClasses_50E2A0())
}
func Nox_xxx_allocMonsterRelatedArrays_50D780() int {
	return int(C.nox_xxx_allocMonsterRelatedArrays_50D780())
}
func Nox_xxx_allocVoteArray_5066D0() int {
	return int(C.nox_xxx_allocVoteArray_5066D0())
}
func Nox_xxx_monsterList_517520() int {
	return int(C.nox_xxx_monsterList_517520())
}
func Sub_416920() {
	C.sub_416920()
}
func Nox_xxx_allocPendingOwnsArray_516EE0() int {
	return int(C.nox_xxx_allocPendingOwnsArray_516EE0())
}
func Sub_421B10() {
	C.sub_421B10()
}
func Sub_516F10() {
	C.sub_516F10()
}
func Sub_4259F0() {
	C.sub_4259F0()
}
func Sub_4ECA90() {
	C.sub_4ECA90()
}
func Sub_506720() {
	C.sub_506720()
}
func Sub_50D820() {
	C.sub_50D820()
}
func Nox_xxx_deleteShopInventories_50E300() {
	C.nox_xxx_deleteShopInventories_50E300()
}
func Sub_416950() {
	C.sub_416950()
}
func Sub_56F3B0() {
	C.sub_56F3B0()
}
func Sub_51D0E0() {
	C.sub_51D0E0()
}
func Sub_502DF0() {
	C.sub_502DF0()
}
func Sub_481410() {
	C.sub_481410()
}
func Sub_4D0550(a1 string) {
	C.sub_4D0550(internCStr(a1))
}
func Sub_4165F0(a1 int, a2 int) {
	C.sub_4165F0(C.int(a1), C.int(a2))
}
func Sub_4D7430() int {
	return int(C.sub_4D7430())
}
func Sub_4D76F0() int {
	return int(C.sub_4D76F0())
}
func Sub_4D76E0(a1 int) {
	C.sub_4D76E0(C.int(a1))
}
func Nox_xxx_ruleSetNoRespawn_40A5E0(a1 int) {
	C.nox_xxx_ruleSetNoRespawn_40A5E0(C.int(a1))
}
func Sub_455C10() {
	C.sub_455C10()
}
func Sub_456050() {
	C.sub_456050()
}
func Nox_xxx_mapFindCrown_4CFC30() {
	C.nox_xxx_mapFindCrown_4CFC30()
}
func Nox_xxx_mapInfoSetCapflag_417EA0() int {
	return int(C.nox_xxx_mapInfoSetCapflag_417EA0())
}
func Nox_xxx_mapInfoSetFlagball_417F30() int {
	return int(C.nox_xxx_mapInfoSetFlagball_417F30())
}
func Nox_xxx_mapInfoSetKotr_4180D0() int {
	return int(C.nox_xxx_mapInfoSetKotr_4180D0())
}
func Nox_xxx_teamAssignFlags_418640() {
	C.nox_xxx_teamAssignFlags_418640()
}
func Nox_xxx_toggleAllTeamFlags_418690(a1 int) {
	C.nox_xxx_toggleAllTeamFlags_418690(C.int(a1))
}
func Sub_410730() {
	C.sub_410730()
}
func Nox_xxx_j_allocHitArray_511840() {
	C.nox_xxx_j_allocHitArray_511840()
}
func Nox_xxx_decayDestroy_5117B0() {
	C.nox_xxx_decayDestroy_5117B0()
}
func Sub_510E50() {
	C.sub_510E50()
}
func Sub_4D1610() {
	C.sub_4D1610()
}
func Sub_4EC5B0() {
	C.sub_4EC5B0()
}
func Sub_50E360() {
	C.sub_50E360()
}
func Sub_50D7E0() {
	C.sub_50D7E0()
}
func Sub_4E4F80() {
	C.sub_4E4F80()
}
func Sub_4ECFE0() {
	C.sub_4ECFE0()
}
func Sub_511E20() {
	C.sub_511E20()
}
func Sub_516F30() {
	C.sub_516F30()
}
func Nox_xxx_playerBotCreate_4FA700(u *server.Object) {
	C.nox_xxx_playerBotCreate_4FA700(asObjectC(u))
}
func Sub_4F7950(u *server.Object) {
	C.sub_4F7950(asObjectC(u))
}
func Sub_4E5BF0(a1 bool) {
	C.sub_4E5BF0(C.int(bool2int(a1)))
}
func Nox_xxx_Fn_4FCAC0(a1 bool, a2 int) {
	C.nox_xxx_Fn_4FCAC0(C.int(bool2int(a1)), C.int(a2))
}
func Nox_xxx_shopCancelSession_510DC0(a1 *server.TradeSession) {
	C.nox_xxx_shopCancelSession_510DC0(unsafe.Pointer(a1))
}
func Sub_4FCEB0(a1 bool) {
	C.sub_4FCEB0(C.int(bool2int(a1)))
}

var _ = [1]struct{}{}[60-unsafe.Sizeof(server.TileDef{})]

func Get_nox_tile_defs_arr() []server.TileDef {
	ptr := (*server.TileDef)(unsafe.Pointer(&C.nox_tile_defs_arr[0]))
	sz := int(C.nox_tile_def_cnt)
	return unsafe.Slice(ptr, len(C.nox_tile_defs_arr))[:sz]
}
