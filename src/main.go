package opennox

import (
	"errors"
	"flag"
	"fmt"
	"image"
	"image/color"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"unsafe"

	"github.com/noxworld-dev/opennox-lib/platform"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"

	"github.com/noxworld-dev/opennox-lib/client/keybind"
	"github.com/noxworld-dev/opennox-lib/common"
	"github.com/noxworld-dev/opennox-lib/datapath"
	"github.com/noxworld-dev/opennox-lib/env"
	"github.com/noxworld-dev/opennox-lib/log"

	"github.com/noxworld-dev/opennox/v1/client"
	"github.com/noxworld-dev/opennox/v1/client/gui"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"github.com/noxworld-dev/opennox/v1/common/serial"
	"github.com/noxworld-dev/opennox/v1/common/sound"
	"github.com/noxworld-dev/opennox/v1/internal/version"
	"github.com/noxworld-dev/opennox/v1/legacy"
	"github.com/noxworld-dev/opennox/v1/legacy/client/audio/ail"
	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc/handles"
)

func init() {
	http.Handle("/metrics", promhttp.Handler())
	viper.SetDefault(configManualSpellCastDelay, 0.5)
	promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "opennox_version",
		Help: "OpenNox version",
	}, []string{"vers", "commit"}).WithLabelValues(
		version.Version(), version.Commit(),
	).Inc()
}

const (
	configManualSpellCastDelay = "extensions.spells.manual_cast.timeout"
)

var (
	isServer      bool
	isServerQuest bool
	serverExec    []string
	onDataPathSet []func()
)

func registerOnDataPathSet(fnc func()) {
	onDataPathSet = append(onDataPathSet, fnc)
}

// Nox only works on 32bit
var _ = [1]struct{}{}[unsafe.Sizeof(int(0))-4]

func writeLogsToDir(dir string) error {
	if dir == "" {
		dir = filepath.Dir(os.Args[0])
		if sdir := env.AppUserDir(); sdir != "" {
			dir = sdir
		}
		dir = filepath.Join(dir, "logs")
	}
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	name := "opennox.log"
	if isDedicatedServer {
		name = "opennox-server.log"
	}
	return log.WriteToFile(filepath.Join(dir, name))
}

