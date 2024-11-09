package opennox

import (
	"context"
	"errors"
	"fmt"
	"image"
	"path/filepath"
	"strings"
	"unsafe"

	"github.com/noxworld-dev/nat"
	"github.com/noxworld-dev/noxcrypt"
	"github.com/noxworld-dev/opennox-lib/common"
	"github.com/noxworld-dev/opennox-lib/datapath"
	"github.com/noxworld-dev/opennox-lib/ifs"
	"github.com/noxworld-dev/opennox-lib/log"
	"github.com/noxworld-dev/opennox-lib/noxnet"
	"github.com/noxworld-dev/opennox-lib/object"
	"github.com/noxworld-dev/opennox-lib/script"
	"github.com/noxworld-dev/opennox-lib/spell"
	"github.com/noxworld-dev/opennox-lib/strman"
	"github.com/noxworld-dev/opennox-lib/types"

	"github.com/noxworld-dev/opennox-lib/console"

	"github.com/noxworld-dev/opennox/v1/common/config"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"github.com/noxworld-dev/opennox/v1/common/sound"
	"github.com/noxworld-dev/opennox/v1/internal/cryptfile"
	"github.com/noxworld-dev/opennox/v1/internal/netlist"
	"github.com/noxworld-dev/opennox/v1/legacy"
	"github.com/noxworld-dev/opennox/v1/legacy/cnxz"
	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc"
	"github.com/noxworld-dev/opennox/v1/server"
)

var (
	noxServer *Server
)

func init() {
	legacy.GetServer = func() legacy.Server {
		return noxServer
	}
}

func nox_xxx_checkGameFlagPause_413A50() bool {
	return noxflags.HasGame(noxflags.GamePause)
}

func sub_40A1A0() int {
	return bool2int((legacy.Sub_40A180(noxflags.GetGame()) != 0 || noxServer.flag3592) &&
		memmap.Uint32(0x587000, 4660) != 0 &&
		platformTicks() > memmap.Uint64(0x5D4594, 3468))
}

func init() {
	nat.Log = log.New("nat")
	nat.LogUPNP = log.New("nat-upnp")
}

func NewServer(pr console.Printer, sm *strman.StringManager) *Server {
	s := &Server{
		Server: server.New(pr, sm),
	}
	s.Server.ExtServer = unsafe.Pointer(s)
	s.Server.CurrentMapXxx = s.nox_server_currentMapGetFilename_409B30
	s.Server.CurrentMapYyy = legacy.Nox_xxx_mapGetMapName_409B40
	s.NetSendPacketXxx = legacy.Nox_xxx_netSendPacket_4E5030
	s.NetStr.OnDiscover = nox_xxx_servNetInitialPackets_552A80_discover
	s.NetStr.OnExtPacket = MixRecvFromReplacer
	s.NetStr.GetMaxPlayers = func() int {
		return s.getServerMaxPlayers()
	}
	s.NetXferLocal = netXferLocal
	s.ObjectByNetCode = s.getObjectFromNetCode
	configBoolPtr(config.KeyNetXor, "NOX_NET_XOR", true, &s.NetStr.Xor)
	configBoolPtr(config.KeyNetPortForward, "NOX_NET_NAT", true, &s.Server.UseNAT)
	configHiddenBoolPtr(config.KeyNetDebug, "NOX_DEBUG_NET", &s.NetStr.Debug)
	s.initMetrics()
	s.abilities.Init(s)
	s.ai.Init(s)
	s.noxScript.Init(s)
	s.MapSend.OnEndReceive = legacy.Sub_4DE410
	s.Objs.XFerInvLight = legacy.Get_nox_xxx_XFerInvLight_4F5AA0()
	return s
}

type Server struct {
	*server.Server
	spells          serverSpells
	abilities       serverAbilities
	srvReg          srvReg
	noxScript       noxScript
	ai              aiData
	quest           questServer
	springs         serverSprings
	mapSwitchWPName string

	flag1548704 bool
	flag3592    bool
}

func (s *Server) S() *server.Server {
	return s.Server
}

func (s *Server) GetFlag1548704() bool {
	return s.flag1548704
}

func (s *Server) GetFlag3592() bool {
	return s.flag3592
}

func (s *Server) NoxScriptC() legacy.NoxScript {
	return &s.noxScript
}

func (s *Server) Update() bool {
	defer noxPerfmon.startProfileServer()()
	if !s.Server.Update() {
		if debugMainloop {
			gameLog.Println("gameStateFunc exit")
		}
		return false
	}
	return true
}

func (s *Server) updateUnits() { // nox_xxx_updateUnits_51B100
	s.updateUnitsAAA()
	s.updateUnitsBBB()
	s.updateUnitsCallUpdate()
	s.updateCollide()
	s.updateUnitsCCC()
	legacy.Nox_xxx_updateUnits_51B100_D()
	legacy.Nox_xxx_decay_511750()
	legacy.Nox_server_checkVictory_509A60()
}

func (s *Server) updateCollide() {
	legacy.Nox_xxx_allocHitArray_5486D0()
	for i := 0; i < 5; i++ {
		legacy.Nox_xxx_updateObjectsVelocity_5118A0(0.2)
		legacy.Sub_548B60()
	}
	for legacy.Get_dword_5d4594_2488604() != 0 {
		obj := legacy.Sub_537700()
		if int8(uint8(obj.Class())) >= 0 {
			s.Objs.AddToUpdatable(obj)
		}
	}
	legacy.Nox_xxx_collide_548740()
}

