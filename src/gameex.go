package opennox

import (
	"bufio"
	"encoding/binary"
	"image"
	"net"
	"net/netip"
	"os"
	"strconv"
	"strings"
	"unsafe"

	"github.com/noxworld-dev/opennox-lib/client/keybind"
	"github.com/noxworld-dev/opennox-lib/datapath"
	"github.com/noxworld-dev/opennox-lib/ifs"
	"github.com/noxworld-dev/opennox-lib/log"
	"github.com/noxworld-dev/opennox-lib/player"

	"github.com/noxworld-dev/opennox/v1/client/gui"
	"github.com/noxworld-dev/opennox/v1/common/config"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/sound"
	"github.com/noxworld-dev/opennox/v1/internal/netstr"
	"github.com/noxworld-dev/opennox/v1/legacy"
	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc"
)

var modifyWndPntr *gui.Window
var gameex = struct {
	Log        *log.Logger
	configPath string
}{
	Log: log.New("gameex"),
}

func gameexDefaults() {
	config.Global.SetDefault(configKeyGameExAutoShield, true)
	config.Global.SetDefault(configKeyGameExGreatSwordBlock, false)
	config.Global.SetDefault(configKeyGameExWarWeaponScroll, true)
	config.Global.SetDefault(configKeyGameExBerserkShieldBlock, false)
}

const (
	configKeyGameExAutoShield         = "extensions.warrior.auto_shield"
	configKeyGameExGreatSwordBlock    = "extensions.warrior.greatsword_block"
	configKeyGameExWarWeaponScroll    = "extensions.warrior.weapon_scroll"
	configKeyGameExBerserkShieldBlock = "extensions.warrior.berserk_shield_block"
)

func gameexLoadConfig() {
	for _, f := range []struct {
		Key  string
		Flag int
	}{
		{Key: configKeyGameExAutoShield, Flag: 0x2},
		{Key: configKeyGameExGreatSwordBlock, Flag: 0x4},
		{Key: configKeyGameExWarWeaponScroll, Flag: 0x8},
		{Key: configKeyGameExBerserkShieldBlock, Flag: 0x10},
	} {
		if config.Global.GetBool(f.Key) {
			*legacy.Get_gameex_flags_ptr() |= uint32(f.Flag)
		} else {
			*legacy.Get_gameex_flags_ptr() &^= uint32(f.Flag)
		}
	}
}

func gameexReadConfig(path string) error {
	defer func() {
		gameexDefaults()
		gameexLoadConfig()
	}()
	path = datapath.Data(path)
	gameex.configPath = path
	f, err := ifs.Open(path)
	if os.IsNotExist(err) {
		gameex.Log.Println("no config file")
		return nil
	} else if err != nil {
		gameex.Log.Println(err)
		return err
	}
	defer f.Close()
	gameex.Log.Println("reading config", path)

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if len(line) == 0 || line[0] == '-' || line[0] == '#' {
			continue
		}
		sub := strings.SplitN(line, "=", 2)
		if len(sub) < 2 {
			gameex.Log.Printf("malformed option: %q", line)
			continue
		}
		key := strings.TrimSpace(sub[0])
		val := strings.TrimSpace(sub[1])
		switch key {
		case "AUTO_SHIELD":
			if !config.Global.IsSet(configKeyGameExAutoShield) {
				v, err := strconv.ParseBool(val)
				if err != nil {
					return err
				}
				config.Global.Set(configKeyGameExAutoShield, v)
			}
		case "GREAT_SWORD_BLOKING_WALK", "GREAT_SWORD_BLOCKING_WALK":
			if !config.Global.IsSet(configKeyGameExGreatSwordBlock) {
				v, err := strconv.ParseBool(val)
				if err != nil {
					return err
				}
				config.Global.Set(configKeyGameExGreatSwordBlock, v)
			}
		case "MOUSE_KEYBOARD_ROLL":
			if !config.Global.IsSet(configKeyGameExWarWeaponScroll) {
				v, err := strconv.ParseBool(val)
				if err != nil {
					return err
				}
				config.Global.Set(configKeyGameExWarWeaponScroll, v)
			}
		case "BERSERKER_SHIED_BLOCK", "BERSERKER_SHIELD_BLOCK":
			if !config.Global.IsSet(configKeyGameExBerserkShieldBlock) {
				v, err := strconv.ParseBool(val)
				if err != nil {
					return err
				}
				config.Global.Set(configKeyGameExBerserkShieldBlock, v)
			}
		case "EXTENSION_MESSAGES":
			v, err := strconv.ParseBool(val)
			if err != nil {
				return err
			}
			if v {
				*legacy.Get_gameex_flags_ptr() |= 0x20
			} else {
				*legacy.Get_gameex_flags_ptr() &^= 0x20
			}
		case "PANEL1", "PANEL2", "PANEL3", "PANEL4", "PANEL5", "TRAPKEY":
		default:
			gameex.Log.Printf("unrecognized option: %q", key)
		}
	}
	return sc.Err()
}

