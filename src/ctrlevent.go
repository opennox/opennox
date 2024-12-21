package opennox

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"image"
	"math"
	"strings"
	"unsafe"

	"github.com/noxworld-dev/opennox-lib/cfg"
	"github.com/noxworld-dev/opennox-lib/client/keybind"
	"github.com/noxworld-dev/opennox-lib/common"
	"github.com/noxworld-dev/opennox-lib/player"

	"github.com/noxworld-dev/opennox/v1/client/gui"
	"github.com/noxworld-dev/opennox/v1/client/input"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"github.com/noxworld-dev/opennox/v1/common/sound"
	"github.com/noxworld-dev/opennox/v1/legacy"
	"github.com/noxworld-dev/opennox/v1/server"
)

var (
	allowEmotionsInSolo = true
	wasManualSpellInvertKeyPressed = false
)

func init() {
	configBoolPtr("game.extensions.solo_allow_emotes", "", true, &allowEmotionsInSolo)
}

const ctrlEventCap = 128

type noxCtrlEvent struct {
	tick uint32
	server.PlayerCtrl
}

type CtrlEventBinding struct {
	keys    []keybind.Key
	events  []keybind.Event
	prev    *CtrlEventBinding
	next    *CtrlEventBinding
	field20 *CtrlEventBinding
	field21 *CtrlEventBinding
	flag22  bool
	frame   uint32
}

func (c *CtrlEventBinding) defKey() keybind.Key {
	if len(c.keys) == 0 {
		return 0
	}
	return c.keys[0]
}

func (c *CtrlEventBinding) defEvent() keybind.Event {
	if len(c.events) == 0 {
		return 0
	}
	return c.events[0]
}

type CtrlEventHandler struct {
	ticks       uint32
	playerDir   uint32
	bufA        [ctrlEventCap]noxCtrlEvent
	bufB        [ctrlEventCap]noxCtrlEvent // TODO: size a guess
	bindings    *CtrlEventBinding
	indA        int
	indB        int
	indC        int
	indD        int
	flags747848 byte
	flags750956 bool
	flags754064 uint32
	netBuf      []byte
}

func (c *CtrlEventHandler) Reset() {
	c.bindings = nil
	c.flags750956 = false
	c.indA = 0
	c.indB = 0
	c.indC = 0
	c.indD = 0
}

func (c *CtrlEventHandler) nox_xxx_clientControl_42D6B0(mpos image.Point, a4 *CtrlEventBinding) {
	c.ticks = uint32(platformTicks())
	if noxflags.HasGame(noxflags.GameHost) && noxflags.HasGame(noxflags.GameOnline) {
		c.ticks += nox_ctrlevent_add_ticks_42E630()
	}
	if !noxflags.HasGame(noxflags.GameHost) {
		c.indA = 0
	}
	c.nox_xxx_clientControl_42D6B0_A(a4)
	c.nox_xxx_clientControl_42D6B0_orientation(mpos)
	if getPlayerClass() != player.Warrior {
		legacy.Nox_xxx_guiSpellTargetClickCheckSend_45DBB0()
	}
	if c.flags754064&4 != 0 {
		c.nox_ctrlevent_action_42E670(player.CCSpellPatternEnd, 0)
	}
	c.nox_xxx_clientControl_42D6B0_C()
	c.flags754064 = 0
	c.nox_xxx_clientControl_42D6B0_B()
}

func (c *CtrlEventHandler) nox_xxx_clientControl_42D6B0_orientation(mpos image.Point) {
	// calculates player orientation
	x := mpos.X
	y := mpos.Y
	if v15 := legacy.Nox_xxx_spriteGetMB_476F80(); v15 != nil {
		y = sub_4739D0(v15.Pos().Y)
	}
	wsz := videoGetWindowSize()
	cx := x - wsz.X/2
	cy := y - wsz.Y/2
	rad := math.Atan2(float64(cy), float64(cx))

	// TODO: support gamepad again
	//if inpHandler.Gamepad().InRelative() {
	//	p := inpHandler.Gamepad().StickRel().Posf()
	//	rad = math.Atan2(float64(p.Y), float64(p.X))
	//}

	// represent as integer
	ang := int((rad+2*math.Pi)*128.0/math.Pi + 0.5)
	if ang < 0 {
		ang += int((uint(255-ang) >> 8) << 8)
	}
	if ang >= 256 {
		ang -= int((uint(ang) >> 8) << 8)
	}
	c.playerDir = uint32(ang)
	c.nox_ctrlevent_action_42E670(player.CCOrientation, c.playerDir)
}