func (s *Server) updateUnitsCCC() {
	for obj := s.Objs.UpdatableList; obj != nil; obj = obj.UpdatableNext {
		obj.PrevPos = obj.PosVec
		obj.PosVec = obj.NewPos
		obj.ForceVec = types.Pointf{}

		obj.Direction1 = obj.Direction2
		if obj.Field541 > 4 {
			obj.Field541 = 4
		}
		obj.SpeedCur = (obj.SpeedBonus + obj.SpeedBase) * (1 - 0.2*float32(obj.Field541))
		if obj.HasEnchant(server.ENCHANT_SLOWED) {
			obj.SpeedCur *= 0.5
		}
		if obj.Field541 != 0 || obj.Poison540 != 0 {
			if obj.Field542 > 0 {
				obj.Field542--
				if obj.Field542 == 0 {
					if obj.Field541 != 0 {
						obj.Field541--
					}
					if obj.Poison540 != 0 && !obj.Flags().Has(object.FlagDead) {
						legacy.Nox_xxx_updatePoison_4EE8F0(obj, 1)
					}
					obj.Field542 = 1000
				}
			}
		}
		legacy.Nox_xxx_updateUnitBuffs_4FF620(obj)
		if v31 := obj.Poison540; v31 != 0 {
			if h := obj.HealthData; h != nil && h.Max > 0 && h.Cur > 0 {
				dmg := 1
				if noxflags.HasGame(noxflags.GameModeQuest) {
					dmg += 1
				}
				if h.Field16 == 0 || (noxServer.Frame()-h.Field16) > 60 {
					if v31 > 8 || noxServer.Frame()%uint32(128>>(v31-1)) == 0 {
						obj.CallDamage(nil, nil, dmg, object.DamagePoison)
					}
				}
			}
		}
	}
}

func (s *Server) nox_xxx_updateServer_4D2DA0(a1 uint64) {
	if legacy.Get_dword_5d4594_528252() == 1 && s.Frame() == uint32(legacy.Get_dword_5d4594_528260()) {
		legacy.Nox_xxx_reconAttempt_41E390()
	}
	legacy.Sub_5096F0()
	if noxflags.HasGame(noxflags.GameFlag4) {
		return
	}
	if noxflags.HasGame(noxflags.GameOnline) {
		legacy.Sub_416720()
		if !noxflags.HasGame(noxflags.GameModeChat) && sub_409F40(0x2000) {
			s.Players.Camper.Update()
		}
	}
	if noxflags.HasGame(noxflags.GameModeElimination) {
		if legacy.Sub_40AA00() != 0 {
			if !s.flag1548704 {
				s.sub_4D2FF0()
			}
		} else {
			s.flag1548704 = false
		}
		if !noxflags.HasGame(noxflags.GameSuddenDeath) && !s.flag3592 {
			for pl := s.Players.First(); pl != nil; pl = s.Players.Next(pl) {
				if int(pl.Field2140) > 0 {
					if legacy.Sub_40AA00() == 0 {
						break
					}
					if noxflags.HasGamePlay(noxflags.GameplayFlag4) {
						if s.Teams.Count() >= legacy.Sub_40AA40() {
							break
						}
						for tm := s.Teams.First(); tm != nil; tm = s.Teams.Next(tm) {
							if legacy.Nox_xxx_countNonEliminatedPlayersInTeam_40A830(tm) == 1 {
								s.ServStartCountdown(int(s.Balance.Float("SuddenDeathCountdown")), "Settings.c:SuddenDeathImminent")
								break
							}
						}
					} else {
						if legacy.Sub_40A770() < legacy.Sub_40AA40() {
							s.ServStartCountdown(int(s.Balance.Float("SuddenDeathCountdown")), "Settings.c:SuddenDeathImminent")
						}
					}
					break
				}
			}
		}
	}
	if legacy.Sub_40A6B0() != 0 {
		sst := getServerSettings()
		s.NetPrintCompToAll(int(*(*uint32)(unsafe.Pointer(&sst.LatencyCompensationA66))))
		legacy.Sub_40A6A0(0)
	}
	if (a1 - *memmap.PtrUint64(0x5D4594, 1548692)) > 0x1F4 {
		legacy.Nox_xxx_netReportAllLatency_4D3050()
		*memmap.PtrUint64(0x5D4594, 1548692) = a1
	}
	if noxflags.HasGame(noxflags.GameModeChat) && legacy.Sub_40A740() == 0 && s.Teams.Count() != 0 && !noxflags.HasGamePlay(noxflags.GameplayFlag2) {
		legacy.Sub_4183C0()
	}
	if noxflags.HasGame(noxflags.GameModeQuest) {
		legacy.Sub_4D7150()
		legacy.Sub_4D71F0()
		legacy.Nox_server_checkWarpGate_4D7600()
		sub_4DCE00()
		legacy.Sub_4D7A80()
	}
}

func (s *Server) sub_4D2FF0() {
	if legacy.Sub_40AA70(nil) != 0 {
		s.flag1548704 = true
		return
	}
	for pl := s.Players.First(); pl != nil; pl = s.Players.Next(pl) {
		if pl.Field3680&1 != 0 {
			legacy.Nox_xxx_netNeedTimestampStatus_4174F0(pl, 256)
		}
	}
	s.flag1548704 = true
}

func (s *Server) ServStartCountdown(sec int, id strman.ID) { // nox_xxx_servStartCountdown_40A2A0
	*memmap.PtrUint64(0x5D4594, 3468) = platformTicks() + 1000*uint64(sec)
	legacy.Sub_40A1F0(1)
	if id != "" {
		s.NetPrintLineToAll(id)
	}
	s.flag3592 = true
}

func (s *Server) Sub40A040settings(a1 int, min int) {
	v2 := legacy.Sub_409A70(a1)
	if memmap.Uint8(0x5D4594, 3500+uintptr(v2)) == byte(min) {
		return
	}
	if !(!noxflags.HasGame(noxflags.GameSuddenDeath) && !s.flag3592) {
		str := s.Strings().GetStringInFile("NotInSuddenDeath", "settings.c")
		nox_xxx_printCentered_445490(str)
		return
	}
	legacy.Nox_server_gameSettingsUpdated_40A670()
	if nox_client_isConnected() {
		if min == 0 {
			str := s.Strings().GetStringInFile("parsecmd.c:timedisabled", "settings.c")
			nox_xxx_printCentered_445490(str)
		} else {
			str := s.Strings().GetStringInFile("parsecmd.c:timeset", "settings.c")
			nox_xxx_printCentered_445490(fmt.Sprintf(str, min))
			if legacy.Nox_xxx_gamePlayIsAnyPlayers_40A8A0() == 0 {
				str := s.Strings().GetStringInFile("TimeLimitDeferred", "settings.c")
				nox_xxx_printCentered_445490(str)
			}
		}
	}
	*memmap.PtrUint8(0x5D4594, 3500+uintptr(v2)) = byte(min)
	*memmap.PtrUint64(0x5D4594, 3468) = platformTicks() + 60000*uint64(min)
}

