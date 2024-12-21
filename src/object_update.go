package opennox

import (
	"unsafe"

	"github.com/noxworld-dev/opennox-lib/object"
	"github.com/noxworld-dev/opennox-lib/player"
	"github.com/noxworld-dev/opennox-lib/script"
	"github.com/noxworld-dev/opennox-lib/spell"
	"github.com/noxworld-dev/opennox-lib/types"

	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"github.com/noxworld-dev/opennox/v1/common/sound"
	"github.com/noxworld-dev/opennox/v1/legacy"
	"github.com/noxworld-dev/opennox/v1/server"
)

var (
	noxPixieObjID            int
	spellTimeout             uint32
	resetCountdownPerPhoneme = false
)

func init() {
	configBoolPtr("game.extensions.reset_countdown_per_phoneme", "", false, &resetCountdownPerPhoneme)
}

func nox_xxx_updatePlayer_4F8100(u *server.Object) {
	s := noxServer
	ud := u.UpdateDataPlayer()
	h := u.HealthData
	for i := 0; i < 4; i++ {
		p := asObjectS(ud.Field29[i])
		if p != nil && p.Flags().Has(object.FlagDestroyed) {
			ud.Field29[i] = nil
		}
	}
	if u.Flags().Has(object.FlagDestroyed) {
		return
	}
	if noxflags.HasGame(noxflags.GameModeQuest) && ud.Trade70 != nil {
		u.ForceVec = types.Pointf{}
		u.VelVec = types.Pointf{}
	}
	if noxflags.HasGame(noxflags.GameModeQuest) && ud.Field137 != 0 && ud.Player.Index() != server.HostPlayerIndex && (s.Frame()-ud.Field137 > s.SecToFrames(30)) {
		sub_4DCFB0(u)
		return
	}
	v2 := 0
	if ud.Field19_1 != 0 {
		ud.Field19_1--
	} else {
		if ud.Field19_0 != 0 {
			v2 = 1000 * (int(ud.Field19_0) - int(h.Cur)) / int(ud.Field19_0)
		}
		ud.Field19_0 = h.Cur
		if v2 > 0 {
			ud.Field19_1 = 7
		}
	}
	if noxflags.HasGame(noxflags.GameSuddenDeath) {
		playerSuddedDeath4F9E70(u)
	}
	sub_4F9ED0(u)
	pl := ud.Player
	u2 := pl.CameraTarget()
	if u2 == nil {
		u2 = pl.PlayerUnit
	}
	pl.SetPos3632(u2.Pos())
	if ud.Field40_0 != 0 {
		ud.Field40_0--
	}
	u.NeedSync()
	if ud.Field20_1 != 0 {
		ud.Field20_1--
	}
	if !u.Flags().Has(object.FlagDead) {
		if v2 > 0 {
			v14 := u.Field131
			if pl.Info().IsFemale() {
				if v14 == 5 {
					s.Audio.EventObj(sound.SoundHumanFemaleHurtPoison, u, 0, 0)
				} else if v2 <= 450 {
					if v2 <= 70 {
						s.Audio.EventObj(sound.SoundHumanFemaleHurtLight, u, 0, 0)
					} else {
						s.Audio.EventObj(sound.SoundHumanFemaleHurtMedium, u, 0, 0)
					}
				} else {
					s.Audio.EventObj(sound.SoundHumanFemaleHurtHeavy, u, 0, 0)
				}
			} else {
				if v14 == 5 {
					s.Audio.EventObj(sound.SoundHumanMaleHurtPoison, u, 0, 0)
				} else if v2 <= 450 {
					if v2 <= 70 {
						s.Audio.EventObj(sound.SoundHumanMaleHurtLight, u, 0, 0)
					} else {
						s.Audio.EventObj(sound.SoundHumanMaleHurtMedium, u, 0, 0)
					}
				} else {
					s.Audio.EventObj(sound.SoundHumanMaleHurtHeavy, u, 0, 0)
				}
			}
		}
		if ud.Field22_3 < 100 {
			ud.Field22_3 += uint8(100 / s.TickRate())
		}
	}
	if ud.SpellCastStart != 0 && ud.Field47_0 == 0 && (s.Frame()-ud.SpellCastStart) > spellTimeout {
		s.PlayerSpell(u) // (manual?) spell casting
		ud.SpellCastStart = 0
	}
	nox_xxx_playerInventory_4F8420(u)
	if oa1, ov68, ok := s.unitUpdatePlayerImplA(u); ok {
		s.unitUpdatePlayerImplB(u, oa1, ov68)
	}
	if u.HasEnchant(server.ENCHANT_RUN) && ud.State != server.PlayerState1 {
		nox_xxx_playerSetState_4FA020(u, server.PlayerState5)
	}
	legacy.Nox_xxx_questCheckSecretArea_421C70(u)
	s.abilities.harpoon.UpdatePlayer(u)
}