func (c *CtrlEventHandler) nox_xxx_clientControl_42D6B0_A(a4 *CtrlEventBinding) {

	fps := noxServer.TickRate()
	pl := getCurPlayer()

	var FoundManualInvertKey = false //stores whether or not the invert key was held this frame

	for it := a4; it != nil; it = it.field21 {
		for _, k := range it.events {
			switch k {
			case keybind.EventAction:
				cl := noxClient
				switch cl.Nox_client_getCursorType() {
				case gui.CursorShop:
					v10 := legacy.Nox_xxx_clientGetSpriteAtCursor_476F90()
					legacy.Nox_xxx_clientTrade_42E850(v10)
				case gui.CursorTalk:
					v11 := legacy.Nox_xxx_clientGetSpriteAtCursor_476F90()
					if v11.Flags70()&0x10 != 0 {
						legacy.Nox_xxx_clientTalk_42E7B0(v11)
					}
				case gui.CursorUse:
					v12 := legacy.Nox_xxx_clientGetSpriteAtCursor_476F90()
					legacy.Nox_xxx_clientCollideOrUse_42E810(v12)
				default:
					c.nox_ctrlevent_action_42E670(player.CCAction, 0) // regular attack?
				}
			case keybind.EventMoveForward:
				var flags uint32
				flags |= 0x1
				if c.flags754064&0x8 != 0 {
					flags |= 0x2
				}
				c.nox_ctrlevent_action_42E670(player.CCMoveForward, flags)
			case keybind.EventMoveBackward, keybind.EventMoveLeft, keybind.EventMoveRight:
				var code player.CtrlCode
				switch k {
				case keybind.EventMoveBackward:
					code = player.CCMoveBackward
				case keybind.EventMoveLeft:
					code = player.CCMoveLeft
				case keybind.EventMoveRight:
					code = player.CCMoveRight
				}
				var flags uint32
				if c.flags754064&0x1 != 0 {
					flags |= 0x1
				}
				if c.flags754064&0x8 != 0 {
					flags |= 0x2
				}
				c.nox_ctrlevent_action_42E670(code, flags)
			case keybind.EventJump:
				c.nox_ctrlevent_action_42E670(player.CCJump, 0)
			case keybind.EventToggleInventory:
				c.nox_ctrlevent_action_42E670(player.CCInventory, 0)
			case keybind.EventChat:
				c.nox_ctrlevent_action_42E670(player.CCChat, 0)
			case keybind.EventTeamChat:
				if pl.Field3680&0x1 == 0 {
					c.nox_ctrlevent_action_42E670(player.CCTeamChat, 0)
				}
			case keybind.EventToggleBook:
				c.nox_ctrlevent_action_42E670(player.CCReadSpellbook, 0)
			case keybind.EventToggleConsole:
				c.nox_ctrlevent_action_42E670(player.CCToggleConsole, 0)
			case keybind.EventIncreaseWindowSize:
				if !nox_xxx_checkGameFlagPause_413A50() {
					c.nox_ctrlevent_action_42E670(player.CCIncreaseWindowSize, 0)
				}
			case keybind.EventDecreaseWindowSize:
				if !nox_xxx_checkGameFlagPause_413A50() {
					c.nox_ctrlevent_action_42E670(player.CCDecreaseWindowSize, 0)
				}
			case keybind.EventIncreaseGamma:
				c.nox_ctrlevent_action_42E670(player.CCIncreaseGamma, 0)
			case keybind.EventDecreaseGamma:
				c.nox_ctrlevent_action_42E670(player.CCDecreaseGamma, 0)
			case keybind.EventToggleMap:
				c.nox_ctrlevent_action_42E670(player.CCReadMap, 0)
			case keybind.EventCastQueued:
				if !nox_xxx_guiSpellTest_45D9C0() {
					targ := legacy.Nox_xxx_packetGetMarshall_476F40()
					c.nox_ctrlevent_action_42E780(player.CCCastQueuedSpell, targ)
				}
			case keybind.EventPhonemeUN:
				if !nox_xxx_guiSpellTest_45D9C0() {
					c.nox_ctrlevent_action_42E780(player.CCSpellGestureUp, 0)
				}
			case keybind.EventPhonemeZO:
				if !nox_xxx_guiSpellTest_45D9C0() {
					c.nox_ctrlevent_action_42E780(player.CCSpellGestureDown, 0)
				}
			case keybind.EventPhonemeET:
				if !nox_xxx_guiSpellTest_45D9C0() {
					c.nox_ctrlevent_action_42E780(player.CCSpellGestureLeft, 0)
				}
			case keybind.EventPhonemeCHA:
				if !nox_xxx_guiSpellTest_45D9C0() {
					c.nox_ctrlevent_action_42E780(player.CCSpellGestureRight, 0)
				}
			case keybind.EventPhonemeIN:
				if !nox_xxx_guiSpellTest_45D9C0() {
					c.nox_ctrlevent_action_42E780(player.CCSpellGestureUpperRight, 0)
				}
			case keybind.EventPhonemeKA:
				if !nox_xxx_guiSpellTest_45D9C0() {
					c.nox_ctrlevent_action_42E780(player.CCSpellGestureUpperLeft, 0)
				}
			case keybind.EventPhonemeDO:
				if !nox_xxx_guiSpellTest_45D9C0() {
					c.nox_ctrlevent_action_42E780(player.CCSpellGestureLowerRight, 0)
				}
			case keybind.EventPhonemeRO:
				if !nox_xxx_guiSpellTest_45D9C0() {
					c.nox_ctrlevent_action_42E780(player.CCSpellGestureLowerLeft, 0)
				}
			case keybind.EventSpellEnd:
				if !nox_xxx_guiSpellTest_45D9C0() {
					c.nox_ctrlevent_action_42E780(player.CCSpellPatternEnd, 0)
				}
			case keybind.EventCastRecent:
				if !nox_xxx_guiSpellTest_45D9C0() {
					targ := legacy.Nox_xxx_packetGetMarshall_476F40()
					c.nox_ctrlevent_action_42E780(player.CCCastMostRecentSpell, targ)
				}
			case keybind.EventInvokeSlot1:
				c.nox_ctrlevent_action_42E780(player.CCCastSpell1, 0)
			case keybind.EventInvokeSlot2:
				c.nox_ctrlevent_action_42E780(player.CCCastSpell2, 0)
			case keybind.EventInvokeSlot3:
				c.nox_ctrlevent_action_42E780(player.CCCastSpell3, 0)
			case keybind.EventInvokeSlot4:
				c.nox_ctrlevent_action_42E780(player.CCCastSpell4, 0)
			case keybind.EventInvokeSlot5:
				c.nox_ctrlevent_action_42E780(player.CCCastSpell5, 0)
			case keybind.EventSpellSet1, keybind.EventSpellSet2, keybind.EventSpellSet3, keybind.EventSpellSet4, keybind.EventSpellSet5:
				legacy.Nox_xxx_clientUpdateButtonRow_45E110(int(k - keybind.EventSpellSet1))
			case keybind.EventMapZoomIn:
				c.nox_ctrlevent_action_42E670(player.CCMapZoomIn, 0)
			case keybind.EventMapZoomOut:
				c.nox_ctrlevent_action_42E670(player.CCMapZoomOut, 0)
			case keybind.EventSwapWeapons:
				c.nox_ctrlevent_action_42E780(player.CCNextWeapon, 0)
			case keybind.EventQuickHealth:
				c.nox_ctrlevent_action_42E780(player.CCQuickHealthPotion, 0)
			case keybind.EventQuickMana:
				c.nox_ctrlevent_action_42E780(player.CCQuickManaPotion, 0)
			case keybind.EventQuickCurePoison:
				c.nox_ctrlevent_action_42E780(player.CCQuickCurePoisonPotion, 0)
			case keybind.EventNextSpellSet:
				c.nox_ctrlevent_action_42E780(player.CCNextSpellSet, 0)
			case keybind.EventPreviousSpellSet:
				c.nox_ctrlevent_action_42E780(player.CCPreviousSpellSet, 0)
			case keybind.EventSelectSpellSet:
				c.nox_ctrlevent_action_42E780(player.CCSelectSpellSet, 0)
			case keybind.EventPlaceTrapBomber:
				c.nox_ctrlevent_action_42E780(player.CCBuildTrap, 0)
			case keybind.EventTrapDrop:
				if inputKeyCheckTimeout(k, fps/4) {
					inputSetKeyTimeout(k)
					gameexDropTrap()
				}
			case keybind.EventTrapBomber1, keybind.EventTrapBomber2, keybind.EventTrapBomber3:
				if inputKeyCheckTimeout(k, fps/4) {
					inputSetKeyTimeout(k)
					legacy.Nox_client_trapSetSelect_4604B0(int(k - keybind.EventTrapBomber1))
					clientPlaySoundSpecial(sound.SoundChangeSpellbar, 100)
				}
			case keybind.EventQuit:
				c.nox_ctrlevent_action_42E670(player.CCQuit, 0)
			case keybind.EventToggleQuitMenu:
				c.nox_ctrlevent_action_42E670(player.CCQuitMenu, 0)
			case keybind.EventToggleServerMenu:
				c.nox_ctrlevent_action_42E670(player.CCServerOptions, 0)
			case keybind.EventTaunt:
				if (allowEmotionsInSolo || noxflags.HasGame(noxflags.GameOnline)) && inputKeyCheckTimeoutLegacy(21, fps) {
					inputSetKeyTimeoutLegacy(21)
					c.nox_ctrlevent_action_42E670(player.CCTaunt, 0)
				}
			case keybind.EventLaugh:
				if (allowEmotionsInSolo || noxflags.HasGame(noxflags.GameOnline)) && inputKeyCheckTimeoutLegacy(20, 2*fps) {
					inputSetKeyTimeoutLegacy(20)
					c.nox_ctrlevent_action_42E670(player.CCLaugh, 0)
				}
			case keybind.EventPoint:
				if (allowEmotionsInSolo || noxflags.HasGame(noxflags.GameOnline)) && inputKeyCheckTimeoutLegacy(22, fps) {
					inputSetKeyTimeoutLegacy(22)
					c.nox_ctrlevent_action_42E670(player.CCPoint, 0)
				}
			case keybind.EventInvertSpellTarget:
				c.nox_ctrlevent_action_42E670(player.CCInvertSpellTarget, 0)
			case keybind.EventInvertSpellTargetManualCast:
				if inputKeyCheckTimeout(k, fps/4) {
					if !wasManualSpellInvertKeyPressed {
						fmt.Println("Just pressed") //Key was not pressed, but is pressed now
						wasManualSpellInvertKeyPressed = true
					}
					inputSetKeyTimeout(k)
				}
				FoundManualInvertKey = true
			case keybind.EventToggleRank:
				if noxflags.HasGame(noxflags.GameOnline) {
					c.nox_ctrlevent_action_42E670(player.CCToggleRank, 0)
				}
			case keybind.EventToggleNetstat:
				if !nox_xxx_checkGameFlagPause_413A50() {
					c.nox_ctrlevent_action_42E670(player.CCToggleNetstat, 0)
				}
			case keybind.EventToggleGUI:
				c.nox_ctrlevent_action_42E670(player.CCToggleGUI, 0)
			case keybind.EventAutoSave:
				if noxflags.HasGame(noxflags.GameModeCoop) && nox_xxx_guiCursor_477600() == 0 {
					c.nox_ctrlevent_action_42E670(player.CCAutoSave, 0)
				}
			case keybind.EventAutoLoad:
				if noxflags.HasGame(noxflags.GameModeCoop) && nox_xxx_guiCursor_477600() == 0 {
					c.nox_ctrlevent_action_42E670(player.CCAutoLoad, 0)
				}
			case keybind.EventScreenShot:
				c.nox_ctrlevent_action_42E670(player.CCScreenShot, 0)
			case keybind.EventCreaturesBanish:
				if inputKeyCheckTimeout(k, fps/4) {
					inputSetKeyTimeout(k)
					legacy.Nox_client_orderCreature(0, 0)
				}
			case keybind.EventCreaturesGuard:
				if inputKeyCheckTimeout(k, fps/4) {
					inputSetKeyTimeout(k)
					legacy.Nox_client_orderCreature(0, 3)
				}
			case keybind.EventCreaturesEscort:
				if inputKeyCheckTimeout(k, fps/4) {
					inputSetKeyTimeout(k)
					legacy.Nox_client_orderCreature(0, 4)
				}
			case keybind.EventCreaturesHunt:
				if inputKeyCheckTimeout(k, fps/4) {
					inputSetKeyTimeout(k)
					legacy.Nox_client_orderCreature(0, 5)
				}
			case keybind.EventAcceptItemsBatch:
				if inputKeyCheckTimeout(k, fps/4) {
					inputSetKeyTimeout(k)
					clientAcceptTradeOrDrop()
				}
			}
		}
	}

	if wasManualSpellInvertKeyPressed && !FoundManualInvertKey {
		wasManualSpellInvertKeyPressed = false
		// TODO: Play spellbar shifting sound

		var a bool = noxClient.ManualCastDirectionInverted
		var b bool = !a
		fmt.Println("Toggling inversion of manual cast from ", a , " to ", b)
		noxClient.ManualCastDirectionInverted = b
	}

	// NOTE: Sending CCInvertSpellTarget to the server flips the player's spellbar arrows, but does not affect the direction of the manual cast
	//if noxClient.ManualCastDirectionInverted {
	//	c.nox_ctrlevent_action_42E670(player.CCInvertSpellTarget, 0)
	//}

}