func (s *Server) nox_xxx_gameTick_4D2580_server_B(ticks uint64) bool {
	s.nox_xxx_updateServer_4D2DA0(ticks)
	s.Nox_server_netMaybeSendInitialPackets_4DEB30()
	s.nox_xxx_netlist_4DEB50()
	if !mainloopContinue {
		return false
	}
	s.MapSend.Update()
	s.NetXfer.Update(s.Frame())
	if !noxflags.HasGame(noxflags.GamePause) {
		s.updateUnits()
		legacy.Sub_4EC720()
		if noxflags.HasGame(noxflags.GameModeQuest) {
			legacy.Sub_50D890()
			legacy.Sub_4E4170()
		}
		nox_xxx_spellBookReact_4FCB70()
		s.abilities.Update()
		s.noxScript.ActRun()
		s.ScriptTick()
		legacy.Nox_xxx_voteUptate_506F30()
		s.deletedObjectsUpdate()
	}
	if err := s.updateRemotePlayers(); err != nil {
		gameLog.Println("update remote players:", err)
		return false
	}
	s.ObjectsAddPending()
	if inputKeyCheckTimeoutLegacy(0x10, s.SecToFrames(10)) {
		s.ProtectTypeCheck()
		inputSetKeyTimeoutLegacy(16)
	}
	if noxflags.HasGame(noxflags.GameOnline) && false && !noxflags.HasGame(noxflags.GameModeChat) && inputKeyCheckTimeoutLegacy(0xF, s.SecToFrames(60)) {
		legacy.Nox_xxx_net_4263C0()
		inputSetKeyTimeoutLegacy(15)
	}
	return true
}

func sub_446040() uint32 {
	return dword_5d4594_825768
}

func (s *Server) nox_xxx_gameTick_4D2580_server_E() {
	if noxflags.HasEngine(noxflags.EngineReplayWrite | noxflags.EngineReplayRead) {
		legacy.Sub_4E76C0()
	}
	if nox_xxx_gameGet_4DB1B0() {
		s.nox_xxx_gameTick_4D2580_server_D()
	}
	sub_4139C0()
	if nox_xxx_serverIsClosing446180() {
		sub_446190()
	}
	if sub_446030() && s.Frame() > s.SecToFrames(5)+sub_446040() {
		sub_446380()
	}
	if !noxflags.HasGame(noxflags.GamePause) {
		s.IncFrame()
	}
	legacy.Nox_xxx_protectData_56F5C0()
	s.maybeInitPlayerUnits()
	s.maybeRegisterGameOnline() // TODO: not exactly the right place
	s.maybeCallMapInit()
	s.maybeCallMapEntry()
	s.abilities.sub_4FC680()
	if playerInd := s.Players.ByInd(server.HostPlayerIndex); playerInd != nil {
		if unit := playerInd.PlayerUnit; unit != nil {
			s.spells.walls.associateSavedWalls(unit)
		}
	}
	if legacy.Nox_xxx_get_57AF20() != 0 && legacy.Sub_57B140() {
		legacy.Sub_57B0A0()
	}
	if s.nox_xxx_isQuest_4D6F50() {
		s.switchQuestIfRequested4D6FD0()
		sub_4DCF20()
	}
}

func (s *Server) maybeCallMapInit() {
	if s.ShouldCallMapInit && s.Players.HasUnits() {
		s.scriptOnEvent(script.EventMapInitialize)
		s.ShouldCallMapInit = false
	}
}

func (s *Server) maybeCallMapEntry() {
	if s.ShouldCallMapEntry && s.Players.HasUnits() {
		s.scriptOnEvent(script.EventMapEntry)
		s.ShouldCallMapEntry = false
	}
}

func getServerSettings() *server.Settings { // sub_416640
	return memmap.PtrT[server.Settings](0x5D4594, 371516)
}

func gameIsClosed() bool { // sub_416A00
	return getServerSettings().Flags100.Has(server.ServerClosed)
}

func serverSetClosed() { // sub_4169E0
	getServerSettings().Flags100 |= server.ServerClosed
}

func serverSetOpen() { // sub_4169F0
	getServerSettings().Flags100 &^= server.ServerClosed
}

func (s *Server) updateRemotePlayers() error {
	for _, pl := range s.Players.List() {
		if pl.PlayerUnit == nil {
			continue
		}
		sec := 30
		if pl.Field3680&0x10 != 0 {
			sec = 90
		}
		if s.Frame()-uint32(pl.Frame3596) > s.SecToFrames(sec) {
			m := uint32(pl.NetCodeVal)
			// TODO: passing Go pointer
			legacy.Nox_xxx_netInformTextMsg2_4DA180(3, unsafe.Pointer(&m))
			var buf [1]byte
			buf[0] = byte(noxnet.MSG_TIMEOUT_NOTIFICATION)
			s.NetStr.ByPlayer(pl).SendReliable(buf[:])
			s.PlayerDisconnect(pl, 3)
		}
		if pl.Field3680&0x80 != 0 {
			s.PlayerDisconnect(pl, 4)
		}
		if (pl.Field3676 != 3) || (pl.Field3680&0x10 == 0) {
			buf, err := noxnet.AppendPacket(nil, &noxnet.MsgTimestamp{
				T: uint16(s.Frame()),
			})
			if err != nil {
				panic(err)
			}
			s.NetList.AddToMsgListCli(pl.PlayerIndex(), netlist.Kind1, buf)
		} else {
			if uint32(pl.PlayerUnit.Ind()) == legacy.DeadWord { // see #401
				pl.PlayerUnit = nil
				mainloopStop()
				return fmt.Errorf("player unit deleted (%s)", pl.Name())
			}
			s.nox_xxx_netUpdate_518EE0(pl.PlayerUnit)
		}
		if pl.PlayerUnit == s.Players.HostUnit() {
			legacy.Nox_xxx_netImportant_4E5770(byte(pl.Index()), 1)
		} else if legacy.Get_dword_5d4594_2650652() == 0 || (s.Frame()%uint32(nox_xxx_rateGet_40A6C0()) == 0) || noxflags.HasGame(noxflags.GameFlag4) {
			s.NetStr.ByPlayer(pl).FlushAndPoll()
		}
	}
	return nil
}

