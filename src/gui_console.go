package opennox

import (
	"context"
	"fmt"
	"image"
	"image/color"
	"log/slog"
	"strings"

	"github.com/opennox/libs/client/keybind"
	"github.com/opennox/libs/console"
	"github.com/opennox/libs/log"

	"github.com/opennox/opennox/v1/client/gui"
	noxflags "github.com/opennox/opennox/v1/common/flags"
	"github.com/opennox/opennox/v1/common/memmap"
	"github.com/opennox/opennox/v1/internal/version"
	"github.com/opennox/opennox/v1/legacy"
	"github.com/opennox/opennox/v1/legacy/common/alloc"
)

var (
	guiCon = &guiConsole{c: noxConsole, translucent: true}
)

func init() {
	consoleMux.Add(guiCon)
	log.AddHandler(guiCon)
	guiCon.c.Register(&console.Command{
		Token: "clear", HelpID: "clearhelp",
		Flags: console.ClientServer,
		Func:  guiCon.cmdClear,
	})
	guiCon.c.Register(&console.Command{
		Token: "lock", HelpID: "lockhelp",
		Flags: console.ClientServer,
		Func:  guiCon.cmdLock,
	})
	guiCon.c.Register(&console.Command{
		Token: "unlock", HelpID: "unlockhelp",
		Flags: console.ClientServer,
		Func:  guiCon.cmdUnlock,
	})
	legacy.GetConsole = func() *console.Console {
		return noxConsole
	}
}

func nox_gui_console_flagXxx_451410() int {
	return int((^(guiCon.root.GetFlags() >> 4)) & 1)
}

func nox_gui_console_Hide_4512B0() int {
	return bool2int(guiCon.Hide())
}

var _ log.Handler = (*guiConsole)(nil)

type guiConsole struct {
	c           *console.Console
	password    string
	root        *gui.Window
	scrollbox   *gui.Window
	input       *gui.Window
	translucent bool
	wantsPass   bool
	enabled     bool
}

func (c *guiConsole) Enable(v bool) {
	c.enabled = v
}

func (c *guiConsole) Enabled(ctx context.Context, level slog.Level) bool {
	const cur = slog.LevelWarn
	return c.enabled && level >= cur
}

func (c *guiConsole) Handle(ctx context.Context, r slog.Record) error {
	if !c.Enabled(ctx, r.Level) {
		return nil
	}
	const logFlags = log.PrintSys
	var buf strings.Builder
	if err := log.PrintRecord(&buf, logFlags, r); err != nil {
		return err
	}
	cl := console.ColorLightYellow
	switch r.Level {
	case slog.LevelDebug:
		cl = console.ColorDarkGrey
	case slog.LevelInfo:
		cl = console.ColorLightGrey
	case slog.LevelWarn:
		cl = console.ColorYellow
	case slog.LevelError:
		cl = console.ColorRed
	}
	c.Print(cl, buf.String())
	return nil
}

func (c *guiConsole) Print(cl console.Color, str string) {
	if c.enabled && c.scrollbox != nil {
		c.scrollbox.Func94(&WindowEvent0x400d{Str: str, Val: int(cl)})
	}
}

func (c *guiConsole) Printf(cl console.Color, format string, args ...interface{}) {
	str := fmt.Sprintf(format, args...)
	c.Print(cl, str)
}

func (c *guiConsole) PrintOrError(cl console.Color, str string) {
	if str != "" {
		c.Print(cl, str)
		return
	}
	str = c.c.Strings().GetStringInFile("InternalError", "guicon.c")
	if str != "" {
		c.Print(cl, str)
	}
}

func (c *guiConsole) Clear() {
	c.scrollbox.Func94(gui.AsWindowEvent(0x400F, 0, 0))
}