func clientAcceptTradeOrDrop() {
	dialog := legacy.Get_nox_gui_itemAmount_dialog_1319228()
	if dialog == nil {
		return
	}
	if dialog.GetFlags().IsHidden() {
		return
	}
	accept := dialog.ChildByID(3606)
	legacy.Sub_4C01C0(0, 16391, accept.C(), 0)
}

func clientSetPhonemeFrame(a1 int) {
	*memmap.PtrUint32(0x5D4594, 1096596+4*uintptr(a1)) = noxServer.Frame()
}

func (c *Client) nox_client_increaseViewport_4766E0() {
	c.nox_draw_setCutSize(0, 2)
}

func (c *Client) nox_client_decreaseViewport_4766F0() {
	c.nox_draw_setCutSize(0, -2)
}

func (c *CtrlEventHandler) nox_xxx_clientControl_42D6B0_B() {
	indStart := 0
	indEnd := 0
	if noxflags.HasGame(noxflags.GameHost) {
		indStart = c.indB
		indEnd = c.indA
		if c.indA < c.indB {
			indEnd = c.indA + cap(c.bufA)
		}
	} else {
		indEnd = c.indA
	}
	for i := indStart; i < indEnd; i++ {
		ce := &c.bufA[i%cap(c.bufA)]
		if !ce.Active {
			continue
		}
		switch ce.Code {
		case player.CCSpellGestureUp: // un
			if !noxflags.HasGame(noxflags.GameModeChat) {
				clientPlaySoundSpecial(sound.SoundSpellPhonemeUp, 100)
				clientSetPhonemeFrame(1)
			}
		case player.CCSpellGestureDown: // zo
			if !noxflags.HasGame(noxflags.GameModeChat) {
				clientPlaySoundSpecial(sound.SoundSpellPhonemeDown, 100)
				clientSetPhonemeFrame(6)
			}
		case player.CCSpellGestureLeft: // et
			if !noxflags.HasGame(noxflags.GameModeChat) {
				clientPlaySoundSpecial(sound.SoundSpellPhonemeLeft, 100)
				clientSetPhonemeFrame(3)
			}
		case player.CCSpellGestureRight: // cha
			if !noxflags.HasGame(noxflags.GameModeChat) {
				clientPlaySoundSpecial(sound.SoundSpellPhonemeRight, 100)
				clientSetPhonemeFrame(4)
			}
		case player.CCSpellGestureUpperRight: // in
			if !noxflags.HasGame(noxflags.GameModeChat) {
				clientPlaySoundSpecial(sound.SoundSpellPhonemeUpRight, 100)
				clientSetPhonemeFrame(2)
			}
		case player.CCSpellGestureUpperLeft: // ka
			if !noxflags.HasGame(noxflags.GameModeChat) {
				clientPlaySoundSpecial(sound.SoundSpellPhonemeUpLeft, 100)
				clientSetPhonemeFrame(0)
			}
		case player.CCSpellGestureLowerRight: // do
			if !noxflags.HasGame(noxflags.GameModeChat) {
				clientPlaySoundSpecial(sound.SoundSpellPhonemeDownRight, 100)
				clientSetPhonemeFrame(7)
			}
		case player.CCSpellGestureLowerLeft: // ro
			if !noxflags.HasGame(noxflags.GameModeChat) {
				clientPlaySoundSpecial(sound.SoundSpellPhonemeDownLeft, 100)
				clientSetPhonemeFrame(5)
			}
		case player.CCChat:
			legacy.Nox_client_chatStart_46A430(0)
			ce.Active = false
		case player.CCTeamChat:
			legacy.Nox_client_chatStart_46A430(1)
			ce.Active = false
		case player.CCReadSpellbook:
			legacy.Nox_client_toggleSpellbook_45AC70()
			ce.Active = false
		case player.CCToggleConsole:
			guiCon.Toggle()
			ce.Active = false
		case player.CCIncreaseWindowSize:
			clientPlaySoundSpecial(sound.SoundShellClick, 100)
			noxClient.nox_client_increaseViewport_4766E0()
			ce.Active = false
		case player.CCDecreaseWindowSize:
			clientPlaySoundSpecial(sound.SoundShellClick, 100)
			noxClient.nox_client_decreaseViewport_4766F0()
			ce.Active = false
		case player.CCIncreaseGamma:
			clientPlaySoundSpecial(sound.SoundShellClick, 100)
			updateGamma(1)
			ce.Active = false
		case player.CCDecreaseGamma:
			clientPlaySoundSpecial(sound.SoundShellClick, 100)
			updateGamma(-1)
			ce.Active = false
		case player.CCQuit:
			nox_client_quit_4460C0()
			ce.Active = false
		case player.CCQuitMenu:
			if sub_450560() {
				nox_savegame_sub_46D580()
			} else {
				legacy.Sub_42EB90(1)
			}
			ce.Active = false
		case player.CCReadMap:
			nox_client_toggleMap_473610()
			ce.Active = false
		case player.CCInventory:
			legacy.Nox_client_toggleInventory_467C60()
			ce.Active = false
		case player.CCCastSpell1, player.CCCastSpell2, player.CCCastSpell3, player.CCCastSpell4, player.CCCastSpell5:
			legacy.Nox_client_invokeSpellSlot_45DA50(int(ce.Code - player.CCCastSpell1))
			ce.Active = false
		case player.CCMapZoomIn:
			legacy.Nox_client_mapZoomIn_4724E0()
			ce.Active = false
		case player.CCMapZoomOut:
			legacy.Nox_client_mapZoomOut_472500()
			ce.Active = false
		case player.CCNextWeapon:
			legacy.Nox_client_invAlterWeapon_4672C0()
			ce.Active = false
		case player.CCQuickHealthPotion:
			legacy.Nox_client_quickHealthPotion_472220()
			ce.Active = false
		case player.CCQuickManaPotion:
			legacy.Nox_client_quickManaPotion_472240()
			ce.Active = false
		case player.CCQuickCurePoisonPotion:
			legacy.Nox_client_quickCurePoisonPotion_472260()
			ce.Active = false
		case player.CCNextSpellSet:
			legacy.Nox_client_spellSetNext_4604F0()
			ce.Active = false
		case player.CCPreviousSpellSet:
			legacy.Nox_client_spellSetPrev_460540()
			ce.Active = false
		case player.CCSelectSpellSet:
			legacy.Nox_client_spellSetSelect_460590()
			ce.Active = false
		case player.CCBuildTrap:
			legacy.Nox_client_buildTrap_45E040()
			ce.Active = false
		case player.CCServerOptions:
			if !(noxflags.HasGame(noxflags.GameFlag4) || !noxflags.HasGame(noxflags.GameOnline)) {
				legacy.Nox_xxx_guiServerOptsLoad_457500()
			}
			ce.Active = false
		case player.CCLaugh:
			nox_xxx_netClientSendSocial(server.HostPlayerIndex, 1, 0, 1)
		case player.CCTaunt:
			nox_xxx_netClientSendSocial(server.HostPlayerIndex, 2, 0, 1)
		case player.CCPoint:
			nox_xxx_netClientSendSocial(server.HostPlayerIndex, 4, 0, 1)
		case player.CCInvertSpellTarget:
			legacy.Sub_460630()
			ce.Active = false
		case player.CCToggleRank:
			clientPlaySoundSpecial(sound.SoundShellClick, 100)
			sub_4703F0()
			ce.Active = false
		case player.CCToggleNetstat:
			noxClient.guiFPS.sub_470A60()
			clientPlaySoundSpecial(sound.SoundShellClick, 100)
			ce.Active = false
		case player.CCToggleGUI:
			if !noxflags.HasEngine(noxflags.EngineNoRendering) {
				nox_client_renderGUI_80828 = !nox_client_renderGUI_80828
				nox_xxx_xxxRenderGUI_587000_80832 = nox_client_renderGUI_80828
				clientPlaySoundSpecial(sound.SoundShellClick, 100)
			}
			ce.Active = false
		case player.CCAutoSave:
			if noxflags.HasGame(noxflags.GameModeCoop) {
				if legacy.Nox_xxx_game_4DCCB0() != 0 {
					clientPlaySoundSpecial(sound.SoundShellClick, 100)
					SaveCoop(common.SaveAuto)
				} else {
					clientPlaySoundSpecial(sound.SoundPermanentFizzle, 100)
				}
			}
			ce.Active = false
		case player.CCAutoLoad:
			if noxflags.HasGame(noxflags.GameModeCoop) {
				if legacy.Nox_xxx_game_4DCCB0() != 0 {
					clientPlaySoundSpecial(sound.SoundShellClick, 100)
					sub_413A00(1)
					v41 := strMan.GetStringInFile("GUIQuit.c:ReallyLoadMessage", "Ctrlevnt.c")
					v40 := strMan.GetStringInFile("SelChar.c:LoadLabel", "Ctrlevnt.c")
					NewDialogWindow(nil, v40, v41, gui.DialogYesButton|gui.DialogNoButton|gui.DialogFlag6, func() {
						sub_413A00(0)
						clientLoadCoopGame(common.SaveAuto)
					}, func() {
						sub_413A00(0)
					})
				} else {
					clientPlaySoundSpecial(sound.SoundPermanentFizzle, 100)
				}
			}
			ce.Active = false
		case player.CCScreenShot:
			makeScreenshot()
			ce.Active = false
		}
	}
}

