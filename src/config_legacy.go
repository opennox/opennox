package opennox

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unsafe"

	"github.com/noxworld-dev/opennox-lib/cfg"
	"github.com/noxworld-dev/opennox-lib/common"
	"github.com/noxworld-dev/opennox-lib/env"
	"github.com/noxworld-dev/opennox-lib/ifs"

	"github.com/noxworld-dev/opennox/v1/common/config"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"github.com/noxworld-dev/opennox/v1/legacy"
	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc"
	"github.com/noxworld-dev/opennox/v1/server"
)

var (
	legacyUnlockSurface = false
	legacyGamma2Set     = false
	legacyGamma2        = 1.0
	legacyGamma         = 50
	legacyProfiled      = 0
)

func nox_common_writecfgfile(path string) {
	if err := writeConfigLegacy(path); err != nil {
		config.Global.Log.Error("cannot write legacy config", "path", path, "err", err)
	}
}

func noxConfigRead(path string, skip bool) error {
	c := noxClient
	config.Global.Log.Info("reading legacy config", "path", path)
	c.ctrl.Reset()
	f, err := ifs.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	if err := parseLegacyConfig(f, skip); err != nil {
		return err
	}
	return nil
}

func nox_common_readcfgfile(path string, skip bool) error {
	c := noxClient
	if env.IsE2E() {
		c.ctrl.Reset()
		return parseLegacyConfig(strings.NewReader(e2eInputConf), skip)
	}
	err := noxConfigRead(path, skip)
	if err == nil {
		return nil
	} else if !os.IsNotExist(err) {
		return err
	}
	// not found: read default and write config back
	orig := err
	err = noxConfigRead("default.cfg", skip)
	if os.IsNotExist(err) {
		return orig
	} else if err != nil {
		return err
	}
	writeConfigLegacy("nox.cfg")
	return nil
}

func parseLegacyConfig(r io.Reader, skip bool) error {
	file, err := cfg.Parse(r)
	if err != nil {
		return fmt.Errorf("config parsing failed: %w", err)
	}
	if !skip {
		var sect cfg.Section
		if len(file.Sections) > 0 {
			sect = file.Sections[0]
		}
		if err := nox_common_parsecfg_all(sect); err != nil {
			return err
		}
	}
	var sect cfg.Section
	if len(file.Sections) > 1 {
		sect = file.Sections[1]
	}
	if err := nox_common_skipcfgfile_4331E0(sect, skip); err != nil {
		return err
	}
	return nil
}