func (s *Server) unitUpdatePlayerImplA(u *server.Object) (a1, v68 bool, _ bool) {
	ud := u.UpdateDataPlayer()
	pl := ud.Player
	switch ud.State {
	default:
		return a1, v68, true
	case server.PlayerState0, server.PlayerState5:
		if legacy.Nox_xxx_playerCanMove_4F9BC0(u) == 0 {
			return a1, v68, true
		}
		if pl.Field3656 != 0 {
			if pl.Info().IsFemale() {
				s.Audio.EventObj(sound.SoundHumanFemaleExertionHeavy, u, 0, 0)
			} else {
				s.Audio.EventObj(sound.SoundHumanMaleExertionHeavy, u, 0, 0)
			}
			s.NetInformTextMsg(pl.PlayerIndex(), 13, 3)
			return a1, v68, true
		}
		v68 = true
		dp := pl.CursorPos().Sub(u.Pos())
		dx := float64(dp.X)
		dy := float64(dp.Y)
		a1 = false
		const runCursorDist = 100
		if !(ud.State != server.PlayerState5 && (dy*dy+dx*dx <= runCursorDist*runCursorDist) || s.Abils.IsActive(u, server.AbilityTreadLightly)) {
			// switch from walking to running
			a1 = true
			u.SpeedCur *= 2
			v67, v69 := s.PlayerAnimFrames(6)
			v25a := int(u.NetCode) + int(noxServer.Frame())
			v25 := v25a / (v69 + 1) % v67
			if !(v25 <= ((v25a-1)/(v69+1)%v67) || v25 != 2 && v25 != 8) {
				tiles := legacy.Get_nox_tile_defs_arr()
				if ti := legacy.Nox_xxx_tileNFromPoint_411160(u.Pos()); ti >= 0 && ti < len(tiles) {
					// emit sound based on the tile material
					switch tiles[ti].Field36 {
					case 2:
						// nop
					case 8:
						s.Audio.EventObj(sound.SoundRunOnWood, u, 0, 0)
					case 64:
						s.Audio.EventObj(sound.SoundRunOnDirt, u, 0, 0)
					case 128:
						s.Audio.EventObj(sound.SoundRunOnWater, u, 0, 0)
					case 0x400:
						s.Audio.EventObj(sound.SoundRunOnSnow, u, 0, 0)
					case 0x800:
						s.Audio.EventObj(sound.SoundRunOnMud, u, 0, 0)
					case 0x4000:
						// nop
					default:
						s.Audio.EventObj(sound.SoundRunOnStone, u, 0, 0)
					}
				}
			}
			if s.Rand.Logic.IntClamp(0, 100) <= 1 {
				s.Audio.EventObj(sound.SoundHumanMaleExertionLight, u, 0, 0)
			}
		}
		if legacy.Sub_4F9AB0(u) == 0 {
			if u.HasEnchant(server.ENCHANT_CONFUSED) {
				u.Direction2 = legacy.Nox_xxx_playerConfusedGetDirection_4F7A40(u)
			}
			// update force based on direction, speed, etc
			u.ForceVec = u.ForceVec.Add(u.Direction2.Vec().Mul(u.SpeedCur))
		}
		if ud.State == server.PlayerState0 {
			v67, v69 := s.PlayerAnimFrames(4)
			v31 := int(u.NetCode) + int(noxServer.Frame())
			v32 := (v31 - 1) / (v69 + 1) % v67
			v33 := v31 / (v69 + 1) % v67
			if (!s.Abils.IsActive(u, server.AbilityTreadLightly) || a1) && v33 != v32 && (v33 == 3 || v33 == 9) {
				tiles := legacy.Get_nox_tile_defs_arr()
				if ti := legacy.Nox_xxx_tileNFromPoint_411160(u.Pos()); ti >= 0 && ti < len(tiles) {
					switch tiles[ti].Field36 {
					case 2:
						// nop
					case 8:
						s.Audio.EventObj(sound.SoundWalkOnWood, u, 0, 0)
					case 64:
						s.Audio.EventObj(sound.SoundWalkOnDirt, u, 0, 0)
					case 128:
						s.Audio.EventObj(sound.SoundWalkOnWater, u, 0, 0)
					case 0x400:
						s.Audio.EventObj(sound.SoundWalkOnSnow, u, 0, 0)
					case 0x800:
						s.Audio.EventObj(sound.SoundWalkOnMud, u, 0, 0)
					case 0x4000:
						// nop
					default:
						s.Audio.EventObj(sound.SoundWalkOnStone, u, 0, 0)
					}
				}
			}
		}
		return a1, v68, true
	case server.PlayerState1:
		if legacy.Nox_xxx_playerAttack_538960(u) == 0 {
			if pl.WeaponEquip&4 != 0 {
				nox_xxx_playerSetState_4FA020(u, server.PlayerState14)
				u.Field34 = s.Frame()
			} else {
				nox_xxx_playerSetState_4FA020(u, server.PlayerState13)
				pl.Field8 &^= 0xff
			}
		}
		return a1, v68, true
	case server.PlayerState2:
		v67, v69 := s.PlayerAnimFrames(21)
		ud.Field59_0 = uint8((int(s.Frame()) - int(u.Field34)) / (v69 + 1))
		if int(ud.Field59_0) >= v67 {
			ud.Field59_0 = uint8(v67 - 1)
		}
		return a1, v68, true
	case server.PlayerState3:
		if (int(s.Frame()) - int(u.Field34)) > int(s.TickRate()) {
			nox_xxx_playerSetState_4FA020(u, server.PlayerState4)
			ud.Field60 &^= 0x20
			u.Field34 = s.Frame()
			u.ObjFlags |= object.FlagShort | object.FlagAllowOverlap
			u.VelVec = types.Pointf{}
			u.ForceVec = types.Pointf{}
			u.Pos24 = types.Pointf{}
			s.CallOnPlayerDeath(u.ControllingPlayer(), nil) // TODO: set the killer
			s.scriptOnEvent(script.EventPlayerDeath)
		}
		return a1, v68, false
	case server.PlayerState4:
		if (int(s.Frame()) - int(u.Field34)) <= int(s.TickRate())/2 {
			return a1, v68, false
		}
		v41 := int(legacy.Nox_xxx_servGamedataGet_40A020(1024))
		if !noxflags.HasGame(noxflags.GameModeElimination) || (v41 <= 0) || (int(pl.Field2140) < v41) {
			if noxflags.HasGame(noxflags.GameOnline) && (pl.Field3680&1 == 0) {
				cb := s.Players.Control.Player(pl.Index())
				for it := cb.First(); it != nil; it = cb.Next() {
					if it.Code == player.CCAction {
						cb.Reset()
						legacy.Nox_xxx_playerRespawn_4F7EF0(u)
						return a1, v68, true
					}
				}
			}
			if legacy.Nox_server_doPlayersAutoRespawn_40A5F0() == 0 {
				return a1, v68, false
			}
			legacy.Nox_xxx_playerRespawn_4F7EF0(u)
			return a1, v68, true
		}
		if pl.Field3680&1 != 0 {
			a1 = pl.CameraTarget() != nil
			pl.CameraUnlock()
			for _, it := range s.Players.ListUnits() {
				pl2 := s.Players.ByID(int(it.NetCode))
				if !it.Flags().Has(object.FlagDead) && (pl2.Field3680&1 == 0) {
					pl.CameraToggle(it)
				}
			}
		} else {
			legacy.Nox_xxx_netNeedTimestampStatus_4174F0(pl, 32)
			s.PlayerGoObserver(pl, false, false)
			pl.CameraUnlock()
			s.nox_xxx_playerLeaveObsByObserved_4E60A0(u)
			if legacy.Sub_4F9E10(u) == 0 {
				for _, it := range s.Players.ListUnits() {
					pl2 := s.Players.ByID(int(it.NetCode))
					if !it.Flags().Has(object.FlagDead) && (pl2.Field3680&1 == 0) {
						pl.CameraToggle(it)
					}
				}
			}
		}
		return a1, v68, false
	case server.PlayerState10:
		ud.Field59_0 = 0
		return a1, v68, true
	case server.PlayerState12:
		v67, v69 := s.PlayerAnimFrames(3)
		v49 := (int(s.Frame()) - int(u.Field34)) / (v69 + 1)

		found := false
		for _, it := range s.Players.ListUnits() {
			ud2 := it.UpdateDataPlayer()
			if ud2.HarpoonTarg == u {
				found = true
				break
			}
		}
		if !found {
			u.ForceVec = u.Direction1.Vec().Mul(2 * u.SpeedCur)
		}
		if v49 >= v67 {
			// stop hovering after a jump?
			nox_xxx_playerSetState_4FA020(u, server.PlayerState0)
			u.ObjFlags &= 0xFFFFBFFF
			u.Field34 = s.Frame()
		}
		a1 = v69 != 0
		return a1, v68, false
	case server.PlayerState13:
		u.ObjFlags &= 0xFFFFBFFE
		if legacy.Sub_4F9A80(u) != 0 {
			nox_xxx_playerSetState_4FA020(u, server.PlayerState0)
		}
		if noxflags.HasGame(noxflags.GameModeChat) || (pl.ArmorEquip&0x3000000 == 0) || legacy.Nox_xxx_monsterTestBlockShield_533E70(u) == 0 &&
			(int(s.Frame())-int(u.Field34)) <= int(s.TickRate())/4 {
			return a1, v68, true
		}
		nox_xxx_playerSetState_4FA020(u, server.PlayerState15)
		ud.Field59_0 = 0
		return a1, v68, true
	case server.PlayerState14:
		v69, _ := s.PlayerAnimFrames(33)
		ud.Field59_0 = uint8(v69 - 1)
		if int(s.Frame())-int(u.Field34) > int(s.TickRate()) {
			nox_xxx_playerSetState_4FA020(u, server.PlayerState13)
		}
		return a1, v68, true
	case server.PlayerState15:
		v67, v69 := s.PlayerAnimFrames(40)
		ud.Field59_0 = uint8((int(s.Frame()) - int(u.Field34)) / (v69 + 1))
		if int(ud.Field59_0) >= v67 {
			nox_xxx_playerSetState_4FA020(u, server.PlayerState16)
			ud.Field59_0 = uint8(v67 - 1)
		}
		return a1, v68, true
	case server.PlayerState16:
		v69, _ := s.PlayerAnimFrames(40)
		ud.Field59_0 = uint8(v69 - 1)
		return a1, v68, true
	case server.PlayerState17:
		v67, v69 := s.PlayerAnimFrames(40)
		v11 := v67 - (int(s.Frame())-int(u.Field34))/(v69+1)
		if v11 >= v67 {
			ud.Field59_0 = uint8(v67 - 1)
		} else {
			if v11 <= 0 {
				v11 = 0
				nox_xxx_playerSetState_4FA020(u, server.PlayerState13)
			}
			ud.Field59_0 = uint8(v11)
		}
		return a1, v68, true
	case server.PlayerState18:
		v67, v69 := s.PlayerAnimFrames(48)
		ud.Field59_0 = uint8((int(s.Frame()) - int(u.Field34)) / (v69 + 1))
		if int(ud.Field59_0) >= v67 {
			nox_xxx_playerSetState_4FA020(u, server.PlayerState13)
		}
		return a1, v68, true
	case server.PlayerState19:
		v67, v69 := s.PlayerAnimFrames(49)
		ud.Field59_0 = uint8((int(s.Frame()) - int(u.Field34)) / (v69 + 1))
		if int(ud.Field59_0) >= v67 {
			nox_xxx_playerSetState_4FA020(u, server.PlayerState13)
		}
		return a1, v68, true
	case server.PlayerState20:
		v67, v69 := s.PlayerAnimFrames(47)
		ud.Field59_0 = uint8((int(s.Frame()) - int(u.Field34)) / (v69 + 1))
		if int(ud.Field59_0) >= v67 {
			nox_xxx_playerSetState_4FA020(u, server.PlayerState13)
		}
		return a1, v68, true
	case server.PlayerState21:
		v69, v67 := s.PlayerAnimFrames(30)
		ud.Field59_0 = uint8((int(s.Frame()) - int(u.Field34)) / (v67 + 1))
		if int(ud.Field59_0) >= v69 {
			nox_xxx_playerSetState_4FA020(u, server.PlayerState13)
		}
		return a1, v68, true
	case server.PlayerState22:
		v69, _ := s.PlayerAnimFrames(31)
		ud.Field59_0 = uint8(v69 - 1)
		return a1, v68, true
	case server.PlayerState23:
		v67, v69 := s.PlayerAnimFrames(50)
		ud.Field59_0 = uint8((int(s.Frame()) - int(u.Field34)) / (v69 + 1))
		if int(ud.Field59_0) >= v67 {
			nox_xxx_playerSetState_4FA020(u, server.PlayerState13)
		}
		return a1, v68, true
	case server.PlayerState24:
		v67, v69 := s.PlayerAnimFrames(19)
		ud.Field59_0 = uint8((int(s.Frame()) - int(u.Field34)) / (v69 + 1))
		if int(ud.Field59_0) >= v67 {
			nox_xxx_playerSetState_4FA020(u, server.PlayerState13)
		}
		return a1, v68, true
	case server.PlayerStateShakeFist:
		v67, v69 := s.PlayerAnimFrames(20)
		ud.Field59_0 = uint8((int(s.Frame()) - int(u.Field34)) / (v69 + 1))
		if int(ud.Field59_0) >= v67 {
			nox_xxx_playerSetState_4FA020(u, server.PlayerState13)
		}
		return a1, v68, true
	case server.PlayerStateLaugh:
		v67, v69 := s.PlayerAnimFrames(15)
		ud.Field59_0 = uint8((int(s.Frame()) - int(u.Field34)) / (v69 + 1))
		if int(ud.Field59_0) >= v67 {
			nox_xxx_playerSetState_4FA020(u, server.PlayerState13)
		}
		return a1, v68, true
	case server.PlayerState27:
		v67, v69 := s.PlayerAnimFrames(16)
		ud.Field59_0 = uint8((int(s.Frame()) - int(u.Field34)) / (v69 + 1))
		if int(ud.Field59_0) >= v67/2 {
			nox_xxx_playerSetState_4FA020(u, server.PlayerStatePoint)
			ud.Field59_0 = uint8(v67 / 2)
		}
		return a1, v68, true
	case server.PlayerStatePoint:
		v67, _ := s.PlayerAnimFrames(16)
		ud.Field59_0 = uint8(v67 / 2)
		if (int(s.Frame()) - int(u.Field34)) > 0x14 {
			nox_xxx_playerSetState_4FA020(u, server.PlayerState29)
			ud.Field59_0 = uint8(v67 / 2)
		}
		return a1, v68, true
	case server.PlayerState29:
		v67, v69 := s.PlayerAnimFrames(16)
		ud.Field59_0 = uint8(v67/2 + (int(s.Frame())-int(u.Field34))/(v69+1))
		if int(ud.Field59_0) >= v67 {
			nox_xxx_playerSetState_4FA020(u, server.PlayerState13)
		}
		return a1, v68, true
	case server.PlayerState30:
		v67, v69 := s.PlayerAnimFrames(52)
		ud.Field59_0 = uint8((int(s.Frame()) - int(u.Field34)) / (v69 + 1))
		if int(ud.Field59_0) >= v67 {
			nox_xxx_playerSetState_4FA020(u, server.PlayerState13)
			ud.Field41 = 0
		}
		return a1, v68, true
	case server.PlayerState32:
		v67, _ := s.PlayerAnimFrames(54)
		ud.Field59_0 = uint8(v67 / 2)
		if (int(s.Frame()) - int(u.Field34)) > 0x14 {
			nox_xxx_playerSetState_4FA020(u, server.PlayerState33)
			ud.Field59_0 = uint8(v67 / 2)
		}
		return a1, v68, true
	case server.PlayerState33:
		v67, v69 := s.PlayerAnimFrames(54)
		ud.Field59_0 = uint8(v67/2 + (int(s.Frame())-int(u.Field34))/(v69+1))
		if int(ud.Field59_0) >= v67 {
			nox_xxx_playerSetState_4FA020(u, server.PlayerState13)
		}
		return a1, v68, true
	}
}

