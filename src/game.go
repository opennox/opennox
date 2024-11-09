package opennox

import (
	"context"
	"errors"
	"fmt"
	"image"
	"math"
	"path/filepath"
	"strconv"
	"strings"
	"unsafe"

	"github.com/noxworld-dev/lobby"
	"github.com/noxworld-dev/xwis"

	"github.com/noxworld-dev/opennox-lib/common"
	"github.com/noxworld-dev/opennox-lib/console"
	"github.com/noxworld-dev/opennox-lib/env"
	"github.com/noxworld-dev/opennox-lib/ifs"
	"github.com/noxworld-dev/opennox-lib/log"
	"github.com/noxworld-dev/opennox-lib/object"
	"github.com/noxworld-dev/opennox-lib/script"
	"github.com/noxworld-dev/opennox-lib/types"

	"github.com/noxworld-dev/opennox/v1/client"
	"github.com/noxworld-dev/opennox/v1/client/gui"
	"github.com/noxworld-dev/opennox/v1/common/config"
	"github.com/noxworld-dev/opennox/v1/common/discover"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"github.com/noxworld-dev/opennox/v1/common/ntype"
	"github.com/noxworld-dev/opennox/v1/common/unit/ai"
	"github.com/noxworld-dev/opennox/v1/internal/netlist"
	"github.com/noxworld-dev/opennox/v1/internal/netstr"
	"github.com/noxworld-dev/opennox/v1/internal/version"
	"github.com/noxworld-dev/opennox/v1/legacy"
	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc"
	"github.com/noxworld-dev/opennox/v1/server"
)

var (
	useXWIS                   = true
	gameLog                   = log.New("game")
	nox_game_playState_811372 int
	movieFilesStackCur        int
	movieFilesStack           [2]string
	dword_587000_311372       = -1
	dword_5d4594_2516476      byte
	gameIsNotMultiplayer      bool
	gameIsSwitchToSolo        bool
	defaultChatMap            string
)

func init() {
	configBoolPtr(config.KeyNetRegister, "NOX_XWIS", true, &useXWIS)
	configStrPtr(config.KeyDefChatMap, "NOX_CHAT_MAP", "so_lod", &defaultChatMap)
	gui.RegisterState(client.StateMovies, "Movies", nox_game_rollLogoAndStart_4AB1F0)
	gui.RegisterState(client.StateMainMenu, "MainMenu", nox_game_showMainMenu4A1C00)
	gui.RegisterState(client.StateCharSelect, "CharSelect", func() bool {
		legacy.Sub_4A7A70(1)
		return nox_game_showSelChar4A4DB0()
	})
	noxCmdShow.Register(&console.Command{
		Token:  "game",
		HelpID: "showgamehelp",
		Flags:  console.ClientServer,
		Func: func(ctx context.Context, c *console.Console, tokens []string) bool {
			if len(tokens) != 0 {
				return false
			}
			s := noxServer
			players := s.Players.Count()
			if noxflags.HasEngine(noxflags.EngineNoRendering) {
				players--
			}
			subVer := ""
			if noxHighRes {
				subVer = " HD"
			}
			c.Printf(console.ColorRed, "OpenNox%s %s", subVer, version.ClientVersion())
			if noxflags.HasGame(noxflags.GameOnline) {
				sname := s.getServerName()
				str := c.Strings().GetStringInFile("Name", "parsecmd.c")
				c.Printf(console.ColorRed, "%s %s", str, sname)

				v16 := legacy.Nox_xxx_guiServerOptionsGetGametypeName_4573C0(noxflags.GetGame())
				str = c.Strings().GetStringInFile("Type", "parsecmd.c")
				c.Printf(console.ColorRed, "%s %s", str, v16)

				p := getCurrentSettings2()
				ind := p.Field52
				timeLimit := legacy.Sub_40A180(noxflags.GameFlag(ind))
				lessons := legacy.Nox_xxx_servGamedataGet_40A020(ind)
				maxPlayers := s.getServerMaxPlayers()
				mname := s.nox_server_currentMapGetFilename_409B30()
				format := c.Strings().GetStringInFile("GameInfo", "parsecmd.c")
				format = strings.ReplaceAll(format, "%S", "%s")
				c.Printf(console.ColorRed, format, mname, players, maxPlayers, lessons, timeLimit)
				c.Print(console.ColorRed, s.GetExtIP(s.NetStr.Host()).String())
			} else {
				maxPlayers := s.getServerMaxPlayers()
				mname := s.nox_server_currentMapGetFilename_409B30()
				format := c.Strings().GetStringInFile("GameInfo", "parsecmd.c")
				format = strings.ReplaceAll(format, "%S", "%s")
				c.Printf(console.ColorRed, format, mname, players, maxPlayers, 0, 0)
			}
			return true
		},
	})
}

func gameGetPlayState() int {
	return nox_game_playState_811372
}

func gameSetPlayState(st int) {
	gameLog.Println("play state:", st)
	nox_game_playState_811372 = st
}

func nox_game_setMovieFile_4CB230(name string) (string, bool) {
	filename := filepath.Join("movies", name)
	_, err := ifs.Stat(filename)
	if err != nil {
		gameLog.Printf("movie not found: %q", filename)
		return "", false
	}
	gameLog.Printf("movie found: %q", filename)
	return filename, true
}

func pushMovieFile(path string) {
	if movieFilesStackCur < 2 {
		movieFilesStack[movieFilesStackCur] = path
		movieFilesStackCur++
	}
}

func sub_578E00() {
	dword_587000_311372 = -1
}

func sub_578DF0() byte {
	return dword_5d4594_2516476
}

func sub_578DE0(v byte) {
	dword_5d4594_2516476 = v
}
func nox_xxx_GetEndgameDialog() string {
	if dword_587000_311372 == 0 {
		if dword_5d4594_2516476&2 == 0 {
			return "gmcap14e.wav"
		}
		if dword_5d4594_2516476&4 == 0 {
			return "gmcap13e.wav"
		}
	} else if dword_587000_311372 == 1 {
		if dword_5d4594_2516476&1 == 0 {
			return "gmcap17e.wav"
		}
		if dword_5d4594_2516476&4 == 0 {
			return "gmcap18e.wav"
		}
	} else if dword_587000_311372 == 2 {
		if dword_5d4594_2516476&2 == 0 {
			return "gmcap16e.wav"
		}
		if dword_5d4594_2516476&1 == 0 {
			return "gmcap15e.wav"
		}
	}
	return ""
}

func sub_578CD0() {
	if dword_587000_311372 != -1 {
		names := []string{
			"warrior", "wizard", "conjurer",
		}
		name := names[dword_587000_311372] + ".vqa"
		if path, ok := nox_game_setMovieFile_4CB230(name); ok {
			pushMovieFile(path)
			sub4B0640(sub_578C60)
			if err := noxClient.DrawGeneral(true); err != nil {
				videoLog.Println(err)
			}
		} else {
			sub_578C60()
		}
	}
}

func sub_578C60() {
	if legacy.Sub_44E560() != nil {
		legacy.Nox_client_lockScreenBriefing_450160(255, 1, 0)
		sub_4A2530()
	}
}