func gameex_makeExtensionPacket(buf []byte, opcode uint16, needsPlayer bool) {
	binary.LittleEndian.PutUint16(buf[0:], 0xF13A) // extension packet code
	binary.LittleEndian.PutUint16(buf[2:], opcode)
	if needsPlayer {
		binary.LittleEndian.PutUint32(buf[4:], uint32(legacy.ClientPlayerNetCode())) // playerNetCode
	}
}

func MixRecvFromReplacer(pc net.PacketConn, buf1 []byte, from netip.AddrPort) {
	op := binary.LittleEndian.Uint16(buf1[2:])
	buf1 = buf1[4:]
	switch op {
	case 0: // warrior weapon scroll
		if (legacy.Get_gameex_flags()>>3)&1 != 0 {
			ti := binary.LittleEndian.Uint32(buf1)
			buf1 = buf1[4:]
			v10 := noxServer.getObjectFromNetCode(int(ti))
			v9 := buf1[0]
			buf1 = buf1[1:]
			v36 := (uint32(v9)>>4)&1 != 0
			if v36 {
				v9 &= 0xEF
			}
			if v10.ControllingPlayer().PlayerClass() == player.Warrior || v36 {
				if legacy.Mix_MouseKeyboardWeaponRoll(v10, int8(v9)) != 0 {
					var buf [4]byte
					binary.LittleEndian.PutUint16(buf[0:], 0xF13A)
					binary.LittleEndian.PutUint16(buf[2:], 2)
					netstr.WriteTo(pc, buf[:4], from)
				}
			}
		}
	case 2: // clientPlaySoundSpecial
		clientPlaySoundSpecial(sound.SoundNextWeapon, 100)
	case 4:
		if noxflags.HasGame(noxflags.GameHost) {
			if (legacy.Get_gameex_flags()>>5)&1 != 0 {
				data := buf1[:alloc.StrLenS(buf1)]
				v43, free := alloc.Make([]byte{}, 2+len(data)+1)
				defer free()
				copy(v43[2:], data)
				if legacy.PlayerInfoStructParser_0(unsafe.Pointer(&v43[0])) != 0 {
					var buf [22]byte
					binary.LittleEndian.PutUint16(buf[0:], 0xF13A)
					binary.LittleEndian.PutUint16(buf[2:], 6)
					alloc.StrCopy(buf[4:22], noxServer.getServerName())
					netstr.WriteTo(pc, buf[:22], from)
				}
			}
		}
	case 5:
		if (legacy.Get_gameex_flags()>>5)&1 != 0 {
			nox_xxx_printCentered_445490(alloc.GoStringS(buf1))
			clientPlaySoundSpecial(sound.SoundManaClick, 100)
		}
	case 7:
		nox_xxx_printCentered_445490(alloc.GoStringS(buf1))
		clientPlaySoundSpecial(sound.SoundManaClick, 100)
	case 8:
		if noxflags.HasGame(noxflags.GameHost) && ((legacy.Get_gameex_flags()>>5)&1 != 0) {
			data := buf1[:alloc.StrLenS(buf1)]
			v43, free := alloc.Make([]byte{}, 2+len(data)+1)
			defer free()
			v39, free39 := alloc.New(int32(0))
			defer free39()
			copy(v43[2:], data)
			if legacy.PlayerInfoStructParser_1(unsafe.Pointer(&v43[0]), v39) != 0 {
				// TODO: decipher it

				//buf[2] = 7
				//v29 = len(v43[2:]) - 1
				//v30 = &buf[v29+4]
				//v31 = strlen(v30)
				//if v31 {
				//	v32 = (*byte)(v44 - (*byte)(v30))
				//	for {
				//		v33 = *v30
				//		v30[uint32(v32)] = *v30
				//		v30++
				//		if !(v33) {
				//			break
				//		}
				//	}
				//	memset(v26, 0, v29+v31)
				//	strcpy(v26, v44)
				//}
			}
		}
	case 9:
		if (legacy.Get_gameex_flags()>>3)&1 != 0 {
			ti := binary.LittleEndian.Uint32(buf1)
			v35 := noxServer.getObjectFromNetCode(int(ti))
			legacy.PlayerDropATrap(v35)
		}
	}
}

