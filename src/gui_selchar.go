package opennox

import (
	"encoding/binary"
	"fmt"
	"image"
	"io"
	"path/filepath"
	"sort"
	"strings"
	"time"
	"unsafe"

	crypt "github.com/noxworld-dev/noxcrypt"
	"github.com/noxworld-dev/opennox-lib/common"
	"github.com/noxworld-dev/opennox-lib/datapath"
	"github.com/noxworld-dev/opennox-lib/ifs"
	"github.com/noxworld-dev/opennox-lib/log"
	"github.com/noxworld-dev/opennox-lib/player"
	"github.com/noxworld-dev/opennox-lib/strman"

	"github.com/noxworld-dev/opennox/v1/client"
	"github.com/noxworld-dev/opennox/v1/client/gui"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"github.com/noxworld-dev/opennox/v1/common/sound"
	"github.com/noxworld-dev/opennox/v1/internal/cryptfile"
	"github.com/noxworld-dev/opennox/v1/legacy"
	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc"
	"github.com/noxworld-dev/opennox/v1/legacy/common/ccall"
	"github.com/noxworld-dev/opennox/v1/legacy/dialog"
	"github.com/noxworld-dev/opennox/v1/server"
)

const NOX_SAVEGAME_XXX_MAX = 14

var (
	saveLog                   = log.New("save")
	nox_xxx_saves_arr         []legacy.Nox_savegame_xxx
	nox_savegame_arr_1064948  [NOX_SAVEGAME_XXX_MAX]legacy.Nox_savegame_xxx
	nox_savegame_name_1307752 string
	nox_wnd_xxx_1307748       *gui.Anim
	winSelSave                *gui.Window
	winCharList               *gui.Window
	winCharListNames          *gui.Window
	winCharListStyle          *gui.Window
	dword_5d4594_1082856      *gui.Window
	dword_5d4594_1082860      *gui.Window
	dword_5d4594_1082864      *gui.Window
	dword_5d4594_1082868      *gui.Window
	dword_5d4594_1082880      *gui.Window
)

func sub_4A1BE0(a1 int) int {
	return legacy.Nox_xxx_wnd_46ABB0(guiOptsBack, a1)
}

func sub_41E300(a1 int) int {
	if a1 == 11 && noxflags.HasGame(noxflags.GameFlag29) {
		if legacy.Get_dword_5d4594_528256() != 0 {
			legacy.Sub_41E470()
		} else {
			legacy.Nox_xxx_reconStart_41E400()
		}
		return 0
	}
	if legacy.Get_dword_5d4594_527988() != a1 {
		legacy.Set_dword_5d4594_527988(a1)
		legacy.Sub_41DA10(a1)
		if legacy.Get_dword_5d4594_527988() == 11 {
			sub_40E0A0()
		}
	}
	return 1
}

func sub_40E0A0() {
	legacy.Set_dword_5d4594_10984(1)
}

func sub_4A50A0() int {
	nox_wnd_xxx_1307748.SetState(gui.AnimOut)
	gui.SetAnimGlobalState(gui.AnimOut)
	clientPlaySoundSpecial(sound.SoundShellSlideOut, 100)
	return 1
}

func sub_4A50D0() int {
	v0 := nox_wnd_xxx_1307748.Func13Ptr
	nox_wnd_xxx_1307748.Free()
	nox_wnd_xxx_1307748 = nil
	winSelSave.Destroy()
	winSelSave = nil
	if v0 != nil {
		ccall.CallIntVoid(unsafe.Pointer(v0))
		return 1
	}
	nox_client_resetScreenParticles_431510()
	noxClient.GUI.Draw()
	if !noxflags.HasGame(noxflags.GameOnline) {
		nox_client_guiXxxDestroy_4A24A0()
		return 1
	}
	if true {
		nox_xxx_serverHost_43B4D0()
		return 1
	}
	if false {
		sub_43B670()
		return 1
	}
	return 1
}

func nox_xxx_particlesLoadColor_4313E0() {
	*memmap.PtrUint32(0x5D4594, 806012) = nox_color_rgb_4344A0(255, 255, 200)
	*memmap.PtrUint32(0x5D4594, 806004) = nox_color_rgb_4344A0(50, 150, 255)
	*memmap.PtrUint32(0x5D4594, 806036) = nox_color_rgb_4344A0(200, 255, 255)
	*memmap.PtrUint32(0x5D4594, 806024) = nox_color_rgb_4344A0(255, 200, 100)
	*memmap.PtrUint32(0x5D4594, 806040) = nox_color_rgb_4344A0(200, 255, 200)
	*memmap.PtrUint32(0x5D4594, 806020) = nox_color_rgb_4344A0(255, 255, 0)
	*memmap.PtrUint32(0x5D4594, 806028) = nox_color_rgb_4344A0(0, 0, 255)
	*memmap.PtrUint32(0x5D4594, 806016) = nox_color_rgb_4344A0(128, 128, 255)
	*memmap.PtrUint32(0x5D4594, 806008) = nox_color_rgb_4344A0(255, 128, 64)
	*memmap.PtrUint32(0x5D4594, 806032) = nox_color_rgb_4344A0(0, 255, 0)
}

func nox_client_initScreenParticles_431390() {
	if legacy.Get_nox_alloc_screenParticles_806044() == nil {
		legacy.Set_nox_alloc_screenParticles_806044(alloc.NewClass("ScreenParticles", unsafe.Sizeof(legacy.Nox_screenParticle{}), 2000).UPtr())
		if legacy.Get_nox_alloc_screenParticles_806044() == nil {
			panic("cannot init particles")
		}
	}
	nox_client_resetScreenParticles_431510()
}