func nox_common_parsecfg_all(sect cfg.Section) error {
	configSetVolume(VolumeMax, VolumeFX)
	configSetVolume(VolumeMax, VolumeDialog)
	configSetVolume(VolumeMax, VolumeMusic)
	for _, kv := range sect {
		switch kv.Key {
		case "Version":
			v, err := strconv.ParseUint(kv.Value, 10, 32)
			if err != nil {
				return fmt.Errorf("cannot parse %s: %w", kv.Key, err)
			}
			*memmap.PtrUint32(0x587000, 81284) = uint32(v)
		case "VideoMode":
			sub := strings.SplitN(kv.Value, " ", 3)
			if len(sub) < 2 {
				return fmt.Errorf("cannot parse %s: %q", kv.Key, kv.Value)
			}
			w, err := strconv.Atoi(sub[0])
			if err != nil {
				return fmt.Errorf("cannot parse %s width: %w", kv.Key, err)
			}
			h, err := strconv.Atoi(sub[1])
			if err != nil {
				return fmt.Errorf("cannot parse %s height: %w", kv.Key, err)
			}
			// TODO: should set resolution in the new config if it's not set
			_, _ = w, h
		case "VideoSize":
			v, err := strconv.Atoi(kv.Value)
			if err != nil {
				return fmt.Errorf("cannot parse %s: %w", kv.Key, err)
			}
			nox_video_setCutSize(v)
		case "Gamma":
			v, err := strconv.Atoi(kv.Value)
			if err != nil {
				return fmt.Errorf("cannot parse %s: %w", kv.Key, err)
			}
			legacyGamma = v
		case "Gamma2":
			v, err := strconv.ParseFloat(kv.Value, 64)
			if err != nil {
				return fmt.Errorf("cannot parse %s: %w", kv.Key, err)
			}
			legacyGamma2Set = true
			legacyGamma2 = v
		case "FXVolume":
			if err := configParseVolume(kv.Value, VolumeFX); err != nil {
				return fmt.Errorf("cannot parse %s: %w", kv.Key, err)
			}
		case "DialogVolume":
			if err := configParseVolume(kv.Value, VolumeDialog); err != nil {
				return fmt.Errorf("cannot parse %s: %w", kv.Key, err)
			}
		case "MusicVolume":
			if err := configParseVolume(kv.Value, VolumeMusic); err != nil {
				return fmt.Errorf("cannot parse %s: %w", kv.Key, err)
			}
		case "LastServer":
			alloc.StrCopy(unsafe.Slice((*byte)(memmap.PtrOff(0x5D4594, 806060)), 22), kv.Value)
		case "ServerName":
			name := kv.Value
			if memmap.Uint32(0x587000, 81284) == 0x10000 && name == "NoxWorld" {
				name = "User_Game"
			}
			name = strings.ReplaceAll(name, "_", " ")
			legacy.Nox_xxx_gameSetServername_40A440(name)
		case "UnlockSurface":
			v, err := strconv.Atoi(kv.Value)
			if err != nil {
				return fmt.Errorf("cannot parse %s: %w", kv.Key, err)
			}
			legacyUnlockSurface = v != 0
		case "SoftShadowEdge":
			v, err := strconv.Atoi(kv.Value)
			if err != nil {
				return fmt.Errorf("cannot parse %s: %w", kv.Key, err)
			}
			if v != 0 {
				noxflags.SetEngine(noxflags.EngineSoftShadowEdge)
			} else {
				noxflags.UnsetEngine(noxflags.EngineSoftShadowEdge)
			}
		case "DrawFrontWalls":
			v, err := strconv.Atoi(kv.Value)
			if err != nil {
				return fmt.Errorf("cannot parse %s: %w", kv.Key, err)
			}
			nox_client_drawFrontWalls_80812 = v != 0
		case "TranslucentFrontWalls":
			v, err := strconv.Atoi(kv.Value)
			if err != nil {
				return fmt.Errorf("cannot parse %s: %w", kv.Key, err)
			}
			legacy.Set_nox_client_translucentFrontWalls_805844(bool2int(v != 0))
		case "HighResFrontWalls":
			v, err := strconv.Atoi(kv.Value)
			if err != nil {
				return fmt.Errorf("cannot parse %s: %w", kv.Key, err)
			}
			legacy.Set_nox_client_highResFrontWalls_80820(bool2int(v != 0))
		case "HighResFloors":
			v, err := strconv.Atoi(kv.Value)
			if err != nil {
				return fmt.Errorf("cannot parse %s: %w", kv.Key, err)
			}
			legacy.Set_nox_client_highResFloors_154952(bool2int(v != 0))
		case "LockHighResFloors":
			v, err := strconv.Atoi(kv.Value)
			if err != nil {
				return fmt.Errorf("cannot parse %s: %w", kv.Key, err)
			}
			nox_client_lockHighResFloors_1193152 = v != 0
		case "TexturedFloors":
			v, err := strconv.Atoi(kv.Value)
			if err != nil {
				return fmt.Errorf("cannot parse %s: %w", kv.Key, err)
			}
			nox_client_texturedFloors_154956 = v != 0
			nox_xxx_tileSetDrawFn_481420()
			if v != 0 {
				legacy.Set_dword_5d4594_1193156(0)
			}
			nox_client_texturedFloors2_154960 = nox_client_texturedFloors_154956
		case "TranslucentConsole":
			v, err := strconv.Atoi(kv.Value)
			if err != nil {
				return fmt.Errorf("cannot parse %s: %w", kv.Key, err)
			}
			guiCon.translucent = v != 0
		case "RenderGlow":
			v, err := strconv.Atoi(kv.Value)
			if err != nil {
				return fmt.Errorf("cannot parse %s: %w", kv.Key, err)
			}
			noxClient.r.Part.RenderGlow = v != 0
		case "RenderGUI":
			v, err := strconv.Atoi(kv.Value)
			if err != nil {
				return fmt.Errorf("cannot parse %s: %w", kv.Key, err)
			}
			nox_client_renderGUI_80828 = v != 0
		case "FadeObjects":
			v, err := strconv.Atoi(kv.Value)
			if err != nil {
				return fmt.Errorf("cannot parse %s: %w", kv.Key, err)
			}
			legacy.Set_nox_client_fadeObjects_80836(bool2int(v != 0))
		case "TrackData":
			v, err := strconv.Atoi(kv.Value)
			if err != nil {
				return fmt.Errorf("cannot parse %s: %w", kv.Key, err)
			}
			sub_578DE0(byte(v))
		case "RenderBubbles":
			v, err := strconv.Atoi(kv.Value)
			if err != nil {
				return fmt.Errorf("cannot parse %s: %w", kv.Key, err)
			}
			legacy.Set_nox_client_renderBubbles_80844(bool2int(v != 0))
		case "SysopPassword":
			legacy.Nox_xxx_sysopSetPass_40A610(kv.Value)
		case "ServerPassword":
			sst := getServerSettings()
			alloc.StrCopyZero16(sst.ServerPass78[:], kv.Value)
		case "Profiled":
			v, err := strconv.Atoi(kv.Value)
			if err != nil {
				return fmt.Errorf("cannot parse %s: %w", kv.Key, err)
			}
			legacyProfiled = v
		case "SanctuaryHelp":
			v, err := strconv.Atoi(kv.Value)
			if err != nil {
				return fmt.Errorf("cannot parse %s: %w", kv.Key, err)
			}
			legacy.Set_nox_server_sanctuaryHelp_54276(bool2int(v != 0))
		case "MaxPacketLossPct":
			v, err := strconv.ParseUint(kv.Value, 10, 32)
			if err != nil {
				return fmt.Errorf("cannot parse %s: %w", kv.Key, err)
			}
			*memmap.PtrUint32(0x587000, 81280) = uint32(v)
			noxServer.NetStr.MaxPacketLoss = int(float64(v) * 0.0099999998 * 10.0)
		case "SendMessageOfTheDay":
			v, err := strconv.Atoi(kv.Value)
			if err != nil {
				return fmt.Errorf("cannot parse %s: %w", kv.Key, err)
			}
			legacy.Set_nox_server_sendMotd_108752(bool2int(v != 0))
		case "InternetUpdateRate":
			v, err := strconv.Atoi(kv.Value)
			if err != nil {
				return fmt.Errorf("cannot parse %s: %w", kv.Key, err)
			}
			legacy.Nox_xxx_rateUpdate_40A6D0(v)
			if nox_xxx_rateGet_40A6C0() <= 0 || nox_xxx_rateGet_40A6C0() > 3 {
				legacy.Nox_xxx_rateUpdate_40A6D0(1)
			}
		case "ConnectionType":
			v, err := strconv.Atoi(kv.Value)
			if err != nil {
				return fmt.Errorf("cannot parse %s: %w", kv.Key, err)
			}
			found := false
			for _, t := range configConnTypes {
				if t.Val == v {
					found = true
					break
				}
			}
			if found {
				if v < 0 || v > 4 {
					v = 4
				}
				legacy.Set_nox_server_connectionType_3596(v)
			}
		case "MapCycle":
			v, err := strconv.Atoi(kv.Value)
			if err != nil {
				return fmt.Errorf("cannot parse %s: %w", kv.Key, err)
			}
			legacy.Sub_4D0D90(bool2int(v != 0))
		case "LANFilters":
			err := sub_432E50(kv.Value, 0)
			if err != nil {
				return fmt.Errorf("cannot parse %s: %w", kv.Key, err)
			}
		case "INETFilters":
			err := sub_432E50(kv.Value, 1)
			if err != nil {
				return fmt.Errorf("cannot parse %s: %w", kv.Key, err)
			}
		case "LessonLimit":
			max := len(configModeCodes)
			sub := strings.SplitN(kv.Value, ",", max)
			if len(sub) != max {
				return fmt.Errorf("invalid lessons value: %q", kv.Value)
			}
			for i, s := range sub {
				v, err := strconv.Atoi(s)
				if err != nil {
					return fmt.Errorf("cannot parse %s: %w", kv.Key, err)
				}
				legacy.Sub_409FB0_settings(uint16(configModeCodes[i]), uint16(v))
			}
		case "TimeLimit":
			max := len(configModeCodes)
			sub := strings.SplitN(kv.Value, ",", max)
			if len(sub) != max {
				return fmt.Errorf("invalid time limit value: %q", kv.Value)
			}
			for i, s := range sub {
				v, err := strconv.Atoi(s)
				if err != nil {
					return fmt.Errorf("cannot parse %s: %w", kv.Key, err)
				}
				noxServer.Sub40A040settings(int(configModeCodes[i]), v)
			}
		case "PlayerSkeletons":
			v, err := strconv.Atoi(kv.Value)
			if err != nil {
				return fmt.Errorf("cannot parse %s: %w", kv.Key, err)
			}
			sst := getServerSettings()
			*(*uint32)(unsafe.Pointer(&sst.PlayerSkeletons58)) = uint32(bool2int(v > 0))
		case "BroadcastGestures":
			v, err := strconv.Atoi(kv.Value)
			if err != nil {
				return fmt.Errorf("cannot parse %s: %w", kv.Key, err)
			}
			sst := getServerSettings()
			*(*uint32)(unsafe.Pointer(&sst.BroadcastGestures62)) = uint32(bool2int(v > 0))
		case "LatencyCompensation":
			sub := strings.SplitN(kv.Value, " ", 2)
			if len(sub) != 2 {
				return fmt.Errorf("invalid lattency compensation value: %q", kv.Value)
			}
			n1, err := strconv.Atoi(sub[0])
			if err != nil {
				return fmt.Errorf("cannot parse %s: %w", kv.Key, err)
			}
			n2, err := strconv.Atoi(sub[1])
			if err != nil {
				return fmt.Errorf("cannot parse %s: %w", kv.Key, err)
			}
			sst := getServerSettings()
			*(*uint32)(unsafe.Pointer(&sst.LatencyCompensationA66)) = uint32(n1)
			*(*uint32)(unsafe.Pointer(&sst.LatencyCompensationB70)) = uint32(n2)
		case "Closed":
			v, err := strconv.Atoi(kv.Value)
			if err != nil {
				return fmt.Errorf("cannot parse %s: %w", kv.Key, err)
			}
			sst := getServerSettings()
			if v > 0 {
				sst.Flags100 |= server.ServerClosed
			} else {
				sst.Flags100 &^= server.ServerClosed
			}
		case "Private":
			v, err := strconv.Atoi(kv.Value)
			if err != nil {
				return fmt.Errorf("cannot parse %s: %w", kv.Key, err)
			}
			sst := getServerSettings()
			if v > 0 {
				sst.Flags100 |= server.ServerPrivate
			} else {
				sst.Flags100 &^= server.ServerPrivate
			}
		case "AudioThreshold":
			v, err := strconv.Atoi(kv.Value)
			if err != nil {
				return fmt.Errorf("cannot parse %s: %w", kv.Key, err)
			}
			if v < 0 {
				v = 0
			}
			if v > 100 {
				v = 100
			}
			sst := getServerSettings()
			*(*uint32)(unsafe.Pointer(&sst.AudioThreshold)) = uint32(v)
		case "MaxPlayers":
			v, err := strconv.Atoi(kv.Value)
			if err != nil {
				return fmt.Errorf("cannot parse %s: %w", kv.Key, err)
			}
			if v < 1 {
				v = 1
			}
			if v > common.MaxPlayers {
				v = common.MaxPlayers
			}
			sst := getServerSettings()
			sst.MaxPlayers104 = byte(v)
			legacy.Nox_xxx_servSetPlrLimit_409F80(v)
		case "RestrictedClasses":
			v, err := strconv.Atoi(kv.Value)
			if err != nil {
				return fmt.Errorf("cannot parse %s: %w", kv.Key, err)
			}
			sst := getServerSettings()
			sst.Flags100 |= server.ServerFlags(v) & server.ServerRestrictClasses
		case "RestrictedPing":
			sub := strings.SplitN(kv.Value, " ", 2)
			if len(sub) != 2 {
				return fmt.Errorf("invalid ping limit value: %q", kv.Value)
			}
			n1, err := strconv.Atoi(sub[0])
			if err != nil {
				return fmt.Errorf("cannot parse %s: %w", kv.Key, err)
			}
			n2, err := strconv.Atoi(sub[1])
			if err != nil {
				return fmt.Errorf("cannot parse %s: %w", kv.Key, err)
			}
			sst := getServerSettings()
			*(*uint16)(unsafe.Pointer(&sst.RestrictedPingMind105)) = uint16(n1)
			*(*uint16)(unsafe.Pointer(&sst.RestrictedPingMax107)) = uint16(n2)
		case "LimitMaxRes":
			v, err := strconv.Atoi(kv.Value)
			if err != nil {
				return fmt.Errorf("cannot parse %s: %w", kv.Key, err)
			}
			sst := getServerSettings()
			if v > 0 {
				sst.Flags102 |= server.ServerLimitMaxRes
			} else {
				sst.Flags102 &^= server.ServerLimitMaxRes
			}
		case "ItemRespawn":
			v, err := strconv.Atoi(kv.Value)
			if err != nil {
				return fmt.Errorf("cannot parse %s: %w", kv.Key, err)
			}
			if v != 0 {
				legacy.Sub_409E70(2)
			} else {
				legacy.Sub_409EC0(2)
			}
		case "CamperAlarm":
			v, err := strconv.Atoi(kv.Value)
			if err != nil {
				return fmt.Errorf("cannot parse %s: %w", kv.Key, err)
			}
			if v != 0 {
				legacy.Sub_409E70(0x2000)
			} else {
				legacy.Sub_409EC0(0x2000)
			}
		case "MinKickVotes":
			v, err := strconv.Atoi(kv.Value)
			if err != nil {
				return fmt.Errorf("cannot parse %s: %w", kv.Key, err)
			}
			*memmap.PtrUint32(0x587000, 229980) = uint32(v)
		case "ResetQuestMinVotes":
			v, err := strconv.Atoi(kv.Value)
			if err != nil {
				return fmt.Errorf("cannot parse %s: %w", kv.Key, err)
			}
			legacy.Set_nox_server_resetQuestMinVotes_229988(v)
		case "KickQuestPlayerMinVotes":
			v, err := strconv.Atoi(kv.Value)
			if err != nil {
				return fmt.Errorf("cannot parse %s: %w", kv.Key, err)
			}
			legacy.Set_nox_server_kickQuestPlayerMinVotes_229992(v)
		case "Fullscreen":
			v, err := strconv.Atoi(kv.Value)
			if err != nil {
				return fmt.Errorf("cannot parse %s: %w", kv.Key, err)
			}
			g_fullscreen_cfg = v
			if noxClient.GetWindowMode() <= -4 {
				noxClient.UpdateFullScreen(g_fullscreen_cfg)
			}
			changeWindowedOrFullscreen()
		case "Stretched":
			v, err := strconv.Atoi(kv.Value)
			if err != nil {
				return fmt.Errorf("cannot parse %s: %w", kv.Key, err)
			}
			g_scaled_cfg = v
			noxClient.setStretchIfNotSet(v != 0)
		case "InputSensitivity":
			v, err := strconv.ParseFloat(kv.Value, 32)
			if err != nil {
				return fmt.Errorf("cannot parse %s: %w", kv.Key, err)
			}
			noxClient.SetSensitivity(float32(v))
		default:
			for _, ci := range configCycleIndexes {
				if ci.Name == kv.Key {
					v, err := strconv.Atoi(kv.Value)
					if err != nil {
						return fmt.Errorf("cannot parse %s: %w", kv.Key, err)
					}
					legacy.Sub_4D0DC0(ci.Flags, v)
					break
				}
			}
		}
	}
	return nil
}