func RunArgs(args []string) (gerr error) {
	defer func() {
		switch r := recover().(type) {
		case ErrExit:
			if r != 0 {
				gerr = r
			}
		default:
			panic(r)
		case nil:
			// ok
		}
	}()
	flags := flag.NewFlagSet("", flag.ContinueOnError)
	// TODO: add missing flag descriptions
	rconDefHost := ""
	if env.IsDevMode() {
		rconDefHost = rconDefaultHost
	}
	var (
		fConfig     = flags.String("config", "", "use specified config file")
		fData       = flags.String("data", "", "explicitly set Nox data dir")
		fServer     = flags.Bool("serveronly", false, "run the server only")
		fWindow     = flags.Bool("window", false, "window")
		fSWindow    = flags.Bool("swindow", false, "swindow")
		fFullScreen = flags.Bool("fullscreen", false, "fullscreen")
		fMinimize   = flags.Bool("minimize", false, "minimize")
		fNoLimit    = flags.Bool("nolimit", false, "nolimit")
		fDrop       = flags.Int("drop", 0, "drop given percent of network packets")
		fNoText     = flags.Bool("notext", false, "notext")
		fLock       = flags.Bool("lock", false, "lock")
		fSafe       = flags.Bool("safe", false, "run in safe mode")
		fNoAudio    = flags.Bool("noaudio", false, "disable audio")
		fNoMMX      = flags.Bool("noMMX", false, "noMMX")
		fNoThreads  = flags.Bool("nothread", false, "nothread")
		fNoFloor    = flags.Bool("noFloor", false, "noFloor")
		fNoDraw     = flags.Bool("noDraw", false, "noDraw")
		fPort       = flags.Int("port", common.GamePort, "server port number to use")
		fClientPort = flags.Int("clientport", 0, "clientport")
		fNoSoft     = flags.Bool("nosoft", false, "nosoft")
		fNoSkip     = flags.Bool("noskip", false, "noskip")
		// TODO: replace with -serveronly once we figure out all the details
		fAutoServer = flags.Bool("autosrv", isDedicatedServer, "automatically start the server")
		fAutoQuest  = flags.Bool("autoquest", false, "automatically start the quest game")
		fAutoExec   = flags.String("autoexec", "load estate", "run the specified command at server startup")
		fRecord     = flags.String("record", "", "record the game to a given file")
		fReplay     = flags.String("replay", "", "replay game recording from a given file")
		fPProf      = flags.String("pprof", "", "enable pprof")
		fRcon       = flags.String("rcon", rconDefHost, "enable remote console")
		fRconPass   = flags.String("rcon-pass", "", "remote console password")
		fLogs       = flags.String("logs", "", "directory for logs")
	)
	if err := flags.Parse(args[1:]); err != nil {
		return err
	}
	if !env.IsE2E() {
		if err := writeLogsToDir(*fLogs); err != nil {
			log.Println("cannot persist logs:", err)
		}
		defer log.Close()
	} else {
		e2eInit()
		defer e2eStop()
	}
	version.LogVersion()

	handles.Init()
	defer handles.Release()
	if !isDedicatedServer {
		go version.Latest() // prefetch
	}
	if env.IsDevMode() || version.IsDev() || env.IsE2E() || *fPProf != "" {
		go func() {
			host := *fPProf
			if host == "" {
				host = "127.0.0.1:6060"
			}
			log.Printf("starting pprof server on %q", host)
			if err := http.ListenAndServe(host, nil); err != nil {
				log.Printf("failed to start pprof: %v", err)
			}
		}()
	}
	noxServer = NewServer(noxConsole, strMan)
	var err error
	noxClient, err = NewClient(noxConsole, noxServer)
	if err != nil {
		return err
	}
	if err := readConfig(*fConfig); err != nil {
		return fmt.Errorf("cannot read config: %w", err)
	}
	defer maybeWriteConfig()
	if cur := viper.GetString(configNoxSerial); cur == "" {
		cur = serial.Generate()
		viper.Set(configNoxSerial, cur)
		writeConfigLater()
	}
	if env.IsE2E() {
		*fNoLimit = true
		mainloopHook = e2eRun
	}
	if path := *fData; path != "" {
		datapath.SetData(path)
	} else if path = os.Getenv("NOX_DATA"); path != "" {
		datapath.SetData(path)
	} else if path = viper.GetString(configNoxDataPath); path != "" {
		datapath.SetData(path)
		if filepath.IsAbs(path) && strings.HasPrefix(path, filepath.Dir(viper.ConfigFileUsed())) {
			if rel, err := filepath.Rel(filepath.Dir(viper.ConfigFileUsed()), path); err == nil {
				viper.Set(configNoxDataPath, rel)
				writeConfigLater()
			}
		}
	} else {
		path = datapath.FindData()
		if path == "" {
			err := errors.New("cannot find Nox data dir")
			datapath.Log.Println(err)
			if runtime.GOOS == "windows" {
				errorMessage("Nox game data directory not found. Please install Nox from Origin or GoG.\nIf the problem persists, please check https://noxworld-dev.github.io/opennox-docs/opennox/install/windows/index.html")
			} else {
				errorMessage("Nox game data directory not found. Please install Nox from GoG (via Lutris or Heroic).\nIf the problem persists, please check https://noxworld-dev.github.io/opennox-docs/opennox/install/linux/index.html")
			}
			return err
		}
		datapath.SetData(path)
		if filepath.IsAbs(path) && strings.HasPrefix(path, filepath.Dir(viper.ConfigFileUsed())) {
			if rel, err := filepath.Rel(filepath.Dir(viper.ConfigFileUsed()), path); err == nil {
				path = rel
			}
		}
		viper.Set(configNoxDataPath, path)
		writeConfigLater()
	}
	maybeWriteConfig()
	if env.IsE2E() {
		viper.Reset() // defaults only
		serial.SetSerial(serial.Generate())
	} else {
		cur := viper.GetString(configNoxSerial)
		if cur == "" {
			cur = serial.Generate()
		}
		serial.SetSerial(cur)
	}
	for _, fnc := range onDataPathSet {
		fnc()
	}
	if err := os.Chdir(datapath.Data()); err != nil {
		return err
	}
	legacy.InitBlobData()
	isServer = *fAutoServer
	isServerQuest = *fAutoQuest
	serverExec = strings.Split(*fAutoExec, ";")
	if err := noxServer.Strings().ReadFile("nox.str"); err != nil {
		return fmt.Errorf("failed to load strings file: %w", err)
	}
	if !*fServer && !*fNoDraw {
		err = noxClient.initSeat(image.Point{X: noxDefaultWidth, Y: noxDefaultHeight})
		if err != nil {
			return err
		}
	}
	if *fWindow {
		dword_5d4594_805860 = false
		noxClient.UpdateFullScreen(-2)
	}
	if *fSWindow {
		dword_5d4594_805860 = true
		noxClient.UpdateFullScreen(-3)
	}
	if *fFullScreen {
		noxClient.UpdateFullScreen(-1)
	}
	if *fMinimize {
		*memmap.PtrUint32(0x5D4594, 805864) = 1
	}

	legacy.Set_nox_gameDisableMapDraw_5d4594_2650672(0)
	noxClient.GameAddStateCode(client.StateMovies)
	noxflags.ResetGame()
	noxflags.OnGameChange(func(f noxflags.GameFlag) {
		if f.Has(noxflags.GameServerSettings) {
			legacy.Nox_server_gameSettingsUpdated_40A670()
		}
	})
	noxflags.OnGameSet(func(f noxflags.GameFlag) {
		log.Printf("game flag set: %v", f)
		legacy.Nox_xxx_guiChatShowHide_445730(noxflags.GetGame().Mode() != noxflags.GameModeChat)
		if f.Has(noxflags.GameSuddenDeath) && noxflags.HasGame(noxflags.GameHost) {
			noxServer.NetPrintLineToAll("Settings.c:SuddenDeathStart")
		}
	})
	noxflags.OnGameUnset(func(f noxflags.GameFlag) {
		log.Printf("game flag unset: %v", f)
		if f.Has(noxflags.GameSuddenDeath) {
			noxServer.flag3592 = false
		}
	})
	noxflags.SetGame(noxflags.GameHost | noxflags.GameClient)
	noxflags.SetEngine(noxflags.EngineSoftShadowEdge)
	legacy.Sub_43AF50()
	noxServer.SetTickRate(30)
	noxServer.SetInitialFrame()
	nox_ticks_reset_416D40()
	noxServer.SetUpdateFunc(nil)
	noxClient.SetDrawFunc(nil)
	noxServer.SetUpdateFunc2(nil)
	noxflags.SetGame(noxflags.GameModeArena)
	if *fNoLimit {
		useFrameLimit = false
	}
	if *fServer {
		nox_enable_audio = 0
		noxflags.SetEngine(noxflags.EngineNoRendering |
			noxflags.EngineNoTextRendering |
			noxflags.EngineNoFloorRendering)
	}
	if v := *fDrop; v != 0 {
		noxServer.NetStr.PacketDrop = v
	}
	if *fNoText {
		noxflags.SetEngine(noxflags.EngineNoTextRendering)
	}
	if *fLock {
		noxflags.SetEngine(noxflags.EngineLockResolution)
	}
	if *fSafe {
		noxflags.UnsetEngine(noxflags.EngineSoftShadowEdge)
		noxflags.SetEngine(noxflags.EngineWindowed)
		nox_enable_audio = 0
		legacyUnlockSurface = true
		*memmap.PtrUint32(0x5D4594, 805840) = 1
		nox_enable_threads = false
		mode := image.Point{
			X: noxDefaultWidth,
			Y: noxDefaultHeight,
		}
		noxClient.videoSetGameMode(mode)
		videoSetMenuMode(mode)
	}
	if *fNoAudio {
		nox_enable_audio = 0
	}
	if *fNoMMX {
		*memmap.PtrUint32(0x5D4594, 805840) = 1
	}
	if *fNoThreads {
		nox_enable_threads = false
	}
	if *fNoFloor {
		noxflags.SetEngine(noxflags.EngineNoFloorRendering)
	}
	if *fNoDraw {
		noxflags.SetEngine(noxflags.EngineNoRendering)
	}
	if v := *fPort; v > 0 {
		noxServer.SetServerPort(v)
	}
	if v := *fClientPort; v > 0 {
		nox_xxx_setClientNetPort_40A410(v)
	}
	if *fNoSoft {
		noxflags.SetEngine(noxflags.EngineNoSoftLights)
	}
	if fname := *fRecord; fname != "" {
		if err := noxServer.nox_xxx_replayStartSave_4D3370(fname); err != nil {
			return err
		}
	} else if fname = *fReplay; fname != "" {
		if err := noxServer.nox_xxx_replayFileOpen_4D34C0(fname); err != nil {
			return err
		}
	}
	legacy.Nox_xxx_servSetPlrLimit_409F80(32)

	// manual spell cast timeout (in seconds)
	msmul := viper.GetFloat64(configManualSpellCastDelay)
	// manual spell cast timeout (in frames)
	spellTimeout = noxServer.SecToFramesF(msmul)

	if err := nox_common_scanAllMaps_4D07F0(); err != nil {
		return fmt.Errorf("cannot find maps: %w", err)
	}
	keyBinding = keybind.New(noxClient.Strings())
	nox_xxx_mapSetDataDefault_416500()
	if err := nox_common_readcfgfile("nox.cfg", false); err != nil {
		return fmt.Errorf("failed to load config file: %w", err)
	}
	if env.IsE2E() {
		noxClient.videoSetGameMode(image.Point{X: 1024, Y: 768})
	} else {
		noxClient.videoSetGameMode(image.Point{
			X: viper.GetInt(configVideoWidth),
			Y: viper.GetInt(configVideoHeight),
		})
	}
	if err := gameexReadConfig("game_ex.cfg"); err != nil {
		return fmt.Errorf("failed to load gameex config file: %w", err)
	}
	inputClearKeyTimeouts()
	noxServer.Rand.Reset()
	legacy.Nox_xxx_loadLook_415D50()
	legacy.Nox_xxx_loadModifyers_4158C0()
	initConsole(noxServer.Strings())
	legacy.Sub_4D11A0()
	if !isDedicatedServer {
		videoResizeView(image.Point{})
		if err := noxClient.gameResetVideoMode(true, true); err != nil {
			return fmt.Errorf("failed to update video mode: %w", err)
		}
		noxClient.r.ClearScreen(color.Black)
	} else {
		enableGUIDrawing(false)
		videoInitStub()
	}
	if err := noxClient.r.Bag.ReadVideoBag(); err != nil {
		return err
	}
	if legacy.Sub_431370() == 0 {
		return fmt.Errorf("sub_431370 failed")
	}
	if nox_audio_initall(nox_enable_audio) == 0 && nox_enable_audio != 0 {
		return fmt.Errorf("failed to init audio")
	}
	if err := nox_xxx_loadAllBinFiles_415470(); err != nil {
		thingsLog.Println(err)
		return fmt.Errorf("failed to load bin files: %w", err)
	}
	if err := nox_xxx_initInput_430190(); err != nil {
		return fmt.Errorf("failed to init input: %w", err)
	}
	if noxServer.Walls.Init() == 0 {
		return fmt.Errorf("failed to init map")
	}
	if legacy.Nox_xxx_tileAlloc_410F60_init() == 0 {
		return fmt.Errorf("failed to init tiles")
	}
	legacy.Nox_xxx_initSinCosTables_414C90()
	legacy.Nox_xxx_loadMapCycle_4D0A30()
	platform.RandSeedTime()
	legacy.Nox_xxx_mapSelectFirst_4D0E00()
	noxServer.NetList.Init()
	noxServer.NetXfer.Init(32, xferDataCallback40AF90)
	legacy.Sub_4134D0()
	if v := noxClient.Strings().Lang(); v == 6 || v == 8 {
		noxClient.r.SetBold(false)
	}
	legacy.Sub_413920()
	nox_client_initScreenParticles_431390()
	mainloopNoSkip = *fNoSkip
	if host := *fRcon; host != "" {
		rconLog.Println("starting rcon at", host)
		rcon, err := NewRemoteConsole(host, execConsoleCmdAuthed, RconOptions{
			Pass: *fRconPass,
		})
		if err != nil {
			if !env.IsDevMode() {
				return err
			}
			rconLog.Println(err)
		} else {
			defer rcon.Close()
			consoleMux.Add(rcon)
			go func() {
				if err := rcon.Serve(); err != nil {
					rconLog.Println(err)
				}
			}()
		}
	}
	cmainLoop()
	return nil
}