func (s *Server) unitUpdatePlayerImplB(u *server.Object, a1, v68 bool) {
	ud := u.UpdateDataPlayer()
	pl := ud.Player
	orientationOnly := false
	cb := s.Players.Control.Player(pl.Index())
	if cb.IsEmpty() {
		goto LABEL_247
	}
	if (ud.State == server.PlayerState0 || ud.State == server.PlayerState5) && legacy.Sub_4F9A80(u) == 0 {
		nox_xxx_playerSetState_4FA020(u, server.PlayerState13)
		u.Field34 = s.Frame()
	}
	ud.Field60 &^= 0x2 | 0x4 | 0x8 | 0x10
	if pl.Field3680&3 != 0 {
		goto LABEL_247
	}
	orientationOnly = s.Spells.Dur.Sub4FEE50(spell.SPELL_FORCE_OF_NATURE, u)
	for it := cb.First(); it != nil; it = cb.Next() {
		if orientationOnly && it.Code != player.CCOrientation {
			continue
		}
		switch it.Code {
		case player.CCOrientation:
			if !u.HasEnchant(server.ENCHANT_FREEZE) &&
				(!noxflags.HasGame(noxflags.GameModeQuest) || ud.Trade70 == nil) &&
				!s.Abils.IsActive(u, server.AbilityBerserk) {
				u.Direction2 = server.Dir16(it.Uint16())
			}
		case player.CCMoveForward, player.CCMoveBackward, player.CCMoveLeft, player.CCMoveRight:
			if legacy.Nox_xxx_playerCanMove_4F9BC0(u) != 0 {
				legacy.Nox_xxx_cancelAllSpells_4FEE90(u)
				if !s.Abils.IsActive(u, server.AbilityBerserk) &&
					(ud.State != server.PlayerState1 || (pl.WeaponEquip&0x47F0000 != 0) && legacy.Nox_common_mapPlrActionToStateId_4FA2B0(u) != 29) {
					if ud.State == server.PlayerState16 {
						nox_xxx_playerSetState_4FA020(u, server.PlayerState17)
					} else {
						if a1 {
							nox_xxx_playerSetState_4FA020(u, server.PlayerState5)
						} else {
							nox_xxx_playerSetState_4FA020(u, server.PlayerState0)
						}
						if it.Uint8()&2 != 0 {
							ud.Field60 |= 0x1
						} else {
							ud.Field60 &^= 0x1
						}
						switch it.Code {
						case player.CCMoveForward:
							ud.Field60 |= 0x8
						case player.CCMoveBackward:
							ud.Field60 |= 0x10
						case player.CCMoveLeft:
							ud.Field60 |= 0x4
						case player.CCMoveRight:
							ud.Field60 |= 0x2
						}
						u.Field34 = s.Frame()
					}
				}
			}
		case player.CCAction:
			if legacy.Nox_xxx_playerCanAttack_4F9C40(u) != 0 {
				if !noxflags.HasGame(noxflags.GameModeChat) && legacy.Nox_xxx_checkWinkFlags_4F7DF0(u) == 0 {
					legacy.Nox_xxx_playerInputAttack_4F9C70(u)
				}
				if ud.State == server.PlayerState10 {
					nox_xxx_playerSetState_4FA020(u, server.PlayerState13)
				}
			}
		case player.CCJump:
			if legacy.Nox_xxx_playerCanMove_4F9BC0(u) == 0 || s.Abils.IsActive(u, server.AbilityBerserk) ||
				s.Abils.IsActiveVal(u, server.AbilityWarcry) {
				break
			}
			legacy.Nox_xxx_cancelAllSpells_4FEE90(u)
			if pl.Field3656 != 0 {
				if pl.Info().IsFemale() {
					s.Audio.EventObj(sound.SoundHumanFemaleExertionHeavy, u, 0, 0)
				} else {
					s.Audio.EventObj(sound.SoundHumanMaleExertionHeavy, u, 0, 0)
				}
				s.NetInformTextMsg(pl.PlayerIndex(), 13, 3)
			} else if legacy.Nox_xxx_playerSubStamina_4F7D30(u, 90) != 0 {
				if u.HasEnchant(server.ENCHANT_CONFUSED) {
					u.Direction2 = legacy.Nox_xxx_playerConfusedGetDirection_4F7A40(u)
				}
				u.ObjFlags |= 0x4000
				nox_xxx_playerSetState_4FA020(u, server.PlayerState12)
				u.Field34 = s.Frame()
				return
			}
		case player.CCSpellGestureUp, player.CCSpellGestureDown,
			player.CCSpellGestureLeft, player.CCSpellGestureRight,
			player.CCSpellGestureUpperRight, player.CCSpellGestureUpperLeft,
			player.CCSpellGestureLowerRight, player.CCSpellGestureLowerLeft:
			// If the appropriate flag is set, reset countdown for manual casting
			// every phoneme press
			u.PlayerActionPhoneme(it.Code, resetCountdownPerPhoneme)
		case player.CCSpellPatternEnd:
			nox_xxx_playerSetState_4FA020(u, server.PlayerState13)
			if !noxflags.HasGame(noxflags.GameModeChat) {
				if ud.SpellCastStart != 0 {
					s.PlayerSpell(u)
					ud.SpellCastStart = 0
				} else {
					targ := s.getObjectFromNetCode(int(it.Uint32()))
					legacy.Nox_xxx_playerDoSchedSpell_4FB0E0(u, targ)
				}
			}
		case player.CCCastQueuedSpell:
			nox_xxx_playerSetState_4FA020(u, server.PlayerState13)
			if !noxflags.HasGame(noxflags.GameModeChat) {
				if ud.SpellCastStart != 0 {
					s.PlayerSpell(u)
					ud.SpellCastStart = 0
				}
				ud.Field55 = pl.CursorVec.X
				ud.Field56 = pl.CursorVec.Y
				targ := s.getObjectFromNetCode(int(it.Uint32()))
				legacy.Nox_xxx_playerDoSchedSpell_4FB0E0(u, targ)
			}
		case player.CCCastMostRecentSpell:
			if !noxflags.HasGame(noxflags.GameModeChat) {
				if ud.SpellCastStart != 0 {
					s.PlayerSpell(u)
					ud.SpellCastStart = 0
				}
				ud.Field55 = pl.CursorVec.X
				ud.Field56 = pl.CursorVec.Y
				targ := s.getObjectFromNetCode(int(it.Uint32()))
				legacy.Nox_xxx_playerDoSchedSpellQueue_4FB1D0(u, targ)
			}
		}
	}

LABEL_247:
	if v68 && ud.State != server.PlayerState0 && ud.State != server.PlayerState5 {
		if s.Abils.IsActive(u, server.AbilityTreadLightly) {
			s.abilities.DisableAbility(u, server.AbilityTreadLightly)
		}
	}
}