func nox_xxx_freeScreenParticles_4314D0() {
	if legacy.Get_nox_alloc_screenParticles_806044() != nil {
		alloc.AsClass(legacy.Get_nox_alloc_screenParticles_806044()).Free()
	}
	legacy.Set_nox_alloc_screenParticles_806044(nil)
	legacy.Set_nox_screenParticles_head(nil)
	legacy.Set_dword_5d4594_806052(nil)
}

func nox_client_resetScreenParticles_431510() {
	if legacy.Get_nox_alloc_screenParticles_806044() != nil {
		alloc.AsClass(legacy.Get_nox_alloc_screenParticles_806044()).FreeAllObjects()
	}
	legacy.Set_nox_screenParticles_head(nil)
	legacy.Set_dword_5d4594_806052(nil)
	nox_xxx_particlesLoadColor_4313E0()
}

func sub_43B670() {
	sst := getServerSettings()
	if legacy.Get_nox_game_createOrJoin_815048() != 0 {
		sub_44A400()
		legacy.Nox_client_xxx_switchChatMap_43B510()
		legacy.Nox_client_guiXxx_43A9D0()
		nox_client_guiXxxDestroy_4A24A0()
		noxServer.ai.nox_xxx_gameSetAudioFadeoutMb(int(*(*uint32)(unsafe.Pointer(&sst.AudioThreshold))))
	} else {
		legacy.Sub_43B630()
	}
	legacy.Sub_49FF20()
}

func sub_46D6F0() int {
	nox_savegame_arr_1064948 = [NOX_SAVEGAME_XXX_MAX]legacy.Nox_savegame_xxx{}
	win := dword_5d4594_1082856
	if win.GetFlags().IsHidden() {
		return 0
	}
	win.Hide()
	legacy.Nox_xxx_wnd_46ABB0(win, 0)
	win.Capture(false)
	sub_413A00(0)
	return 1
}

func sub_413A00(a1 int) {
	if noxflags.HasGame(noxflags.GameModeCoop) {
		if a1 != 0 {
			noxflags.SetGame(noxflags.GamePause)
		} else {
			if legacy.Get_dword_5d4594_251744() == 0 {
				noxflags.UnsetGame(noxflags.GamePause)
				nox_ticks_reset_416D40()
			}
		}
	}
}

func sub_477530(a1 bool) {
	if a1 {
		*memmap.PtrUint32(0x5D4594, 1096652) = uint32(legacy.Nox_xxx_bookHideMB_45ACA0(0))
		*memmap.PtrUint32(0x5D4594, 1096664) = uint32(legacy.Sub_467C10())
		*memmap.PtrUint32(0x5D4594, 1096660) = uint32(legacy.Nox_xxx_guiServerOptionsTryHide_4574D0())
		nox_client_renderGUI_80828 = false
		*memmap.PtrUint32(0x5D4594, 1096656) = uint32(sub_473670())
		if *memmap.PtrUint32(0x5D4594, 1096656) == 1 {
			nox_client_toggleMap_473610()
		}
		guiCon.Hide()
		legacy.Sub_446780()
		legacy.Sub_4AD9B0(1)
		legacy.Sub_4C35B0(1)
		sub_46D6F0()
		legacy.Sub_460660()
		*memmap.PtrUint32(0x5D4594, 1096672) = 1
	} else {
		*memmap.PtrUint32(0x5D4594, 1096672) = 0
		if *memmap.PtrUint32(0x5D4594, 1096652) != 0 {
			legacy.Nox_xxx_bookShowMB_45AD70(0)
		}
		if *memmap.PtrUint32(0x5D4594, 1096664) != 0 {
			legacy.Sub_467BB0()
		}
		if *memmap.PtrUint32(0x5D4594, 1096660) != 0 {
			legacy.Nox_xxx_guiServerOptsLoad_457500()
		}
		if *memmap.PtrUint32(0x5D4594, 1096656) != 0 {
			nox_client_toggleMap_473610()
		}
		if nox_xxx_xxxRenderGUI_587000_80832 {
			nox_client_renderGUI_80828 = true
		}
	}
}

func sub_44A400() {
	c := noxClient
	if win1 := nox_gui_curDialog_830224; win1 != nil {
		if legacy.Get_dword_5d4594_830236() != nil {
			legacy.Get_dword_5d4594_830236().Capture(true)
			legacy.Set_dword_5d4594_830236(nil)
		}
		win1.StackPop()
		win1.Capture(false)
		win1.Destroy()
		if legacy.Get_dword_5d4594_830232() != nil {
			win1.GUI().Focus(legacy.Get_dword_5d4594_830232())
		}
		nox_gui_curDialog_830224 = nil
		dword_5d4594_830228 = nil
		c.Nox_client_setCursorType(gui.CursorSelect)
		c.SetMouseBounds(image.Rect(0, 0, nox_win_width-1, nox_win_height-1))
	}
}

func nox_game_showSelChar_4A4DB0() int {
	if nox_game_showSelChar4A4DB0() {
		return 1
	}
	return 0
}

