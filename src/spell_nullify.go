package opennox

import (
	"github.com/opennox/libs/spell"

	"github.com/opennox/opennox/v1/server"
)

func castNullify(spellID spell.ID, _, a3, _ *server.Object, args *server.SpellAcceptArg, lvl int) int {
	return castBuffSpell(spellID, server.ENCHANT_ANTI_MAGIC, lvl, args.Obj, spellBuffConf{
		Dur: noxServer.abilities.defs[server.AbilityWarcry].duration, Orig: a3, Offensive: true,
	})
}
