package legacy

import (
	"image"
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/legacy/common/ccall"
	"github.com/noxworld-dev/opennox/v1/server"
)

var (
	Sub_526CA0 func(a1 string) int
)

var _ = [1]struct{}{}[12332-unsafe.Sizeof(server.WallDef{})]

func asWallP(p unsafe.Pointer) *server.Wall {
	return (*server.Wall)(p)
}

// nox_server_getWallAtGrid_410580
func nox_server_getWallAtGrid_410580(x, y int) unsafe.Pointer {
	return GetServer().S().Walls.GetWallAtGrid(image.Pt(x, y)).C()
}

// nox_xxx_wall_4105E0
func nox_xxx_wall_4105E0(x, y int) unsafe.Pointer {
	return GetServer().S().Walls.GetWallAtGrid2(image.Pt(x, y)).C()
}

// nox_xxx_wallCreateAt_410250
func nox_xxx_wallCreateAt_410250(x, y int) unsafe.Pointer {
	return GetServer().S().Walls.CreateAtGrid(image.Pt(x, y)).C()
}

// nox_xxx_mapDelWallAtPt_410430
func nox_xxx_mapDelWallAtPt_410430(x, y int) {
	GetServer().S().Walls.DeleteAtGrid(image.Pt(x, y))
}

// sub_4106A0
func sub_4106A0(y int) unsafe.Pointer {
	return GetServer().S().Walls.IndexByY(y).C()
}

// nox_xxx_wallForeachFn_410640
func nox_xxx_wallForeachFn_410640(cfnc unsafe.Pointer, data unsafe.Pointer) {
	GetServer().S().Walls.EachWallXxx(func(it *server.Wall) {
		ccall.CallVoidPtr2(cfnc, it.C(), data)
	})
}

// sub_57B500
func sub_57B500(x, y int, flags char) char {
	return char(GetServer().S().Sub_57B500(image.Pt(x, y), byte(int8(flags))))
}

// sub_4D72C0
func sub_4D72C0() int {
	return int(bool2int(GetServer().S().Doors.Sub_4D72C0()))
}

// sub_4D72B0
func sub_4D72B0(v int) {
	GetServer().S().Doors.Sub_4D72B0(v != 0)
}

// nox_xxx_wallFlags
func nox_xxx_wallFlags(ind int) uint32 {
	return GetServer().S().Walls.DefByInd(ind).Flags32
}

// nox_xxx_wallGetBrickObj_410E60
func nox_xxx_wallGetBrickObj_410E60(ind int, ind2 int) *char {
	return internCStr(GetServer().S().Walls.DefByInd(ind).BrickObject(ind2))
}

// nox_xxx_getWallSprite_46A3B0
func nox_xxx_getWallSprite_46A3B0(ind int, a2 int, a3 int, a4 int) unsafe.Pointer {
	return GetServer().S().Walls.DefByInd(ind).Sprite(a2, a3, a4)
}

// nox_xxx_getWallDrawOffset_46A3F0
func nox_xxx_getWallDrawOffset_46A3F0(ind int, a2 int, a3 int, a4 int, px, py *int) {
	v := GetServer().S().Walls.DefByInd(ind).DrawOffset(a2, a3, a4)
	*px = int(v.X)
	*py = int(v.Y)
}

// nox_xxx_mapWallMaxVariation_410DD0
func nox_xxx_mapWallMaxVariation_410DD0(ind int, a2 int, a3 int) byte {
	return GetServer().S().Walls.DefByInd(ind).Variations(a2, a3)
}

// nox_xxx_map_410E00
func nox_xxx_map_410E00(ind int) byte {
	return GetServer().S().Walls.DefByInd(ind).Field749
}

// nox_xxx_mapWallGetHpByTile_410E20
func nox_xxx_mapWallGetHpByTile_410E20(ind int) byte {
	return GetServer().S().Walls.DefByInd(ind).Health41
}

// nox_xxx_wallGetBrickTypeMB_410E40
func nox_xxx_wallGetBrickTypeMB_410E40(ind int) byte {
	return GetServer().S().Walls.DefByInd(ind).BrickType42
}

// nox_xxx_wallField36
func nox_xxx_wallField36(ind int) uint16 {
	return GetServer().S().Walls.DefByInd(ind).Field36
}

// nox_xxx_wallSoundByTile_410EA0
func nox_xxx_wallSoundByTile_410EA0(ind int) *char {
	return internCStr(GetServer().S().Walls.DefByInd(ind).BreakSound())
}

// nox_xxx_wallFindOpenSound_410EE0
func nox_xxx_wallFindOpenSound_410EE0(ind int) *char {
	return internCStr(GetServer().S().Walls.DefByInd(ind).OpenSound())
}

// nox_xxx_wallFindCloseSound_410F20
func nox_xxx_wallFindCloseSound_410F20(ind int) *char {
	return internCStr(GetServer().S().Walls.DefByInd(ind).CloseSound())
}

// nox_xxx_wallTileByName_410D60
func nox_xxx_wallTileByName_410D60(name *char) int {
	return GetServer().S().Walls.DefIndByName(GoString(name))
}

// sub_526CA0
func sub_526CA0(a1 *char) int {
	return Sub_526CA0(GoString(a1))
}

func Nox_xxx_wallClose_512070(a1 unsafe.Pointer) {
	nox_xxx_wallClose_512070(int(uintptr(a1)))
}
func Nox_xxx_wallOpen_511F80(a1 unsafe.Pointer) {
	nox_xxx_wallOpen_511F80(int(uintptr(a1)))
}
func Nox_xxx_wallToggle_512160(a1 unsafe.Pointer) {
	nox_xxx_wallToggle_512160(int(uintptr(a1)))
}
func Nox_xxx_wallPreDestroyByPtr_5122C0(a1 unsafe.Pointer) {
	nox_xxx_wallPreDestroyByPtr_5122C0(int(uintptr(a1)))
}
