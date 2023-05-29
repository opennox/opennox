package legacy

import (
	"unsafe"

	"github.com/noxworld-dev/opennox-lib/spell"
	"github.com/noxworld-dev/opennox-lib/things"
	"github.com/noxworld-dev/opennox-lib/types"

	"github.com/noxworld-dev/opennox/v1/server"
)

var (
	GetPhonemeTree                    func() unsafe.Pointer
	Nox_xxx_spellAwardAll1_4EFD80     func(p *server.Player)
	Nox_xxx_spellAwardAll2_4EFC80     func(p *server.Player)
	Nox_xxx_spellAwardAll3_4EFE10     func(p *server.Player)
	Nox_xxx_spellGetAud44_424800      func(ind, a2 int) int
	Nox_xxx_spellTitle_424930         func(ind int) (string, bool)
	Nox_xxx_spellDescription_424A30   func(ind int) (string, bool)
	Nox_xxx_spellByTitle_424960       func(ctitle string) int
	Nox_xxx_spellManaCost_4249A0      func(ind, a2 int) int
	Nox_xxx_spellPhonemes_424A20      func(ind, ind2 int) spell.Phoneme
	Nox_xxx_spellIcon_424A90          func(ind int) unsafe.Pointer
	Nox_xxx_spellIconHighlight_424AB0 func(ind int) unsafe.Pointer
	Nox_xxx_spellFirstValid_424AD0    func() int
	Nox_xxx_spellNextValid_424AF0     func(ind int) int
	Nox_xxx_spellIsValid_424B50       func(ind int) bool
	Nox_xxx_spellIsEnabled_424B70     func(ind int) bool
	Nox_xxx_spellEnable_424B90        func(ind int) bool
	Nox_xxx_spellDisable_424BB0       func(ind int) bool
	Nox_xxx_spellCanUseInTrap_424BF0  func(ind int) bool
	Nox_xxx_spellPrice_424C40         func(ind int) int
	Nox_xxx_spellEnableAll_424BD0     func()
	Nox_xxx_castSpellByUser_4FDD20    func(a1 int, a2 *server.Object, a3 unsafe.Pointer) int
)

// nox_xxx_spellGetDefArrayPtr_424820
func nox_xxx_spellGetDefArrayPtr_424820() unsafe.Pointer {
	return unsafe.Pointer(GetPhonemeTree())
}

// nox_xxx_getEnchantSpell_424920
func nox_xxx_getEnchantSpell_424920(enc int) int {
	return int(server.EnchantID(enc).Spell())
}

// nox_xxx_getEnchantName_4248F0
func nox_xxx_getEnchantName_4248F0(enc int) *char {
	return internCStr(server.EnchantID(enc).String())
}

// nox_xxx_enchantByName_424880
func nox_xxx_enchantByName_424880(cname *char) int {
	id, ok := server.ParseEnchant(GoString(cname))
	if !ok {
		return -1
	}
	return int(id)
}

// nox_xxx_spellNameByN_424870
func nox_xxx_spellNameByN_424870(ind int) *char {
	s := spell.ID(ind).String()
	if s == "" {
		return nil
	}
	return internCStr(s)
}

// nox_xxx_spellNameToN_4243F0
func nox_xxx_spellNameToN_4243F0(cid *char) int {
	id := GoString(cid)
	ind := spell.ParseID(id)
	if ind <= 0 {
		return 0
	}
	return int(ind)
}

// nox_xxx_spellAwardAll1_4EFD80
func nox_xxx_spellAwardAll1_4EFD80(p *nox_playerInfo) { Nox_xxx_spellAwardAll1_4EFD80(asPlayerS(p)) }

// nox_xxx_spellAwardAll2_4EFC80
func nox_xxx_spellAwardAll2_4EFC80(p *nox_playerInfo) { Nox_xxx_spellAwardAll2_4EFC80(asPlayerS(p)) }