func nox_game_rollLogoAndStart_4AB1F0() bool {
	if isServer || isServerQuest {
		// FIXME: switch to server state directly
		return startServer()
	}
	noxClient.GamePopState()
	if noxflags.HasGame(noxflags.GameFlag26) {
		nox_game_rollIntroAndStart_4AB170()
		return true
	}
	path, ok := nox_game_setMovieFile_4CB230("wwlogo.vqa")
	if !ok {
		nox_game_rollIntroAndStart_4AB170()
		return true
	}
	pushMovieFile(path)
	sub4B0640(nox_game_rollIntroAndStart_4AB170)
	if err := noxClient.DrawGeneral(false); err != nil {
		videoLog.Println(err)
	}
	return true
}

func nox_game_rollIntroAndStart_4AB170() {
	if sub_578DF0()&0x80 != 0 || noxflags.HasGame(noxflags.GameFlag26) {
		nox_game_rollNoxLogoAndStart_4AB0F0()
		return
	}
	path, ok := nox_game_setMovieFile_4CB230("intro.vqa")
	if !ok {
		nox_game_rollNoxLogoAndStart_4AB0F0()
		return
	}
	pushMovieFile(path)
	sub4B0640(nox_game_rollNoxLogoAndStart_4AB0F0)
	if err := noxClient.DrawGeneral(true); err != nil {
		videoLog.Println(err)
	}
	sub_578DE0(sub_578DF0() | 0x80)
}

func nox_game_rollNoxLogoAndStart_4AB0F0() {
	if noxflags.HasGame(noxflags.GameFlag26) {
		nox_game_showLegal_4CC4E0()
		return
	}
	path, ok := nox_game_setMovieFile_4CB230("noxlogo.vqa")
	if !ok {
		nox_game_showLegal_4CC4E0()
		return
	}
	pushMovieFile(path)
	sub4B0640(nox_game_showLegal_4CC4E0)
	if err := noxClient.DrawGeneral(false); err != nil {
		videoLog.Println(err)
	}
}

func nox_xxx_parseGamedataBinPre_4D1630() int {
	s := noxServer
	if err := s.Balance.Read(); err != nil {
		return 0
	}
	s.abilities.reloadGamedata()
	s.CalcClassStats()
	return bool2int(legacy.Nox_xxx_loadMonsterBin_517010() != 0)
}

func startServer() bool {
	legacy.Set_nox_game_createOrJoin_815048(1)
	noxflags.SetEngine(noxflags.EngineAdmin)
	noxflags.UnsetEngine(noxflags.EngineGodMode)
	noxflags.SetGame(noxflags.GameOnline)
	if !isServerQuest {
		noxflags.SetGame(noxflags.GameNotQuest)
	} else {
		noxflags.UnsetGame(noxflags.GameNotQuest)
	}
	noxflags.UnsetGame(noxflags.GameModeCoop)
	legacy.Sub_461440(0)
	sub4D6F40(false)
	sub_4D6F90(0)
	if !isServerQuest {
		noxServer.nox_xxx_setQuest_4D6F60(0)
	} else {
		noxServer.nox_xxx_setQuest_4D6F60(1)
		legacy.Nox_xxx_cliShowHideTubes_470AA0(1)
	}
	legacy.Sub_4D6F80(0)
	if sub_473670() == 0 {
		nox_client_toggleMap_473610()
	}
	if !isServerQuest {
		legacy.Nox_xxx_cliShowHideTubes_470AA0(0)
	}
	legacy.Nox_xxx_cliSetMinimapZoom_472520(2300)
	legacy.Sub_43AF50()
	if nox_xxx_parseGamedataBinPre_4D1630() == 0 {
		nox_xxx_setContinueMenuOrHost_43DDD0(0)
		nox_client_gui_flag_815132 = 0
		return false
	}
	nox_xxx_serverHost_43B4D0()
	return true
}

func (s *Server) getServerName() string {
	return legacy.Nox_xxx_serverOptionsGetServername_40A4C0()
}

func (s *Server) getServerMap() string {
	return legacy.Nox_xxx_mapGetMapName_409B40()
}

func (s *Server) getServerMaxPlayers() int {
	return legacy.Nox_xxx_servGetPlrLimit_409FA0()
}

func (s *Server) nox_xxx_isQuest_4D6F50() bool {
	return memmap.Uint32(0x5D4594, 1556160) != 0
}

func (s *Server) nox_xxx_setQuest_4D6F60(v int) {
	*memmap.PtrUint32(0x5D4594, 1556160) = uint32(v)
}

func sub_4D6F90(a1 uint32) {
	*memmap.PtrUint32(0x5D4594, 1556104) = a1
}

type srvReg struct {
	cur xwisInfoShort
	srv *discover.RegServer
}

func nox_xxx_serverHost_43B4D0() {
	isHost := legacy.Get_nox_game_createOrJoin_815048() != 0
	gameLog.Println("host:", isHost)
	if isHost {
		// host
		legacy.Nox_client_xxx_switchChatMap_43B510()
		legacy.Nox_client_guiXxx_43A9D0()
		nox_client_guiXxxDestroy_4A24A0()
		noxServer.ai.nox_xxx_gameSetAudioFadeoutMb(0)
	} else {
		// join
		legacy.Sub_43B630()
		nox_client_createSockAndJoin_43B440()
	}
	legacy.Sub_49FF20()
	if isHost {
		noxServer.maybeRegisterGameOnline()
	} else {
		noxServer.maybeStopRegister()
	}
}

type xwisInfoShort struct {
	Name       string
	Map        string
	Port       int
	Flags      noxflags.GameFlag
	Access     lobby.GameAccess
	Players    []lobby.PlayerInfo
	MaxPlayers int
	Level      int
	Res        image.Point
}

func (v *xwisInfoShort) Equal(v2 *xwisInfoShort) bool {
	return v.Name == v2.Name && v.Map == v2.Map &&
		v.Port == v2.Port && v.Flags == v2.Flags &&
		v.Access == v2.Access && len(v.Players) == len(v2.Players) &&
		v.MaxPlayers == v2.MaxPlayers && v.Level == v2.Level &&
		v.Res == v2.Res
}

func (s *Server) getGameInfo() xwisInfoShort {
	access := lobby.AccessOpen
	if sst := getServerSettings(); sst.Flags100.Has(server.ServerPrivate) {
		access = lobby.AccessPassword
	} else if sst.Flags100.Has(server.ServerClosed) {
		access = lobby.AccessClosed
	}
	players := s.Players.List()
	list := make([]lobby.PlayerInfo, 0, len(players))
	for _, p := range players {
		list = append(list, lobby.PlayerInfo{
			Name:  p.Name(),
			Class: p.PlayerClass().String(),
		})
	}
	res := videoGetMaxSize()
	if !isDedicatedServer {
		res = noxClient.videoGetGameMode()
	}
	return xwisInfoShort{
		Port:       s.ServerPort(),
		Name:       s.getServerName(),
		Map:        s.getServerMap(),
		Flags:      noxflags.GetGame(),
		Players:    list,
		MaxPlayers: s.getServerMaxPlayers(),
		Level:      s.nox_game_getQuestStage_4E3CC0(),
		Access:     access,
		Res:        res,
	}
}

func xwisIsQuest(info *xwis.GameInfo) bool {
	flags := int(info.Flags) | int(info.MapType)
	return flags&int(xwis.MapTypeQuest) != 0
}

