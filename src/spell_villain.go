package opennox

import (
	"github.com/opennox/libs/spell"

	"github.com/opennox/opennox/v1/server"
)

func castVillain(spellID spell.ID, _, a3, _ *server.Object, args *server.SpellAcceptArg, lvl int) int {
	return castBuffSpell(spellID, server.ENCHANT_VILLAIN, lvl, args.Obj, spellBuffConf{
		Dur: 12, DurInSec: true, DurLevelMul: true, Orig: a3, Offensive: true,
	})
}