func sub_4E7540(a1, a2 server.Obj) {
	legacy.Sub_4E7540(server.ToObject(a1), server.ToObject(a2))
}

func nox_xxx_playerInventory_4F8420(u *server.Object) {
	for it := u.FirstItem(); it != nil; it = it.NextItem() {
		if it.Flags().Has(object.FlagEquipped) {
			if !legacy.Nox_xxx_playerCheckStrength_4F3180(u, it) {
				asObjectS(u).forceDrop(it)
			}
		}
	}
}

func (s *Server) ApplyForce(obj *server.Object, vec types.Pointf, force float64) {
	asObjectS(obj).applyForce(vec, force)
}

func (obj *Object) applyForce(vec types.Pointf, force float64) { // nox_xxx_objectApplyForce_52DF80
	if !obj.IsMovable() {
		return
	}
	dp := obj.Pos().Sub(vec)
	r := float32(dp.Len() + 0.1)
	f := 10.0 * float32(force) / obj.Mass
	// This weird conversion is how Nox is doing it.
	// Be aware that changing it may cause minor deviation in physics.
	obj.ForceVec.X += float32(float64(dp.X) * float64(f) / float64(r))
	obj.ForceVec.Y += float32(float64(dp.Y) * float64(f) / float64(r))
	if !obj.Class().Has(object.ClassMissile) {
		legacy.Nox_xxx_unitHasCollideOrUpdateFn_537610(obj.SObj())
	}
}