func (c *CtrlEventHandler) nox_xxx_clientControl_42D6B0_C() {
	if noxflags.HasGame(noxflags.GameHost) {
		v20 := c.indB
		cnt := 0
		for i := c.indD; i != v20; i = (i + 1) % cap(c.bufA) {
			c.bufB[cnt] = c.bufA[i]
			cnt++
		}
		if cnt > 0 {
			c.indD = v20
			c.indC = cnt
		}
		v23 := c.indC
		for i := v20; i != c.indA; i = (i + 1) % cap(c.bufA) {
			if v23 <= 0 {
				continue
			}
			p1 := &c.bufA[i]
			for j := 0; j < v23; j++ {
				p2 := &c.bufB[j]
				if p1.Code == p2.Code {
					if p1.Code.CanPauseMode() {
						p1.Active = false
					}
					v23 = c.indC
					break
				}
			}
		}
	} else {
		for i := 0; i < c.indA; i++ {
			p1 := &c.bufA[i]
			p1.Active = true
			if c.indC <= 0 {
				continue
			}
			for j := 0; j < c.indC; j++ {
				p2 := &c.bufB[j]
				if p1.Code == p2.Code {
					if p1.Code.CanPauseMode() {
						p1.Active = false
					}
					break
				}
			}
		}
		n := c.indA
		copy(c.bufB[:n], c.bufA[:n])
		c.indC = c.indA
	}
}