// nox_xxx_spellAwardAll3_4EFE10
func nox_xxx_spellAwardAll3_4EFE10(p *nox_playerInfo) { Nox_xxx_spellAwardAll3_4EFE10(asPlayerS(p)) }

// nox_xxx_spellFlySearchTarget_540610
func nox_xxx_spellFlySearchTarget_540610(cpos *float2, msl *nox_object_t, sflags int, dist float32, a5 int, self *nox_object_t) *nox_object_t {
	var pos *types.Pointf
	if cpos != nil {
		pos = &types.Pointf{X: float32(cpos.field_0), Y: float32(cpos.field_4)}
	}
	return asObjectC(GetServer().Nox_xxx_spellFlySearchTarget(pos, ToObjS(msl), things.SpellFlags(sflags), float32(dist), int(a5), asObjectS(self)))
}

// nox_xxx_spellGetAud44_424800
func nox_xxx_spellGetAud44_424800(ind, a2 int) int { return Nox_xxx_spellGetAud44_424800(ind, a2) }

// nox_xxx_spellTitle_424930
func nox_xxx_spellTitle_424930(ind int) *wchar2_t {
	s, ok := Nox_xxx_spellTitle_424930(ind)
	if !ok {
		return nil
	}
	return internWStr(s)
}

// nox_xxx_spellDescription_424A30
func nox_xxx_spellDescription_424A30(ind int) *wchar2_t {
	s, ok := Nox_xxx_spellDescription_424A30(ind)
	if !ok {
		return nil
	}
	return internWStr(s)
}

// nox_xxx_spellByTitle_424960
func nox_xxx_spellByTitle_424960(ctitle *wchar2_t) int {
	return Nox_xxx_spellByTitle_424960(GoWString(ctitle))
}

// nox_xxx_spellManaCost_4249A0
func nox_xxx_spellManaCost_4249A0(ind, a2 int) int { return Nox_xxx_spellManaCost_4249A0(ind, a2) }

// nox_xxx_spellPhonemes_424A20
func nox_xxx_spellPhonemes_424A20(ind, ind2 int) char {
	return char(Nox_xxx_spellPhonemes_424A20(ind, ind2))
}

// nox_xxx_spellHasFlags_424A50
func nox_xxx_spellHasFlags_424A50(ind, flags int) bool {
	return bool(GetServer().SpellHasFlags(spell.ID(ind), things.SpellFlags(flags)))
}

// nox_xxx_spellFlags_424A70
func nox_xxx_spellFlags_424A70(ind int) uint { return uint(GetServer().SpellFlags(spell.ID(ind))) }

// nox_xxx_spellIcon_424A90
func nox_xxx_spellIcon_424A90(ind int) unsafe.Pointer { return Nox_xxx_spellIcon_424A90(ind) }

// nox_xxx_spellIconHighlight_424AB0
func nox_xxx_spellIconHighlight_424AB0(ind int) unsafe.Pointer {
	return Nox_xxx_spellIconHighlight_424AB0(ind)
}

// nox_xxx_spellFirstValid_424AD0
func nox_xxx_spellFirstValid_424AD0() int { return Nox_xxx_spellFirstValid_424AD0() }

// nox_xxx_spellNextValid_424AF0
func nox_xxx_spellNextValid_424AF0(ind int) int { return Nox_xxx_spellNextValid_424AF0(ind) }

// nox_xxx_spellIsValid_424B50
func nox_xxx_spellIsValid_424B50(ind int) bool { return bool(Nox_xxx_spellIsValid_424B50(ind)) }

// nox_xxx_spellIsEnabled_424B70
func nox_xxx_spellIsEnabled_424B70(ind int) bool { return bool(Nox_xxx_spellIsEnabled_424B70(ind)) }

// nox_xxx_spellEnable_424B90
func nox_xxx_spellEnable_424B90(ind int) bool { return bool(Nox_xxx_spellEnable_424B90(ind)) }

