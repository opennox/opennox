package opennox

import (
	"context"
	"image"
	"strconv"

	"github.com/noxworld-dev/opennox-lib/client/seat"
	"github.com/noxworld-dev/opennox-lib/console"
	"github.com/noxworld-dev/opennox-lib/log"

	"github.com/noxworld-dev/opennox/v1/common/config"
)

const (
	noxDefaultWidth  = seat.DefaultWidth
	noxDefaultHeight = seat.DefaultHeight
	noxDefaultDepth  = seat.DefaultDepth
)

const (
	gammaMin = 0.1
	gammaDef = 1.0
	gammaMax = 3.0
	gammaInc = 0.1
)

var (
	videoLog         = log.New("video")
	noxVideoModeMenu = image.Point{
		X: noxDefaultWidth,
		Y: noxDefaultHeight,
	}
	noxVideoMax = image.Point{
		X: noxMaxWidth,
		Y: noxMaxHeight,
	}
	noxVideoModes = []image.Point{
		{640, 480},
		{800, 600},
		{1024, 768},
		{1280, 720},
		{1920, 1080},
		{2560, 1440},
		{3840, 2160},
	}
	noxVideoGamma = float32(gammaDef)

	nox_video_cutSize = 100

	g_scaled_cfg     = 0
	g_fullscreen_cfg = 0
)

func init() {
	noxConsole.Register(&console.Command{
		Token:  "gamma",
		HelpID: "gammahelp",
		Flags:  console.ClientServer,
		Func: func(ctx context.Context, c *console.Console, tokens []string) bool {
			if len(tokens) == 0 {
				c.Printf(console.ColorLightYellow, "gamma = %d", getGammaSlider())
				return true
			} else if len(tokens) > 1 {
				return false
			}
			s := tokens[0]
			if len(s) == 0 {
				return false
			}
			add := 0
			if s[0] == '+' {
				add = +1
				s = s[1:]
			} else if s[0] == '-' {
				add = -1
				s = s[1:]
			}
			v, err := strconv.Atoi(s)
			if err != nil {
				return false
			}
			gamma := v
			if add != 0 {
				gamma = getGammaSlider() + add*gamma
			}
			setGammaSlider(gamma)
			return true
		},
	})
}

func videoModeID() int {
	mode := noxClient.videoGetGameMode()
	for i, res := range noxVideoModes {
		if mode == res {
			return i
		}
	}
	for i, res := range noxVideoModes {
		if mode.X == res.X {
			return i
		}
	}
	for i, res := range noxVideoModes {
		if mode.Y == res.Y {
			return i
		}
	}
	area := mode.X * mode.Y
	min := 0
	ind := -1
	for i, res := range noxVideoModes {
		area2 := res.X * res.Y
		if ind == -1 || abs(area-area2) < min {
			ind = i
			min = abs(area - area2)
		}
	}
	return ind
}

func videoModeByID(code byte) (image.Point, bool) {
	if int(code) >= len(noxVideoModes) {
		return image.Point{}, false
	}
	return noxVideoModes[code], true
}

func videoGetMenuMode() image.Point {
	return noxVideoModeMenu
}

func videoSetMenuMode(mode image.Point) {
	noxVideoModeMenu = mode
}

func videoSetMaxSize(sz image.Point) {
	noxVideoMax = sz
}

func videoGetMaxSize() image.Point {
	return noxVideoMax
}

func videoResizeView(mode image.Point) {
	max := videoGetMaxSize()
	if mode.X > max.X {
		mode.X = max.X
	}
	if mode.Y > max.Y {
		mode.Y = max.Y
	}
	videoSetWindowSize(mode)
}

func changeWindowedOrFullscreen() {
	cfgUpdateFullScreen()
}

func nox_video_setCutSize(v int) {
	if v < 65 {
		v = 65
	}
	if v > 100 {
		v = 100
	}
	nox_video_cutSize = v
}

func nox_video_getCutSize() int {
	return nox_video_cutSize
}

func getGammaSlider() int {
	return int(100 * (getGamma() - gammaMin) / gammaMax)
}

func setGammaSlider(v int) {
	fv := float32(v) / 100
	setGamma(gammaMin + gammaMax*fv)
}

func setGammaSliderOpts(v int) {
	setGammaSlider(v)
	config.Global.Set(configVideoGamma, getGamma())
}

func updateGamma(value int) {
	var dv float32
	if value > 0 {
		dv = +gammaInc
	} else if value < 0 {
		dv = -gammaInc
	}
	setGamma(getGamma() + dv)
}

func getGamma() float32 {
	return noxVideoGamma
}

func setGamma(v float32) {
	if v < gammaMin {
		v = gammaMin
	} else if v > gammaMax {
		v = gammaMax
	}
	if noxVideoGamma == v {
		return
	}
	videoLog.Printf("gamma: %v", v)
	noxVideoGamma = v
	noxClient.SetScreenGamma(v)
}