func gameFlagsToMode(f noxflags.GameFlag) lobby.GameMode {
	switch {
	case f.Has(noxflags.GameModeKOTR):
		return lobby.ModeKOTR
	case f.Has(noxflags.GameModeCTF):
		return lobby.ModeCTF
	case f.Has(noxflags.GameModeFlagBall):
		return lobby.ModeFlagBall
	case f.Has(noxflags.GameModeChat):
		return lobby.ModeChat
	case f.Has(noxflags.GameModeArena):
		return lobby.ModeArena
	case f.Has(noxflags.GameModeElimination):
		return lobby.ModeElimination
	case f.Has(noxflags.GameModeQuest):
		return lobby.ModeQuest
	case f.Has(noxflags.GameModeCoop):
		return lobby.ModeCoop
	}
	return lobby.ModeCustom
}

func gameAccessToXWIS(v lobby.GameAccess) xwis.Access {
	switch v {
	case lobby.AccessOpen:
		return xwis.AccessOpen
	case lobby.AccessClosed:
		return xwis.AccessClosed
	case lobby.AccessPassword:
		return xwis.AccessPrivate
	}
	return xwis.AccessOpen
}

func (v xwisInfoShort) GameInfo() discover.GameInfo {
	info := discover.GameInfo{
		Game: lobby.Game{
			Name: v.Name,
			Port: v.Port,
			Map:  v.Map,
			Res: lobby.Resolution{
				Width:   v.Res.X,
				Height:  v.Res.Y,
				HighRes: noxHighRes,
			},
			Mode: gameFlagsToMode(v.Flags),
			Vers: version.Version(),
			Players: lobby.PlayersInfo{
				Cur:  len(v.Players),
				Max:  v.MaxPlayers,
				List: v.Players,
			},
		},
		XWIS: xwis.GameInfo{
			Name:       v.Name,
			Map:        v.Map,
			Resolution: xwis.Res1024x768,
			Flags:      xwis.GameFlags(v.Flags),
			Access:     gameAccessToXWIS(v.Access),
			Players:    len(v.Players),
			MaxPlayers: v.MaxPlayers,
		},
	}
	if v.Flags.Has(noxflags.GameModeQuest) {
		info.Game.Quest = &lobby.QuestInfo{Stage: v.Level}
	}
	if xwisIsQuest(&info.XWIS) {
		info.XWIS.FragLimit = v.Level
	}
	return info
}

func (s *Server) maybeRegisterGameOnline() {
	if !useXWIS || !noxflags.HasGame(noxflags.GameOnline) || env.IsE2E() || !(isDedicatedServer || s.Announce) {
		s.maybeStopRegister()
		return
	}
	info := s.getGameInfo()
	if s.srvReg.srv == nil {
		s.srvReg.cur = info
		s.srvReg.srv = discover.NewServer(info.GameInfo())
		return
	}
	if s.srvReg.cur.Equal(&info) {
		return
	}
	s.srvReg.cur = info
	s.srvReg.srv.Update(info.GameInfo())
}

func (s *Server) maybeStopRegister() {
	if s.srvReg.srv != nil {
		s.srvReg.srv.Close()
		s.srvReg.srv = nil
	}
}

func getString10984() string {
	return alloc.GoString((*byte)(memmap.PtrOff(0x85B3FC, 10984)))
}

func initGameSession435CC0() error {
	c := noxClient
	ctx := context.Background()
	legacy.Sub_445450()
	legacy.Sub_45DB90()
	legacy.Sub_41D1A0(0)
	legacy.Nox_xxx_initTime_435570()
	legacy.Set_nox_client_gui_flag_1556112(0)

	c.Objs.Init(5000)

	if legacy.Nox_xxx_allocArrayHealthChanges_49A5F0() == 0 {
		return errors.New("nox_xxx_allocArrayHealthChanges_49A5F0 failed")
	}

	if err := nox_xxx_parseThingBinClient_44C840_read_things(); err != nil {
		return err
	}

	if legacy.Nox_xxx_loadGuides_427070() == 0 {
		return errors.New("nox_xxx_loadGuides_427070 failed")
	}

	if legacy.Sub_494F00() == 0 {
		return errors.New("sub_494F00 failed")
	}

	if !isDedicatedServer {
		if err := nox_game_guiInit_473680(); err != nil {
			return fmt.Errorf("game gui Init failed: %w", err)
		}
	}
	c.Server.NPCs.Init()
	if legacy.Nox_xxx_loadReflSheild_499360() == 0 {
		return errors.New("nox_xxx_loadReflSheild_499360 failed")
	}

	c.nox_xxx_initSight_485F80()
	c.Sight.Init(nox_win_width, nox_win_height)

	c.InitDrawableLists()

	if nox_xxx_allocArrayDrawableFX_495AB0() == 0 {
		return errors.New("nox_xxx_allocArrayDrawableFX_495AB0 failed")
	}

	if legacy.Nox_xxx_allocClassListFriends_495980() == 0 {
		return errors.New("nox_xxx_allocClassListFriends_495980 failed")
	}
	legacy.Sub_4958F0()
	nox_xxx_gameSetCliConnected(false)
	noxflags.SetGame(noxflags.GameFlag24)
	if noxflags.HasGame(noxflags.GameHost) {
		if !isDedicatedServer {
			legacy.Nox_xxx_netPlayerIncomingServ_4DDF60(server.HostPlayerIndex)
		}
	} else {
		c.Nox_xxx_netSendIncomingClient_43CB00()
	}
	c.SetDrawFunc(c.clientDraw)
	gameSetPlayState(3)
	c.Objs.LoadError = false
	sz := videoGetWindowSize()
	vp := c.Viewport()
	vp.Screen = image.Rect(0, 0, sz.X-1, sz.Y-1)
	vp.Size = sz
	vp.Field10 = 0
	vp.Field11 = 0
	vp.Jiggle12 = 0
	v1 := nox_video_getCutSize()
	c.nox_draw_setCutSize(v1, 0)
	if noxflags.HasGame(noxflags.GameModeCoop) {
		sub_41CC00(getString10984())
	} else if c.srv.nox_xxx_isQuest_4D6F50() || sub_4D6F70() {
		if c.srv.nox_xxx_isQuest_4D6F50() || sub_4D6F70() {
			legacy.Sub_460380()
			legacy.Nox_xxx_cliPrepareGameplay1_460E60()
			legacy.Nox_xxx_cliPrepareGameplay2_4721D0()
		}
		if !noxflags.HasGame(noxflags.GameHost) {
			sub_41CC00(getString10984())
		}
	}
	if !isServer {
		legacy.Nox_xxx_plrLoad_41A480(getString10984())
	}
	if isServer && !isDedicatedServer {
		c.srv.PlayerGoObserver(c.Server.Players.List()[0], false, true)
	}
	execConsoleCmdAuthed(ctx, "execrul autoexec.rul")
	if isServer {
		execServerCmd("set cycle on")
		for _, cmd := range serverExec {
			if len(cmd) == 0 {
				continue
			}
			execServerCmd(cmd)
		}
	}
	legacy.Sub_4951C0()
	legacy.Sub_465DE0(0)
	return nil
}

func sub_41CC00(a1 string) {
	sz := 4 + nox_xxx_computeServerPlayerDataBufferSize_41CC50(a1)
	data, _ := alloc.Make([]byte{}, sz)
	legacy.Sub_41CAC0(a1, data)
	noxServer.NetXferSendHost(server.NetXferSaveServer, server.NetXferSaveServerType, data, netXferSendDone, netXferSendAborted)
}

