package opennox

import (
	"github.com/noxworld-dev/opennox-lib/types"

	"github.com/noxworld-dev/opennox/v1/client"
	"github.com/noxworld-dev/opennox/v1/client/gui"
	"github.com/noxworld-dev/opennox/v1/client/noxrender"
	"github.com/noxworld-dev/opennox/v1/common/ntype"
	"github.com/noxworld-dev/opennox/v1/internal/cryptfile"
	"github.com/noxworld-dev/opennox/v1/legacy"
	"github.com/noxworld-dev/opennox/v1/server"
)

func init() {
	legacy.WriteConfigLegacy = nox_common_writecfgfile
	legacy.Sub_57C490_2 = sub_57C490_2
	legacy.Nox_thing_debug_draw = nox_thing_debug_draw
	legacy.Sub_4E8290 = sub_4E8290
	legacy.Nox_client_parseConfigHotkeysLine_42CF50 = nox_client_parseConfigHotkeysLine_42CF50
	legacy.Nox_client_getIntroScreenDuration_44E3B0 = nox_client_getIntroScreenDuration_44E3B0
	legacy.Nox_client_getBriefDuration = nox_client_getBriefDuration
	legacy.Nox_game_exit_xxx2 = nox_game_exit_xxx2
	legacy.Sub_470510 = sub_470510
	legacy.Sub_4703F0 = sub_4703F0
	legacy.Nox_xxx_cliDrawConnectedLoop_43B360 = nox_xxx_cliDrawConnectedLoop_43B360
	legacy.Nox_client_guiXxxDestroy_4A24A0 = nox_client_guiXxxDestroy_4A24A0
	legacy.Nox_client_quit_4460C0 = nox_client_quit_4460C0
	legacy.Nox_gui_makeAnimation_43C5B0 = nox_gui_makeAnimation
	legacy.Nox_client_gui_set_flag_815132 = nox_client_gui_set_flag_815132
	legacy.Nox_client_onClientStatusA = nox_client_onClientStatusA
	legacy.Nox_client_setRenderGUI = nox_client_setRenderGUI
	legacy.Nox_client_getRenderGUI = nox_client_getRenderGUI
	legacy.Nox_gui_console_flagXxx_451410 = nox_gui_console_flagXxx_451410
	legacy.Nox_gui_console_Hide_4512B0 = nox_gui_console_Hide_4512B0
	legacy.Nox_xxx_gLoadImg = nox_xxx_gLoadImg
	legacy.Nox_xxx_gLoadAnim = nox_xxx_gLoadAnim
	legacy.Nox_xxx_dialogMsgBoxCreate_449A10 = NewDialogWindow
	legacy.Sub_449E00 = sub449E00
	legacy.Sub_449E30 = sub449E30
	legacy.Sub_449E60 = sub_449E60
	legacy.Sub_449EA0 = sub449EA0
	legacy.Sub_44A4E0 = sub_44A4E0
	legacy.Sub_44A4B0 = sub_44A4B0
	legacy.Sub_44A360 = sub_44A360
	legacy.Sub_4706C0 = func(a1 int) {
		noxClient.guiFPS.sub_4706C0(a1)
	}
	legacy.Sub_473670 = sub_473670
	legacy.Nox_client_toggleMap_473610 = nox_client_toggleMap_473610
	legacy.Nox_client_refreshServerList_4378B0 = nox_client_refreshServerList_4378B0
	legacy.Sub_438770_waitList = sub_438770_waitList
	legacy.Nox_xxx_createSocketLocal = nox_xxx_createSocketLocal
	legacy.Sub_554D10 = sub_554D10
	legacy.Nox_gui_menu_proc_ext = nox_gui_menu_proc_ext
	legacy.Nox_video_setMenuOptions = nox_video_setMenuOptions
	legacy.Sub_4A19F0 = guiSetBackButtonText
	legacy.Sub_4C3A90 = sub_4C3A90
	legacy.Sub_4CBE70 = sub_4CBE70
	legacy.Sub_4AAA10 = sub_4AAA10
	legacy.Nox_new_window_from_file = nox_new_window_from_file
	legacy.Sub_445B40 = clientLoadCoopAuto445B40
	legacy.Nox_xxx_mapGenMakeInfo_4D5DB0 = nox_xxx_mapGenMakeInfo_4D5DB0
	legacy.Nox_common_checkMapFile = nox_common_checkMapFile
	legacy.Nox_xxx_mapWriteSectionsMB_426E20 = nox_xxx_mapWriteSectionsMB_426E20
	legacy.NetworkLogPrint = networkLogPrint
	legacy.ClientSetServerHost = clientSetServerHost
	legacy.Nox_client_joinGame_438A90 = nox_client_joinGame_438A90
	legacy.Nox_xxx_moveUpdateSpecial_517970 = func(obj *server.Object) {
		noxServer.nox_xxx_moveUpdateSpecial_517970(obj)
	}
	legacy.Nox_xxx_unitMove_4E7010 = func(a1 *server.Object, pos types.Pointf) {
		asObjectS(a1).SetPos(pos)
	}
	legacy.Nox_mapToGameFlags = nox_mapToGameFlags
	legacy.Sub_40A1A0 = sub_40A1A0
	legacy.PlatformTicks = platformTicks
	legacy.Nox_ticks_reset_416D40 = nox_ticks_reset_416D40
	legacy.Nox_xxx_spellAwardAll1_4EFD80 = nox_xxx_spellAwardAll1_4EFD80
	legacy.Nox_xxx_spellAwardAll2_4EFC80 = nox_xxx_spellAwardAll2_4EFC80
	legacy.Nox_xxx_spellAwardAll3_4EFE10 = nox_xxx_spellAwardAll3_4EFE10
	legacy.Nox_xxx_spellTitle_424930 = nox_xxx_spellTitle_424930
	legacy.Nox_xxx_spellDescription_424A30 = nox_xxx_spellDescription_424A30
	legacy.Nox_xxx_spellIcon_424A90 = nox_xxx_spellIcon_424A90
	legacy.Nox_xxx_spellIconHighlight_424AB0 = nox_xxx_spellIconHighlight_424AB0
	legacy.Nox_xxx_castSpellByUser_4FDD20 = nox_xxx_castSpellByUser_4FDD20
	legacy.Nox_xxx_spellWallCreate_4FFA90 = func(sp *server.DurSpell) int {
		return noxServer.spells.walls.Create(sp)
	}
	legacy.Nox_xxx_spellWallUpdate_500070 = func(sp *server.DurSpell) int {
		return noxServer.spells.walls.Update(sp)
	}
	legacy.Nox_xxx_spellWallDestroy_500080 = func(sp *server.DurSpell) {
		noxServer.spells.walls.Destroy(sp)
	}
	legacy.Nox_xxx_spellCastByPlayer_4FEEF0 = func() {
		noxServer.spells.duration.spellCastByPlayer()
	}
	legacy.Nox_exit = nox_exit
	legacy.Nox_xxx_gameGetScreenBoundaries_43BEB0_get_video_mode = nox_xxx_gameGetScreenBoundaries_43BEB0_get_video_mode
	legacy.Sub_4AA9C0 = sub_4AA9C0
	legacy.Nox_server_questAllowDefault = nox_server_questAllowDefault
	legacy.Nox_server_questNextStageThreshold_4D74F0 = nox_server_questNextStageThreshold_4D74F0
	legacy.Sub_4D6F30 = sub_4D6F30
	legacy.GetNetPlayerBufSize = getNetPlayerBufSize
	legacy.Nox_netlist_addToMsgListSrv = nox_netlist_addToMsgListSrv
	legacy.Nox_client_setVersion_409AE0 = nox_client_setVersion_409AE0
	legacy.Nox_video_getCutSize_4766D0 = nox_video_getCutSize_4766D0
	legacy.Nox_video_setCutSize_4766A0 = nox_video_setCutSize_4766A0
	legacy.Nox_video_setGammaSlider = nox_video_setGammaSlider
	legacy.Sub_43BE50_get_video_mode_id = sub_43BE50_get_video_mode_id
	legacy.Get_video_mode_string = get_video_mode_string
	legacy.Nox_getBackbufWidth = nox_getBackbufWidth
	legacy.Nox_getBackbufHeight = nox_getBackbufHeight
	legacy.Nox_video_getFullScreen = nox_video_getFullScreen
	legacy.Nox_video_setFullScreen = nox_video_setFullScreen
	legacy.Sub_430C30_set_video_max = sub_430C30_set_video_max
	legacy.VideoGetMaxSize = videoGetMaxSize
	legacy.Nox_video_callCopyBackBuffer_4AD170 = nox_video_callCopyBackBuffer_4AD170
	legacy.Nox_getBackbufferPitch = nox_getBackbufferPitch
	legacy.Nox_client_clearScreen_440900 = nox_client_clearScreen_440900
	legacy.Nox_draw_setCutSize_476700 = nox_draw_setCutSize_476700
	legacy.Nox_xxx_bookSaveSpellForDragDrop_477640 = nox_xxx_bookSaveSpellForDragDrop_477640
	legacy.Nox_xxx_bookSpellDnDclear_477660 = nox_xxx_bookSpellDnDclear_477660
	legacy.Nox_xxx_bookGetSpellDnDType_477670 = nox_xxx_bookGetSpellDnDType_477670
	legacy.Nox_xxx_cursorSetDraggedItem_477690 = nox_xxx_cursorSetDraggedItem_477690
	legacy.Nox_xxx_cursorResetDraggedItem_4776A0 = nox_xxx_cursorResetDraggedItem_4776A0
	legacy.Sub_478000 = sub_478000
	legacy.Nox_xxx_GetEndgameDialog = nox_xxx_GetEndgameDialog
	legacy.GameGetPlayState = gameGetPlayState
	legacy.ServerCheatGod = serverCheatGod
	legacy.Nox_xxx_serverHost_43B4D0 = nox_xxx_serverHost_43B4D0
	legacy.Nox_xxx_netServerCmd_440950 = nox_xxx_netServerCmd_440950
	legacy.ExecConsoleCmd = execConsoleCmd
	legacy.Nox_xxx_gameIsNotMultiplayer_4DB250 = nox_xxx_gameIsNotMultiplayer_4DB250
	legacy.Nox_xxx_gameSetSwitchSolo_4DB220 = nox_xxx_gameSetSwitchSolo_4DB220
	legacy.Nox_xxx_gameIsSwitchToSolo_4DB240 = nox_xxx_gameIsSwitchToSolo_4DB240
	legacy.Nox_xxx_gameSetWallsDamage_4E25A0 = nox_xxx_gameSetWallsDamage_4E25A0
	legacy.GetDoDamageWalls = func() bool {
		return doDamageWalls
	}
	legacy.Nox_video_inFadeTransition_44E0D0 = nox_video_inFadeTransition_44E0D0
	legacy.Sub_4373A0 = sub_4373A0
	legacy.Sub_43EA20 = sub_43EA20
	legacy.Sub_43E9F0 = sub_43E9F0
	legacy.Sub_43E940 = sub_43E940
	legacy.Sub_43EFD0 = sub_43EFD0
	legacy.Sub_43EC10 = sub_43EC10
	legacy.Sub_43F130 = sub_43F130
	legacy.Sub_43ED00 = sub_43ED00
	legacy.Sub_43D650 = sub_43D650
	legacy.Sub_43D680 = sub_43D680
	legacy.Sub_43D6A0 = sub_43D6A0
	legacy.Sub_44D640 = sub_44D640
	legacy.Sub_44D7E0 = sub_44D7E0
	legacy.Nox_xxx_musicStartPlay_43D6C0 = nox_xxx_musicStartPlay_43D6C0
	legacy.Sub_44D660 = sub_44D660
	legacy.Sub_43F060 = sub_43F060
	legacy.Sub_43EC30 = sub_43EC30
	legacy.Sub_43ECB0 = sub_43ECB0
	legacy.Nox_xxx_updateSprings_5113A0 = nox_xxx_updateSprings_5113A0
	legacy.Nox_xxx_unitIsUnitTT_4E7C80 = nox_xxx_unitIsUnitTT_4E7C80
	legacy.Nox_xxx_updatePlayer_4F8100 = nox_xxx_updatePlayer_4F8100
	legacy.Nox_xxx_updatePixie_53CD20 = nox_xxx_updatePixie_53CD20
	legacy.Nox_xxx_updatePlayerObserver_4E62F0 = nox_xxx_updatePlayerObserver_4E62F0
	legacy.Nox_xxx_updateProjectile_53AC10 = nox_xxx_updateProjectile_53AC10
	legacy.Nox_xxx_updateDeathBall_53D080 = nox_xxx_updateDeathBall_53D080
	legacy.Nox_xxx_collideDeathBall_4E9E90 = nox_xxx_collideDeathBall_4E9E90
	legacy.Nox_xxx_castCounterSpell_52BBB0 = nox_xxx_castCounterSpell_52BBB0
	legacy.Nox_xxx_changeOwner_52BE40 = nox_xxx_changeOwner_52BE40
	legacy.Sub_44A4A0 = sub_44A4A0
	legacy.Nox_script_shouldReadMoreXxx = nox_script_shouldReadMoreXxx
	legacy.Nox_script_shouldReadEvenMoreXxx = nox_script_shouldReadEvenMoreXxx
	legacy.Nox_xxx_xfer_saveObj51DF90 = nox_xxx_xfer_saveObj51DF90
	legacy.Nox_xxx_XFerDefault4F49A0 = nox_xxx_XFerDefault4F49A0
	legacy.Nox_xxx_XFer_ReadShopItem_52A840 = nox_xxx_XFer_ReadShopItem_52A840
	legacy.Nox_xxx_XFer_WriteShopItem_52A5F0 = nox_xxx_XFer_WriteShopItem_52A5F0
	legacy.Sub_4FC670 = sub_4FC670
	legacy.Nox_xxx_playerExecuteAbil_4FBB70 = nox_xxx_playerExecuteAbil_4FBB70
	legacy.Sub_4FC0B0 = sub_4FC0B0
	legacy.Nox_xxx_playerCancelAbils_4FC180 = nox_xxx_playerCancelAbils_4FC180
	legacy.Sub_4FC300 = sub_4FC300
	legacy.Nox_xxx_abilityGetName_0_425260 = nox_xxx_abilityGetName_0_425260
	legacy.Nox_xxx_abilityCooldown_4252D0 = nox_xxx_abilityCooldown_4252D0
	legacy.Sub_4252F0 = sub_4252F0
	legacy.Nox_xxx_spellGetAbilityIcon_425310 = nox_xxx_spellGetAbilityIcon_425310
	legacy.Nox_xxx_bookFirstKnownAbil_425330 = nox_xxx_bookFirstKnownAbil_425330
	legacy.Nox_xxx_bookNextKnownAbil_425350 = nox_xxx_bookNextKnownAbil_425350
	legacy.Sub_425450 = sub_425450
	legacy.Nox_xxx_netAbilRepotState_4D8100 = nox_xxx_netAbilRepotState_4D8100
	legacy.Nox_setImaginaryCaster = func() int {
		return noxServer.nox_setImaginaryCaster()
	}
	legacy.Nox_script_readWriteZzz_541670 = nox_script_readWriteZzz_541670
	legacy.Sub_4C26F0 = sub_4C26F0
	legacy.WinMainMenuAnimOutStartFnc = winMainMenuAnimOutStartFnc
	legacy.Sub_44E320 = sub_44E320
	legacy.WinMainMenuAnimOutDoneFnc = winMainMenuAnimOutDoneFnc
	legacy.Sub_4A24C0 = sub_4A24C0
	legacy.Nox_game_showMainMenu_4A1C00 = nox_game_showMainMenu_4A1C00
	legacy.Sub_43BE40 = sub_43BE40
	legacy.Sub_43BE30 = sub_43BE30
	legacy.Sub_4A18E0 = sub_4A18E0
	legacy.Sub_42BFB0 = sub_42BFB0
	legacy.Nox_xxx_objectTOCgetTT = nox_xxx_objectTOCgetTT
	legacy.Sub_42C310 = sub_42C310
	legacy.Sub_42C2E0 = sub_42C2E0
	legacy.Sub_42C300 = sub_42C300
	legacy.Sub_42BFE0 = sub_42BFE0
	legacy.Sub_4E3AD0 = sub_4E3AD0
	legacy.Sub_4A5E90_A = sub_4A5E90_A
	legacy.Nox_xxx_fireRingEffect_4E05B0 = nox_xxx_fireRingEffect_4E05B0
	legacy.Nox_xxx_blueFREffect_4E05F0 = nox_xxx_blueFREffect_4E05F0
	legacy.Sub_4E3B80 = sub_4E3B80
	legacy.Nox_savegame_rm = deleteSaveDir
	legacy.Nox_client_countPlayerFiles04_4DC7D0 = nox_client_countPlayerFiles04_4DC7D0
	legacy.Nox_xxx_gameGet_4DB1B0 = nox_xxx_gameGet_4DB1B0
	legacy.Sub_4DCC90 = sub_4DCC90
	legacy.Sub_4DB1C0 = sub_4DB1C0
	legacy.Sub_4DCBF0 = sub_4DCBF0
	legacy.Nox_xxx_serverIsClosing_446180 = nox_xxx_serverIsClosing_446180
	legacy.Sub_4DCC10 = sub_4DCC10
	legacy.Sub_4DCFB0 = sub_4DCFB0
	legacy.Sub_4DD0B0 = sub_4DD0B0
	legacy.InputSetKeyTimeoutLegacy = inputSetKeyTimeoutLegacy
	legacy.InputKeyCheckTimeoutLegacy = inputKeyCheckTimeoutLegacy
	legacy.Sub_416120 = sub_416120
	legacy.Sub_416170 = sub_416170
	legacy.Sub_416150 = sub_416150
	legacy.Nox_xxx_keybind_nameByTitle_42E960 = nox_xxx_keybind_nameByTitle_42E960
	legacy.Nox_xxx_bindevent_bindNameByTitle_42EA40 = nox_xxx_bindevent_bindNameByTitle_42EA40
	legacy.Sub_4C3B70 = sub_4C3B70
	legacy.Sub_4CBBF0 = sub_4CBBF0
	legacy.Nox_input_reset_430140 = nox_input_reset_430140
	legacy.Nox_xxx_playerDisconnByPlrID_4DEB00 = func(id ntype.PlayerInd) {
		noxServer.PlayerDisconnectByIndCode4(id)
	}
	legacy.Nox_xxx_playerCallDisconnect_4DEAB0 = func(ind ntype.PlayerInd, v int8) {
		noxServer.PlayerDisconnectByInd(ind, v)
	}
	legacy.Nox_xxx_playerCameraUnlock_4E6040 = nox_xxx_playerCameraUnlock_4E6040
	legacy.Nox_xxx_playerCameraFollow_4E6060 = nox_xxx_playerCameraFollow_4E6060
	legacy.Nox_xxx_playerGetPossess_4DDF30 = nox_xxx_playerGetPossess_4DDF30
	legacy.Nox_xxx_playerGoObserver_4E6860 = nox_xxx_playerGoObserver_4E6860
	legacy.Nox_xxx_playerObserveClear_4DDEF0 = nox_xxx_playerObserveClear_4DDEF0
	legacy.Nox_xxx_playerObserveMonster_4DDE80 = nox_xxx_playerObserveMonster_4DDE80
	legacy.Nox_xxx_gameSetAudioFadeoutMb_501AC0 = nox_xxx_gameSetAudioFadeoutMb_501AC0
	legacy.Nox_xxx_unitUpdateMonster_50A5C0 = nox_xxx_unitUpdateMonster_50A5C0
	legacy.Sub_446380 = sub_446380
	legacy.Nox_xxx_cliUpdateCameraPos_435600 = nox_xxx_cliUpdateCameraPos_435600
	legacy.Sub_437260 = sub_437260
	legacy.Get_nox_client_texturedFloors_154956 = get_nox_client_texturedFloors_154956
	legacy.Sub_480860 = sub_480860
	legacy.Sub_473970 = sub_473970
	legacy.Nox_client_isConnected = nox_client_isConnected
	legacy.Nox_xxx_pickupDefault_4F31E0 = nox_xxx_pickupDefault_4F31E0
	legacy.Nox_objectPickupAudEvent_4F3D50 = nox_objectPickupAudEvent_4F3D50
	legacy.Nox_xxx_pickupPotion_4F37D0 = nox_xxx_pickupPotion_4F37D0
	legacy.Nox_xxx_playerClassCanUseItem_57B3D0 = nox_xxx_playerClassCanUseItem_57B3D0
	legacy.Nox_xxx_useMushroom_53ECE0 = nox_xxx_useMushroom_53ECE0
	legacy.Nox_xxx_useCiderConfuse_53EF00 = nox_xxx_useCiderConfuse_53EF00
	legacy.Nox_xxx_useEnchant_53ED60 = nox_xxx_useEnchant_53ED60
	legacy.Nox_xxx_useCast_53ED90 = nox_xxx_useCast_53ED90
	legacy.Nox_xxx_useConsume_53EE10 = nox_xxx_useConsume_53EE10
	legacy.Nox_xxx_usePotion_53EF70 = nox_xxx_usePotion_53EF70
	legacy.Sub_57B370 = sub_57B370
	legacy.Sub_4A1BE0 = sub_4A1BE0
	legacy.Sub_41E300 = sub_41E300
	legacy.Sub_4A50A0 = sub_4A50A0
	legacy.Sub_4A50D0 = sub_4A50D0
	legacy.Nox_client_resetScreenParticles_431510 = nox_client_resetScreenParticles_431510
	legacy.Sub_46D6F0 = sub_46D6F0
	legacy.Sub_413A00 = sub_413A00
	legacy.Sub_44A400 = sub_44A400
	legacy.Nox_game_showSelChar_4A4DB0 = nox_game_showSelChar_4A4DB0
	legacy.Nox_savegame_sub_46D580 = nox_savegame_sub_46D580
	legacy.Sub_450580 = sub_450580
	legacy.Sub_4DB170 = sub_4DB170
	legacy.Nox_setSaveFileName_4DB130 = setSaveFileName
	legacy.Nox_xxx_harpoonBreakForPlr_537520 = nox_xxx_harpoonBreakForPlr_537520
	legacy.Nox_xxx_collideHarpoon_4EB6A0 = nox_xxx_collideHarpoon_4EB6A0
	legacy.Nox_xxx_updateHarpoon_54F380 = nox_xxx_updateHarpoon_54F380
	legacy.SendXXX_5550D0 = sendXXX_5550D0
	legacy.Nox_xxx_netStatsMultiplier_4D9C20 = func(a1p *server.Object) int {
		return noxServer.NetStatsMultiplier(a1p)
	}
	legacy.Sub_554240 = sub_554240
	legacy.Nox_xxx_netOnPacketRecvCli_48EA70 = nox_xxx_netOnPacketRecvCli_48EA70
	legacy.Sub_43C6E0 = sub_43C6E0
	legacy.Sub_43CF40 = sub_43CF40
	legacy.Sub_43CF70 = sub_43CF70
	legacy.Nox_objectDropAudEvent_4EE2F0 = nox_objectDropAudEvent_4EE2F0
	legacy.Nox_xxx_unitMonsterInit_4F0040 = objectMonsterInit
	legacy.Nox_xxx_checkSummonedCreaturesLimit_500D70 = nox_xxx_checkSummonedCreaturesLimit_500D70
	legacy.Nox_xxx_unitDoSummonAt_5016C0 = nox_xxx_unitDoSummonAt_5016C0
	legacy.Nox_xxx___mkgmtime_538280 = nox_xxx___mkgmtime_538280
	legacy.Sub_4E71F0 = sub_4E71F0
	legacy.Nox_bomberDead_54A150 = nox_bomberDead_54A150
	legacy.Nox_xxx_dieGlyph_54DF30 = nox_xxx_dieGlyph_54DF30
	legacy.Nox_xxx_collideGlyph_4E9A00 = nox_xxx_collideGlyph_4E9A00
	legacy.Sub_41CC00 = sub_41CC00
	legacy.Nox_xxx_playerSendMOTD_4DD140 = nox_xxx_playerSendMOTD_4DD140
	legacy.Sub_526CA0 = sub_526CA0
	legacy.Nox_xxx_playerSetState_4FA020 = nox_xxx_playerSetState_4FA020
	legacy.Nox_xxx_mapSetWallInGlobalDir0pr1_5004D0 = func() {
		noxServer.spells.walls.switchToPrevDir()
	}
	legacy.Nox_xxx_map_5004F0 = func() {
		noxServer.spells.walls.switchToNewDir()
	}
	legacy.Sub_4FF990 = func(a1 uint32) {
		noxServer.spells.walls.disablePlayerBits(a1)
	}
	legacy.Sub_5000B0 = func(a1 *server.Object) int {
		return noxServer.spells.walls.LoadOrSave(cryptfile.Global(), a1)
	}
	legacy.Nox_xxx_monsterCreateFn_54C480 = nox_xxx_monsterCreateFn_54C480
	legacy.Sub_4A1A40 = sub_4A1A40
	legacy.Nox_client_getChatMap_49FF40 = nox_client_getChatMap_49FF40
	legacy.Nox_xxx_soundPlayerDamageSound_5328B0 = Nox_xxx_soundPlayerDamageSound_5328B0
	legacy.Nox_xxx_soundDefaultDamageSound_532E20 = server.Nox_xxx_soundDefaultDamageSound_532E20
	legacy.Nox_thing_monster_draw = func(vp *noxrender.Viewport, dr *client.Drawable) int {
		return noxClient.DrawMonster(vp, dr)
	}
	legacy.Nox_thing_vector_animate_draw = func(vp *noxrender.Viewport, dr *client.Drawable) int {
		return noxClient.DrawAnimVector(vp, dr)
	}
	legacy.Nox_thing_animate_state_draw = func(vp *noxrender.Viewport, dr *client.Drawable) int {
		return noxClient.DrawAnimState(vp, dr)
	}
	legacy.Nox_thing_player_draw = func(vp *noxrender.Viewport, dr *client.Drawable) int {
		return noxClient.DrawPlayer(vp, dr)
	}
	legacy.Nox_thing_npc_draw = func(vp *noxrender.Viewport, dr *client.Drawable) int {
		return noxClient.DrawNPC(vp, dr)
	}
	legacy.Nox_xxx_mapSwitchLevel_4D12E0 = func(a1 bool) {
		noxServer.nox_xxx_mapSwitchLevel_4D12E0(a1)
	}
	legacy.NewButtonOrCheckbox = func(parent *gui.Window, status gui.StatusFlags, px, py, w, h int, draw *gui.WindowData) *gui.Window {
		return gui.NewButtonOrCheckboxRaw(noxClient.GUI, parent, status, px, py, w, h, draw)
	}
}