func nox_game_showSelChar4A4DB0() bool {
	c := noxClient
	c.Nox_client_setCursorType(gui.CursorSelect)
	if sub4D6F30() {
		sub_4D6F90(1)
	}
	c.GameAddStateCode(client.StateCharSelect)
	sub_4A1BE0(1)
	win := newWindowFromFile(c.GUI, "selchar.wnd", nox_xxx_windowSelCharProc_4A5710)
	winSelSave = win
	if win == nil {
		return false
	}
	win.SetFunc93(gui.WrapFuncC(legacy.Get_sub_4A18E0()))
	anim := nox_gui_makeAnimation(win, 0, 0, 0, -440, 0, 20, 0, -40)
	nox_wnd_xxx_1307748 = anim
	if anim == nil {
		return false
	}
	anim.StateID = client.StateCharSelect
	anim.Func12Ptr = legacy.Get_sub_4A50A0()
	anim.FncDoneOutPtr = legacy.Get_sub_4A50D0()
	wlist := win.ChildByID(510)
	wnames := win.ChildByID(511)
	wstyle := win.ChildByID(512)
	winCharList = wlist
	winCharListNames = wnames
	winCharListStyle = wstyle
	wlist.SetFunc94(nox_xxx_windowSelCharProc_4A5710)
	wnames.SetParent(wlist)
	wstyle.SetParent(wlist)
	if noxflags.HasGame(noxflags.GameModeCoop) {
		saveLog.Printf("save win mode: coop")
		wup := win.ChildByID(504)
		wup.Hide()

		wdown := win.ChildByID(505)
		wdown.Hide()

		wblank := win.ChildByID(509)
		v8 := c.Strings().GetStringInFile("LoadLabel", "C:\\NoxPost\\src\\client\\shell\\selchar.c")
		sub_46AEE0(wblank, v8)
	} else {
		saveLog.Printf("save win mode: char")
		wup := win.ChildByID(504)
		wlist.Func94(&WindowEvent0x4018{Win: wup})
		wnames.Func94(&WindowEvent0x4018{Win: wup})
		wstyle.Func94(&WindowEvent0x4018{Win: wup})

		wdown := win.ChildByID(505)
		wlist.Func94(&WindowEvent0x4019{Win: wdown})
		wnames.Func94(&WindowEvent0x4019{Win: wdown})
		wstyle.Func94(&WindowEvent0x4019{Win: wdown})

		wblank := win.ChildByID(509)
		v4 := c.Strings().GetStringInFile("LoadCharLabel", "C:\\NoxPost\\src\\client\\shell\\selchar.c")
		sub_46AEE0(wblank, v4)
	}
	nox_xxx_findAutosaves_4A5150()
	guiSetBackButtonText("OptsBack.wnd:Back")
	if noxflags.HasGame(noxflags.GameOnline) {
		noxflags.SetEngine(noxflags.EngineAdmin)
	} else {
		noxflags.UnsetEngine(noxflags.EngineAdmin)
		legacy.Sub_4602F0()
	}
	if noxflags.HasGame(noxflags.GameFlag26) {
		wload := win.ChildByID(502)
		win.Func94(&WindowEvent0x4007{Win: wload})
	}
	return true
}

func nox_xxx_findAutosaves_4A5150() {
	wlist := winCharList
	wnames := winCharListNames
	wstyle := winCharListStyle
	v1 := winSelSave.ChildByID(501)
	legacy.Nox_xxx_wnd_46ABB0(v1, 1)
	PathName := datapath.Save()
	ifs.Mkdir(PathName)
	if noxflags.HasGame(noxflags.GameModeCoop) {
		saveLog.Printf("scanning path for coop saves: %q", PathName)
		nox_xxx_saves_arr, _ = alloc.Make([]legacy.Nox_savegame_xxx{}, NOX_SAVEGAME_XXX_MAX)
		nox_savegame_sub_46CE40(wlist, wnames, wstyle, nox_xxx_saves_arr)
		return
	}
	saveLog.Printf("scanning path for saves: %q", PathName)
	ifs.Chdir(PathName)
	files, _ := ifs.ReadDir(PathName)
	v0 := 0
	for _, f := range files {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".plr") {
			v0++
		}
	}
	if v0 != 0 {
		nox_xxx_saves_arr, _ = alloc.Make([]legacy.Nox_savegame_xxx{}, v0)
	} else {
		nox_xxx_saves_arr = nil
	}
	v9 := 0
	for _, f := range files {
		if f.IsDir() || !strings.HasSuffix(f.Name(), ".plr") {
			continue
		}
		s := filepath.Join(PathName, f.Name())
		sv := &nox_xxx_saves_arr[v9]
		if legacy.Sub_41A000(s, sv) == 0 {
			saveLog.Printf("%q: failed", f.Name())
		}
		if sub_4A5690(sv) == 1 {
			v9++
		}
	}
	sort.Slice(nox_xxx_saves_arr[:v9], func(i, j int) bool {
		sv1, sv2 := &nox_xxx_saves_arr[i], &nox_xxx_saves_arr[j]
		return sv1.Timestamp.Time().After(sv2.Timestamp.Time())
	})
	for i := 0; i < v9; i++ {
		sv := &nox_xxx_saves_arr[i]
		v28 := fmt.Sprintf("nox.str:%s", saveClasses[sv.Player_class])
		v26 := strMan.GetStringInFile(strman.ID(v28), "C:\\NoxPost\\src\\client\\shell\\selchar.c")
		v24 := alloc.GoString16S(sv.Player_name[:])
		v14 := strMan.GetStringInFile("GuiInv.c:ElaborateNameFormat", "C:\\NoxPost\\src\\client\\shell\\selchar.c")
		v30 := fmt.Sprintf(v14, v24, v26)
		if sv.Flags&0x4 != 0 {
			v15 := int(sv.Stage)
			if v15 < 1 {
				v15 = 1
			}
			v30 += " - "
			v30 += strMan.GetStringInFile("Noxworld.c:Stage", "C:\\NoxPost\\src\\client\\shell\\selchar.c")
			v30 += fmt.Sprintf(" %d", v15)
		}
		wlist.Func94(&WindowEvent0x400d{Str: " ", Val: 14})
		wstyle.Func94(&WindowEvent0x400d{Str: v30, Val: 14})
		var v25 string
		if sv.Flags&0x4 != 0 {
			v25 = strMan.GetStringInFile("Noxworld.c:Quest", "C:\\NoxPost\\src\\client\\shell\\selchar.c")
		} else {
			v25 = strMan.GetStringInFile("Multiplayer", "C:\\NoxPost\\src\\client\\shell\\selchar.c")
		}
		wnames.Func94(&WindowEvent0x400d{Str: v25, Val: 14})
	}
	ifs.Chdir(datapath.Data())
	if v9 != 0 {
		v21 := winSelSave.ChildByID(503)
		v22 := winSelSave.ChildByID(502)
		v20 := winSelSave.ChildByID(501)
		legacy.Nox_xxx_wnd_46ABB0(v21, 1)
		legacy.Nox_xxx_wnd_46ABB0(v22, 1)
		if (noxServer.nox_xxx_isQuest_4D6F50() || sub_4D6F70()) && nox_client_countPlayerFiles04_4DC7D0() >= 14 {
			legacy.Nox_xxx_wnd_46ABB0(v20, 0)
		} else {
			legacy.Nox_xxx_wnd_46ABB0(v20, 1)
		}
	} else {
		v21 := winSelSave.ChildByID(503)
		v22 := winSelSave.ChildByID(502)
		v23 := winSelSave.ChildByID(501)
		legacy.Nox_xxx_wnd_46ABB0(v21, 0)
		legacy.Nox_xxx_wnd_46ABB0(v22, 0)
		legacy.Nox_xxx_wnd_46ABB0(v23, 1)
	}
	if int32(*(*uint16)(unsafe.Add(wnames.WidgetData, 46))) != 0 { // TODO: check widget type
		wnames.Func94(gui.AsWindowEvent(0x4013, 0, 0))
		wstyle.Func94(gui.AsWindowEvent(0x4013, 0, 0))
	}
}