func sub_4464D0(a1 int32) []byte {
	sz := int(memmap.Uint32(0x5D4594, 826040+4*uintptr(a1)))
	return unsafe.Slice(legacy.Get_dword_5d4594_826036(), sz)
}

func nox_xxx_playerSendMOTD_4DD140(ind ntype.PlayerInd) {
	buf := sub_4464D0(1)
	if len(buf) != 0 {
		noxServer.NetXferSend(ind, server.NetXferMOTD, server.NetXferMOTDType, buf, netXferSendDone, netXferSendAborted)
	}
}

func execConsoleCmd(ctx context.Context, cmd string) bool { // nox_server_parseCmdText_443C80
	cmd = strings.TrimSpace(cmd)
	if len(cmd) == 0 {
		return false
	}
	if consoleWaitSysOpPass {
		nox_console_sendSysOpPass_4409D0(cmd)
		consoleWaitSysOpPass = false
		return true
	}
	return execConsoleCmdAuthed(ctx, cmd)
}

func execConsoleCmdAuthed(ctx context.Context, cmd string) bool {
	cmd = strings.TrimSpace(cmd)
	if len(cmd) == 0 {
		return false
	}
	if noxflags.HasGame(noxflags.GameHost) {
		ctx = console.AsServer(ctx)
	} else {
		ctx = console.AsClient(ctx)
	}
	if noxflags.HasEngine(noxflags.EngineNoRendering) {
		ctx = console.AsDedicated(ctx)
	}
	return noxConsole.Exec(ctx, cmd)
}

func execServerCmd(cmd string) {
	cmd = strings.TrimSpace(cmd)
	if len(cmd) == 0 {
		return
	}
	ctx := context.Background()
	ctx = console.AsServer(ctx)
	if noxflags.HasEngine(noxflags.EngineNoRendering) {
		ctx = console.AsDedicated(ctx)
	}
	ctx = console.WithCheats(ctx)
	noxConsole.Exec(ctx, cmd)
}

func serverCmdLoadMap(name string) {
	if len(name) == 0 {
		return
	}
	execServerCmd("load " + name)
}

func (s *Server) nox_server_currentMapGetFilename_409B30() string {
	return memmap.String(0x5D4594, 2598188)
}

func nox_xxx_mapFilenameGetSolo_4DB260() string {
	return dword_5d4594_1559960
}

func nox_client_getChatMap_49FF40() string {
	return defaultChatMap
}

func (s *Server) nox_xxx_servInitialMapLoad_4D17F0() bool {
	legacy.Sub_4E79B0(0)
	if s.nox_server_currentMapGetFilename_409B30() == "" {
		s.nox_xxx_gameSetMapPath_409D70("tutorial.map")
	}
	s.MapSend.Reset()
	if err := s.nox_xxx_mapExitAndCheckNext_4D1860_server(); err != nil {
		gameLog.Println(err)
		return false
	}
	if debugMainloop {
		log.Println("gameStateFunc = nox_xxx_gameTick_4D2580_server")
	}
	s.SetUpdateFunc(s.nox_xxx_gameTick_4D2580_server)
	s.NetList.ResetAllInd(netlist.Kind1)
	noxflags.SetGame(noxflags.GameFlag18)
	legacy.Nox_xxx_netGameSettings_4DEF00()
	legacy.Nox_server_gameUnsetMapLoad_40A690()
	return true
}

func (s *Server) nox_xxx_gameTick_4D2580_server() bool {
	defer s.RunTickHooks()
	ticks := platformTicks()
	v2 := false
	if legacy.Get_dword_5d4594_2650652() == 0 {
		s.NetList.ResetAllInd(netlist.Kind1)
	} else {
		v4 := nox_xxx_rateGet_40A6C0()
		if legacy.Sub_416650() != 0 && legacy.Sub_41E2F0() == 8 {
			v2 = true
		}
		if v4 == 1 || noxflags.HasGame(noxflags.GameFlag4) || s.Frame()%uint32(v4) == 1 {
			s.NetList.ResetAllInd(netlist.Kind1)
		}
	}
	s.Audio.Reset()
	s.NetStr.Update()
	if noxflags.HasEngine(noxflags.EngineReplayRead) {
		s.nox_xxx_replayTickMB_4D3580_net_playback(true)
	}
	if noxflags.HasEngine(noxflags.EngineLogBand) {
		noxPerfmon.LogBandwidth(s.Server, s.NetStr)
	}
	if noxflags.HasGame(noxflags.GameFlag4) {
		nox_xxx_gameTick_4D2580_server_A1()
		s.nox_xxx_gameTick_4D2580_server_A2(v2)
	} else if legacy.Get_dword_5d4594_1548524() == 0 {
		if !s.nox_xxx_gameTick_4D2580_server_B(ticks) {
			return true
		}
	} else if true || !v2 {
		if !s.nox_xxx_gameTick_4D2580_server_C() {
			return false
		}
	}
	s.nox_xxx_gameTick_4D2580_server_E()
	return true
}

func nox_xxx_gameTick_4D2580_server_A1() {
	if memmap.Uint64(0x5D4594, 1548676) == 0 {
		*memmap.PtrUint64(0x5D4594, 1548676) = platformTicks() + 10000
		legacy.Nox_xxx_guiServerOptionsHide_4597E0(0)
		netstr.Log.Println("EndGame")
		legacy.Sub_46DCC0()
	}
}

func (s *Server) nox_xxx_gameTick_4D2580_server_A2(v2 bool) {
	s.Nox_server_netMaybeSendInitialPackets_4DEB30()
	s.nox_xxx_netlist_4DEB50()
	if platformTicks() <= memmap.Uint64(0x5D4594, 1548676) || v2 {
		return
	}
	*memmap.PtrUint32(0x5D4594, 1548676) = 0
	*memmap.PtrUint32(0x5D4594, 1548680) = 0
	legacy.Set_dword_5d4594_1548524(1)
	sub_416170(12)
	noxflags.UnsetGame(noxflags.GameFlag4)
	for _, u := range s.Players.ListUnits() {
		asObjectS(u).dropAllItems()
		legacy.Nox_xxx_playerMakeDefItems_4EF7D0(u, 1, 0)
	}
	s.ObjectsAddPending()
	if noxflags.HasGame(noxflags.GameModeCoopTeam) {
		return
	}
	v7 := getSettings2ByInd(0)
	if v7.Field57 != 0 {
		v8 := legacy.Sub_409B80() + ".map"
		s.nox_xxx_gameSetMapPath_409D70(v8)
		v7.Field57 = 0
	} else if legacy.Sub_4D0D70() != 0 {
		legacy.Nox_xxx_loadMapCycle_4D0A30()
		v11 := legacy.Nox_xxx_getSomeMapName_4D0CF0()
		if v11 != "" {
			v12 := s.nox_server_currentMapGetFilename_409B30()
			if strings.ToLower(v11) == strings.ToLower(v12) {
				v11 = legacy.Nox_xxx_getSomeMapName_4D0CF0()
			}
			if v11 != "" {
				s.nox_xxx_gameSetMapPath_409D70(v11)
			}
		}
	}
}

var _ = [1]struct{}{}[20-unsafe.Sizeof(server.Settings3{})]

var _ = [1]struct{}{}[60-unsafe.Sizeof(server.Settings2{})] // FIXME: must be 58