// nox_xxx_spellDisable_424BB0
func nox_xxx_spellDisable_424BB0(ind int) bool { return bool(Nox_xxx_spellDisable_424BB0(ind)) }

// nox_xxx_spellCanUseInTrap_424BF0
func nox_xxx_spellCanUseInTrap_424BF0(ind int) bool {
	return bool(Nox_xxx_spellCanUseInTrap_424BF0(ind))
}

// nox_xxx_spellPrice_424C40
func nox_xxx_spellPrice_424C40(ind int) int { return Nox_xxx_spellPrice_424C40(ind) }

// nox_xxx_spellEnableAll_424BD0
func nox_xxx_spellEnableAll_424BD0() { Nox_xxx_spellEnableAll_424BD0() }

// nox_xxx_spellAccept_4FD400
func nox_xxx_spellAccept_4FD400(ispellID int, a2, a3p, a4p *nox_object_t, a5p unsafe.Pointer, lvli int) int {
	if GetServer().Nox_xxx_spellAccept4FD400(spell.ID(ispellID), asObjectS(a2), asObjectS(a3p), asObjectS(a4p), (*server.SpellAcceptArg)(a5p), lvli) {
		return 1
	}
	return 0
}

// nox_xxx_castSpellByUser_4FDD20
func nox_xxx_castSpellByUser_4FDD20(a1 int, a2 *nox_object_t, a3 unsafe.Pointer) int {
	return Nox_xxx_castSpellByUser_4FDD20(a1, asObjectS(a2), a3)
}
func Nox_xxx_spellCastByBook_4FCB80() {
	nox_xxx_spellCastByBook_4FCB80()
}
func Nox_xxx_playerResetProtectionCRC_56F7D0(a1 uintptr, a2 int) {
	nox_xxx_playerResetProtectionCRC_56F7D0(int(a1), int(a2))
}
func Nox_xxx_playerAwardSpellProtectionCRC_56FCE0(a1 uintptr, a2 int, a3 int) {
	nox_xxx_playerAwardSpellProtectionCRC_56FCE0(int(a1), int(a2), int(a3))
}
func Nox_xxx_playerApplyProtectionCRC_56FD50(a1 uintptr, a2 unsafe.Pointer, a3 int) {
	nox_xxx_playerApplyProtectionCRC_56FD50(int(a1), a2, int(a3))
}
func Nox_xxx_spellGrantToPlayer_4FB550(a1 *server.Object, a2 spell.ID, a3 int, a4 int, a5 int) int {
	return int(nox_xxx_spellGrantToPlayer_4FB550(asObjectC(a1), int(a2), int(a3), int(a4), int(a5)))
}
func Nox_xxx_gameCaptureMagic_4FDC10(a1 spell.ID, a2 *server.Object) int {
	return int(nox_xxx_gameCaptureMagic_4FDC10(int(a1), asObjectC(a2)))
}
func Nox_spells_call_intint6_go(a1 unsafe.Pointer, a2 spell.ID, a3 *server.Object, a4 *server.Object, a5 *server.Object, a6 *server.SpellAcceptArg, a7 int) int {
	return int(nox_spells_call_intint6_go((*[0]byte)(a1), int(a2), asObjectC(a3), asObjectC(a4), asObjectC(a5), unsafe.Pointer(a6), int(a7)))
}
func Nox_xxx_createSpellFly_4FDDA0(a1 *server.Object, a2 *server.Object, a3 spell.ID) {
	nox_xxx_createSpellFly_4FDDA0(asObjectC(a1), asObjectC(a2), int(a3))
}
func Nox_xxx_spellGetPower_4FE7B0(a1 spell.ID, a2 *server.Object) int {
	return int(nox_xxx_spellGetPower_4FE7B0(int(a1), asObjectC(a2)))
}

func Nox_xxx_spellArachna_52DC80(spellID spell.ID, a2, a3, a4 *server.Object, sa *server.SpellAcceptArg, lvl int) int {
	return Nox_spells_call_intint6_go(nox_xxx_spellArachna_52DC80, spellID, a2.SObj(), a3.SObj(), a4.SObj(), sa, lvl)
}