func playerSuddedDeath4F9E70(u *server.Object) {
	v1 := memmap.Uint32(0x5D4594, 1392)
	h := u.HealthData
	if !u.Flags().Has(object.FlagDead) && h != nil && h.Cur != 0 && (noxServer.Frame()%(v1*noxServer.TickRate()/uint32(h.Max))) == 0 {
		legacy.Nox_xxx_unitDamageClear_4EE5E0(u, 1)
	}
}

func sub_4F9ED0(u *server.Object) {
	s := u.Server()
	ud := u.UpdateDataPlayer()
	h := u.HealthData
	if u.Flags().Has(object.FlagDead) {
		return
	}
	if h != nil && (s.Frame()-u.Frame134) > s.TickRate() {
		if h.Cur < h.Max && h.Max != 0 {
			div := 300 * s.TickRate() / uint32(h.Max)
			if div == 0 {
				div = 1
			}
			if s.Frame()%div == 0 {
				legacy.Nox_xxx_unitAdjustHP_4EE460(u, 1)
			}
		}
	}
	if ud.ManaCur < ud.ManaMax && ud.ManaMax != 0 {
		div := 300 * s.TickRate() / uint32(ud.ManaMax)
		if div == 0 {
			div = 1
		}
		if s.Frame()%div == 0 {
			legacy.Nox_xxx_playerManaAdd_4EEB80(u, 1)
		}
	}
}

