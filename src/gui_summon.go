package opennox

/*
extern int dword_5d4594_3799524;
*/
import "C"
import (
	"image"
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/common/memmap"
)

//export sub_4C26F0
func sub_4C26F0(a1 unsafe.Pointer) C.int {
	win := asWindowP(a1)
	pos := win.GlobalPos()
	x, y := pos.X, pos.Y
	sz := win.Size()
	w, h := sz.X, sz.Y
	r := noxrend
	r.DrawRectFilledAlpha(x, y, w, h)
	r.Data().SetColor2(memmap.Uint32(0x85B3FC, 944))
	r.nox_client_drawAddPoint_49F500(image.Point{X: x + 1, Y: y})
	r.nox_xxx_rasterPointRel_49F570(image.Point{X: w - 2, Y: 0})
	r.DrawLineFromPoints()
	r.nox_client_drawAddPoint_49F500(image.Point{X: x + 1, Y: y + h})
	r.nox_xxx_rasterPointRel_49F570(image.Point{X: w - 2, Y: 0})
	r.DrawLineFromPoints()
	r.nox_client_drawAddPoint_49F500(image.Point{X: x, Y: y + 1})
	r.nox_xxx_rasterPointRel_49F570(image.Point{X: 0, Y: h - 2})
	r.DrawLineFromPoints()
	r.nox_client_drawAddPoint_49F500(image.Point{X: x + w, Y: y + 1})
	r.nox_xxx_rasterPointRel_49F570(image.Point{X: 0, Y: h - 2})
	r.DrawLineFromPoints()
	C.dword_5d4594_3799524 = 1
	return 1
}