func Nox_xxx_castBurn_52C3E0(spellID spell.ID, a2, a3, a4 *server.Object, sa *server.SpellAcceptArg, lvl int) int {
	return Nox_spells_call_intint6_go(nox_xxx_castBurn_52C3E0, spellID, a2.SObj(), a3.SObj(), a4.SObj(), sa, lvl)
}

func Nox_xxx_spellCastCleansingFlame_52D5C0(spellID spell.ID, a2, a3, a4 *server.Object, sa *server.SpellAcceptArg, lvl int) int {
	return Nox_spells_call_intint6_go(nox_xxx_spellCastCleansingFlame_52D5C0, spellID, a2.SObj(), a3.SObj(), a4.SObj(), sa, lvl)
}

func Nox_xxx_castConfuse_52C1E0(spellID spell.ID, a2, a3, a4 *server.Object, sa *server.SpellAcceptArg, lvl int) int {
	return Nox_spells_call_intint6_go(nox_xxx_castConfuse_52C1E0, spellID, a2.SObj(), a3.SObj(), a4.SObj(), sa, lvl)
}

func Nox_xxx_castCounterSpell_52BBB0(spellID spell.ID, a2, a3, a4 *server.Object, sa *server.SpellAcceptArg, lvl int) int {
	return Nox_spells_call_intint6_go(nox_xxx_castCounterSpell_52BBB0, spellID, a2.SObj(), a3.SObj(), a4.SObj(), sa, lvl)
}

func Nox_xxx_castCurePoison_52CDB0(spellID spell.ID, a2, a3, a4 *server.Object, sa *server.SpellAcceptArg, lvl int) int {
	return Nox_spells_call_intint6_go(nox_xxx_castCurePoison_52CDB0, spellID, a2.SObj(), a3.SObj(), a4.SObj(), sa, lvl)
}

func Nox_xxx_castEquake_52DE40(spellID spell.ID, a2, a3, a4 *server.Object, sa *server.SpellAcceptArg, lvl int) int {
	return Nox_spells_call_intint6_go(nox_xxx_castEquake_52DE40, spellID, a2.SObj(), a3.SObj(), a4.SObj(), sa, lvl)
}

func Nox_xxx_castFireball_52C790(spellID spell.ID, a2, a3, a4 *server.Object, sa *server.SpellAcceptArg, lvl int) int {
	return Nox_spells_call_intint6_go(nox_xxx_castFireball_52C790, spellID, a2.SObj(), a3.SObj(), a4.SObj(), sa, lvl)
}

func Nox_xxx_castFist_52D3C0(spellID spell.ID, a2, a3, a4 *server.Object, sa *server.SpellAcceptArg, lvl int) int {
	return Nox_spells_call_intint6_go(nox_xxx_castFist_52D3C0, spellID, a2.SObj(), a3.SObj(), a4.SObj(), sa, lvl)
}

func Nox_xxx_castFumble_52C060(spellID spell.ID, a2, a3, a4 *server.Object, sa *server.SpellAcceptArg, lvl int) int {
	return Nox_spells_call_intint6_go(nox_xxx_castFumble_52C060, spellID, a2.SObj(), a3.SObj(), a4.SObj(), sa, lvl)
}

func Sub_52BEB0(spellID spell.ID, a2, a3, a4 *server.Object, sa *server.SpellAcceptArg, lvl int) int {
	return Nox_spells_call_intint6_go(sub_52BEB0, spellID, a2.SObj(), a3.SObj(), a4.SObj(), sa, lvl)
}

func Sub_52DD50(spellID spell.ID, a2, a3, a4 *server.Object, sa *server.SpellAcceptArg, lvl int) int {
	return Nox_spells_call_intint6_go(sub_52DD50, spellID, a2.SObj(), a3.SObj(), a4.SObj(), sa, lvl)
}