func (c *CtrlEventHandler) nox_ctrlevent_action_42E670(code player.CtrlCode, data uint32) {
	if !noxflags.HasGame(noxflags.GameHost) {
		if c.indA >= cap(c.bufA) {
			return
		}
	} else {
		if c.indA+1 == c.indB {
			return
		}
	}
	if getPlayerClass() != player.Warrior || !code.IsMagicRelated() {
		j := c.indA
		if noxflags.HasGame(noxflags.GameHost) && code.CanPauseMode() {
			for i := c.indB; i != c.indA; i = (i + 1) % cap(c.bufA) {
				if c.bufA[i].Code == code {
					return
				}
			}
		}
		p := &c.bufA[j]
		p.tick = c.ticks
		p.PlayerCtrl = server.PlayerCtrl{
			Code: code, Data: data, Active: true,
		}
		c.indA = j + 1
		if noxflags.HasGame(noxflags.GameHost) {
			c.indA %= cap(c.bufA)
		}
	}
}

func (c *CtrlEventHandler) nox_ctrlevent_action_42E780(code player.CtrlCode, data uint32) {
	if p := noxClient.ClientPlayerUnit(); p != nil && p.Flags()&2 == 0 {
		if !nox_xxx_checkGameFlagPause_413A50() {
			c.nox_ctrlevent_action_42E670(code, data)
		}
	}
}