func (s *Server) nox_xxx_secretWallCheckUnits_517F00(rect types.Rectf, fnc func(it unsafe.Pointer)) {
	for it := nox_xxx_wallSecretGetFirstWall_410780(); it != nil; it = nox_xxx_wallSecretNext_410790(it) {
		x := float64(*(*int32)(unsafe.Add(it, 4)) * common.GridStep)
		y := float64(*(*int32)(unsafe.Add(it, 8)) * common.GridStep)
		if x > float64(rect.Min.X) && x < float64(rect.Max.X) &&
			y > float64(rect.Min.Y) && y < float64(rect.Max.Y) {
			fnc(it)
		}
	}
}

func (s *Server) nox_xxx_netUpdate_518EE0(u *server.Object) {
	ud := u.UpdateDataPlayer()
	pl := ud.Player
	pind := pl.PlayerIndex()
	s.NetList.InitByInd(pind)
	if pind != server.HostPlayerIndex && ((s.Frame()+uint32(pind))%s.SecToFrames(15)) == 0 {
		legacy.Nox_xxx_netReportUnitHeight_4D9020(pind, u)
	}
	if legacy.Get_dword_5d4594_2650652() == 0 || (s.Frame()%uint32(nox_xxx_rateGet_40A6C0())) == 0 || noxflags.HasGame(noxflags.GameFlag4) {
		if pl.Field3680&0x40 != 0 {
			buf, err := noxnet.AppendPacket(nil, &noxnet.MsgFullTimestamp{
				T: s.Frame(),
			})
			if err != nil {
				panic(err)
			}
			nox_netlist_addToMsgListSrv(pind, buf)
			legacy.Nox_xxx_playerUnsetStatus_417530(pl, 64)
		} else {
			buf, err := noxnet.AppendPacket(nil, &noxnet.MsgTimestamp{
				T: uint16(s.Frame()),
			})
			if err != nil {
				panic(err)
			}
			nox_netlist_addToMsgListSrv(pind, buf)
		}
	}
	if legacy.Get_dword_5d4594_2650652() == 0 || u == s.Players.HostUnit() || noxflags.HasGame(noxflags.GameFlag4) || (s.Frame()%uint32(nox_xxx_rateGet_40A6C0())) == 0 {
		if pl.Field3680&3 != 0 || noxflags.HasEngine(noxflags.EngineReplayRead) {
			if !s.nox_xxx_netPlayerObjSendCamera_519330(u) {
				return
			}
			if noxflags.HasEngine(noxflags.EngineReplayRead) {
				legacy.Nox_xxx_netPlayerObjSend_518C30(u, u, 1, 1)
			}
		} else {
			if legacy.Nox_xxx_netPlayerObjSend_518C30(u, u, 1, 1) == 0 {
				return
			}
		}
		dp := types.Ptf(float32(pl.Field10), float32(pl.Field12)).Add(types.Ptf(100, 100))
		p1 := pl.Pos3632().Sub(dp)
		p2 := pl.Pos3632().Add(dp)
		rect := types.RectFromPointsf(p1, p2)
		s.Map.EachObjAndMissileInRect(rect, func(it *server.Object) bool {
			s.nox_xxx_unitAroundPlayerFn_5193B0(it, u)
			return true
		})

		dp = types.Ptf(float32(pl.Field10), float32(pl.Field12)).Add(types.Ptf(128, 128))
		p1 = pl.Pos3632().Sub(dp)
		p2 = pl.Pos3632().Add(dp)
		rect = types.RectFromPointsf(p1, p2)
		s.nox_xxx_secretWallCheckUnits_517F00(rect, func(it unsafe.Pointer) {
			s.sub_519660(it, u)
		})
		if legacy.Sub_519710(unsafe.Pointer(ud)) != 0 {
			s.sub_519760(u, rect)
		}
		if s.Frame()&8 != 0 {
			plBit := uint32(1 << pl.Index())
			for it := s.Objs.First(); it != nil; it = it.Next() {
				if !it.Class().HasAny(object.ClassClientPersist|object.ClassImmobile) && legacy.Nox_xxx_playerMapTracksObj_4173D0(pl.Index(), it) == 0 && (float64(it.CollideP1.X) > float64(rect.Max.X) || float64(it.CollideP2.X) < float64(rect.Min.X) || float64(it.CollideP1.Y) > float64(rect.Max.Y) || float64(it.CollideP2.Y) < float64(rect.Min.Y)) {
					if it.Field37&plBit != 0 {
						s.Nox_xxx_netObjectOutOfSight_528A60(pl.Index(), it)
						it.Field38 |= plBit
						it.Field37 &^= plBit
					}
				}
			}
			for it := s.Objs.MissileList; it != nil; it = it.Next() {
				if !it.Class().HasAny(object.ClassClientPersist|object.ClassImmobile) && legacy.Nox_xxx_playerMapTracksObj_4173D0(pl.Index(), it) == 0 && (float64(it.CollideP1.X) > float64(rect.Max.X) || float64(it.CollideP2.X) < float64(rect.Min.X) || float64(it.CollideP1.Y) > float64(rect.Max.Y) || float64(it.CollideP2.Y) < float64(rect.Min.Y)) {
					if it.Field37&plBit != 0 {
						s.Nox_xxx_netObjectOutOfSight_528A60(pl.Index(), it)
						it.Field38 |= plBit
						it.Field37 &^= plBit
					}
				}
			}
		}
	}
	if legacy.Get_dword_5d4594_2650652() == 0 || (s.Frame()%uint32(nox_xxx_rateGet_40A6C0())) == 0 || noxflags.HasGame(noxflags.GameFlag4) {
		s.spells.walls.changeOrAddRemoteWalls(pl)
		legacy.Sub_511100(pl.Index())
	}
	legacy.Nox_xxx_netUpdateRemotePlr_501CA0(u)
}

func (s *Server) sub_519760(u *server.Object, rect types.Rectf) {
	ud := u.UpdateDataPlayer()
	pl := ud.Player
	pind := pl.PlayerIndex()
	obj := s.Players.Sub_4172C0(pind)
	if obj == nil {
		return
	}
	if obj.Flags().Has(object.FlagDestroyed) {
		s.Players.Nox_xxx_netMinimapUnmark4All_417430(obj)
	} else if float64(obj.PosVec.X) < float64(rect.Min.X) || float64(obj.PosVec.X) > float64(rect.Max.X) || float64(obj.PosVec.Y) < float64(rect.Min.Y) || float64(obj.PosVec.Y) > float64(rect.Max.Y) {
		obj.Field38 |= uint32(1 << pind)
		legacy.Nox_xxx_netSendObjects2Plr_519410(u, obj)
		legacy.Nox_xxx_netReportUnitHeight_4D9020(pind, obj)
		ud.Field67 = s.Frame()
	}
}