func nox_xxx_updatePixie_53CD20(u *server.Object) {
	s := u.Server()
	ss := noxServer
	ud := unsafe.Slice((*uint32)(u.UpdateData), 7)
	if memmap.Uint32(0x5D4594, 2488696) == 0 {
		dt := s.Balance.Float("PixieReturnTimeout")
		*memmap.PtrUint32(0x5D4594, 2488696) = s.SecToFramesF(dt)
	}

	if deadline := ud[5]; deadline != 0 && s.Frame() > deadline {
		asObjectS(u).Delete()
		return
	}

	if targ := asObjectS(*(**server.Object)(unsafe.Pointer(&ud[1]))); targ != nil {
		if targ.Flags().HasAny(object.FlagDestroyed | object.FlagDead) {
			ud[1] = 0
		}
	}
	owner := u.ObjOwner
	if u.Flags().Has(object.FlagEnabled) {
		if s.Frame()-u.Field34 > s.TickRate()/4 {
			targ := s.PixieFindTarget(u)
			*(**server.Object)(unsafe.Pointer(&ud[1])) = targ
			if targ == owner {
				ud[1] = 0
			}
			u.Field34 = s.Frame()
		}
	} else {
		ud[1] = 0
	}
	if owner != nil && owner.Class().HasAny(object.ClassPlayer) && owner.Flags().HasAny(object.FlagNoUpdate) {
		ud[1] = 0
	}
	if targ := asObjectS(*(**server.Object)(unsafe.Pointer(&ud[1]))); targ != nil {
		server.PixieIdleAnimate(u, targ.Pos().Sub(u.Pos()), 32)
	} else {
		s.Map.EachMissileInCircle(u.PosVec, 200.0, func(it *server.Object) bool {
			if noxPixieObjID == 0 {
				noxPixieObjID = noxServer.Types.PixieID()
			}
			if int(it.TypeInd) != noxPixieObjID {
				return true
			}
			if it != u && u.ObjOwner == it.ObjOwner {
				server.PixieIdleAnimate(u, it.Pos().Sub(u.Pos()), 16)
			}
			return true
		})
		if owner != nil {
			pos1, pos2 := u.Pos(), owner.Pos()
			if s.MapTraceRay9(pos1, pos2) {
				server.PixieIdleAnimate(u, pos2.Sub(pos1), 25)
			}
		} else {
			pos1, pos2 := u.Pos(), u.Pos39
			if s.MapTraceRay9(pos1, pos2) {
				server.PixieIdleAnimate(u, pos2.Sub(pos1), 25)
			}
		}
	}
	u.Float28 = 0.9
	u.ForceVec = u.Direction2.Vec().Mul(u.SpeedCur)
	if (s.Frame()&8 == 0) && owner != nil {
		if s.MapTraceVision(u, owner) {
			ud[6] = s.Frame()
		}
		if s.Frame()-ud[6] > memmap.Uint32(0x5D4594, 2488696) {
			ss.nox_xxx_teleportPixie_4FD050(u, owner)
			ud[6] = s.Frame()
		}
	}
}