func (c *guiConsole) Init(sz image.Point) *gui.Window {
	*memmap.PtrInt32(0x5D4594, 833704) = int32(sz.X) - 1
	*memmap.PtrInt32(0x5D4594, 833708) = int32(sz.Y) / 2
	c.root = noxClient.GUI.NewWindowRaw(nil, 0x38, 0, 0, sz.X-1, sz.Y/2, nil)
	c.root.SetAllFuncs(func(win *gui.Window, ev gui.WindowEvent) gui.WindowEventResp {
		return nil
	}, func(win *gui.Window, a2 *gui.WindowData) int {
		r := noxClient.r
		pos := win.GlobalPos()
		if win.GetFlags().Has(gui.StatusImage) {
			r.DrawImageAt(a2.BackgroundImage(), pos)
			return 1
		}
		wsz := win.Size()
		if _, _, _, alpha := a2.BackgroundColor().RGBA(); alpha != 0 {
			if c.translucent {
				r.DrawRectFilledAlpha(pos.X, pos.Y, wsz.X, wsz.Y)
				return 1
			}
			r.DrawRectFilledOpaque(pos.X, pos.Y, wsz.X, wsz.Y, a2.BackgroundColor())
		}
		return 1
	}, nil)
	c.root.DrawData().SetBackgroundColor(color.Black)

	drawData, freeDraw := alloc.New(gui.WindowData{})
	defer freeDraw()

	drawData.SetHighlightColor(color.Transparent)
	drawData.SetSelectedColor(color.Transparent)
	drawData.SetBackgroundImage(nil)
	drawData.SetDisabledImage(nil)
	drawData.SetEnabledColor(nox_color_gray2)
	drawData.SetTextColor(nox_color_gray2)
	drawData.SetText(version.ClientVersion())

	drawData.Style = gui.StyleScrollListBox

	scrollData, freeScr := alloc.New(gui.ScrollListBoxData{})
	defer freeScr()
	scrollData.Count = 128
	scrollData.Line_height = 10
	scrollData.Field_1 = 1
	scrollData.Field_2 = 1
	scrollData.Field_3 = 1
	scrollData.Field_4 = 0
	c.scrollbox = legacy.Nox_gui_newScrollListBox_4A4310(c.root, 1152, 0, 0, int(memmap.Int32(0x5D4594, 833704)), int(memmap.Int32(0x5D4594, 833708))-20, drawData, scrollData)
	drawData.SetDisabledColor(nox_color_gray2)
	drawData.SetEnabledColor(nox_color_gray2)
	drawData.SetHighlightColor(nox_color_gray2)
	drawData.SetSelectedColor(nox_color_gray2)
	drawData.SetTextColor(nox_color_gray2)

	inpData, freeInp := alloc.New(gui.EntryFieldData{})
	defer freeInp()
	inpData.Text[0] = 0
	inpData.Field_1024 = 0
	inpData.Field_1028 = 0
	inpData.Field_1032 = 0
	inpData.Field_1036 = 0
	inpData.Field_1040 = 128
	inpData.Field_1042 = memmap.Int16(0x5D4594, 833704)

	drawData.SetText("")
	drawData.Style = gui.StyleEntryField
	c.input = legacy.Nox_gui_newEntryField_488500(c.root, 16393, 0, int(memmap.Int32(0x5D4594, 833708))-20, int(memmap.Int32(0x5D4594, 833704)), 20, drawData, inpData)
	c.input.SetFunc93(c.inputProc)
	c.wantsPass = false
	c.password = ""
	return c.root
}

func (c *guiConsole) Hide() bool {
	if c.root.GetFlags().IsHidden() {
		return false
	}
	if noxClient.GUI.Focused() == c.input {
		noxClient.GUI.Focus(nil)
	}
	c.root.Hide()
	c.root.Flags &= 0xFFFFFFF7
	c.input.Flags &= 0xFFFFFFF7
	c.scrollbox.Flags &= 0xFFFFFFF7
	c.input.DrawData().Field0 &= 0xFFFFFFFB
	c.input.DrawData().Field0 &= 0xFFFFFFFD
	noxClient.GUI.ValYYY = 1
	return true
}

func (c *guiConsole) ReloadColors() {
	gray := nox_color_gray2
	black := nox_color_black_2650656
	white := nox_color_white_2523948
	if c.root != nil {
		c.root.DrawData().SetBackgroundColor(black)
	}
	if c.scrollbox != nil {
		dd := c.scrollbox.DrawData()
		dd.SetBackgroundColor(color.Transparent)
		dd.SetDisabledColor(color.Transparent)
		dd.SetEnabledColor(gray)
		dd.SetHighlightColor(gray)
		dd.SetSelectedColor(color.Transparent)
		dd.SetTextColor(gray)

		scr := (*gui.ScrollListBoxData)(c.scrollbox.WidgetData)
		if v2 := legacy.AsWindowP(scr.Field_9); v2 != nil {
			dd2 := v2.DrawData()
			dd2.SetBackgroundColor(black)
			dd2.SetDisabledColor(black)
			dd2.SetEnabledColor(gray)
			dd2.SetHighlightColor(black)
			dd2.SetSelectedColor(gray)
			if v3 := v2.Field100(); v3 != nil {
				dd3 := v3.DrawData()
				dd3.SetBackgroundColor(black)
				dd3.SetDisabledColor(black)
				dd3.SetEnabledColor(gray)
				dd3.SetHighlightColor(black)
				dd3.SetSelectedColor(black)
			}
		}
		if v4 := legacy.AsWindowP(scr.Field_7); v4 != nil {
			dd4 := v4.DrawData()
			dd4.SetBackgroundColor(black)
			dd4.SetDisabledColor(black)
			dd4.SetEnabledColor(gray)
			dd4.SetHighlightColor(white)
			dd4.SetSelectedColor(gray)
			dd4.SetTextColor(gray)
		}
		if v5 := legacy.AsWindowP(scr.Field_8); v5 != nil {
			dd5 := v5.DrawData()
			dd5.SetBackgroundColor(black)
			dd5.SetDisabledColor(black)
			dd5.SetEnabledColor(gray)
			dd5.SetHighlightColor(white)
			dd5.SetSelectedColor(gray)
			dd5.SetTextColor(gray)
		}
	}
	if c.input != nil {
		dd := c.input.DrawData()
		dd.SetBackgroundColor(black)
		dd.SetDisabledColor(gray)
		dd.SetEnabledColor(gray)
		dd.SetHighlightColor(gray)
		dd.SetSelectedColor(gray)
		dd.SetTextColor(gray)
	}
}