func sub_4A5690(sv *legacy.Nox_savegame_xxx) int {
	if sv.Flags&4 == 0 {
		if !noxServer.nox_xxx_isQuest_4D6F50() && !sub_4D6F70() {
			return 1
		}
		if sv.Flags&4 == 0 {
			return 0
		}
	}
	if noxServer.nox_xxx_isQuest_4D6F50() || sub_4D6F70() {
		return 1
	}
	return 0
}

func sub_4A5C70() {
	if noxflags.HasGame(noxflags.GameModeCoop) {
		deleteSaveDir(nox_savegame_name_1307752, true)
	} else {
		ind := memmap.Uint32(0x5D4594, 1307772)
		path := alloc.GoStringS(nox_xxx_saves_arr[ind].Path[:])
		ifs.Remove(path)
	}
	winCharList.Func94(gui.AsWindowEvent(0x400F, 0, 0))
	winCharListNames.Func94(gui.AsWindowEvent(0x400F, 0, 0))
	winCharListStyle.Func94(gui.AsWindowEvent(0x400F, 0, 0))
	nox_xxx_findAutosaves_4A5150()
}

var saveClasses = []string{
	"Warrior",
	"Wizard",
	"Conjurer",
}

func nox_savegame_sub_46CE40(wlist, wnames, wstyles *gui.Window, sarr []legacy.Nox_savegame_xxx) int {
	const stringsFile = "C:\\NoxPost\\src\\client\\Gui\\GUISave.c"
	PathName := datapath.Save()
	ifs.Mkdir(PathName)
	wlist.Func94(gui.AsWindowEvent(0x400F, 0, 0))
	wnames.Func94(gui.AsWindowEvent(0x400F, 0, 0))
	wstyles.Func94(gui.AsWindowEvent(0x400F, 0, 0))
	v45 := filepath.Join(PathName, common.SaveAuto, common.PlayerFile)
	var v4 int32 = 0
	if legacy.Sub_41A000(v45, &sarr[0]) != 0 {
		v4++
	}
	for i := int32(1); i < NOX_SAVEGAME_XXX_MAX; i++ {
		v45 = filepath.Join(PathName, fmt.Sprintf(common.SaveFormat, i), common.PlayerFile)
		if legacy.Sub_41A000(v45, &sarr[i]) != 0 {
			v4++
		}
	}

	swin := dword_5d4594_1082856
	for v9 := 0; v9 < NOX_SAVEGAME_XXX_MAX; v9++ {
		sv := &sarr[v9]
		spath := alloc.GoStringS(sv.Path[:])
		if spath == "" {
			wlist.Func94(&WindowEvent0x400d{Str: " ", Val: 3})
			str := " "
			if v9 != 0 {
				str = strMan.GetStringInFile("EmptySlot", stringsFile)
			}
			wnames.Func94(&WindowEvent0x400d{Str: str, Val: 3})
			wstyles.Func94(&WindowEvent0x400d{Str: str, Val: 3})
			continue
		}
		playerName := alloc.GoString16S(sv.Player_name[:])
		sname := datapath.SaveNameFromPath(spath)
		class := int(sv.Player_class)
		classNameID := fmt.Sprintf("nox.str:%s", saveClasses[class])
		v39 := strMan.GetStringInFile(strman.ID(classNameID), stringsFile)
		v13 := strMan.GetStringInFile("the", stringsFile)
		v42 := fmt.Sprintf(v13, playerName, v39)
		if sname == common.SaveAuto {
			v42 += " "
			v14 := strMan.GetStringInFile("AutoSaveLabel", stringsFile)
			v42 += v14
		} else {
			v42 += " - "
			var v17 string
			switch sv.Stage {
			case 1:
				if class == 0 {
					v17 = strMan.GetStringInFile("WolChat.c:Nox_Lobby_3", stringsFile)
				} else if class != 1 {
					v17 = strMan.GetStringInFile("Noxworld.wnd:VillageIx", stringsFile)
				} else {
					v17 = strMan.GetStringInFile("NPC:Horvath", stringsFile)
				}
			case 2:
				if class == 0 {
					v17 = strMan.GetStringInFile("War02a.scr:Sign5", stringsFile)
				} else if class != 1 {
					v17 = strMan.GetStringInFile("NPC:Aldwyn", stringsFile)
				} else {
					v17 = strMan.GetStringInFile("Noxworld.wnd:Galava", stringsFile)
				}
			case 3:
				if class == 0 {
					v17 = strMan.GetStringInFile("Noxworld.wnd:VillageIx", stringsFile)
				} else if class != 1 {
					v17 = strMan.GetStringInFile("Noxworld.wnd:Mines", stringsFile)
				} else {
					v17 = strMan.GetStringInFile("WolChat.c:Nox_Lobby_9", stringsFile)
				}
			case 4:
				v17 = strMan.GetStringInFile("Noxworld.wnd:FieldValor", stringsFile)
			case 5:
				v17 = strMan.GetStringInFile("Noxworld.wnd:HamletBrin", stringsFile)
			case 6:
				v17 = strMan.GetStringInFile("NPC:Horrendous", stringsFile)
			case 7:
				v17 = strMan.GetStringInFile("Noxworld.wnd:Galava", stringsFile)
			case 8:
				v17 = strMan.GetStringInFile("Noxworld.wnd:TempleIx", stringsFile)
			case 9:
				v17 = strMan.GetStringInFile("Noxworld.wnd:DismalSwamp", stringsFile)
			case 10:
				v17 = strMan.GetStringInFile("Noxworld.wnd:LandDead", stringsFile)
			case 11:
				if class == 0 {
					v17 = strMan.GetStringInFile("WolChat.c:Nox_Lobby_47", stringsFile)
				} else if class == 1 {
					v17 = strMan.GetStringInFile("NPC:Hecubah", stringsFile)
				} else {
					v17 = strMan.GetStringInFile("Noxworld.wnd:GrokkTorr", stringsFile)
				}
			default:
				v17 = strMan.GetStringInFile("GuiInv.c:IdentifyUnknown", stringsFile)
			}
			v42 += v17
		}
		v21 := sub46CD70(sv)
		stime := sv.Timestamp.Time().Format("06-01-02, 15:04 ")
		wlist.Func94(&WindowEvent0x400d{Str: " ", Val: v21})
		wstyles.Func94(&WindowEvent0x400d{Str: v42, Val: v21})
		wnames.Func94(&WindowEvent0x400d{Str: stime, Val: v21})
	}
	v22 := swin.ChildByID(503)
	v24 := swin.ChildByID(502)
	if v4 <= 0 {
		legacy.Nox_xxx_wnd_46ABB0(v22, 0)
		legacy.Nox_xxx_wnd_46ABB0(v24, 0)
	} else {
		legacy.Nox_xxx_wnd_46ABB0(v22, 1)
		legacy.Nox_xxx_wnd_46ABB0(v24, 1)
	}
	ind := nox_savegame_findLatestSave_46CDC0(sarr)
	if ind == -1 {
		return -1
	}
	ev2 := gui.AsWindowEvent(0x4013, uintptr(ind), 0)
	wlist.Func94(ev2)
	wnames.Func94(ev2)
	wstyles.Func94(ev2)
	return ind
}