func nox_ctrlevent_add_ticks_42E630() uint32 {
	v0 := sub_416640()
	switch *(*uint32)(unsafe.Add(unsafe.Pointer(&v0[0]), 66)) {
	case 1:
		return legacy.Sub_554290()
	case 2:
		return legacy.Sub_554300()
	case 3:
		return *(*uint32)(unsafe.Add(unsafe.Pointer(&v0[0]), 70))
	}
	return 0
}

func (c *CtrlEventHandler) Sub_42E8E0_go(ev keybind.Event, ind int) string {
	for v2 := c.bindings; v2 != nil; v2 = v2.next {
		for _, e := range v2.events {
			if e != ev {
				continue
			}
			ind--
			if ind <= 0 {
				if key := v2.defKey(); key.IsValid() {
					return key.Title(strMan)
				}
			}
		}
	}
	return ""
}

func (c *CtrlEventHandler) addBinding(ce *CtrlEventBinding) {
	ce.prev = nil
	ce.next = c.bindings
	if c.bindings != nil {
		c.bindings.prev = ce
	}
	c.bindings = ce
}

func (c *CtrlEventHandler) listBindings() []*CtrlEventBinding {
	var tail *CtrlEventBinding
	for it := c.bindings; it != nil; it = it.next {
		tail = it
	}
	if tail == nil {
		return nil
	}
	var out []*CtrlEventBinding
	for it := tail; it != nil; it = it.prev {
		out = append(out, it)
	}
	return out
}