var _ = [1]struct{}{}[168-unsafe.Sizeof(server.Settings{})]

func getCurrentSettings2Ind() int { // sub_416580
	return int(memmap.Int32(0x5D4594, 371688))
}
func setCurrentSettings2Ind(ind int) {
	*memmap.PtrInt32(0x5D4594, 371688) = int32(ind)
}
func getSettings2ByInd(ind int) *server.Settings2 { // nox_xxx_cliGamedataGet_416590
	return memmap.PtrT[server.Settings2](0x5D4594, 371380+uintptr(ind)*58)
}
func getCurrentSettings2() *server.Settings2 { // sub_4165B0
	ind := getCurrentSettings2Ind()
	return getSettings2ByInd(ind)
}
func sub_4165D0(ind int) *server.Settings2 {
	setCurrentSettings2Ind(ind)
	return getSettings2ByInd(ind)
}

func nox_xxx_gameSetNoMPFlag_4DB230(a1 int) {
	gameIsNotMultiplayer = a1 != 0
}

func nox_xxx_gameIsNotMultiplayer_4DB250() bool {
	return gameIsNotMultiplayer
}

func nox_xxx_gameSetSwitchSolo_4DB220(a1 int) {
	gameIsSwitchToSolo = a1 != 0
}

func nox_xxx_gameIsSwitchToSolo_4DB240() bool {
	return gameIsSwitchToSolo
}

func (s *Server) nox_xxx_gameTick_4D2580_server_D() {
	pl := s.Players.ByInd(server.HostPlayerIndex)
	if pl == nil {
		return
	}
	u := pl.PlayerUnit
	if u == nil || sub_4DCC10(u) != 1 {
		return
	}
	savedone := false
	dead := u.Flags().Has(object.FlagDead)
	if !dead {
		s.scriptOnEvent(script.EventMapShutdown)
		noxflags.SetGame(noxflags.GameFlag28)
		savedone = saveCoopGame(saveName1557900)
		noxflags.UnsetGame(noxflags.GameFlag28)
		if !savedone && noxflags.HasGame(noxflags.GameClient) {
			v35 := strMan.GetStringInFile("GUISave.c:SaveErrorTitle", "C:\\NoxPost\\src\\Server\\System\\server.c")
			NewDialogWindow(nil, v35, v35, gui.DialogOKButton|gui.DialogFlag6, nil, nil)
		}
	}
	v28 := sub_4DB1C0()
	if dead || !savedone {
		if v28 != nil && !savedone {
			asObjectS(u).SetPos(legacy.AsPointf(unsafe.Pointer(*(*uintptr)(unsafe.Add(v28, 700)) + 80)))
		}
	} else if v28 != nil {
		v30 := alloc.GoString(*(**byte)(unsafe.Add(v28, 700)))
		v31, err := nox_client_checkSaveMapExistsTmp(v30)
		if err == nil && v31 != "" {
			nox_xxx_gameSetSwitchSolo_4DB220(1)
			nox_xxx_gameSetNoMPFlag_4DB230(1)
			nox_xxx_gameSetSoloSavePath_4DB270(v31)
		} else {
			gameLog.Println("check tmp map exists:", err)
		}
		s.SwitchMap(v30)
	}
	sub_4DB170(false, v28, 0)
}

func (s *Server) nox_xxx_gameTick_4D2580_server_C() bool {
	legacy.Sub_4EDD70()
	s.sub_417160()
	legacy.Sub_4573B0()
	if noxflags.HasGamePlay(noxflags.GameplayFlag2) && !noxflags.HasGame(noxflags.GameFlag15|noxflags.GameFlag16) &&
		noxflags.HasGamePlay(noxflags.GameplayFlag4) && !noxflags.HasGame(noxflags.GameModeChat) {
		legacy.Sub_4181F0(1)
	}
	if noxflags.HasGame(noxflags.GameModeQuest) && false && !noxflags.HasGame(noxflags.GameModeChat) {
		legacy.Sub_4264D0()
	}
	noxflags.SetGame(noxflags.GameFlag28)
	noxAudioServeT(500)
	err := s.nox_xxx_mapExitAndCheckNext_4D1860_server()
	noxAudioServe()
	noxflags.UnsetGame(noxflags.GameFlag28)
	mname := s.getServerMap()
	if err != nil {
		gameLog.Println(err)
		format := s.Strings().GetStringInFile("MapAccessError", "C:\\NoxPost\\src\\Server\\System\\server.c")
		legacy.PrintToPlayers(fmt.Sprintf(format, mname))
		//v36 := strMan.GetStringInFile("Error", "C:\\NoxPost\\src\\Server\\System\\server.c")
		//v15 := strMan.GetStringInFile("MapCorrupt", "C:\\NoxPost\\src\\Server\\System\\server.c")
		nox_xxx_setContinueMenuOrHost_43DDD0(0)
		return false
	}
	crc := nox_xxx_mapCrcGetMB_409B00()
	s.nox_xxx_netUseMap_4DEE00(mname+".map", crc)
	if false {
		legacy.Sub_416690()
		if noxflags.HasGame(noxflags.GameModeChat) {
			if noxflags.HasGame(noxflags.GameFlag16) {
				s.TeamsRemoveActive(true)
			}
			noxflags.UnsetGame(noxflags.GameFlag15 | noxflags.GameFlag16)
		} else {
			legacy.Sub_426060()
			legacy.Sub_41D6C0()
			inputSetKeyTimeoutLegacy(15)
		}
	}
	legacy.Nox_xxx_guiServerOptionsHide_4597E0(0)
	return true
}

