package legacy

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"unsafe"
)

func nox_xxx_netReportGuideAward_4D8000(a1 int32, a2 int8, a3 int8, a4 int32) int32 {
	var (
		result int32
		v5     int32
	)
	result = a1
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4 != 0 {
		v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 1)) = uint8(a2)
		*((*uint8)(unsafe.Pointer(&a1))) = 209
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 2)) = uint8(a3)
		if a4 != 0 {
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 2)) = uint8(int8(int32(a3) | 0x80))
		}
		result = nox_xxx_netSendPacket1_4E5390(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v5 + 276))) + 2064)))), int32(uintptr(unsafe.Pointer(&a1))), 3, 0, 1)
	}
	return result
}
func nox_xxx_awardBeastGuide_4FAE80_magic_plyrgide(a1 int32, a2 int32, a3 int32) int32 {
	var (
		v3 int32
		v4 int32
		v5 int32
		v6 int32
		v7 *uint32
	)
	_ = v7
	var v8 *uint8
	var v9 *uint32
	var v10 int32
	var v11 *int32
	var v12 int32
	var i *byte
	var v14 int32
	var v16 *wchar2_t
	var v17 *uint8
	v3 = a1
	if (int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8)))) & 4) == 0 {
		return 0
	}
	if a2 <= 0 || a2 >= 41 {
		v16 = nox_strman_loadString_40F1D0(internCStr("AwardGuideError"), nil, internCStr("C:\\NoxPost\\src\\Server\\Magic\\PlyrGide.c"), 39)
		nox_xxx_netSendLineMessage_4D9EB0(a1, v16)
		return 0
	}
	v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 276))))
	v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + a2*4 + 4244))))
	v7 = (*uint32)(unsafe.Pointer(uintptr(v5 + a2*4 + 4244)))
	if v6 != 0 {
		return 0
	}
	*v7 = 1
	nox_xxx_playerAwardSpellProtectionCRC_56FCE0(int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 276))) + 4640)))), a2, int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 276))) + uint32(a2*4) + 4244)))))
	if a3 != 0 {
		nox_xxx_aud_501960(227, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
		nox_xxx_netSendRewardNotify_4FAD50(a1, 1, a1, int8(a2))
	}
	v8 = (*uint8)(memmap.PtrOff(0x587000, 216292))
	v17 = (*uint8)(memmap.PtrOff(0x587000, 216292))
	if *memmap.PtrUint32(0x587000, 216292) != 0 {
		for {
			v9 = *(**uint32)(unsafe.Pointer(v8))
			if uint32(a2) == **(**uint32)(unsafe.Pointer(v8)) {
				v10 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v9), 4*1)))
				v11 = (*int32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(v9), 4*1))))
				if v10 != 0 {
					for {
						*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 276))) + uint32(v10*4) + 4244))) = 1
						nox_xxx_playerAwardSpellProtectionCRC_56FCE0(int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 276))) + 4640)))), *v11, int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 276))) + uint32(*v11*4) + 4244)))))
						v10 = *(*int32)(unsafe.Add(unsafe.Pointer(v11), 4*1))
						v11 = (*int32)(unsafe.Add(unsafe.Pointer(v11), 4*1))
						if v10 == 0 {
							break
						}
					}
					v8 = v17
				}
			}
			v12 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v8))), 4*1))))
			v8 = (*uint8)(unsafe.Add(unsafe.Pointer(v8), 4))
			v17 = v8
			if v12 == 0 {
				break
			}
		}
	}
	if a3 != 0 {
		for i = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0())); i != nil; i = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(i))))))))) {
			v14 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), 4*514))))
			if v14 != v3 {
				if v14 != 0 {
					nox_xxx_netSendRewardNotify_4FAD50(v14, 1, v3, int8(a2))
				}
			}
		}
	}
	nox_xxx_netReportGuideAward_4D8000(v3, int8(a2), int8(a3), 0)
	return 1
}