type ErrExit int

func (e ErrExit) Error() string {
	return fmt.Sprintf("exit code: %d", int(e))
}

func nox_exit(exitCode int) {
	panic(ErrExit(exitCode))
}

func nox_xxx_gameGetScreenBoundaries_43BEB0_get_video_mode(w, h, d *int) {
	mode := noxClient.videoGetGameMode()
	if w != nil {
		*w = mode.X
	}
	if h != nil {
		*h = mode.Y
	}
	if d != nil {
		*d = 16
	}
}

func videoUpdateGameMode(mode image.Point) {
	noxClient.videoSetGameMode(mode)
	changeWindowedOrFullscreen()
}

func sub_4AA9C0() int {
	legacy.Sub_44D8F0()
	if !env.IsE2E() {
		videoUpdateGameMode(guiOptionsRes)
	}
	writeConfigLegacy("nox.cfg")
	legacy.Get_nox_wnd_xxx_1309740().SetState(gui.AnimOut)
	gui.SetAnimGlobalState(gui.AnimOut)
	clientPlaySoundSpecial(sound.SoundShellSlideOut, 100)
	return 1
}

func cleanup() {
	log.Println("cleanup")
	writeConfigLegacy("nox.cfg")
	nox_xxx_freeScreenParticles_4314D0()
	sub_413960()
	legacy.Sub_431380()
	noxServer.Modif.Nox_xxx_freeWeaponArmorDefAndModifs_413060()
	sub_4311B0()
	noxClient.nox_video_freeFloorBuffer_430EC0()
	nox_xxx_freeKeyboard_430210()
	legacy.Nox_xxx_tileFree_410FC0_free()
	noxServer.Walls.Free()
	nox_video_bagFree_42F4D0()
	legacy.Sub_42EDC0()
	noxClient.ctrl.Reset()
	Nox_strman_free_410020()
	noxServer.NetList.Free()
	legacy.Sub_4D11D0()
	legacy.Sub_4D0DA0()
	noxServer.NetXfer.Free()
	legacy.Nox_common_maplist_free_4D0970()
	ail.Shutdown()
}