var wndEntryNames = [5][35]uint16{
	{119, 97, 114, 114, 105, 111, 114, 115, 32, 108, 105, 107, 101, 32, 115, 104, 105, 101, 108, 100, 115, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{119, 97, 108, 107, 105, 110, 103, 32, 119, 105, 116, 104, 32, 115, 119, 111, 114, 100, 32, 98, 108, 111, 99, 107, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{119, 101, 97, 112, 111, 110, 32, 114, 111, 108, 108, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{115, 104, 105, 101, 108, 100, 32, 38, 32, 98, 101, 114, 115, 101, 114, 107, 101, 114, 32, 98, 108, 111, 99, 107, 105, 110, 103, 0, 0, 0, 0, 0, 0, 0, 0},
	{101, 120, 116, 101, 110, 115, 105, 111, 110, 32, 109, 101, 115, 115, 97, 103, 101, 115, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
}

func gameexDropTrap() {
	c := noxClient
	if noxflags.HasGame(noxflags.GameFlag3 | noxflags.GameModeCoopTeam) {
		if legacy.Get_dword_5d4594_1064868() != 0 || c.GUI.Captured() != nil {
			return
		}
		if noxflags.HasGame(noxflags.GameHost) { // checkGameFlags isServer
			v9 := noxServer.getObjectFromNetCode(legacy.ClientPlayerNetCode())
			legacy.PlayerDropATrap(v9)
		} else {
			// TODO: this currently relies on extension packets, which should not be required for this
			//       it can be done the "natural way": find the trap in the client-side data structures
			//       and ask server to drop that item, as the client does when doing the same manually
			buf, freeBuf := alloc.Make([]byte{}, 10)
			defer freeBuf()
			gameex_makeExtensionPacket(buf, 9, true)
			c.Conn.SendSelfRaw(buf[:8])
		}
	}
}

func call_OnLibraryNotice_265(arg3 int) {
	c := noxClient
	// toggles weapons by mouse wheel
	// autoshield is actually implemented in appendix of nox_xxx_playerDequipWeapon_53A140
	a2a := bool2int(arg3 > 0) // scroll weapons back or forth
	if !c.GUI.GameexCheck() {
		return
	}
	if (legacy.Get_gameex_flags()>>3)&1 == 0 {
		return
	}
	if !noxflags.HasGame(noxflags.GameFlag3 | noxflags.GameModeCoopTeam) {
		return
	}
	if noxflags.HasGame(noxflags.GameHost) {
		if pl := c.srv.Players.Host(); pl != nil && pl.PlayerClass() == player.Warrior && pl.PlayerUnit != nil {
			if legacy.Mix_MouseKeyboardWeaponRoll(pl.PlayerUnit, int8(a2a)) != 0 {
				clientPlaySoundSpecial(sound.SoundNextWeapon, 100)
			}
		}
	} else {
		var buf [10]byte
		gameex_makeExtensionPacket(buf[:], 0, true)
		buf[8] = byte(a2a)
		c.Conn.SendSelfRaw(buf[:9])
	}
}

func gameexOnKeyboardPress(kcode keybind.Key) {
	c := noxClient
	if ((legacy.Get_gameex_flags()>>3)&1 != 0) && (kcode == keybind.KeyLBracket || kcode == keybind.KeyRBracket) {
		v8 := byte(bool2int(kcode == keybind.KeyLBracket))
		// checks some gameFlags that are yet undiscovered
		if noxflags.HasGame(noxflags.GameFlag3 | noxflags.GameModeCoopTeam) {
			if legacy.Get_dword_5d4594_1064868() != 0 || noxClient.GUI.Captured() != nil {
				return
			}
			if noxflags.HasGame(noxflags.GameHost) { // isServer
				if u := noxServer.Players.HostUnit(); u != nil && legacy.Mix_MouseKeyboardWeaponRoll(u, int8(v8)) != 0 {
					clientPlaySoundSpecial(sound.SoundNextWeapon, 100)
				}
			} else {
				buf, freeBuf := alloc.Make([]byte{}, 10)
				defer freeBuf()
				gameex_makeExtensionPacket(buf, 0, true)
				buf[8] = v8 | 0x10 // TODO: should it be just v8?
				c.Conn.SendSelfRaw(buf[:9])
			}
		}
	}
	if kcode == keybind.KeyF8 { // TODO: should be configurable, or actually just set defaults and remove this window
		if !noxflags.HasGame(noxflags.GameHost) {
			nox_xxx_printCentered_445490("only server can change these options")
			clientPlaySoundSpecial(sound.SoundPermanentFizzle, 100)
			return
		}
		if !noxflags.HasGame(noxflags.GameFlag3 | noxflags.GameModeCoopTeam) {
			return
		}
		clientPlaySoundSpecial(sound.SoundShellClick, 100)
		if modifyWndPntr != nil {
			destroyGameExWindow()
			modifyWndPntr = nil
		} else {
			modifyWndPntr = newWindowFromString(noxClient.GUI, gameexModifWnd, modifyWndInputHandler)
			if modifyWndPntr == nil {
				return
			}
			if noxflags.HasGame(noxflags.GameModeCoopTeam) {
				modifyWndPntr.ChildByID(1938).Hide()
				v11 := modifyWndPntr.ChildByID(1524)
				legacy.Nox_xxx_wnd_46ABB0(v11, 0)
			}
			modifyWndPntr.SetParent(nil)
			a2b := modifyWndPntr.ChildByID(1981)
			for i := 0; i < 5; i++ {
				wstr := alloc.GoString16S(wndEntryNames[i][:])
				id := uint(1520 + i)
				a2b.Func94(&WindowEvent0x400d{Str: wstr, Val: -1})
				if legacy.GetFlagValueFromFlagIndex(int(id)-1519)&legacy.Get_gameex_flags() != 0 {
					v14 := modifyWndPntr.ChildByID(id)
					v14.DrawData().Field0 |= 0x4
				} else {
					v15 := modifyWndPntr.ChildByID(id)
					v15.DrawData().Field0 &= 0xFFFFFFFB
				}
			}
		}
	}
}

func destroyGameExWindow() {
	modifyWndPntr.StackPop()
	modifyWndPntr.Destroy()
	modifyWndPntr = nil
}

func modifyWndInputHandler(a1 *gui.Window, ev gui.WindowEvent) gui.WindowEventResp {
	switch ev := ev.(type) {
	case *WindowEvent0x4007:
		clientPlaySoundSpecial(sound.SoundButtonPress, 100)
		switch ev.Win.ID() {
		case 1937:
			destroyGameExWindow()
		case 1938:
			if !noxflags.HasGame(noxflags.GameModeCoopTeam) {
				legacy.Sub_4BDFD0()
				legacy.Get_dword_5d4594_1316972().SetPos(image.Point{X: 200, Y: 100})
			}
		case 1520:
			if (legacy.Get_gameex_flags()>>1)&1 != 0 {
				*legacy.Get_gameex_flags_ptr() &= 0xFFFFFFFD
			} else {
				*legacy.Get_gameex_flags_ptr() |= 0x2
			}
		case 1521:
			if (legacy.Get_gameex_flags()>>2)&1 != 0 {
				*legacy.Get_gameex_flags_ptr() &= 0xFFFFFFFB
			} else {
				*legacy.Get_gameex_flags_ptr() |= 0x4
			}
		case 1522:
			if (legacy.Get_gameex_flags()>>3)&1 != 0 {
				*legacy.Get_gameex_flags_ptr() &= 0xFFFFFFF7
			} else {
				*legacy.Get_gameex_flags_ptr() |= 0x8
			}
		case 1523:
			if (legacy.Get_gameex_flags()>>4)&1 != 0 {
				*legacy.Get_gameex_flags_ptr() &= 0xFFFFFFEF
			} else {
				*legacy.Get_gameex_flags_ptr() |= 0x10
			}
		case 1524:
			if (legacy.Get_gameex_flags()>>5)&1 != 0 {
				*legacy.Get_gameex_flags_ptr() &= 0xFFFFFFDF
			} else {
				*legacy.Get_gameex_flags_ptr() |= 0x20
			}
		}
		return nil
	}
	return nil
}

const gameexModifWnd = `
ENABLEDCOLOR = TRANSPARENT;
HILITECOLOR = 255 192 192;
DISABLEDCOLOR = 128 128 128;
SELECTEDCOLOR = 255 128 128;
TEXTCOLOR = 255 255 255;
BACKGROUNDCOLOR = 0 0 0;

WINDOW
  1930 50 55 200 150 USER;
  STATUS = ENABLED+ABOVE+BORDER;

  CHILD
	WINDOW
	  1520 8 22 15 10 CHECKBOX (Object allowed);
      STATUS = ENABLED+IMAGE;
	  STYLE = MOUSETRACK;
	  SELECTEDIMAGE = UICheckBoxLit;
	  HILITEIMAGE = NULL;
	  BACKGROUNDIMAGE = UICheckBox;
	  ENABLEDIMAGE = NULL;
	  DISABLEDIMAGE = UICheckBoxDis;
	  TOOLTIP = objlst.wnd:ToggleTT;
	END

	WINDOW
	  1521 8 32 15 10 CHECKBOX (Object allowed);
      STATUS = ENABLED+IMAGE;
	  SELECTEDIMAGE = UICheckBoxLit;
	  HILITEIMAGE = NULL;
	  BACKGROUNDIMAGE = UICheckBox;
	  ENABLEDIMAGE = NULL;
	  DISABLEDIMAGE = UICheckBoxDis;
	  STYLE = MOUSETRACK;
	  TOOLTIP = objlst.wnd:ToggleTT;
	END

	WINDOW
	  1522 8 42 15 10 CHECKBOX (Object allowed);
      STATUS = ENABLED+IMAGE;
	  SELECTEDIMAGE = UICheckBoxLit;
	  HILITEIMAGE = NULL;
	  BACKGROUNDIMAGE = UICheckBox;
	  ENABLEDIMAGE = NULL;
	  DISABLEDIMAGE = UICheckBoxDis;
	  STYLE = MOUSETRACK;
	  TOOLTIP = objlst.wnd:ToggleTT;
	END

	WINDOW
	  1523 8 52 15 10 CHECKBOX (Object allowed);
      STATUS = ENABLED+IMAGE;
	  SELECTEDIMAGE = UICheckBoxLit;
	  HILITEIMAGE = NULL;
	  BACKGROUNDIMAGE = UICheckBox;
	  ENABLEDIMAGE = NULL;
	  DISABLEDIMAGE = UICheckBoxDis;
	  STYLE = MOUSETRACK;
	  TOOLTIP = objlst.wnd:ToggleTT;
	END

	WINDOW
	  1524 8 62 15 10 CHECKBOX (Object allowed);
      STATUS = ENABLED+IMAGE;
	  SELECTEDIMAGE = UICheckBoxLit;
	  HILITEIMAGE = NULL;
	  BACKGROUNDIMAGE = UICheckBox;
	  ENABLEDIMAGE = NULL;
	  DISABLEDIMAGE = UICheckBoxDis;
	  STYLE = MOUSETRACK;
	  TOOLTIP = objlst.wnd:ToggleTT;
	END

        WINDOW
          1937 65 120 70 20 PUSHBUTTON;
          STATUS = ENABLED+IMAGE;
          STYLE = MOUSETRACK+TABSTOP;
	  SELECTEDIMAGE = UIButtonLgLit;
	  HILITEIMAGE = UIButtonLgLit;
	  BACKGROUNDIMAGE = UIButtonLg;
	  ENABLEDIMAGE = NULL;
	  DISABLEDIMAGE = UIButtonLgDis;
          TEXT = WolFind.wnd:Done;
       END

        WINDOW
          1938 115 80 70 20 PUSHBUTTON;
          STATUS = ENABLED+IMAGE;
          STYLE = MOUSETRACK+TABSTOP;
	  SELECTEDIMAGE = UIButtonLgLit;
	  HILITEIMAGE = UIButtonLgLit;
	  BACKGROUNDIMAGE = UIButtonLg;
	  ENABLEDIMAGE = NULL;
	  DISABLEDIMAGE = UIButtonLgDis;
          TEXT = servopts.wnd:Advanced;
       END

        WINDOW
          1981 22 21 150 60 SCROLLLISTBOX
          STATUS = ONE_LINE;
          DISABLEDCOLOR = TRANSPARENT;
          DATA = 26 10 0 1 0 0 1;
          FONT = default;
       END
END
`
