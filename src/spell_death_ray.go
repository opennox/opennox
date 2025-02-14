package opennox

import (
	"image"
	"math"

	"github.com/opennox/libs/noxnet/netmsg"
	"github.com/opennox/libs/object"
	"github.com/opennox/libs/spell"

	"github.com/opennox/opennox/v1/legacy"
	"github.com/opennox/opennox/v1/server"
)

var (
	spellDRayVioletSpark int
)

func castDeathRay(spellID spell.ID, a2, a3, a4 *server.Object, a5 *server.SpellAcceptArg, lvl int) int {
	s := noxServer
	if a5 == nil || a4 == nil {
		return 0
	}
	pos4 := a4.Pos()
	pos16 := a5.Pos
	if !s.MapTraceRay9(pos4, pos16) {
		if a4.Class().Has(object.ClassPlayer) && a3.Class().Has(object.ClassPlayer) {
			s.NetInformTextMsg(asObjectS(a3).ControllingPlayer().PlayerIndex(), 0, 2)
		}
		return 0
	}
	dmg := int(s.Balance.Float("DeathRayDamage"))
	rin := float32(s.Balance.Float("DeathRayInRadius"))
	rout := float32(s.Balance.Float("DeathRayOutRadius"))
	s.Nox_xxx_mapDamageUnitsAround(pos16, rout, rin, dmg, object.DamageZapRay, a3, nil, false)
	s.Nox_xxx_netSendRayFx_5232F0(netmsg.MSG_FX_DEATH_RAY, pos4.Point(), pos16.Point())
	snd := s.Spells.DefByInd(spellID).GetCastSound()
	s.Audio.EventObj(snd, a4, 0, 0)
	legacy.Nox_xxx_sMakeScorch_537AF0(pos16, 1)
	return 1
}

func (c *Client) clientFXDeathRay(p1, p2 image.Point) {
	if !nox_client_isConnected() {
		return
	}
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y
	dist := int(math.Sqrt(float64(dx*dx+dy*dy))) + 1
	if dist < 0 {
		return
	}
	if spellDRayVioletSpark == 0 {
		spellDRayVioletSpark = c.Things.IndByID("VioletSpark")
	}
	for i := 0; i < dist/2+1; i++ {
		pos := p1.Add(image.Point{
			X: (2 * dx * i) / dist,
			Y: (2 * dy * i) / dist,
		})
		dr := c.Nox_xxx_spriteLoadAdd_45A360_drawable(spellDRayVioletSpark, pos)
		if dr == nil {
			continue
		}
		r1 := c.srv.Rand.Other.Int(0, 255)
		r2 := c.srv.Rand.Other.Int(1, 200)
		expire := c.srv.Rand.Other.Int(20, 40)
		z := c.srv.Rand.Other.Int(15, 30)
		vz := c.srv.Rand.Other.Int(-4, 4)
		dr.Field_74_4 = byte(r1)
		d := dr.UnionEffect()
		d.Field_108 = uint32(pos.X) << 12
		d.Field_109 = uint32(pos.Y) << 12
		d.Field_110 = uint32(r2)
		d.Field_111 = c.srv.Frame()
		d.Field_112 = c.srv.Frame() + uint32(expire)
		dr.ZVal = uint16(z)
		dr.VelZ = int8(vz)
		c.Objs.List34Add(dr)
	}
}