func sub46CD70(sv *legacy.Nox_savegame_xxx) int {
	if sv.Flags&0x8 != 0 {
		return 10
	}
	buf := datapath.SaveNameFromPath(alloc.GoStringS(sv.Path[:]))
	return bool2int(buf != common.SaveAuto) + (NOX_SAVEGAME_XXX_MAX - 1)
}

func nox_savegame_findLatestSave_46CDC0(sarr []legacy.Nox_savegame_xxx) int {
	var (
		ind    = -1
		latest time.Time
	)
	for i := 0; i < len(sarr); i++ {
		sv := &sarr[i]
		spath := alloc.GoStringS(sv.Path[:])
		if spath == "" {
			continue
		}
		t := sv.Timestamp.Time()
		if ind == -1 || t.After(latest) {
			ind = i
			latest = t
		}
	}
	return ind
}

func sub_4D6F70() bool {
	return memmap.Uint32(0x5D4594, 1556164) != 0
}

func sub_450560() bool {
	return legacy.Get_dword_5d4594_831260() != 0
}

func sub_450570() bool {
	return legacy.Get_dword_5d4594_831220() != 0
}

func nox_savegame_sub_46D580() {
	c := noxClient
	if noxflags.HasGame(noxflags.GameModeQuest) {
		return
	}
	nox_savegame_arr_1064948 = [NOX_SAVEGAME_XXX_MAX]legacy.Nox_savegame_xxx{}
	nox_savegame_sub_46CE40(dword_5d4594_1082860, dword_5d4594_1082864, dword_5d4594_1082868, nox_savegame_arr_1064948[:])
	win1 := dword_5d4594_1082856
	win1.ShowModal()
	legacy.Nox_xxx_wnd_46ABB0(win1, 1)
	win1.Capture(true)
	v1 := win1.ChildByID(501)
	if sub_450560() && !sub_450570() {
		legacy.Nox_xxx_wnd_46ABB0(v1, 0)
	} else {
		switch c.ClientPlayerUnit().AnimInd {
		case 1, 2, 51:
			legacy.Nox_xxx_wnd_46ABB0(v1, 0)
		default:
			legacy.Nox_xxx_wnd_46ABB0(v1, 1)
		}
	}
	var (
		v4 *gui.Window
		v5 string
	)
	if sub_450560() {
		v3 := strMan.GetStringInFile("SaveGUIQuit", "C:\\NoxPost\\src\\client\\Gui\\GUISave.c")
		sub_46AEE0(dword_5d4594_1082880, v3)
		v4 = win1.ChildByID(509)
		v5 = strMan.GetStringInFile("LoadLabel", "C:\\NoxPost\\src\\client\\Gui\\GUISave.c")
	} else {
		v6 := strMan.GetStringInFile("SaveGUICancel", "C:\\NoxPost\\src\\client\\Gui\\GUISave.c")
		sub_46AEE0(dword_5d4594_1082880, v6)
		v4 = win1.ChildByID(509)
		v5 = strMan.GetStringInFile("SaveLoadLabel", "C:\\NoxPost\\src\\client\\Gui\\GUISave.c")
	}
	sub_46AEE0(v4, v5)
	sub_413A00(1)
}

