package opennox

import (
	"github.com/opennox/libs/spell"

	"github.com/opennox/opennox/v1/server"
)

func castRun(spellID spell.ID, _, a3, _ *server.Object, args *server.SpellAcceptArg, lvl int) int {
	return castBuffSpell(spellID, server.ENCHANT_RUN, lvl, args.Obj, spellBuffConf{
		DurOpt: "RunEnchantDuration", Orig: a3, Offensive: true,
	})
}