func (s *Server) sub_519660(it unsafe.Pointer, u *server.Object) {
	pl := u.ControllingPlayer()
	v2 := uint32(1 << pl.Index())
	isSet := (v2 & *(*uint32)(unsafe.Add(it, 28))) != 0
	var exp bool
	switch *(*uint8)(unsafe.Add(it, 21)) {
	case 1, 4:
		exp = false
	case 2, 3:
		exp = true
	default:
		exp = u != nil
	}
	if isSet != exp {
		wl := s.Walls.GetWallAtGrid(image.Pt(int(*(*uint32)(unsafe.Add(it, 4))), int(*(*uint32)(unsafe.Add(it, 8)))))
		if exp {
			legacy.Sub_4DF120(wl.C())
			*(*uint32)(unsafe.Add(it, 28)) |= v2
		} else {
			legacy.Sub_4DF180(wl.C())
			*(*uint32)(unsafe.Add(it, 28)) &^= v2
		}
	}
}

func (s *Server) nox_xxx_unitAroundPlayerFn_5193B0(it, u *server.Object) {
	ud := u.UpdateDataPlayer()
	pl := ud.Player
	if u == it {
		legacy.Nox_xxx_netUpdateObjectSpecial_527E50(u, it)
		if pl.Field3680&0x1 == 0 {
			return
		}
	}
	if !noxflags.HasGame(noxflags.GameOnline) || ud.Field68 != s.Frame() {
		legacy.Nox_xxx_netSendObjects2Plr_519410(u, it)
	}
}

func (s *Server) newSession() error {
	gameLog.Println("new server session")
	legacy.Sub_4D15C0()
	legacy.Set_dword_5d4594_2649712(0x80000000)
	s.Players.SetHost(nil, nil)
	legacy.Sub_4D7B40()
	legacy.Sub_41E4B0(0)
	s.Objs.ResetObjectScriptIDs()
	legacy.Sub_56F1C0()
	s.Players.ResetAll()
	s.NetList.ResetAll()
	legacy.Sub_4E4EF0()
	legacy.Sub_4E4ED0()
	s.Audio.Init(s.Server)
	s.Audio.OnSound(func(id sound.ID, kind int, obj *server.Object, pos types.Pointf) {
		if kind != 2 {
			s.ai.NewSound(id, obj, pos)
		}
	})
	if err := s.nox_read_things_alternative_4E2B60(); err != nil {
		return err
	}
	legacy.Nox_motd_4463E0(1)
	s.TeamsReset()
	legacy.Sub_4259C0()
	s.Players.Camper.Reset()
	if legacy.Sub_518770() == 0 {
		return errors.New("sub_518770 failed")
	}
	noxflags.HasGame(noxflags.GameFlag22)
	if !s.Objs.Init(5000) {
		return errors.New("nox_xxx_allocClassArrayObjects_4E3360 failed")
	}
	s.Map.Init()
	s.AI.Paths.Init(s.Server)
	s.Spells.Init()
	s.spells.Init(s)
	s.Abils.Reset()
	if err := nox_xxx_allocSpellRelatedArrays_4FC9B0(); err != nil {
		return err
	}
	s.MapGroups.Init()
	if legacy.Nox_xxx_allocItemRespawnArray_4ECA60() == 0 {
		return errors.New("nox_xxx_allocItemRespawnArray_4ECA60 failed")
	}
	if legacy.Nox_xxx_registerShopClasses_50E2A0() == 0 {
		return errors.New("nox_xxx_registerShopClasses_50E2A0 failed")
	}
	if legacy.Nox_xxx_allocMonsterRelatedArrays_50D780() == 0 {
		return errors.New("nox_xxx_allocMonsterRelatedArrays_50D780 failed")
	}
	if legacy.Nox_xxx_allocVoteArray_5066D0() == 0 {
		return errors.New("nox_xxx_allocVoteArray_5066D0 failed")
	}
	if legacy.Nox_xxx_monsterList_517520() == 0 {
		return errors.New("nox_xxx_monsterList_517520 failed")
	}
	legacy.Sub_416920()
	if !noxflags.HasGame(noxflags.GameModeCoop) {
		if err := s.listen(context.Background(), s.ServerPort()); err != nil {
			return err
		}
	}
	if legacy.Nox_xxx_allocPendingOwnsArray_516EE0() == 0 {
		return errors.New("nox_xxx_allocPendingOwnsArray_516EE0 failed")
	}
	legacy.Sub_421B10()
	sub_4DB0A0()
	legacy.Sub_4D0F30()
	return s.StartServices(isDedicatedServer)
}

func (s *Server) nox_xxx_servEndSession_4D3200() {
	s.ai.Free()
	sub_4DB100()
	legacy.Sub_421B10()
	legacy.Sub_516F10()
	s.nox_xxx_replayStopSave_4D33B0()
	s.nox_xxx_replayStopReadMB_4D3530()
	s.Players.ResetAll()
	legacy.Sub_446490(1)
	legacy.Sub_4259F0()
	s.nox_xxx_mapSwitchLevel_4D12E0(false)
	s.nox_xxx_mapLoad_40A380()
	legacy.Sub_4E4DE0()
	s.Map.Debug.Reset()
	s.MapGroups.Free()
	s.springs.Reset()
	s.abilities.Free()
	s.spells.Free()
	s.Spells.Free()
	nox_xxx_freeSpellRelated_4FCA80()
	s.AI.Paths.Free()
	s.Map.Free()
	s.Audio.Free()
	legacy.Sub_4ECA90()
	legacy.Sub_506720()
	legacy.Sub_50D820()
	legacy.Nox_xxx_deleteShopInventories_50E300()
	legacy.Sub_416950()
	s.Objs.FreeObjects()
	s.FreeObjectTypes()
	nox_xxx_free_42BF80()
	if !noxflags.HasGame(noxflags.GameModeCoop) {
		s.Nox_server_netCloseHandler_4DEC60()
	}
	legacy.Sub_56F3B0()
	s.NetList.ResetAll()
	_ = ifs.Remove(datapath.Save("_temp_.dat"))
}

func sub_4D3C30() {
	noxServer.Nox_xxx_free503F40()
	legacy.Sub_51D0E0()
	legacy.Sub_502DF0()
}