var configCycleIndexes = []struct {
	Name  string
	Flags uint32
}{
	{Name: "DeathmatchMapCycleIndex", Flags: 0x100},
	{Name: "HighlanderMapCycleIndex", Flags: 0x400},
	{Name: "CaptureMapCycleIndex", Flags: 0x20},
	{Name: "FlagballMapCycleIndex", Flags: 0x40},
	{Name: "KingOfTheRealmMapCylceIndex", Flags: 0x10},
	{Name: "QuestMapCycleIndex", Flags: 0x1000},
}

var configModeCodes = []uint32{
	0x100,
	0x400,
	0x20,
	0x40,
	0x10,
}

var configConnTypes = []struct {
	Name string
	Val  int
}{
	{Name: "Modem", Val: 0x4},
	{Name: "ISDN", Val: 0x3},
	{Name: "Cable/T1", Val: 0x1},
	{Name: "Undefined", Val: 0x0},
}

type VolumeControl int

const (
	VolumeMax = 0x4000
)

const (
	VolumeFX     = VolumeControl(0)
	VolumeDialog = VolumeControl(1)
	VolumeMusic  = VolumeControl(2)
)

func configParseVolume(val string, ind VolumeControl) error {
	v, err := strconv.ParseUint(val, 10, 32)
	if err != nil {
		return err
	}
	if v < 0 {
		v = 0
	} else if v > VolumeMax {
		v = VolumeMax
	}
	configSetVolume(int(v), ind)
	return nil
}