func sub_4311B0() {
	sub_43DCC0()
	sub_4312B0()
	sub_43D970()
	legacy.Sub_44D8C0()
	legacy.Sub_451970()
	if dword_5d4594_805980 != nil {
		dword_5d4594_805980.Free()
		dword_5d4594_805980 = nil
	}
	legacy.Set_dword_587000_81128(nil)
	legacy.Sub_431270()
	if dword_5d4594_1193336 != 0 {
		legacy.Sub_4875F0()
		legacy.Sub_4870A0()
		dword_5d4594_1193336 = 0
	}
}

func sub_4312B0() {
	legacy.Sub_431290()
	legacy.Sub_44D8F0()
	legacy.Sub_43D990()
}

func sub_43D970() {
	if legacy.Get_dword_5d4594_816340() != 0 {
		sub_43DCC0()
		legacy.Set_dword_5d4594_816340(0)
	}
}

func sub_43DCC0() {
	if legacy.Get_dword_587000_93156() == 0 {
		return
	}
	if legacy.Get_dword_5d4594_816340() == 0 {
		return
	}
	legacy.Sub_43D990()
	for legacy.Get_dword_5d4594_816348() != 0 {
		if legacy.Get_dword_5d4594_816348() == 3 {
			sub_43D650()
			legacy.Set_dword_5d4594_816348(0)
		}
		noxClient.sub4312C0()
	}
}