func (s *Server) nox_server_loadMapFile_4CF5F0(mname string, noCrypt bool) error {
	gameLog.Printf("loading map %q", mname)
	cntGameMap.WithLabelValues(filepath.Base(mname)).Inc()
	legacy.Sub_481410()
	s.ObjectsAddPending()
	s.Nox_xxx_waypoint_5799C0()
	if mname == "" {
		return errors.New("empty map name")
	}
	if strings.ToLower(mname) == "#return" {
		mname = alloc.GoString((*byte)(memmap.PtrOff(0x5D4594, 1523080)))
	} else if strings.HasPrefix(mname, "#") {
		v3 := datapath.Data()
		legacy.Sub_4D39F0(v3)
		v13 := mname[1:]
		if i := strings.IndexByte(mname, '.'); i > 0 {
			v13 = v3[:i]
		}
		legacy.Sub_4D42E0(v13)
		v12 := fmt.Sprintf("$%s.map", v13)
		s.nox_xxx_gameSetMapPath_409D70(v12)
		if legacy.Nox_xxx_mapGenStart_4D4320() == 0 {
			s.nox_xxx_mapSwitchLevel_4D12E0(true)
			return errors.New("nox_xxx_mapGenStart_4D4320 failed")
		}
		sub_4D3C30()
		mname = v12
	}
	var fname string
	if strings.ContainsAny(mname, "\\/") {
		fname = mname
	} else {
		dir := strings.TrimSuffix(mname, filepath.Ext(mname))
		fname = filepath.Join("maps", dir, mname)
	}
	if _, err := ifs.Stat(fname); err != nil {
		zname := strings.TrimSuffix(fname, filepath.Ext(mname)) + ".nxz"
		if _, err := ifs.Stat(zname); err != nil {
			return err
		}
		if err := cnxz.DecompressFile(zname, fname); err != nil {
			return fmt.Errorf("cannot decompress map %q: %w", zname, err)
		}
	}
	v8 := s.getServerMap()
	nox_common_checkMapFile(v8)
	ckey := crypt.MapKey
	if noCrypt {
		ckey = -1
	}
	err := cryptfile.OpenGlobal(fname, cryptfile.ReadOnly, ckey)
	if err != nil {
		return err
	}
	cf := cryptfile.Global()
	crc, err := mapReadCryptHeader(cf)
	if err != nil {
		cryptfile.Close()
		return err
	}
	nox_xxx_mapSetCrcMB_409B10(crc)
	// Script VM init must be done before map sections are read.
	// This is done because object section will need to bind to script function names.
	// Doing VMs init after map load is too late in this case.
	s.vmsInitMap()
	if err := s.nox_xxx_serverParseEntireMap_4CFCE0(cf); err != nil {
		cryptfile.Close()
		gameLog.Println("server read map:", err)
		return err
	}
	s.noxScript.nox_xxx_scriptRunFirst_507290()
	cryptfile.Close()
	if !noxflags.HasGame(noxflags.GameFlag22) {
		s.nox_xxx_mapReadSetFlags_4CF990()
		if false {
			legacy.Sub_416690()
		}
		noxflags.UnsetGame(noxflags.GameSuddenDeath)
		legacy.Sub_470680()
		legacy.Sub_4D0550(fname)
		legacy.Sub_4161E0()
		if !noxflags.HasGame(noxflags.GameModeChat) {
			legacy.Sub_4165F0(0, 1)
		}
	}
	alloc.StrCopy(unsafe.Slice((*byte)(memmap.PtrOff(0x5D4594, 1523080)), 1024), mname)
	return nil
}

func (s *Server) maybeInitPlayerUnits() {
	if !s.ShouldCallMapInit && !s.ShouldCallMapEntry {
		return
	}
	if len(s.Players.ListUnits()) == 0 {
		return
	}
	if noxflags.HasGame(noxflags.GameModeQuest) {
		if s.nox_game_getQuestStage_4E3CC0() == 1 {
			legacy.Nox_game_sendQuestStage_4D6960(255)
			legacy.Sub_4D7440(1)
			legacy.Sub_4D60B0()
		} else if !sub4D6F30() || legacy.Sub_4D7430() != 0 {
			if legacy.Sub_4D76F0() == 1 {
				legacy.Sub_4D6880(255, 1)
				legacy.Sub_4D76E0(0)
				legacy.Sub_4D7440(1)
				legacy.Sub_4D60B0()
			} else {
				fname := datapath.Save("_temp_.dat")
				for _, u := range s.Players.ListUnits() {
					ud := u.UpdateDataPlayer()
					plx := ud.Player
					pi := plx.PlayerIndex()
					if plx.Field4792 == 1 && ud.Field138 == 0 && savePlayerData(fname, pi) {
						v5 := sub_419EE0(pi)
						s.Nox_xxx_sendGauntlet_4DCF80(pi, 1)
						if !sub41CFA0(fname, pi) && !v5 {
							s.Nox_xxx_sendGauntlet_4DCF80(pi, 0)
						}
						ifs.Remove(fname)
					}
					legacy.Sub_4D6770(pi)
				}
				legacy.Sub_4D6880(255, 0)
				legacy.Sub_4D7440(1)
				legacy.Sub_4D60B0()
			}
		} else {
			legacy.Nox_game_sendQuestStage_4D6960(255)
			legacy.Sub_4D7440(1)
			legacy.Sub_4D60B0()
		}
	} else {
		s.Nox_xxx_netMsgFadeBegin_4D9800(true, true)
	}
	if noxflags.HasGame(noxflags.GameOnline) && !noxflags.HasGame(noxflags.GameModeChat) {
		for _, u := range s.Players.ListUnits() {
			plx := u.ControllingPlayer()
			if plx.PlayerIndex() != server.HostPlayerIndex && plx.Field3680&1 == 0 {
				asObjectS(u).ApplyEnchant(server.ENCHANT_INVULNERABLE, 0, 5)
			}
		}
	}
}