func nox_game_guiInit_473680() error {
	*memmap.PtrPtr(0x5D4594, 1096420) = unsafe.Pointer(nox_xxx_gLoadImg("CursorBitmap").C())
	if legacy.Sub_455C30() == 0 {
		return errors.New("sub_455C30 failed")
	}
	if legacy.Sub_456070() == 0 {
		return errors.New("sub_456070 failed")
	}
	if noxClient.guiFPS.initGuiFPS(noxClient.Render(), noxClient.GUI) == 0 {
		return errors.New("sub_470710 failed")
	}
	if legacy.Nox_xxx_guiHealthManaInit_4714E0() == 0 {
		return errors.New("nox_xxx_guiHealthManaInit_4714E0 failed")
	}
	if legacy.Nox_xxx_bookInit_45B9D0() == 0 {
		return errors.New("nox_xxx_bookInit_45B9D0 failed")
	}
	if legacy.Sub_476E20() == nil {
		return errors.New("sub_476E20 failed")
	}
	if legacy.Sub_4BFAD0() == 0 {
		return errors.New("sub_4BFAD0 failed")
	}
	tmp := legacy.Nox_xxx_wndCreateInventoryMB_465E00()
	*memmap.PtrUint32(0x5D4594, 1096328) = tmp
	if tmp == 0 {
		return errors.New("nox_xxx_wndCreateInventoryMB_465E00 failed")
	}
	if legacy.Nox_game_initOptionsInGame_4ADAD0() == 0 {
		return errors.New("nox_game_initOptionsInGame_4ADAD0 failed")
	}
	if legacy.Sub_48D000_initGuiKick() == 0 {
		// This window is not present in Nox Demo and this must not prevent the game from booting
		gameLog.Printf("failed to initialize GuiKick.wnd")
	}
	if legacy.Sub_4C3760() == 0 {
		return errors.New("sub_4C3760 failed")
	}
	if nox_savegame_sub_46C730() == 0 {
		return errors.New("nox_savegame_sub_46C730 failed")
	}
	if legacy.Sub_4C09D0() == 0 {
		return errors.New("sub_4C09D0 failed")
	}
	if legacy.Sub_478110() == 0 {
		return errors.New("sub_478110 failed")
	}
	if legacy.Sub_49B3E0() == 0 {
		// This window is not present in Nox Demo and this must not prevent the game from booting
		gameLog.Printf("failed to initialize GGOver.wnd")
	}
	if legacy.Sub_4BFC90() == 0 {
		// This window is not present in Nox Demo and this must not prevent the game from booting
		gameLog.Printf("failed to initialize SKey.wnd")
	}
	if legacy.Nox_gui_itemAmount_init_4BFEF0() == 0 {
		return errors.New("nox_gui_itemAmount_init_4BFEF0 failed")
	}
	if legacy.Sub_4799A0() == 0 {
		return errors.New("sub_4799A0 failed")
	}
	if legacy.Nox_xxx_cliPrepareGameplay1_460E60() == 0 {
		return errors.New("nox_xxx_cliPrepareGameplay1_460E60 failed")
	}
	vsz := videoGetWindowSize()
	*memmap.PtrPtr(0x5D4594, 1096352) = unsafe.Pointer(guiCon.Init(vsz).C())
	if memmap.Uint32(0x5D4594, 1096352) == 0 {
		return errors.New("nox_gui_console_Create_450C70 failed")
	}
	if legacy.Sub_46A730() == nil {
		return errors.New("sub_46A730 failed")
	}
	if legacy.Sub_44E560() == nil {
		return errors.New("sub_44E560 failed")
	}
	if legacy.Sub_4C3500() == 0 {
		return errors.New("sub_4C3500 failed")
	}
	tmp = legacy.Nox_xxx_guiDrawRank_46E870()
	*memmap.PtrUint32(0x5D4594, 1096340) = tmp
	if tmp == 0 {
		return errors.New("nox_xxx_guiDrawRank_46E870 failed")
	}
	tmp = legacy.Nox_xxx_guiMotdLoad_4465C0()
	*memmap.PtrUint32(0x5D4594, 1096324) = tmp
	if tmp == 0 {
		return errors.New("nox_xxx_guiMotdLoad_4465C0 failed")
	}
	if legacy.Nox_xxx_guiSummonCreatureLoad_4C1D80() == 0 {
		return errors.New("nox_xxx_guiSummonCreatureLoad_4C1D80 failed")
	}
	if nox_xxx_wndLoadQuitMenu_445790() == 0 {
		return errors.New("nox_xxx_wndLoadQuitMenu_445790 failed")
	}
	if legacy.Sub_4AB260() == 0 {
		return errors.New("sub_4AB260 failed")
	}
	if legacy.Nox_xxx_guiChatIconLoad_445650() == 0 {
		return errors.New("nox_xxx_guiChatIconLoad_445650 failed")
	}
	if legacy.Sub_4C3390() == 0 {
		return errors.New("sub_4C3390 failed")
	}
	if legacy.Sub_48C980() == 0 {
		return errors.New("sub_48C980 failed")
	}
	guiCon.Enable(true)
	legacy.Sub_4AB4A0(0)
	legacy.Sub_4AB4D0(0)
	if !nox_client_renderGUI_80828 || noxflags.HasEngine(noxflags.EngineNoRendering) {
		legacy.Sub_4721A0(0)
		legacy.Sub_460EA0(0)
		legacy.Nox_window_set_visible_unk5(0)
		legacy.Sub_45D500(0)
		legacy.Sub_455A00(0)
		legacy.Sub_455F10(0)
	}
	return nil
}

func (s *Server) nox_xxx_mapFindPlayerStart_4F7AB0(a2 *server.Object) types.Pointf {
	return legacy.Nox_xxx_mapFindPlayerStart_4F7AB0(a2)
}