func configSetVolume(val int, ind VolumeControl) {
	switch ind {
	case VolumeFX:
		*memmap.PtrInt32(0x587000, 155048) = int32(val)
	case VolumeDialog:
		*memmap.PtrInt32(0x587000, 155052) = int32(val)
	case VolumeMusic:
		*memmap.PtrInt32(0x587000, 155056) = int32(val)
	}
}

func configGetVolume(ind VolumeControl) int {
	switch ind {
	case VolumeFX:
		return int(memmap.Int32(0x587000, 155048))
	case VolumeDialog:
		return int(memmap.Int32(0x587000, 155052))
	case VolumeMusic:
		return int(memmap.Int32(0x587000, 155056))
	}
	return VolumeMax
}

func configGetVolumeGain(ind VolumeControl) float32 {
	return float32(configGetVolume(ind)) / float32(VolumeMax)
}

func sub_432E50(val string, ind int) error {
	sub := strings.SplitN(val, ",", 12)
	if len(sub) > 12 {
		return fmt.Errorf("invalid filter values: %q", val)
	}
	arr, freeArr := alloc.Make([]int32{}, 12)
	defer freeArr()
	for i, s := range sub {
		v, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		arr[i] = int32(v)
	}
	legacy.Sub_489FF0(ind, int(arr[0]), unsafe.Pointer(&arr[1]))
	return nil
}