func (s *Server) nox_xxx_teleportPixie_4FD050(u *server.Object, owner *server.Object) {
	pos := owner.Pos()
	u.PosVec = pos
	u.PrevPos = pos
	u.NewPos = pos
	s.nox_xxx_moveUpdateSpecial_517970(u)
}

func (s *Server) nox_xxx_teleportAllPixies_4FD090(u *server.Object) {
	for it := u.FirstOwned516(); it != nil; it = it.NextOwned512() {
		if int(it.TypeInd) != noxPixieObjID {
			continue
		}
		if it.Flags().HasAny(object.FlagDead) {
			continue
		}
		if *(*uint32)(unsafe.Add(it.UpdateData, 4)) == 0 {
			s.nox_xxx_teleportPixie_4FD050(it, u)
		}
	}
}

func nox_xxx_updatePlayerObserver_4E62F0(a1p *server.Object) {
	s := noxServer
	u := asObjectS(a1p)
	ud := u.UpdateDataPlayer()
	pl := ud.Player
	for i := range ud.Field29 {
		it := asObjectS(ud.Field29[i])
		if it != nil && it.Flags().Has(object.FlagDestroyed) {
			ud.Field29[i] = nil
		}
	}
	u.NeedSync()
	if targ := pl.CameraTarget(); targ != nil {
		pl.SetPos3632(targ.Pos())
	}
	cb := s.Players.Control.Player(pl.Index())
	if cb.First() == nil {
		return
	}
	pl.Field3688 = 0
	for it := cb.First(); it != nil; it = cb.Next() {
		if it.Code == player.CCMoveForward {
			if pl.Field3672 == 0 {
				pl.Field3688 = 1
				if pl.Field3692 == 0 {
					s.PlayerLeaveMonsterObserver(pl)
				}
				it.Active = false
			} else if pl.Field3672 == 1 {
				const max = 30
				dp := pl.Pos3632().Sub(pl.CursorPos())
				opos := pl.Pos3632()
				if dp.X > max {
					opos.X -= (dp.X - max) * 0.1
				} else if dp.X < -max {
					opos.X -= (dp.X + max) * 0.1
				}
				if dp.Y > max {
					opos.Y -= (dp.Y - max) * 0.1
				} else if dp.Y < -max {
					opos.Y -= (dp.Y + max) * 0.1
				}
				if s.Map.ValidIndexPos(opos) {
					pl.SetPos3632(opos)
				}
			}
			continue
		}
		if it.Code != player.CCAction {
			if it.Code != player.CCJump {
				continue
			}
			if pl.ObserveTarget() == nil && !noxflags.HasGame(noxflags.GameModeQuest) {
				pos2 := pl.Pos3632()
				var (
					found *Object
					min   = float32(1e+08)
				)
				rect := types.RectFromPointsf(pos2.Sub(types.Pointf{X: 100, Y: 100}), pos2.Add(types.Pointf{X: 100, Y: 100}))
				s.Map.EachObjInRect(rect, func(obj *server.Object) bool {
					if obj.Class().Has(object.ClassMonster) && obj.ObjOwner != nil && obj.ObjOwner == u.SObj() {
						dp := obj.Pos().Sub(pos2)
						dist := dp.X*dp.X + dp.Y*dp.Y
						if dist < min {
							found = asObjectS(obj)
							min = dist
						}
					}
					return true
				})
				if found != nil && found.SObj() != pl.CameraTarget() {
					pl.CameraToggle(found)
					pl.Field3672 = 0
				} else {
					pl.CameraUnlock()
					pl.Field3672 = 1
				}
				continue
			}
		}
		if legacy.Get_dword_5d4594_2650652() != 0 && noxflags.HasGame(noxflags.GameFlag15|noxflags.GameFlag16) && legacy.Sub_509CF0(&pl.Field2096Buf[0], pl.PlayerClass(), pl.Field2068) == 0 {
			s.NetInformTextMsg(pl.PlayerIndex(), 17, 0)
			it.Active = false
			continue
		}
		if pl.Field3680&0x20 != 0 {
			s.PlayerLeaveMonsterObserver(pl)
			it.Active = false
			continue
		}
		if noxflags.HasGame(noxflags.GameModeQuest) {
			if pl.Field4792 == 0 {
				if ud.Field138 == 1 {
					s.NetPriMsgToPlayer(u.SObj(), "MainBG.wnd:Loading", 0)
				} else {
					pl.Field4792 = uint32(legacy.Sub_4E4100())
					if pl.Field4792 == 1 {
						legacy.Sub_4D79C0(u.SObj())
					} else {
						s.NetPriMsgToPlayer(u.SObj(), "GeneralPrint:QuestGameFull", 0)
					}
				}
			}
			if ud.Field79 != 0 {
				legacy.Sub_4D7480(u.SObj())
				continue
			}
			if ud.Field78 != 0 {
				s.PlayerLeaveMonsterObserver(pl)
				it.Active = false
				continue
			}
			if pl.Field4792 == 0 {
				s.PlayerLeaveMonsterObserver(pl)
				it.Active = false
				continue
			}
		}
		v13 := legacy.Nox_xxx_gamePlayIsAnyPlayers_40A8A0() != 0
		if legacy.Sub_40A740() != 0 || noxflags.HasGame(noxflags.GameFlag16) || (pl.Field3680&0x100 != 0) && v13 {
			if legacy.Sub_40AA70(pl) == 0 {
				s.PlayerLeaveMonsterObserver(pl)
				it.Active = false
				continue
			}
		}
		if noxflags.HasEngine(noxflags.EngineNoRendering) && u.SObj() == s.Players.HostUnit() {
			it.Active = false
			continue
		}
		if pl.ObserveTarget() == nil {
			legacy.Sub_4DF3C0(pl)
			legacy.Nox_xxx_playerLeaveObserver_0_4E6AA0(pl)
			pl.CameraUnlock()
			if !noxflags.HasGame(noxflags.GameModeQuest) {
				v22 := s.nox_xxx_mapFindPlayerStart_4F7AB0(pl.PlayerUnit)
				asObjectS(pl.PlayerUnit).SetPos(v22)
			}
			it.Active = false
			continue
		}
		u.observeClear()
		it.Active = false
	}
	pl.Field3692 = pl.Field3688
}

func nox_xxx_updateProjectile_53AC10(obj *server.Object) {
	if (noxServer.Frame() - obj.Field32) > 40 {
		obj.CallCollide(0, 0)
		asObjectS(obj).Delete()
	}
}
