package opennox

import (
	"github.com/opennox/libs/object"
	"github.com/opennox/libs/spell"

	"github.com/opennox/opennox/v1/server"
)

func castInfravision(spellID spell.ID, _, _, _ *server.Object, args *server.SpellAcceptArg, lvl int) int {
	return castBuffSpell(spellID, server.ENCHANT_INFRAVISION, lvl, args.Obj, spellBuffConf{
		DurOpt: "InfravisionEnchantDuration",
	})
}

func nox_xxx_warriorInfravis_540110(u *server.Object, dur int) {
	if u == nil {
		return
	}
	if !u.Class().Has(object.ClassPlayer) {
		return
	}
	s := noxServer
	if ud := u.UpdateDataPlayer(); ud != nil {
		asObjectS(u).ApplyEnchant(server.ENCHANT_INFRAVISION, dur, int(ud.Player.SpellLvl[server.AbilityInfravis]))
		s.abilities.netAbilReportActive(u, server.AbilityInfravis, true)
	}
}