func nox_common_skipcfgfile_4331E0(sect cfg.Section, skip bool) error {
	for _, kv := range sect {
		if nox_client_parseConfigHotkeysLine(kv.Key, kv.Value) == 0 {
			return fmt.Errorf("cannot parse binding %s: %q", kv.Key, kv.Value)
		}
	}
	return nil
}

func writeConfigLegacy(path string) error {
	if env.IsE2E() {
		return nil
	}
	config.Global.Log.Info("writing legacy config", "path", path)
	cfile := &cfg.File{Sections: make([]cfg.Section, 2)}
	writeConfigLegacyMain(&cfile.Sections[0])
	writeConfigHotkeys(&cfile.Sections[1])
	f, err := ifs.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return cfile.WriteTo(f)
}

func writeConfigLegacyMain(sect *cfg.Section) {
	sst := getServerSettings()
	if _, ok := sect.Get("Version"); !ok {
		sect.Set("Version", "65540")
	}
	mode := noxClient.videoGetGameMode()
	sect.Set("VideoMode", fmt.Sprintf("%d %d 16", mode.X, mode.Y))
	sect.Set("Stretched", strconv.Itoa(g_scaled_cfg))
	sect.Set("Fullscreen", strconv.Itoa(g_fullscreen_cfg))
	sect.Set("VideoSize", strconv.Itoa(nox_video_getCutSize()))
	sect.Set("Gamma", strconv.Itoa(legacyGamma))
	if legacyGamma2Set {
		sect.Set("Gamma2", strconv.FormatFloat(legacyGamma2, 'g', -1, 32))
	}
	sect.Set("InputSensitivity", strconv.FormatFloat(float64(noxClient.GetSensitivity()), 'g', -1, 32))

	v := 0
	if legacy.Sub_453070() == 1 {
		v = int(*(*uint32)(unsafe.Add(unsafe.Pointer(legacy.Get_dword_587000_127004()), 4)) >> 16)
	}
	sect.Set("FXVolume", strconv.Itoa(v))

	v = 0
	if legacy.Sub_44D990() == 1 {
		v = int(*(*uint32)(unsafe.Add(unsafe.Pointer(legacy.Get_dword_587000_122852()), 4)) >> 16)
	}
	sect.Set("DialogVolume", strconv.Itoa(v))

	v = 0
	if legacy.Sub_43DC30() == 1 {
		v = int(*(*uint32)(unsafe.Add(unsafe.Pointer(legacy.Get_dword_587000_93164()), 4)) >> 16)
	}
	sect.Set("MusicVolume", strconv.Itoa(v))

	sect.Set("LastServer", alloc.GoString((*byte)(memmap.PtrOff(0x5D4594, 806060))))
	sect.Set("ServerName", sub_433890())
	sect.Set("UnlockSurface", strconv.Itoa(bool2int(legacyUnlockSurface)))
	sect.Set("SoftShadowEdge", strconv.Itoa(bool2int(noxflags.HasEngine(noxflags.EngineSoftShadowEdge))))
	sect.Set("DrawFrontWalls", strconv.Itoa(bool2int(nox_client_drawFrontWalls_80812)))
	sect.Set("TranslucentFrontWalls", strconv.Itoa(int(legacy.Get_nox_client_translucentFrontWalls_805844())))
	sect.Set("HighResFrontWalls", strconv.Itoa(int(legacy.Get_nox_client_highResFrontWalls_80820())))
	sect.Set("HighResFloors", strconv.Itoa(int(legacy.Get_nox_client_highResFloors_154952())))
	sect.Set("LockHighResFloors", strconv.Itoa(bool2int(nox_client_lockHighResFloors_1193152)))
	sect.Set("TexturedFloors", strconv.Itoa(bool2int(nox_client_texturedFloors_154956)))
	sect.Set("TranslucentConsole", strconv.Itoa(bool2int(guiCon.translucent)))
	sect.Set("RenderGlow", strconv.Itoa(bool2int(noxClient.r.Part.RenderGlow)))
	sect.Set("RenderGUI", strconv.Itoa(bool2int(nox_client_renderGUI_80828)))
	sect.Set("FadeObjects", strconv.Itoa(int(legacy.Get_nox_client_fadeObjects_80836())))
	sect.Set("RenderBubbles", strconv.Itoa(int(legacy.Get_nox_client_renderBubbles_80844())))
	sect.Set("TrackData", strconv.Itoa(int(sub_578DF0())))
	sect.Set("SysopPassword", legacy.Nox_xxx_sysopGetPass_40A630())
	sect.Set("ServerPassword", alloc.GoString16S(sst.ServerPass78[:]))
	sect.Set("Profiled", strconv.Itoa(bool2int(legacyProfiled != 0)))
	sect.Set("SanctuaryHelp", strconv.Itoa(int(legacy.Get_nox_server_sanctuaryHelp_54276())))
	sect.Set("MaxPacketLossPct", strconv.Itoa(int(memmap.Uint32(0x587000, 81280))))
	sect.Set("SendMessageOfTheDay", strconv.Itoa(int(legacy.Get_nox_server_sendMotd_108752())))
	sect.Set("MapCycle", strconv.Itoa(legacy.Sub_4D0D70()))
	sect.Set("ConnectionType", strconv.Itoa(int(legacy.Get_nox_server_connectionType_3596())))
	sect.Set("InternetUpdateRate", strconv.Itoa(nox_xxx_rateGet_40A6C0()))
	sect.Set("LessonLimit", sub_4337B0())
	sect.Set("TimeLimit", sub_433820())
	sect.Set("PlayerSkeletons", strconv.Itoa(int(*(*uint32)(unsafe.Pointer(&sst.PlayerSkeletons58)))))
	sect.Set("BroadcastGestures", strconv.Itoa(int(*(*uint32)(unsafe.Pointer(&sst.BroadcastGestures62)))))
	sect.Set("LatencyCompensation", fmt.Sprintf("%d %d", *(*uint32)(unsafe.Pointer(&sst.LatencyCompensationA66)), *(*uint32)(unsafe.Pointer(&sst.LatencyCompensationB70))))
	sect.Set("Closed", strconv.Itoa(bool2int(sst.Flags100.Has(server.ServerClosed))))
	sect.Set("Private", strconv.Itoa(bool2int(sst.Flags100.Has(server.ServerPrivate))))
	sect.Set("AudioThreshold", strconv.Itoa(int(*(*uint32)(unsafe.Pointer(&sst.AudioThreshold)))))
	sect.Set("MaxPlayers", strconv.Itoa(noxServer.getServerMaxPlayers()))
	sect.Set("RestrictedClasses", strconv.Itoa(int(sst.Flags100&server.ServerRestrictClasses)))
	sect.Set("RestrictedPing", fmt.Sprintf("%d %d", *(*uint16)(unsafe.Pointer(&sst.RestrictedPingMind105)), *(*uint16)(unsafe.Pointer(&sst.RestrictedPingMax107))))
	sect.Set("LimitMaxRes", strconv.Itoa(bool2int(sst.Flags102.Has(server.ServerLimitMaxRes))))
	sect.Set("CamperAlarm", strconv.Itoa(bool2int((legacy.Nox_xxx_getServerSubFlags_409E60()>>13)&0x1 != 0)))
	sect.Set("ItemRespawn", strconv.Itoa(bool2int(sub_409F40(2))))
	sect.Set("MinKickVotes", strconv.Itoa(int(memmap.Uint32(0x587000, 229980))))
	sect.Set("ResetQuestMinVotes", strconv.Itoa(int(legacy.Get_nox_server_resetQuestMinVotes_229988())))
	sect.Set("KickQuestPlayerMinVotes", strconv.Itoa(int(legacy.Get_nox_server_kickQuestPlayerMinVotes_229992())))

	v16, arr := sub_48A020(0)
	sect.Set("LANFilters", fmt.Sprintf("%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d",
		v16, arr[0], arr[1], arr[2], arr[3], arr[4], arr[5], arr[6], arr[7], arr[8], arr[9], arr[10],
	))
	v16, arr = sub_48A020(1)
	sect.Set("INETFilters", fmt.Sprintf("%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d",
		int(v16), arr[0], arr[1], arr[2], arr[3], arr[4], arr[5], arr[6], arr[7], arr[8], arr[9], arr[10],
	))
	for _, ci := range configCycleIndexes {
		sect.Set(ci.Name, strconv.Itoa(legacy.Sub_4D0DE0(ci.Flags)))
	}
}

func sub_48A020(ind int) (int, []int32) {
	v := int(memmap.Uint32(0x5D4594, 1193372+4*uintptr(ind)))
	arr := unsafe.Slice((*int32)(memmap.PtrOff(0x5D4594, 1193388+44*uintptr(ind))), 11)
	return v, arr
}

func sub_4337B0() string {
	var arr []string
	for _, m := range configModeCodes {
		arr = append(arr, strconv.Itoa(legacy.Nox_xxx_servGamedataGet_40A020(uint16(m))))
	}
	return strings.Join(arr, ",")
}

func sub_433820() string {
	var arr []string
	for _, m := range configModeCodes {
		arr = append(arr, strconv.Itoa(legacy.Sub_40A180(noxflags.GameFlag(m))))
	}
	return strings.Join(arr, ",")
}

func sub_433890() string {
	name := noxServer.getServerName()
	name = strings.ReplaceAll(name, " ", "_")
	return name
}