func Nox_xxx_castLock_52CE90(spellID spell.ID, a2, a3, a4 *server.Object, sa *server.SpellAcceptArg, lvl int) int {
	return Nox_spells_call_intint6_go(nox_xxx_castLock_52CE90, spellID, a2.SObj(), a3.SObj(), a4.SObj(), sa, lvl)
}

func Sub_52CA80(spellID spell.ID, a2, a3, a4 *server.Object, sa *server.SpellAcceptArg, lvl int) int {
	return Nox_spells_call_intint6_go(sub_52CA80, spellID, a2.SObj(), a3.SObj(), a4.SObj(), sa, lvl)
}

func Sub_52CBD0(spellID spell.ID, a2, a3, a4 *server.Object, sa *server.SpellAcceptArg, lvl int) int {
	return Nox_spells_call_intint6_go(sub_52CBD0, spellID, a2.SObj(), a3.SObj(), a4.SObj(), sa, lvl)
}

func Nox_xxx_castMeteor_52D9D0(spellID spell.ID, a2, a3, a4 *server.Object, sa *server.SpellAcceptArg, lvl int) int {
	return Nox_spells_call_intint6_go(nox_xxx_castMeteor_52D9D0, spellID, a2.SObj(), a3.SObj(), a4.SObj(), sa, lvl)
}

func Nox_xxx_castMeteorShower_52D8A0(spellID spell.ID, a2, a3, a4 *server.Object, sa *server.SpellAcceptArg, lvl int) int {
	return Nox_spells_call_intint6_go(nox_xxx_castMeteorShower_52D8A0, spellID, a2.SObj(), a3.SObj(), a4.SObj(), sa, lvl)
}

func Nox_xxx_castPixies_540440(spellID spell.ID, a2, a3, a4 *server.Object, sa *server.SpellAcceptArg, lvl int) int {
	return Nox_spells_call_intint6_go(nox_xxx_castPixies_540440, spellID, a2.SObj(), a3.SObj(), a4.SObj(), sa, lvl)
}

func Nox_xxx_castPoison_52C720(spellID spell.ID, a2, a3, a4 *server.Object, sa *server.SpellAcceptArg, lvl int) int {
	return Nox_spells_call_intint6_go(nox_xxx_castPoison_52C720, spellID, a2.SObj(), a3.SObj(), a4.SObj(), sa, lvl)
}

func Nox_xxx_castPull_52BFA0(spellID spell.ID, a2, a3, a4 *server.Object, sa *server.SpellAcceptArg, lvl int) int {
	return Nox_spells_call_intint6_go(nox_xxx_castPull_52BFA0, spellID, a2.SObj(), a3.SObj(), a4.SObj(), sa, lvl)
}

func Nox_xxx_castPush_52C000(spellID spell.ID, a2, a3, a4 *server.Object, sa *server.SpellAcceptArg, lvl int) int {
	return Nox_spells_call_intint6_go(nox_xxx_castPush_52C000, spellID, a2.SObj(), a3.SObj(), a4.SObj(), sa, lvl)
}

func Nox_xxx_castSpellWinkORrestoreHealth_52BF20(spellID spell.ID, a2, a3, a4 *server.Object, sa *server.SpellAcceptArg, lvl int) int {
	return Nox_spells_call_intint6_go(nox_xxx_castSpellWinkORrestoreHealth_52BF20, spellID, a2.SObj(), a3.SObj(), a4.SObj(), sa, lvl)
}

func Sub_52BF50(spellID spell.ID, a2, a3, a4 *server.Object, sa *server.SpellAcceptArg, lvl int) int {
	return Nox_spells_call_intint6_go(sub_52BF50, spellID, a2.SObj(), a3.SObj(), a4.SObj(), sa, lvl)
}