func nox_client_parseConfigHotkeysLine_42CF50(a1 string) int {
	r := strings.NewReader(a1)
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if len(line) == 0 || line[0] == '-' || line[0] == '#' {
			continue
		}
		i := strings.IndexByte(line, '=')
		if i < 0 {
			input.Log.Printf("invalid config line: %q", line)
			continue
		}
		key, val := strings.TrimSpace(line[:i]), strings.TrimSpace(line[i+1:])
		return nox_client_parseConfigHotkeysLine(key, val)
	}
	return 0
}

func nox_client_parseConfigHotkeysLine(key, val string) int {
	c := noxClient
	switch key {
	case "MousePickup":
		val = strings.ToLower(val)
		for i, s := range noxMouseSelectOpt {
			if val == strings.ToLower(s) {
				legacy.Sub_430AA0(i)
				return 1
			}
		}
		legacy.Sub_430AA0(0)
		return 1
	}
	ce := new(CtrlEventBinding)
	for _, s := range strings.Split(key, "+") {
		s = strings.TrimSpace(s)
		if k := keybind.KeyByName(s); k != 0 {
			ce.keys = append(ce.keys, k)
		}
	}
	for _, s := range strings.Split(val, "+") {
		s = strings.TrimSpace(s)
		if b := keyBinding.EventByName(s); b != nil && b.Event != 0 {
			ce.events = append(ce.events, b.Event)
		}
	}
	c.ctrl.addBinding(ce)
	return 1
}