func (s *Server) SwitchMap(fname string) {
	gameLog.Printf("switch map: %q", fname)
	ptr2408 := unsafe.Slice((*byte)(memmap.PtrOff(0x973F18, 2408)), 1464)

	var v5 [1464]byte
	copy(v5[:], ptr2408)

	name := fname
	if ext := filepath.Ext(fname); ext != "" {
		name = strings.TrimSuffix(name, ext)
	}
	name = strings.ToLower(name)
	nox_common_checkMapFile(name)
	legacy.Sub_4CFDF0(int(memmap.Uint32(0x973F18, 3800)))
	copy(ptr2408, v5[:])
	legacy.Set_dword_5d4594_1548524(1)
	mname := fname
	if i := strings.IndexByte(fname, ':'); i >= 0 {
		s.mapSwitchWPName = mname[i+1:]
		mname = mname[:i]
	} else {
		s.mapSwitchWPName = ""
	}
	if s.MapSend.Active() != 0 {
		s.MapSend.AbortAll(0)
	}
	mname = strings.ToLower(mname)
	s.nox_xxx_gameSetMapPath_409D70(mname)
}

func nox_gameModeFromMapPtr(a1 unsafe.Pointer) noxflags.GameFlag {
	v := *(*uint32)(unsafe.Pointer(uintptr(a1) + 1392))
	if v&0x4 != 0 {
		return noxflags.GameModeArena
	}
	if v&0x20 != 0 {
		return noxflags.GameModeElimination
	}
	if v&0x8 != 0 {
		return noxflags.GameModeCTF
	}
	if v&0x10 != 0 {
		return noxflags.GameModeKOTR
	}
	if v&0x40 != 0 {
		return noxflags.GameModeFlagBall
	}
	if v&0x2 != 0 {
		return noxflags.GameModeQuest
	}
	if v&0x1 != 0 {
		return noxflags.GameModeCoopTeam
	}
	return noxflags.GameModeChat
}

func nox_mapToGameFlags(v int) noxflags.GameFlag {
	var out noxflags.GameFlag
	if v&1 != 0 {
		out |= noxflags.GameModeCoopTeam
	}
	if v&2 != 0 {
		out |= noxflags.GameModeQuest
	}
	if v&4 != 0 {
		out |= noxflags.GameModeArena
	}
	if v&0x20 != 0 {
		out |= noxflags.GameModeElimination
	}
	if v&8 != 0 {
		out |= noxflags.GameModeCTF
	}
	if v&0x10 != 0 {
		out |= noxflags.GameModeKOTR
	}
	if v&0x40 != 0 {
		out |= noxflags.GameModeFlagBall
	}
	if v < 0 {
		out |= noxflags.GameModeChat
	}
	return out
}

func nox_xxx_mapGetTypeMB_4CFFA0(a1 unsafe.Pointer) noxflags.GameFlag {
	val := *(*int32)(unsafe.Pointer(uintptr(a1) + 1392))
	return nox_mapToGameFlags(int(val))
}

func (s *Server) nox_xxx_mapReadSetFlags_4CF990() {
	stt := getCurrentSettings2()
	if noxflags.HasGame(noxflags.GameModeElimination) && (memmap.Int32(0x973F18, 3800) < 0 || (stt.Field52&0x400 == 0)) {
		legacy.Nox_xxx_ruleSetNoRespawn_40A5E0(0)
	}
	legacy.Sub_455C10()
	legacy.Sub_456050()
	if noxflags.HasGame(noxflags.GameModeQuest) && memmap.Int32(0x973F18, 3800) < 0 {
		s.sub_4D6B10(true)
		s.TeamsRemoveActive(true)
	}
	mapname := s.getServerMap()
	gameLog.Printf("checking map flags for %q", filepath.Base(mapname))
	if err := nox_common_checkMapFile(mapname); err != nil {
		gameLog.Println("check map file:", err)
		if !noxflags.HasGame(noxflags.GameModeCoopTeam) {
			noxflags.UnsetGame(noxflags.GameModeMask)
			noxflags.SetGame(noxflags.GameModeArena)
			legacy.Sub_4D0D90(1)
		}
		return
	}
	mapType := nox_xxx_mapGetTypeMB_4CFFA0(memmap.PtrOff(0x973F18, 2408))
	vv := memmap.Int32(0x973F18, 3800)
	if vv&1 != 0 {
		gameLog.Println("setting coop mode")
		noxflags.UnsetGame(noxflags.GameModeMask)
		s.createCoopTeam()
		noxflags.SetGame(noxflags.GameModeCoop)
	} else if vv&2 != 0 {
		gameLog.Println("setting quest mode")
		isChat := noxflags.HasGame(noxflags.GameModeChat)
		noxflags.UnsetGame(noxflags.GameModeMask)
		noxflags.SetGame(noxflags.GameModeQuest)
		if isChat {
			s.setupQuestGame()
		}
	} else if vv >= 0 {
		if noxflags.GameFlag(stt.Field52)&mapType == 0 {
			stt.Field52 = uint16(nox_gameModeFromMapPtr(memmap.PtrOff(0x973F18, 2408)) | noxflags.GameFlag(stt.Field52)&0xE80F)
		}
		if stt.Field52&0x10 == 0 {
			legacy.Nox_xxx_mapFindCrown_4CFC30()
		}
		mode := noxflags.GameFlag(stt.Field52)
		if mode.Has(noxflags.GameModeCTF) {
			gameLog.Println("setting CTF mode")
			if legacy.Nox_xxx_mapInfoSetCapflag_417EA0() != 0 {
				noxflags.UnsetGame(noxflags.GameModeMask)
				noxflags.SetGame(noxflags.GameModeCTF)
			}
		} else if mode.Has(noxflags.GameModeFlagBall) {
			gameLog.Println("setting flagball mode")
			if legacy.Nox_xxx_mapInfoSetFlagball_417F30() != 0 {
				noxflags.UnsetGame(noxflags.GameModeMask)
				noxflags.SetGame(noxflags.GameModeFlagBall)
				s.Spells.Enable(spell.SPELL_WALL, false)
			}
		} else if mode.Has(noxflags.GameModeKOTR) {
			gameLog.Println("setting KOTR mode")
			if legacy.Nox_xxx_mapInfoSetKotr_4180D0() != 0 {
				noxflags.UnsetGame(noxflags.GameModeMask)
				noxflags.SetGame(noxflags.GameModeKOTR)
			}
		} else if mode.Has(noxflags.GameModeElimination) {
			gameLog.Println("setting elimination mode")
			if !noxflags.HasGame(noxflags.GameModeElimination) {
				legacy.Nox_xxx_ruleSetNoRespawn_40A5E0(1)
			}
			noxflags.UnsetGame(noxflags.GameModeMask)
			noxflags.SetGame(noxflags.GameModeElimination)
		} else {
			gameLog.Println("setting arena mode")
			noxflags.UnsetGame(noxflags.GameModeMask)
			noxflags.SetGame(noxflags.GameModeArena)
		}
	} else {
		gameLog.Println("setting chat mode")
		legacy.Sub_40A1F0(0)
		noxflags.UnsetGame(noxflags.GameModeMask)
		noxflags.SetGame(noxflags.GameModeChat)
		if s.Teams.Count() != 0 {
			legacy.Nox_xxx_teamAssignFlags_418640()
			if !noxflags.HasGamePlay(noxflags.GameplayFlag2) && !noxflags.HasGame(noxflags.GameFlag16) {
				legacy.Nox_xxx_toggleAllTeamFlags_418690(1)
			}
		}
	}
}