func Nox_xxx_useShock_52C5A0(spellID spell.ID, a2, a3, a4 *server.Object, sa *server.SpellAcceptArg, lvl int) int {
	return Nox_spells_call_intint6_go(nox_xxx_useShock_52C5A0, spellID, a2.SObj(), a3.SObj(), a4.SObj(), sa, lvl)
}

func Nox_xxx_castStun_52C2C0(spellID spell.ID, a2, a3, a4 *server.Object, sa *server.SpellAcceptArg, lvl int) int {
	return Nox_spells_call_intint6_go(nox_xxx_castStun_52C2C0, spellID, a2.SObj(), a3.SObj(), a4.SObj(), sa, lvl)
}

func Nox_xxx_castTelekinesis_52D330(spellID spell.ID, a2, a3, a4 *server.Object, sa *server.SpellAcceptArg, lvl int) int {
	return Nox_spells_call_intint6_go(nox_xxx_castTelekinesis_52D330, spellID, a2.SObj(), a3.SObj(), a4.SObj(), sa, lvl)
}

func Nox_xxx_castToxicCloud_52DB60(spellID spell.ID, a2, a3, a4 *server.Object, sa *server.SpellAcceptArg, lvl int) int {
	return Nox_spells_call_intint6_go(nox_xxx_castToxicCloud_52DB60, spellID, a2.SObj(), a3.SObj(), a4.SObj(), sa, lvl)
}

func Sub_52CCD0(spellID spell.ID, a2, a3, a4 *server.Object, sa *server.SpellAcceptArg, lvl int) int {
	return Nox_spells_call_intint6_go(sub_52CCD0, spellID, a2.SObj(), a3.SObj(), a4.SObj(), sa, lvl)
}

func Get_nox_xxx_spellBlink2_530310() unsafe.Pointer {
	return nox_xxx_spellBlink2_530310
}

func Get_nox_xxx_spellBlink1_530380() unsafe.Pointer {
	return nox_xxx_spellBlink1_530380
}

func Get_sub_52F460() unsafe.Pointer {
	return sub_52F460
}

func Get_nox_xxx_charmCreature1_5011F0() unsafe.Pointer {
	return nox_xxx_charmCreature1_5011F0
}

func Get_nox_xxx_charmCreatureFinish_5013E0() unsafe.Pointer {
	return nox_xxx_charmCreatureFinish_5013E0
}

func Get_nox_xxx_charmCreature2_501690() unsafe.Pointer {
	return nox_xxx_charmCreature2_501690
}

func Get_nox_xxx_spellTurnUndeadCreate_531310() unsafe.Pointer {
	return nox_xxx_spellTurnUndeadCreate_531310
}

func Get_nox_xxx_spellTurnUndeadUpdate_531410() unsafe.Pointer {
	return nox_xxx_spellTurnUndeadUpdate_531410
}

func Get_nox_xxx_spellTurnUndeadDelete_531420() unsafe.Pointer {
	return nox_xxx_spellTurnUndeadDelete_531420
}

func Get_nox_xxx_spellDrainMana_52E210() unsafe.Pointer {
	return nox_xxx_spellDrainMana_52E210
}

func Get_nox_xxx_spellEnergyBoltStop_52E820() unsafe.Pointer {
	return nox_xxx_spellEnergyBoltStop_52E820
}

func Get_nox_xxx_spellEnergyBoltTick_52E850() unsafe.Pointer {
	return nox_xxx_spellEnergyBoltTick_52E850
}

func Get_nullsub_29() unsafe.Pointer {
	return nullsub_29
}

func Get_nox_xxx_firewalkTick_52ED40() unsafe.Pointer {
	return nox_xxx_firewalkTick_52ED40
}

func Get_sub_52EF30() unsafe.Pointer {
	return sub_52EF30
}

func Get_sub_52EFD0() unsafe.Pointer {
	return sub_52EFD0
}

func Get_sub_52F1D0() unsafe.Pointer {
	return sub_52F1D0
}

func Get_sub_52F220() unsafe.Pointer {
	return sub_52F220
}