func writeConfigHotkeys(sect *cfg.Section) {
	c := noxClient
	v1 := int(legacy.Nox_client_mousePriKey_430AF0())
	sect.Set("MousePickup", noxMouseSelectOpt[v1])
	for _, it := range c.ctrl.listBindings() {
		var keys []string
		for _, k := range it.keys {
			if !k.IsValid() {
				continue
			}
			keys = append(keys, k.String())
		}
		var vals []string
		for _, k := range it.events {
			if b := keyBinding.EventByCode(k); b != nil && b.Name != "" {
				vals = append(vals, b.Name)
			}
		}
		sect.Set(strings.Join(keys, " + "), strings.Join(vals, " + "))
	}
}

func (c *CtrlEventHandler) writeToNetBuffer() {
	buf := c.netBuf[:0]
	ticks := uint32(platformTicks())
	if noxflags.HasGame(noxflags.GameHost) {
		for i := c.indB; i != c.indA; i = (i + 1) % cap(c.bufA) {
			p := &c.bufA[i]
			if p.tick > ticks+50 {
				break
			}
			if !p.Active {
				continue
			}
			buf = append(buf, byte(p.Code), 0, 0, 0)
			if sz := p.Code.DataSize(); sz != 0 {
				var b [4]byte
				binary.LittleEndian.PutUint32(b[:], p.Data)
				buf = append(buf, b[:sz]...)
			}
		}
		c.indB = c.indA
	} else {
		for i := 0; i < c.indA; i++ {
			p := &c.bufA[i]
			if !p.Active {
				continue
			}
			buf = append(buf, byte(p.Code), 0, 0, 0)
			if sz := p.Code.DataSize(); sz != 0 {
				var b [4]byte
				binary.LittleEndian.PutUint32(b[:], p.Data)
				buf = append(buf, b[:sz]...)
			}
		}
	}
	c.netBuf = buf
}

func (c *CtrlEventHandler) hasDefBinding(ev keybind.Event, key keybind.Key) bool {
	for it := c.bindings; it != nil; it = it.next {
		if it.defEvent() == ev && it.defKey() == key {
			return true
		}
	}
	return false
}