func (s *Server) nox_xxx_mapExitAndCheckNext_4D1860_server() error {
	if noxflags.HasGame(noxflags.GameClient) {
		noxClient.Nox_client_setCursorType(gui.CursorBusy)
	}
	legacy.Sub_4D22B0()
	s.Nox_xxx_netMsgFadeBegin_4D9800(false, true)
	s.scriptOnEvent(script.EventMapExit)
	v2 := s.getServerMap()
	legacy.Sub_500510(v2)
	s.nox_xxx_mapSwitchLevel_4D12E0(true)
	for _, obj := range s.Objs.AllNonMissiles() {
		obj.ObjFlags |= object.FlagMarked
	}
	for _, obj := range s.Objs.AllMissiles() {
		obj.ObjFlags |= object.FlagMarked
	}
	if noxflags.HasGame(noxflags.GameModeCoop) {
		s.Spells.EnableAll()
		s.Sub4537F0()
	}
	var merr error
	if nox_xxx_gameIsSwitchToSolo_4DB240() {
		v5 := nox_xxx_mapFilenameGetSolo_4DB260()
		merr = s.nox_server_loadMapFile_4CF5F0(v5, false)
	} else {
		stt := getCurrentSettings2()
		if noxflags.HasGame(noxflags.GameOnline) {
			if stt.Field57 == 0 {
				stt2, freeStt := alloc.New(server.Settings2{})
				defer freeStt()
				stt2.Field0 = stt.Field0
				v8 := legacy.Sub_459870()
				legacy.Sub_57A1E0(stt2, "user.rul", v8, 3, noxflags.GameFlag(stt.Field52))
				v9 := legacy.Sub_459870()
				legacy.Sub_57AAA0("user.rul", stt, v9)
			}
			s.Spells.EnableAll()
			s.Sub4537F0()
		}
		v10 := s.nox_server_currentMapGetFilename_409B30()
		merr = s.nox_server_loadMapFile_4CF5F0(v10, false)
		if noxflags.HasGame(noxflags.GameOnline) && !noxflags.HasGame(noxflags.GameModeChat) {
			v13 := 0
			if legacy.Nox_xxx_gamePlayIsAnyPlayers_40A8A0() != 0 {
				v12 := noxflags.GetGame()
				if legacy.Sub_40A180(v12) != 0 {
					v13 = 1
				}
			}
			legacy.Sub_40A1F0(v13)
			legacy.Sub_40A250()
			v60 := noxflags.GetGame()
			v58 := legacy.Sub_459870()
			stt0 := getSettings2ByInd(0)
			legacy.Sub_57A1E0(stt0, "user.rul", v58, 3, v60)
			sub_4D2230()
		}
	}
	if merr != nil {
		*memmap.PtrUint32(0x5D4594, 1548520) = 1
		return merr
	}
	if !noxflags.HasGame(noxflags.GameModeCoop) {
		s.MapSend.ReadMapFile()
	}
	s.ObjectsAddPending()
	for _, k := range s.Players.ListUnits() {
		legacy.Sub_4EF660(k)
		v61 := s.nox_xxx_mapFindPlayerStart_4F7AB0(k)
		if noxflags.HasGame(noxflags.GameModeChat) && s.Teams.Count() != 0 {
			if !noxflags.HasGamePlay(noxflags.GameplayFlag2) && !noxflags.HasGame(noxflags.GameFlag16) {
				if t := k.Team(); t != nil {
					v61 = s.RandomReachablePointAround(50.0, legacy.AsPointf(unsafe.Add(t.Field_72, 56)))
				}
			}
		}
		asObjectS(k).SetPos(v61)
		if !noxflags.HasGame(noxflags.GameModeCoopTeam) {
			plx := k.ControllingPlayer()
			plx.Lessons = 0
			plx.Field2140 = 0
			plx.Field2144 = s.Frame()
			s.Nox_xxx_netReportLesson_4D8EF0(k)
		}
	}
	s.AI.Paths.Sub50AFA0()
	s.AI.Paths.IndexObjects()
	*memmap.PtrUint32(0x5D4594, 1548528) = 0
	legacy.Set_dword_5d4594_1548532(nil)
	if memmap.Uint32(0x5D4594, 1548708) == 0 {
		*memmap.PtrUint32(0x5D4594, 1548708) = uint32(s.Types.ByID("PlayerStart").Ind())
	}
	if memmap.Uint32(0x5D4594, 1548712) == 0 {
		*memmap.PtrUint32(0x5D4594, 1548712) = uint32(s.Types.ByID("GameBallStart").Ind())
	}
	var starts struct {
		playerN int
		flagN   int
		ballN   int
	}
	for _, v22 := range s.Objs.AllNonMissiles() {
		if v22.Class().Has(0x10000000) {
			v22.Field34 = 0
			legacy.Set_dword_5d4594_1548532(v22.CObj())
			starts.flagN++
		}
		ti := uint32(v22.TypeInd)
		if ti == memmap.Uint32(0x5D4594, 1548712) {
			starts.ballN++
		}
		if ti == memmap.Uint32(0x5D4594, 1548708) {
			starts.playerN++
		}
	}
	if starts.playerN == 0 {
		v24 := s.Strings().GetStringInFile("StartingPositionError", "C:\\NoxPost\\src\\Server\\System\\server.c")
		legacy.PrintToPlayers(v24)
		return errors.New("cannot find player starting position")
	}
	if noxflags.HasGame(noxflags.GameModeCTF|noxflags.GameModeFlagBall) && starts.flagN < 2 {
		v24 := s.Strings().GetStringInFile("FlagCountError", "C:\\NoxPost\\src\\Server\\System\\server.c")
		legacy.PrintToPlayers(v24)
		return errors.New("invalid flag starting position(s)")
	}
	if noxflags.HasGame(noxflags.GameModeFlagBall) && starts.ballN < 1 {
		v24 := s.Strings().GetStringInFile("BallStartCountError", "C:\\NoxPost\\src\\Server\\System\\server.c")
		legacy.PrintToPlayers(v24)
		return errors.New("invalid ball starting position(s)")
	}

	if noxflags.HasGamePlay(noxflags.GameplayFlag4) || noxflags.HasGame(noxflags.GameFlag16) {
		s.Teams.ResetYyy()
	}
	s.Map.Debug.Each([]string{"spring", "SentryGlobe"}, func(key, val string) {
		switch key {
		case "spring":
			val = strings.TrimSpace(val)
			if sub := strings.Fields(val); len(sub) >= 2 {
				ind1, _ := strconv.Atoi(sub[0])
				obj1 := s.Objs.GetObjectByInd(ind1)
				ind2, _ := strconv.Atoi(sub[1])
				obj2 := s.Objs.GetObjectByInd(ind2)
				if obj1 != nil && obj2 != nil {
					s.springs.Add(obj1, obj2)
				}
			}
		case "SentryGlobe":
			val = strings.TrimSpace(val)
			if sub := strings.Fields(val); len(sub) >= 3 {
				v34, _ := strconv.Atoi(sub[0])
				if v35 := s.Objs.GetObjectByInd(v34); v35 != nil {
					v36 := unsafe.Slice((*float32)(v35.UpdateData), 3)
					v38, _ := strconv.ParseFloat(sub[1], 64)
					v36[0] = float32(v38 * memmap.Float64(0x581450, 10008))
					v36[1] = v36[0]
					v40, _ := strconv.ParseFloat(sub[2], 64)
					v36[2] = float32(v40 * memmap.Float64(0x581450, 10000) / float64(s.TickRate()))
				}
			}
		}
	})
	legacy.Set_dword_5d4594_1548524(0)
	v41 := s.getServerMap()
	legacy.Sub_500510(v41)
	if !nox_xxx_gameIsSwitchToSolo_4DB240() {
		s.ShouldCallMapInit = true
	}
	noxflags.SetGame(noxflags.GameFlag28)
	noxAudioServeT(500)
	v42 := nox_xxx_gameIsSwitchToSolo_4DB240()
	legacy.Sub_4DBA30(v42)
	noxAudioServe()
	noxflags.UnsetGame(noxflags.GameFlag28)
	s.ShouldCallMapEntry = true
	if s.mapSwitchWPName != "" {
		if wp := s.WPs.ByID(s.mapSwitchWPName); wp != nil {
			gameLog.Printf("moving player to waypoint: %q", s.mapSwitchWPName)
			wpos := wp.Pos()
			for _, u := range s.Players.ListUnits() {
				asObjectS(u).SetPos(wpos)
			}
		} else {
			gameLog.Printf("cannot find map waypoint %q!", s.mapSwitchWPName)
		}
	}
	s.mapSwitchWPName = ""
	if sub_4DCC00() {
		for _, m := range s.Players.ListUnits() {
			for _, np := range m.GetOwned516() {
				if legacy.Nox_xxx_isUnit_4E5B50(np) != 0 {
					ud := np.UpdateDataMonster()
					v61 := s.RandomReachablePointAround(50.0, m.Pos())
					asObjectS(np).SetPos(v61)
					ud.Field97 = 0
					np.ClearActionStack()
					np.Obj130 = nil
					ud.CurrentEnemy = nil
					ud.Field304 = 0
					*(*byte)(unsafe.Add(unsafe.Pointer(ud), 1129)) = 0
					ud.Field0 = 0xDEADFACE
					p := m.Pos()
					np.MonsterPushAction(ai.ACTION_ESCORT, p.X, p.Y, m)
					if np.Class().Has(object.ClassMonster) && np.SubClass().AsMonster().HasAny(object.MonsterNPC|object.MonsterFemaleNPC) {
						np.Nox_xxx_setNPCColor_4E4A90(0, &ud.Color[0])
					}
				} else if m.Class().Has(1) && legacy.Sub_4E5B80(m) != 0 {
					legacy.Sub_4E81D0(m)
				}
			}
		}
		sub_4DCBF0(0)
	}
	_ = object.FlagMarked
	for _, obj := range s.Objs.AllNonMissiles() {
		obj.ObjFlags &= 0x7FFFFFFF
	}
	for _, obj := range s.Objs.AllMissiles() {
		obj.ObjFlags &= 0x7FFFFFFF
	}
	if noxflags.HasGame(noxflags.GameModeKOTR) && noxflags.HasGamePlay(noxflags.GameplayFlag4) {
		legacy.Sub_4D2160()
	}
	if noxflags.HasGame(noxflags.GameModeQuest) {
		s.nox_server_questMapNextLevel()
	}
	if noxflags.HasGame(noxflags.GameModeCoop) && legacy.Nox_xxx_mapLoadRequired_4DCC80() == 0 {
		SaveCoopX(common.SaveAuto, 30)
	}
	legacy.Nox_xxx_mapLoadOrSaveMB_4DCC70(0)
	if noxflags.HasGame(noxflags.GameModeCoop) {
		sub_413980(30)
	}
	if noxflags.HasGame(noxflags.GameModeQuest) {
		legacy.Sub_4D7140(s.Frame() + s.SecToFrames(1))
	} else {
		legacy.Sub_4D7140(0)
	}
	nox_ticks_reset_416D40()
	legacy.Sub_40A970()
	s.flag1548704 = false
	if noxflags.HasGame(noxflags.GameModeQuest) {
		legacy.Sub_4D71E0(0)
		if s.Doors.Sub_4D72C0() {
			s.Doors.Sub_4D72B0(false)
			v57 := s.Doors.Sub_4D72C0()
			s.Sub_4D7280(255, v57)
		}
	}
	return nil
}