func nox_xxx_windowSelCharProc_4A5710(a1 *gui.Window, e gui.WindowEvent) gui.WindowEventResp {
	switch e := e.(type) {
	case gui.WindowDestroy:
		if e.Targ.ID() == 500 {
			nox_xxx_saves_arr = nil
		}
		return nil
	case *WindowEvent0x4005:
		clientPlaySoundSpecial(sound.SoundShellSelect, 100)
		return gui.RawEventResp(1)
	case *WindowEvent0x4010:
		ev2 := gui.AsWindowEvent(0x4013, e.Val, 0)
		winCharList.Func94(ev2)
		winCharListNames.Func94(ev2)
		winCharListStyle.Func94(ev2)
		return nil
	case *WindowEvent0x4007:
		switch e.Win.ID() {
		case 501:
			noxServer.Objs.SetFirstObjectScriptID(1000000000)
			sub_4A50A0()
			nox_wnd_xxx_1307748.Func13Ptr = legacy.Get_nox_game_showSelClass_4A4840()
		case 502:
			v7 := winCharListNames.WidgetData
			v10 := *(*int32)(unsafe.Add(v7, 48))
			if v10 == -1 {
				break
			}
			sv := &nox_xxx_saves_arr[v10]
			spath := alloc.GoStringS(sv.Path[:])
			if spath == "" {
				saveLog.Printf("save path is empty for slot %d", v10)
				clientPlaySoundSpecial(sound.SoundNoCanDo, 100)
				clientPlaySoundSpecial(sound.SoundShellClick, 100)
				return gui.RawEventResp(1)
			}
			v20 := datapath.SaveNameFromPath(spath)
			saveLog.Printf("loading slot %d: %q (%q, %q)", v10, v20, spath, alloc.GoStringS(sv.Map_name[:]))
			var v23 legacy.Nox_savegame_xxx
			if (!noxflags.HasGame(noxflags.GameModeCoop) || copySaveDir(v20, common.SaveTmp) == nil) && legacy.Sub_41A000(alloc.GoStringS(sv.Path[:]), &v23) != 0 {

				v23d := memmap.PtrT[legacy.Nox_savegame_xxx](0x85B3FC, 10980)
				*v23d = v23
				noxClient.GamePopState()
				switch getPlayerClass() {
				case player.Warrior:
					noxServer.nox_xxx_gameSetMapPath_409D70("war01a.map")
				case player.Wizard:
					noxServer.nox_xxx_gameSetMapPath_409D70("wiz01a.map")
				case player.Conjurer:
					noxServer.nox_xxx_gameSetMapPath_409D70("con01a.map")
				}
				if noxflags.HasGame(noxflags.GameModeCoop) {
					nox_xxx_gameSetSwitchSolo_4DB220(1)
					nox_xxx_gameSetNoMPFlag_4DB230(1)
					mname := alloc.GoStringS(sv.Map_name[:])
					fbase := fmt.Sprintf("%s.map", mname)
					v22, err := nox_client_checkSaveMapExistsTmp(fbase)
					if err != nil {
						v22 = datapath.Maps(mname, fbase)
					}
					nox_xxx_gameSetSoloSavePath_4DB270(v22)
					noxServer.nox_xxx_gameSetMapPath_409D70(fbase)
					legacy.Nox_xxx_mapLoadOrSaveMB_4DCC70(1)
					v13, _ := sub41D090(alloc.GoStringS(sv.Path[:]))
					noxServer.Objs.SetFirstObjectScriptID(server.ObjectScriptID(v13))
				} else if sub4D6F30() {
					sub_4DCE60(int(sv.Stage))
					noxServer.setQuestMapName(alloc.GoStringS(sv.Map_name[:]))
					v14, _ := sub41D090(alloc.GoStringS(sv.Path[:]))
					noxServer.Objs.SetFirstObjectScriptID(server.ObjectScriptID(v14))
				}
				sub4A24C0(false)
				sub_4A50A0()
				nox_wnd_xxx_1307748.Func13Ptr = nil
			}
		case 503:
			v7 := winCharListNames.WidgetData
			v5 := *(*int32)(unsafe.Add(v7, 48))
			if v5 == -1 {
				break
			}
			sv := &nox_xxx_saves_arr[v5]
			spath := alloc.GoStringS(sv.Path[:])
			var (
				v17     func()
				v16     gui.DialogFlags
				v6, v15 string
			)
			*memmap.PtrInt32(0x5D4594, 1307772) = v5
			npath := datapath.SaveNameFromPath(spath)
			nox_savegame_name_1307752 = npath
			if noxflags.HasGame(noxflags.GameModeCoop) && nox_savegame_name_1307752 == common.SaveAuto {
				clientPlaySoundSpecial(sound.SoundNoCanDo, 100)
				v17 = nil
				v16 = 33
				v15 = strMan.GetStringInFile("GUISave.c:AutoSaveDeleteNotAllowed", "C:\\NoxPost\\src\\client\\shell\\selchar.c")
				v6 = strMan.GetStringInFile("GUISave.c:AutoSaveDeleteTitle", "C:\\NoxPost\\src\\client\\shell\\selchar.c")
			} else {
				if spath == "" {
					clientPlaySoundSpecial(sound.SoundNoCanDo, 100)
					clientPlaySoundSpecial(sound.SoundShellClick, 100)
					return gui.RawEventResp(1)
				}
				v17 = sub_4A5C70
				v16 = 56
				if noxflags.HasGame(noxflags.GameModeCoop) {
					v15 = strMan.GetStringInFile("GUISave.c:DeleteSaveMessage", "C:\\NoxPost\\src\\client\\shell\\selchar.c")
					v6 = strMan.GetStringInFile("GUISave.c:DeleteSaveTitle", "C:\\NoxPost\\src\\client\\shell\\selchar.c")
				} else {
					v15 = strMan.GetStringInFile("GUISave.c:DeleteCharacterMessage", "C:\\NoxPost\\src\\client\\shell\\selchar.c")
					v6 = strMan.GetStringInFile("GUISave.c:DeleteCharacterTitle", "C:\\NoxPost\\src\\client\\shell\\selchar.c")
				}
			}
			NewDialogWindow(nil, v6, v15, v16, v17, nil)
		case 504, 505:
			ev2 := &WindowEvent0x4000{Win: e.Win}
			winCharList.Func94(ev2)
			winCharListNames.Func94(ev2)
			winCharListStyle.Func94(ev2)
		}
		clientPlaySoundSpecial(sound.SoundShellClick, 100)
		return gui.RawEventResp(1)
	}
	return nil
}