func Get_sub_52F2E0() unsafe.Pointer {
	return sub_52F2E0
}

func Get_nox_xxx_onStartLightning_52F820() unsafe.Pointer {
	return nox_xxx_onStartLightning_52F820
}

func Get_nox_xxx_onFrameLightning_52F8A0() unsafe.Pointer {
	return nox_xxx_onFrameLightning_52F8A0
}

func Get_sub_530100() unsafe.Pointer {
	return sub_530100
}

func Get_nox_xxx_castShield1_52F5A0() unsafe.Pointer {
	return nox_xxx_castShield1_52F5A0
}

func Get_sub_52F650() unsafe.Pointer {
	return sub_52F650
}

func Get_sub_52F670() unsafe.Pointer {
	return sub_52F670
}

func Get_nox_xxx_spellCreateMoonglow_531A00() unsafe.Pointer {
	return nox_xxx_spellCreateMoonglow_531A00
}

func Get_sub_531AF0() unsafe.Pointer {
	return sub_531AF0
}

func Get_nox_xxx_manaBomb_530F90() unsafe.Pointer {
	return nox_xxx_manaBomb_530F90
}

func Get_nox_xxx_manaBombBoom_5310C0() unsafe.Pointer {
	return nox_xxx_manaBombBoom_5310C0
}

func Get_sub_531290() unsafe.Pointer {
	return sub_531290
}

func Get_nox_xxx_plasmaSmth_531580() unsafe.Pointer {
	return nox_xxx_plasmaSmth_531580
}

func Get_nox_xxx_plasmaShot_531600() unsafe.Pointer {
	return nox_xxx_plasmaShot_531600
}

func Get_sub_5319E0() unsafe.Pointer {
	return sub_5319E0
}

func Get_sub_531490() unsafe.Pointer {
	return sub_531490
}

func Get_sub_5314F0() unsafe.Pointer {
	return sub_5314F0
}

func Get_sub_531560() unsafe.Pointer {
	return sub_531560
}

func Get_nox_xxx_summonStart_500DA0() unsafe.Pointer {
	return nox_xxx_summonStart_500DA0
}

func Get_nox_xxx_summonFinish_5010D0() unsafe.Pointer {
	return nox_xxx_summonFinish_5010D0
}

func Get_nox_xxx_summonCancel_5011C0() unsafe.Pointer {
	return nox_xxx_summonCancel_5011C0
}

func Get_sub_530CA0() unsafe.Pointer {
	return sub_530CA0
}

func Get_sub_530D30() unsafe.Pointer {
	return sub_530D30
}

func Get_nox_xxx_spellTagCreature_530160() unsafe.Pointer {
	return nox_xxx_spellTagCreature_530160
}

func Get_sub_530250() unsafe.Pointer {
	return sub_530250
}

func Get_sub_530270() unsafe.Pointer {
	return sub_530270
}

func Get_sub_5305D0() unsafe.Pointer {
	return sub_5305D0
}

func Get_sub_530650() unsafe.Pointer {
	return sub_530650
}

func Get_nox_xxx_castTele_530820() unsafe.Pointer {
	return nox_xxx_castTele_530820
}

func Get_sub_530880() unsafe.Pointer {
	return sub_530880
}

func Get_sub_530A30_spell_execdur() unsafe.Pointer {
	return sub_530A30_spell_execdur
}

func Get_nox_xxx_castTTT_530B70() unsafe.Pointer {
	return nox_xxx_castTTT_530B70
}

func Get_nox_xxx_spellWallCreate_4FFA90() unsafe.Pointer {
	return nox_xxx_spellWallCreate_4FFA90
}

func Get_nox_xxx_spellWallUpdate_500070() unsafe.Pointer {
	return nox_xxx_spellWallUpdate_500070
}

func Get_nox_xxx_spellWallDestroy_500080() unsafe.Pointer {
	return nox_xxx_spellWallDestroy_500080
}