func sub_4D2230() {
	if noxflags.HasGame(0xC000) {
		if legacy.Nox_xxx_servGamedataGet_40A020(uint16(noxflags.GetGame())) == 0 {
			if legacy.Sub_40A180(noxflags.GetGame()) == 0 {
				legacy.Sub_409FB0_settings(uint16(noxflags.GetGame()), 0xF)
			}
		}
	} else if noxflags.HasGame(1024) && legacy.Nox_xxx_servGamedataGet_40A020(1024) == 0 {
		legacy.Sub_409FB0_settings(1024, 0xF)
	}
}

func sub_4DCC00() bool {
	return dword_5d4594_1563064
}

func sub_413980(a1 int) {
	if a1 != 0 {
		sub_413A00(1)
		legacy.Set_dword_5d4594_251744(a1 + int(noxClient.GetInputSeq()))
	} else {
		legacy.Set_dword_5d4594_251744(0)
	}
}

func sub_4139C0() {
	if legacy.Get_dword_5d4594_251744() != 0 {
		if int(noxClient.GetInputSeq()) >= int(legacy.Get_dword_5d4594_251744()) {
			legacy.Set_dword_5d4594_251744(0)
			if !sub_450560() {
				sub_413A00(0)
			}
		}
	}
}

func (s *Server) sub_417160() {
	for _, it := range s.Players.List() {
		legacy.Nox_xxx_playerUnsetStatus_417530(it, 16)
	}
}

func nox_xxx_calcDistance_4E6C00(obj1, obj2 server.Obj) float32 {
	return legacy.Nox_xxx_calcDistance_4E6C00(server.ToObject(obj1), server.ToObject(obj2))
}

var doDamageWalls = true

func nox_xxx_gameSetWallsDamage_4E25A0(v int) {
	doDamageWalls = v != 0
}

func (s *Server) Nox_xxx_mapDamageUnitsAround(pos types.Pointf, r1, r2 float32, dmg int, dtyp object.DamageType, who *server.Object, a7 server.Obj, damageWalls bool) {
	rr := r1
	if r1 < r2 {
		rr = r2
	}
	rect := types.Rectf{
		Min: pos.Sub(types.Ptf(rr, rr)),
		Max: pos.Add(types.Ptf(rr, rr)),
	}
	s.Map.EachObjInRect(rect, func(it *server.Object) bool {
		u := it
		if u == who && !damageWalls {
			return true
		}
		if u == server.ToObject(a7) {
			return true
		}
		pos2 := u.Pos()
		dx := pos2.X - pos.X
		dy := pos2.Y - pos.Y
		dist := float32(math.Sqrt(float64(dy*dy + dx*dx)))
		if dist > rr {
			return true
		}
		if !s.MapTraceRay(pos, pos2, server.MapTraceFlag1) {
			return true
		}
		rdmg := float32(dmg)
		if dist >= r2 {
			rdmg *= 1.0 - (dist-r2)/(rr-r2)
		}
		asObjectS(u).CallDamage(who, nil, int(rdmg), dtyp)
		return true
	})
	wrect := image.Rect(
		int(rect.Min.X)/common.GridStep,
		int(rect.Min.Y)/common.GridStep,
		int(rect.Max.X)/common.GridStep,
		int(rect.Max.Y)/common.GridStep,
	)
	s.Nox_xxx_mapDamageToWalls_534FC0(wrect, pos, r1, dmg, dtyp, who)
	doDamageWalls = true
}

func sub_473840() {
	guiCon.Enable(false)
	legacy.Nox_gui_itemAmount_free_4C03E0()
	sub_46CCB0()
	legacy.Sub_4AE3B0()
	legacy.Sub_48D450()
	legacy.Sub_4C4220()
	legacy.Nox_xxx_closeP2PTradeWnd_4C12A0()
	legacy.Sub_4BFD10()
	legacy.Sub_49B490()
	legacy.Sub_478F80()
	legacy.Sub_479D10()
	legacy.Sub_4AB470()
	legacy.Sub_4C34A0()
	legacy.Sub_445770()
	legacy.Sub_456240()
	legacy.Sub_455EE0()
	sub_460E90()
	legacy.Sub_4505E0()
	legacy.Sub_46A860()
	legacy.Sub_49C7A0()
	legacy.Nox_xxx_guiServerOptionsHide_4597E0(1)
	legacy.Sub_467980()
	noxClient.GUI.DestroyAll()
	nox_client_renderGUI_80828 = nox_xxx_xxxRenderGUI_587000_80832
}

func sub_460E90() {
	if legacy.Sub_460D40() {
		legacy.Sub_460D50()
	}
}

func (s *Server) nox_xxx_netlist_4DEB50() {
	if !noxflags.HasGame(noxflags.GameClient) {
		return
	}
	if noxflags.HasEngine(noxflags.EngineReplayRead) {
		s.nox_xxx_replayTickMB_4D3580_net_playback(false)
		s.NetList.ResetByInd(server.HostPlayerIndex, netlist.Kind0)
	} else if !isDedicatedServer {
		s.NetList.HandlePacketsA(server.HostPlayerIndex, netlist.Kind0, func(data []byte) {
			if len(data) == 0 {
				return
			}
			s.onPacketRaw(server.HostPlayerIndex, data)
		})
	}
}

func (c *Client) nox_game_checkStateSwitch_43C1E0() {
	if c.GameGetStateCode() < 0 {
		return
	}
	p := gui.FindAnimForStateID(c.GameGetStateCode())
	if p != nil {
		p.Func12()
		p.Func13Ptr = legacy.Get_nox_game_switchStates_43C0A0()
		c.GamePopState()
	}
}

func (c *Client) nox_game_checkStateOptions_43C220() {
	if c.GameGetStateCode() < 0 {
		return
	}
	p := gui.FindAnimForStateID(c.GameGetStateCode())
	if p != nil {
		p.Func12()
		p.Func13Ptr = legacy.Get_nox_game_showOptions_4AA6B0()
	}
}

func (c *Client) nox_game_checkStateMenu_43C2F0() {
	if c.GameGetStateCode() < 0 {
		return
	}
	p := gui.FindAnimForStateID(c.GameGetStateCode())
	if p != nil {
		if c.GameGetStateCode() == client.StateMainMenu {
			nox_game_showMainMenu_4A1C00()
		} else {
			p.Func12()
			_ = nox_game_showMainMenu_4A1C00
			p.Func13Ptr = legacy.Get_nox_game_showMainMenu_4A1C00()
		}
	}
	c.GamePopStateUntil(client.StateMainMenu)
}