func (c *guiConsole) SetPass(pass string) {
	if pass != "" {
		c.password = pass
	}
}

func (c *guiConsole) ResetPass() {
	c.password = ""
	c.wantsPass = false
}

func (c *guiConsole) Toggle() {
	if !c.root.GetFlags().Has(gui.StatusHidden) {
		c.Hide()
		return
	}
	if legacy.Nox_gui_xxx_check_446360() == 0 {
		c.root.ShowModal()
		c.root.Flags |= 8
		c.scrollbox.Flags |= 8
		c.input.Flags |= 8
		c.input.Flags |= 1
		noxClient.GUI.Focus(c.input)
		if c.password != "" {
			c.Clear()
			s := c.c.Strings().GetStringInFile("ENTERPASSWORD", "C:\\NoxPost\\src\\Client\\Gui\\guicon.c")
			c.c.Print(console.ColorWhite, s)
			c.wantsPass = true
		}
	}
}

func (c *guiConsole) inputProc(win *gui.Window, ev gui.WindowEvent) gui.WindowEventResp {
	cl := noxClient
	switch ev := ev.(type) {
	case gui.WindowKeyPress:
		if cl.ctrl.hasDefBinding(keybind.EventToggleConsole, ev.Key) {
			if ev.Pressed {
				c.Toggle()
			}
			return gui.RawEventResp(1)
		}
		switch ev.Key {
		case keybind.KeyEsc:
			if ev.Pressed {
				legacy.Nox_xxx_consoleEsc_49B7A0()
			}
			return gui.RawEventResp(1)
		case keybind.KeyEnter:
			if ev.Pressed {
				c.nox_gui_console_Enter_450FD0()
			}
			return gui.RawEventResp(1)
		}
	}
	return legacy.Nox_xxx_wndEditProc_487D70(win, ev)
}

func (c *guiConsole) nox_gui_console_Enter_450FD0() {
	wd := (*gui.EntryFieldData)(c.input.WidgetData)
	if wd.Field_1044 != 0 {
		return
	}
	cmd := eventRespStr(c.input.Func94(gui.AsWindowEvent(0x401d, 0, 0)))
	if c.wantsPass && c.password != "" {
		if cmd != c.password {
			s := strMan.GetStringInFile("INVALIDPASSWORD", "C:\\NoxPost\\src\\Client\\Gui\\guicon.c")
			c.PrintOrError(console.ColorWhite, s)
		} else {
			c.wantsPass = false
			s := strMan.GetStringInFile("ConsoleUnlocked", "C:\\NoxPost\\src\\Client\\Gui\\guicon.c")
			c.PrintOrError(console.ColorRed, s)
		}
		c.input.Func94(gui.AsWindowEvent(0x401e, uintptr(memmap.PtrOff(0x5D4594, 833744)), 0))
		return
	}
	c.Print(console.ColorWhite, "> "+cmd)
	if cmd != "" {
		if noxflags.HasEngine(noxflags.EngineReplayWrite) {
			nox_xxx_replaySaveConsole(cmd)
		}
		execConsoleCmd(context.Background(), cmd)
	}
	c.input.Func94(gui.AsWindowEvent(0x401e, uintptr(memmap.PtrOff(0x5D4594, 833748)), 0))
}

func (c *guiConsole) cmdClear(ctx context.Context, _ *console.Console, tokens []string) bool {
	if len(tokens) != 0 {
		return false
	}
	c.Clear()
	return true
}

func (c *guiConsole) cmdLock(ctx context.Context, _ *console.Console, tokens []string) bool {
	if len(tokens) != 1 {
		return false
	}
	c.SetPass(tokens[0])
	c.Clear()
	s := c.c.Strings().GetStringInFile("consolelocked", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
	c.Print(console.ColorWhite, s)
	return true
}

func (c *guiConsole) cmdUnlock(ctx context.Context, _ *console.Console, tokens []string) bool {
	if len(tokens) != 0 {
		return false
	}
	c.ResetPass()
	s := c.c.Strings().GetStringInFile("consoleunlocked", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
	c.Print(console.ColorWhite, s)
	return true
}
