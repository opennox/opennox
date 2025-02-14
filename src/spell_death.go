package opennox

import (
	"github.com/opennox/libs/spell"

	"github.com/opennox/opennox/v1/server"
)

func castDeath(spellID spell.ID, _, a3, _ *server.Object, args *server.SpellAcceptArg, lvl int) int {
	return castBuffSpell(spellID, server.ENCHANT_DEATH, lvl, args.Obj, spellBuffConf{
		Dur: 450, Once: true, Orig: a3, Offensive: true,
	})
}
