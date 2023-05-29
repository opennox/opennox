package legacy

import (
	"image"
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/client"
)

var (
	Nox_video_getCutSize_4766D0             func() int
	Nox_video_setCutSize_4766A0             func(v int)
	Nox_video_setGammaSlider                func(v int)
	Sub_43BE50_get_video_mode_id            func() int
	Get_video_mode_string                   func(id int) string
	Nox_getBackbufWidth                     func() int
	Nox_getBackbufHeight                    func() int
	Nox_video_getFullScreen                 func() int
	Nox_video_setFullScreen                 func(v int)
	Sub_430C30_set_video_max                func(w, h int)
	VideoGetMaxSize                         func() image.Point
	Nox_video_callCopyBackBuffer_4AD170     func()
	Nox_getBackbufferPitch                  func() int
	Nox_client_clearScreen_440900           func()
	Nox_draw_setCutSize_476700              func(cutPerc int, a2 int)
	Nox_xxx_bookSaveSpellForDragDrop_477640 func(a1, a2 int)
	Nox_xxx_bookSpellDnDclear_477660        func()
	Nox_xxx_bookGetSpellDnDType_477670      func() int
	Nox_xxx_cursorSetDraggedItem_477690     func(a1 *client.Drawable)
	Nox_xxx_cursorResetDraggedItem_4776A0   func()
	Sub_478000                              func() int
)

type nox_render_data_t = nox_render_data_t

// nox_video_getCutSize_4766D0
func nox_video_getCutSize_4766D0() int { return Nox_video_getCutSize_4766D0() }

// nox_video_setCutSize_4766A0
func nox_video_setCutSize_4766A0(v int) { Nox_video_setCutSize_4766A0(v) }

// nox_video_setGammaSlider
func nox_video_setGammaSlider(v int) { Nox_video_setGammaSlider(v) }

// sub_43BE50_get_video_mode_id
func sub_43BE50_get_video_mode_id() int { return Sub_43BE50_get_video_mode_id() }

// get_video_mode_string
func get_video_mode_string(id int) *wchar2_t { return internWStr(Get_video_mode_string(id)) }

// nox_getBackbufWidth
func nox_getBackbufWidth() int { return Nox_getBackbufWidth() }

// nox_getBackbufHeight
func nox_getBackbufHeight() int { return Nox_getBackbufHeight() }

// nox_video_getFullScreen
func nox_video_getFullScreen() int { return Nox_video_getFullScreen() }

// nox_video_setFullScreen
func nox_video_setFullScreen(v int) { Nox_video_setFullScreen(v) }

// sub_430C30_set_video_max
func sub_430C30_set_video_max(w, h int) { Sub_430C30_set_video_max(w, h) }

// nox_xxx_screenGetSize_430C50_get_video_max
func nox_xxx_screenGetSize_430C50_get_video_max(pw, ph *int) {
	sz := VideoGetMaxSize()
	*pw = int(sz.X)
	*ph = int(sz.Y)
}

// nox_video_callCopyBackBuffer_4AD170
func nox_video_callCopyBackBuffer_4AD170() { Nox_video_callCopyBackBuffer_4AD170() }

// nox_getBackbufferPitch
func nox_getBackbufferPitch() int { return Nox_getBackbufferPitch() }

// nox_client_clearScreen_440900
func nox_client_clearScreen_440900() { Nox_client_clearScreen_440900() }

// nox_draw_setCutSize_476700
func nox_draw_setCutSize_476700(cutPerc int, a2 int) { Nox_draw_setCutSize_476700(int(cutPerc), a2) }

// nox_xxx_bookSaveSpellForDragDrop_477640
func nox_xxx_bookSaveSpellForDragDrop_477640(a1, a2 int) {
	Nox_xxx_bookSaveSpellForDragDrop_477640(a1, a2)
}

// nox_xxx_bookSpellDnDclear_477660
func nox_xxx_bookSpellDnDclear_477660() { Nox_xxx_bookSpellDnDclear_477660() }

// nox_xxx_bookGetSpellDnDType_477670
func nox_xxx_bookGetSpellDnDType_477670() int { return Nox_xxx_bookGetSpellDnDType_477670() }

// nox_xxx_cursorSetDraggedItem_477690
func nox_xxx_cursorSetDraggedItem_477690(a1 *nox_drawable) {
	Nox_xxx_cursorSetDraggedItem_477690(asDrawable(a1))
}

// nox_xxx_cursorResetDraggedItem_4776A0
func nox_xxx_cursorResetDraggedItem_4776A0() { Nox_xxx_cursorResetDraggedItem_4776A0() }

// sub_478000
func sub_478000() int { return Sub_478000() }
func Sub_4AEE30() {
	sub_4AEE30()
}
func Nox_xxx_guiSpell_460650() int {
	return int(nox_xxx_guiSpell_460650())
}
func Sub_4611A0() int {
	return int(sub_4611A0())
}
func Sub_467CD0() {
	sub_467CD0()
}
func Nox_xxx_setSomeFunc_48A210(a1 unsafe.Pointer) {
	nox_xxx_setSomeFunc_48A210(int(uintptr(a1)))
}
func Sub_49F6D0(a1 int) int {
	return int(sub_49F6D0(int(a1)))
}
func Sub_430B50(a1 int, a2 int, a3 int, a4 int) {
	sub_430B50(int(a1), int(a2), int(a3), int(a4))
}
func Sub_495A80(a1 uint32) int {
	return int(sub_495A80(int(a1)))
}

func Get_sub_47FCE0() unsafe.Pointer {
	return sub_47FCE0
}