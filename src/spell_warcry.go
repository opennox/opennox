package opennox

import (
	"github.com/opennox/libs/object"
	"github.com/opennox/libs/player"
	"github.com/opennox/libs/spell"

	"github.com/opennox/opennox/v1/legacy"
	"github.com/opennox/opennox/v1/legacy/common/alloc"
	"github.com/opennox/opennox/v1/server"
)

func nox_xxx_warriorWarcry_53FF40(u *server.Object) {
	if u == nil {
		return
	}
	const dist = 300.0
	s := noxServer
	if u.Class().Has(object.ClassPlayer) {
		if ud := u.UpdateDataPlayer(); ud != nil {
			ud.Field59_0 = 0
		}
	}
	nox_xxx_playerSetState_4FA020(u, server.PlayerState1)
	legacy.Nox_xxx_spell_4FE680(u, dist)
	s.Map.EachObjInCircle(u.Pos(), dist, func(it *server.Object) bool {
		if !it.Class().HasAny(object.MaskUnits) {
			return true
		}
		if it.Flags().HasAny(object.FlagDestroyed | object.FlagDead) {
			return true
		}
		u2 := asObjectS(it)
		if u2.Class().Has(object.ClassPlayer) {
			pl := u2.ControllingPlayer()
			if pl == nil {
				return true
			}
			if pl.PlayerClass() == player.Warrior {
				return true
			}
			if pl.Field3680&0x1 != 0 {
				return true
			}
		} else if u2.Class().Has(object.ClassMonster) {
			ud := u2.UpdateDataMonster()
			if !ud.StatusFlags.Has(object.MonStatusCanCastSpells) {
				return true
			}
		}
		if !u.TeamPtr().SameAs(u2.TeamPtr()) {
			if s.MapTraceVision(u2.SObj(), u) {
				pos := u2.Pos()
				sa, free := alloc.New(server.SpellAcceptArg{})
				defer free()
				*sa = server.SpellAcceptArg{
					Obj: u2.SObj(),
					Pos: pos,
				}
				s.Nox_xxx_spellAccept4FD400(spell.SPELL_NULLIFY, u, u, u, sa, 5)
			}
		}
		return true
	})
	asObjectS(u).DisableEnchant(server.ENCHANT_INVISIBLE)
	asObjectS(u).DisableEnchant(server.ENCHANT_INVULNERABLE)
	s.Spells.Dur.CancelFor(spell.SPELL_OVAL_SHIELD, u)
	s.abilities.netAbilReportActive(u, server.AbilityWarcry, true)
}
