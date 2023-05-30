package legacy

import (
	"unsafe"

	"github.com/gotranspile/cxgo/runtime/stdio"

	"github.com/noxworld-dev/opennox/v1/common/memmap"
)

func mem_getPtrSize(base uint32, off uint32, size uint32) unsafe.Pointer {
	switch base {
	case 0x581450:
		if off+size <= uint32(23472) {
			return unsafe.Pointer(&byte_581450[off])
		}
	case 0x587000:
		if off+size <= uint32(316820) {
			return unsafe.Pointer(&byte_587000[off])
		}
	case 0x5D4594:
		if off+size <= uint32(2598284) {
			return unsafe.Pointer(&byte_5D4594[off])
		}
	case 0x8531A0:
		if off+size <= uint32(2592) {
			return unsafe.Pointer(&byte_8531A0[off])
		}
	case 0x973CE0:
		if off+size <= uint32(568) {
			return unsafe.Pointer(&byte_973CE0[off])
		}
	case 0x973F18:
		if off+size <= uint32(44881) {
			return unsafe.Pointer(&byte_973F18[off])
		}
	case 0x85B3FC:
		if off+size <= uint32(1029636) {
			return unsafe.Pointer(&byte_85B3FC[off])
		}
	case 0x852978:
		if off+size <= uint32(40) {
			return unsafe.Pointer(&byte_852978[off])
		}
	case 0x973A20:
		if off+size <= uint32(704) {
			return unsafe.Pointer(&byte_973A20[off])
		}
	}
	stdio.Fprintf(stdio.Stderr(), "Invalid memory access! Requested = %x+%d[%d]\n", base, off, size)
	panic("abort")
	return nil
}
func mem_getPtr(base uint32, off uint32) unsafe.Pointer {
	return mem_getPtrSize(base, off, uint32(4))
}
func mem_getU8Ptr(base uintptr, off uint32) *uint8 {
	return (*uint8)(memmap.PtrSizeOff(base, uintptr(off), 1))
}
func mem_getI8Ptr(base uintptr, off uint32) *int8 {
	return (*int8)(memmap.PtrSizeOff(base, uintptr(off), 1))
}
func mem_getU16Ptr(base uintptr, off uint32) *uint16 {
	return (*uint16)(memmap.PtrSizeOff(base, uintptr(off), 2))
}
func mem_getI16Ptr(base uintptr, off uint32) *int16 {
	return (*int16)(memmap.PtrSizeOff(base, uintptr(off), 2))
}
func mem_getU32Ptr(base uintptr, off uint32) *uint32 {
	return (*uint32)(memmap.PtrSizeOff(base, uintptr(off), 4))
}
func mem_getI32Ptr(base uintptr, off uint32) *int32 {
	return (*int32)(memmap.PtrSizeOff(base, uintptr(off), 4))
}
func mem_getU64Ptr(base uintptr, off uint32) *uint64 {
	return (*uint64)(memmap.PtrSizeOff(base, uintptr(off), 8))
}
func mem_getI64Ptr(base uintptr, off uint32) *int64 {
	return (*int64)(memmap.PtrSizeOff(base, uintptr(off), 8))
}
func mem_getFloatPtr(base uintptr, off uint32) *float32 {
	return (*float32)(memmap.PtrSizeOff(base, uintptr(off), 4))
}
func mem_getDoublePtr(base uintptr, off uint32) *float64 {
	return (*float64)(memmap.PtrSizeOff(base, uintptr(off), 8))
}