func (s *Server) nox_xxx_moveUpdateSpecial_517970(u *server.Object) {
	s.sub_517870(u)
	if s.Map.ValidIndexPos(u.NewPos) {
		s.Map.AddObjectToIndex(u)
	} else {
		if u.Class().Has(object.ClassPlayer) {
			gameLog.Printf("attempting to delete player unit; stopping the map")
			mainloopStop()
			return
		}
		asObjectS(u).Delete()
	}
}

func (s *Server) sub_517870(obj *server.Object) {
	if !obj.Flags().Has(object.FlagPartitioned) {
		return
	}
	s.Map.Sub5178E0(false, &obj.ObjIndexBase)
	if !obj.Class().Has(object.ClassMissile) {
		for i := range obj.ObjIndex[:obj.ObjIndexCur] {
			s.Map.Sub5178E0(true, &obj.ObjIndex[i])
		}
		obj.ObjIndexCur = 0
	}
	obj.ObjFlags &^= object.FlagPartitioned
}

func sub_4DB0A0() {
	questPlayerFile = ""
	dword_5d4594_1563044 = false
	dword_5d4594_1563048 = false
	dword_5d4594_1563080 = false
	dword_5d4594_1563084 = nil
	dword_5d4594_1563088 = 0
	dword_5d4594_1563092 = 0
	legacy.Set_dword_5d4594_1563096(0)
	*memmap.PtrUint32(0x5D4594, 1563100) = 0
	noxServer.quest.mapName = ""
	dword_5d4594_1563064 = false
	questPlayerSet = false
}

func sub_469B90(a1 [3]uint32) {
	*memmap.PtrUint32(0x587000, 142296) = a1[0]
	*memmap.PtrUint32(0x587000, 142300) = a1[1]
	*memmap.PtrUint32(0x587000, 142304) = a1[2]
}

func (s *Server) nox_xxx_mapSwitchLevel_4D12E0(a1 bool) {
	noxflags.SetGame(noxflags.GameFlag20)
	defer noxflags.UnsetGame(noxflags.GameFlag20)

	legacy.Sub_516F30()
	legacy.Sub_421B10()
	acl := [3]uint32{
		25, 25, 25,
	}
	sub_469B90(acl)
	if noxflags.HasGame(noxflags.GameClient) {
		sub_4349C0(acl)
	}
	s.scriptsReset()
	if noxflags.HasGame(noxflags.GameModeCoop) {
		legacy.Sub_4FCEB0(a1)
	} else {
		legacy.Sub_4FCEB0(false)
	}
	s.spells.walls.Reset()
	for _, pu := range s.Players.ListUnits() {
		ud := pu.UpdateDataPlayer()
		legacy.Sub_4F7950(pu)
		ud.Field74 = 0
		asObjectS(pu).Freeze(false)
		ud.Field40_0 = 0
		if ud.Trade70 != nil {
			legacy.Nox_xxx_shopCancelSession_510DC0(ud.Trade70)
		}
		ud.Trade70 = nil
		if pu.Update == legacy.Get_nox_xxx_updatePlayerMonsterBot_4FAB20() {
			legacy.Nox_xxx_playerBotCreate_4FA700(pu)
		}
	}
	for {
		s.ObjectsAddPending()
		legacy.Sub_4E5BF0(a1)
		s.spells.duration.spellCastByPlayer()
		s.FinalizeDeletingObjects()
		if s.Objs.Pending == nil {
			break
		}
	}
	for obj := s.Objs.First(); obj != nil; obj = obj.Next() {
		obj.Obj130 = nil
		if legacy.Nox_xxx_isUnit_4E5B50(obj) != 0 && obj.Class().Has(object.ClassMonster) {
			ud := obj.UpdateDataMonster()
			ud.ScriptEnemySighted.Func = -1
			ud.ScriptLookingForEnemy.Func = -1
			ud.ScriptDeath.Func = -1
			ud.ScriptChangeFocus.Func = -1
			ud.ScriptIsHit.Func = -1
			ud.ScriptRetreat.Func = -1
			ud.ScriptCollision.Func = -1
			ud.ScriptHearEnemy.Func = -1
			ud.ScriptEndOfWaypoint.Func = -1
			ud.ScriptLostEnemy.Func = -1
			ud.Field98 = 0
			ud.Field101 = 0
		}
	}
	s.ai.Reset()
	for obj := s.Objs.MissileList; obj != nil; obj = obj.Next() {
		if legacy.Sub_4E5B80(obj) != 0 {
			legacy.Sub_4E81D0(obj)
		}
	}
	legacy.Sub_4ECFE0()
	legacy.Sub_511E20()
	s.Walls.Reset()
	if a1 {
		legacy.Nox_xxx_Fn_4FCAC0(a1, 1)
	} else {
		legacy.Nox_xxx_Fn_4FCAC0(false, 0)
	}
	legacy.Nox_xxx_mapSwitchLevel_4D12E0_tileFree()
	legacy.Sub_410730()
	s.Walls.ClearBreakable()
	s.Nox_xxx_waypointDeleteAll_579DD0()
	legacy.Nox_xxx_j_allocHitArray_511840()
	legacy.Nox_xxx_decayDestroy_5117B0()
	s.springs.Reset()
	s.Map.Debug.Reset()
	s.MapGroups.Reset()
	legacy.Sub_510E50()
	legacy.Sub_4D1610()
	legacy.Sub_4EC5B0()
	legacy.Sub_50E360()
	legacy.Sub_50D7E0()
	legacy.Sub_4E4F80()
}