func sub41D090(path string) (uint32, error) {
	cf, err := cryptfile.OpenFile(path, cryptfile.ReadOnly, crypt.SaveKey)
	if err != nil {
		return 0, err
	}
	defer cf.Close()
	var buf [4]byte
	for {
		_, err := cf.ReadWrite(buf[:4])
		if err == io.EOF {
			return 0, nil
		} else if err != nil {
			return 0, err
		}
		a1 := binary.LittleEndian.Uint32(buf[:])
		if a1 == 0 {
			return 0, nil
		}
		cf.ReadMaybeAlign(buf[:4])
		v3 := binary.LittleEndian.Uint32(buf[:])
		if a1 == 10 {
			return sub_41D110(cf)
		}
		cf.Seek(int64(v3), io.SeekCurrent)
	}
}

func sub_41D110(cf *cryptfile.CryptFile) (uint32, error) {
	if !noxflags.HasGame(noxflags.GameModeCoop) {
		return 0, nil
	}
	var buf [4]byte
	v2 := uint16(5)
	binary.LittleEndian.PutUint16(buf[:], v2)
	_, err := cf.ReadWrite(buf[:2])
	v2 = binary.LittleEndian.Uint16(buf[:])
	if int16(v2) <= 5 && int16(v2) >= 5 {
		buf = [4]byte{}
		_, err = cf.ReadWrite(buf[:4])
		v3 := binary.LittleEndian.Uint32(buf[:])
		return v3, err
	}
	return uint32(v2), err
}

func nox_savegame_sub_46C730() int {
	win := newWindowFromFile(noxClient.GUI, "selchar.wnd", nox_savegame_sub_46C920)
	dword_5d4594_1082856 = win
	if win == nil {
		return 0
	}
	win.Flags |= 32
	dword_5d4594_1082856.SetFunc93(func(win *gui.Window, ev gui.WindowEvent) gui.WindowEventResp {
		return gui.RawEventResp(1)
	})
	dword_5d4594_1082860 = dword_5d4594_1082856.ChildByID(510)
	dword_5d4594_1082864 = dword_5d4594_1082856.ChildByID(511)
	dword_5d4594_1082868 = dword_5d4594_1082856.ChildByID(512)
	dword_5d4594_1082860.SetFunc94(nox_savegame_sub_46C920)
	dword_5d4594_1082864.SetParent(dword_5d4594_1082860)
	dword_5d4594_1082868.SetParent(dword_5d4594_1082860)
	swin1 := dword_5d4594_1082856.ChildByID(501)
	swin2 := dword_5d4594_1082856.ChildByID(502)
	dword_5d4594_1082880 = dword_5d4594_1082856.ChildByID(503)
	v1 := dword_5d4594_1082856.ChildByID(504)
	v1.Hide()
	v2 := dword_5d4594_1082856.ChildByID(505)
	v2.Hide()
	for i := uint(500); i <= 512; i++ {
		if v4 := dword_5d4594_1082856.ChildByID(i); v4 != nil {
			v4.Flags |= 1024
		}
	}
	v5 := strMan.GetStringInFile("SaveGUISave", "C:\\NoxPost\\src\\client\\Gui\\GUISave.c")
	sub_46AEE0(swin1, v5)
	v6 := strMan.GetStringInFile("SaveGUILoad", "C:\\NoxPost\\src\\client\\Gui\\GUISave.c")
	sub_46AEE0(swin2, v6)
	dword_5d4594_1082856.SetPos(image.Point{
		X: (nox_win_width - dword_5d4594_1082856.Size().X) / 2,
	})
	dword_5d4594_1082856.Hide()
	legacy.Nox_xxx_wnd_46ABB0(dword_5d4594_1082856, 0)
	return 1
}

