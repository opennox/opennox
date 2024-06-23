//go:build !server

package opennox

import (
	"errors"
	"image"
	"image/color"

	"github.com/noxworld-dev/opennox-lib/ifs"

	"github.com/noxworld-dev/opennox/v1/client/gui"
	"github.com/noxworld-dev/opennox/v1/client/noxmovie"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"github.com/noxworld-dev/opennox/v1/legacy"
)

func (c *Client) DrawGeneral(a1 bool) error {
	dword_5d4594_1311936 = true
	*memmap.PtrUint32(0x5D4594, 1311932) = uint32(bool2int(a1))
	// FIXME
	v1 := false
	videoLog.Println("DrawGeneralStart")
	if /*noxflags.HasEngine(noxflags.EngineWindowed) ||*/ v1 /*|| nox_video_renderTargetFlags&0x10 != 0*/ {
		videoLog.Println("DrawGeneralSkip")
		sub_4B05D0()
		return nil
	}
	legacy.Sub_431290()
	legacy.Sub_43DBD0()
	legacy.Sub_44D8F0()
	for legacy.Sub_43DC40() != 0 || legacy.Sub_44D930() {
		c.sub4312C0()
	}
	v12 := sub_48B3E0(false)
	//inpHandler.UnacquireMouse()

	c.playMovieFile(movieFilesStack[0])

	legacy.Sub_43DBE0()
	//inpHandler.AcquireMouse()
	sub_48B3E0(v12)
	sub_4B05D0()
	return nil
}

func (c *Client) playMovieFile(name string) {
	videoLog.Printf("playMovieFile: %q", name)
	if f, err := ifs.Open(name); err == nil {
		defer f.Close()
		if plr, err := noxmovie.NewPlayerWithHandle(f, noxClient.Seat, audioDev); err == nil {
			defer plr.Close()
			plr.SetAudioGain(configGetVolumeGain(VolumeMusic))
			plr.Start()
			plr.Play()
		}
	}
}

func (c *Client) clientDraw() bool {
	return c.nox_xxx_client_435F80_draw()
}

func (c *Client) copyPixBuffer() {
	c.Win.CopyBuffer(noxPixBuffer.img)
	*memmap.PtrUint32(0x973A20, 496)++
}

func (c *Client) resetRenderer(sz image.Point, init bool) error {
	if nox_video_renderTargetFlags&4 == 0 && !init {
		if err := c.Win.Reset(sz); err != nil {
			return err
		}
	}
	if nox_video_setBackBufSizes_48A3D0(sz) == 0 {
		return errors.New("nox_video_setBackBufSizes_48A3D0 failed")
	}
	return nil
}

func (c *Client) gameResetVideoMode(inMenu, force bool) error {
	var mode image.Point
	if inMenu {
		if debugMainloop {
			videoLog.Printf("gameUpdateVideoMode: menu (%s)", caller(1))
		}
		mode = videoGetMenuMode()
	} else {
		if debugMainloop {
			videoLog.Printf("gameUpdateVideoMode: game (%s)", caller(1))
		}
		mode = c.videoGetGameMode()
	}
	videoLog.Printf("mode switch: %+v (menu: %v)", mode, inMenu)
	videoResizeView(mode)
	if !force && mode == noxPixBuffer.img.Size() {
		return nil
	}
	if err := recreateBuffersAndTarget(mode); err != nil {
		return err
	}
	c.nox_xxx_loadPal_4A96C0_video_read_palette()
	legacy.Sub_461520()
	return nil
}

func (c *Client) nox_xxx_loadPal_4A96C0_video_read_palette() {
	nox_xxx_loadDefColor_4A94A0()
	for i := 0; i < 6; i++ {
		c.r.Data().SetMaterial(i, color.White)
	}
}

func (c *Client) sub_43C060() bool {
	c.nox_xxx_loadPal_4A96C0_video_read_palette()
	c.guiLoadBorderImages()
	c.nox_xxx_wndLoadMainBG_4A2210()
	c.Nox_client_setCursorType(gui.CursorSelect)
	sub_48B3E0(true)
	c.Nox_video_stopAllFades44E040()
	legacy.Sub_43E8C0(0)
	return legacy.Nox_xxx_compassGenStrings_4A9C80() != 0
}

func nox_video_setBackBufSizes_48A3D0(sz image.Point) int {
	*memmap.PtrUint32(0x973F18, 2368) = 0

	*memmap.PtrUint32(0x973F18, 136) = uint32(2 * sz.X)
	*memmap.PtrUint32(0x973F18, 2368) = 1

	*memmap.PtrUint32(0x973F18, 2348) = uint32(sz.X / 2)
	return 1
}
