package legacy

/*
#include "GAME1.h"
#include "GAME1_1.h"
#include "GAME1_2.h"
#include "GAME1_3.h"
#include "GAME2.h"
#include "GAME2_1.h"
#include "GAME3.h"
#include "GAME3_1.h"
#include "GAME3_2.h"
#include "GAME3_3.h"
#include "GAME4.h"
#include "GAME4_1.h"
#include "GAME4_2.h"
#include "GAME4_3.h"
#include "GAME5.h"
#include "GAME5_2.h"
#include "client__system__parsecmd.h"
#include "common__object__modifier.h"
#include "server__system__cscrfunc.h"
#include "server__script__script.h"
#include "client__draw__animdraw.h"
#include "client__draw__canidraw.h"
#include "client__draw__slavedraw.h"
#include "client__draw__staticdraw.h"
#include "client__drawable__update__dball.h"
#include "client__drawable__update__drainup.h"
#include "client__drawable__update__healup.h"
#include "client__drawable__update__manabomb.h"
#include "client__drawable__update__mmislup.h"
#include "client__drawable__update__mtailup.h"
#include "client__drawable__update__sparklup.h"
#include "client__drawable__update__telwake.h"
#include "client__drawable__update__vortexup.h"
extern void* nox_xxx_aClosewoodengat_587000_133480;
extern void* dword_587000_155144;
extern void* dword_587000_127004;
extern void* dword_587000_93164;
extern void* dword_587000_81128;
*/
import "C"
import (
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/common/memmap/nox/blobdata"
)