func nox_savegame_sub_46C920(win1 *gui.Window, ev gui.WindowEvent) gui.WindowEventResp {
	c := noxClient
	switch ev := ev.(type) {
	case *WindowEvent0x4010:
		ev2 := gui.AsWindowEvent(0x4013, ev.Val, 0)
		dword_5d4594_1082860.Func94(ev2)
		dword_5d4594_1082864.Func94(ev2)
		dword_5d4594_1082868.Func94(ev2)
		return nil
	case *WindowEvent0x4007:
		clientPlaySoundSpecial(sound.SoundButtonPress, 100)
		switch ev.Win.ID() {
		case 501:
			saveNum := *(*int32)(unsafe.Add(dword_5d4594_1082864.WidgetData, 48))
			if saveNum < 0 {
				return nil
			}
			v9 := c.ClientPlayerUnit().Flags()
			if (v9 & 0x8000) != 0 {
				v10 := win1.ChildByID(501)
				legacy.Nox_xxx_wnd_46ABB0(v10, 0)
				return nil
			}
			if alloc.GoStringS(nox_savegame_arr_1064948[saveNum].Path[:]) != "" {
				path := datapath.SaveNameFromPath(alloc.GoStringS(nox_savegame_arr_1064948[saveNum].Path[:]))
				pathArr := memmap.PtrT[[16]byte](0x5D4594, 1082840)
				alloc.StrCopy(pathArr[:], path)
				dword_5d4594_1082856.Capture(false)
				msg := strMan.GetStringInFile("GUISave.c:OverwriteSaveMessage", "GUISave.c")
				title := strMan.GetStringInFile("GUISave.c:OverwriteSaveTitle", "GUISave.c")
				NewDialogWindow(dword_5d4594_1082856, title, msg, gui.DialogYesButton|gui.DialogNoButton|gui.DialogFlag6, func() {
					name := alloc.GoStringS(pathArr[:])
					SaveCoop(name)
					sub_46D6F0()
				}, func() {
					dword_5d4594_1082856.Capture(true)
				})
				return nil
			}
			name := common.SaveAuto
			if saveNum != 0 {
				name = fmt.Sprintf(common.SaveFormat, saveNum)
			}
			SaveCoop(name)
			sub_46D6F0()
			return nil
		case 502:
			v6 := *(*int32)(unsafe.Add(dword_5d4594_1082864.WidgetData, 48))
			if v6 >= 0 && alloc.GoStringS(nox_savegame_arr_1064948[v6].Path[:]) != "" {
				if legacy.Nox_xxx_playerAnimCheck_4372B0() != 0 {
					nox_savegame_sub_46CBD0()
					return nil
				}
				v12 := strMan.GetStringInFile("GUIQuit.c:ReallyLoadMessage", "C:\\NoxPost\\src\\client\\Gui\\GUISave.c")
				v7 := strMan.GetStringInFile("SelChar.c:LoadLabel", "C:\\NoxPost\\src\\client\\Gui\\GUISave.c")
				NewDialogWindow(nil, v7, v12, 24, nox_savegame_sub_46CBD0, sub_44A400)
				return nil
			}
			clientPlaySoundSpecial(sound.SoundNoCanDo, 100)
			return nil
		case 503:
			if sub_450560() {
				sub_450580()
				sub_43DDA0()
				nox_game_exit_xxx2()
				sub_446060()
				sub_46D6F0()
				return nil
			}
			sub_46D6F0()
			return nil
		}
		return nil
	}
	return nil
}

func nox_savegame_sub_46CBD0() {
	i := *(*int32)(unsafe.Add(dword_5d4594_1082864.WidgetData, 48))
	if alloc.GoStringS(nox_savegame_arr_1064948[i].Path[:]) == "" {
		return
	}
	v3 := datapath.SaveNameFromPath(alloc.GoStringS(nox_savegame_arr_1064948[i].Path[:]))
	if !clientLoadCoopGame(v3) {
		v1 := strMan.GetStringInFile("SaveErrorTitle", "GUISave.c")
		NewDialogWindow(nil, v1, v1, gui.DialogOKButton|gui.DialogFlag6, nil, nil)
	}
	if sub_450560() {
		sub_4505B0()
	}
	sub_44A400()
	sub_46D6F0()
}

func sub_4505B0() {
	sub_450580()
	legacy.Set_nox_gameDisableMapDraw_5d4594_2650672(0)
	v0 := int(nox_client_getIntroScreenDuration_44E3B0())
	noxClient.r.FadeInScreen(v0, true, sub_44E320)
	legacy.Set_nox_gameDisableMapDraw_5d4594_2650672(1)
}

func sub_43DDA0() {
	*memmap.PtrUint32(0x5D4594, 816344) = 0
	legacy.Sub_43D9E0(memmap.PtrOff(0x5D4594, 816060))
}

func sub_450580() {
	dialog.Sub_44D8F0()
	*memmap.PtrUint32(0x5D4594, 832488) = 1
	legacy.Set_dword_5d4594_831224(0)
	*memmap.PtrUint32(0x5D4594, 831292) = 0
	*memmap.PtrUint32(0x5D4594, 831296) = 0
}

func sub_4DB170(a1 bool, a2 unsafe.Pointer, a3 int) {
	dword_5d4594_1563092 = uint32(a3)
	dword_5d4594_1563088 = noxServer.Frame()
	dword_5d4594_1563084 = a2
	dword_5d4594_1563080 = a1
	legacy.Set_dword_5d4594_1563096(bool2int(a2 != nil))
	if !a1 {
		sub_4DCBD0(0)
	}
}

func sub_4DCBD0(a1 int) {
	*memmap.PtrUint32(0x5D4594, 1563076) = uint32(a1)
}

func sub_46CCB0() {
	dword_5d4594_1082856.Destroy()
	dword_5d4594_1082856 = nil
	dword_5d4594_1082864 = nil
	dword_5d4594_1082868 = nil
}
