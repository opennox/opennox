//go:build !server

package opennox

import (
	"image"
	"unicode/utf16"

	"github.com/tawesoft/golib/v2/dialog"

	"github.com/noxworld-dev/opennox-lib/client/seat/sdl"
	"github.com/noxworld-dev/opennox-lib/env"

	"github.com/noxworld-dev/opennox/v1/client/input"
	"github.com/noxworld-dev/opennox/v1/client/render"
	"github.com/noxworld-dev/opennox/v1/common/config"
	"github.com/noxworld-dev/opennox/v1/internal/version"
	"github.com/noxworld-dev/opennox/v1/legacy"
)

func init() {
	config.Global.SetDefault(configVideoFiltering, true)
	config.Global.SetDefault(configVideoStretch, false)
}

func (c *Client) initSeat(sz image.Point) error {
	sst, err := sdl.New("OpenNox "+version.ClientVersion(), sz)
	if err != nil {
		return err
	}
	c.Seat = sst
	if env.IsE2E() {
		c.Seat = e2eWrapSeat(c.Seat)
	}
	c.Win, err = render.New(c.Seat)
	if err != nil {
		_ = c.Seat.Close()
		return err
	}

	inp := input.New(c.Seat, false, c.Strings().Lang())
	c.Inp = inp

	inp.OnQuit(mainloopStop)
	inp.OnToggleFullScreen(c.Win.ToggleWindowMode)
	inp.OnKeyPress(gameexOnKeyboardPress)
	inp.OnMouseWheel(func(dv int) {
		// mix event handler is triggered only for wheel events
		call_OnLibraryNotice_265(dv)
	})
	inp.OnInputString(func(str string) {
		for _, c := range utf16.Encode([]rune(str)) {
			legacy.NoxInputOnChar(c)
		}
	})

	c.Win.OnViewResize(inp.SetWinSize)
	OnPixBufferResize(inp.SetDrawWinSize)

	c.Win.SetFiltering(config.Global.GetBool(configVideoFiltering))
	c.Win.SetStretched(config.Global.GetBool(configVideoStretch))
	if err != nil {
		return err
	}
	sst.SetGamma(getGamma())
	return nil
}

func (c *Client) freeSeat() {
	if c.Seat != nil {
		c.Seat.Close()
		c.Seat = nil
	}
}

func errorMessage(format string, args ...any) {
	dialog.Error(format, args...)
}