func InitBlobData() {
	blobdata.Init(&blobdata.Ptrs{
		Ptr_nox_xxx_aClosewoodengat_587000_133480: &C.nox_xxx_aClosewoodengat_587000_133480,
		Ptr_dword_587000_155144:                   &C.dword_587000_155144,
		Ptr_dword_587000_127004:                   &C.dword_587000_127004,
		Ptr_dword_587000_93164:                    &C.dword_587000_93164,
		Ptr_dword_587000_81128:                    &C.dword_587000_81128,

		Ptr_sub_4DFB50:                            unsafe.Pointer(C.sub_4DFB50),
		Ptr_nox_xxx_effectSpeedEngage_4DFC30:      unsafe.Pointer(C.nox_xxx_effectSpeedEngage_4DFC30),
		Ptr_sub_4DFD10:                            unsafe.Pointer(C.sub_4DFD10),
		Ptr_nox_xxx_buff_4DFD80:                   unsafe.Pointer(C.nox_xxx_buff_4DFD80),
		Ptr_nox_xxx_checkPoisonProtectEnch_4DFDE0: unsafe.Pointer(C.nox_xxx_checkPoisonProtectEnch_4DFDE0),
		Ptr_sub_4E0140:                            unsafe.Pointer(C.sub_4E0140),
		Ptr_sub_41C280:                            unsafe.Pointer(C.sub_41C280),
		Ptr_nox_xxx_parseFileInfoData_41C3B0:      unsafe.Pointer(C.nox_xxx_parseFileInfoData_41C3B0),
		Ptr_sub_41C780:                            unsafe.Pointer(C.sub_41C780),
		Ptr_sub_43EC30:                            unsafe.Pointer(C.sub_43EC30),
		Ptr_sub_43ECB0:                            unsafe.Pointer(C.sub_43ECB0),
		Ptr_sub_43ED00:                            unsafe.Pointer(C.sub_43ED00),
		Ptr_sub_43EFD0:                            unsafe.Pointer(C.sub_43EFD0),
		Ptr_sub_43F010:                            unsafe.Pointer(C.sub_43F010),
		Ptr_sub_43F030:                            unsafe.Pointer(C.sub_43F030),
		Ptr_sub_43F050:                            unsafe.Pointer(C.sub_43F050),
		Ptr_sub_43F060:                            unsafe.Pointer(C.sub_43F060),
		Ptr_sub_43F0D0:                            unsafe.Pointer(C.sub_43F0D0),
		Ptr_sub_43E940:                            unsafe.Pointer(C.sub_43E940),
		Ptr_sub_43E9F0:                            unsafe.Pointer(C.sub_43E9F0),
		Ptr_sub_43EA20:                            unsafe.Pointer(C.sub_43EA20),
		Ptr_sub_43EC10:                            unsafe.Pointer(C.sub_43EC10),
		Ptr_sub_4A7A80:                            unsafe.Pointer(C.sub_4A7A80),
		Ptr_sub_4A7AC0:                            unsafe.Pointer(C.sub_4A7AC0),
		Ptr_sub_4A7B00:                            unsafe.Pointer(C.sub_4A7B00),
		Ptr_sub_4A7B40:                            unsafe.Pointer(C.sub_4A7B40),
		Ptr_sub_4A7BA0:                            unsafe.Pointer(C.sub_4A7BA0),
		Ptr_sub_4A7BC0:                            unsafe.Pointer(C.sub_4A7BC0),
		Ptr_sub_4A7C00:                            unsafe.Pointer(C.sub_4A7C00),
		Ptr_sub_4A7C40:                            unsafe.Pointer(C.sub_4A7C40),
		Ptr_sub_4A7C60:                            unsafe.Pointer(C.sub_4A7C60),
		Ptr_sub_4A7CE0:                            unsafe.Pointer(C.sub_4A7CE0),
		Ptr_sub_4A7D00:                            unsafe.Pointer(C.sub_4A7D00),
		Ptr_sub_4A7D50:                            unsafe.Pointer(C.sub_4A7D50),
		Ptr_nox_xxx_updDrawColorlight_4CE390:      unsafe.Pointer(C.nox_xxx_updDrawColorlight_4CE390),
		Ptr_nox_xxx_updDrawCloud_4CE1D0:           unsafe.Pointer(C.nox_xxx_updDrawCloud_4CE1D0),
		Ptr_sub_4CE360:                            unsafe.Pointer(C.sub_4CE360),
		Ptr_nox_xxx_updDrawDBall_4CDF80:           unsafe.Pointer(C.nox_xxx_updDrawDBall_4CDF80),
		Ptr_sub_4CE0A0:                            unsafe.Pointer(C.sub_4CE0A0),
		Ptr_nox_xxx_updDrawDBallCharge_4CE0C0:     unsafe.Pointer(C.nox_xxx_updDrawDBallCharge_4CE0C0),
		Ptr_nox_xxx_updDrawMagic_4CDD80:           unsafe.Pointer(C.nox_xxx_updDrawMagic_4CDD80),
		Ptr_nox_xxx_updDrawSparkleTrail_4CDBF0:    unsafe.Pointer(C.nox_xxx_updDrawSparkleTrail_4CDBF0),
		Ptr_nox_xxx_updDrawMagicMissile_4CD9E0:    unsafe.Pointer(C.nox_xxx_updDrawMagicMissile_4CD9E0),
		Ptr_nox_xxx_updDrawTeleportWake_4CD8D0:    unsafe.Pointer(C.nox_xxx_updDrawTeleportWake_4CD8D0),
		Ptr_sub_4CD690:                            unsafe.Pointer(C.sub_4CD690),
		Ptr_sub_4CD450:                            unsafe.Pointer(C.sub_4CD450),
		Ptr_sub_4CD400:                            unsafe.Pointer(C.sub_4CD400),
		Ptr_sub_4CCE70:                            unsafe.Pointer(C.sub_4CCE70),
		Ptr_sub_4CD090:                            unsafe.Pointer(C.sub_4CD090),
		Ptr_sub_4CD0C0:                            unsafe.Pointer(C.sub_4CD0C0),
		Ptr_sub_4CD0F0:                            unsafe.Pointer(C.sub_4CD0F0),
		Ptr_sub_4CD120:                            unsafe.Pointer(C.sub_4CD120),
		Ptr_nox_xxx_updDrawFist_4CCDB0:            unsafe.Pointer(C.nox_xxx_updDrawFist_4CCDB0),
		Ptr_sub_4CCD00:                            unsafe.Pointer(C.sub_4CCD00),
		Ptr_nox_xxx_updDrawUndeadKiller_4CCCF0:    unsafe.Pointer(C.nox_xxx_updDrawUndeadKiller_4CCCF0),
		Ptr_nox_xxx_updDrawManabombCharge_4CCAC0:  unsafe.Pointer(C.nox_xxx_updDrawManabombCharge_4CCAC0),
		Ptr_nox_xxx_updDrawVortexSource_4CC950:    unsafe.Pointer(C.nox_xxx_updDrawVortexSource_4CC950),
		Ptr_sub_4CA650:                            unsafe.Pointer(C.sub_4CA650),
		Ptr_nox_xxx_updDrawMonsterGen_4BC920:      unsafe.Pointer(C.nox_xxx_updDrawMonsterGen_4BC920),
		Ptr_sub_454120:                            unsafe.Pointer(C.sub_454120),
		Ptr_sub_453750:                            unsafe.Pointer(C.sub_453750),

		Ptr_nox_xxx_strikeOgre_549220:           unsafe.Pointer(C.nox_xxx_strikeOgre_549220),
		Ptr_nox_xxx_strikeScorpion_5495B0:       unsafe.Pointer(C.nox_xxx_strikeScorpion_5495B0),
		Ptr_nox_xxx_strikeVileZombie_549700:     unsafe.Pointer(C.nox_xxx_strikeVileZombie_549700),
		Ptr_nox_xxx_strikeStoneGolem_5497E0:     unsafe.Pointer(C.nox_xxx_strikeStoneGolem_5497E0),
		Ptr_nox_xxx_strikeMechGolem_549960:      unsafe.Pointer(C.nox_xxx_strikeMechGolem_549960),
		Ptr_nox_xxx_strikeWasp_549980:           unsafe.Pointer(C.nox_xxx_strikeWasp_549980),
		Ptr_nox_xxx_strikeSpider_549BC0:         unsafe.Pointer(C.nox_xxx_strikeSpider_549BC0),
		Ptr_nox_xxx_strikeSpittingSpider_549CA0: unsafe.Pointer(C.nox_xxx_strikeSpittingSpider_549CA0),
		Ptr_nox_xxx_strikeGhost_549A60:          unsafe.Pointer(C.nox_xxx_strikeGhost_549A60),
		Ptr_nox_xxx_strikeBomber_549BB0:         unsafe.Pointer(C.nox_xxx_strikeBomber_549BB0),
		Ptr_nox_xxx_strikeMonsterDefault_549380: unsafe.Pointer(C.nox_xxx_strikeMonsterDefault_549380),
		Ptr_sub_549D80:                          unsafe.Pointer(C.sub_549D80),
		Ptr_sub_549E00:                          unsafe.Pointer(C.sub_549E00),
		Ptr_sub_549E70:                          unsafe.Pointer(C.sub_549E70),
		Ptr_sub_549E90:                          unsafe.Pointer(C.sub_549E90),
		Ptr_sub_549FA0:                          unsafe.Pointer(C.sub_549FA0),
		Ptr_nox_bomberDead_54A150:               unsafe.Pointer(C.nox_bomberDead_54A150),
		Ptr_sub_54A250:                          unsafe.Pointer(C.sub_54A250),
		Ptr_sub_54A310:                          unsafe.Pointer(C.sub_54A310),
		Ptr_sub_54A750:                          unsafe.Pointer(C.sub_54A750),
		Ptr_nox_xxx_monsterDeadTroll_54A270:     unsafe.Pointer(C.nox_xxx_monsterDeadTroll_54A270),
		Ptr_sub_54A890:                          unsafe.Pointer(C.sub_54A890),
		Ptr_sub_54A900:                          unsafe.Pointer(C.sub_54A900),
		Ptr_sub_54A7D0:                          unsafe.Pointer(C.sub_54A7D0),
		Ptr_sub_54A850:                          unsafe.Pointer(C.sub_54A850),
		Ptr_sub_54A950:                          unsafe.Pointer(C.sub_54A950),
	})
}
