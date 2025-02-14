package opennox

import (
	"github.com/opennox/libs/object"
	"github.com/opennox/libs/spell"
	"github.com/opennox/libs/things"
	"github.com/opennox/libs/types"

	"github.com/opennox/opennox/v1/server"
)

type spellMissiles struct {
	s    *Server
	proj map[spell.ID]int // map[spellID]typeID, 0x5D4594, 2489136
}

func (sp *spellMissiles) Init(s *Server) {
	sp.s = s
	sp.proj = make(map[spell.ID]int)
}

func (sp *spellMissiles) Free() {
	sp.proj = nil
}

func (sp *spellMissiles) Cast(spellID spell.ID, a2, owner, caster *server.Object, a5 *server.SpellAcceptArg, lvl int) int {
	spl := sp.s.Spells.DefByInd(spellID)
	opts := spl.Def.Missiles.Level(lvl)
	typ, ok := sp.proj[spellID]
	if !ok {
		typ = sp.s.Types.IndByID(opts.Projectile)
		sp.proj[spellID] = typ
	}
	curCnt := owner.CountSubOfType(typ)
	var cnt, maxCnt int
	if opts.Count <= 0 {
		// it's intentionally loading this variable twice
		// looks previously there were two separate config values for it
		cnt = int(sp.s.Balance.FloatInd("MagicMissileCount", lvl-1))
		maxCnt = int(sp.s.Balance.FloatInd("MagicMissileCount", lvl-1))
	} else {
		cnt, maxCnt = opts.Count, opts.Count
	}
	if curCnt+cnt > maxCnt {
		cnt = maxCnt - curCnt
	}
	if cnt <= 0 {
		sp.s.NetPriMsgToPlayer(owner, "mmissile.c:TooManyMissiles", 0)
		return 0
	}
	opts.Count = cnt
	sp.CastCustom(spellID, owner, caster, opts)
	return 1
}

func (sp *spellMissiles) CastCustom(spellID spell.ID, owner, caster *server.Object, opts things.MissilesSpell) {
	cpos := caster.Pos()
	cvel := caster.Vel()
	rdist := caster.Shape.Circle.R + opts.Offset
	for i := 0; i < opts.Count; i++ {
		doff := int16(opts.Spread * uint16((i+1)/2))
		if i%2 == 1 {
			doff = -doff
		}
		dir := server.RoundDir(int(int16(caster.Direction1) + doff))
		dv := dir.Vec()
		p2 := cpos.Add(cvel).Add(dv.Mul(rdist))
		if !sp.s.MapTraceRay(cpos, p2, server.MapTraceFlag1|server.MapTraceFlag3) {
			continue
		}
		msl := sp.s.NewObjectByTypeID(opts.Projectile)
		mud := msl.UpdateDataMissile()
		sp.s.CreateObjectAt(msl, owner, p2)
		mspeed := float32(sp.s.Rand.Logic.FloatClamp(opts.SpeedRndMin, opts.SpeedRndMax) * float64(msl.Speed()))
		msl.SpeedCur = mspeed
		msl.SetDir(dir)
		msl.VelVec = cvel.Add(dv.Mul(mspeed * opts.VelMult))
		var ppos *types.Pointf
		if caster.Class().Has(object.ClassPlayer) {
			pl := caster.ControllingPlayer()
			cur := pl.CursorPos()
			ppos = &cur
		}
		targ := sp.s.Nox_xxx_spellFlySearchTarget(ppos, msl, 0x20, opts.SearchDist, 0, owner)
		mud.Owner = owner
		mud.Target = targ
		mud.SpellID = int32(spellID)
	}
	aud := sp.s.Spells.DefByInd(spellID).GetCastSound()
	sp.s.Audio.EventObj(aud, caster, 0, 0)
}
