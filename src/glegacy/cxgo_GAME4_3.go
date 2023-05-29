package legacy

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/gotranspile/cxgo/runtime/stdio"
	"math"
	"unsafe"
)

func nox_xxx_onFrameLightning_52F8A0(a1 float32) int32 {
	var (
		source              int32
		v2                  int32
		v4                  int32
		v5                  int32
		v6                  int32
		v7                  int32
		owner               int32
		spellLevel          int32
		v10                 int32
		v11                 int32
		v12                 int32
		target              int32
		index               int32
		v15                 int32
		v16                 int32
		v17                 int32
		v18                 int32
		secondBounceTarget  int32
		v20                 int32
		v21                 *uint32
		j                   *uint32
		v23                 int32
		v24                 int32
		v25                 int32
		v26                 int32
		v27                 int32
		v28                 int8
		v29                 uint8
		v30                 int32
		v31                 int32
		i                   int32
		v33                 int32
		v34                 int32
		range1              float32
		range2              float32
		range3              float32
		range4              float32
		lightningSearchTime float32
		range5              float32
	)
	_ = range5
	var lightningRange float32
	var damage float32
	source = int32(*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&a1))), 4*0)))
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr((*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&a1))), 4*0))) + 16))))
	if v2 != 0 {
		if nox_xxx_testUnitBuffs_4FF350((*nox_object_t)(unsafe.Pointer(uintptr(v2))), 8) != 0 {
			return 1
		}
	} else if *(*uint32)(unsafe.Pointer(uintptr((*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&a1))), 4*0))) + 20))) == 0 {
		return 1
	}
	lightningRange = float32(nox_xxx_gamedataGetFloat_419D40(internCStr("LightningRange")))
	if *(*uint32)(unsafe.Pointer(uintptr(source + 20))) != 0 {
		*memmap.PtrUint32(0x5D4594, 2487820) = *(*uint32)(unsafe.Pointer(uintptr(source + 28)))
		*memmap.PtrUint32(0x5D4594, 2487824) = *(*uint32)(unsafe.Pointer(uintptr(source + 32)))
		nox_xxx_unitsGetInCircle_517F90((*float2)(unsafe.Pointer(uintptr(source+28))), lightningRange, unsafe.Pointer(funAddr(nox_xxx_lightningSpellTrapEffect_530020)), unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(source + 16))))))
		return 1
	}
	if int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(source + 16))) + 8))))&4 != 0 && int32(nox_xxx_unitGetOldMana_4EEC80(int32(*(*uint32)(unsafe.Pointer(uintptr(source + 16)))))) == 0 {
		return 1
	}
	if (gameFrame()-*(*uint32)(unsafe.Pointer(uintptr(source + 60)))) > 2 && sub_4E6BD0(int32(*(*uint32)(unsafe.Pointer(uintptr(source + 16))))) != 0 {
		return 1
	}
	v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(source + 16))))
	if int32(*(*uint8)(unsafe.Pointer(uintptr(v4 + 8))))&2 != 0 && sub_4FEA70(v4, (*float2)(unsafe.Pointer(uintptr(source+28)))) != 0 {
		return 1
	}
	v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(source + 104))))
	if v5 != 0 {
		for {
			v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 116))))
			sub_4FE980(unsafe.Pointer(uintptr(v5)))
			v5 = v6
			if v6 == 0 {
				break
			}
		}
	}
	v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(source + 8))))
	*(*uint32)(unsafe.Pointer(uintptr(source + 104))) = *(*uint32)(unsafe.Pointer(uintptr(source + 108)))
	*(*uint32)(unsafe.Pointer(uintptr(source + 108))) = 0
	nox_xxx_lightningTarget_5d4594_2487908 = 0
	nox_xxx_lightningTargetArrayIndex_5d4594_2487904 = 0
	owner = int32(*(*uint32)(unsafe.Pointer(uintptr(source + 16))))
	*memmap.PtrUint32(0x5D4594, 2487844) = 0
	spellLevel = int32(*memmap.PtrUint32(0x587000, uint32(v7*4)+260380))
	*memmap.PtrUint32(0x5D4594, 2487848) = 0
	nox_xxx_lightningOwner_5d4594_2487900 = uint32(owner)
	*memmap.PtrUint32(0x5D4594, 2487852) = 0
	*memmap.PtrUint32(0x5D4594, 2487856) = 0
	*memmap.PtrUint32(0x5D4594, 2487860) = 0
	v10 = int32(*(*uint32)(unsafe.Pointer(uintptr(source + 16))))
	if (int32(*(*uint8)(unsafe.Pointer(uintptr(v10 + 8))))&4) == 0 || (func() bool {
		v11 = int32(*(*uint32)(unsafe.Pointer(uintptr(v10 + 748))))
		return (func() int32 {
			v12 = int32(*(*uint32)(unsafe.Pointer(uintptr(v11 + 288))))
			return v12
		}()) == 0
	}()) || (func() bool {
		if nox_xxx_unitIsEnemyTo_5330C0((*nox_object_t)(unsafe.Pointer(uintptr(v10))), (*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v11 + 288))))))) == 0 || nox_xxx_calcDistance_4E6C00((*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(source + 16)))))), (*nox_object_t)(unsafe.Pointer(uintptr(v12)))) > float64(lightningRange) {
			target = int32(nox_xxx_lightningTarget_5d4594_2487908)
		} else {
			target = v12
			nox_xxx_lightningTarget_5d4594_2487908 = uint32(v12)
		}
		return target == 0
	}()) {
		*(*float32)(unsafe.Pointer(&nox_xxx_lightningClosestTargetDistance_5d4594_2487912)) = lightningRange * lightningRange
		nox_xxx_unitsGetInCircle_517F90((*float2)(unsafe.Pointer(uintptr(source+28))), lightningRange, unsafe.Pointer(funAddr(nox_xxx_lightningCanAttackCheck_52FF10)), unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(source + 16))))))
		target = int32(nox_xxx_lightningTarget_5d4594_2487908)
		if nox_xxx_lightningTarget_5d4594_2487908 == 0 {
			for i = int32(*(*uint32)(unsafe.Pointer(uintptr(source + 104)))); i != 0; i = int32(*(*uint32)(unsafe.Pointer(uintptr(i + 116)))) {
				if *(*uint32)(unsafe.Pointer(uintptr(i + 48))) != 0 {
					nox_xxx_netStopRaySpell_4FEF90(i, *(**uint32)(unsafe.Pointer(uintptr(i + 48))))
				}
			}
			v33 = int32(*(*uint32)(unsafe.Pointer(uintptr(source + 104))))
			if v33 != 0 {
				for {
					v34 = int32(*(*uint32)(unsafe.Pointer(uintptr(v33 + 116))))
					sub_4FE980(unsafe.Pointer(uintptr(v33)))
					v33 = v34
					if v34 == 0 {
						break
					}
				}
			}
			*(*uint32)(unsafe.Pointer(uintptr(source + 104))) = 0
			return 0
		}
	}
	index = int32(nox_xxx_lightningTargetArrayIndex_5d4594_2487904)
	*memmap.PtrUint32(0x5D4594, nox_xxx_lightningTargetArrayIndex_5d4594_2487904*4+2487844) = uint32(target)
	nox_xxx_lightningTargetArrayIndex_5d4594_2487904 = uint32(index + 1)
	if spellLevel > 1 {
		nox_xxx_lightningTarget_5d4594_2487908 = 0
		*(*float32)(unsafe.Pointer(&nox_xxx_lightningClosestTargetDistance_5d4594_2487912)) = lightningRange * lightningRange
		range1 = float32(float64(lightningRange) * 0.94999999)
		nox_xxx_unitsGetInCircle_517F90((*float2)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x5D4594, 2487844)+56))), range1, unsafe.Pointer(funAddr(nox_xxx_lightningCanAttackCheck_52FF10)), unsafe.Pointer(uintptr(*memmap.PtrInt32(0x5D4594, 2487844))))
		if nox_xxx_lightningTarget_5d4594_2487908 != 0 {
			v15 = int32(nox_xxx_lightningTargetArrayIndex_5d4594_2487904)
			*memmap.PtrUint32(0x5D4594, nox_xxx_lightningTargetArrayIndex_5d4594_2487904*4+2487844) = nox_xxx_lightningTarget_5d4594_2487908
			nox_xxx_lightningTargetArrayIndex_5d4594_2487904 = uint32(v15 + 1)
		}
	}
	if spellLevel > 2 {
		nox_xxx_lightningTarget_5d4594_2487908 = 0
		*(*float32)(unsafe.Pointer(&nox_xxx_lightningClosestTargetDistance_5d4594_2487912)) = lightningRange * lightningRange
		range2 = float32(float64(lightningRange) * 0.89999998)
		nox_xxx_unitsGetInCircle_517F90((*float2)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x5D4594, 2487844)+56))), range2, unsafe.Pointer(funAddr(nox_xxx_lightningCanAttackCheck_52FF10)), unsafe.Pointer(uintptr(*memmap.PtrInt32(0x5D4594, 2487844))))
		if nox_xxx_lightningTarget_5d4594_2487908 != 0 {
			v16 = int32(nox_xxx_lightningTargetArrayIndex_5d4594_2487904)
			*memmap.PtrUint32(0x5D4594, nox_xxx_lightningTargetArrayIndex_5d4594_2487904*4+2487844) = nox_xxx_lightningTarget_5d4594_2487908
			nox_xxx_lightningTargetArrayIndex_5d4594_2487904 = uint32(v16 + 1)
		}
	}
	if *memmap.PtrUint32(0x5D4594, 2487848) != 0 {
		if spellLevel > 3 {
			nox_xxx_lightningTarget_5d4594_2487908 = 0
			*(*float32)(unsafe.Pointer(&nox_xxx_lightningClosestTargetDistance_5d4594_2487912)) = lightningRange * lightningRange
			range3 = float32(float64(lightningRange) * 0.85000002)
			nox_xxx_unitsGetInCircle_517F90((*float2)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x5D4594, 2487848)+56))), range3, unsafe.Pointer(funAddr(nox_xxx_lightningCanAttackCheck_52FF10)), unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(mem_getU32Ptr(0x5D4594, 2487848))))))
			if nox_xxx_lightningTarget_5d4594_2487908 != 0 {
				v17 = int32(nox_xxx_lightningTargetArrayIndex_5d4594_2487904)
				*memmap.PtrUint32(0x5D4594, nox_xxx_lightningTargetArrayIndex_5d4594_2487904*4+2487844) = nox_xxx_lightningTarget_5d4594_2487908
				nox_xxx_lightningTargetArrayIndex_5d4594_2487904 = uint32(v17 + 1)
			}
		}
	}
	if *memmap.PtrUint32(0x5D4594, 2487852) != 0 {
		if spellLevel > 4 {
			nox_xxx_lightningTarget_5d4594_2487908 = 0
			range5 = lightningRange * lightningRange
			*(*float32)(unsafe.Pointer(&nox_xxx_lightningClosestTargetDistance_5d4594_2487912)) = lightningRange * lightningRange
			range4 = float32(float64(lightningRange) * 0.80000001)
			nox_xxx_unitsGetInCircle_517F90((*float2)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x5D4594, 2487852)+56))), range4, unsafe.Pointer(funAddr(nox_xxx_lightningCanAttackCheck_52FF10)), unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(mem_getU32Ptr(0x5D4594, 2487852))))))
			if nox_xxx_lightningTarget_5d4594_2487908 != 0 {
				v18 = int32(nox_xxx_lightningTargetArrayIndex_5d4594_2487904)
				*memmap.PtrUint32(0x5D4594, nox_xxx_lightningTargetArrayIndex_5d4594_2487904*4+2487844) = nox_xxx_lightningTarget_5d4594_2487908
				nox_xxx_lightningTargetArrayIndex_5d4594_2487904 = uint32(v18 + 1)
			}
		}
	}
	nox_xxx_lightningSpellDuration_52FFD0(source, int32(*(*uint32)(unsafe.Pointer(uintptr(source + 16)))), *memmap.PtrInt32(0x5D4594, 2487844))
	if spellLevel > 1 && *memmap.PtrUint32(0x5D4594, 2487848) != 0 {
		nox_xxx_lightningSpellDuration_52FFD0(source, *memmap.PtrInt32(0x5D4594, 2487844), int32(*memmap.PtrUint32(0x5D4594, 2487848)))
	}
	secondBounceTarget = int32(*memmap.PtrUint32(0x5D4594, 2487852))
	if spellLevel > 2 && *memmap.PtrUint32(0x5D4594, 2487852) != 0 {
		nox_xxx_lightningSpellDuration_52FFD0(source, *memmap.PtrInt32(0x5D4594, 2487844), int32(*memmap.PtrUint32(0x5D4594, 2487852)))
		secondBounceTarget = int32(*memmap.PtrUint32(0x5D4594, 2487852))
	}
	if spellLevel > 3 && *memmap.PtrUint32(0x5D4594, 2487856) != 0 {
		if *memmap.PtrUint32(0x5D4594, 2487848) != 0 {
			nox_xxx_lightningSpellDuration_52FFD0(source, int32(*memmap.PtrUint32(0x5D4594, 2487848)), int32(*memmap.PtrUint32(0x5D4594, 2487856)))
			secondBounceTarget = int32(*memmap.PtrUint32(0x5D4594, 2487852))
			goto LABEL_55
		}
		if secondBounceTarget != 0 {
			nox_xxx_lightningSpellDuration_52FFD0(source, secondBounceTarget, int32(*memmap.PtrUint32(0x5D4594, 2487856)))
			secondBounceTarget = int32(*memmap.PtrUint32(0x5D4594, 2487852))
			goto LABEL_55
		}
	}
LABEL_55:
	if spellLevel > 4 {
		if *memmap.PtrUint32(0x5D4594, 2487860) != 0 {
			if secondBounceTarget != 0 || (func() int32 {
				secondBounceTarget = int32(*memmap.PtrUint32(0x5D4594, 2487848))
				return secondBounceTarget
			}()) != 0 {
				nox_xxx_lightningSpellDuration_52FFD0(source, secondBounceTarget, *memmap.PtrInt32(0x5D4594, 2487860))
			}
		}
	}
	if *memmap.PtrUint32(0x5D4594, 2487844) == 0 {
		return 0
	}
	damage = float32(nox_xxx_gamedataGetFloat_419D40(internCStr("LightningDamage")) + float64(*(*float32)(unsafe.Pointer(uintptr(source + 76)))))
	v20 = nox_float2int(damage)
	*(*float32)(unsafe.Pointer(uintptr(source + 76))) = float32(float64(damage) - float64(v20))
	v21 = *(**uint32)(unsafe.Pointer(uintptr(source + 108)))
	for j = *(**uint32)(unsafe.Pointer(uintptr(source + 104))); v21 != nil; v21 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v21), 4*29))))) {
		if j != nil {
			v23 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(j), 4*12)))
			if *(*uint32)(unsafe.Add(unsafe.Pointer(v21), 4*12)) != uint32(v23) || *(*uint32)(unsafe.Add(unsafe.Pointer(v21), 4*4)) != *(*uint32)(unsafe.Add(unsafe.Pointer(j), 4*4)) {
				if v23 != 0 {
					nox_xxx_netStopRaySpell_4FEF90(int32(uintptr(unsafe.Pointer(j))), (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(j), 4*12))))))
				}
				nox_xxx_netStartDurationRaySpell_4FF130(int32(uintptr(unsafe.Pointer(v21))))
			}
			j = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(j), 4*29)))))
		} else {
			nox_xxx_netStartDurationRaySpell_4FF130(int32(uintptr(unsafe.Pointer(v21))))
		}
		if v20 > 0 {
			(*(*func(uint32, uint32, uint32, int32, int32))(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v21), 4*12)) + 716))))(*(*uint32)(unsafe.Add(unsafe.Pointer(v21), 4*12)), *(*uint32)(unsafe.Pointer(uintptr(source + 16))), 0, v20, 17)
		}
		v24 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v21), 4*12)))
		if *(*uint32)(unsafe.Pointer(uintptr(v24 + 16)))&0x8020 != 0 {
			nox_xxx_netSendPointFx_522FF0(-127, (*float2)(unsafe.Pointer(uintptr(v24+56))))
		}
	}
	for ; j != nil; j = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(j), 4*29))))) {
		if *(*uint32)(unsafe.Add(unsafe.Pointer(j), 4*12)) != 0 {
			nox_xxx_netStopRaySpell_4FEF90(int32(uintptr(unsafe.Pointer(j))), (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(j), 4*12))))))
		}
	}
	v25 = int32(*(*uint32)(unsafe.Pointer(uintptr(source + 16))))
	if int32(*(*uint8)(unsafe.Pointer(uintptr(v25 + 8))))&4 != 0 {
		v26 = int32(*(*uint32)(unsafe.Pointer(uintptr(source + 72))))
		if v26 != 0 {
			v27 = int32(*(*uint32)(unsafe.Pointer(uintptr(v26 + 736))))
			v28 = int8(*(*uint8)(unsafe.Pointer(uintptr(v27 + 108))))
			if int32(v28) == 0 {
				return 1
			}
			v29 = uint8(int8(int32(v28) - 1))
			*(*uint8)(unsafe.Pointer(uintptr(v27 + 108))) = v29
			*(*uint32)(unsafe.Pointer(uintptr(v27 + 112))) = uint32(int32(v29) * 100 / int32(*(*uint8)(unsafe.Pointer(uintptr(v27 + 109)))))
			v30 = int32(*(*uint32)(unsafe.Pointer(uintptr(source + 16))))
			if v30 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(v30 + 8))))&4 != 0 {
				v31 = int32(*(*uint32)(unsafe.Pointer(uintptr(v30 + 748))))
				nox_xxx_playerSetState_4FA020((*nox_object_t)(unsafe.Pointer(uintptr(v30))), 22)
				nox_xxx_netReportCharges_4D82B0(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v31 + 276))) + 2064)))), (*nox_object_t)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(uintptr(source + 72))))), int8(*(*uint8)(unsafe.Pointer(uintptr(v27 + 108)))), int8(*(*uint8)(unsafe.Pointer(uintptr(v27 + 109)))))
			}
			if int32(*(*uint8)(unsafe.Pointer(uintptr(v27 + 108)))) == 0 {
				return 1
			}
		} else {
			nox_xxx_playerSetState_4FA020((*nox_object_t)(unsafe.Pointer(uintptr(v25))), 10)
			nox_xxx_playerManaSub_4EEBF0(int32(*(*uint32)(unsafe.Pointer(uintptr(source + 16)))), 1)
			if int32(nox_xxx_unitGetOldMana_4EEC80(int32(*(*uint32)(unsafe.Pointer(uintptr(source + 16)))))) == 0 {
				return 1
			}
		}
	}
	if (gameFrame() % (gameFPS() / 3)) == 0 {
		nox_xxx_aud_501960(78, (*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(source + 16)))))), 0, 0)
		nox_xxx_aud_501960(78, (*nox_object_t)(unsafe.Pointer(uintptr(*memmap.PtrInt32(0x5D4594, 2487844)))), 0, 0)
	}
	lightningSearchTime = float32(nox_xxx_gamedataGetFloat_419D40(internCStr("LightningSearchTime")))
	*(*uint32)(unsafe.Pointer(uintptr(source + 68))) = gameFrame() + uint32(nox_float2int(lightningSearchTime))
	return 0
}
func nox_xxx_lightningCanAttackCheck_52FF10(target int32, source int32) {
	var (
		owner              int32
		v3                 int32
		index              int32
		ptrTargetFromArray *uint8
		xDistance          float64
		yDistance          float64
		distance           float64
	)
	if *(*uint32)(unsafe.Pointer(uintptr(target + 8)))&0x20006 != 0 {
		owner = int32(nox_xxx_lightningOwner_5d4594_2487900)
		if nox_xxx_lightningOwner_5d4594_2487900 != 0 {
			if nox_xxx_unitIsEnemyTo_5330C0((*nox_object_t)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&nox_xxx_lightningOwner_5d4594_2487900))))), (*nox_object_t)(unsafe.Pointer(uintptr(target)))) == 0 {
				return
			}
			owner = int32(nox_xxx_lightningOwner_5d4594_2487900)
		}
		if (int32(*(*uint8)(unsafe.Pointer(uintptr(target + 8))))&2) == 0 || (func() bool {
			v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(target + 12))))
			return (v3 & 0x8000) == 0
		}()) {
			if (*(*uint32)(unsafe.Pointer(uintptr(target + 16)))&0x8020) == 0 && target != source && target != owner {
				index = 0
				if *(*int32)(unsafe.Pointer(&nox_xxx_lightningTargetArrayIndex_5d4594_2487904)) > 0 {
					ptrTargetFromArray = (*uint8)(memmap.PtrOff(0x5D4594, 2487844))
					for *(*uint32)(unsafe.Pointer(ptrTargetFromArray)) != uint32(target) {
						index++
						ptrTargetFromArray = (*uint8)(unsafe.Add(unsafe.Pointer(ptrTargetFromArray), 4))
						if index >= *(*int32)(unsafe.Pointer(&nox_xxx_lightningTargetArrayIndex_5d4594_2487904)) {
							goto LABEL_14
						}
					}
					return
				}
			LABEL_14:
				if nox_xxx_unitCanInteractWith_5370E0((*nox_object_t)(unsafe.Pointer(uintptr(source))), (*nox_object_t)(unsafe.Pointer(uintptr(target))), 0) != 0 {
					xDistance = float64(*(*float32)(unsafe.Pointer(uintptr(target + 56))) - *(*float32)(unsafe.Pointer(uintptr(source + 56))))
					yDistance = float64(*(*float32)(unsafe.Pointer(uintptr(target + 60))) - *(*float32)(unsafe.Pointer(uintptr(source + 60))))
					distance = yDistance*yDistance + xDistance*xDistance
					if distance < float64(*(*float32)(unsafe.Pointer(&nox_xxx_lightningClosestTargetDistance_5d4594_2487912))) {
						*(*float32)(unsafe.Pointer(&nox_xxx_lightningClosestTargetDistance_5d4594_2487912)) = float32(distance)
						nox_xxx_lightningTarget_5d4594_2487908 = uint32(target)
					}
				}
			}
		}
	}
}
func nox_xxx_lightningSpellTrapEffect_530020(a1 int32, a2 int32) {
	var (
		v2 int32
		v3 int32
		v4 float32
		v5 float32
		v6 int32
		v7 int32
		v8 float32
		v9 float4
	)
	if a1 != a2 {
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))))
		if v2&6 != 0 {
			if (*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))) & 0x8020) == 0 {
				if (v2&2) == 0 || (func() bool {
					v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12))))
					return (v3 & 0x8000) == 0
				}()) {
					if a2 == 0 || nox_xxx_unitIsEnemyTo_5330C0((*nox_object_t)(unsafe.Pointer(uintptr(a2))), (*nox_object_t)(unsafe.Pointer(uintptr(a1)))) != 0 {
						v4 = *(*float32)(unsafe.Pointer(uintptr(a1 + 56)))
						v9.field_4 = *mem_getFloatPtr(0x5D4594, 2487824)
						v9.field_0 = *mem_getFloatPtr(0x5D4594, 2487820)
						v5 = *(*float32)(unsafe.Pointer(uintptr(a1 + 60)))
						v9.field_8 = v4
						v9.field_C = v5
						if nox_xxx_mapTraceRay_535250(&v9, nil, nil, 9) != 0 {
							v8 = float32(nox_xxx_gamedataGetFloat_419D40(internCStr("LightningGlyphDamage")))
							v6 = nox_float2int(v8)
							(*(*func(int32, uint32, uint32, int32, int32))(unsafe.Pointer(uintptr(a1 + 716))))(a1, 0, 0, v6, 17)
							nox_xxx_netSendPointFx_522FF0(-127, (*float2)(unsafe.Pointer(uintptr(a1+56))))
							v7 = nox_xxx_spellGetAud44_424800(43, 0)
							nox_xxx_aud_501960(v7, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
						}
					}
				}
			}
		}
	}
}
func sub_530100(a1 *uint32) int8 {
	var (
		v1 int32
		v2 int32
		v3 int32
		v4 int32
		v5 int32
		v6 int32
	)
	v1 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*27)))
	if v1 != 0 {
		for {
			v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 116))))
			sub_4FE980(unsafe.Pointer(uintptr(v1)))
			v1 = v2
			if v2 == 0 {
				break
			}
		}
	}
	v3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*26)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*27)) = 0
	if v3 != 0 {
		for {
			v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 116))))
			sub_4FE980(unsafe.Pointer(uintptr(v3)))
			v3 = v4
			if v4 == 0 {
				break
			}
		}
	}
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*26)) = 0
	v5 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*18)))
	if v5 != 0 {
		v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 736))))
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(v6 + 96))))
		*((*uint8)(unsafe.Pointer(&v3))) = uint8(int8(v3 & 0xFB))
		*(*uint32)(unsafe.Pointer(uintptr(v6 + 96))) = uint32(v3)
	}
	return int8(v3)
}
func nox_xxx_spellTagCreature_530160(a1 *uint32) int32 {
	var (
		v1  int32
		v2  int32
		v3  int32
		v4  *uint32
		v5  int16
		v6  int32
		v7  int16
		v9  float32
		v10 [7]byte
	)
	v1 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*4)))
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 748))))
	if v1 == 0 {
		return 1
	}
	if *(*uint32)(unsafe.Pointer(uintptr(v1 + 16)))&0x8020 != 0 {
		return 1
	}
	if (int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 8)))) & 4) == 0 {
		return 1
	}
	v3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*12)))
	if v3 == 0 || *(*uint32)(unsafe.Pointer(uintptr(v3 + 16)))&0x8020 != 0 {
		return 1
	}
	v9 = float32(nox_xxx_gamedataGetFloat_419D40(internCStr("TagDurationPerLevel")))
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*17)) = gameFrame() + *(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*2))*uint32(nox_float2int(v9))
	nox_xxx_netMarkMinimapObject_417190(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 2064)))), (*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*12))))), 1)
	v4 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*12)))))
	v10[0] = 210
	v5 = int16(uint16(nox_xxx_netGetUnitCodeServ_578AC0((*nox_object_t)(unsafe.Pointer(v4)))))
	v6 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*12)))
	*(*uint16)(unsafe.Pointer(&v10[1])) = uint16(v5)
	v7 = int16(*(*uint16)(unsafe.Pointer(uintptr(v6 + 4))))
	v10[5] = 1
	*(*uint16)(unsafe.Pointer(&v10[3])) = uint16(v7)
	v10[6] = 1
	nox_xxx_netSendPacket0_4E5420(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 2064)))), unsafe.Pointer(&v10[0]), 7, 0, 1)
	return 0
}
func sub_530250(a1 int32) uint32 {
	var (
		v1     int32
		result uint32
	)
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 48))))
	if v1 != 0 {
		result = ((*(*uint32)(unsafe.Pointer(uintptr(v1 + 16))) & math.MaxUint8) >> 5) & 1
	} else {
		result = 1
	}
	return result
}
func sub_530270(a1 int32) int32 {
	var (
		result int32
		v2     int32
		v3     *uint32
		v4     int16
		v5     [7]byte
	)
	result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))))
	if result != 0 {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(result + 8))))&4 != 0 {
			v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(result + 748))))
			result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 48))))
			if result != 0 {
				if (int32(*(*uint8)(unsafe.Pointer(uintptr(result + 8)))) & 4) == 0 {
					nox_xxx_netUnmarkMinimapObj_417300(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 2064)))), (*nox_object_t)(unsafe.Pointer(uintptr(result))), 1)
				}
				v3 = *(**uint32)(unsafe.Pointer(uintptr(a1 + 48)))
				v5[0] = 210
				*(*uint16)(unsafe.Pointer(&v5[1])) = uint16(nox_xxx_netGetUnitCodeServ_578AC0((*nox_object_t)(unsafe.Pointer(v3))))
				v4 = int16(*(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 48))) + 4))))
				v5[5] = 2
				*(*uint16)(unsafe.Pointer(&v5[3])) = uint16(v4)
				v5[6] = 1
				result = nox_xxx_netSendPacket0_4E5420(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 2064)))), unsafe.Pointer(&v5[0]), 7, 0, 1)
			}
		}
	}
	return result
}
func nox_xxx_spellBlink2_530310(a1 *uint32) int32 {
	var (
		result int32
		v2     float32
	)
	if !nox_common_gameFlags_check_40A5C0(4096) || (func() bool {
		result = 1
		return *(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*5)) != 1
	}()) {
		if nox_common_gameFlags_check_40A5C0(2048) {
			v2 = float32(nox_xxx_gamedataGetFloatTable_419D70(internCStr("TeleportDelay"), int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*2))-1)))
			*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*17)) = gameFrame() + uint32(nox_float2int(v2))
			result = 0
		} else {
			result = 0
			*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*17)) = gameFrame() + 1
		}
	}
	return result
}
func nox_xxx_spellBlink1_530380(a1 *int32) int32 {
	var (
		v1  int32
		v3  int32
		v4  *float2
		v5  int32
		v6  bool
		v7  int32
		v8  int32
		v9  int32
		v10 int32
		v11 int32
		v12 int32
		v13 float2
	)
	v1 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))
	if v1 == 0 {
		return 1
	}
	if *(*uint32)(unsafe.Pointer(uintptr(v1 + 16)))&0x8020 != 0 {
		return 1
	}
	if uint32(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*17))-1) != gameFrame() {
		return 0
	}
	if nox_xxx_testUnitBuffs_4FF350((*nox_object_t)(unsafe.Pointer(uintptr(v1))), 14) != 0 {
		nox_xxx_aud_501960(231, (*nox_object_t)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))))), 0, 0)
		return 1
	}
	if nox_common_gameFlags_check_40A5C0(4096) && (func() int32 {
		v3 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))
		return int32(*(*uint8)(unsafe.Pointer(uintptr(v3 + 8)))) & 4
	}()) != 0 {
		v4 = *(**float2)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 748))) + 308)))
		if v4 != nil {
			sub_4ED970(60.0, (*float2)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(float2{})*7)), &v13)
		} else {
			nox_xxx_mapFindPlayerStart_4F7AB0(&v13, (*nox_object_t)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))))))
		}
	} else if nox_xxx_waypoint_579F00((*uint32)(unsafe.Pointer(&v13)), *(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))) == 0 {
		nox_xxx_mapFindPlayerStart_4F7AB0(&v13, (*nox_object_t)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))))))
	}
	nox_xxx_spellTeleportCreateWake_530560(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12)), (*int32)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))+56))), (*uint32)(unsafe.Pointer(&v13)))
	nox_xxx_netSendPointFx_522FF0(-119, (*float2)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))+56))))
	v9 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))
	v5 = nox_xxx_spellGetAud44_424800(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*1)), 0)
	nox_xxx_aud_501960(v5, (*nox_object_t)(unsafe.Pointer(uintptr(v9))), 0, 0)
	if nox_xxx_testUnitBuffs_4FF350((*nox_object_t)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))))), 0) == 0 {
		nox_xxx_netSendPointFx_522FF0(-119, (*float2)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))+56))))
		nox_xxx_netSendPointFx_522FF0(-119, &v13)
	}
	nox_xxx_teleportToMB_4E7190((*uint8)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))))), &v13.field_0)
	v6 = nox_xxx_testUnitBuffs_4FF350((*nox_object_t)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))))), 0) == 0
	v7 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))
	if v6 {
		nox_xxx_netSendPointFx_522FF0(-119, (*float2)(unsafe.Pointer(uintptr(v7+56))))
		v12 = 0
		v11 = 0
		v10 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))
	} else {
		if (int32(*(*uint8)(unsafe.Pointer(uintptr(v7 + 8)))) & 4) == 0 {
			sub_4E7540((*nox_object_t)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*4))))), (*nox_object_t)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))))))
			return 1
		}
		v12 = int32(*(*uint32)(unsafe.Pointer(uintptr(v7 + 36))))
		v11 = 2
		v10 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))
	}
	v8 = nox_xxx_spellGetAud44_424800(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*1)), 0)
	nox_xxx_aud_501960(v8, (*nox_object_t)(unsafe.Pointer(uintptr(v10))), v11, v12)
	sub_4E7540((*nox_object_t)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*4))))), (*nox_object_t)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))))))
	return 1
}
func nox_xxx_spellTeleportCreateWake_530560(a1 int32, a2 *int32, a3 *uint32) *uint32 {
	var (
		v3     int32
		result *uint32
		v5     *uint32
		v6     *uint32
	)
	v3 = int32(*memmap.PtrUint32(0x5D4594, 2487916))
	if *memmap.PtrUint32(0x5D4594, 2487916) == 0 {
		v3 = nox_xxx_getNameId_4E3AA0(internCStr("TeleportWake"))
		*memmap.PtrUint32(0x5D4594, 2487916) = uint32(v3)
	}
	result = (*uint32)(unsafe.Pointer(nox_xxx_newObjectWithTypeInd_4E3450(v3)))
	v5 = result
	if result != nil {
		v6 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(result), 4*175)))))
		*v6 = *a3
		*(*uint32)(unsafe.Add(unsafe.Pointer(v6), 4*1)) = *(*uint32)(unsafe.Add(unsafe.Pointer(a3), 4*1))
		nox_xxx_createAt_4DAA50((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v5)))))), (*nox_object_t)(unsafe.Pointer(uintptr(a1))), *(*float32)(unsafe.Pointer(a2)), *((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(a2))), unsafe.Sizeof(float32(0))*1))))
		result = (*uint32)(unsafe.Pointer(uintptr(gameFrame() + gameFPS())))
		*(*uint32)(unsafe.Add(unsafe.Pointer(v5), 4*34)) = gameFrame() + gameFPS()
	}
	return result
}
func sub_5305D0(a1 *uint32) int32 {
	var (
		v1 int32
		v2 int32
		v4 float32
	)
	v1 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*3)))
	if int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 8))))&4 != 0 {
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 748))))
		if v2 == 0 || *(*uint32)(unsafe.Pointer(uintptr(uint32(v2) + *(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*1))*4 - 372))) == 0 {
			return 1
		}
	}
	if nox_common_gameFlags_check_40A5C0(2048) {
		v4 = float32(nox_xxx_gamedataGetFloatTable_419D70(internCStr("TeleportDelay"), int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*2))-1)))
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*17)) = gameFrame() + uint32(nox_float2int(v4))
	} else {
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*17)) = gameFrame() + 1
	}
	return 0
}
func sub_530650(a1 *int32) int32 {
	var (
		v1  int32
		v2  int32
		v4  int32
		v5  int32
		v6  int32
		v7  int32
		v8  int32
		v9  int32
		v10 int32
		v11 int32
		v12 int8
		v13 int32
		v14 int32
		v15 int32
		v16 int32
		v17 int32
		v18 float2
	)
	v1 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*3))
	if v1 == 0 {
		return 1
	}
	if int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 16))))&0x20 != 0 {
		return 1
	}
	v2 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))
	if v2 == 0 || *(*uint32)(unsafe.Pointer(uintptr(v2 + 16)))&0x8020 != 0 || v2 != v1 && *(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*5)) == 0 {
		return 1
	}
	if uint32(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*17))-1) != gameFrame() {
		return 0
	}
	if nox_xxx_testUnitBuffs_4FF350((*nox_object_t)(unsafe.Pointer(uintptr(v2))), 14) != 0 {
		nox_xxx_aud_501960(231, (*nox_object_t)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))))), 0, 0)
		return 1
	}
	v4 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*3))
	if int32(*(*uint8)(unsafe.Pointer(uintptr(v4 + 8))))&4 != 0 {
		v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 748))))
		v6 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*1)) - 122
		v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + v6*4 + 116))))
		if v7 == 0 {
			return 1
		}
		v8 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))
		v18 = *(*float2)(unsafe.Pointer(uintptr(v8 + 56)))
		nox_xxx_spellTeleportCreateWake_530560(v8, (*int32)(unsafe.Pointer(uintptr(v8+56))), (*uint32)(unsafe.Pointer(uintptr(v7+56))))
		v14 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))
		v9 = nox_xxx_spellGetAud44_424800(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*1)), 1)
		nox_xxx_aud_501960(v9, (*nox_object_t)(unsafe.Pointer(uintptr(v14))), 0, 0)
		nox_xxx_teleportToMB_4E7190((*uint8)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))))), (*float32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v5 + v6*4 + 116)))+56))))
		if nox_xxx_testUnitBuffs_4FF350((*nox_object_t)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))))), 0) != 0 {
			v10 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))
			if (int32(*(*uint8)(unsafe.Pointer(uintptr(v10 + 8)))) & 4) == 0 {
				goto LABEL_18
			}
			v17 = int32(*(*uint32)(unsafe.Pointer(uintptr(v10 + 36))))
			v16 = 2
			v15 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))
		} else {
			nox_xxx_netSendPointFx_522FF0(-119, &v18)
			nox_xxx_netSendPointFx_522FF0(-119, (*float2)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))+56))))
			v17 = 0
			v16 = 0
			v15 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))
		}
		v11 = nox_xxx_spellGetAud44_424800(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*1)), 1)
		nox_xxx_aud_501960(v11, (*nox_object_t)(unsafe.Pointer(uintptr(v15))), v16, v17)
	LABEL_18:
		*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v5 + v6*4 + 116))) + 136))) = gameFrame()
		v12 = int8(int32(*(*uint8)(unsafe.Pointer(uintptr(v6 + v5 + 156)))) - 1)
		*(*uint8)(unsafe.Pointer(uintptr(v6 + v5 + 156))) = uint8(v12)
		if int32(v12) == 0 {
			nox_xxx_netSendPointFx_522FF0(-127, (*float2)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v5 + v6*4 + 116)))+56))))
			nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v5 + v6*4 + 116)))))))
			*(*uint32)(unsafe.Pointer(uintptr(v5 + v6*4 + 116))) = 0
		}
	}
	v13 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*4))
	if v13 == 0 || int32(*(*uint8)(unsafe.Pointer(uintptr(v13 + 16))))&0x20 != 0 {
		return 1
	}
	sub_4E7540((*nox_object_t)(unsafe.Pointer(uintptr(v13))), (*nox_object_t)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))))))
	return 1
}
func nox_xxx_castTele_530820(a1 int32) int32 {
	var (
		result int32
		v2     float32
	)
	if nox_common_gameFlags_check_40A5C0(2048) {
		v2 = float32(nox_xxx_gamedataGetFloatTable_419D70(internCStr("TeleportDelay"), int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))-1)))
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 68))) = gameFrame() + uint32(nox_float2int(v2))
		result = 0
	} else {
		result = 0
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 68))) = gameFrame() + 1
	}
	return result
}
func sub_530880(a1 *int32) int32 {
	var (
		v1 int32
		v2 int32
		v4 int32
		v5 int32
		v6 int32
		v7 int32
		v8 *uint32
		v9 int32
	)
	_ = v9
	var v10 int32
	var v11 int32
	var v12 int8
	var v13 int32
	var v14 float2
	v1 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*4))
	if v1 == 0 {
		return 1
	}
	if *(*uint32)(unsafe.Pointer(uintptr(v1 + 16)))&0x8020 != 0 {
		return 1
	}
	v2 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))
	if v2 == 0 || *(*uint32)(unsafe.Pointer(uintptr(v2 + 16)))&0x8020 != 0 {
		return 1
	}
	if uint32(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*17))-1) != gameFrame() {
		return 0
	}
	if nox_xxx_testUnitBuffs_4FF350((*nox_object_t)(unsafe.Pointer(uintptr(v2))), 14) != 0 {
		nox_xxx_aud_501960(231, (*nox_object_t)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))))), 0, 0)
		return 1
	}
	v4 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*4))
	if int32(*(*uint8)(unsafe.Pointer(uintptr(v4 + 8))))&4 != 0 {
		v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 748))))
		v6 = 0
		v7 = 4
		v8 = (*uint32)(unsafe.Pointer(uintptr(v5 + 116)))
		for {
			if *v8 != 0 {
				v6++
			}
			v8 = (*uint32)(unsafe.Add(unsafe.Pointer(v8), 4*1))
			v7--
			if v7 == 0 {
				break
			}
		}
		if v6 == 0 {
			return 1
		}
		v9 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))
		v14 = *(*float2)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12)) + 56)))
		for {
			v10 = nox_common_randomInt_415FA0(0, 3)
			if *(*uint32)(unsafe.Pointer(uintptr(v5 + v10*4 + 116))) != 0 {
				break
			}
		}
		nox_xxx_spellTeleportCreateWake_530560(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12)), (*int32)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))+56))), (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v5 + v10*4 + 116)))+56))))
		nox_xxx_teleportToMB_4E7190((*uint8)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))))), (*float32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v5 + v10*4 + 116)))+56))))
		if nox_xxx_testUnitBuffs_4FF350((*nox_object_t)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))))), 0) == 0 {
			nox_xxx_netSendPointFx_522FF0(-119, &v14)
			nox_xxx_netSendPointFx_522FF0(-119, (*float2)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))+56))))
		}
		v13 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))
		v11 = nox_xxx_spellGetAud44_424800(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*1)), 0)
		nox_xxx_aud_501960(v11, (*nox_object_t)(unsafe.Pointer(uintptr(v13))), 0, 0)
		v12 = int8(int32(*(*uint8)(unsafe.Pointer(uintptr(v10 + v5 + 156)))) - 1)
		*(*uint8)(unsafe.Pointer(uintptr(v10 + v5 + 156))) = uint8(v12)
		if int32(v12) == 0 {
			nox_xxx_netSendPointFx_522FF0(-127, (*float2)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v5 + v10*4 + 116)))+56))))
			nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v5 + v10*4 + 116)))))))
			*(*uint32)(unsafe.Pointer(uintptr(v5 + v10*4 + 116))) = 0
		}
	}
	sub_4E7540((*nox_object_t)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*4))))), (*nox_object_t)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))))))
	return 1
}
func nox_xxx_castTTT_530B70(a1 *int32) int32 {
	var (
		v1  int32
		v2  int32
		v4  int32
		v5  int32
		v6  int32
		v7  int32
		v8  int32
		v9  int32
		v10 int32
	)
	v1 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))
	if v1 == 0 {
		return 1
	}
	v2 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*4))
	if v2 == 0 || *(*uint32)(unsafe.Pointer(uintptr(v1 + 16)))&0x8020 != 0 || *(*uint32)(unsafe.Pointer(uintptr(v2 + 16)))&0x8020 != 0 {
		return 1
	}
	if uint32(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*17))-1) != gameFrame() {
		return 0
	}
	if nox_xxx_testUnitBuffs_4FF350((*nox_object_t)(unsafe.Pointer(uintptr(v1))), 14) != 0 {
		nox_xxx_aud_501960(231, (*nox_object_t)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))))), 0, 0)
		return 1
	}
	nox_xxx_netSendPointFx_522FF0(-119, (*float2)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))+56))))
	v7 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))
	v4 = nox_xxx_spellGetAud44_424800(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*1)), 0)
	nox_xxx_aud_501960(v4, (*nox_object_t)(unsafe.Pointer(uintptr(v7))), 0, 0)
	nox_xxx_spellTeleportCreateWake_530560(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12)), (*int32)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))+56))), (*uint32)(unsafe.Pointer((*int32)(unsafe.Add(unsafe.Pointer(a1), 4*13)))))
	nox_xxx_teleportToMB_4E7190((*uint8)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))))), (*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(a1))), unsafe.Sizeof(float32(0))*13)))
	if nox_xxx_testUnitBuffs_4FF350((*nox_object_t)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))))), 0) != 0 {
		v6 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))
		if (int32(*(*uint8)(unsafe.Pointer(uintptr(v6 + 8)))) & 4) == 0 {
			sub_4E7540((*nox_object_t)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*4))))), (*nox_object_t)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))))))
			return 1
		}
		v10 = int32(*(*uint32)(unsafe.Pointer(uintptr(v6 + 36))))
		v9 = 2
		v8 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))
		v5 = nox_xxx_spellGetAud44_424800(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*1)), 0)
	} else {
		nox_xxx_netSendPointFx_522FF0(-119, (*float2)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))+56))))
		v10 = 0
		v9 = 0
		v8 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))
		v5 = nox_xxx_spellGetAud44_424800(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*1)), 0)
	}
	nox_xxx_aud_501960(v5, (*nox_object_t)(unsafe.Pointer(uintptr(v8))), v9, v10)
	sub_4E7540((*nox_object_t)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*4))))), (*nox_object_t)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))))))
	return 1
}
func sub_530CA0(a1 int32) int32 {
	var (
		v1     int32
		v2     int32
		v3     int32
		result int32
		v5     float32
	)
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))))
	if v1 == 0 {
		return 1
	}
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 88))))&0x20 != 0 {
		return 1
	}
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 48))))
	if v2 == 0 {
		return 1
	}
	if *(*uint32)(unsafe.Pointer(uintptr(v2 + 16)))&0x8020 != 0 {
		return 1
	}
	if v1 == v2 {
		return 1
	}
	if int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 8))))&2 != 0 {
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 12))))
		if v3&0x4000 != 0 {
			return 1
		}
	}
	if nox_common_gameFlags_check_40A5C0(2048) {
		v5 = float32(nox_xxx_gamedataGetFloatTable_419D70(internCStr("TeleportDelay"), int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))-1)))
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 68))) = gameFrame() + uint32(nox_float2int(v5))
		result = 0
	} else {
		result = 0
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 68))) = gameFrame() + 1
	}
	return result
}
func sub_530D30(a1 *int32) int32 {
	var (
		v1  *int32
		v2  int32
		v3  int32
		v5  int32
		v6  int32
		v7  *float2
		v8  int32
		v9  int32
		v10 int32
		v11 int32
		v12 int32
		v13 *float32
		v14 float4
		v18 int32
	)
	v1 = a1
	v2 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))
	if v2 == 0 {
		return 1
	}
	v3 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*4))
	if v3 == 0 || *(*uint32)(unsafe.Pointer(uintptr(v2 + 16)))&0x8020 != 0 || *(*uint32)(unsafe.Pointer(uintptr(v3 + 16)))&0x8020 != 0 {
		return 1
	}
	if uint32(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*17))-1) != gameFrame() {
		return 0
	}
	if nox_xxx_testUnitBuffs_4FF350((*nox_object_t)(unsafe.Pointer(uintptr(v2))), 14) != 0 || nox_xxx_testUnitBuffs_4FF350((*nox_object_t)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*4))))), 14) != 0 {
		nox_xxx_aud_501960(231, (*nox_object_t)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))))), 0, 0)
		nox_xxx_aud_501960(231, (*nox_object_t)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*4))))), 0, 0)
		return 1
	}
	if *(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*5)) != 0 {
		goto LABEL_23
	}
	if nox_xxx_unitCanInteractWith_5370E0((*nox_object_t)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*4))))), (*nox_object_t)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*12))))), 0) == 0 {
		nox_xxx_netPriMsgToPlayer_4DA2C0((*nox_object_t)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*4))))), internCStr("ExecDur.c:NeedClearLOSForSwap"), 0)
		return 1
	}
	v5 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*4))
	if (int32(*(*uint8)(unsafe.Pointer(uintptr(v5 + 8))))&4) != 0 && (func() int32 {
		v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 748))))
		v14.field_0 = float32(float64(*(*float32)(unsafe.Pointer(uintptr(v5 + 56)))) - float64(*(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v6 + 276))) + 10)))))
		v14.field_4 = float32(float64(*(*float32)(unsafe.Pointer(uintptr(v5 + 60)))) - float64(*(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v6 + 276))) + 12)))))
		v14.field_8 = float32(float64(*(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v6 + 276))) + 10)))) + float64(*(*float32)(unsafe.Pointer(uintptr(v5 + 56)))))
		v18 = int32(*(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v6 + 276))) + 12))))
		v7 = (*float2)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(v1), 4*12)) + 56)))
		v14.field_C = float32(float64(v18) + float64(*(*float32)(unsafe.Pointer(uintptr(v5 + 60)))))
		return sub_428220(v7, &v14)
	}()) == 0 {
		nox_xxx_netPriMsgToPlayer_4DA2C0((*nox_object_t)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(v1), 4*4))))), internCStr("ExecDur.c:NeedClearLOSForSwap"), 0)
		return 1
	}
LABEL_23:
	v8 = *(*int32)(unsafe.Add(unsafe.Pointer(v1), 4*12))
	v14.field_0 = *(*float32)(unsafe.Pointer(uintptr(v8 + 56)))
	v13 = (*float32)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(v1), 4*4)) + 56)))
	v14.field_4 = *(*float32)(unsafe.Pointer(uintptr(v8 + 60)))
	nox_xxx_teleportToMB_4E7190((*uint8)(unsafe.Pointer(uintptr(v8))), v13)
	nox_xxx_teleportToMB_4E7190((*uint8)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(v1), 4*4))))), &v14.field_0)
	if nox_xxx_testUnitBuffs_4FF350((*nox_object_t)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(v1), 4*12))))), 0) == 0 && nox_xxx_testUnitBuffs_4FF350((*nox_object_t)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(v1), 4*4))))), 0) == 0 {
		nox_xxx_netSendPointFx_522FF0(-119, (*float2)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(v1), 4*4))+56))))
		nox_xxx_netSendPointFx_522FF0(-119, (*float2)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(v1), 4*12))+56))))
	}
	v11 = *(*int32)(unsafe.Add(unsafe.Pointer(v1), 4*12))
	v9 = nox_xxx_spellGetAud44_424800(*(*int32)(unsafe.Add(unsafe.Pointer(v1), 4*1)), 0)
	nox_xxx_aud_501960(v9, (*nox_object_t)(unsafe.Pointer(uintptr(v11))), 0, 0)
	v12 = *(*int32)(unsafe.Add(unsafe.Pointer(v1), 4*4))
	v10 = nox_xxx_spellGetAud44_424800(*(*int32)(unsafe.Add(unsafe.Pointer(v1), 4*1)), 0)
	nox_xxx_aud_501960(v10, (*nox_object_t)(unsafe.Pointer(uintptr(v12))), 0, 0)
	sub_4E7540((*nox_object_t)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(v1), 4*4))))), (*nox_object_t)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(v1), 4*12))))))
	return 1
}
func nox_xxx_manaBomb_530F90(a1 *uint32) int32 {
	var (
		v1 int32
		v2 *uint32
		v4 float32
		v5 float32
	)
	if *memmap.PtrUint32(0x5D4594, 2487920) == 0 {
		*memmap.PtrUint32(0x5D4594, 2487920) = uint32(nox_xxx_getNameId_4E3AA0(internCStr("ManaBombCharge")))
	}
	v4 = float32(nox_xxx_gamedataGetFloatTable_419D70(internCStr("ManaBombInitPower"), int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*2))-1)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*18)) = uint32(nox_float2int(v4))
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*19)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*21)) = 0
	v1 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*4)))
	if v1 != 0 && (int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 8))))&4) == 0 || *(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*5)) != 0 {
		v5 = float32(nox_xxx_gamedataGetFloat_419D40(internCStr("ManaBombGlyphDuration")))
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*17)) = gameFrame() + uint32(nox_float2int(v5))
	} else {
		nox_xxx_buffApplyTo_4FF380((*nox_object_t)(unsafe.Pointer(uintptr(v1))), 5, int16(int32(uint16(gameFPS()))*10), 5)
		nox_xxx_buffApplyTo_4FF380((*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*4))))), 14, int16(int32(uint16(gameFPS()))*10), 5)
		nox_xxx_buffApplyTo_4FF380((*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*4))))), 29, int16(int32(uint16(gameFPS()))*10), 5)
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*20)) = *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*4)) + 120)))
		*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*4)) + 120))) = 1203982323
		*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*4)) + 84))) = 0
		*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*4)) + 80))) = 0
		*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*4)) + 92))) = 0
		*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*4)) + 88))) = 0
		*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*4)) + 100))) = 0
		*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*4)) + 96))) = 0
	}
	v2 = (*uint32)(unsafe.Pointer(nox_xxx_newObjectWithTypeInd_4E3450(*memmap.PtrInt32(0x5D4594, 2487920))))
	if v2 != nil {
		nox_xxx_createAt_4DAA50((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v2)))))), nil, *((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(a1))), unsafe.Sizeof(float32(0))*7))), *((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(a1))), unsafe.Sizeof(float32(0))*8))))
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*19)) = uint32(uintptr(unsafe.Pointer(v2)))
	}
	return 0
}
func nox_xxx_manaBombBoom_5310C0(a1 *int32) int32 {
	var (
		v1  int32
		v2  int32
		v4  int32
		v5  bool
		v6  int32
		v7  float32
		v8  int32
		v9  int32
		v10 float32
		v11 float32
		v12 int32
		v13 float32
		v14 float32
		v15 float2
	)
	v1 = 0
	v2 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*4))
	if v2 == 0 && *(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*5)) == 0 {
		return 1
	}
	if uint32(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*17))-1) == gameFrame() {
		v1 = 1
	}
	if v2 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 8))))&4 != 0 && int32(nox_xxx_unitGetOldMana_4EEC80(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*4)))) == 0 {
		v1 = 1
	}
	if *(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*19)) != 0 {
		if *(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*5)) == 0 && (func() int32 {
			v4 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*4))
			return v4
		}()) != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(v4 + 8))))&4 != 0 {
			if int32(uint16(nox_xxx_unitGetOldMana_4EEC80(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*4))))) >= 0xF {
				goto LABEL_18
			}
		} else if (uint32(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*17))) - gameFrame()) >= 0xA {
			goto LABEL_18
		}
		nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*19))))))
		*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*19)) = 0
	}
LABEL_18:
	v5 = v1 == 0
	v6 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*18))
	if !v5 {
		if *(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*5)) != 0 {
			v7 = *((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(a1))), unsafe.Sizeof(float32(0))*8)))
			v15.field_0 = *((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(a1))), unsafe.Sizeof(float32(0))*7)))
		} else {
			v8 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*4))
			if v8 == 0 {
				return 1
			}
			v15.field_0 = *(*float32)(unsafe.Pointer(uintptr(v8 + 56)))
			v7 = *(*float32)(unsafe.Pointer(uintptr(v8 + 60)))
		}
		v15.field_4 = v7
		nox_xxx_gameSetWallsDamage_4E25A0(1)
		v12 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*4))
		v11 = float32(nox_xxx_gamedataGetFloat_419D40(internCStr("ManaBombInRadius")))
		v10 = float32(nox_xxx_gamedataGetFloat_419D40(internCStr("ManaBombOutRadius")))
		nox_xxx_mapDamageUnitsAround_4E25B0(&v15.field_0, v10, v11, v6, 15, (*nox_object_t)(unsafe.Pointer(uintptr(v12))), nil)
		v13 = float32(nox_xxx_gamedataGetFloat_419D40(internCStr("ManaBombShakeMag")))
		v9 = nox_float2int(v13)
		nox_xxx_earthquakeSend_4D9110(&v15.field_0, v9)
		nox_xxx_netSendPointFx_522FF0(-127, &v15)
		nox_xxx_netSendPointFx_522FF0(-102, &v15)
		nox_xxx_aud_501960(81, (*nox_object_t)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*4))))), 0, 0)
		*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*21)) = 1
		return 1
	}
	v14 = float32(nox_xxx_gamedataGetFloatTable_419D70(internCStr("ManaBombDeltaPower"), *(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*2))-1))
	*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*18)) = v6 + nox_float2int(v14)
	if *(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*5)) == 0 {
		if sub_4E7BC0(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*4))) != 0 {
			nox_xxx_playerManaSub_4EEBF0(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*4)), *(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*2)))
		}
	}
	return 0
}
func sub_531290(a1 int32) int32 {
	var (
		v1     int32
		result int32
	)
	if *(*uint32)(unsafe.Pointer(uintptr(a1 + 76))) != 0 {
		nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 76)))))))
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 76))) = 0
	}
	if *(*uint32)(unsafe.Pointer(uintptr(a1 + 20))) == 0 {
		v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))))
		if v1 != 0 {
			if int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 8))))&4 != 0 {
				nox_xxx_spellBuffOff_4FF5B0((*nox_object_t)(unsafe.Pointer(uintptr(v1))), 5)
				nox_xxx_spellBuffOff_4FF5B0((*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16)))))), 14)
				nox_xxx_spellBuffOff_4FF5B0((*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16)))))), 29)
				*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))) + 120))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 80)))
			}
		}
	}
	result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 84))))
	if result == 0 {
		result = nox_xxx_netSendPointFx_522FF0(-93, (*float2)(unsafe.Pointer(uintptr(a1+28))))
	}
	return result
}
func nox_xxx_spellTurnUndeadCreate_531310(a1 *uint32) int32 {
	var (
		v1 int32
		v2 int32
		v3 *uint32
		v4 *uint32
		v5 int32
		v6 float64
		v8 float32
		v9 float2
	)
	v8 = float32(nox_xxx_gamedataGetFloatTable_419D70(internCStr("TurnUndeadKillPoints"), int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*2))-1)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*18)) = uint32(nox_float2int(v8))
	if *memmap.PtrUint32(0x5D4594, 2487924) == 0 {
		*memmap.PtrUint32(0x5D4594, 2487924) = uint32(nox_xxx_getNameId_4E3AA0(internCStr("UndeadKiller")))
	}
	if *(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*5)) != 0 {
		v1 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*6)))
	} else {
		v1 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*4)))
	}
	v2 = 0
	v9 = *(*float2)(unsafe.Pointer(uintptr(v1 + 56)))
	for {
		v3 = (*uint32)(unsafe.Pointer(nox_xxx_newObjectWithTypeInd_4E3450(*memmap.PtrInt32(0x5D4594, 2487924))))
		v4 = v3
		if v3 != nil {
			*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), 4*175))))) = uint32(uintptr(unsafe.Pointer(a1)))
			nox_xxx_createAt_4DAA50((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v3)))))), (*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*4))))), v9.field_0, v9.field_4)
			*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v4))), unsafe.Sizeof(uint16(0))*63))) = uint16(int16(v2))
			v5 = int32(int16(v2)) * 8
			*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v4))), unsafe.Sizeof(uint16(0))*62))) = uint16(int16(v2))
			*((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v4))), unsafe.Sizeof(float32(0))*20))) = float32(float64(*mem_getFloatPtr(0x587000, uint32(v5)+194136)) * 4.0)
			v6 = float64(*mem_getFloatPtr(0x587000, uint32(v5)+194140)) * 4.0
			*(*uint32)(unsafe.Add(unsafe.Pointer(v4), 4*28)) = 0
			*((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v4))), unsafe.Sizeof(float32(0))*21))) = float32(v6)
		}
		v2 += 6
		if v2 >= 256 {
			break
		}
	}
	nox_xxx_netSendPointFx_522FF0(-96, &v9)
	return 0
}
func nox_xxx_spellTurnUndeadUpdate_531410() int32 {
	return 0
}
func nox_xxx_spellTurnUndeadDelete_531420(a1 int32) int32 {
	var (
		result int32
		i      int32
	)
	if *memmap.PtrUint32(0x5D4594, 2487928) == 0 {
		*memmap.PtrUint32(0x5D4594, 2487928) = uint32(nox_xxx_getNameId_4E3AA0(internCStr("UndeadKiller")))
	}
	result = int32(uintptr(unsafe.Pointer(nox_server_getFirstObject_4DA790())))
	for i = result; result != 0; i = result {
		if uint32(*(*uint16)(unsafe.Pointer(uintptr(i + 4)))) == *memmap.PtrUint32(0x5D4594, 2487928) && **(**uint32)(unsafe.Pointer(uintptr(i + 700))) == uint32(a1) {
			nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(i))))
		}
		result = int32(uintptr(unsafe.Pointer(nox_server_getNextObject_4DA7A0((*nox_object_t)(unsafe.Pointer(uintptr(i)))))))
	}
	return result
}
func sub_531490(a1 *uint32) int32 {
	var (
		v1     int32
		v2     int32
		result int32
	)
	v1 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*12)))
	v2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*2)) * 20 * gameFPS())
	if v1 == 0 {
		return 1
	}
	if (int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 8)))) & 4) == 0 {
		return 1
	}
	sub_4FF310((*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*12))))))
	nox_xxx_buffApplyTo_4FF380((*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*12))))), 27, int16(v2), int8(uint8(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*2)))))
	result = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*17)) = uint32(v2) + gameFrame()
	return result
}
func sub_5314F0(a1 int32) int32 {
	var (
		v1     int32
		result int32
		v3     int32
	)
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 48))))
	if v1 == 0 {
		return 1
	}
	if nox_xxx_testUnitBuffs_4FF350((*nox_object_t)(unsafe.Pointer(uintptr(v1))), 8) != 0 {
		return 1
	}
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 48))))
	if v3 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(v3 + 8))))&2 != 0 && sub_4FEA70(v3, (*float2)(unsafe.Pointer(uintptr(a1+28)))) != 0 {
		result = 1
	} else {
		result = bool2int32((*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 48))) + 16))) & 0x8020) != 0)
	}
	return result
}
func sub_531560(a1 int32) int32 {
	var result int32
	result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 48))))
	if result != 0 {
		result = nox_xxx_spellBuffOff_4FF5B0((*nox_object_t)(unsafe.Pointer(uintptr(result))), 27)
	}
	return result
}
func nox_xxx_plasmaSmth_531580(a1 int32) int32 {
	var (
		v1 int32
		v2 int32
	)
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 72))) = 0
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 76))) = 0
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 48))) = 0
	nox_xxx_netSendPointFx_522FF0(-125, (*float2)(unsafe.Pointer(uintptr(a1+28))))
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))))
	if v1 == 0 || (int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 8))))&4) == 0 {
		return 0
	}
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 748))) + 104))))
	if v2 != 0 && *(*uint32)(unsafe.Pointer(uintptr(v2 + 12)))&0x4000000 != 0 {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 736))) + 96))))&4 != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(a1 + 72))) = uint32(v2)
			*(*uint8)(unsafe.Pointer(uintptr(a1 + 88))) |= 2
		}
		return 0
	}
	return 1
}
func nox_xxx_plasmaShot_531600(a1 int32) int32 {
	var (
		v1  int32
		v2  int32
		v3  int32
		v4  int32
		v5  int32
		v6  int32
		v7  int32
		v8  int32
		v9  float64
		v10 int32
		v11 int32
		v12 int32
		v13 int32
		v14 uint8
		v15 uint8
		v16 int32
		v17 int32
		v19 *uint32
		v20 float32
		v21 float32
	)
	if *memmap.PtrUint32(0x5D4594, 2487936) == 0 {
		*memmap.PtrUint32(0x5D4594, 2487936) = uint32(nox_xxx_getNameId_4E3AA0(internCStr("Hecubah")))
		*memmap.PtrUint32(0x5D4594, 2487940) = uint32(nox_xxx_getNameId_4E3AA0(internCStr("HecubahWithOrb")))
	}
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))))
	if v1 == 0 || int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 88))))&0x20 != 0 || int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 8))))&2 != 0 && sub_4FEA70(v1, (*float2)(unsafe.Pointer(uintptr(a1+28)))) != 0 {
		return 1
	}
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 48))))
	if v2 == 0 {
		goto LABEL_16
	}
	if *(*uint32)(unsafe.Pointer(uintptr(v2 + 16)))&0x8020 != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 48))) = 0
		goto LABEL_16
	}
	if nox_server_testTwoPointsAndDirection_4E6E50((*float2)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16)))+56))), int32(*(*int16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))) + 124)))), (*float2)(unsafe.Pointer(uintptr(v2+56))))&2 != 0 {
		if *(*uint32)(unsafe.Pointer(uintptr(a1 + 76))) == 0 {
			*(*uint32)(unsafe.Pointer(uintptr(a1 + 76))) = gameFPS() * 3
		}
	} else {
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 76))) = 0
	}
	if nox_xxx_calcDistance_4E6C00((*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 48)))))), (*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))))))) > 400.0 || nox_xxx_unitCanInteractWith_5370E0((*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16)))))), (*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 48)))))), 0) == 0 {
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 48))) = 0
		goto LABEL_16
	} else {
		goto LABEL_A
	}
LABEL_16:
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))) + 748))) + 288))))
	dword_5d4594_2487932 = 0
	if v3 == 0 {
		*memmap.PtrUint32(0x587000, 260404) = 1209810944
		nox_xxx_unitsGetInCircle_517F90((*float2)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16)))+56))), 400.0, unsafe.Pointer(funAddr(sub_531920)), unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))))))
	} else {
		if nox_xxx_unitIsEnemyTo_5330C0((*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16)))))), (*nox_object_t)(unsafe.Pointer(uintptr(v3)))) != 0 && nox_xxx_calcDistance_4E6C00((*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16)))))), (*nox_object_t)(unsafe.Pointer(uintptr(v3)))) <= 400.0 {
			dword_5d4594_2487932 = uint32(v3)
		}
		if dword_5d4594_2487932 == 0 {
			*memmap.PtrUint32(0x587000, 260404) = 1209810944
			nox_xxx_unitsGetInCircle_517F90((*float2)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16)))+56))), 400.0, unsafe.Pointer(funAddr(sub_531920)), unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))))))
		}
	}
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 48))) = dword_5d4594_2487932
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 76))) = 0
LABEL_A:
	v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 76))))
	if v4 != 0 {
		v5 = v4 - 1
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 76))) = uint32(v5)
		if v5 == 0 {
			*(*uint32)(unsafe.Pointer(uintptr(a1 + 48))) = 0
		}
	}
	v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 48))))
	v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 36))))
	if v6 != 0 {
		if v6 != v7 {
			if v7 != 0 {
				nox_xxx_netStopRaySpell_4FEF90(a1, *(**uint32)(unsafe.Pointer(uintptr(a1 + 36))))
			}
			nox_xxx_netStartDurationRaySpell_4FF130(a1)
		}
		v8 = int32(*(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 48))) + 4))))
		if uint32(uint16(int16(v8))) == *memmap.PtrUint32(0x5D4594, 2487936) || uint32(v8) == *memmap.PtrUint32(0x5D4594, 2487940) {
			v9 = nox_xxx_gamedataGetFloat_419D40(internCStr("PlasmaDamageHecubah"))
		} else {
			v9 = nox_xxx_gamedataGetFloat_419D40(internCStr("PlasmaDamage"))
		}
		v20 = float32(v9)
		v10 = nox_float2int(v20)
		(*(*func(uint32, uint32, uint32, int32, int32))(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 48))) + 716))))(*(*uint32)(unsafe.Pointer(uintptr(a1 + 48))), *(*uint32)(unsafe.Pointer(uintptr(a1 + 16))), 0, v10, 14)
		v11 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 48))))
		if *(*uint32)(unsafe.Pointer(uintptr(v11 + 16)))&0x8020 != 0 {
			nox_xxx_netSendPointFx_522FF0(-125, (*float2)(unsafe.Pointer(uintptr(v11+56))))
		}
		v19 = *(**uint32)(unsafe.Pointer(uintptr(a1 + 16)))
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 36))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 48)))
		nox_xxx_playerSetState_4FA020((*nox_object_t)(unsafe.Pointer(v19)), 22)
		if (gameFrame() % (gameFPS() / 3)) == 0 {
			nox_xxx_aud_501960(98, (*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16)))))), 0, 0)
			nox_xxx_aud_501960(98, (*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 48)))))), 0, 0)
		}
		if *(*uint32)(unsafe.Pointer(uintptr(a1 + 76))) == 0 {
			v21 = float32(nox_xxx_gamedataGetFloat_419D40(internCStr("PlasmaSearchTime")))
			*(*uint32)(unsafe.Pointer(uintptr(a1 + 68))) = gameFrame() + uint32(nox_float2int(v21))
		}
		v12 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 72))))
		if v12 != 0 {
			v13 = int32(*(*uint32)(unsafe.Pointer(uintptr(v12 + 736))))
			v14 = *(*uint8)(unsafe.Pointer(uintptr(v13 + 108)))
			if int32(v14) <= 0 {
				return 1
			}
			v15 = uint8(int8(int32(v14) - 1))
			*(*uint8)(unsafe.Pointer(uintptr(v13 + 108))) = v15
			*(*uint32)(unsafe.Pointer(uintptr(v13 + 112))) = uint32(int32(v15) * 100 / int32(*(*uint8)(unsafe.Pointer(uintptr(v13 + 109)))))
			v16 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))))
			if v16 != 0 {
				if int32(*(*uint8)(unsafe.Pointer(uintptr(v16 + 8))))&4 != 0 {
					v17 = int32(*(*uint32)(unsafe.Pointer(uintptr(v16 + 748))))
					nox_xxx_playerSetState_4FA020((*nox_object_t)(unsafe.Pointer(uintptr(v16))), 22)
					nox_xxx_netReportCharges_4D82B0(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v17 + 276))) + 2064)))), (*nox_object_t)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(uintptr(a1 + 72))))), int8(*(*uint8)(unsafe.Pointer(uintptr(v13 + 108)))), int8(*(*uint8)(unsafe.Pointer(uintptr(v13 + 109)))))
				}
			}
			if int32(*(*uint8)(unsafe.Pointer(uintptr(v13 + 108)))) <= 0 {
				return 1
			}
		}
	} else if v7 != 0 {
		nox_xxx_netStopRaySpell_4FEF90(a1, *(**uint32)(unsafe.Pointer(uintptr(a1 + 36))))
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 36))) = 0
	}
	return 0
}
func sub_531920(a1 int32, a2 int32) {
	var (
		v2 int32
		v3 int32
		v4 float64
		v5 float64
		v6 float64
	)
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))))
	if uint32(v2)&0x20006 != 0 {
		if (*(*uint32)(unsafe.Pointer(uintptr(a1 + 16)))&0x8020) == 0 && a1 != a2 {
			if (v2&2) == 0 || (func() bool {
				v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12))))
				return (v3 & 0x8000) == 0
			}()) {
				if nox_xxx_unitIsEnemyTo_5330C0((*nox_object_t)(unsafe.Pointer(uintptr(a2))), (*nox_object_t)(unsafe.Pointer(uintptr(a1)))) != 0 && nox_server_testTwoPointsAndDirection_4E6E50((*float2)(unsafe.Pointer(uintptr(a2+56))), int32(*(*int16)(unsafe.Pointer(uintptr(a2 + 124)))), (*float2)(unsafe.Pointer(uintptr(a1+56))))&1|0xC != 0 && nox_xxx_unitCanInteractWith_5370E0((*nox_object_t)(unsafe.Pointer(uintptr(a2))), (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0) != 0 {
					v4 = float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 56))) - *(*float32)(unsafe.Pointer(uintptr(a2 + 56))))
					v5 = float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 60))) - *(*float32)(unsafe.Pointer(uintptr(a2 + 60))))
					v6 = v5*v5 + v4*v4
					if v6 < float64(*mem_getFloatPtr(0x587000, 260404)) {
						*mem_getFloatPtr(0x587000, 260404) = float32(v6)
						dword_5d4594_2487932 = uint32(a1)
					}
				}
			}
		}
	}
}
func sub_5319E0(a1 int32) int32 {
	var result int32
	result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 72))))
	if result != 0 {
		result = int32(*(*uint32)(unsafe.Pointer(uintptr(result + 736))))
		*(*uint32)(unsafe.Pointer(uintptr(result + 96))) &= 0xFFFFFFFB
	}
	return result
}
func nox_xxx_spellCreateMoonglow_531A00(a1 *uint32) int32 {
	var (
		v1 int16
		v2 int32
		v3 *uint32
		v4 int32
		v5 int32
		v7 float32
		v8 float32
		v9 float32
	)
	v7 = float32(nox_xxx_gamedataGetFloat_419D40(internCStr("MoonglowEnchantmentDuration")))
	v1 = int16(nox_float2int(v7))
	v2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*12)))
	if v2 == 0 || *(*uint32)(unsafe.Pointer(uintptr(v2 + 16)))&0x8020 != 0 {
		return 1
	}
	if (int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 8)))) & 4) != 4 {
		nox_xxx_buffApplyTo_4FF380((*nox_object_t)(unsafe.Pointer(uintptr(v2))), 15, v1, int8(uint8(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*2)))))
		return 1
	}
	v3 = (*uint32)(unsafe.Pointer(nox_xxx_newObjectByTypeID_4E3810(internCStr("Moonglow"))))
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*18)) = uint32(uintptr(unsafe.Pointer(v3)))
	if v3 == nil {
		return 1
	}
	v4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*12)))
	v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 748))) + 276))))
	if int32(*(*uint8)(unsafe.Pointer(uintptr(v5 + 3680))))&0x10 != 0 {
		v8 = float32(float64(*(*int32)(unsafe.Pointer(uintptr(v5 + 2284)))))
		v9 = float32(float64(*(*int32)(unsafe.Pointer(uintptr(v5 + 2288)))))
	} else {
		v8 = 2944.0
		v9 = 2944.0
	}
	nox_xxx_createAt_4DAA50((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v3)))))), (*nox_object_t)(unsafe.Pointer(uintptr(v4))), v8, v9)
	nox_xxx_buffApplyTo_4FF380((*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*12))))), 1, v1, int8(uint8(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*2)))))
	return 0
}
func sub_531AF0(a1 int32) int32 {
	var result int32
	result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 48))))
	if result != 0 {
		if (int32(*(*uint8)(unsafe.Pointer(uintptr(result + 8)))) & 4) == 4 {
			if *(*uint32)(unsafe.Pointer(uintptr(a1 + 72))) != 0 {
				nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 72)))))))
			}
			*(*uint32)(unsafe.Pointer(uintptr(a1 + 72))) = 0
			result = nox_xxx_spellBuffOff_4FF5B0((*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 48)))))), 1)
		} else {
			result = nox_xxx_spellBuffOff_4FF5B0((*nox_object_t)(unsafe.Pointer(uintptr(result))), 15)
		}
	}
	return result
}
func nox_xxx_TODOsomeCallingMeleeAttack_531B40(a1 int32, a2 int32) *int32 {
	var (
		v2     int32
		result *int32
		v4     int32
		v5     float64
		v6     int32
		v7     int32
	)
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	if nox_xxx_testUnitBuffs_4FF350((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 29) != 0 || (func() *int32 {
		result = (*int32)(unsafe.Pointer(uintptr(nox_xxx_mobCastRelated2_540D90(a1, a2))))
		return result
	}()) == nil {
		if nox_xxx_monsterCanShoot_534280(a1) != 0 {
			v4 = nox_xxx_monsterCanMelee_534220(a1)
			v7 = a2
			v6 = a1
			if v4 != 0 && (func() bool {
				v5 = nox_xxx_calcDistance_4E6C00((*nox_object_t)(unsafe.Pointer(uintptr(a1))), (*nox_object_t)(unsafe.Pointer(uintptr(a2))))
				v7 = a2
				v6 = a1
				return v5 < float64(*(*float32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 484))) + 212))))*0.5
			}()) {
				result = nox_xxx_monsterAction_531C60(a1, a2)
			} else {
				result = sub_531D50(v6, v7)
			}
		} else if nox_xxx_monsterCanMelee_534220(a1) != 0 {
			result = nox_xxx_monsterAction_531C60(a1, a2)
		} else {
			result = (*int32)(unsafe.Pointer(uintptr(nox_xxx_monsterCanCast_534300((*nox_object_t)(unsafe.Pointer(uintptr(a1)))))))
			if result == nil {
				result = sub_531C10(a1, a2)
			}
		}
	}
	return result
}
func sub_531C10(a1 int32, a2 int32) *int32 {
	var (
		v2     *int32
		v3     *int32
		result *int32
	)
	v2 = (*int32)(nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 60, internCStr(__FILE__), __LINE__))
	if v2 != nil {
		*(*int32)(unsafe.Add(unsafe.Pointer(v2), 4*1)) = a2
	}
	v3 = (*int32)(nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 42, internCStr(__FILE__), __LINE__))
	if v3 != nil {
		*(*int32)(unsafe.Add(unsafe.Pointer(v3), 4*1)) = a2
	}
	result = (*int32)(nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 7, internCStr(__FILE__), __LINE__))
	if result != nil {
		*(*int32)(unsafe.Add(unsafe.Pointer(result), 4*1)) = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 56))))
		*(*int32)(unsafe.Add(unsafe.Pointer(result), 4*2)) = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 60))))
		*(*int32)(unsafe.Add(unsafe.Pointer(result), 4*3)) = a2
	}
	return result
}
func nox_xxx_monsterAction_531C60(a1 int32, a2 int32) *int32 {
	var (
		v2     int32
		v3     *int32
		v4     *int32
		v5     *int32
		v6     float64
		v7     *int32
		v8     *int32
		v9     *int32
		v10    int32
		result *int32
	)
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	v3 = (*int32)(nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 60, internCStr(__FILE__), __LINE__))
	if v3 != nil {
		*(*int32)(unsafe.Add(unsafe.Pointer(v3), 4*1)) = a2
	}
	v4 = (*int32)(nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 42, internCStr(__FILE__), __LINE__))
	if v4 != nil {
		*(*int32)(unsafe.Add(unsafe.Pointer(v4), 4*1)) = a2
	}
	if nox_xxx_monsterCanShoot_534280(a1) != 0 {
		v5 = (*int32)(nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 50, internCStr(__FILE__), __LINE__))
		if v5 != nil {
			v6 = float64(*(*float32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 484))) + 212)))) * 0.60000002
			*(*int32)(unsafe.Add(unsafe.Pointer(v5), 4*3)) = a2
			*((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v5))), unsafe.Sizeof(float32(0))*1))) = float32(v6)
		}
	}
	v7 = (*int32)(nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 45, internCStr(__FILE__), __LINE__))
	if v7 != nil {
		*(*int32)(unsafe.Add(unsafe.Pointer(v7), 4*1)) = a2
	}
	nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 16, internCStr(__FILE__), __LINE__)
	v8 = (*int32)(nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 26, internCStr(__FILE__), __LINE__))
	if v8 != nil {
		*(*int32)(unsafe.Add(unsafe.Pointer(v8), 4*1)) = a2
	}
	v9 = (*int32)(nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 49, internCStr(__FILE__), __LINE__))
	if v9 != nil {
		v10 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 484))) + 112))))
		*(*int32)(unsafe.Add(unsafe.Pointer(v9), 4*3)) = a2
		*(*int32)(unsafe.Add(unsafe.Pointer(v9), 4*1)) = v10
	}
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 12))))&0x10 != 0 {
		nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 65, internCStr(__FILE__), __LINE__)
		nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 40, internCStr(__FILE__), __LINE__)
	}
	result = (*int32)(nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 7, internCStr(__FILE__), __LINE__))
	if result != nil {
		*(*int32)(unsafe.Add(unsafe.Pointer(result), 4*1)) = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 56))))
		*(*int32)(unsafe.Add(unsafe.Pointer(result), 4*2)) = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 60))))
		*(*int32)(unsafe.Add(unsafe.Pointer(result), 4*3)) = a2
	}
	return result
}
func sub_531D50(a1 int32, a2 int32) *int32 {
	var (
		v2     int32
		v3     *int32
		v4     *int32
		v5     *int32
		v6     *int32
		result *int32
		v8     *int32
		v9     *int32
		v10    int32
	)
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	v3 = (*int32)(nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 60, internCStr(__FILE__), __LINE__))
	if v3 != nil {
		*(*int32)(unsafe.Add(unsafe.Pointer(v3), 4*1)) = a2
	}
	v4 = (*int32)(nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 45, internCStr(__FILE__), __LINE__))
	if v4 != nil {
		*(*int32)(unsafe.Add(unsafe.Pointer(v4), 4*1)) = a2
	}
	v5 = (*int32)(nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 17, internCStr(__FILE__), __LINE__))
	if v5 != nil {
		*(*int32)(unsafe.Add(unsafe.Pointer(v5), 4*1)) = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 56))))
		*(*int32)(unsafe.Add(unsafe.Pointer(v5), 4*2)) = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 60))))
		*(*int32)(unsafe.Add(unsafe.Pointer(v5), 4*3)) = a2
	}
	v6 = (*int32)(nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 26, internCStr(__FILE__), __LINE__))
	if v6 != nil {
		*(*int32)(unsafe.Add(unsafe.Pointer(v6), 4*1)) = a2
	}
	result = (*int32)(unsafe.Pointer(uintptr(sub_534710(a1))))
	if result == nil {
		v8 = (*int32)(nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 47, internCStr(__FILE__), __LINE__))
		if v8 != nil {
			*(*int32)(unsafe.Add(unsafe.Pointer(v8), 4*1)) = a2
		}
		v9 = (*int32)(nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 49, internCStr(__FILE__), __LINE__))
		if v9 != nil {
			v10 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 484))) + 212))))
			*(*int32)(unsafe.Add(unsafe.Pointer(v9), 4*3)) = a2
			*(*int32)(unsafe.Add(unsafe.Pointer(v9), 4*1)) = v10
		}
		nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 40, internCStr(__FILE__), __LINE__)
		result = (*int32)(nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 7, internCStr(__FILE__), __LINE__))
		if result != nil {
			*(*int32)(unsafe.Add(unsafe.Pointer(result), 4*1)) = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 56))))
			*(*int32)(unsafe.Add(unsafe.Pointer(result), 4*2)) = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 60))))
			*(*int32)(unsafe.Add(unsafe.Pointer(result), 4*3)) = a2
		}
	}
	return result
}
func nox_xxx_mobActionFightStart_531E20(a1 uint32) int32 {
	var (
		v1 *int32
		v2 int32
		v3 int32
	)
	v1 = *(**int32)(unsafe.Pointer(uintptr(a1 + 748)))
	v2 = int32(uintptr(nox_xxx_monsterGetSoundSet_424300((*nox_object_t)(unsafe.Pointer(uintptr(a1))))))
	if v2 != 0 {
		nox_xxx_aud_501960(int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 20)))), (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
	}
	nox_xxx_scriptCallByEventBlock_502490(unsafe.Pointer((*int32)(unsafe.Add(unsafe.Pointer(v1), 4*310))), unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(v1), 4*299)))), unsafe.Pointer(uintptr(a1)), 13)
	v3 = *(*int32)(unsafe.Add(unsafe.Pointer(v1), 4*360))
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v3))), 1)) |= 1
	*(*int32)(unsafe.Add(unsafe.Pointer(v1), 4*360)) = v3
	nox_xxx_frameCounterSetCopy_5281E0()
	nox_xxx_unitUpdateSightMB_5281F0((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
	return sub_534750(int32(a1))
}
func sub_531E90(a1 int32) int32 {
	var (
		v1 int32
		v2 int32
	)
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 1440))))
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v2))), 1)) &= 0xFE
	*(*uint32)(unsafe.Pointer(uintptr(v1 + 1440))) = uint32(v2)
	return sub_534780(a1)
}
func nox_xxx_mobActionFight_531EC0(a1 int32) int8 {
	var (
		v1 int32
		v2 int32
		v3 *int32
		v4 *float32
		v5 float64
		v6 float64
	)
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	v2 = v1 + int32((*(*byte)(unsafe.Pointer(uintptr(v1 + 544)))+23)*24)
	if gameFrame()-*(*uint32)(unsafe.Pointer(uintptr(v1 + int32((*(*byte)(unsafe.Pointer(uintptr(v1 + 544)))+23)*24) + 12))) > (gameFPS() * 10) {
		*((*uint8)(unsafe.Pointer(&v3))) = uint8(int8(nox_xxx_monsterPopAction_50A160((*nox_object_t)(unsafe.Pointer(uintptr(a1))))))
		return int8(uintptr(unsafe.Pointer(v3)))
	}
	if *(*uint32)(unsafe.Pointer(uintptr(v1 + 1196))) != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(v2 + 12))) = gameFrame()
		if nox_xxx_checkIsKillable_528190((*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 1196))))))) != 0 {
			if nox_xxx_testUnitBuffs_4FF350((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 29) != 0 || (func() *int32 {
				v3 = (*int32)(unsafe.Pointer(uintptr(nox_xxx_monsterBuffSelf_540B90(a1))))
				return v3
			}()) == nil && (func() *int32 {
				v3 = (*int32)(unsafe.Pointer(uintptr(nox_xxx_monsterCastOffensive_540F20(a1, int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 1196))))))))
				return v3
			}()) == nil {
				*((*uint8)(unsafe.Pointer(&v3))) = uint8(uint32(uintptr(unsafe.Pointer(nox_xxx_TODOsomeCallingMeleeAttack_531B40(a1, int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 1196)))))))))
			}
			return int8(uintptr(unsafe.Pointer(v3)))
		}
		*((*uint8)(unsafe.Pointer(&v3))) = uint8(int8(nox_xxx_monsterPopAction_50A160((*nox_object_t)(unsafe.Pointer(uintptr(a1))))))
		return int8(uintptr(unsafe.Pointer(v3)))
	}
	if sub_534710(a1) != 0 {
		*((*uint8)(unsafe.Pointer(&v3))) = uint8(int8(nox_xxx_monsterPopAction_50A160((*nox_object_t)(unsafe.Pointer(uintptr(a1))))))
		return int8(uintptr(unsafe.Pointer(v3)))
	}
	v4 = (*float32)(unsafe.Pointer(uintptr(v2 + 4)))
	*memmap.PtrUint32(0x5D4594, 2487944) = 0
	nox_xxx_unitsGetInCircle_517F90((*float2)(unsafe.Pointer(uintptr(v2+4))), 30.0, unsafe.Pointer(funAddr(sub_532040)), unsafe.Pointer(uintptr(v1)))
	if *memmap.PtrUint32(0x5D4594, 2487944) != 0 {
		nox_xxx_monsterPopAction_50A160((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
		v3 = *(**int32)(unsafe.Pointer(uintptr(v1 + 1200)))
		if *(**int32)(unsafe.Pointer(uintptr(v1 + 392))) == v3 {
			*(*uint32)(unsafe.Pointer(uintptr(v1 + 388))) = 0
		}
	} else {
		v5 = float64(*v4 - *(*float32)(unsafe.Pointer(uintptr(a1 + 56))))
		v6 = float64(*(*float32)(unsafe.Pointer(uintptr(v2 + 8))) - *(*float32)(unsafe.Pointer(uintptr(a1 + 60))))
		if v6*v6+v5*v5 < 64.0 {
			*((*uint8)(unsafe.Pointer(&v3))) = uint8(int8(nox_xxx_monsterPopAction_50A160((*nox_object_t)(unsafe.Pointer(uintptr(a1))))))
			return int8(uintptr(unsafe.Pointer(v3)))
		}
		nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 56, internCStr(__FILE__), __LINE__)
		v3 = (*int32)(nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 7, internCStr(__FILE__), __LINE__))
		if v3 != nil {
			*(*int32)(unsafe.Add(unsafe.Pointer(v3), 4*1)) = int32(*(*uint32)(unsafe.Pointer(v4)))
			*(*int32)(unsafe.Add(unsafe.Pointer(v3), 4*2)) = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))))
			*(*int32)(unsafe.Add(unsafe.Pointer(v3), 4*3)) = 0
		}
	}
	return int8(uintptr(unsafe.Pointer(v3)))
}
func sub_532040(a1 int32, a2 int32) int32 {
	var (
		result int32
		v3     int32
	)
	result = a1
	if *(*uint32)(unsafe.Pointer(uintptr(a1 + 36))) == *(*uint32)(unsafe.Pointer(uintptr(a2 + 1200))) {
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))))
		if (v3 & 0x8000) != 0 {
			*memmap.PtrUint32(0x5D4594, 2487944) = 1
		}
	}
	return result
}
func nox_xxx_monsterShieldBlockStart_532070(a1 int32) int8 {
	var (
		v1 int32
		v2 uint32
	)
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + uint32((*(*byte)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + 544)))+23)*24))
	if nox_xxx_monsterTestBlockShield_533E70((*nox_object_t)(unsafe.Pointer(uintptr(a1)))) != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(v1 + 4))) = gameFrame() + (gameFPS() >> 1)
	}
	v2 = *(*uint32)(unsafe.Pointer(uintptr(v1 + 4)))
	if gameFrame() > v2 {
		nox_xxx_monsterPopAction_50A160((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
		*((*uint8)(unsafe.Pointer(&v2))) = *(*uint8)(unsafe.Pointer(uintptr(a1 + 12)))
		if (v2 & 0x10) == 0 {
			*((*uint8)(unsafe.Pointer(&v2))) = uint8(uint32(uintptr(nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 22, internCStr(__FILE__), __LINE__))))
		}
	}
	return int8(uint8(v2))
}
func nox_xxx_monsterShieldBlockStop_5320E0(a1 int32) int8 {
	var result int8
	result = int8(a1)
	if int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + 483)))) != 0 {
		result = int8(nox_xxx_monsterPopAction_50A160((*nox_object_t)(unsafe.Pointer(uintptr(a1)))))
	}
	return result
}
func nox_ai_action_pop_532100(a1 int32) int8 {
	return int8(nox_xxx_monsterPopAction_50A160((*nox_object_t)(unsafe.Pointer(uintptr(a1)))))
}
func sub_532110(a1 int32) int8 {
	var result int8
	result = int8(a1)
	if int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + 483)))) != 0 {
		result = int8(nox_xxx_monsterPopAction_50A160((*nox_object_t)(unsafe.Pointer(uintptr(a1)))))
	}
	return result
}
func nox_xxx_mobActionMelee1_532130(a1 int32) *int32 {
	var (
		v1     int32
		v2     *int32
		result *int32
		v4     int32
		v5     float64
		v6     int32
		v7     int32
		v8     int32
		v9     *int32
		v10    *int32
		v11    int32
		v12    int32
		v13    int32
		v14    float4
	)
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	if gameFrame() >= uint32(*(*int32)(unsafe.Pointer(uintptr(v1 + 512)))) {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 12))))&0x10 != 0 {
			v4 = nox_xxx_weaponGetStaminaByType_4F7E80(int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 2056)))))
			if v4 > int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 1128)))) {
				*(*uint8)(unsafe.Pointer(uintptr(v1 + 1128))) -= uint8(int8(v4))
			} else {
				*(*uint8)(unsafe.Pointer(uintptr(v1 + 1128))) = 0
			}
		}
		v14.field_0 = *(*float32)(unsafe.Pointer(uintptr(a1 + 56))) - *(*float32)(unsafe.Pointer(&dword_587000_261388))
		v14.field_4 = *(*float32)(unsafe.Pointer(uintptr(a1 + 60))) - *(*float32)(unsafe.Pointer(&dword_587000_261388))
		v14.field_8 = *(*float32)(unsafe.Pointer(&dword_587000_261388)) + *(*float32)(unsafe.Pointer(uintptr(a1 + 56)))
		v5 = float64(*(*float32)(unsafe.Pointer(&dword_587000_261388)) + *(*float32)(unsafe.Pointer(uintptr(a1 + 60))))
		dword_5d4594_2487948 = 0
		v14.field_C = float32(v5)
		*mem_getFloatPtr(0x5D4594, 2487952) = float32(float64(*(*float32)(unsafe.Pointer(&dword_587000_261388))) + 1.0)
		nox_xxx_getUnitsInRect_517C10(&v14, unsafe.Pointer(funAddr(sub_532390)), unsafe.Pointer(uintptr(a1)))
		if dword_5d4594_2487948 != 0 && nox_xxx_unitIsEnemyTo_5330C0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), (*nox_object_t)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_2487948)))))) == 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 12))))&0x10 != 0 && (func() int32 {
			v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 2064))))
			return v6
		}()) != 0 && (func() int32 {
			v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(v6 + 12))))
			return int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v7))), 1))) & 0x40
		}()) != 0 {
			v13 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 36))))
			v8 = int32(uintptr(unsafe.Pointer(nox_xxx_getUnitName_4E39D0((*nox_object_t)(unsafe.Pointer(uintptr(a1)))))))
			nox_ai_debug_printf_5341A0(internCStr("%d: %s(#%d) Tried to MELEE_ATTACK but friend in the way\n"), gameFrame(), v8, v13)
			nox_xxx_monsterPopAction_50A160((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
			v9 = (*int32)(nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 27, internCStr(__FILE__), __LINE__))
			if v9 != nil {
				*(*int32)(unsafe.Add(unsafe.Pointer(v9), 4*1)) = int32(*(*int16)(unsafe.Pointer(uintptr(a1 + 124))))
			}
			v10 = (*int32)(nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 41, internCStr(__FILE__), __LINE__))
			if v10 != nil {
				*(*int32)(unsafe.Add(unsafe.Pointer(v10), 4*1)) = int32(gameFrame() + uint32(nox_common_randomInt_415FA0(int32(gameFPS()>>2), int32(gameFPS()>>1))))
			}
			result = (*int32)(nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 24, internCStr(__FILE__), __LINE__))
			if result != nil {
				v11 = int32(dword_5d4594_2487948 + 56)
				*(*int32)(unsafe.Add(unsafe.Pointer(result), 4*1)) = int32(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_2487948 + 56))))
				v12 = int32(*(*uint32)(unsafe.Pointer(uintptr(v11 + 4))))
				*(*int32)(unsafe.Add(unsafe.Pointer(result), 4*3)) = 0
				*(*int32)(unsafe.Add(unsafe.Pointer(result), 4*2)) = v12
			}
		} else {
			sub_5341D0(a1)
			*(*uint32)(unsafe.Pointer(uintptr(a1 + 136))) = gameFrame()
			*(*uint32)(unsafe.Pointer(uintptr(v1 + 512))) = gameFrame() + uint32(nox_common_randomInt_415FA0(int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 484))) + 128)))), int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 484))) + 132))))))
			result = (*int32)(nox_xxx_monsterGetSoundSet_424300((*nox_object_t)(unsafe.Pointer(uintptr(a1)))))
			if result != nil {
				nox_xxx_aud_501960(*(*int32)(unsafe.Add(unsafe.Pointer(result), 4*6)), (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
			}
		}
	} else {
		v2 = (*int32)(nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 50, internCStr(__FILE__), __LINE__))
		if v2 != nil {
			*((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v2))), unsafe.Sizeof(float32(0))*1))) = float32(float64(*(*float32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 484))) + 112)))) * 1.2)
			*(*int32)(unsafe.Add(unsafe.Pointer(v2), 4*3)) = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 1196))))
		}
		result = (*int32)(nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 1, internCStr(__FILE__), __LINE__))
		if result != nil {
			*(*int32)(unsafe.Add(unsafe.Pointer(result), 4*1)) = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 512))))
		}
	}
	return result
}
func sub_532390(a1 int32, a2 int32) {
	var (
		v2 int32
		v3 int32
		v4 float64
		v5 float64
		v6 *float32
		v7 float64
		v8 float32
		v9 float32
	)
	v2 = a1
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&6 != 0 {
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))))
		if (v3 & 0x8000) == 0 {
			v4 = float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 56))) - *(*float32)(unsafe.Pointer(uintptr(a2 + 56))))
			v5 = float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 60))) - *(*float32)(unsafe.Pointer(uintptr(a2 + 60))))
			v8 = float32(v5)
			v9 = float32(math.Sqrt(v5*float64(v8)+v4*v4) + 9.9999997e-05)
			v7 = v4 / float64(v9)
			if float64(v9) < float64(*mem_getFloatPtr(0x5D4594, 2487952)) {
				v6 = mem_getFloatPtr(0x587000, uint32(int32(*(*int16)(unsafe.Pointer(uintptr(a2 + 124))))*8)+194136)
				if float64(v8/v9**(*float32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(float32(0))*1)))+v7*float64(*v6) > 0.5 {
					dword_5d4594_2487948 = uint32(v2)
					*mem_getFloatPtr(0x5D4594, 2487952) = v9
				}
			}
		}
	}
}
func nox_xxx_mobActionMeleeAtt_532440(a1 int32) int8 {
	var (
		v1  int32
		v2  int32
		v3  int32
		v4  int32
		v5  func(int32) int32
		v6  int32
		v7  int32
		v8  int32
		v10 int32
	)
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 12))))&0x10 != 0 {
		if *(*uint32)(unsafe.Pointer(uintptr(v1 + 1440)))&0x20000 != 0 {
			nox_xxx_mobMorphToPlayer_4FAAF0((*uint32)(unsafe.Pointer(uintptr(a1))))
		}
		v2 = nox_xxx_playerAttack_538960((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 1440))))
		if uint32(v3)&0x20000 != 0 {
			*((*uint8)(unsafe.Pointer(&v3))) = uint8(nox_xxx_mobMorphFromPlayer_4FAAC0((*uint32)(unsafe.Pointer(uintptr(a1)))))
		}
		if v2 == 0 {
			*((*uint8)(unsafe.Pointer(&v3))) = uint8(int8(nox_xxx_monsterPopAction_50A160((*nox_object_t)(unsafe.Pointer(uintptr(a1))))))
		}
	} else {
		v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 484))))
		v5 = *(*func(int32) int32)(unsafe.Pointer(uintptr(v4 + 236)))
		if v5 != nil {
			if uint32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 481)))) == *(*uint32)(unsafe.Pointer(uintptr(v4 + 108))) && int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 482)))) == 0 {
				v7 = v5(a1)
				v8 = int32(uintptr(nox_xxx_monsterGetSoundSet_424300((*nox_object_t)(unsafe.Pointer(uintptr(a1))))))
				if v8 != 0 {
					if v7 != 0 {
						nox_xxx_aud_501960(int32(*(*uint32)(unsafe.Pointer(uintptr(v8 + 32)))), (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
					} else {
						nox_xxx_aud_501960(int32(*(*uint32)(unsafe.Pointer(uintptr(v8 + 36)))), (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
					}
				}
			}
			*((*uint8)(unsafe.Pointer(&v3))) = *(*uint8)(unsafe.Pointer(uintptr(v1 + 483)))
			if int32(uint8(int8(v3))) != 0 {
				*((*uint8)(unsafe.Pointer(&v3))) = uint8(int8(nox_xxx_monsterPopAction_50A160((*nox_object_t)(unsafe.Pointer(uintptr(a1))))))
			}
		} else {
			v10 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 36))))
			v6 = int32(uintptr(unsafe.Pointer(nox_xxx_getUnitName_4E39D0((*nox_object_t)(unsafe.Pointer(uintptr(a1)))))))
			nox_ai_debug_printf_5341A0(internCStr("%d: %s(#%d) Tried to MELEE_ATTACK but cannot\n"), gameFrame(), v6, v10)
			*((*uint8)(unsafe.Pointer(&v3))) = uint8(int8(nox_xxx_monsterPopAction_50A160((*nox_object_t)(unsafe.Pointer(uintptr(a1))))))
		}
	}
	return int8(v3)
}
func sub_532540(a1 int32) *uint32 {
	var (
		v1     *uint32
		v2     *int32
		result *uint32
	)
	v1 = *(**uint32)(unsafe.Pointer(uintptr(a1 + 748)))
	if gameFrame() >= *(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*128)) {
		sub_5341D0(a1)
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 136))) = gameFrame()
		*(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*128)) = gameFrame() + uint32(nox_common_randomInt_415FA0(int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*121)) + 220)))), int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*121)) + 224))))))
		result = (*uint32)(nox_xxx_monsterGetSoundSet_424300((*nox_object_t)(unsafe.Pointer(uintptr(a1)))))
		if result != nil {
			nox_xxx_aud_501960(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(result), 4*10))), (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
		}
	} else {
		v2 = (*int32)(nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 50, internCStr(__FILE__), __LINE__))
		if v2 != nil {
			*((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v2))), unsafe.Sizeof(float32(0))*1))) = float32(float64(*(*float32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*121)) + 212)))) * 1.2)
			*(*int32)(unsafe.Add(unsafe.Pointer(v2), 4*3)) = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*299)))
		}
		result = (*uint32)(nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 1, internCStr(__FILE__), __LINE__))
		if result != nil {
			*(*uint32)(unsafe.Add(unsafe.Pointer(result), 4*1)) = *(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*128))
		}
	}
	return result
}
func nox_xxx_mobActionMissileAtt_532610(a1 int32) int8 {
	var (
		v1  int32
		v2  int32
		v3  int32
		v4  *int32
		v5  int32
		v6  *float32
		v7  int32
		v8  float32
		v9  float64
		v10 float64
		v11 int32
		v12 float32
		v13 int16
		v14 int32
		v16 float32
		v17 float32
		v18 float32
		v19 float32
		v20 float32
		v21 float32
		v22 float2
		v24 float4
		v25 float32
	)
	v1 = a1
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 12))))&0x10 != 0 {
		v3 = nox_xxx_playerAttack_538960((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
		if v3 == 0 {
			*((*uint8)(unsafe.Pointer(&v3))) = uint8(int8(nox_xxx_monsterPopAction_50A160((*nox_object_t)(unsafe.Pointer(uintptr(a1))))))
		}
	} else {
		v4 = (*int32)(unsafe.Pointer(uintptr(v2 + int32((*(*byte)(unsafe.Pointer(uintptr(v2 + 544)))+23)*24))))
		v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 484))))
		if uint32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 481)))) == *(*uint32)(unsafe.Pointer(uintptr(v5 + 216))) && int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 482)))) == 0 {
			v6 = (*float32)(unsafe.Pointer(nox_xxx_newObjectByTypeID_4E3810((*byte)(unsafe.Pointer(uintptr(v5 + 148))))))
			if v6 != nil {
				v7 = *(*int32)(unsafe.Add(unsafe.Pointer(v4), 4*3))
				if v7 != 0 {
					nox_xxx_projAddVelocitySmth_533080(a1, v7, *(*float32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(float32(0))*136)), int32(uintptr(unsafe.Pointer(&v22))))
				} else {
					v8 = *(*float32)(unsafe.Pointer((*int32)(unsafe.Add(unsafe.Pointer(v4), 4*2))))
					v22.field_0 = *(*float32)(unsafe.Pointer((*int32)(unsafe.Add(unsafe.Pointer(v4), 4*1))))
					v22.field_4 = v8
				}
				v9 = float64(v22.field_0 - *(*float32)(unsafe.Pointer(uintptr(a1 + 56))))
				v10 = float64(v22.field_4 - *(*float32)(unsafe.Pointer(uintptr(a1 + 60))))
				v18 = float32(v10)
				v11 = int32(*(*int16)(unsafe.Pointer(uintptr(a1 + 124)))) * 8
				v12 = *(*float32)(unsafe.Pointer(uintptr(a1 + 60)))
				v24.field_0 = *(*float32)(unsafe.Pointer(uintptr(a1 + 56)))
				v24.field_4 = v12
				v25 = float32(math.Sqrt(v10*float64(v18)+v9*v9) + 0.1)
				v16 = float32(float64(*(*float32)(unsafe.Pointer(uintptr(v1 + 176)))) + 4.0)
				v20 = float32(v9 * float64(*(*float32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(float32(0))*136))) / float64(v25))
				v21 = v18 * *(*float32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(float32(0))*136)) / v25
				v17 = v16**mem_getFloatPtr(0x587000, uint32(v11)+194136) + *(*float32)(unsafe.Pointer(uintptr(v1 + 56)))
				v19 = v16**mem_getFloatPtr(0x587000, uint32(v11)+194140) + *(*float32)(unsafe.Pointer(uintptr(v1 + 60)))
				v24.field_8 = v17 + v20
				v24.field_C = v19 + v21
				if nox_xxx_mapTraceRay_535250(&v24, nil, nil, 5) != 0 {
					nox_xxx_createAt_4DAA50((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v6)))))), (*nox_object_t)(unsafe.Pointer(uintptr(v1))), v17, v19)
					*(*float32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(float32(0))*20)) = v20
					*(*float32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(float32(0))*21)) = v21
					v13 = int16(*(*uint16)(unsafe.Pointer(uintptr(v1 + 124))))
					*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v6))), unsafe.Sizeof(uint16(0))*62))) = uint16(v13)
					*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v6))), unsafe.Sizeof(uint16(0))*63))) = uint16(v13)
				} else {
					nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v6)))))))
				}
			}
			v14 = int32(uintptr(nox_xxx_monsterGetSoundSet_424300((*nox_object_t)(unsafe.Pointer(uintptr(v1))))))
			if v14 != 0 {
				nox_xxx_aud_501960(int32(*(*uint32)(unsafe.Pointer(uintptr(v14 + 44)))), (*nox_object_t)(unsafe.Pointer(uintptr(v1))), 0, 0)
			}
		}
		*((*uint8)(unsafe.Pointer(&v3))) = *(*uint8)(unsafe.Pointer(uintptr(v2 + 483)))
		if int32(uint8(int8(v3))) != 0 {
			*((*uint8)(unsafe.Pointer(&v3))) = uint8(int8(nox_xxx_monsterPopAction_50A160((*nox_object_t)(unsafe.Pointer(uintptr(v1))))))
		}
	}
	return int8(v3)
}
func nox_xxx_monsterPlayHurtSound_532800(a1p *nox_object_t) int8 {
	var (
		a1 int32 = int32(uintptr(unsafe.Pointer(a1p)))
		v1 int32
		v2 int32
	)
	*((*uint8)(unsafe.Pointer(&v1))) = *(*uint8)(unsafe.Pointer(uintptr(a1 + 8)))
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	if v1&2 != 0 {
		*((*uint8)(unsafe.Pointer(&v1))) = uint8(gameFrame())
		if gameFrame() >= uint32(*(*int32)(unsafe.Pointer(uintptr(v2 + 532)))) {
			*(*uint32)(unsafe.Pointer(uintptr(v2 + 532))) = gameFrame() + uint32(nox_common_randomInt_415FA0(int32(gameFPS()*2), int32(gameFPS()*4)))
			v1 = int32(uintptr(nox_xxx_monsterGetSoundSet_424300((*nox_object_t)(unsafe.Pointer(uintptr(a1))))))
			if v1 != 0 {
				nox_xxx_aud_501960(int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 8)))), (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
			}
		}
	}
	return int8(v1)
}
func sub_532880(a1 int32) int32 {
	var result int32
	result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	if *(*uint32)(unsafe.Pointer(uintptr(result + 520))) == 0 {
		*(*uint32)(unsafe.Pointer(uintptr(result + 520))) = gameFrame()
	}
	return result
}
func nox_xxx_soundPlayerDamageSound_5328B0(a1 int32, a2 int32) int32 {
	var (
		v2 uint16
		i  int32
	)
	v2 = 1
	if a2 == 0 {
		return 1
	}
	for i = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 504)))); i != 0; i = int32(*(*uint32)(unsafe.Pointer(uintptr(i + 496)))) {
		if *(*uint32)(unsafe.Pointer(uintptr(i + 8)))&0x2000000 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(i + 12))))&4 != 0 {
			if sub_4133D0(i) != 0 {
				v2 = 0x2000
			} else if int32(*(*uint16)(unsafe.Pointer(uintptr(i + 24)))) > int32(v2) {
				v2 = *(*uint16)(unsafe.Pointer(uintptr(i + 24)))
			}
		}
	}
	sub_532930(a1, v2, *(*uint16)(unsafe.Pointer(uintptr(a2 + 24))))
	return 1
}
func sub_532930(a1 int32, a2 uint16, a3 uint16) {
	var result int32
	_ = result
	result = int32(a2)
	if int32(a2) <= 32 {
		if int32(a2) == 32 {
			result = int32(a3)
			if int32(a3) > 16 {
				if int32(a3) == 32 {
					nox_xxx_aud_501960(848, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
				} else if int32(a3) == 0x2000 {
					nox_xxx_aud_501960(868, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
				}
			} else if int32(a3) == 16 {
				nox_xxx_aud_501960(858, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
			} else {
				result = int32(a3) - 1
				switch a3 {
				case 1:
					fallthrough
				case 2:
					fallthrough
				case 4:
					nox_xxx_aud_501960(863, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
				case 8:
					nox_xxx_aud_501960(853, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
				default:
					return
				}
			}
		} else {
			result = int32(a2) - 1
			switch a2 {
			case 1:
				fallthrough
			case 2:
				fallthrough
			case 4:
				result = int32(a3)
				if int32(a3) > 16 {
					if int32(a3) == 32 {
						nox_xxx_aud_501960(852, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
					} else if int32(a3) == 0x2000 {
						nox_xxx_aud_501960(872, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
					}
				} else if int32(a3) == 16 {
					nox_xxx_aud_501960(862, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
				} else {
					result = int32(a3) - 1
					switch a3 {
					case 1:
						fallthrough
					case 2:
						fallthrough
					case 4:
						nox_xxx_aud_501960(867, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
					case 8:
						nox_xxx_aud_501960(857, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
					default:
						return
					}
				}
			case 8:
				result = int32(a3)
				if int32(a3) > 16 {
					if int32(a3) == 32 {
						nox_xxx_aud_501960(850, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
					} else if int32(a3) == 0x2000 {
						nox_xxx_aud_501960(870, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
					}
				} else if int32(a3) == 16 {
					nox_xxx_aud_501960(860, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
				} else {
					result = int32(a3) - 1
					switch a3 {
					case 1:
						fallthrough
					case 2:
						fallthrough
					case 4:
						nox_xxx_aud_501960(865, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
					case 8:
						nox_xxx_aud_501960(855, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
					default:
						return
					}
				}
			case 0x10:
				result = int32(a3)
				if int32(a3) > 16 {
					if int32(a3) == 32 {
						nox_xxx_aud_501960(851, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
					} else if int32(a3) == 0x2000 {
						nox_xxx_aud_501960(871, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
					}
				} else if int32(a3) == 16 {
					nox_xxx_aud_501960(861, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
				} else {
					result = int32(a3) - 1
					switch a3 {
					case 1:
						fallthrough
					case 2:
						fallthrough
					case 4:
						nox_xxx_aud_501960(866, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
					case 8:
						nox_xxx_aud_501960(856, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
					default:
						return
					}
				}
			default:
				return
			}
		}
		return
	}
	if int32(a2) != 64 && int32(a2) != 1024 {
		if int32(a2) != 0x2000 {
			return
		}
		result = int32(a3)
		if int32(a3) > 16 {
			if int32(a3) == 32 {
				nox_xxx_aud_501960(868, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
				return
			}
			if int32(a3) != 0x2000 {
				return
			}
		} else if int32(a3) != 16 {
			result = int32(a3) - 1
			switch a3 {
			case 1:
				fallthrough
			case 2:
				fallthrough
			case 4:
				nox_xxx_aud_501960(872, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
			case 8:
				nox_xxx_aud_501960(870, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
			default:
				return
			}
			return
		}
		nox_xxx_aud_501960(871, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
		return
	}
	result = int32(a3)
	if int32(a3) > 16 {
		if int32(a3) == 32 {
			nox_xxx_aud_501960(849, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
		} else if int32(a3) == 0x2000 {
			nox_xxx_aud_501960(869, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
		}
	} else if int32(a3) == 16 {
		nox_xxx_aud_501960(859, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
	} else {
		result = int32(a3) - 1
		switch a3 {
		case 1:
			fallthrough
		case 2:
			fallthrough
		case 4:
			nox_xxx_aud_501960(864, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
		case 8:
			nox_xxx_aud_501960(854, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
		default:
			return
		}
	}
}
func nox_xxx_soundDefaultDamageSound_532E20(a1p *nox_object_t, a2p *nox_object_t) int32 {
	var (
		a1 int32 = int32(uintptr(unsafe.Pointer(a1p)))
		a2 int32 = int32(uintptr(unsafe.Pointer(a2p)))
		v3 int32
	)
	if a1 != 0 && *(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))&0x1001000 != 0 && *(*uint32)(unsafe.Pointer(uintptr(a1 + 492))) != 0 {
		return 1
	}
	if a2 != 0 && int32(*(*uint16)(unsafe.Pointer(uintptr(a2 + 24)))) != 0x4000 && sub_532F70(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 524))))) != 0 {
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 556))))
		if v3 != 0 && int32(*(*uint16)(unsafe.Pointer(uintptr(v3 + 4)))) != 0 && sub_532FB0(int16(*(*uint16)(unsafe.Pointer(uintptr(a1 + 24))))) != 0 {
			sub_532EC0(a1, *(*uint16)(unsafe.Pointer(uintptr(a1 + 24))))
			return 1
		}
		sub_532930(a1, *(*uint16)(unsafe.Pointer(uintptr(a1 + 24))), *(*uint16)(unsafe.Pointer(uintptr(a2 + 24))))
	}
	return 1
}
func sub_532EC0(a1 int32, a2 uint16) {
	if int32(a2) <= 64 {
		if int32(a2) != 64 {
			switch a2 {
			case 8:
				nox_xxx_aud_501960(875, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
				return
			case 0x10:
				nox_xxx_aud_501960(876, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
				return
			case 0x20:
				nox_xxx_aud_501960(873, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
			}
			return
		}
		nox_xxx_aud_501960(874, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
		return
	}
	if int32(a2) == 1024 {
		nox_xxx_aud_501960(874, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
		return
	}
	if int32(a2) == 4096 {
		nox_xxx_aud_501960(877, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
	}
}
func sub_532F70(a1 int32) int32 {
	var result int32
	switch a1 {
	case 0:
		fallthrough
	case 2:
		fallthrough
	case 3:
		fallthrough
	case 10:
		fallthrough
	case 11:
		result = 1
	default:
		result = 0
	}
	return result
}
func sub_532FB0(a1 int16) int32 {
	return bool2int32(int32(a1) == 8 || int32(a1) == 32 || int32(a1) == 64)
}
func sub_532FE0(a1 uint16, a2 int32) int32 {
	if a2 != 0 {
		sub_532930(a2, a1, *(*uint16)(unsafe.Pointer(uintptr(a2 + 24))))
	}
	return 1
}
func sub_533010(a1 uint16, a2 int32) int32 {
	if a2 != 0 {
		sub_532EC0(a2, a1)
	}
	return 1
}
func nox_xxx_projAddVelocitySmth_533080(a1 int32, a2 int32, a3 float32, a4 int32) int32 {
	var (
		result int32
		v5     float64
		v6     float64
		v7     float64
	)
	result = a2
	v5 = float64(*(*float32)(unsafe.Pointer(uintptr(a2 + 56))) - *(*float32)(unsafe.Pointer(uintptr(a1 + 56))))
	v6 = float64(*(*float32)(unsafe.Pointer(uintptr(a2 + 60))) - *(*float32)(unsafe.Pointer(uintptr(a1 + 60))))
	v7 = math.Sqrt(v6*v6+v5*v5) / float64(a3)
	*(*float32)(unsafe.Pointer(uintptr(a4))) = float32(v7*float64(*(*float32)(unsafe.Pointer(uintptr(a2 + 80)))) + float64(*(*float32)(unsafe.Pointer(uintptr(a2 + 56)))))
	*(*float32)(unsafe.Pointer(uintptr(a4 + 4))) = float32(v7*float64(*(*float32)(unsafe.Pointer(uintptr(a2 + 84)))) + float64(*(*float32)(unsafe.Pointer(uintptr(a2 + 60)))))
	return result
}
func nox_xxx_mobActionToAnimation_533790(a1 int32) int32 {
	var (
		v1     int32
		v2     int8
		v3     uint32
		result int32
	)
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	v2 = int8(*(*uint8)(unsafe.Pointer(uintptr(v1 + 544))))
	if int32(v2) == -1 {
		v3 = 8
	} else {
		switch *(*uint32)(unsafe.Pointer(uintptr(v1 + (int32(v2)+23)*24))) {
		case 7:
			fallthrough
		case 8:
			fallthrough
		case 0xA:
			fallthrough
		case 0xD:
			fallthrough
		case 0x1D:
			fallthrough
		case 0x24:
			fallthrough
		case 0x25:
			v3 = (*(*uint32)(unsafe.Pointer(uintptr(v1 + 1440)))&0x4000 | 0x30000) >> 14
		case 9:
			v3 = 12
		case 0x10:
			v3 = 1
		case 0x11:
			v3 = 3
		case 0x12:
			fallthrough
		case 0x13:
			fallthrough
		case 0x14:
			v3 = 7
		case 0x15:
			fallthrough
		case 0x17:
			v3 = 5
		case 0x16:
			v3 = 6
		case 0x18:
			v3 = 13
		case 0x1E:
			v3 = 9
		case 0x1F:
			v3 = 10
		case 0x21:
			fallthrough
		case 0x23:
			v3 = 14
		case 0x22:
			v3 = 15
		default:
			v3 = 8
		}
	}
	if nox_xxx_unitIsMimic_534840(a1) != 0 && v3 == 8 {
		if *(*uint32)(unsafe.Pointer(uintptr(v1 + 1440)))&0x40000 != 0 {
			return 0
		}
		return int32(v3)
	}
	if nox_xxx_unitIsPlant_534A10(a1) != 0 && v3 == 8 {
		if *(*uint32)(unsafe.Pointer(uintptr(v1 + 1196))) == 0 {
			return 14
		}
		return int32(v3)
	}
	if nox_xxx_unitIsZombie_534A40(a1) == 0 {
		return int32(v3)
	}
	if v3 != 9 {
		return int32(v3)
	}
	result = 15
	if (*(*uint32)(unsafe.Pointer(uintptr(v1 + 1440))) & 0x80000) == 0 {
		return int32(v3)
	}
	return result
}
func nox_xxx_orderUnit_533900(ownerp *nox_object_t, creaturep *nox_object_t, orderType int32) {
	var (
		owner    int32 = int32(uintptr(unsafe.Pointer(ownerp)))
		creature int32 = int32(uintptr(unsafe.Pointer(creaturep)))
		i        int32
	)
	if owner != 0 {
		if creature != 0 {
			nox_xxx_enactUnitOrder_5339A0(owner, creature, orderType)
		} else {
			if int32(*(*uint8)(unsafe.Pointer(uintptr(owner + 8))))&4 != 0 && (orderType == 4 || orderType == 3 || orderType == 5) {
				nox_xxx_orderUnitLocal_500C70(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(owner + 748))) + 276))) + 2064)))), orderType)
			}
			for i = int32(*(*uint32)(unsafe.Pointer(uintptr(owner + 516)))); i != 0; i = int32(*(*uint32)(unsafe.Pointer(uintptr(i + 512)))) {
				if int32(*(*uint8)(unsafe.Pointer(uintptr(i + 8))))&2 != 0 {
					if int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(i + 748))) + 1440))))&0x80 != 0 {
						nox_xxx_enactUnitOrder_5339A0(owner, i, orderType)
					}
				}
			}
		}
	}
}
func nox_xxx_enactUnitOrder_5339A0(source int32, unit int32, orderId int32) {
	var (
		v3        *uint32
		v4        int32
		v5        *uint32
		sfxIdle   int32
		v7        uint32
		sfxGuard  int32
		v9        int32
		v10       *int32
		v11       int32
		sfxEscort int32
		v13       uint32
		v14       *int32
		v15       int32
		sfxHunt   int32
		v17       uint32
	)
	v3 = *(**uint32)(unsafe.Pointer(uintptr(unit + 748)))
	if source != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(unit + 8))))&2 != 0 {
		if nox_xxx_unitIsZombie_534A40(unit) != 0 && orderId == 0 || (func() bool {
			v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(unit + 16))))
			return (v4 & 0x8000) == 0
		}()) {
			if *(*uint32)(unsafe.Add(unsafe.Pointer(v3), 4*121)) != 0 || (func() bool {
				v5 = (*uint32)(nox_xxx_monsterDefByTT_517560(int32(*(*uint16)(unsafe.Pointer(uintptr(unit + 4))))))
				return (func() uint32 {
					p := (*uint32)(unsafe.Add(unsafe.Pointer(v3), 4*121))
					*(*uint32)(unsafe.Add(unsafe.Pointer(v3), 4*121)) = uint32(uintptr(unsafe.Pointer(v5)))
					return *p
				}()) != 0
			}()) {
				switch orderId {
				case 0:
					if *(*uint32)(unsafe.Pointer(uintptr(unit + 508))) == uint32(source) {
						nox_xxx_banishUnit_5017F0(unit)
					}
				case 1:
					if int32(*(*uint8)(unsafe.Pointer(uintptr(source + 8))))&4 != 0 {
						nox_xxx_playerObserveMonster_4DDE80((*nox_object_t)(unsafe.Pointer(uintptr(source))), (*nox_object_t)(unsafe.Pointer(uintptr(unit))))
					}
				case 2:
					if int32(*(*uint8)(unsafe.Pointer(uintptr(source + 8))))&4 != 0 {
						nox_xxx_monsterCmdSend_528BD0(unit, source, internCStr("MonUtil.c:idle"), 0)
					}
					sfxIdle = int32(uintptr(nox_xxx_monsterGetSoundSet_424300((*nox_object_t)(unsafe.Pointer(uintptr(unit))))))
					if sfxIdle != 0 {
						nox_xxx_aud_501960(int32(*(*uint32)(unsafe.Pointer(uintptr(sfxIdle + 68)))), (*nox_object_t)(unsafe.Pointer(uintptr(unit))), 0, 0)
					}
					v7 = *(*uint32)(unsafe.Add(unsafe.Pointer(v3), 4*360)) & 0xFFFFFFBF
					*(*uint32)(unsafe.Add(unsafe.Pointer(v3), 4*326)) = 1056964608
					*(*uint32)(unsafe.Add(unsafe.Pointer(v3), 4*360)) = v7
					nox_xxx_monsterClearActionStack_50A3A0((*nox_object_t)(unsafe.Pointer(uintptr(unit))))
					nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(unit))), 0, internCStr(__FILE__), __LINE__)
				case 3:
					if nox_xxx_monsterIsMoveing_534320(unit) != 0 {
						if int32(*(*uint8)(unsafe.Pointer(uintptr(source + 8))))&4 != 0 {
							nox_xxx_monsterCmdSend_528BD0(unit, source, internCStr("MonUtil.c:guarding"), 0)
						}
						sfxGuard = int32(uintptr(nox_xxx_monsterGetSoundSet_424300((*nox_object_t)(unsafe.Pointer(uintptr(unit))))))
						if sfxGuard != 0 {
							nox_xxx_aud_501960(int32(*(*uint32)(unsafe.Pointer(uintptr(sfxGuard + 68)))), (*nox_object_t)(unsafe.Pointer(uintptr(unit))), 0, 0)
						}
						*(*uint32)(unsafe.Add(unsafe.Pointer(v3), 4*326)) = 1056964608
						if nox_xxx_monsterCanShoot_534280(unit) != 0 {
							v9 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), 4*360)))
							*((*uint8)(unsafe.Pointer(&v9))) = uint8(int8(v9 | 0x40))
							*(*uint32)(unsafe.Add(unsafe.Pointer(v3), 4*360)) = uint32(v9)
						}
						*(*uint32)(unsafe.Add(unsafe.Pointer(v3), 4*328)) = 1132068864
						nox_xxx_monsterClearActionStack_50A3A0((*nox_object_t)(unsafe.Pointer(uintptr(unit))))
						v10 = (*int32)(nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(unit))), 4, internCStr(__FILE__), __LINE__))
						if v10 != nil {
							*(*int32)(unsafe.Add(unsafe.Pointer(v10), 4*1)) = int32(*(*uint32)(unsafe.Pointer(uintptr(unit + 56))))
							*(*int32)(unsafe.Add(unsafe.Pointer(v10), 4*2)) = int32(*(*uint32)(unsafe.Pointer(uintptr(unit + 60))))
							*(*int32)(unsafe.Add(unsafe.Pointer(v10), 4*3)) = int32(*(*int16)(unsafe.Pointer(uintptr(unit + 124))))
						}
					}
				case 4:
					v11 = int32(*(*uint32)(unsafe.Pointer(uintptr(unit + 748))))
					if nox_xxx_monsterIsMoveing_534320(unit) != 0 {
						if int32(*(*uint8)(unsafe.Pointer(uintptr(source + 8))))&4 != 0 {
							nox_xxx_monsterCmdSend_528BD0(unit, source, internCStr("MonUtil.c:escorting"), 0)
						}
						sfxEscort = int32(uintptr(nox_xxx_monsterGetSoundSet_424300((*nox_object_t)(unsafe.Pointer(uintptr(unit))))))
						if sfxEscort != 0 {
							nox_xxx_aud_501960(int32(*(*uint32)(unsafe.Pointer(uintptr(sfxEscort + 68)))), (*nox_object_t)(unsafe.Pointer(uintptr(unit))), 0, 0)
						}
						v13 = *(*uint32)(unsafe.Pointer(uintptr(v11 + 1440))) & 0xFFFFFFBF
						*(*uint32)(unsafe.Pointer(uintptr(v11 + 1304))) = 1062501089
						*(*uint32)(unsafe.Pointer(uintptr(v11 + 1440))) = v13
						nox_xxx_monsterClearActionStack_50A3A0((*nox_object_t)(unsafe.Pointer(uintptr(unit))))
						v14 = (*int32)(nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(unit))), 3, internCStr(__FILE__), __LINE__))
						if v14 != nil {
							*(*int32)(unsafe.Add(unsafe.Pointer(v14), 4*1)) = int32(*(*uint32)(unsafe.Pointer(uintptr(source + 56))))
							v15 = int32(*(*uint32)(unsafe.Pointer(uintptr(source + 60))))
							*(*int32)(unsafe.Add(unsafe.Pointer(v14), 4*3)) = source
							*(*int32)(unsafe.Add(unsafe.Pointer(v14), 4*2)) = v15
						}
					}
				case 5:
					if nox_xxx_monsterIsMoveing_534320(unit) != 0 {
						sfxHunt = int32(uintptr(nox_xxx_monsterGetSoundSet_424300((*nox_object_t)(unsafe.Pointer(uintptr(unit))))))
						if sfxHunt != 0 {
							nox_xxx_aud_501960(int32(*(*uint32)(unsafe.Pointer(uintptr(sfxHunt + 68)))), (*nox_object_t)(unsafe.Pointer(uintptr(unit))), 0, 0)
						}
						if int32(*(*uint8)(unsafe.Pointer(uintptr(source + 8))))&4 != 0 {
							nox_xxx_monsterCmdSend_528BD0(unit, source, internCStr("MonUtil.c:Hunting"), 0)
						}
						v17 = *(*uint32)(unsafe.Add(unsafe.Pointer(v3), 4*360)) & 0xFFFFFFBF
						*(*uint32)(unsafe.Add(unsafe.Pointer(v3), 4*326)) = 1062501089
						*(*uint32)(unsafe.Add(unsafe.Pointer(v3), 4*360)) = v17
						nox_xxx_monsterClearActionStack_50A3A0((*nox_object_t)(unsafe.Pointer(uintptr(unit))))
						nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(unit))), 5, internCStr(__FILE__), __LINE__)
					}
				default:
					return
				}
			}
		}
	}
}
func nox_xxx_mobCalcDir_533CC0(a1 int32, a2 *float32) {
	var v2 float2
	if a1 != 0 {
		v2.field_0 = *a2 - *(*float32)(unsafe.Pointer(uintptr(a1 + 56)))
		v2.field_4 = *(*float32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(float32(0))*1)) - *(*float32)(unsafe.Pointer(uintptr(a1 + 60)))
		*(*uint16)(unsafe.Pointer(uintptr(a1 + 126))) = uint16(int16(nox_xxx_math_509ED0(&v2)))
	}
}
func nox_xxx_unitNPCActionToAnim_533D00(a1 int32) *uint8 {
	var (
		v1     int32
		v2     int32
		v3     int32
		v4     int32
		result *uint8
		v6     int32
	)
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 12))))&0x10 != 0 {
		v2 = int32(*(*byte)(unsafe.Pointer(uintptr(v1 + 544))) + 23)
		a1 = 0
		v6 = 0
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + v2*24))))
		switch v3 {
		case 16:
			fallthrough
		case 17:
			if *(*uint32)(unsafe.Pointer(uintptr(v1 + 2056)))&0xFFFFFFFC != 0 {
				v4 = sub_4FA280(int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 2056))) & 0xFFFFFFFC))
			} else {
				v4 = int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 2068))))
			}
			nox_xxx_animPlayerGetFrameRange_4F9F90(v4, &a1, &v6)
		case 18:
			fallthrough
		case 19:
			fallthrough
		case 20:
			nox_xxx_animPlayerGetFrameRange_4F9F90(21, &a1, &v6)
		case 30:
			nox_xxx_animPlayerGetFrameRange_4F9F90(1, &a1, &v6)
		case 31:
			nox_xxx_animPlayerGetFrameRange_4F9F90(2, &a1, &v6)
		case 21:
			nox_xxx_animPlayerGetFrameRange_4F9F90(40, &a1, &v6)
		case 23:
			if *(*uint32)(unsafe.Pointer(uintptr(v1 + 2056)))&0x7FF8000 != 0 {
				nox_xxx_animPlayerGetFrameRange_4F9F90(30, &a1, &v6)
			} else {
				nox_xxx_animPlayerGetFrameRange_4F9F90(47, &a1, &v6)
			}
		}
		*memmap.PtrUint32(0x5D4594, 2487968) = 0
		*memmap.PtrUint32(0x5D4594, 2487964) = 0
		*memmap.PtrUint32(0x5D4594, 2487972) = 0
		*memmap.PtrUint32(0x5D4594, 2487976) = 0
		*memmap.PtrUint8(0x5D4594, 2487973) = uint8(int8(a1))
		*memmap.PtrUint8(0x5D4594, 2487974) = uint8(int8(v6))
		result = (*uint8)(memmap.PtrOff(0x5D4594, 2487964))
	} else if *(*uint32)(unsafe.Pointer(uintptr(v1 + 476))) != 0 {
		result = (*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 476))) + uint32(nox_xxx_mobActionToAnimation_533790(a1)*16))))
	} else {
		result = nil
	}
	return result
}
func nox_xxx_monsterTestBlockShield_533E70(a1 *nox_object_t) int32 {
	*memmap.PtrUint32(0x5D4594, 2487956) = 0
	*memmap.PtrUint32(0x5D4594, 2487988) = 1315859240
	nox_xxx_getMissilesInCircle_518170((*float2)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))+56))), 100.0, unsafe.Pointer(funAddr(sub_533EB0)), a1)
	return int32(*memmap.PtrUint32(0x5D4594, 2487956))
}
func sub_533EB0(a1 int32, a2 int32) {
	var (
		v2  int32
		v3  float32
		v4  int32
		v5  float64
		v6  float64
		v7  *float32
		v8  float64
		v9  float64
		v10 float32
		v11 float32
		a3  float2
		v13 float32
		v14 float32
	)
	v2 = a1
	v3 = *(*float32)(unsafe.Pointer(uintptr(a1 + 76)))
	a3.field_0 = *(*float32)(unsafe.Pointer(uintptr(a1 + 72)))
	a3.field_4 = v3
	if sub_54E6F0(a2, a1) != 0 {
		if nox_server_testTwoPointsAndDirection_4E6E50((*float2)(unsafe.Pointer(uintptr(a2+56))), int32(*(*int16)(unsafe.Pointer(uintptr(a2 + 124)))), &a3)&1 != 0 {
			v4 = int32(*(*int16)(unsafe.Pointer(uintptr(a2 + 124))))
			v5 = float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 80))) * *mem_getFloatPtr(0x587000, uint32(v4*8)+194136))
			v6 = float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 84))) * *mem_getFloatPtr(0x587000, uint32(v4*8)+194140))
			v7 = mem_getFloatPtr(0x587000, uint32(v4*8)+194136)
			if v5+v6 < 0.0 {
				v10 = *(*float32)(unsafe.Pointer(uintptr(a1 + 56))) - *(*float32)(unsafe.Pointer(uintptr(a2 + 56)))
				v11 = *(*float32)(unsafe.Pointer(uintptr(a1 + 60))) - *(*float32)(unsafe.Pointer(uintptr(a2 + 60)))
				v8 = float64(*v7*v11 + -*(*float32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(float32(0))*1))*v10)
				if v8 < 0.0 {
					v8 = -v8
				}
				if v8 < 20.0 {
					v9 = math.Sqrt(float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 80)))**(*float32)(unsafe.Pointer(uintptr(a1 + 80))) + *(*float32)(unsafe.Pointer(uintptr(a1 + 84)))**(*float32)(unsafe.Pointer(uintptr(a1 + 84)))))
					v13 = float32(float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 80)))) / v9)
					v14 = float32(math.Sqrt(float64(v11*v11 + v10*v10)))
					if float64(*(*float32)(unsafe.Pointer(uintptr(v2 + 84))))/v9*float64(-(v11/v14))+float64(v13*(-(v10/v14))) > 0.69999999 {
						if nox_xxx_unitCanInteractWith_5370E0((*nox_object_t)(unsafe.Pointer(uintptr(a2))), (*nox_object_t)(unsafe.Pointer(uintptr(v2))), 0) != 0 {
							if float64(v14) < float64(*mem_getFloatPtr(0x5D4594, 2487988)) {
								*memmap.PtrUint32(0x5D4594, 2487956) = uint32(v2)
								*mem_getFloatPtr(0x5D4594, 2487988) = v14
							}
						}
					}
				}
			}
		}
	}
}
func sub_534020(a1 int32) int32 {
	return int32((*(*uint32)(unsafe.Pointer(uintptr(a1 + 12))) >> 10) & 1)
}
func nox_xxx_monsterMoveAudio_534030(a1 int32) int8 {
	var (
		v1  int32
		v2  int8
		v3  int32
		v4  uint32
		v5  uint32
		v6  uint32
		v7  int32
		v9  uint32
		v10 int32
	)
	v1 = a1
	v2 = int8(*(*uint8)(unsafe.Pointer(uintptr(a1 + 12))))
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	v10 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	if int32(v2)&0x30 != 0 {
		nox_xxx_animPlayerGetFrameRange_4F9F90(4, (*int32)(unsafe.Pointer(&v9)), &a1)
		v4 = *(*uint32)(unsafe.Pointer(uintptr(v1 + 36))) + gameFrame()
		v5 = v4 / uint32(a1+1) % v9
		v6 = (v4 - 1) / uint32(a1+1) / v9
		if v5 != (v4-1)/uint32(a1+1)%v9 {
			v6 = *(*uint32)(unsafe.Pointer(uintptr(v10 + 484)))
			if v5 == *(*uint32)(unsafe.Pointer(uintptr(v6 + 100))) || v5 == *(*uint32)(unsafe.Pointer(uintptr(v6 + 104))) {
				v6 = uint32(uintptr(nox_xxx_monsterGetSoundSet_424300((*nox_object_t)(unsafe.Pointer(uintptr(v1))))))
				if v6 != 0 {
					nox_xxx_aud_501960(int32(*(*uint32)(unsafe.Pointer(uintptr(v6 + 72)))), (*nox_object_t)(unsafe.Pointer(uintptr(v1))), 0, 0)
				}
			}
		}
	} else {
		v6 = *(*uint32)(unsafe.Pointer(uintptr(v3 + 484)))
		v7 = int32(*(*uint8)(unsafe.Pointer(uintptr(v3 + 481))))
		if uint32(v7) == *(*uint32)(unsafe.Pointer(uintptr(v6 + 100))) || uint32(v7) == *(*uint32)(unsafe.Pointer(uintptr(v6 + 104))) {
			*((*uint8)(unsafe.Pointer(&v6))) = *(*uint8)(unsafe.Pointer(uintptr(v3 + 482)))
			if int32(uint8(v6)) == 0 {
				v6 = uint32(uintptr(nox_xxx_monsterGetSoundSet_424300((*nox_object_t)(unsafe.Pointer(uintptr(a1))))))
				if v6 != 0 {
					nox_xxx_aud_501960(int32(*(*uint32)(unsafe.Pointer(uintptr(v6 + 72)))), (*nox_object_t)(unsafe.Pointer(uintptr(v1))), 0, 0)
				}
			}
		}
	}
	return int8(uint8(v6))
}
func sub_534120(a1 int32, a2 *float2) int32 {
	return bool2int32(float64(*mem_getFloatPtr(0x587000, uint32(int32(*(*int16)(unsafe.Pointer(uintptr(a1 + 124))))*8)+194140)*a2.field_4+*mem_getFloatPtr(0x587000, uint32(int32(*(*int16)(unsafe.Pointer(uintptr(a1 + 124))))*8)+194136)*a2.field_0) > 0.89999998)
}
func nox_ai_debug_printf_5341A0(a1 *byte, _rest ...interface{}) {
	var va libc.ArgList
	va.Start(a1, _rest)
	if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_ENABLE_SHOW_AI)) {
		nox_vsprintf((*byte)(memmap.PtrOff(0x5D4594, 2487996)), a1, va)
		nox_ai_debug_print((*byte)(memmap.PtrOff(0x5D4594, 2487996)))
	}
}
func sub_5341D0(a1 int32) int32 {
	nox_xxx_spellBuffOff_4FF5B0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0)
	return nox_xxx_spellBuffOff_4FF5B0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 23)
}
func nox_xxx_monsterCanMelee_534220(a1 int32) int32 {
	var v1 int32
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	if nox_xxx_monsterCanCast_534300((*nox_object_t)(unsafe.Pointer(uintptr(a1)))) != 0 {
		return 0
	}
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 12))))&0x10 != 0 {
		return bool2int32(nox_xxx_monsterCanShoot_534280(a1) == 0)
	}
	if float64(*(*float32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 484))) + 112)))) <= 0.0 {
		return 0
	} else {
		return 1
	}
}
func nox_xxx_monsterCanShoot_534280(a1 int32) int32 {
	var (
		v1     int32
		result int32
	)
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 12))))&0x10 != 0 {
		result = bool2int32((*(*uint32)(unsafe.Pointer(uintptr(v1 + 2056))) & 0x47F00FE) != 0)
	} else {
		result = bool2int32(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 484))) + 148)))) != 0)
	}
	return result
}
func nox_xxx_monsterHasShield_5342C0(a1 int32) int32 {
	var (
		v1     int32
		result int32
	)
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 12))))&0x10 != 0 {
		result = bool2int32((*(*uint32)(unsafe.Pointer(uintptr(v1 + 2060))) & 0x3000000) != 0)
	} else {
		result = int32((*(*uint32)(unsafe.Pointer(uintptr(v1 + 1440))) >> 2) & 1)
	}
	return result
}
func nox_xxx_monsterCanCast_534300(a1 *nox_object_t) int32 {
	return int32((*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1))) + 748))) + 1440))) >> 5) & 1)
}
func nox_xxx_monsterIsMoveing_534320(a1 int32) int32 {
	return bool2int32(float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 548)))) >= 0.0099999998)
}
func sub_534340(a1 int32) int32 {
	var v1 int32
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + uint32((*(*byte)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + 544)))+23)*24)))))
	return bool2int32(v1 == 0 || v1 == 1 || v1 == 4 || v1 == 25 || v1 == 26 || v1 == 27 || v1 == 23)
}
func nox_xxx_monsterCanAttackAtWill_534390(a1 *nox_object_t) int32 {
	return bool2int32(float64(*(*float32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1))) + 748))) + 1304)))) > 0.66000003)
}
func sub_5343C0(a1 int32) int32 {
	var v1 int32
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	return bool2int32(float64(*(*float32)(unsafe.Pointer(uintptr(v1 + 1304)))) < 0.66000003 && float64(*(*float32)(unsafe.Pointer(uintptr(v1 + 1304)))) > 0.33000001)
}
func sub_534400(a1 int32) int32 {
	var v1 int32
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	return bool2int32(float64(*(*float32)(unsafe.Pointer(uintptr(v1 + 1304)))) < 0.33000001 && float64(*(*float32)(unsafe.Pointer(uintptr(v1 + 1304)))) > 0.079999998)
}
func sub_534440(a1 int32) int32 {
	return bool2int32(float64(*(*float32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + 1304)))) < 0.079999998)
}
func sub_534470(a1 int32) float64 {
	return float64(*(*float32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + 484))) + 112))))
}
func sub_5345B0(a1 int32) *byte {
	var (
		v1 *uint8
		v2 int32
	)
	v1 = (*uint8)(memmap.PtrOff(0x587000, 262056))
	for {
		if *(*uint32)(unsafe.Pointer(v1)) == uint32(a1) {
			v2 = 0
			for {
				if v2 == a1 {
					return *(**byte)(memmap.PtrOff(0x587000, uint32(v2*4)+261768))
				}
				if func() int32 {
					p := &v2
					*p++
					return *p
				}() >= 39 {
					goto LABEL_6
				}
			}
		}
	LABEL_6:
		v1 = (*uint8)(unsafe.Add(unsafe.Pointer(v1), 4))
		if int32(uintptr(unsafe.Pointer(v1))) >= int32(uintptr(memmap.PtrOff(0x587000, 262072))) {
			return *(**byte)(memmap.PtrOff(0x587000, 261920))
		}
	}
}
func nox_xxx_actionNByNameMB_5345F0(a1 *byte) int32 {
	var (
		v1 int32
		v2 **byte
	)
	v1 = 0
	v2 = (**byte)(memmap.PtrOff(0x587000, 261768))
	for libc.StrCmp(*v2, a1) != 0 {
		v2 = (**byte)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof((*byte)(nil))*1))
		v1++
		if int32(uintptr(unsafe.Pointer(v2))) >= int32(uintptr(memmap.PtrOff(0x587000, 261924))) {
			return 38
		}
	}
	return v1
}
func sub_534650(a1 int32) *byte {
	var v1 int32
	v1 = 0
	for v1 != a1 {
		if func() int32 {
			p := &v1
			*p++
			return *p
		}() >= 72 {
			return nil
		}
	}
	return *(**byte)(memmap.PtrOff(0x587000, uint32(v1*4)+261768))
}
func nox_xxx_actionByName_534670(a1 *byte) int32 {
	var (
		v1 int32
		v2 **byte
	)
	v1 = 0
	v2 = (**byte)(memmap.PtrOff(0x587000, 261768))
	for libc.StrCmp(*v2, a1) != 0 {
		v2 = (**byte)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof((*byte)(nil))*1))
		v1++
		if int32(uintptr(unsafe.Pointer(v2))) >= int32(uintptr(memmap.PtrOff(0x587000, 262056))) {
			return 0
		}
	}
	return v1
}
func sub_5346D0(a1 int32) int32 {
	var result int32
	result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	*(*uint32)(unsafe.Pointer(uintptr(result + 8))) = 0
	*(*uint32)(unsafe.Pointer(uintptr(result + 296))) = 0
	return result
}
func nox_xxx_monsterResetEnemy_5346F0(a1 int32) int32 {
	var result int32
	result = a1
	*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + 1196))) = 0
	return result
}
func sub_534710(a1 int32) int32 {
	var (
		v1     int32
		result int32
	)
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	if int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 1440))))&0x40 != 0 {
		result = 1
	} else {
		result = bool2int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + int32((*(*byte)(unsafe.Pointer(uintptr(v1 + 544)))+23)*24)))) == 4)
	}
	return result
}
func sub_534750(a1 int32) int32 {
	var (
		v1     int32
		result int32
	)
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	result = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 1440))))
	if (uint32(result) & 0x10000) == 0 {
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&result))), 1)) |= 0x40
		*(*uint32)(unsafe.Pointer(uintptr(v1 + 1440))) = uint32(result)
	}
	return result
}
func sub_534780(a1 int32) int32 {
	var (
		v1     int32
		result int32
	)
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	result = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 1440))))
	if (result & 0x8000) == 0 {
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&result))), 1)) &= 0xBF
		*(*uint32)(unsafe.Pointer(uintptr(v1 + 1440))) = uint32(result)
	}
	return result
}
func sub_5347A0(a1 *nox_object_t) int32 {
	return int32((*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1))) + 748))) + 1440))) >> 9) & 1)
}
func sub_5347C0(a1 int32) int32 {
	var (
		v1     *uint16
		v2     uint16
		result int32
	)
	v1 = *(**uint16)(unsafe.Pointer(uintptr(a1 + 556)))
	result = 0
	if v1 != nil {
		v2 = *(*uint16)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint16(0))*2))
		if int32(*v1) < int32(v2) {
			if int32(v2) != 0 {
				result = 1
			}
		}
	}
	return result
}
func nox_xxx_isNotPoisoned_5347F0(a1 int32) int32 {
	return bool2int32(int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 540)))) != 0)
}
func nox_xxx_mobGetMoveAttemptTime_534810(a1 *nox_object_t) int32 {
	return bool2int32(gameFrame()-*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1))) + 748))) + 508))) < (gameFPS() * 3))
}
func nox_xxx_unitIsMimic_534840(a1 int32) int32 {
	var v1 int32
	v1 = int32(*memmap.PtrUint32(0x5D4594, 2488524))
	if *memmap.PtrUint32(0x5D4594, 2488524) == 0 {
		v1 = nox_xxx_getNameId_4E3AA0(internCStr("Mimic"))
		*memmap.PtrUint32(0x5D4594, 2488524) = uint32(v1)
	}
	return bool2int32(int32(*(*uint16)(unsafe.Pointer(uintptr(a1 + 4)))) == v1)
}
func nox_xxx_monsterMimicCheckMorph_534950(a1p *nox_object_t) {
	var (
		a1 int32 = int32(uintptr(unsafe.Pointer(a1p)))
		v1 int32
		v2 int32
		v3 int32
		v4 float64
		v5 float64
	)
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + int32((*(*byte)(unsafe.Pointer(uintptr(v1 + 544)))+23)*24)))))
	v3 = v1 + int32((*(*byte)(unsafe.Pointer(uintptr(v1 + 544)))+23)*24)
	if v2 != 0 && (v2 != 4 || (func() bool {
		v4 = float64(*(*float32)(unsafe.Pointer(uintptr(v3 + 4))) - *(*float32)(unsafe.Pointer(uintptr(a1 + 56))))
		v5 = float64(*(*float32)(unsafe.Pointer(uintptr(v3 + 8))) - *(*float32)(unsafe.Pointer(uintptr(a1 + 60))))
		return v5*v5+v4*v4 > 64.0
	}())) {
		if v2 != 34 && *(*uint32)(unsafe.Pointer(uintptr(v1 + 1440)))&0x40000 != 0 {
			nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 61, internCStr(__FILE__), __LINE__)
			nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 34, internCStr(__FILE__), __LINE__)
			nox_xxx_aud_501960(460, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
			return
		}
	} else if (*(*uint32)(unsafe.Pointer(uintptr(v1 + 1440)))&0x40000) == 0 && (gameFrame()-*(*uint32)(unsafe.Pointer(uintptr(v1 + 548)))) > uint32(int32(gameFPS())) {
		nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 61, internCStr(__FILE__), __LINE__)
		nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 33, internCStr(__FILE__), __LINE__)
		nox_xxx_aud_501960(460, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
		return
	}
}
func nox_xxx_unitIsPlant_534A10(a1 int32) int32 {
	var v1 int32
	v1 = int32(*memmap.PtrUint32(0x5D4594, 2488528))
	if *memmap.PtrUint32(0x5D4594, 2488528) == 0 {
		v1 = nox_xxx_getNameId_4E3AA0(internCStr("CarnivorousPlant"))
		*memmap.PtrUint32(0x5D4594, 2488528) = uint32(v1)
	}
	return bool2int32(int32(*(*uint16)(unsafe.Pointer(uintptr(a1 + 4)))) == v1)
}
func nox_xxx_unitIsZombie_534A40(a1 int32) int32 {
	var v1 int32
	if *memmap.PtrUint32(0x5D4594, 2488532) == 0 {
		*memmap.PtrUint32(0x5D4594, 2488532) = uint32(nox_xxx_getNameId_4E3AA0(internCStr("Zombie")))
		*memmap.PtrUint32(0x5D4594, 2488536) = uint32(nox_xxx_getNameId_4E3AA0(internCStr("VileZombie")))
	}
	v1 = int32(*(*uint16)(unsafe.Pointer(uintptr(a1 + 4))))
	return bool2int32(uint32(uint16(int16(v1))) == *memmap.PtrUint32(0x5D4594, 2488532) || uint32(v1) == *memmap.PtrUint32(0x5D4594, 2488536))
}
func nox_xxx_mobActionGetUp_534A90(a1 int32) int8 {
	var result int8
	result = int8(a1)
	if int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + 483)))) != 0 {
		result = int8(nox_xxx_monsterPopAction_50A160((*nox_object_t)(unsafe.Pointer(uintptr(a1)))))
	}
	return result
}
func nox_xxx_mobRaiseZombie_534AB0(a1 int32) uint32 {
	var result uint32
	result = uint32(nox_xxx_unitIsZombie_534A40(a1))
	if result != 0 {
		result = uint32(nox_xxx_mobActionGet_50A020(a1))
		if result == 31 {
			nox_xxx_monsterPopAction_50A160((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
			nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 61, internCStr(__FILE__), __LINE__)
			nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 35, internCStr(__FILE__), __LINE__)
			nox_xxx_aud_501960(469, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
			nox_xxx_unitHPsetOnMax_4EE6F0(a1)
			result = *(*uint32)(unsafe.Pointer(uintptr(a1 + 16))) & 0xFFFF7FA7
			*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))) = result
		}
	}
	return result
}
func nox_xxx_damageToMap_534BC0(a1 int32, a2 int32, a3 int32, a4 int32, a5 *nox_object_t) int32 {
	var (
		v5     *uint8
		v6     uint8
		v7     int32
		v8     int32
		result int32
		v10    uint8
		v11    int32
		v12    int32
		v13    int32
		v14    *byte
		v15    *uint32
		v16    float32
		v17    int2
	)
	v5 = (*uint8)(nox_server_getWallAtGrid_410580(a1, a2))
	if *memmap.PtrUint32(0x5D4594, 2488556) == 0 {
		*memmap.PtrUint32(0x5D4594, 2488556) = uint32(nox_xxx_wallTileByName_410D60(internCStr("MagicWallSystemUseOnly")))
	}
	if v5 == nil {
		return 0
	}
	v6 = *(*uint8)(unsafe.Add(unsafe.Pointer(v5), 4))
	if int32(v6)&0x20 != 0 {
		return 0
	}
	if (int32(v6) & 8) == 0 {
		sub_532FE0(nox_xxx_wallField36(int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v5), 1)))), int32(uintptr(unsafe.Pointer(a5))))
		return 0
	}
	if !nox_common_gameFlags_check_40A5C0(4096) || a5 == nil || uint32(*(*uint8)(unsafe.Add(unsafe.Pointer(v5), 1))) != *memmap.PtrUint32(0x5D4594, 2488556) || (func() int32 {
		v7 = int32(uintptr(unsafe.Pointer(nox_xxx_findParentChainPlayer_4EC580(a5))))
		return v7
	}()) == 0 || (func() bool {
		v8 = 99999999
		return (int32(*(*uint8)(unsafe.Pointer(uintptr(v7 + 8)))) & 4) == 0
	}()) {
		v8 = a3
	}
	if int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v5), 7)))-v8 > 0 {
		v10 = uint8(int8(int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v5), 7))) - v8))
		v11 = int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v5), 1)))
		*(*uint8)(unsafe.Add(unsafe.Pointer(v5), 7)) = v10
		*(*float32)(unsafe.Pointer(&v17.field_0)) = float32(float64(a1)*23.0 + 11.5)
		*(*float32)(unsafe.Pointer(&v17.field_4)) = float32(float64(a2)*23.0 + 11.5)
		v12 = int32(nox_xxx_wallGetBrickTypeMB_410E40(v11))
		if v12 != 0 {
			v13 = nox_common_randomInt_415FA0(0, v12-1)
			v14 = nox_xxx_wallGetBrickObj_410E60(int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v5), 1))), v13)
			v15 = (*uint32)(unsafe.Pointer(nox_xxx_newObjectByTypeID_4E3810(v14)))
			if v15 != nil {
				nox_xxx_createAt_4DAA50((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v15)))))), nil, *(*float32)(unsafe.Pointer(&v17.field_0)), *(*float32)(unsafe.Pointer(&v17.field_4)))
				v16 = float32(nox_common_randomFloat_416030(10.0, 20.0))
				nox_xxx_objectApplyForce_52DF80((*float32)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(&v17)))))), (*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v15)))))), v16)
			}
		}
		sub_533010(nox_xxx_wallField36(int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v5), 1)))), int32(uintptr(unsafe.Pointer(a5))))
		result = 0
	} else {
		v17.field_0 = a1
		v17.field_4 = a2
		result = nox_xxx_wallPreDestroy_534DA0(&v17.field_0)
	}
	return result
}
func nox_xxx_wallPreDestroy_534DA0(a1 *int32) int32 {
	var (
		v1  int32
		v2  int8
		v3  float64
		v4  *byte
		v5  int32
		v6  int32
		v7  int32
		v8  int32
		v9  int32
		v10 *byte
		v11 *uint32
		v13 int32
		v14 float32
		v15 float32
		v16 float32
		v17 float2
	)
	v1 = int32(uintptr(nox_server_getWallAtGrid_410580(*a1, *(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*1)))))
	if *memmap.PtrUint32(0x5D4594, 2488560) == 0 {
		*memmap.PtrUint32(0x5D4594, 2488560) = uint32(nox_xxx_wallTileByName_410D60(internCStr("MagicWallSystemUseOnly")))
	}
	if v1 == 0 {
		return 0
	}
	v2 = int8(*(*uint8)(unsafe.Pointer(uintptr(v1 + 4))))
	if (int32(v2)&8) == 0 || int32(v2)&0x20 != 0 {
		return 0
	}
	v13 = int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 1))))
	v17.field_0 = float32(float64(*a1)*23.0 + 11.5)
	v3 = float64(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*1)))
	*(*uint8)(unsafe.Pointer(uintptr(v1 + 7))) = 0
	v17.field_4 = float32(v3*23.0 + 11.5)
	v4 = nox_xxx_wallSoundByTile_410EA0(v13)
	v5 = nox_xxx_utilFindSound_40AF50(v4)
	nox_xxx_audCreate_501A30(v5, &v17, 0, 0)
	nox_xxx_netSendPointFx_522FF0(-118, &v17)
	if !nox_common_gameFlags_check_40A5C0(4096) {
		v6 = int32(nox_xxx_wallGetBrickTypeMB_410E40(int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 1))))))
		if v6 != 0 {
			v7 = nox_common_randomInt_415FA0(0, v6-1)
			v8 = nox_common_randomInt_415FA0(3, 6)
			if v8 > 0 {
				v9 = v8
				for {
					v10 = nox_xxx_wallGetBrickObj_410E60(int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 1)))), v7)
					v11 = (*uint32)(unsafe.Pointer(nox_xxx_newObjectByTypeID_4E3810(v10)))
					if v11 != nil {
						v15 = float32(nox_common_randomFloat_416030(-2.0, 2.0) + float64(v17.field_4))
						v14 = float32(nox_common_randomFloat_416030(-2.0, 2.0) + float64(v17.field_0))
						nox_xxx_createAt_4DAA50((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v11)))))), nil, v14, v15)
						v16 = float32(nox_common_randomFloat_416030(4.0, 10.0))
						nox_xxx_objectApplyForce_52DF80((*float32)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(&v17)))))), (*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v11)))))), v16)
					}
					if func() int32 {
						p := &v7
						*p++
						return *p
					}() >= int32(nox_xxx_wallGetBrickTypeMB_410E40(int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 1)))))) {
						v7 = 0
					}
					v9--
					if v9 == 0 {
						break
					}
				}
			}
		}
	}
	if uint32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 1)))) == *memmap.PtrUint32(0x5D4594, 2488560) {
		nox_xxx_wallDestroyMagicwallSysuse_4FF840(v1)
	} else {
		nox_xxx_wallDestroyedByWallid_410520(int16(*(*uint16)(unsafe.Pointer(uintptr(v1 + 10)))))
		nox_xxx_wallSendDestroyed_4DF0A0(v1, 32)
	}
	return 1
}
func nox_xxx_mapDamageToWalls_534FC0(a1 *int4, a2p unsafe.Pointer, a3 float32, a4 int32, a5 int32, a6p unsafe.Pointer) bool {
	var (
		a2  int32 = int32(uintptr(a2p))
		a6  int32 = int32(uintptr(a6p))
		v6  *int4
		v7  int32
		v8  int32
		v9  int32
		v10 int32
		v11 int32
		v12 int32
		v13 int32
		v14 int32
		v15 *byte
		v16 *float32
		v17 float64
		v18 float64
		v19 float64
		v20 float64
		v21 float64
		v22 float32
		v23 *int32
		v25 int32
		v26 float32
		v27 float32
		v28 float32
		v29 int32
		v30 float32
		v31 float32
		v32 int32
		v33 int32
		v34 *float32
		v35 float32
		v36 [2]int32
		v37 float32
		v38 float4
		v41 [256]byte
	)
	v6 = a1
	v7 = 0
	v25 = 0
	if a1 == nil || a2 == 0 {
		return false
	}
	v8 = a6
	v29 = 0
	if a6 != 0 {
		v9 = int32(*(*uint32)(unsafe.Pointer(uintptr(a6 + 8))))
		if v9&6 != 0 {
			v10 = int32(*(*int16)(unsafe.Pointer(uintptr(a6 + 124))))
			v29 = 1
		} else {
			if (uint32(v9) & 0x1001000) == 0 {
				goto LABEL_10
			}
			v11 = int32(*(*uint32)(unsafe.Pointer(uintptr(a6 + 492))))
			if v11 == 0 {
				goto LABEL_10
			}
			v10 = int32(*(*int16)(unsafe.Pointer(uintptr(v11 + 124))))
			v29 = 1
		}
		v34 = mem_getFloatPtr(0x587000, uint32(v10*8)+194136)
	}
LABEL_10:
	v12 = a1.field_C
	v13 = a1.field_4
	v33 = a1.field_4
	v37 = a3 * a3
	if v13 <= v12 {
		for {
			v14 = v6.field_0
			v32 = v6.field_0
			if v6.field_0 <= v6.field_8 {
				v15 = &v41[v7*8+4]
				for {
					if nox_server_getWallAtGrid_410580(v14, v13) != nil {
						v16 = (*float32)(unsafe.Pointer(uintptr(func() int32 {
							if v29 != 0 {
								return a6 + 56
							}
							return a2
						}())))
						v17 = float64(v32)
						v35 = float32(v17)
						v26 = float32(v17*23.0 + 11.5 - float64(*v16))
						v18 = float64(v33)
						v31 = float32(v18)
						v19 = v18*23.0 + 11.5 - float64(*(*float32)(unsafe.Add(unsafe.Pointer(v16), unsafe.Sizeof(float32(0))*1)))
						v28 = float32(v19)
						v20 = v19*float64(v28) + float64(v26*v26)
						v30 = float32(v20)
						if v20 < float64(v37) {
							if v29 == 0 || float64(v30) <= 0.0 || (func() bool {
								v21 = math.Sqrt(float64(v30)) + 0.5
								v27 = float32(float64(v26) / v21)
								return float64(v28)/v21*float64(*(*float32)(unsafe.Add(unsafe.Pointer(v34), unsafe.Sizeof(float32(0))*1)))+float64(v27**v34) > *(*float64)(unsafe.Pointer(&qword_581450_9544))
							}()) {
								v22 = *(*float32)(unsafe.Pointer(uintptr(a2 + 4)))
								v38.field_0 = *(*float32)(unsafe.Pointer(uintptr(a2)))
								v38.field_4 = v22
								v38.field_8 = float32(float64(v35)*23.0 + 11.5)
								v36[1] = 0
								v36[0] = 0
								v38.field_C = float32(float64(v31)*23.0 + 11.5)
								if (nox_xxx_mapTraceRay_535250(&v38, nil, (*int2)(unsafe.Pointer(&v36[0])), 1) == 1 || v36[0] == v14 && v36[1] == v13) && v25 < 32 {
									*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v15))), -int(4*1)))) = uint32(v14)
									*(*uint32)(unsafe.Pointer(v15)) = uint32(v13)
									v25++
									v15 = (*byte)(unsafe.Add(unsafe.Pointer(v15), 8))
								}
							}
						}
					}
					v32 = func() int32 {
						p := &v14
						*p++
						return *p
					}()
					if v14 > a1.field_8 {
						break
					}
				}
				v7 = v25
				v8 = a6
				v6 = a1
			}
			v33 = func() int32 {
				p := &v13
				*p++
				return *p
			}()
			if v13 > v6.field_C {
				break
			}
		}
	}
	if v7 > 0 {
		v23 = (*int32)(unsafe.Pointer(&v41[0]))
		for {
			nox_xxx_damageToMap_534BC0(*v23, *(*int32)(unsafe.Add(unsafe.Pointer(v23), 4*1)), a4, a5, (*nox_object_t)(unsafe.Pointer(uintptr(v8))))
			v23 = (*int32)(unsafe.Add(unsafe.Pointer(v23), 4*2))
			if func() int32 {
				p := &v7
				*p--
				return *p
			}() == 0 {
				break
			}
			v8 = a6
		}
		v7 = v25
	}
	return v7 > 0
}
func sub_536130(a1 *byte, a2 *int32) *byte {
	var (
		result *byte
		v3     int32
		v4     *byte
	)
	result = libc.StrTok(a1, internCStr(" "))
	if result != nil {
		v3 = nox_xxx_enchantByName_424880(result)
		*a2 = v3
		if v3 == -1 {
			result = nil
		} else {
			v4 = libc.StrTok(nil, internCStr(" "))
			*(*int32)(unsafe.Add(unsafe.Pointer(a2), 4*1)) = int32(libc.Atoi(libc.GoString(v4)))
			result = (*byte)(unsafe.Pointer(uintptr(1)))
		}
	}
	return result
}
func sub_536180(a1 *byte, a2 *int32) *byte {
	var result *byte
	result = libc.StrTok(a1, internCStr(" "))
	if result != nil {
		*a2 = nox_xxx_spellNameToN_4243F0(result)
		result = (*byte)(unsafe.Pointer(uintptr(1)))
	}
	return result
}
func sub_5361B0(a1 *byte, a2 int32) *byte {
	var (
		result *byte
		v3     int8
		v4     float64
	)
	*(*uint32)(unsafe.Pointer(uintptr(a2))) = 1
	result = libc.StrTok(a1, internCStr(" "))
	if result != nil {
		v3 = int8(libc.Atoi(libc.GoString(result)))
		*(*uint8)(unsafe.Pointer(uintptr(a2 + 109))) = uint8(v3)
		*(*uint8)(unsafe.Pointer(uintptr(a2 + 108))) = uint8(v3)
		*(*uint32)(unsafe.Pointer(uintptr(a2 + 112))) = 100
		result = libc.StrTok(nil, internCStr(" "))
		if result != nil {
			v4 = float64(gameFPS())
			*(*uint32)(unsafe.Pointer(uintptr(a2 + 100))) = uint32(int32(int64(v4 / libc.Atof(libc.GoString(result)))))
			result = libc.StrTok(nil, internCStr(" "))
			if result != nil {
				*(*uint32)(unsafe.Pointer(uintptr(a2 + 92))) = uint32(nox_xxx_spellNameToN_4243F0(result))
				result = (*byte)(unsafe.Pointer(uintptr(1)))
			}
		}
	}
	return result
}
func sub_536260(a1 *byte, a2 int32) *byte {
	var (
		result *byte
		v3     int8
		v4     *byte
		v5     int32
		v6     *byte
		v7     float64
		v8     float64
	)
	*(*uint32)(unsafe.Pointer(uintptr(a2))) = 0
	result = libc.StrTok(a1, internCStr(" "))
	if result != nil {
		v3 = int8(libc.Atoi(libc.GoString(result)))
		*(*uint8)(unsafe.Pointer(uintptr(a2 + 109))) = uint8(v3)
		*(*uint8)(unsafe.Pointer(uintptr(a2 + 108))) = uint8(v3)
		*(*uint32)(unsafe.Pointer(uintptr(a2 + 112))) = 100
		result = libc.StrTok(nil, internCStr(" "))
		if result != nil {
			libc.StrCpy((*byte)(unsafe.Pointer(uintptr(a2+4))), result)
			*(*uint32)(unsafe.Pointer(uintptr(a2 + 84))) = 0
			result = libc.StrTok(nil, internCStr(" "))
			if result != nil {
				v7 = float64(gameFPS())
				v8 = libc.Atof(libc.GoString(result))
				if v8 == 0.0 {
					*(*uint32)(unsafe.Pointer(uintptr(a2 + 100))) = 0
				} else {
					*(*uint32)(unsafe.Pointer(uintptr(a2 + 100))) = uint32(int32(int64(v7 / v8)))
				}
				v4 = libc.StrTok(nil, internCStr(" "))
				if v4 != nil && libc.StrCmp(v4, internCStr("MULTI_SHOT")) == 0 {
					v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 96))))
					*((*uint8)(unsafe.Pointer(&v5))) = uint8(int8(v5 | 1))
					*(*uint32)(unsafe.Pointer(uintptr(a2 + 96))) = uint32(v5)
				}
				v6 = libc.StrTok(nil, internCStr(" "))
				if v6 != nil {
					*(*uint32)(unsafe.Pointer(uintptr(a2 + 88))) = uint32(nox_xxx_utilFindSound_40AF50(v6))
				}
				result = (*byte)(unsafe.Pointer(uintptr(1)))
			}
		}
	}
	return result
}
func sub_536390(a1 *byte, a2 *int32) *byte {
	var result *byte
	result = libc.StrTok(a1, internCStr(" "))
	if result != nil {
		*a2 = int32(libc.Atoi(libc.GoString(result)))
		result = (*byte)(unsafe.Pointer(uintptr(1)))
	}
	return result
}
func sub_5363C0(a1 *byte, a2 *int32) *byte {
	var result *byte
	result = libc.StrTok(a1, internCStr(" "))
	if result != nil {
		*a2 = int32(libc.Atoi(libc.GoString(result)))
		result = (*byte)(unsafe.Pointer(uintptr(1)))
	}
	return result
}
func sub_5364E0(a1 *byte, a2 int32) int32 {
	var (
		v2     uint32
		v3     int8
		v4     *byte
		v5     *byte
		result int32
		v7     [256]byte
	)
	stdio.Sscanf(a1, "%s", &v7[0])
	v2 = uint32(libc.StrLen(&v7[0]) + 1)
	v3 = int8(uint8(v2))
	v2 >>= 2
	alloc.Memcpy(unsafe.Pointer(uintptr(a2+16)), unsafe.Pointer(&v7[0]), int(v2*4))
	v5 = &v7[v2*4]
	v4 = (*byte)(unsafe.Pointer(uintptr(uint32(a2+16) + v2*4)))
	*((*uint8)(unsafe.Pointer(&v2))) = uint8(v3)
	result = 1
	alloc.Memcpy(unsafe.Pointer(v4), unsafe.Pointer(v5), int(v2&3))
	*(*uint32)(unsafe.Pointer(uintptr(a2 + 12))) = 0
	return result
}
func sub_536550(a1 *byte, a2 *uint32) int32 {
	stdio.Sscanf(a1, "%f %f", a2, (*uint32)(unsafe.Add(unsafe.Pointer(a2), 4*2)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(a2), 4*1)) = *a2
	return 1
}
func sub_536580(a1 *byte, a2 int32) int32 {
	stdio.Sscanf(a1, "%d %d %d", a2, a2+4, a2+8)
	return 1
}
func sub_5365B0(a1 *byte, a2 int32) int32 {
	var (
		v2 *byte
		v3 *byte
	)
	v2 = libc.StrTok(a1, internCStr(" "))
	if v2 != nil {
		*(*uint32)(unsafe.Pointer(uintptr(a2 + 36))) = uint32(nox_xxx_utilFindSound_40AF50(v2))
	}
	v3 = libc.StrTok(nil, internCStr(" "))
	if v3 != nil {
		*(*uint32)(unsafe.Pointer(uintptr(a2 + 40))) = uint32(nox_xxx_utilFindSound_40AF50(v3))
	}
	return 1
}
func sub_536600(a1 *byte, a2 int32) int32 {
	stdio.Sscanf(a1, "%d", a2)
	return 1
}
func sub_536B40(a1 *byte, a2 int32) int32 {
	var v3 [64]byte
	stdio.Sscanf(a1, "%s %s", a2, &v3[0])
	*(*uint32)(unsafe.Pointer(uintptr(a2 + 128))) = uint32(nox_xxx_utilFindSound_40AF50(&v3[0]))
	return 1
}
func sub_536D80(a1 *byte, a2 int32) int32 {
	stdio.Sscanf(a1, "%d", a2)
	return 1
}
func sub_536DA0(a1 *byte, a2 *int32) int32 {
	var (
		v2 int32
		v4 [256]byte
	)
	stdio.Sscanf(a1, "%s", &v4[0])
	v2 = nox_xxx_utilFindSound_40AF50(&v4[0])
	*a2 = v2
	return bool2int32(v2 != 0)
}
func sub_536DE0(a1 *byte, a2 *uint8) int32 {
	stdio.Sscanf(a1, "%d", &a1)
	*a2 = uint8(uintptr(unsafe.Pointer(a1)))
	return 1
}
func nox_xxx_collideDamageLoad_536E10(a1 *byte, a2 int32) int32 {
	var (
		v2 *byte
		v3 *byte
		v4 int32
	)
	v2 = libc.StrTok(a1, internCStr(" "))
	*(*uint8)(unsafe.Pointer(uintptr(a2))) = uint8(int8(libc.Atoi(libc.GoString(v2))))
	v3 = libc.StrTok(nil, internCStr(" "))
	v4 = nox_xxx_parseDamageTypeByName_4E0A00(v3)
	*(*uint32)(unsafe.Pointer(uintptr(a2 + 4))) = uint32(v4)
	return bool2int32(v4 != 18)
}
func sub_536E50(a1 *byte, a2 *uint8) int32 {
	var v2 *byte
	v2 = libc.StrTok(a1, internCStr(" "))
	*a2 = uint8(int8(libc.Atoi(libc.GoString(v2))))
	return 1
}
func sub_536E80(a1 *byte, a2 *int32) int32 {
	var (
		v2 *byte
		v3 *byte
	)
	v2 = libc.StrTok(a1, internCStr(" "))
	*a2 = int32(libc.Atoi(libc.GoString(v2)))
	v3 = libc.StrTok(a1, internCStr(" "))
	*(*int32)(unsafe.Add(unsafe.Pointer(a2), 4*1)) = int32(libc.Atoi(libc.GoString(v3)))
	return 1
}
func nox_xxx_traceRay_5374B0(a1 *float4) int32 {
	return nox_xxx_mapTraceRay_535250(a1, nil, nil, 9)
}
func sub_537540(a1 int32) {
	var i *uint32
	if a1 != 0 {
		for i = (*uint32)(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0())); i != nil; i = (*uint32)(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(i))))))))) {
			if *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(i), 4*187)) + 132))) == uint32(a1) {
				nox_xxx_harpoonBreakForPlr_537520((*nox_object_t)(unsafe.Pointer(i)))
			}
		}
	}
}
func sub_537580(a1 int32) int32 {
	return int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 464)))) & 1
}
func sub_5375A0(a1 int32) {
	var (
		v1 int32
		v2 int32
		v3 int8
	)
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 464))))&1 != 0 {
		v1 = int32(dword_5d4594_2488604)
		v2 = 0
		if dword_5d4594_2488604 != 0 {
			for v1 != a1 {
				v2 = v1
				v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 460))))
				if v1 == 0 {
					return
				}
			}
			if v1 != 0 {
				if v2 != 0 {
					*(*uint32)(unsafe.Pointer(uintptr(v2 + 460))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 460)))
				} else {
					dword_5d4594_2488604 = *(*uint32)(unsafe.Pointer(uintptr(a1 + 460)))
				}
				if uint32(a1) == dword_5d4594_2488608 {
					dword_5d4594_2488608 = uint32(v2)
				}
				v3 = int8(*(*uint8)(unsafe.Pointer(uintptr(a1 + 464))))
				*(*uint32)(unsafe.Pointer(uintptr(a1 + 460))) = math.MaxUint32
				*(*uint8)(unsafe.Pointer(uintptr(a1 + 464))) = uint8(int8(int32(v3) & 0xFE))
			}
		}
	}
}
func nox_xxx_unitHasCollideOrUpdateFn_537610(a1p *nox_object_t) int8 {
	var (
		a1 int32 = int32(uintptr(unsafe.Pointer(a1p)))
		v1 int32
		v2 int32
		v3 int32
		v4 func(int32, int32)
		v5 int32
	)
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 744))))
	if v1 != 0 || (func() int32 {
		v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 696))))
		return v1
	}()) != 0 && (int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 16))))&0x40) == 0 {
		if (func() bool {
			v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))))
			return (uint32(v2) & 0x400000) == 0
		}()) && (int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 16))))&8) == 0 || (func() int32 {
			v3 = nox_xxx_getNameId_4E3AA0(internCStr("Spike"))
			v1 = nox_xxx_getNameId_4E3AA0(internCStr("PeriodicSpike"))
			v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))))
			return v2 & 0xE080
		}()) != 0 || (func() bool {
			v4 = *(*func(int32, int32))(unsafe.Pointer(uintptr(a1 + 696)))
			return funAddr(v4) == funAddr(nox_xxx_collideFist_4EADF0)
		}()) || funAddr(v4) == funAddr(nox_xxx_collideUndeadKiller_4EBD40) || (func() bool {
			v5 = int32(*(*uint16)(unsafe.Pointer(uintptr(a1 + 4))))
			return int32(uint16(int16(v5))) == v3
		}()) || v5 == v1 {
			if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 16))))&4 != 0 {
				if v2&0x2008 != 0 {
					sub_50B500()
				}
				nullsub_30(uint32(a1))
				*((*uint8)(unsafe.Pointer(&v1))) = *(*uint8)(unsafe.Pointer(uintptr(a1 + 464)))
				if (v1 & 1) == 0 {
					if dword_5d4594_2488608 != 0 {
						*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_2488608 + 460))) = uint32(a1)
					} else {
						dword_5d4594_2488604 = uint32(a1)
					}
					dword_5d4594_2488608 = uint32(a1)
					*((*uint8)(unsafe.Pointer(&v1))) = uint8(int8(int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 464)))) | 1))
					*(*uint32)(unsafe.Pointer(uintptr(a1 + 460))) = 0
					*(*uint8)(unsafe.Pointer(uintptr(a1 + 464))) = uint8(int8(v1))
				}
			}
		}
	}
	return int8(v1)
}
func sub_537700() *nox_object_t {
	var (
		result int32
		v1     *uint32
	)
	_ = v1
	result = int32(dword_5d4594_2488604)
	v1 = (*uint32)(unsafe.Pointer(uintptr(dword_5d4594_2488604 + 460)))
	dword_5d4594_2488604 = *(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_2488604 + 460)))
	if dword_5d4594_2488604 == 0 {
		dword_5d4594_2488608 = 0
	}
	*v1 = math.MaxUint32
	*(*uint8)(unsafe.Pointer(uintptr(result + 464))) &= 0xFE
	return (*nox_object_t)(unsafe.Pointer(uintptr(result)))
}
func sub_537740() int32 {
	return int32(dword_5d4594_2488604)
}
func sub_537750(a1 int32) int32 {
	var result int32
	result = a1
	if a1 != 0 {
		result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 460))))
	}
	return result
}
func sub_537760() uint32 {
	if dword_5d4594_2488620 != 0 {
		return uint32(uintptr(memmap.PtrOff(0x5D4594, 2488612)))
	}
	return 0
}
func sub_537770(a1p *nox_object_t) {
	var (
		a1 int32 = int32(uintptr(unsafe.Pointer(a1p)))
		v1 int32
		v3 int32
		v4 float2
	)
	*((*uint8)(unsafe.Pointer(&v1))) = *memmap.PtrUint8(0x5D4594, 2488624)
	if *memmap.PtrUint32(0x5D4594, 2488624) == 0 {
		*memmap.PtrUint32(0x5D4594, 2488624) = uint32(nox_xxx_getNameId_4E3AA0(internCStr("SmallFist")))
		*memmap.PtrUint32(0x5D4594, 2488628) = uint32(nox_xxx_getNameId_4E3AA0(internCStr("MediumFist")))
		v1 = nox_xxx_getNameId_4E3AA0(internCStr("LargeFist"))
		*memmap.PtrUint32(0x5D4594, 2488632) = uint32(v1)
	}
	if (int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 16)))) & 0x60) == 0 {
		dword_5d4594_2488620 = 0
		*((*uint8)(unsafe.Pointer(&v1))) = uint8(nox_xxx_projectileTraceHit_537850(a1, &v3, &v4))
		if int32(uint8(int8(v1))) != 0 {
			if v3 == 0 || (func() bool {
				v1 = int32(*(*uint16)(unsafe.Pointer(uintptr(v3 + 4))))
				return uint32(uint16(int16(v1))) != *memmap.PtrUint32(0x5D4594, 2488624)
			}()) && uint32(v1) != *memmap.PtrUint32(0x5D4594, 2488628) && uint32(v1) != *memmap.PtrUint32(0x5D4594, 2488632) {
				(*(*func(int32, int32, *float2))(unsafe.Pointer(uintptr(a1 + 696))))(a1, v3, &v4)
				*((*uint8)(unsafe.Pointer(&v1))) = uint8(int8(v3))
				dword_5d4594_2488620 = 0
				if v3 != 0 {
					v4.field_0 = -v4.field_0
					v4.field_4 = -v4.field_4
					*((*uint8)(unsafe.Pointer(&v1))) = uint8(int8((*(*func(int32, int32, *float2) int32)(unsafe.Pointer(uintptr(v3 + 696))))(v3, a1, &v4)))
				}
			}
		}
	}
}
func nox_xxx_projectileTraceHit_537850(a1 int32, a2 *int32, a3 *float2) int8 {
	var (
		v3  int32
		v4  int32
		v5  float64
		v6  float64
		v7  int32
		v8  float32
		v9  int32
		v10 float32
		v11 float32
		v12 float32
		v13 bool
		v14 int32
		v15 int32
		v16 float32
		v17 float64
		v18 int32
		v19 float64
		v20 float64
		v21 float64
		v22 float64
		v23 float64
		v25 float32
		v26 float32
		v27 float2
		v28 float2
		v29 [2]int32
		a2a float2
		v31 float2
		a3a int2
		a1a float4
		v34 int32
		v35 int32
	)
	v3 = a1
	v4 = 0
	*(*float32)(unsafe.Pointer(&v34)) = *(*float32)(unsafe.Pointer(uintptr(a1 + 64))) - *(*float32)(unsafe.Pointer(uintptr(a1 + 56)))
	v5 = float64(*(*float32)(unsafe.Pointer(uintptr(v3 + 68))) - *(*float32)(unsafe.Pointer(uintptr(v3 + 60))))
	v26 = float32(v5)
	v6 = v5*float64(v26) + float64(*(*float32)(unsafe.Pointer(&v34))**(*float32)(unsafe.Pointer(&v34)))
	if v6 <= 36.0 {
		v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 56))))
		v8 = *(*float32)(unsafe.Pointer(uintptr(v3 + 64)))
		v28.field_4 = *(*float32)(unsafe.Pointer(uintptr(v3 + 68)))
		v29[0] = v7
		v28.field_0 = v8
		v29[1] = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 60))))
		v9 = sub_54E810(v3, &v28, int32(uintptr(unsafe.Pointer(&v29[0]))))
		if v9 != 0 {
			v4 = v9
			v27.field_0 = *(*float32)(unsafe.Pointer(&v29[0])) - *(*float32)(unsafe.Pointer(uintptr(v9 + 56)))
			v27.field_4 = *(*float32)(unsafe.Pointer(&v29[1])) - *(*float32)(unsafe.Pointer(uintptr(v9 + 60)))
		}
	} else {
		v15 = nox_double2int(math.Sqrt(v6*0.027777778)) + 1
		v16 = *(*float32)(unsafe.Pointer(uintptr(v3 + 60)))
		v17 = float64(v15)
		v29[0] = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 56))))
		v18 = 0
		*(*float32)(unsafe.Pointer(&v29[1])) = v16
		*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v28.field_0))), 4*0)) = uint32(v29[0])
		v28.field_4 = v16
		*(*float32)(unsafe.Pointer(&v35)) = float32(float64(*(*float32)(unsafe.Pointer(&v34))) / v17)
		v27.field_0 = float32(float64(v26) / v17)
		if v15 > 0 {
			for {
				v28.field_0 = v28.field_0 + *(*float32)(unsafe.Pointer(&v35))
				v28.field_4 = v28.field_4 + v27.field_0
				v9 = sub_54E810(v3, &v28, int32(uintptr(unsafe.Pointer(&v29[0]))))
				if v9 != 0 {
					v4 = v9
					v27.field_0 = *(*float32)(unsafe.Pointer(&v29[0])) - *(*float32)(unsafe.Pointer(uintptr(v9 + 56)))
					v27.field_4 = *(*float32)(unsafe.Pointer(&v29[1])) - *(*float32)(unsafe.Pointer(uintptr(v9 + 60)))
					break
				}
				v18++
				*(*float2)(unsafe.Pointer(&v29[0])) = v28
				if v18 >= v15 {
					break
				}
			}
		}
	}
	v10 = *(*float32)(unsafe.Pointer(uintptr(v3 + 60)))
	v11 = *(*float32)(unsafe.Pointer(uintptr(v3 + 64)))
	a1a.field_0 = *(*float32)(unsafe.Pointer(uintptr(v3 + 56)))
	v12 = *(*float32)(unsafe.Pointer(uintptr(v3 + 68)))
	a1a.field_4 = v10
	a1a.field_8 = v11
	a1a.field_C = v12
	if nox_xxx_mapTraceRay_535250(&a1a, &a2a, &a3a, 5) != 0 {
		v13 = false
	} else {
		*(*int2)(memmap.PtrOff(0x5D4594, 2488612)) = a3a
		*(*float2)(unsafe.Pointer(&a1a.field_8)) = a2a
		dword_5d4594_2488620 = 1
		v13 = sub_57CDB0(&a3a, &a1a.field_0, &v31) != 0
		v14 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 60))))
		*(*uint32)(unsafe.Pointer(uintptr(v3 + 64))) = *(*uint32)(unsafe.Pointer(uintptr(v3 + 56)))
		*(*uint32)(unsafe.Pointer(uintptr(v3 + 68))) = uint32(v14)
	}
	if v4 != 0 {
		if !v13 {
			*a3 = v27
			*a2 = v4
			return 1
		}
		v19 = float64(*(*float32)(unsafe.Pointer(uintptr(v3 + 56))) - *(*float32)(unsafe.Pointer(uintptr(v4 + 56))))
		v20 = float64(*(*float32)(unsafe.Pointer(uintptr(v3 + 60))) - *(*float32)(unsafe.Pointer(uintptr(v4 + 60))))
		v21 = v20*v20 + v19*v19
		v22 = float64(*(*float32)(unsafe.Pointer(uintptr(v3 + 56))) - a2a.field_0)
		v23 = float64(*(*float32)(unsafe.Pointer(uintptr(v3 + 60))) - a2a.field_4)
		if v21 < v23*v23+v22*v22 {
			*a3 = v27
			*a2 = v4
			return 1
		}
		v25 = v31.field_4
		a3.field_0 = v31.field_0
		a3.field_4 = v25
		*a2 = 0
		return 1
	}
	if v13 {
		v25 = v31.field_4
		a3.field_0 = v31.field_0
		a3.field_4 = v25
		*a2 = 0
		return 1
	}
	return 0
}
func nox_xxx_sMakeScorch_537AF0(a1 *float32, a2 int32) {
	var (
		result *uint32
		v3     *uint32
		v4     int32
		v5     int32
		v6     int32
		v7     int32
	)
	if *memmap.PtrUint32(0x5D4594, 2488636) == 0 {
		nox_xxx_scorchInit_537BD0()
	}
	if a2 != 0 {
		if a2 == 1 {
			v6 = int32(*memmap.PtrUint32(0x587000, uint32(nox_common_randomInt_415FA0(0, 0)*8)+276836))
			result = (*uint32)(unsafe.Pointer(nox_xxx_newObjectWithTypeInd_4E3450(v6)))
		} else {
			result = (*uint32)(unsafe.Pointer(uintptr(a2 - 2)))
			if a2 != 2 {
				return
			}
			v5 = int32(*memmap.PtrUint32(0x587000, uint32(nox_common_randomInt_415FA0(0, 0)*8)+276844))
			result = (*uint32)(unsafe.Pointer(nox_xxx_newObjectWithTypeInd_4E3450(v5)))
		}
	} else {
		v7 = int32(*memmap.PtrUint32(0x587000, uint32(nox_common_randomInt_415FA0(0, 0)*8)+276828))
		result = (*uint32)(unsafe.Pointer(nox_xxx_newObjectWithTypeInd_4E3450(v7)))
	}
	v3 = result
	if result != nil {
		nox_xxx_createAt_4DAA50((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(result)))))), nil, *a1, *((*float32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(float32(0))*1))))
		if nox_common_gameFlags_check_40A5C0(4096) {
			v4 = nox_common_randomInt_415FA0(5, 8)
		} else {
			v4 = nox_common_randomInt_415FA0(10, 20)
		}
		nox_xxx_unitSetDecayTime_511660((*nox_object_t)(unsafe.Pointer(v3)), int32(gameFPS()*uint32(v4)))
	}
}
func nox_xxx_scorchInit_537BD0() int32 {
	var result int32
	*memmap.PtrUint32(0x587000, 276828) = uint32(nox_xxx_getNameId_4E3AA0(*(**byte)(memmap.PtrOff(0x587000, 276824))))
	*memmap.PtrUint32(0x587000, 276836) = uint32(nox_xxx_getNameId_4E3AA0(*(**byte)(memmap.PtrOff(0x587000, 276832))))
	result = nox_xxx_getNameId_4E3AA0(*(**byte)(memmap.PtrOff(0x587000, 276840)))
	*memmap.PtrUint32(0x587000, 276844) = uint32(result)
	*memmap.PtrUint32(0x5D4594, 2488636) = 1
	return result
}
func nox_xxx_playerPreAttackEffects_538290(a1 int32, a2 int32, a3 int32, a4 int32) int32 {
	var (
		result int32
		v5     int32
		v6     int32
		v7     *int32
		v8     int32
		v9     func(int32, int32, int32, int32, int32)
		v10    int32
	)
	result = a3
	if a3 != 0 {
		v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a3 + 692))))
		v6 = a1
		if nox_xxx_CheckGameplayFlags_417DA0(1) || a2 == 0 || (int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 8))))&6) == 0 || (func() int32 {
			result = nox_xxx_unitIsEnemyTo_5330C0((*nox_object_t)(unsafe.Pointer(uintptr(a2))), (*nox_object_t)(unsafe.Pointer(uintptr(a1))))
			return result
		}()) != 0 {
			result = nox_xxx_testUnitBuffs_4FF350((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 23)
			if result == 0 {
				result = nox_xxx_testUnitBuffs_4FF350((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 27)
				if result == 0 {
					v7 = (*int32)(unsafe.Pointer(uintptr(v5 + 8)))
					v10 = 2
					for {
						v8 = *v7
						if *v7 != 0 {
							v9 = *(*func(int32, int32, int32, int32, int32))(unsafe.Pointer(uintptr(v8 + 52)))
							if v9 != nil {
								v9(v8, a3, a2, v6, a4)
							}
						}
						v7 = (*int32)(unsafe.Add(unsafe.Pointer(v7), 4*1))
						result = func() int32 {
							p := &v10
							*p--
							return *p
						}()
						if v10 == 0 {
							break
						}
					}
				}
			}
		}
	}
	return result
}
func nox_xxx_playerTraceAttack_538330(a1 int32, a2 int32) int32 {
	var (
		v2  int32
		v3  int32
		v4  int32
		v5  int32
		v6  int32
		v8  int32
		v9  int32
		v12 int32
		a3  float32
		v15 float32
		v16 float32
		v17 float32
		v18 float32
		v19 float32
		v20 int32
		v21 int32
		a1a int4
		v23 float32
	)
	v2 = a1
	v21 = 0
	if a1 == 0 || a2 == 0 {
		return 0
	}
	dword_5d4594_2488656 = 0
	dword_5d4594_2488660 = 0
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 28))))
	if v3 != 0 && (func() int32 {
		v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 12))))
		return int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v4))), 1))) & 0x40
	}()) != 0 {
		sub_518040(a2+16, *(*float32)(unsafe.Pointer(uintptr(a2 + 8))), int32(funAddr(sub_538510)), a2)
		v21 = 25
	} else {
		dword_5d4594_2488660 = 0
		dword_5d4594_2488652 = *(*uint32)(unsafe.Pointer(uintptr(a2 + 8)))
		sub_518040(a1+56, *(*float32)(unsafe.Pointer(uintptr(a2 + 8))), int32(funAddr(sub_5386A0)), a1)
		if dword_5d4594_2488660 != 0 {
			sub_538510(*(*int32)(unsafe.Pointer(&dword_5d4594_2488660)), a2)
		}
	}
	v23 = float32(float64(v21))
	v15 = *(*float32)(unsafe.Pointer(uintptr(a2 + 16))) - *(*float32)(unsafe.Pointer(uintptr(a2 + 8))) - v23
	v5 = nox_float2int(v15)
	v16 = *(*float32)(unsafe.Pointer(uintptr(a2 + 20))) - *(*float32)(unsafe.Pointer(uintptr(a2 + 8))) - v23
	a1a.field_0 = v5 / 23
	v6 = nox_float2int(v16)
	v17 = *(*float32)(unsafe.Pointer(uintptr(a2 + 8))) + *(*float32)(unsafe.Pointer(uintptr(a2 + 16))) + v23
	a1a.field_4 = v6 / 23
	v8 = nox_float2int(v17)
	v18 = *(*float32)(unsafe.Pointer(uintptr(a2 + 20))) + *(*float32)(unsafe.Pointer(uintptr(a2 + 8))) + v23
	a1a.field_8 = v8 / 23
	v9 = nox_float2int(v18)
	v12 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 28))))
	a1a.field_C = v9 / 23
	if v12 == 0 {
		v12 = v2
	}
	a3 = v23 + *(*float32)(unsafe.Pointer(uintptr(a2 + 8)))
	nox_xxx_mapDamageToWalls_534FC0(&a1a, unsafe.Pointer(uintptr(a2+16)), a3, int32(int64(float64(*(*float32)(unsafe.Pointer(uintptr(a2))))+0.5)), int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 4)))), unsafe.Pointer(uintptr(v12)))
	if *(*uint32)(unsafe.Pointer(uintptr(a2 + 28))) != 0 {
		if dword_5d4594_2488656 != 0 {
			v20 = int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 4))))
			v19 = float32(nox_xxx_gamedataGetFloat_419D40(internCStr("ItemDamagePercentage")) * float64(*(*float32)(unsafe.Pointer(uintptr(a2)))))
			nox_xxx_playerDamageWeapon_4E1560(int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 28)))), int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 12)))), *(*int32)(unsafe.Pointer(&dword_5d4594_2488660)), *(*int32)(unsafe.Pointer(&dword_5d4594_2488660)), v19, v20)
		}
	}
	return int32(dword_5d4594_2488656)
}
func sub_538510(a1 int32, a2 int32) {
	var (
		v2  int32
		v3  int32
		v4  int32
		v5  float32
		v6  float32
		v7  float32
		v8  int32
		v9  int32
		v10 int32
		v11 int8
		v12 float4
	)
	if a1 != 0 {
		if a2 != 0 {
			v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 12))))
			if v2 != 0 {
				if a1 != v2 {
					if int32(uint8(int8(nox_server_testTwoPointsAndDirection_4E6E50((*float2)(unsafe.Pointer(uintptr(v2+56))), int32(*(*int16)(unsafe.Pointer(uintptr(v2 + 124)))), (*float2)(unsafe.Pointer(uintptr(a1+56)))))))&int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 32)))) != 0 {
						v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))))
						if (v3&0x8040) == 0 && (*(*uint32)(unsafe.Pointer(uintptr(a2 + 24))) != 0 || (v3&8) == 0) {
							if int32(*(*uint16)(unsafe.Pointer(uintptr(a1 + 24)))) != 0x4000 {
								dword_5d4594_2488656 = 1
							}
							v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 12))))
							v12.field_0 = *(*float32)(unsafe.Pointer(uintptr(v4 + 56)))
							v5 = *(*float32)(unsafe.Pointer(uintptr(a1 + 60)))
							v6 = *(*float32)(unsafe.Pointer(uintptr(v4 + 60)))
							v7 = *(*float32)(unsafe.Pointer(uintptr(a1 + 56)))
							v12.field_4 = v6
							v12.field_8 = v7
							v12.field_C = v5
							if nox_xxx_mapTraceRay_535250(&v12, nil, nil, 5) != 0 {
								nox_xxx_playerPreAttackEffects_538290(a1, int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 12)))), int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 28)))), a2)
								(*(*func(int32, uint32, uint32, uint32, uint32))(unsafe.Pointer(uintptr(a1 + 716))))(a1, *(*uint32)(unsafe.Pointer(uintptr(a2 + 12))), *(*uint32)(unsafe.Pointer(uintptr(a2 + 28))), uint32(int32(int64(float64(*(*float32)(unsafe.Pointer(uintptr(a2))))+0.5))), uint32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 4)))))
								if nox_common_gameFlags_check_40A5C0(2048) {
									if int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 12))) + 8))))&4 != 0 {
										if (int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8)))) & 2) == 0 {
											v8 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 556))))
											if v8 != 0 {
												if int32(*(*uint16)(unsafe.Pointer(uintptr(v8 + 4)))) != 0 {
													v9 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))))
													if (v9&0x8000) == 0 && (v9&0x20) == 0 {
														nox_xxx_netSendPointFx_522FF0(-117, (*float2)(unsafe.Pointer(uintptr(a1+56))))
													}
												}
											}
										}
									}
								}
								if *(*uint32)(unsafe.Pointer(uintptr(a2 + 28))) == 0 {
									v10 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 12))))
									if int32(*(*uint8)(unsafe.Pointer(uintptr(v10 + 8))))&4 != 0 {
										v11 = int8(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v10 + 748))) + 276))) + 8))))
									} else {
										v11 = int8(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v10 + 748))) + 2068))))
									}
									if int32(v11) == 25 {
										nox_xxx_objectApplyForce_52DF80((*float32)(unsafe.Pointer(uintptr(v10+56))), (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 20.0)
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func sub_5386A0(a3 int32, a2 int32) {
	var (
		v2 int32
		v3 int32
		v4 float64
		v5 float64
		v6 *float32
		v7 float64
		a4 float2
	)
	if a2 != a3 {
		if a2 != 0 {
			if a3 != 0 {
				v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a3 + 16))))
				if (v2&0x8049) == 0 && (int32(*(*uint8)(unsafe.Pointer(uintptr(a3 + 8))))&6 != 0 || (v2&0x10) == 0 || (v2&0x80) != 0) && (nox_xxx_unitIsEnemyTo_5330C0((*nox_object_t)(unsafe.Pointer(uintptr(a2))), (*nox_object_t)(unsafe.Pointer(uintptr(a3)))) != 0 || (int32(*(*uint8)(unsafe.Pointer(uintptr(a3 + 8))))&6) == 0) {
					v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a3 + 16))))
					if (v3 & 0x8000) == 0 {
						if nox_xxx_unitCanInteractWith_5370E0((*nox_object_t)(unsafe.Pointer(uintptr(a2))), (*nox_object_t)(unsafe.Pointer(uintptr(a3))), 1) != 0 {
							if float64(*(*float32)(unsafe.Pointer(&dword_5d4594_2488652))) > 0.0 {
								a4.field_0 = *(*float32)(unsafe.Pointer(uintptr(a3 + 56))) - *(*float32)(unsafe.Pointer(uintptr(a2 + 56)))
								v4 = float64(*(*float32)(unsafe.Pointer(uintptr(a3 + 60))) - *(*float32)(unsafe.Pointer(uintptr(a2 + 60))))
								a4.field_4 = float32(v4)
								v5 = math.Sqrt(v4*float64(a4.field_4) + float64(a4.field_0*a4.field_0))
								if v5 != 0.0 {
									v6 = mem_getFloatPtr(0x587000, uint32(int32(*(*int16)(unsafe.Pointer(uintptr(a2 + 124))))*8)+194136)
									if float64(a4.field_4)/v5*float64(*(*float32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(float32(0))*1)))+float64(a4.field_0)/v5*float64(*v6) > 0.5 {
										if *(*uint32)(unsafe.Pointer(uintptr(a3 + 172))) == 2 {
											v5 = v5 - float64(*(*float32)(unsafe.Pointer(uintptr(a3 + 176))))
										} else if *(*uint32)(unsafe.Pointer(uintptr(a3 + 172))) == 3 {
											v7 = sub_54A990((*float2)(unsafe.Pointer(uintptr(a2+56))), *(*float32)(unsafe.Pointer(&dword_5d4594_2488652)), a3, &a4)
											if v7 < 0.0 {
												return
											}
											v5 = float64(*(*float32)(unsafe.Pointer(&dword_5d4594_2488652))) - v7
										}
										if v5 < 0.0 {
											v5 = 0.0
										}
										if (v5 < float64(*(*float32)(unsafe.Pointer(&dword_5d4594_2488652))) || dword_5d4594_2488660 != 0 && (int32(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_2488660 + 8))))&2) == 0 && (int32(*(*uint8)(unsafe.Pointer(uintptr(a3 + 8))))&2) == 2) && (dword_5d4594_2488660 == 0 || (int32(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_2488660 + 8))))&2) == 0) {
											*(*float32)(unsafe.Pointer(&dword_5d4594_2488652)) = float32(v5)
											dword_5d4594_2488660 = uint32(a3)
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func nox_xxx_itemApplyAttackEffect_538840(a1 int32, a2 int32, a3 int32) int32 {
	var (
		v3     int32
		v4     *int32
		v5     int32
		v6     func(int32, int32, int32, uint32, int32)
		result int32
		v8     int32
	)
	v3 = a1
	v8 = 4
	v4 = *(**int32)(unsafe.Pointer(uintptr(v3 + 692)))
	for {
		v5 = *v4
		if *v4 != 0 {
			v6 = *(*func(int32, int32, int32, uint32, int32))(unsafe.Pointer(uintptr(v5 + 40)))
			if v6 != nil {
				v6(v5, v3, a2, 0, a3)
			}
		}
		v4 = (*int32)(unsafe.Add(unsafe.Pointer(v4), 4*1))
		result = func() int32 {
			p := &v8
			*p--
			return *p
		}()
		if v8 == 0 {
			break
		}
	}
	return result
}
func nox_xxx_playerAttack_538960(a1p *nox_object_t) int32 {
	var (
		a1  int32 = int32(uintptr(unsafe.Pointer(a1p)))
		v1  *float32
		v2  int32
		v3  int32
		v4  int32
		v5  int32
		v6  int8
		v7  float32
		v8  float32
		v9  int8
		v10 float64
		v11 int32
		v12 int32
		v13 float64
		v14 int16
		v15 int32
		v16 int32
		v17 float64
		v18 int32
		v19 int32
		v20 int8
		v21 int32
		v22 float64
		v23 float64
		v24 int32
		v25 int32
		v26 int32
		v27 float64
		v28 float64
		v29 float64
		v30 int32
		v31 float32
		v32 float64
		v33 *uint32
		v34 int32
		v35 int32
		v36 int16
		v37 float64
		v38 int32
		v39 float32
		v40 *byte
		v41 float64
		v42 *uint32
		v43 int32
		v44 int16
		v45 int8
		v46 int32
		v47 float64
		v48 float64
		v49 int32
		v50 float64
		v51 float64
		v52 int32
		v53 float64
		v54 float64
		v55 int32
		v56 int32
		v57 int32
		v58 float64
		v59 float64
		v60 int32
		v61 float64
		v62 float64
		v63 *uint32
	)
	_ = v63
	var v64 int32
	var v65 int8
	var v66 uint8
	var v67 *uint32
	_ = v67
	var v68 int32
	var v69 bool
	var v70 int32
	var v71 int32
	var v72 uint8
	var v73 int32
	var v75 float32
	var v76 int16 = 0
	var v77 int32 = 0
	var v78 int32
	var v79 int32
	var v80 int32
	var v81 int32
	var v82 [36]byte
	var v83 float4
	var v86 int32
	var v87 [88]byte
	v1 = nil
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))))
	if v2&4 != 0 {
		v79 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(v79 + 276))))
		v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v79 + 104))))
		v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 4))))
		v6 = int8(*(*uint8)(unsafe.Pointer(uintptr(v3 + 8))))
		*((*uint8)(unsafe.Pointer(&v76))) = *(*uint8)(unsafe.Pointer(uintptr(v79 + 236)))
	} else {
		if (v2&2) == 0 || (int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 12))))&0x10) == 0 {
			return 0
		}
		v79 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
		v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(v79 + 2056))))
		v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v79 + 2064))))
		v6 = int8(*(*uint8)(unsafe.Pointer(uintptr(v79 + 2068))))
		*((*uint8)(unsafe.Pointer(&v76))) = *(*uint8)(unsafe.Pointer(uintptr(v79 + 481)))
	}
	*((*uint8)(unsafe.Pointer(&v80))) = uint8(v6)
	if v4 != 0 {
		v1 = (*float32)(nox_xxx_getProjectileClassById_413250(int32(*(*uint16)(unsafe.Pointer(uintptr(v4 + 4))))))
		if v1 == nil {
			return 0
		}
		*(*uint32)(unsafe.Pointer(uintptr(v4 + 56))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 56)))
		*(*uint32)(unsafe.Pointer(uintptr(v4 + 60))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 60)))
		*(*uint32)(unsafe.Pointer(uintptr(v4 + 72))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 56)))
		*(*uint32)(unsafe.Pointer(uintptr(v4 + 76))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 60)))
	} else if int32(v6) == 0 {
		*((*uint8)(unsafe.Pointer(&v80))) = uint8(int8(nox_common_randomInt_415FA0(23, 24)))
		if (int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8)))) & 4) == 0 {
			*(*uint8)(unsafe.Pointer(uintptr(v79 + 2068))) = uint8(int8(v80))
		} else {
			if int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v79 + 276))) + 2251)))) == 0 && nox_common_randomInt_415FA0(0, 100) >= 75 {
				*((*uint8)(unsafe.Pointer(&v80))) = 25
			}
			if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4 != 0 {
				*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v79 + 276))) + 8))) = uint8(int8(v80))
			} else {
				*(*uint8)(unsafe.Pointer(uintptr(v79 + 2068))) = uint8(int8(v80))
			}
		}
	}
	*(*uint32)(unsafe.Pointer(&v82[28])) = uint32(v4)
	v81 = nox_xxx_unitGetStrength_4F9FD0(a1)
	if nox_common_playerIsAbilityActive_4FC250((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 2) != 0 && nox_xxx_probablyWarcryCheck_4FC3E0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 2) != 0 {
		nox_xxx_animPlayerGetFrameRange_4F9F90(46, &v77, &v78)
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v76))), unsafe.Sizeof(int16(0))-1)) = uint8((gameFrame() - *(*uint32)(unsafe.Pointer(uintptr(a1 + 136)))) / uint32(v78+1))
		if int32(v76) == 770 {
			v7 = *(*float32)(unsafe.Pointer(uintptr(a1 + 56)))
			v8 = *(*float32)(unsafe.Pointer(uintptr(a1 + 60)))
			*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v83.field_0))), 4*0)) = uint32(a1)
			v83.field_4 = v7
			v83.field_8 = v8
			nox_xxx_earthquakeSend_4D9110((*float32)(unsafe.Pointer(uintptr(a1+56))), 15)
			nox_xxx_unitsGetInCircle_517F90((*float2)(unsafe.Pointer(uintptr(a1+56))), 300.0, unsafe.Pointer(funAddr(nox_xxx_warcryStunMonsters_539B90)), unsafe.Pointer(uintptr(a1)))
			nox_xxx_castCounterSpell_52BBB0(13, a1, a1, a1)
		}
		if int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v76))), unsafe.Sizeof(int16(0))-1))) >= v77 {
			sub_4FC440((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 2)
		}
		goto LABEL_159
	}
	if nox_common_playerIsAbilityActive_4FC250((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 1) != 0 {
		if nox_xxx_testUnitBuffs_4FF350((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 25) == 0 && nox_xxx_testUnitBuffs_4FF350((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 5) == 0 {
			nox_xxx_animPlayerGetFrameRange_4F9F90(45, &v77, &v78)
			v9 = int8(uint8((gameFrame() - *(*uint32)(unsafe.Pointer(uintptr(a1 + 136)))) / uint32(v78+1)))
			v10 = float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 548)))) * 6.0
			*(*float32)(unsafe.Pointer(uintptr(a1 + 544))) = float32(v10)
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v76))), unsafe.Sizeof(int16(0))-1)) = uint8(v9)
			v11 = int32(*(*int16)(unsafe.Pointer(uintptr(a1 + 124)))) * 8
			v12 = v77 - 1
			*(*float32)(unsafe.Pointer(uintptr(a1 + 88))) = float32(v10*float64(*mem_getFloatPtr(0x587000, uint32(v11)+194136)) + float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 88)))))
			*(*float32)(unsafe.Pointer(uintptr(a1 + 92))) = float32(v10*float64(*mem_getFloatPtr(0x587000, uint32(v11)+194140)) + float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 92)))))
			if int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v76))), unsafe.Sizeof(int16(0))-1))) >= v12 {
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v76))), unsafe.Sizeof(int16(0))-1)) = 0
			}
			goto LABEL_159
		}
		return 0
	}
	if v4 == 0 {
		nox_xxx_animPlayerGetFrameRange_4F9F90(int32(uint8(int8(v80))), &v77, &v78)
		if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4 != 0 && *(*uint32)(unsafe.Pointer(uintptr(v79))) == 0 {
			*(*uint32)(unsafe.Pointer(uintptr(v79))) = gameFrame() + uint32(v77*(v78+1))
		}
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v76))), unsafe.Sizeof(int16(0))-1)) = uint8((gameFrame() - *(*uint32)(unsafe.Pointer(uintptr(a1 + 136)))) / uint32(v78+1))
		if int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v76))), unsafe.Sizeof(int16(0))-1))) < v77 {
			goto LABEL_159
		}
		if int32(uint8(int8(v80))) < 0x17 {
			v13 = 0.0
			v14 = 0
		} else if int32(uint8(int8(v80))) <= 0x18 {
			v13 = 0.039999999
			v14 = 5
		} else if int32(uint8(int8(v80))) == 25 {
			v13 = 0.039999999
			v14 = 10
		} else {
			v13 = 0.0
			v14 = 0
		}
		v15 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 56))))
		v16 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 60))))
		*(*float32)(unsafe.Pointer(&v87[64])) = float32(v13)
		*(*uint16)(unsafe.Pointer(&v87[72])) = uint16(v14)
		*(*uint32)(unsafe.Pointer(&v82[16])) = uint32(v15)
		*(*uint32)(unsafe.Pointer(&v82[20])) = uint32(v16)
		*(*uint16)(unsafe.Pointer(&v87[60])) = 0
		*(*float32)(unsafe.Pointer(&v82[0])) = float32(nox_xxx_calcBoltDamage_4EF1E0(v81, int32(uintptr(unsafe.Pointer(&v87[0])))))
		v17 = float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 176)))) + 20.0
		v82[4] = 10
		*(*uint32)(unsafe.Pointer(&v82[12])) = uint32(a1)
		*(*uint32)(unsafe.Pointer(&v82[24])) = 0
		*(*float32)(unsafe.Pointer(&v82[8])) = float32(v17)
		v82[32] = 1
		*(*uint32)(unsafe.Pointer(&v82[28])) = 0
		v18 = nox_xxx_playerTraceAttack_538330(a1, int32(uintptr(unsafe.Pointer(&v82[0]))))
		if v18 == 0 {
			nox_xxx_aud_501960(879, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
		}
		goto LABEL_159
	}
	if uint32(v5)&0x47F8000 != 0 {
		v19 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 736))))
		v86 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 736))))
		if (v5&0x8000) != 0 || (func() int32 {
			v20 = int8(*(*uint8)(unsafe.Pointer(uintptr(v19 + 96))))
			v80 = 31
			return int32(v20) & 2
		}()) != 0 {
			v80 = 29
		}
		nox_xxx_animPlayerGetFrameRange_4F9F90(v80, &v77, &v78)
		if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4 != 0 && *(*uint32)(unsafe.Pointer(uintptr(v79))) == 0 {
			if v80 == 29 {
				*(*uint32)(unsafe.Pointer(uintptr(v79))) = gameFrame() + uint32(v77*(v78+1))
			} else {
				*(*uint32)(unsafe.Pointer(uintptr(v79))) = gameFrame()
			}
		}
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v76))), unsafe.Sizeof(int16(0))-1)) = uint8((gameFrame() - *(*uint32)(unsafe.Pointer(uintptr(a1 + 136)))) / uint32(v78+1))
		if int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v76))), unsafe.Sizeof(int16(0))-1))) >= v77 {
			if uint32(v5)&0x47F0000 != 0 {
				*(*uint32)(unsafe.Pointer(uintptr(v86 + 96))) &= 0xFFFFFFFD
			}
			goto LABEL_159
		}
		if v80 != 29 || int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v76))), unsafe.Sizeof(int16(0))-1))) != v77/2 || int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v76))), unsafe.Sizeof(int16(0))-1))) <= int32(uint8(int8(v76))) {
			if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&2 != 0 {
				if v80 != 29 && int32(v76) == 256 {
					v24 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 736))))
					nox_xxx_useByNetCode_53F8E0(a1, v4)
					if int32(*(*uint8)(unsafe.Pointer(uintptr(v24 + 108)))) == 0 {
						if int32(*(*uint8)(unsafe.Pointer(uintptr(v24 + 109)))) != 0 {
							nox_xxx_equipWeaponNPC_53A030(a1, v4)
						}
					}
				}
			}
			goto LABEL_159
		}
		v21 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 60))))
		*(*uint32)(unsafe.Pointer(&v82[16])) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 56)))
		*(*uint32)(unsafe.Pointer(&v82[20])) = uint32(v21)
		*(*float32)(unsafe.Pointer(&v82[0])) = float32(nox_xxx_calcBoltDamage_4EF1E0(v81, int32(uintptr(unsafe.Pointer(v1)))))
		v22 = float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 176))))
		v82[4] = 0
		*(*uint32)(unsafe.Pointer(&v82[12])) = uint32(a1)
		v23 = v22 + float64(*(*float32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(float32(0))*17)))
		*(*uint32)(unsafe.Pointer(&v82[24])) = 0
		v82[32] = 1
		*(*float32)(unsafe.Pointer(&v82[8])) = float32(v23)
		nox_xxx_itemApplyAttackEffect_538840(v4, a1, int32(uintptr(unsafe.Pointer(&v82[0]))))
		v18 = nox_xxx_playerTraceAttack_538330(a1, int32(uintptr(unsafe.Pointer(&v82[0]))))
		if v18 == 0 {
			nox_xxx_aud_501960(879, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
		}
		goto LABEL_159
	}
	if uint32(v5)&0x7800000 != 0 {
		v25 = bool2int32((uint32(v5)&0x3800000) != 0) + 31
		nox_xxx_animPlayerGetFrameRange_4F9F90(v25, &v77, &v78)
		if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4 != 0 && *(*uint32)(unsafe.Pointer(uintptr(v79))) == 0 {
			*(*uint32)(unsafe.Pointer(uintptr(v79))) = gameFrame() + uint32(v77*(v78+1))
		}
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v76))), unsafe.Sizeof(int16(0))-1)) = uint8((gameFrame() - *(*uint32)(unsafe.Pointer(uintptr(a1 + 136)))) / uint32(v78+1))
		if v25 != 32 || int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v76))), unsafe.Sizeof(int16(0))-1))) != v77/2 || int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v76))), unsafe.Sizeof(int16(0))-1))) <= int32(uint8(int8(v76))) {
			goto LABEL_159
		}
		v26 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 60))))
		*(*uint32)(unsafe.Pointer(&v82[16])) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 56)))
		*(*uint32)(unsafe.Pointer(&v82[20])) = uint32(v26)
		*(*float32)(unsafe.Pointer(&v82[0])) = float32(nox_xxx_calcBoltDamage_4EF1E0(v81, int32(uintptr(unsafe.Pointer(v1)))))
		v27 = float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 176))))
		v82[4] = 0
		*(*uint32)(unsafe.Pointer(&v82[12])) = uint32(a1)
		v28 = v27 + float64(*(*float32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(float32(0))*17)))
		*(*uint32)(unsafe.Pointer(&v82[24])) = 0
		v82[32] = 1
		*(*float32)(unsafe.Pointer(&v82[8])) = float32(v28)
		nox_xxx_itemApplyAttackEffect_538840(v4, a1, int32(uintptr(unsafe.Pointer(&v82[0]))))
		v18 = nox_xxx_playerTraceAttack_538330(a1, int32(uintptr(unsafe.Pointer(&v82[0]))))
		if v18 == 0 {
			nox_xxx_aud_501960(879, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
		}
		goto LABEL_159
	}
	if (v5 & 0x40) == 0 {
		if (v5 & 0x80) == 0 {
			if v5&0x200 != 0 {
				nox_xxx_animPlayerGetFrameRange_4F9F90(28, &v77, &v78)
				if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4 != 0 && *(*uint32)(unsafe.Pointer(uintptr(v79))) == 0 {
					*(*uint32)(unsafe.Pointer(uintptr(v79))) = gameFrame() + uint32(v77*(v78+1))
				}
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v76))), unsafe.Sizeof(int16(0))-1)) = uint8((gameFrame() - *(*uint32)(unsafe.Pointer(uintptr(a1 + 136)))) / uint32(v78+1))
				if int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v76))), unsafe.Sizeof(int16(0))-1))) == v77/2 && int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v76))), unsafe.Sizeof(int16(0))-1))) > int32(uint8(int8(v76))) {
					v46 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 60))))
					*(*uint32)(unsafe.Pointer(&v82[16])) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 56)))
					*(*uint32)(unsafe.Pointer(&v82[20])) = uint32(v46)
					*(*float32)(unsafe.Pointer(&v82[0])) = float32(nox_xxx_calcBoltDamage_4EF1E0(v81, int32(uintptr(unsafe.Pointer(v1)))))
					v47 = float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 176))))
					v82[4] = 0
					*(*uint32)(unsafe.Pointer(&v82[12])) = uint32(a1)
					v48 = v47 + float64(*(*float32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(float32(0))*17)))
					*(*uint32)(unsafe.Pointer(&v82[24])) = 0
					v82[32] = 1
					*(*float32)(unsafe.Pointer(&v82[8])) = float32(v48)
					nox_xxx_itemApplyAttackEffect_538840(v4, a1, int32(uintptr(unsafe.Pointer(&v82[0]))))
					if nox_xxx_playerTraceAttack_538330(a1, int32(uintptr(unsafe.Pointer(&v82[0])))) == 0 {
						nox_xxx_aud_501960(880, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
					}
				}
			} else if v5&0x100 != 0 {
				nox_xxx_animPlayerGetFrameRange_4F9F90(27, &v77, &v78)
				if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4 != 0 && *(*uint32)(unsafe.Pointer(uintptr(v79))) == 0 {
					*(*uint32)(unsafe.Pointer(uintptr(v79))) = gameFrame() + uint32(v77*(v78+1))
				}
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v76))), unsafe.Sizeof(int16(0))-1)) = uint8((gameFrame() - *(*uint32)(unsafe.Pointer(uintptr(a1 + 136)))) / uint32(v78+1))
				if int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v76))), unsafe.Sizeof(int16(0))-1))) == v77/2 && int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v76))), unsafe.Sizeof(int16(0))-1))) > int32(uint8(int8(v76))) {
					v49 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 60))))
					*(*uint32)(unsafe.Pointer(&v82[16])) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 56)))
					*(*uint32)(unsafe.Pointer(&v82[20])) = uint32(v49)
					*(*float32)(unsafe.Pointer(&v82[0])) = float32(nox_xxx_calcBoltDamage_4EF1E0(v81, int32(uintptr(unsafe.Pointer(v1)))))
					v50 = float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 176))))
					v82[4] = 0
					*(*uint32)(unsafe.Pointer(&v82[12])) = uint32(a1)
					v51 = v50 + float64(*(*float32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(float32(0))*17)))
					*(*uint32)(unsafe.Pointer(&v82[24])) = 0
					v82[32] = 1
					*(*float32)(unsafe.Pointer(&v82[8])) = float32(v51)
					nox_xxx_itemApplyAttackEffect_538840(v4, a1, int32(uintptr(unsafe.Pointer(&v82[0]))))
					if nox_xxx_playerTraceAttack_538330(a1, int32(uintptr(unsafe.Pointer(&v82[0])))) == 0 {
						nox_xxx_aud_501960(881, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
					}
				}
			} else if v5&0x400 != 0 {
				nox_xxx_animPlayerGetFrameRange_4F9F90(37, &v77, &v78)
				if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4 != 0 && *(*uint32)(unsafe.Pointer(uintptr(v79))) == 0 {
					*(*uint32)(unsafe.Pointer(uintptr(v79))) = gameFrame() + uint32(v77*(v78+1))
				}
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v76))), unsafe.Sizeof(int16(0))-1)) = uint8((gameFrame() - *(*uint32)(unsafe.Pointer(uintptr(a1 + 136)))) / uint32(v78+1))
				if int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v76))), unsafe.Sizeof(int16(0))-1))) == v77/2 && int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v76))), unsafe.Sizeof(int16(0))-1))) > int32(uint8(int8(v76))) {
					v52 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 60))))
					*(*uint32)(unsafe.Pointer(&v82[16])) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 56)))
					*(*uint32)(unsafe.Pointer(&v82[20])) = uint32(v52)
					*(*float32)(unsafe.Pointer(&v82[0])) = float32(nox_xxx_calcBoltDamage_4EF1E0(v81, int32(uintptr(unsafe.Pointer(v1)))))
					v53 = float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 176))))
					v82[4] = 0
					*(*uint32)(unsafe.Pointer(&v82[12])) = uint32(a1)
					v54 = v53 + float64(*(*float32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(float32(0))*17)))
					*(*uint32)(unsafe.Pointer(&v82[24])) = 0
					v82[32] = 1
					*(*float32)(unsafe.Pointer(&v82[8])) = float32(v54)
					nox_xxx_itemApplyAttackEffect_538840(v4, a1, int32(uintptr(unsafe.Pointer(&v82[0]))))
					if nox_xxx_playerTraceAttack_538330(a1, int32(uintptr(unsafe.Pointer(&v82[0])))) == 0 {
						nox_xxx_aud_501960(881, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
					}
				}
			} else if v5&0x4000 != 0 {
				nox_xxx_animPlayerGetFrameRange_4F9F90(39, &v77, &v78)
				if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4 != 0 && *(*uint32)(unsafe.Pointer(uintptr(v79))) == 0 {
					*(*uint32)(unsafe.Pointer(uintptr(v79))) = gameFrame() + uint32(v77*(v78+1))
				}
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v76))), unsafe.Sizeof(int16(0))-1)) = uint8((gameFrame() - *(*uint32)(unsafe.Pointer(uintptr(a1 + 136)))) / uint32(v78+1))
				if int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v76))), unsafe.Sizeof(int16(0))-1))) == v77/2 && int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v76))), unsafe.Sizeof(int16(0))-1))) > int32(uint8(int8(v76))) {
					v55 = int32(*(*int16)(unsafe.Pointer(uintptr(a1 + 124)))) * 8
					*(*float32)(unsafe.Pointer(&v82[16])) = float32(float64(*mem_getFloatPtr(0x587000, uint32(v55)+194136))*35.0 + float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 56)))))
					*(*float32)(unsafe.Pointer(&v82[20])) = float32(float64(*mem_getFloatPtr(0x587000, uint32(v55)+194140))*35.0 + float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 60)))))
					*(*float32)(unsafe.Pointer(&v82[0])) = float32(nox_xxx_calcBoltDamage_4EF1E0(v81, int32(uintptr(unsafe.Pointer(v1)))))
					v82[4] = 2
					*(*uint32)(unsafe.Pointer(&v82[12])) = uint32(a1)
					*(*float32)(unsafe.Pointer(&v82[8])) = *(*float32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(float32(0))*17))
					*(*uint32)(unsafe.Pointer(&v82[24])) = 1
					v82[32] = 1
					nox_xxx_itemApplyAttackEffect_538840(v4, a1, int32(uintptr(unsafe.Pointer(&v82[0]))))
					nox_xxx_playerTraceAttack_538330(a1, int32(uintptr(unsafe.Pointer(&v82[0]))))
					v75 = float32(float64(v81) * 0.1)
					v56 = nox_float2int(v75)
					nox_xxx_earthquakeSend_4D9110((*float32)(unsafe.Pointer(uintptr(a1+56))), v56)
					nox_xxx_aud_501960(882, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
				}
			} else if v5&0x800 != 0 {
				nox_xxx_animPlayerGetFrameRange_4F9F90(26, &v77, &v78)
				if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4 != 0 && *(*uint32)(unsafe.Pointer(uintptr(v79))) == 0 {
					*(*uint32)(unsafe.Pointer(uintptr(v79))) = gameFrame() + uint32(v77*(v78+1))
				}
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v76))), unsafe.Sizeof(int16(0))-1)) = uint8((gameFrame() - *(*uint32)(unsafe.Pointer(uintptr(a1 + 136)))) / uint32(v78+1))
				if int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v76))), unsafe.Sizeof(int16(0))-1))) == v77/2 && int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v76))), unsafe.Sizeof(int16(0))-1))) > int32(uint8(int8(v76))) {
					v57 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 60))))
					*(*uint32)(unsafe.Pointer(&v82[16])) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 56)))
					*(*uint32)(unsafe.Pointer(&v82[20])) = uint32(v57)
					*(*float32)(unsafe.Pointer(&v82[0])) = float32(nox_xxx_calcBoltDamage_4EF1E0(v81, int32(uintptr(unsafe.Pointer(v1)))))
					v58 = float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 176))))
					v82[4] = 2
					*(*uint32)(unsafe.Pointer(&v82[12])) = uint32(a1)
					v59 = v58 + float64(*(*float32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(float32(0))*17)))
					*(*uint32)(unsafe.Pointer(&v82[24])) = 1
					v82[32] = 1
					*(*float32)(unsafe.Pointer(&v82[8])) = float32(v59)
					nox_xxx_itemApplyAttackEffect_538840(v4, a1, int32(uintptr(unsafe.Pointer(&v82[0]))))
					if nox_xxx_playerTraceAttack_538330(a1, int32(uintptr(unsafe.Pointer(&v82[0])))) == 0 {
						nox_xxx_aud_501960(884, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
					}
				}
			} else if v5&0x3000 != 0 {
				nox_xxx_animPlayerGetFrameRange_4F9F90(35, &v77, &v78)
				if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4 != 0 && *(*uint32)(unsafe.Pointer(uintptr(v79))) == 0 {
					*(*uint32)(unsafe.Pointer(uintptr(v79))) = gameFrame() + uint32(v77*(v78+1))
				}
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v76))), unsafe.Sizeof(int16(0))-1)) = uint8((gameFrame() - *(*uint32)(unsafe.Pointer(uintptr(a1 + 136)))) / uint32(v78+1))
				if int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v76))), unsafe.Sizeof(int16(0))-1))) == v77/2 && int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v76))), unsafe.Sizeof(int16(0))-1))) > int32(uint8(int8(v76))) {
					v60 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 60))))
					*(*uint32)(unsafe.Pointer(&v82[16])) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 56)))
					*(*uint32)(unsafe.Pointer(&v82[20])) = uint32(v60)
					*(*float32)(unsafe.Pointer(&v82[0])) = float32(nox_xxx_calcBoltDamage_4EF1E0(v81, int32(uintptr(unsafe.Pointer(v1)))))
					v61 = float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 176))))
					v82[4] = 0
					*(*uint32)(unsafe.Pointer(&v82[12])) = uint32(a1)
					v62 = v61 + float64(*(*float32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(float32(0))*17)))
					*(*uint32)(unsafe.Pointer(&v82[24])) = 1
					v82[32] = 1
					*(*float32)(unsafe.Pointer(&v82[8])) = float32(v62)
					nox_xxx_itemApplyAttackEffect_538840(v4, a1, int32(uintptr(unsafe.Pointer(&v82[0]))))
					if nox_xxx_playerTraceAttack_538330(a1, int32(uintptr(unsafe.Pointer(&v82[0])))) == 0 {
						nox_xxx_aud_501960(883, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
					}
				}
			} else if v5&4 != 0 {
				nox_xxx_animPlayerGetFrameRange_4F9F90(33, &v77, &v78)
				if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4 != 0 {
					v63 = (*uint32)(unsafe.Pointer(uintptr(v79)))
					if *(*uint32)(unsafe.Pointer(uintptr(v79))) == 0 {
						v64 = nox_xxx_itemCheckReadinessEffect_4E0960(v4)
						*v63 = gameFrame() + uint32(v77*(v78+1)) - uint32(v64)
					}
				}
				v65 = int8(uint8((gameFrame() - *(*uint32)(unsafe.Pointer(uintptr(a1 + 136)))) / uint32(v78+1)))
				v66 = uint8(int8(nox_xxx_itemCheckReadinessEffect_4E0960(v4) + int32(v65)))
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v76))), unsafe.Sizeof(int16(0))-1)) = v66
				if int32(v66) >= v77-1 && int32(v66) > int32(uint8(int8(v76))) {
					nox_xxx_shootBowCrossbow1_539BD0(a1, v4)
					*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v76))), unsafe.Sizeof(int16(0))-1)) = uint8(int8(v77))
				}
			} else if v5&8 != 0 {
				nox_xxx_animPlayerGetFrameRange_4F9F90(34, &v77, &v78)
				if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4 != 0 {
					v67 = (*uint32)(unsafe.Pointer(uintptr(v79)))
					if *(*uint32)(unsafe.Pointer(uintptr(v79))) == 0 {
						v68 = nox_xxx_itemCheckReadinessEffect_4E0960(v4)
						*v67 = gameFrame() + uint32(v77*(v78+1)) - uint32(v68)
					}
				}
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v76))), unsafe.Sizeof(int16(0))-1)) = uint8((gameFrame() - *(*uint32)(unsafe.Pointer(uintptr(a1 + 136)))) / uint32(v78+1))
				v69 = int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v76))), unsafe.Sizeof(int16(0))-1))) < 1
				if int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v76))), unsafe.Sizeof(int16(0))-1))) == 1 {
					if int32(uint8(int8(v76))) == 0 {
						nox_xxx_shootBowCrossbow1_539BD0(a1, v4)
					}
					v69 = false
				}
				if !v69 {
					*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v76))), unsafe.Sizeof(int16(0))-1)) += uint8(int8(nox_xxx_itemCheckReadinessEffect_4E0960(v4)))
				}
			}
			goto LABEL_159
		}
		nox_xxx_animPlayerGetFrameRange_4F9F90(44, &v77, &v78)
		if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4 != 0 && *(*uint32)(unsafe.Pointer(uintptr(v79))) == 0 {
			*(*uint32)(unsafe.Pointer(uintptr(v79))) = gameFrame() + uint32(v77*(v78+1))
		}
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v76))), unsafe.Sizeof(int16(0))-1)) = uint8((gameFrame() - *(*uint32)(unsafe.Pointer(uintptr(a1 + 136)))) / uint32(v78+1))
		if int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v76))), unsafe.Sizeof(int16(0))-1))) != v77/2 || int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v76))), unsafe.Sizeof(int16(0))-1))) <= int32(uint8(int8(v76))) {
			goto LABEL_159
		}
		v37 = float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 176)))) + 4.0
		v38 = int32(*(*int16)(unsafe.Pointer(uintptr(a1 + 124)))) * 8
		v39 = *(*float32)(unsafe.Pointer(uintptr(a1 + 60)))
		v40 = *(**byte)(unsafe.Pointer(uintptr(v4 + 736)))
		v41 = v37*float64(*mem_getFloatPtr(0x587000, uint32(v38)+194136)) + float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 56))))
		v83.field_0 = *(*float32)(unsafe.Pointer(uintptr(a1 + 56)))
		v83.field_4 = v39
		v83.field_8 = float32(v41)
		v83.field_C = float32(v37*float64(*mem_getFloatPtr(0x587000, uint32(v38)+194140)) + float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 60)))))
		if nox_xxx_mapTraceRay_535250(&v83, nil, nil, 4) != 0 {
			v42 = (*uint32)(unsafe.Pointer(nox_xxx_newObjectByTypeID_4E3810(internCStr("FanChakramInMotion"))))
			v43 = int32(uintptr(unsafe.Pointer(v42)))
			if v42 == nil {
				return 0
			}
			*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v42), 4*175)) + 4))) = uint32(a1)
			nox_xxx_createAt_4DAA50((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v42)))))), (*nox_object_t)(unsafe.Pointer(uintptr(a1))), v83.field_8, v83.field_C)
			nox_xxx_modifSetItemAttrs_4E4990((*nox_object_t)(unsafe.Pointer(uintptr(v43))), *(**int32)(unsafe.Pointer(uintptr(v4 + 692))))
			*(*float32)(unsafe.Pointer(uintptr(v43 + 80))) = *mem_getFloatPtr(0x587000, uint32(int32(*(*int16)(unsafe.Pointer(uintptr(a1 + 124))))*8)+194136) * *(*float32)(unsafe.Pointer(uintptr(v43 + 544)))
			*(*float32)(unsafe.Pointer(uintptr(v43 + 84))) = *mem_getFloatPtr(0x587000, uint32(int32(*(*int16)(unsafe.Pointer(uintptr(a1 + 124))))*8)+194140) * *(*float32)(unsafe.Pointer(uintptr(v43 + 544)))
			v44 = int16(*(*uint16)(unsafe.Pointer(uintptr(a1 + 124))))
			*(*uint16)(unsafe.Pointer(uintptr(v43 + 124))) = uint16(v44)
			*(*uint16)(unsafe.Pointer(uintptr(v43 + 126))) = uint16(v44)
			nox_xxx_aud_501960(891, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
			if *(*byte)(unsafe.Add(unsafe.Pointer(v40), 2)) == 0 {
				v45 = int8(*(*byte)(unsafe.Add(unsafe.Pointer(v40), 1)) - 1)
				*(*byte)(unsafe.Add(unsafe.Pointer(v40), 1)) = byte(v45)
				if int32(v45) != 0 {
					if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4 != 0 {
						nox_xxx_netReportCharges_4D82B0(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v79 + 276))) + 2064)))), (*nox_object_t)(unsafe.Pointer(uintptr(v4))), v45, int8(*v40))
					}
				} else {
					sub_4ED0C0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), (*nox_object_t)(unsafe.Pointer(uintptr(v4))))
					nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(v4))))
					sub_539FB0((*uint32)(unsafe.Pointer(uintptr(a1))))
				}
			}
			goto LABEL_159
		}
		nox_xxx_aud_501960(323, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
		goto LABEL_159
	}
	nox_xxx_animPlayerGetFrameRange_4F9F90(44, &v77, &v78)
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4 != 0 && *(*uint32)(unsafe.Pointer(uintptr(v79))) == 0 {
		*(*uint32)(unsafe.Pointer(uintptr(v79))) = gameFrame() + uint32(v77*(v78+1))
	}
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v76))), unsafe.Sizeof(int16(0))-1)) = uint8((gameFrame() - *(*uint32)(unsafe.Pointer(uintptr(a1 + 136)))) / uint32(v78+1))
	if int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v76))), unsafe.Sizeof(int16(0))-1))) != v77/2 || int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v76))), unsafe.Sizeof(int16(0))-1))) <= int32(uint8(int8(v76))) {
		goto LABEL_159
	}
	v29 = float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 176)))) + 4.0
	v30 = int32(*(*int16)(unsafe.Pointer(uintptr(a1 + 124)))) * 8
	v31 = *(*float32)(unsafe.Pointer(uintptr(a1 + 60)))
	v32 = v29*float64(*mem_getFloatPtr(0x587000, uint32(v30)+194136)) + float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 56))))
	v83.field_0 = *(*float32)(unsafe.Pointer(uintptr(a1 + 56)))
	v83.field_4 = v31
	v83.field_8 = float32(v32)
	v83.field_C = float32(v29*float64(*mem_getFloatPtr(0x587000, uint32(v30)+194140)) + float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 60)))))
	if nox_xxx_mapTraceRay_535250(&v83, nil, nil, 5) == 0 {
		nox_xxx_aud_501960(323, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
		goto LABEL_159
	}
	v33 = (*uint32)(unsafe.Pointer(nox_xxx_newObjectByTypeID_4E3810(internCStr("RoundChakramInMotion"))))
	v34 = int32(uintptr(unsafe.Pointer(v33)))
	if v33 == nil {
		return 0
	}
	v35 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v33), 4*187)))
	sub_4ED0C0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), (*nox_object_t)(unsafe.Pointer(uintptr(v4))))
	nox_xxx_createAt_4DAA50((*nox_object_t)(unsafe.Pointer(uintptr(v34))), (*nox_object_t)(unsafe.Pointer(uintptr(a1))), v83.field_8, v83.field_C)
	nox_xxx_inventoryPutImpl_4F3070((*nox_object_t)(unsafe.Pointer(uintptr(v34))), (*nox_object_t)(unsafe.Pointer(uintptr(v4))), 1)
	nox_xxx_modifSetItemAttrs_4E4990((*nox_object_t)(unsafe.Pointer(uintptr(v34))), *(**int32)(unsafe.Pointer(uintptr(v4 + 692))))
	*(*float32)(unsafe.Pointer(uintptr(v34 + 80))) = *mem_getFloatPtr(0x587000, uint32(int32(*(*int16)(unsafe.Pointer(uintptr(a1 + 124))))*8)+194136) * *(*float32)(unsafe.Pointer(uintptr(v34 + 544)))
	*(*float32)(unsafe.Pointer(uintptr(v34 + 84))) = *mem_getFloatPtr(0x587000, uint32(int32(*(*int16)(unsafe.Pointer(uintptr(a1 + 124))))*8)+194140) * *(*float32)(unsafe.Pointer(uintptr(v34 + 544)))
	v36 = int16(*(*uint16)(unsafe.Pointer(uintptr(a1 + 124))))
	*(*uint16)(unsafe.Pointer(uintptr(v34 + 124))) = uint16(v36)
	*(*uint16)(unsafe.Pointer(uintptr(v34 + 126))) = uint16(v36)
	*(*uint8)(unsafe.Pointer(uintptr(v35 + 4))) = 4
	*(*uint32)(unsafe.Pointer(uintptr(v35 + 16))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 56)))
	*(*uint32)(unsafe.Pointer(uintptr(v35 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 60)))
	*(*uint8)(unsafe.Pointer(uintptr(v35 + 24))) = 2
	nox_xxx_aud_501960(891, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
LABEL_159:
	v70 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))))
	if v70&4 != 0 {
		v71 = v79
		v72 = *(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v76))), unsafe.Sizeof(int16(0))-1))
		*(*uint8)(unsafe.Pointer(uintptr(v79 + 236))) = *(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v76))), unsafe.Sizeof(int16(0))-1))
		if int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v76))), unsafe.Sizeof(int16(0))-1))) >= v77 {
			*(*uint8)(unsafe.Pointer(uintptr(v71 + 236))) = uint8(int8(v77 - 1))
		}
	} else if v70&2 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 12))))&0x10 != 0 {
		v73 = v79
		*(*uint8)(unsafe.Pointer(uintptr(v79 + 481))) = *(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v76))), unsafe.Sizeof(int16(0))-1))
		v72 = *(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v76))), unsafe.Sizeof(int16(0))-1))
		if int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v76))), unsafe.Sizeof(int16(0))-1))) >= v77 {
			*(*uint8)(unsafe.Pointer(uintptr(v73 + 481))) = uint8(int8(v77 - 1))
		}
	} else {
		v72 = *(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v76))), unsafe.Sizeof(int16(0))-1))
	}
	return bool2int32(int32(v72) < v77)
}
func nox_xxx_warcryStunMonsters_539B90(a1 int32, a2 int32) int16 {
	var result int16
	result = int16(a2)
	if a2 != 0 {
		result = int16(a1)
		if a1 != 0 {
			if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&2 != 0 && *(*uint32)(unsafe.Pointer(uintptr(a1 + 12)))&0x20000 != 0 && (*(*uint32)(unsafe.Pointer(uintptr(a1 + 16)))&0x8020) == 0 {
				nox_xxx_buffApplyTo_4FF380((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 5, 90, 3)
			}
		}
	}
	return result
}
func nox_xxx_shootBowCrossbow1_539BD0(a1 int32, a2 int32) int32 {
	var (
		v2 int32
		v3 *uint8
		v4 *byte
		v6 *uint32
		v7 int32
		v8 int32
	)
	v2 = nox_xxx_weaponInventoryEquipFlags_415820((*nox_object_t)(unsafe.Pointer(uintptr(a2))))
	v3 = *(**uint8)(unsafe.Pointer(uintptr(a2 + 736)))
	v4 = (*byte)(unsafe.Pointer(uintptr(v2)))
	if int32(*v3) == 0 {
		v6 = *(**uint32)(unsafe.Pointer(uintptr(a1 + 504)))
		if v6 != nil {
			for {
				v7 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), 4*4)))
				if v7&0x100 != 0 {
					if nox_xxx_weaponInventoryEquipFlags_415820((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v6))))))) == 2 {
						v8 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), 4*184)))
						if int32(*(*uint8)(unsafe.Pointer(uintptr(v8 + 1)))) != 0 || int32(*(*uint8)(unsafe.Pointer(uintptr(v8 + 2)))) == 1 {
							break
						}
					}
				}
				v6 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), 4*124)))))
				if v6 == nil {
					goto LABEL_14
				}
			}
			nox_xxx_shootBowCrossbow2_539D80(a1, int32(uintptr(unsafe.Pointer(v6))), a2, v4)
			return 1
		}
	LABEL_14:
		if nox_common_gameFlags_check_40A5C0(4096) && uintptr(unsafe.Pointer(v4)) == uintptr(4) {
			nox_xxx_shootBowCrossbow2_539D80(a1, 0, a2, (*byte)(unsafe.Pointer(uintptr(4))))
		}
		if (int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8)))) & 4) == 0 {
			return 0
		}
		if nox_xxx_playerTryReloadQuiver_539FF0((*uint32)(unsafe.Pointer(uintptr(a1)))) == 1 {
			nox_xxx_netPriMsgToPlayer_4DA2C0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), internCStr("pattack.c:ReloadQuiver"), 0)
			*v3 = 0
		} else {
			if nox_common_gameFlags_check_40A5C0(4096) && uintptr(unsafe.Pointer(v4)) == uintptr(4) {
				goto LABEL_25
			}
			nox_xxx_netPriMsgToPlayer_4DA2C0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), internCStr("pattack.c:NoQuiver"), 0)
		}
		if uintptr(unsafe.Pointer(v4)) != uintptr(4) {
			if uintptr(unsafe.Pointer(v4)) == uintptr(8) {
				nox_xxx_aud_501960(888, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
			}
			goto LABEL_29
		}
	LABEL_25:
		if !nox_common_gameFlags_check_40A5C0(4096) {
			nox_xxx_aud_501960(887, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
		}
	LABEL_29:
		nox_xxx_playerSetState_4FA020((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 13)
		return 0
	}
	nox_xxx_netPriMsgToPlayer_4DA2C0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), internCStr("pattack.c:ReloadingQuiver"), 0)
	if uintptr(unsafe.Pointer(v4)) == uintptr(4) {
		if !nox_common_gameFlags_check_40A5C0(4096) {
			nox_xxx_aud_501960(887, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
			*v3--
			return 0
		}
	} else if uintptr(unsafe.Pointer(v4)) == uintptr(8) {
		nox_xxx_aud_501960(888, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
	}
	*v3--
	return 0
}
func nox_xxx_shootBowCrossbow2_539D80(a1 int32, a2 int32, a3 int32, a4 *byte) *uint32 {
	var (
		v4     float64
		v5     *byte
		v6     int32
		v7     float64
		v8     float64
		v9     float32
		result *uint32
		v11    *uint32
		v12    int32
		v13    int16
		v14    int32
		v15    int8
		v16    int8
		v17    float32
		v18    float32
		v19    float4
	)
	v4 = float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 176)))) + 4.0
	if a2 != 0 {
		v5 = *(**byte)(unsafe.Pointer(uintptr(a2 + 736)))
	} else {
		v5 = a4
	}
	v6 = int32(*(*int16)(unsafe.Pointer(uintptr(a1 + 124)))) * 8
	v7 = v4 * float64(*mem_getFloatPtr(0x587000, uint32(v6)+194136))
	v19.field_4 = *(*float32)(unsafe.Pointer(uintptr(a1 + 60)))
	v17 = float32(v7 + float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 56)))))
	v8 = v4 * float64(*mem_getFloatPtr(0x587000, uint32(v6)+194140))
	v9 = *(*float32)(unsafe.Pointer(uintptr(a1 + 56)))
	v19.field_8 = v17
	v19.field_0 = v9
	v18 = float32(v8 + float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 60)))))
	v19.field_C = v18
	result = (*uint32)(unsafe.Pointer(uintptr(nox_xxx_mapTraceRay_535250(&v19, nil, nil, 5))))
	if result != nil {
		if a2 != 0 {
			if uintptr(unsafe.Pointer(a4)) == uintptr(4) {
				v11 = (*uint32)(unsafe.Pointer(nox_xxx_newObjectByTypeID_4E3810(internCStr("ArcherArrow"))))
			} else {
				v11 = (*uint32)(unsafe.Pointer(nox_xxx_newObjectByTypeID_4E3810(internCStr("ArcherBolt"))))
			}
		} else {
			v11 = (*uint32)(unsafe.Pointer(nox_xxx_newObjectByTypeID_4E3810(internCStr("WeakArcherArrow"))))
		}
		v12 = int32(uintptr(unsafe.Pointer(v11)))
		if v11 != nil {
			*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v11), 4*175)) + 4))) = uint32(a1)
			nox_xxx_createAt_4DAA50((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v11)))))), (*nox_object_t)(unsafe.Pointer(uintptr(a1))), v17, v18)
			if a2 != 0 {
				nox_xxx_modifSetItemAttrs_4E4990((*nox_object_t)(unsafe.Pointer(uintptr(v12))), *(**int32)(unsafe.Pointer(uintptr(a2 + 692))))
			}
			nox_xxx_shootApplyEffects_539F40(a1, a3, v12)
			*(*float32)(unsafe.Pointer(uintptr(v12 + 80))) = *mem_getFloatPtr(0x587000, uint32(int32(*(*int16)(unsafe.Pointer(uintptr(a1 + 124))))*8)+194136) * *(*float32)(unsafe.Pointer(uintptr(v12 + 544)))
			*(*float32)(unsafe.Pointer(uintptr(v12 + 84))) = *mem_getFloatPtr(0x587000, uint32(int32(*(*int16)(unsafe.Pointer(uintptr(a1 + 124))))*8)+194140) * *(*float32)(unsafe.Pointer(uintptr(v12 + 544)))
			v13 = int16(*(*uint16)(unsafe.Pointer(uintptr(a1 + 124))))
			*(*uint16)(unsafe.Pointer(uintptr(v12 + 124))) = uint16(v13)
			*(*uint16)(unsafe.Pointer(uintptr(v12 + 126))) = uint16(v13)
		}
		if a2 != 0 {
			if *(*byte)(unsafe.Add(unsafe.Pointer(v5), 2)) == 0 {
				if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4 != 0 {
					v14 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
					v15 = int8(*(*byte)(unsafe.Add(unsafe.Pointer(v5), 1)) - 1)
					v16 = int8(*v5)
					*(*byte)(unsafe.Add(unsafe.Pointer(v5), 1)) = byte(v15)
					nox_xxx_netReportCharges_4D82B0(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v14 + 276))) + 2064)))), (*nox_object_t)(unsafe.Pointer(uintptr(a2))), v15, v16)
					if *(*byte)(unsafe.Add(unsafe.Pointer(v5), 1)) == 0 {
						nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(a2))))
					}
				}
			}
		}
		if uintptr(unsafe.Pointer(a4)) == uintptr(4) {
			nox_xxx_aud_501960(885, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
		} else {
			nox_xxx_aud_501960(886, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
		}
	}
	return result
}
func nox_xxx_shootApplyEffects_539F40(a1 int32, a2 int32, a3 int32) int32 {
	var (
		v3     int32
		v4     int32
		v5     *int32
		v6     int32
		v7     func(int32, int32, int32, int32, int32) int32
		result int32
		v9     int32
		v10    int32
	)
	v3 = a3
	v4 = a2
	v10 = int32(*(*uint32)(unsafe.Pointer(uintptr(a3 + 692))))
	v9 = 2
	v5 = (*int32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 692))) + 8)))
	for {
		v6 = *v5
		if *v5 != 0 {
			if funAddr(*(*func(int32, int32, int32, int32) int8)(unsafe.Pointer(uintptr(v6 + 52)))) == funAddr(nox_xxx_recoilEffect_4E0640) {
				*(*uint32)(unsafe.Pointer(uintptr(v10 + 12))) = uint32(v6)
			} else {
				v7 = *(*func(int32, int32, int32, int32, int32) int32)(unsafe.Pointer(uintptr(v6 + 40)))
				if funAddr(v7) == funAddr(nox_xxx_effectProjectileSpeed_4E09B0) {
					v7(v6, v4, a1, 0, v3)
				}
			}
		}
		v5 = (*int32)(unsafe.Add(unsafe.Pointer(v5), 4*1))
		result = func() int32 {
			p := &v9
			*p--
			return *p
		}()
		if v9 == 0 {
			break
		}
	}
	return result
}
func sub_539FB0(a1 *uint32) int32 {
	var item *nox_object_t
	item = (*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*126)))))
	if item == nil {
		return 0
	}
	for nox_xxx_weaponInventoryEquipFlags_415820(item) != 128 {
		item = item.inv_next_item
		if item == nil {
			return 0
		}
	}
	return nox_xxx_playerEquipWeapon_53A420(a1, item, 1, 1)
}
func nox_xxx_playerTryReloadQuiver_539FF0(a1 *uint32) int32 {
	var item *nox_object_t
	item = (*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*126)))))
	if item == nil {
		return 0
	}
	for nox_xxx_weaponInventoryEquipFlags_415820(item) != 2 {
		item = item.inv_next_item
		if item == nil {
			return 0
		}
	}
	return nox_xxx_playerEquipWeapon_53A420(a1, item, 1, 1)
}
func nox_xxx_equipWeaponNPC_53A030(a1 int32, a2 int32) int32 {
	var (
		v2     int32
		result int32
		v4     int32
		v5     int32
	)
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	if (*(*uint32)(unsafe.Pointer(uintptr(a2 + 8))) & 0x1001000) == 0 {
		return 0
	}
	v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 16))))
	if (v4 & 0x100) == 0 {
		return 0
	}
	result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 504))))
	if result == 0 {
		return 0
	}
	for result != a2 {
		result = int32(*(*uint32)(unsafe.Pointer(uintptr(result + 496))))
		if result == 0 {
			return result
		}
	}
	*(*uint8)(unsafe.Pointer(uintptr(v2 + 2068))) = 0
	v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 16))))
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v5))), 1)) &= 0xFE
	*(*uint32)(unsafe.Pointer(uintptr(a2 + 16))) = uint32(v5)
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 12))))&0x10 != 0 {
		nox_xxx_npcSetItemEquipFlags_4E4B20(a1, (*nox_object_t)(unsafe.Pointer(uintptr(a2))), 0)
	}
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 12))))&0xC != 0 {
		sub_53A0F0(a1, 1, 1)
	}
	if (int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 12)))) & 2) == 0 {
		*(*uint32)(unsafe.Pointer(uintptr(v2 + 2064))) = 0
	}
	nox_xxx_itemApplyDisengageEffect_4F3030((*nox_object_t)(unsafe.Pointer(uintptr(a2))), a1)
	sub_4FEB60(a1, a2)
	return 1
}
func sub_53A0F0(a1 int32, a2 int32, a3 int32) {
	var (
		v3 int32
		v4 int32
	)
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 504))))
	if v3 != 0 {
		for {
			v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 16))))
			if v4&0x100 != 0 {
				if nox_xxx_weaponInventoryEquipFlags_415820((*nox_object_t)(unsafe.Pointer(uintptr(v3)))) == 2 {
					break
				}
			}
			v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 496))))
			if v3 == 0 {
				return
			}
		}
		nox_xxx_playerDequipWeapon_53A140((*uint32)(unsafe.Pointer(uintptr(a1))), (*nox_object_t)(unsafe.Pointer(uintptr(v3))), a2, a3)
	}
}
func nox_xxx_playerDequipWeapon_53A140(a1 *uint32, item *nox_object_t, a3 int32, a4 int32) int32 {
	var eflags int32 = nox_xxx_weaponInventoryEquipFlags_415820(item)
	if *(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*2))&2 != 0 {
		return nox_xxx_equipWeaponNPC_53A030(int32(uintptr(unsafe.Pointer(a1))), int32(uintptr(unsafe.Pointer(item))))
	}
	if (*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*2))&4) == 0 || (item.obj_class&0x1001000) == 0 {
		return 0
	}
	var v10 int32 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*187)))
	if int32(*(*uint8)(unsafe.Pointer(uintptr(v10 + 88)))) == 1 {
		nox_xxx_playerSetState_4FA020((*nox_object_t)(unsafe.Pointer(a1)), 13)
	}
	var v11 unsafe.Pointer = unsafe.Pointer(*(**uint32)(unsafe.Pointer(uintptr(v10 + 104))))
	if v11 == nil || v11 != unsafe.Pointer(item) && eflags != 2 {
		return 0
	}
	if eflags&0xC != 0 {
		sub_53A0F0(int32(uintptr(unsafe.Pointer(a1))), a3, a4)
	}
	sub_4FEB60(int32(uintptr(unsafe.Pointer(a1))), int32(uintptr(unsafe.Pointer(item))))
	if eflags == 2 {
		item.obj_flags &= 0xFFFFFEFF
		*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v10 + 276))) + 4))) &= 0xFFFFFFFD
		if a3 != 0 {
			nox_xxx_netReportDequip_4D8590(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v10 + 276))) + 2064)))), item)
		}
		if a4 != 0 {
			nox_xxx_netReportDequip_4D84C0(math.MaxUint8, item)
		}
	} else if *(*uint32)(unsafe.Pointer(uintptr(v10 + 104))) != 0 {
		var v13 int32 = int32(*(*uint32)(unsafe.Pointer(uintptr(v10 + 104))))
		*(*uint32)(unsafe.Pointer(uintptr(v13 + 16))) &= 0xFFFFFEFF
		*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v10 + 276))) + 4))) &= uint32(^nox_xxx_weaponInventoryEquipFlags_415820((*nox_object_t)(unsafe.Pointer(uintptr(v13)))))
		if a3 != 0 {
			nox_xxx_netReportDequip_4D8590(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v10 + 276))) + 2064)))), (*nox_object_t)(unsafe.Pointer(uintptr(v13))))
		}
		if a4 != 0 {
			nox_xxx_netReportDequip_4D84C0(math.MaxUint8, item)
		}
		*(*uint32)(unsafe.Pointer(uintptr(v10 + 104))) = 0
	}
	nox_xxx_itemApplyDisengageEffect_4F3030(item, int32(uintptr(unsafe.Pointer(a1))))
	if gameex_flags&2 != 0 {
		sub_980523((*nox_object_t)(unsafe.Pointer(a1)))
		var v16 int32 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*187)))
		if *(*uint32)(unsafe.Pointer(uintptr(v16 + 108))) == 0 || (uint32(nox_xxx_weaponInventoryEquipFlags_415820((*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v16 + 108))))))))&0x7FFE40C) == 0 {
			var v17 int32 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v16 + 276))) + 2500))))
			if v17 != 0 && int32(uint8(*(*uint32)(unsafe.Pointer(uintptr(v17 + 16))))) == 16 {
				nox_xxx_playerTryEquip_4F2F70((*nox_object_t)(unsafe.Pointer(a1)), (*nox_object_t)(unsafe.Pointer(uintptr(v17))))
			} else {
				var v18 unsafe.Pointer = unsafe.Pointer(sub_9805EB((*nox_object_t)(unsafe.Pointer(a1))))
				if v18 != nil {
					nox_xxx_playerTryEquip_4F2F70((*nox_object_t)(unsafe.Pointer(a1)), (*nox_object_t)(v18))
				}
			}
		}
	}
	return 1
}
func nox_xxx_NPCEquipWeapon_53A2C0(a1 int32, item *nox_object_t) int32 {
	var (
		v2     int32
		result int32
		v4     int32
		v5     int32
		v6     int32
		v7     int32
		v8     int32
	)
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	if (item.obj_class & 0x1001000) == 0 {
		return 0
	}
	v4 = int32(item.obj_flags)
	if v4&0x100 != 0 {
		return 0
	}
	result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 504))))
	if result == 0 {
		return 0
	}
	for unsafe.Pointer(uintptr(result)) != unsafe.Pointer(item) {
		result = int32(*(*uint32)(unsafe.Pointer(uintptr(result + 496))))
		if result == 0 {
			return result
		}
	}
	*(*uint8)(unsafe.Pointer(uintptr(v2 + 2068))) = 0
	if (item.obj_subclass & 0xC) == 0 {
		sub_53A0F0(a1, 1, 1)
	}
	v5 = int32(item.obj_subclass)
	if (v5 & 2) == 0 {
		v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 504))))
		if v6 != 0 {
			for {
				v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(v6 + 16))))
				if v7&0x100 != 0 {
					if *(*uint32)(unsafe.Pointer(uintptr(v6 + 8)))&0x1001000 != 0 && ((v5&0xC) == 0 || (int32(*(*uint8)(unsafe.Pointer(uintptr(v6 + 12))))&2) == 0) {
						break
					}
				}
				v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(v6 + 496))))
				if v6 == 0 {
					goto LABEL_22
				}
			}
			nox_xxx_equipWeaponNPC_53A030(a1, v6)
		}
	}
LABEL_22:
	v8 = int32(item.obj_flags)
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v8))), 1)) |= 1
	item.obj_flags = uint32(v8)
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 12))))&0x10 != 0 {
		nox_xxx_npcSetItemEquipFlags_4E4B20(a1, item, 1)
	}
	if (item.obj_subclass & 2) == 0 {
		*(*uint32)(unsafe.Pointer(uintptr(v2 + 2064))) = uint32(uintptr(unsafe.Pointer(item)))
	}
	nox_xxx_itemApplyEngageEffect_4F2FF0(item, a1)
	if uint32(nox_xxx_weaponInventoryEquipFlags_415820(item))&0x7FFE40C != 0 {
		sub_53A3D0((*uint32)(unsafe.Pointer(uintptr(a1))))
	}
	return 1
}
func sub_53A3D0(a1 *uint32) {
	var (
		i  *uint32
		v2 int32
	)
	if a1 != nil {
		for i = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*126))))); i != nil; i = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(i), 4*124))))) {
			if *(*uint32)(unsafe.Add(unsafe.Pointer(i), 4*2))&0x2000000 != 0 {
				v2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(i), 4*4)))
				if v2&0x100 != 0 {
					if uint32(nox_xxx_unitArmorInventoryEquipFlags_415C70((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(i))))))))&0x3000000 != 0 {
						sub_53E430(a1, (*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(i)))))), 1, 1)
					}
				}
			}
		}
	}
}
func nox_xxx_playerEquipWeapon_53A420(a1 *uint32, item *nox_object_t, a3 int32, a4 int32) int32 {
	var v4 int32 = nox_xxx_weaponInventoryEquipFlags_415820(item)
	if (item.obj_class & 0x1001000) == 0 {
		return 0
	}
	var v5 int32 = int32(item.obj_flags)
	if v5&0x100 != 0 {
		return 0
	}
	var v6 int32 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*2)))
	if v6&2 != 0 {
		return nox_xxx_NPCEquipWeapon_53A2C0(int32(uintptr(unsafe.Pointer(a1))), item)
	}
	if (v6 & 4) == 0 {
		return 0
	}
	var v8 int32 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*187)))
	if nox_xxx_probablyWarcryCheck_4FC3E0((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), 2) != 0 || nox_xxx_probablyWarcryCheck_4FC3E0((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), 1) != 0 {
		return 0
	}
	if nox_xxx_playerClassCanUseItem_57B3D0(item, int8(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v8 + 276))) + 2251))))) == 0 {
		nox_xxx_netPriMsgToPlayer_4DA2C0((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), internCStr("weapon.c:WeaponEquipClassFail"), 0)
		if a4 != 0 {
			nox_xxx_aud_501960(925, (*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), 2, int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*9))))
		}
		return 0
	}
	var v9 bool = nox_xxx_playerCheckStrength_4F3180((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), item)
	if !v9 {
		nox_xxx_netPriMsgToPlayer_4DA2C0((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), internCStr("weapon.c:WeaponEquipStrengthFail"), 0)
		if a4 != 0 {
			nox_xxx_aud_501960(925, (*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), 2, int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*9))))
		}
		return 0
	}
	var result int32 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*126)))
	if result == 0 {
		return 0
	}
	for unsafe.Pointer(uintptr(result)) != unsafe.Pointer(item) {
		result = int32(*(*uint32)(unsafe.Pointer(uintptr(result + 496))))
		if result == 0 {
			return result
		}
	}
	if int32(*(*uint8)(unsafe.Pointer(uintptr(v8 + 88)))) == 1 {
		nox_xxx_playerSetState_4FA020((*nox_object_t)(unsafe.Pointer(a1)), 13)
	}
	if v4 == 2 {
		if (int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v8 + 276))) + 4))))&0xC) == 0 && sub_53A680(int32(uintptr(unsafe.Pointer(a1)))) == 0 {
			nox_xxx_netPriMsgToPlayer_4DA2C0((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), internCStr("weapon.c:BowNotFound"), 0)
			if a4 != 0 {
				nox_xxx_aud_501960(925, (*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), 2, int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*9))))
			}
			return 0
		}
		sub_53A0F0(int32(uintptr(unsafe.Pointer(a1))), 1, 1)
	}
	var v10 int32 = int32(*(*uint32)(unsafe.Pointer(uintptr(v8 + 104))))
	if v10 != 0 && v4 != 2 && nox_xxx_playerDequipWeapon_53A140(a1, (*nox_object_t)(unsafe.Pointer(uintptr(v10))), 1, 1) == 0 {
		return 0
	}
	var v11 int32 = int32(item.obj_flags)
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v11))), 1)) |= 1
	item.obj_flags = uint32(v11)
	*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v8 + 276))) + 4))) |= uint32(v4)
	nox_xxx_netReportEquip_4D8540(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v8 + 276))) + 2064)))), (*uint32)(unsafe.Pointer(item)), a3)
	if v4 != 2 {
		*(*uint32)(unsafe.Pointer(uintptr(v8 + 104))) = uint32(uintptr(unsafe.Pointer(item)))
	}
	var v12 int32 = int32(item.obj_class)
	if v12&0x1000 != 0 && (item.obj_subclass&0x47F0000) != 0 {
		nox_xxx_netReportCharges_4D82B0(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v8 + 276))) + 2064)))), (*nox_object_t)(unsafe.Pointer((*uint32)(unsafe.Pointer(item)))), int8(*(*uint8)(unsafe.Pointer(uintptr(int32(uintptr(item.use_data)) + 108)))), int8(*(*uint8)(unsafe.Pointer(uintptr(int32(uintptr(item.use_data)) + 109)))))
	} else if uint32(v12)&0x1000000 != 0 {
		if v4&0x82 != 0 {
			nox_xxx_netReportCharges_4D82B0(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v8 + 276))) + 2064)))), (*nox_object_t)(unsafe.Pointer((*uint32)(unsafe.Pointer(item)))), int8(*(*uint8)(unsafe.Pointer(uintptr(int32(uintptr(item.use_data)) + 1)))), int8(*(*uint8)(item.use_data)))
		} else if v4&0xC != 0 {
			*(*uint8)(item.use_data) = 0
		}
	}
	nox_xxx_itemApplyEngageEffect_4F2FF0(item, int32(uintptr(unsafe.Pointer(a1))))
	if uint32(v4)&0x7FFE40C != 0 {
		sub_53A3D0(a1)
	}
	return 1
}
func sub_53A680(a1 int32) int32 {
	var item *nox_object_t
	item = *(**nox_object_t)(unsafe.Pointer(uintptr(a1 + 504)))
	if item == nil {
		return 0
	}
	for (nox_xxx_weaponInventoryEquipFlags_415820(item) & 0xC) == 0 {
		item = item.inv_next_item
		if item == nil {
			return 0
		}
	}
	return nox_xxx_playerEquipWeapon_53A420((*uint32)(unsafe.Pointer(uintptr(a1))), item, 1, 1)
}
func sub_53A6C0(a1 int32, item *nox_object_t) {
	var (
		v2 int32
		v3 int16
	)
	if a1 != 0 && item != nil {
		v2 = int32(item.obj_class)
		if v2&0x1000 != 0 {
			nox_xxx_aud_501960(830, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
		} else {
			v3 = int16(item.material)
			if int32(v3)&0x10 != 0 {
				nox_xxx_aud_501960(842, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
			} else if int32(v3)&8 != 0 {
				nox_xxx_aud_501960(844, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
			}
		}
	}
}
func sub_53A720(a1 int32, item *nox_object_t, a3 int32, a4 int32) int32 {
	var (
		v4  *uint32
		v5  int32
		v6  int32
		v8  int32
		v9  int32
		v10 int32
		v11 int32
		v12 float32
		v13 int32
	)
	v4 = (*uint32)(unsafe.Pointer(uintptr(a1)))
	v5 = 0
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4 != 0 {
		if nox_common_gameFlags_check_40A5C0(4096) {
			if item.obj_subclass&0x200000 != 0 {
				v12 = float32(nox_xxx_gamedataGetFloat_419D40(internCStr("ForceOfNatureStaffLimit")))
				v6 = nox_float2int(v12)
				if nox_xxx_inventoryCountObjects_4E7D30(a1, int32(item.obj_flags)) >= v6 {
					nox_xxx_netPriMsgToPlayer_4DA2C0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), internCStr("pickup.c:MaxSameItem"), 0)
					nox_xxx_aud_501960(925, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
					return 0
				}
			}
		}
	}
	if !nox_common_gameFlags_check_40A5C0(2048) && !nox_common_gameFlags_check_40A5C0(4096) && sub_409F40(2) != 0 {
		if (item.obj_subclass&0x82) == 0 && sub_4E7EC0(a1, item) != 0 {
			v5 = 1
		}
		if item.obj_subclass&0x40 != 0 {
			v8 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 516))))
			if v8 != 0 {
				for (*(*uint32)(unsafe.Pointer(uintptr(v8 + 8)))&0x1000000) == 0 || (int32(*(*uint8)(unsafe.Pointer(uintptr(v8 + 12))))&0x40) == 0 {
					v8 = int32(*(*uint32)(unsafe.Pointer(uintptr(v8 + 512))))
					if v8 == 0 {
						goto LABEL_18
					}
				}
				if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4 != 0 {
					nox_xxx_netPriMsgToPlayer_4DA2C0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), internCStr("weapon.c:CannotPickupDuplicateWeapon"), 0)
					nox_xxx_aud_501960(925, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 2, int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 36)))))
				}
				return 0
			}
		}
	LABEL_18:
		if v5 == 1 {
			if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4 != 0 {
				nox_xxx_netPriMsgToPlayer_4DA2C0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), internCStr("weapon.c:CannotPickupDuplicateWeapon"), 0)
				nox_xxx_aud_501960(925, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 2, int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 36)))))
			}
			return 0
		}
	}
	if !nox_common_gameFlags_check_40A5C0(2048) && !nox_common_gameFlags_check_40A5C0(4096) && int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4 != 0 && nox_xxx_playerClassCanUseItem_57B3D0(item, int8(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + 276))) + 2251))))) == 0 {
		nox_xxx_netPriMsgToPlayer_4DA2C0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), internCStr("weapon.c:WeaponEquipClassFail"), 0)
		nox_xxx_aud_501960(925, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 2, int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 36)))))
		return 0
	}
	if nox_xxx_pickupDefault_4F31E0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), item, a3) != 1 {
		return 0
	}
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4 != 0 {
		v9 = 0
		v13 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
		if *(*uint32)(unsafe.Pointer(uintptr(v13 + 104))) == 0 && sub_419E60((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v4))))))) == 0 && nox_xxx_weaponInventoryEquipFlags_415820(item) != 2 {
			v9 = nox_xxx_playerEquipWeapon_53A420(v4, item, a4, 0)
		}
		if sub_419E60((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v4))))))) == 0 && nox_xxx_weaponInventoryEquipFlags_415820(item) == 2 {
			v10 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v13 + 276))) + 4))))
			if v10&0xC != 0 {
				if (v10 & 2) == 0 {
					v9 = nox_xxx_playerEquipWeapon_53A420(v4, item, a4, 0)
				}
			}
		}
		if v9 == 0 {
			v11 = int32(item.obj_class)
			if v11&0x1000 != 0 && (item.obj_subclass&0x47F0000) != 0 {
				nox_xxx_netReportCharges_4D82B0(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v13 + 276))) + 2064)))), item, int8(*(*uint8)(unsafe.Pointer(uintptr(int32(uintptr(item.use_data)) + 108)))), int8(*(*uint8)(unsafe.Pointer(uintptr(int32(uintptr(item.use_data)) + 109)))))
			} else if uint32(v11)&0x1000000 != 0 && (item.obj_subclass&0x82) != 0 {
				nox_xxx_netReportCharges_4D82B0(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v13 + 276))) + 2064)))), item, int8(*(*uint8)(unsafe.Pointer(uintptr(int32(uintptr(item.use_data)) + 1)))), int8(*(*uint8)(item.use_data)))
			}
		}
	}
	sub_53A6C0(int32(uintptr(unsafe.Pointer(v4))), item)
	nox_xxx_decay_5116F0(item)
	return 1
}
func nox_xxx_sendMsgOblivionPickup_53A9C0(a1 int32, item *nox_object_t, a3 int32, a4 int32) int32 {
	var (
		v4 int32
		v5 int32
	)
	v4 = sub_53A720(a1, item, a3, a4)
	if v4 == 1 && int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4 != 0 && sub_419E60((*nox_object_t)(unsafe.Pointer(uintptr(a1)))) == 0 {
		v5 = int32(item.obj_subclass)
		if uint32(v5)&0x800000 != 0 {
			nox_xxx_netPriMsgToPlayer_4DA2C0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), internCStr("weapon.c:PickupHalberdOblivion"), 0)
			nox_xxx_aud_501960(914, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
		} else if uint32(v5)&0x1000000 != 0 {
			nox_xxx_netPriMsgToPlayer_4DA2C0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), internCStr("weapon.c:PickupHeartOblivion"), 0)
			nox_xxx_aud_501960(915, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
		} else if uint32(v5)&0x2000000 != 0 {
			nox_xxx_netPriMsgToPlayer_4DA2C0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), internCStr("weapon.c:PickupWierdlingOblivion"), 0)
			nox_xxx_aud_501960(916, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
		} else if uint32(v5)&0x4000000 != 0 {
			nox_xxx_netPriMsgToPlayer_4DA2C0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), internCStr("weapon.c:PickupOrbOblivion"), 0)
			nox_xxx_aud_501960(917, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
		}
		sub_57AF30(a1, 1)
		nox_xxx_playerTryEquip_4F2F70((*nox_object_t)(unsafe.Pointer(uintptr(a1))), item)
	}
	return v4
}
func sub_53AAB0(a1 int32) {
	var (
		v2 int32
		v3 int16
	)
	if a1 != 0 {
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))))
		if v2&0x1000 != 0 {
			nox_xxx_aud_501960(831, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
		} else {
			v3 = int16(*(*uint16)(unsafe.Pointer(uintptr(a1 + 24))))
			if int32(v3)&0x10 != 0 {
				nox_xxx_aud_501960(843, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
			} else if int32(v3)&8 != 0 {
				nox_xxx_aud_501960(845, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
			}
		}
	}
}
func nox_xxx_dropWeapon_53AB10(a1 int32, a2 *uint32, a3 *int32) int32 {
	if nox_xxx_dropDefault_4ED290((*nox_object_t)(unsafe.Pointer(uintptr(a1))), (*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a2)))))), (*float2)(unsafe.Pointer(a3))) != 1 {
		return 0
	}
	sub_53AAB0(int32(uintptr(unsafe.Pointer(a2))))
	if !nox_common_gameFlags_check_40A5C0(2048) && !nox_common_gameFlags_check_40A5C0(4096) {
		if sub_409F40(2) != 0 {
			nox_xxx_unitSetDecayTime_511660((*nox_object_t)(unsafe.Pointer(a2)), int32(gameFPS()*25))
		}
	}
	return 1
}
func sub_53AB90(a1 int32, a2 int32) {
	var (
		v2 int32
		v3 int32
	)
	if a1 != 0 {
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
		if a2 != 0 {
			if nox_xxx_playerClassCanUseItem_57B3D0((*nox_object_t)(unsafe.Pointer(uintptr(a2))), int8(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 2251))))) == 0 || (func() bool {
				*((*uint8)(unsafe.Pointer(&v3))) = uint8(int8(bool2int32(nox_xxx_playerCheckStrength_4F3180((*nox_object_t)(unsafe.Pointer(uintptr(a1))), (*nox_object_t)(unsafe.Pointer(uintptr(a2)))))))
				return v3 == 0
			}()) {
				nox_xxx_netSendSecondaryWeapon_4D9670(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 2064)))), nil, 1)
			}
		}
		*(*uint32)(unsafe.Pointer(uintptr(v2 + 108))) = uint32(a2)
	}
}
func nox_xxx_updateDoor_53AC50(a1 int32) int8 {
	var (
		v1 *uint32
		v2 int32
		v3 int32
		v4 int32
		v5 int32
		v6 int32
		v7 int32
	)
	v1 = *(**uint32)(unsafe.Pointer(uintptr(a1 + 748)))
	v2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*2)))
	v3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*1)))
	if v2 == v3 {
		if *(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*3)) != uint32(v3) {
			v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12))))
			if v4&4 != 0 {
				if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 24))))&8 != 0 {
					nox_xxx_aud_501960(245, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
				} else {
					nox_xxx_aud_501960(241, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
				}
			} else if v4&1 != 0 {
				nox_xxx_aud_501960(247, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
			} else if v4&0x1000 != 0 {
				nox_xxx_aud_501960(1014, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
			} else {
				nox_xxx_aud_501960(237, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
			}
			goto LABEL_24
		}
		if v2 == v3 {
			goto LABEL_24
		}
	}
	if *(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*3)) == uint32(v3) {
		v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12))))
		if v5&4 != 0 {
			if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 24))))&8 != 0 {
				nox_xxx_aud_501960(246, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
			} else {
				nox_xxx_aud_501960(243, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
			}
		} else if v5&1 != 0 {
			nox_xxx_aud_501960(248, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
		} else if v5&0x1000 != 0 {
			nox_xxx_aud_501960(1015, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
		} else {
			nox_xxx_aud_501960(239, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
		}
		nox_xxx_unitRemoveFromUpdatable_4DA920((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
	}
LABEL_24:
	v6 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*3)))
	if *(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*2)) != uint32(v6) {
		nox_xxx_unitNeedSync_4E44F0((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
		v6 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*3)))
		*(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*2)) = uint32(v6)
	}
	if *(*uint32)(unsafe.Pointer(uintptr(a1 + 16)))&0x1000000 != 0 {
		if (gameFrame() - *(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*11))) > uint32(int32(gameFPS())>>1) {
			v7 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*1)))
			if v6 != v7 {
				if v6-v7+(func() int32 {
					if v6-v7 < 0 {
						return 0x20
					}
					return 0
				}()) >= 16 {
					sub_548860(a1, 2)
				} else {
					sub_548860(a1, -2)
				}
				*((*uint8)(unsafe.Pointer(&v6))) = uint8(nox_xxx_unitHasCollideOrUpdateFn_537610((*nox_object_t)(unsafe.Pointer(uintptr(a1)))))
			}
		}
	}
	return int8(v6)
}
func nox_xxx_updateSpark_53ADC0(a1 int32) {
	var (
		v1 int32
		v2 int32
	)
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 4))))
	if v2 <= 0 {
		nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
	} else {
		*(*uint32)(unsafe.Pointer(uintptr(v1 + 4))) = uint32(v2 - 1)
		if *(*uint32)(unsafe.Pointer(uintptr(v1 + 12))) == 4 {
			*(*uint32)(unsafe.Pointer(uintptr(a1 + 112))) = 1065353216
		} else {
			*(*uint32)(unsafe.Pointer(uintptr(a1 + 112))) = 1064514355
		}
	}
}
func nox_xxx_updateProjTrail_53AEC0(a1 int32) *float32 {
	var (
		v1     *float32
		v2     int32
		v3     int32
		v4     float64
		v5     float64
		v6     int32
		result *float32
		v8     float32
		v9     float32
		v10    float32
		v11    float32
		v12    float32
		v13    float32
		v14    float32
		v15    float32
		v16    int32
		v17    float32
		v18    int32
		v19    float32
		v20    float32
	)
	v1 = (*float32)(unsafe.Pointer(uintptr(a1)))
	v2 = int32(*(*int16)(unsafe.Pointer(uintptr(a1 + 124)))) * 8
	v3 = 0
	v4 = float64(*mem_getFloatPtr(0x587000, uint32(v2)+194136) * *(*float32)(unsafe.Pointer(uintptr(a1 + 544))))
	v18 = 0
	*(*float32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(float32(0))*22)) = float32(v4*0.25 + float64(*(*float32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(float32(0))*22))))
	*(*float32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(float32(0))*23)) = float32(float64(*mem_getFloatPtr(0x587000, uint32(v2)+194140)**(*float32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(float32(0))*136)))*0.25 + float64(*(*float32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(float32(0))*23))))
	v14 = float32(float64(*(*float32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(float32(0))*14))-*(*float32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(float32(0))*18))) * 0.25)
	v15 = float32(float64(*(*float32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(float32(0))*15))-*(*float32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(float32(0))*19))) * 0.25)
	for {
		v5 = float64(v18)
		v19 = float32(v5)
		v12 = float32(v5*float64(v14) + float64(*(*float32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(float32(0))*18))))
		v16 = nox_float2int(v12)
		v13 = v19*v15 + *(*float32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(float32(0))*19))
		v6 = 2
		v20 = float32(float64(nox_float2int(v13)))
		v17 = float32(float64(v16))
		for {
			v11 = float32(nox_common_randomFloat_416030(-2.0, 2.0))
			v10 = float32(nox_common_randomFloat_416030(-2.0, 2.0))
			v9 = float32(nox_common_randomFloat_416030(-4.0, 4.0) + float64(v20))
			v8 = float32(nox_common_randomFloat_416030(-4.0, 4.0) + float64(v17))
			result = nox_xxx_createSpark_54FD80(v8, v9, 1, 6, v10, v11, 0.0, int32(uintptr(unsafe.Pointer(v1))))
			v6--
			if v6 == 0 {
				break
			}
		}
		v18 = func() int32 {
			p := &v3
			*p++
			return *p
		}()
		if v3 >= 4 {
			break
		}
	}
	return result
}
func nox_xxx_updatePush_53B030(a1 int32) {
	nox_xxx_mapPushUnitsAround_52E040(unsafe.Pointer(uintptr(a1+56)), **(**float32)(unsafe.Pointer(uintptr(a1 + 748))), 0, *(*float32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + 8))), nil, 0, 0)
}
func nox_xxx_updateToggle_53B060(a1 *uint32) int8 {
	var (
		v1 *uint32
		v2 uint32
	)
	v1 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*187)))))
	if (*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*4)) & 0x1000000) == 0 {
		v2 = *v1
		*((*uint8)(unsafe.Pointer(&v2))) = uint8(*v1 & 0xF6)
		*v1 = v2
		return int8(uint8(v2))
	}
	if (int32(*(*uint8)(unsafe.Pointer(v1))) & 8) == 0 {
		*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v1))), 8))) = 0
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*34)) = gameFrame()
		nox_xxx_servMarkObjAnimFrame_4E4880(int32(uintptr(unsafe.Pointer(a1))), 0)
	}
	if int32(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v1))), 8)))) != 0 {
		if int32(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v1))), 8)))) != 1 {
			if int32(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v1))), 8)))) == 3 && gameFrame() > *(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*34)) && (int32(*(*uint8)(unsafe.Pointer(v1)))&1) == 0 {
				*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v1))), 8))) = 1
			}
			goto LABEL_18
		}
		if gameFrame() > *(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*34)) && int32(*(*uint8)(unsafe.Pointer(v1)))&1 != 0 {
			nox_xxx_aud_501960(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*10))), (*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), 0, 0)
			nox_xxx_servMarkObjAnimFrame_4E4880(int32(uintptr(unsafe.Pointer(a1))), 0)
			nox_xxx_scriptCallByEventBlock_502490(unsafe.Pointer((*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(v1))), 4*7))), nil, unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1))))), 11)
			if (*v1 & 2) != 0 {
				*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v1))), 8))) = 5
			} else {
				*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v1))), 8))) = 0
			}
			*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*34)) = gameFrame() + gameFPS()
			goto LABEL_18
		}
	} else if gameFrame() > *(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*34)) && int32(*(*uint8)(unsafe.Pointer(v1)))&1 != 0 {
		nox_xxx_aud_501960(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*9))), (*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), 0, 0)
		nox_xxx_servMarkObjAnimFrame_4E4880(int32(uintptr(unsafe.Pointer(a1))), 1)
		nox_xxx_scriptCallByEventBlock_502490(unsafe.Pointer((*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(v1))), 4*5))), unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*1)))), unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1))))), 10)
		*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v1))), 8))) = 3
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*34)) = gameFrame() + gameFPS()
		goto LABEL_18
	}
LABEL_18:
	v2 = *v1
	if *v1&1 != 0 {
		*((*uint8)(unsafe.Pointer(&v2))) = uint8(v2 | 4)
	} else {
		*((*uint8)(unsafe.Pointer(&v2))) = uint8(v2 & 0xFB)
	}
	*v1 = v2
	*v1 = v2&0xFFFFFFFE | 8
	return int8(uint8(v2))
}
func nox_xxx_updateTrigger_53B1B0(a1 int32) int8 {
	var (
		v1 *int32
		v2 int32
		v3 int32
		v4 int32
	)
	v1 = *(**int32)(unsafe.Pointer(uintptr(a1 + 748)))
	if *memmap.PtrUint32(0x5D4594, 2488680) == 0 {
		*memmap.PtrUint32(0x5D4594, 2488680) = uint32(nox_xxx_getNameId_4E3AA0(internCStr("Trigger")))
		*memmap.PtrUint32(0x5D4594, 2488676) = uint32(nox_xxx_getNameId_4E3AA0(internCStr("PressurePlate")))
	}
	v2 = int32(*(*uint16)(unsafe.Pointer(uintptr(a1 + 4))))
	if uint32(uint16(int16(v2))) == *memmap.PtrUint32(0x5D4594, 2488680) || uint32(v2) == *memmap.PtrUint32(0x5D4594, 2488676) {
		v3 = 0
	} else {
		v3 = int32(gameFPS())
	}
	if *(*uint32)(unsafe.Pointer(uintptr(a1 + 16)))&0x1000000 != 0 {
		if (int32(*(*uint8)(unsafe.Pointer(v1))) & 8) == 0 {
			*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v1))), 8))) = 0
			*(*uint32)(unsafe.Pointer(uintptr(a1 + 136))) = gameFrame()
			nox_xxx_servMarkObjAnimFrame_4E4880(a1, 0)
		}
		if int32(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v1))), 8)))) != 0 {
			if int32(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v1))), 8)))) == 1 && gameFrame() > uint32(*(*int32)(unsafe.Pointer(uintptr(a1 + 136)))) && (int32(*(*uint8)(unsafe.Pointer(v1)))&1) == 0 {
				nox_xxx_aud_501960(*(*int32)(unsafe.Add(unsafe.Pointer(v1), 4*10)), (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
				nox_xxx_servMarkObjAnimFrame_4E4880(a1, 0)
				nox_xxx_scriptCallByEventBlock_502490(unsafe.Pointer((*int32)(unsafe.Add(unsafe.Pointer(v1), 4*7))), nil, unsafe.Pointer(uintptr(a1)), 9)
				if (*v1 & 2) != 0 {
					*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v1))), 8))) = 5
				} else {
					*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v1))), 8))) = 0
				}
			}
		} else if int32(*(*uint8)(unsafe.Pointer(v1)))&1 != 0 {
			nox_xxx_aud_501960(*(*int32)(unsafe.Add(unsafe.Pointer(v1), 4*9)), (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
			nox_xxx_servMarkObjAnimFrame_4E4880(a1, 1)
			nox_xxx_scriptCallByEventBlock_502490(unsafe.Pointer((*int32)(unsafe.Add(unsafe.Pointer(v1), 4*5))), unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(v1), 4*1)))), unsafe.Pointer(uintptr(a1)), 8)
			*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v1))), 8))) = 1
			*(*uint32)(unsafe.Pointer(uintptr(a1 + 136))) = uint32(v3) + gameFrame()
		}
		v4 = *v1
		if *v1&1 != 0 {
			*((*uint8)(unsafe.Pointer(&v4))) = uint8(int8(v4 | 4))
		} else {
			*((*uint8)(unsafe.Pointer(&v4))) = uint8(int8(v4 & 0xFB))
		}
		*v1 = v4
		*((*uint8)(unsafe.Pointer(&v4))) = uint8(int8(v4&0xFE | 8))
		*v1 = v4
	} else {
		v4 = *v1
		*((*uint8)(unsafe.Pointer(&v4))) = uint8(int8(*v1 & 0xF6))
		*v1 = v4
	}
	return int8(v4)
}
func sub_53B300(a1 int32) int8 {
	var v1 int32
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))))
	if uint32(v1)&0x1000000 != 0 {
		*((*uint8)(unsafe.Pointer(&v1))) = uint8(int8(v1 & 0xBF))
	} else {
		*((*uint8)(unsafe.Pointer(&v1))) = uint8(int8(v1 | 0x40))
	}
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))) = uint32(v1)
	return int8(v1)
}
func nox_xxx_updateSwitch_53B320(a1 *uint32) int8 {
	var v1 int32
	v1 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*4)))
	if uint32(v1)&0x1000000 != 0 {
		if *(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*174)) != 0 && v1&0x40 != 0 {
			*((*uint8)(unsafe.Pointer(&v1))) = uint8(int8(v1 & 0xBF))
			*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*4)) = uint32(v1)
			*((*uint8)(unsafe.Pointer(&v1))) = uint8(nox_xxx_unitHasCollideOrUpdateFn_537610((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1))))))))
		} else {
			*((*uint8)(unsafe.Pointer(&v1))) = uint8(int8(v1 & 0xBF))
			*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*4)) = uint32(v1)
		}
	} else {
		if *(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*33)) == 0 {
			nox_xxx_unitNeedSync_4E44F0((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))))
			*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*33)) = 1
		}
		v1 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*4)))
		*((*uint8)(unsafe.Pointer(&v1))) = uint8(int8(v1 | 0x40))
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*4)) = uint32(v1)
	}
	return int8(v1)
}
func nox_xxx_updateElevatorShaft_53B380(a1 int32) int8 {
	var (
		v1 int32
		v2 int32
		v3 int32
	)
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 4))))
	if v2 != 0 {
		nox_xxx_unitHasCollideOrUpdateFn_537610((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 4))) + 748))))
		if int32(*(*uint8)(unsafe.Pointer(uintptr(v3 + 12)))) == 1 {
			if *(*uint32)(unsafe.Pointer(uintptr(v3 + 16))) <= 32 {
				nox_xxx_unitsGetInCircle_517F90((*float2)(unsafe.Pointer(uintptr(a1+56))), 64.0, unsafe.Pointer(funAddr(nox_xxx_fnElevatorShaft_53B410)), unsafe.Pointer(uintptr(a1)))
			}
			if int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 12)))) != int32(*(*uint8)(unsafe.Pointer(uintptr(v3 + 12)))) {
				nox_xxx_elevatorAud_53B490(a1, 0)
			}
		} else if int32(*(*uint8)(unsafe.Pointer(uintptr(v3 + 12)))) == 3 && int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 12)))) != int32(*(*uint8)(unsafe.Pointer(uintptr(v3 + 12)))) {
			nox_xxx_elevatorAud_53B490(a1, 1)
			*((*uint8)(unsafe.Pointer(&v2))) = *(*uint8)(unsafe.Pointer(uintptr(v3 + 12)))
			*(*uint8)(unsafe.Pointer(uintptr(v1 + 12))) = uint8(int8(v2))
			return int8(v2)
		}
		*((*uint8)(unsafe.Pointer(&v2))) = *(*uint8)(unsafe.Pointer(uintptr(v3 + 12)))
		*(*uint8)(unsafe.Pointer(uintptr(v1 + 12))) = uint8(int8(v2))
	}
	return int8(v2)
}
func nox_xxx_fnElevatorShaft_53B410(a1 int32, a2 int32) {
	var (
		v2 int32
		v3 int32
		v4 float32
		v5 float32
	)
	if nox_xxx_map_57B850((*float2)(unsafe.Pointer(uintptr(a2+56))), (*float32)(unsafe.Pointer(uintptr(a2+172))), (*float2)(unsafe.Pointer(uintptr(a1+56)))) != 0 {
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 748))))
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 4))) + 748))))
		v4 = float32(float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 104)))) + 64.0 - float64(*(*int32)(unsafe.Pointer(uintptr(v3 + 16)))))
		if sub_419A10(v4) < 10.0 {
			nox_xxx_unitMove_4E7010((*nox_object_t)(unsafe.Pointer(uintptr(a1))), (*float2)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 4)))+56))))
			v5 = float32(float64(*(*int32)(unsafe.Pointer(uintptr(v3 + 16)))))
			nox_xxx_unitRaise_4E46F0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), v5)
		}
	}
}
func nox_xxx_elevatorAud_53B490(a1 int32, a2 int32) {
	var result *uint32
	_ = result
	var v3 int32
	var v4 int32
	var v5 int32
	var v6 int32
	if a2 != 0 {
		result = (*uint32)(unsafe.Pointer(uintptr(a1)))
		v3 = int32(*(*uint16)(unsafe.Pointer(uintptr(a1 + 24))))
		if int32(uint16(int16(v3))) == 8 {
			nox_xxx_aud_501960(257, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
		} else if v3 == 16 {
			v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12))))
			if v4&0x20 != 0 {
				nox_xxx_aud_501960(253, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
			} else if v4&0x40 != 0 {
				nox_xxx_aud_501960(259, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
			} else {
				nox_xxx_aud_501960(251, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
			}
		} else if v3 == 32 {
			if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 12))))&2 != 0 {
				nox_xxx_aud_501960(math.MaxUint8, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
			} else {
				nox_xxx_aud_501960(249, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
			}
		}
	} else {
		result = (*uint32)(unsafe.Pointer(uintptr(a1)))
		v5 = int32(*(*uint16)(unsafe.Pointer(uintptr(a1 + 24))))
		if int32(uint16(int16(v5))) == 8 {
			nox_xxx_aud_501960(258, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
		} else if v5 == 16 {
			v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12))))
			if v6&0x20 != 0 {
				nox_xxx_aud_501960(254, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
			} else if v6&0x40 != 0 {
				nox_xxx_aud_501960(260, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
			} else {
				nox_xxx_aud_501960(252, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
			}
		} else if v5 == 32 {
			if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 12))))&2 != 0 {
				nox_xxx_aud_501960(256, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
			} else {
				nox_xxx_aud_501960(250, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
			}
		}
	}
}
func nox_xxx_updateElevator_53B5D0(a1 *uint32) {
	var (
		v1 int32
		v2 int32
		v3 int32
		v4 int32
	)
	v1 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*187)))
	switch *(*uint8)(unsafe.Pointer(uintptr(v1 + 12))) {
	case 0:
		if *(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*4))&0x1000000 != 0 {
			if (gameFrame() - *(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*34))) > uint32(int32(gameFPS())) {
				*(*uint8)(unsafe.Pointer(uintptr(v1 + 12))) = 3
				nox_xxx_elevatorAud_53B490(int32(uintptr(unsafe.Pointer(a1))), 1)
			}
		}
	case 1:
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 16))))
		if v2 > 0 {
			*(*uint32)(unsafe.Pointer(uintptr(v1 + 16))) = uint32(v2 - 2)
		} else {
			*(*uint8)(unsafe.Pointer(uintptr(v1 + 12))) = 0
			*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*34)) = gameFrame()
		}
		nox_xxx_unitNeedSync_4E44F0((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))))
		if *(*uint32)(unsafe.Pointer(uintptr(v1 + 4))) != 0 {
			nox_xxx_unitNeedSync_4E44F0((*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 4)))))))
		}
		if float64(*(*int32)(unsafe.Pointer(uintptr(v1 + 16)))) <= 20.0 {
			v3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*4)))
			*((*uint8)(unsafe.Pointer(&v3))) = uint8(int8(v3 | 0x10))
			*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*4)) = uint32(v3)
		}
	case 2:
		if *(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*4))&0x1000000 != 0 && (gameFrame()-*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*34))) > uint32(int32(gameFPS())) {
			*(*uint8)(unsafe.Pointer(uintptr(v1 + 12))) = 1
			nox_xxx_elevatorAud_53B490(int32(uintptr(unsafe.Pointer(a1))), 0)
		}
	case 3:
		*(*uint32)(unsafe.Pointer(uintptr(v1 + 16))) += 2
		nox_xxx_unitNeedSync_4E44F0((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))))
		if *(*uint32)(unsafe.Pointer(uintptr(v1 + 4))) != 0 {
			nox_xxx_unitNeedSync_4E44F0((*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 4)))))))
		}
		if float64(*(*int32)(unsafe.Pointer(uintptr(v1 + 16)))) >= 20.0 {
			v4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*4)))
			*((*uint8)(unsafe.Pointer(&v4))) = uint8(int8(v4 & 0xEF))
			*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*4)) = uint32(v4)
		}
		if float64(*(*int32)(unsafe.Pointer(uintptr(v1 + 16)))) >= 32.0 {
			nox_xxx_unitsGetInCircle_517F90((*float2)(unsafe.Add(unsafe.Pointer((*float2)(unsafe.Pointer(a1))), unsafe.Sizeof(float2{})*7)), 64.0, unsafe.Pointer(funAddr(nox_xxx_elevatorFn_53B750)), unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1))))))
		}
		if *(*uint32)(unsafe.Pointer(uintptr(v1 + 16))) >= 64 {
			*(*uint32)(unsafe.Pointer(uintptr(v1 + 16))) = 64
			*(*uint8)(unsafe.Pointer(uintptr(v1 + 12))) = 2
			*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*34)) = gameFrame()
		}
	default:
		return
	}
}
func nox_xxx_elevatorFn_53B750(a1 int32, a2 int32) {
	var (
		v2 int32
		v3 int32
		v4 int32
		v5 float64
		v6 float32
		v7 float32
	)
	if nox_xxx_map_57B850((*float2)(unsafe.Pointer(uintptr(a2+56))), (*float32)(unsafe.Pointer(uintptr(a2+172))), (*float2)(unsafe.Pointer(uintptr(a1+56)))) != 0 {
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 748))))
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 4))))
		if v3 != 0 {
			v6 = float32(float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 104)))) - float64(*(*int32)(unsafe.Pointer(uintptr(v2 + 16)))))
			if sub_419A10(v6) < 10.0 {
				v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 172))))
				if v4 == 3 {
					if float64(*(*float32)(unsafe.Pointer(uintptr(v3 + 184)))) < float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 184)))) || float64(*(*float32)(unsafe.Pointer(uintptr(v3 + 188)))) < float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 188)))) {
						*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))) &= 0xFFEFFFFF
						nox_xxx_unitRaise_4E46F0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0.0)
						return
					}
				} else if v4 == 2 {
					v5 = float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 176))) + *(*float32)(unsafe.Pointer(uintptr(a1 + 176))))
					if v5 > float64(*(*float32)(unsafe.Pointer(uintptr(v3 + 184)))) || v5 > float64(*(*float32)(unsafe.Pointer(uintptr(v3 + 188)))) {
						*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))) &= 0xFFEFFFFF
						nox_xxx_unitRaise_4E46F0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0.0)
						return
					}
				}
				nox_xxx_unitMove_4E7010((*nox_object_t)(unsafe.Pointer(uintptr(a1))), (*float2)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 4)))+56))))
				v7 = float32(float64(int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 16))) - 64)))
				nox_xxx_unitRaise_4E46F0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), v7)
				return
			}
		}
	}
}
func nox_xxx_updatePhantomPlayer_53B860(a1 int32) {
	var (
		v1 *float32
		v2 float64
		v3 float64
		v4 int32
		v5 int16
		v6 bool
	)
	_ = v6
	var v7 float32
	v1 = *(**float32)(unsafe.Pointer(uintptr(a1 + 748)))
	v2 = float64(*(*float32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(v1)) + 56))) - *(*float32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(float32(0))*1)))
	v3 = float64(*(*float32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(v1)) + 60))) - *(*float32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(float32(0))*2)))
	v7 = float32(v3)
	if v3*float64(v7)+v2*v2 <= 160000.0 {
		*(*float32)(unsafe.Pointer(uintptr(a1 + 64))) = float32(float64(*(*float32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(v1)) + 56)))) - v2)
		*(*float32)(unsafe.Pointer(uintptr(a1 + 68))) = *(*float32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(v1)) + 60))) - v7
		v4 = int32(*(*uint32)(unsafe.Pointer(v1)))
		v5 = int16(int32(*(*uint16)(unsafe.Pointer(uintptr(v4 + 124)))) + 128)
		v6 = int32(int16(int32(*(*uint16)(unsafe.Pointer(uintptr(v4 + 124))))-128)) < 0
		*(*uint16)(unsafe.Pointer(uintptr(a1 + 126))) = uint16(v5)
		for int32(*(*uint16)(unsafe.Pointer(uintptr(a1 + 126)))) >= 256 {
			*(*uint16)(unsafe.Pointer(uintptr(a1 + 126))) -= 256
		}
	} else {
		nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
	}
}
func nox_xxx_updateLifetime_53B8F0(unit int32) {
	var deleteOverride func(int32)
	if (gameFrame() - *(*uint32)(unsafe.Pointer(uintptr(unit + 128)))) > **(**uint32)(unsafe.Pointer(uintptr(unit + 748))) {
		deleteOverride = *(*func(int32))(unsafe.Pointer(uintptr(unit + 724)))
		*(*uint32)(unsafe.Pointer(uintptr(unit + 16))) |= 0x8000
		if deleteOverride != nil {
			deleteOverride(unit)
		} else {
			nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(unit))))
		}
	}
}
func nox_xxx_spellFlyUpdate_53B940(a1 int32) {
	var (
		v1  int32
		v2  *int32
		v3  float64
		v4  int32
		v5  int32
		v6  int32
		v7  int32
		v8  int32
		v9  int32
		v10 float64
		v11 float64
		v12 float64
		v13 float64
		v14 float64
		v15 float32
		v16 int32
		v17 float32
		v18 float32
	)
	v1 = a1
	v2 = *(**int32)(unsafe.Pointer(uintptr(a1 + 748)))
	if *(*int32)(unsafe.Add(unsafe.Pointer(v2), 4*1)) != 0 {
		v3 = nox_xxx_gamedataGetFloat_419D40(internCStr("TargetedSpellLifetime"))
	} else {
		v3 = nox_xxx_gamedataGetFloat_419D40(internCStr("UnTargetedSpellLifetime"))
	}
	v15 = float32(v3)
	if gameFrame()-*(*uint32)(unsafe.Pointer(uintptr(a1 + 128))) <= uint32(nox_float2int(v15)) {
		if *v2 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(*v2 + 16))))&0x20 != 0 {
			*v2 = 0
		}
		v4 = *(*int32)(unsafe.Add(unsafe.Pointer(v2), 4*2))
		if v4 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(v4 + 16))))&0x20 != 0 {
			*(*int32)(unsafe.Add(unsafe.Pointer(v2), 4*2)) = 0
		}
		v5 = *(*int32)(unsafe.Add(unsafe.Pointer(v2), 4*1))
		if v5 != 0 {
			v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 16))))
			if v6&0x20 != 0 || (v6&0x8000) != 0 {
				*(*int32)(unsafe.Add(unsafe.Pointer(v2), 4*1)) = 0
			}
		}
		if *(*int32)(unsafe.Add(unsafe.Pointer(v2), 4*1)) == 0 {
			v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 136))))
			if (gameFrame()-uint32(v7)) > uint32(int32(gameFPS())>>2) || uint32(v7) == *(*uint32)(unsafe.Pointer(uintptr(a1 + 128))) {
				v16 = *v2
				v8 = int32(nox_xxx_spellFlags_424A70(*(*int32)(unsafe.Add(unsafe.Pointer(v2), 4*3))))
				*(*int32)(unsafe.Add(unsafe.Pointer(v2), 4*1)) = int32(uintptr(unsafe.Pointer(nox_xxx_spellFlySearchTarget_540610(nil, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), v8, 600.0, 0, (*nox_object_t)(unsafe.Pointer(uintptr(v16)))))))
				*(*uint32)(unsafe.Pointer(uintptr(a1 + 136))) = gameFrame()
			}
		}
		v9 = *(*int32)(unsafe.Add(unsafe.Pointer(v2), 4*1))
		if v9 != 0 {
			v18 = *(*float32)(unsafe.Pointer(uintptr(v9 + 56))) - *(*float32)(unsafe.Pointer(uintptr(a1 + 56)))
			v10 = float64(*(*float32)(unsafe.Pointer(uintptr(v9 + 60))) - *(*float32)(unsafe.Pointer(uintptr(v1 + 60))))
			v17 = float32(v10)
			v11 = float64(nox_double2float(math.Sqrt(v10*float64(v17)+float64(v18*v18)) + *(*float64)(unsafe.Pointer(&qword_581450_10176))))
			v12 = float64(v18 * *(*float32)(unsafe.Pointer(uintptr(v1 + 544))))
			*(*uint32)(unsafe.Pointer(uintptr(v1 + 112))) = 1063675494
			*(*float32)(unsafe.Pointer(uintptr(v1 + 88))) = float32(v12 / v11)
			*(*float32)(unsafe.Pointer(uintptr(v1 + 92))) = float32(float64(v17**(*float32)(unsafe.Pointer(uintptr(v1 + 544)))) / v11)
		} else {
			v13 = float64(nox_double2float(math.Sqrt(float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 80)))**(*float32)(unsafe.Pointer(uintptr(a1 + 80)))+*(*float32)(unsafe.Pointer(uintptr(a1 + 84)))**(*float32)(unsafe.Pointer(uintptr(a1 + 84))))) + *(*float64)(unsafe.Pointer(&qword_581450_10176))))
			v14 = float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 544))) * *(*float32)(unsafe.Pointer(uintptr(a1 + 80))))
			*(*uint32)(unsafe.Pointer(uintptr(a1 + 112))) = 1063675494
			*(*float32)(unsafe.Pointer(uintptr(a1 + 88))) = float32(v14 / v13)
			*(*float32)(unsafe.Pointer(uintptr(a1 + 92))) = float32(float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 544)))**(*float32)(unsafe.Pointer(uintptr(a1 + 84)))) / v13)
		}
	} else {
		nox_xxx_netSendPointFx_522FF0(-106, (*float2)(unsafe.Pointer(uintptr(a1+56))))
		sub_4E71F0((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
	}
}
func nox_xxx_updateAntiSpellProj_53BB00(a1 int32) {
	var (
		v1  int32
		v2  int32
		v3  int32
		v4  uint32
		v5  int32
		v6  int32
		v7  float64
		v8  float64
		v9  float64
		v10 float64
		v11 float64
		v12 float32
		v13 float32
		v14 int32
		v15 float32
		v16 float32
		v17 float32
		v18 int32
	)
	v1 = int32(gameFrame())
	v2 = a1
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	v4 = gameFPS()
	if gameFrame()-*(*uint32)(unsafe.Pointer(uintptr(a1 + 128))) <= (gameFPS() * 5) {
		v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 4))))
		if v5 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(v5 + 16))))&0x20 != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(v3 + 4))) = 0
			v1 = int32(gameFrame())
			v4 = gameFPS()
		}
		if *(*uint32)(unsafe.Pointer(uintptr(v3 + 4))) == 0 && uint32(v1)-*(*uint32)(unsafe.Pointer(uintptr(a1 + 136))) > v4>>2 {
			*memmap.PtrUint32(0x5D4594, 2488668) = 0
			*memmap.PtrUint32(0x5D4594, 2488672) = 1287568416
			nox_xxx_getMissilesInCircle_518170((*float2)(unsafe.Pointer(uintptr(a1+56))), 600.0, unsafe.Pointer(funAddr(sub_53BD10)), (*nox_object_t)(unsafe.Pointer(uintptr(a1))))
			*(*uint32)(unsafe.Pointer(uintptr(v3 + 4))) = *memmap.PtrUint32(0x5D4594, 2488668)
			*(*uint32)(unsafe.Pointer(uintptr(a1 + 136))) = gameFrame()
		}
		v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 4))))
		if v6 != 0 {
			v7 = float64(*(*float32)(unsafe.Pointer(uintptr(v6 + 56))) - *(*float32)(unsafe.Pointer(uintptr(a1 + 56))))
			v8 = float64(*(*float32)(unsafe.Pointer(uintptr(v6 + 60))) - *(*float32)(unsafe.Pointer(uintptr(a1 + 60))))
			v9 = v7
			v10 = math.Sqrt(v8*v8+v7*v7) + 0.1
			*(*float32)(unsafe.Pointer(uintptr(a1 + 88))) = float32(v9 * float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 544)))) / v10)
			*(*float32)(unsafe.Pointer(uintptr(a1 + 92))) = float32(v8 * float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 544)))) / v10)
			if v10 < 10.0 {
				nox_xxx_aud_501960(20, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
				nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
				nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 4)))))))
			}
		} else {
			v11 = math.Sqrt(float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 80)))**(*float32)(unsafe.Pointer(uintptr(a1 + 80)))+*(*float32)(unsafe.Pointer(uintptr(a1 + 84)))**(*float32)(unsafe.Pointer(uintptr(a1 + 84))))) + 0.1
			*(*float32)(unsafe.Pointer(uintptr(a1 + 88))) = float32(float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 544)))**(*float32)(unsafe.Pointer(uintptr(a1 + 80)))) / v11)
			*(*float32)(unsafe.Pointer(uintptr(a1 + 92))) = float32(float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 544)))**(*float32)(unsafe.Pointer(uintptr(a1 + 84)))) / v11)
		}
		v17 = *(*float32)(unsafe.Pointer(uintptr(a1 + 56)))
		v18 = *(*int32)(unsafe.Pointer(uintptr(a1 + 60)))
		v16 = float32(nox_common_randomFloat_416030(-2.0, 2.0))
		v15 = float32(nox_common_randomFloat_416030(-2.0, 2.0))
		v14 = nox_common_randomInt_415FA0(15, 30)
		v13 = float32(nox_common_randomFloat_416030(-4.0, 4.0) + float64(*(*float32)(unsafe.Pointer(&v18))))
		v12 = float32(nox_common_randomFloat_416030(-4.0, 4.0) + float64(v17))
		nox_xxx_createSpark_54FD80(v12, v13, 3, v14, v15, v16, 0.0, v2)
	} else {
		nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
	}
}
func sub_53BD10(a1 int32, a2 int32) {
	var (
		v2 float64
		v3 float64
		v4 float64
	)
	if (int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 16))))&0x20) == 0 && ((int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&1) == 0 || int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 12))))&2 != 0) && a2 != a1 && (*(*uint32)(unsafe.Pointer(uintptr(a2 + 508))) == 0 || nox_xxx_unitHasThatParent_4EC4F0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), (*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 508))))))) == 0) {
		if nox_xxx_unitCanInteractWith_5370E0((*nox_object_t)(unsafe.Pointer(uintptr(a2))), (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0) != 0 {
			v2 = float64(*(*float32)(unsafe.Pointer(uintptr(a2 + 56))) - *(*float32)(unsafe.Pointer(uintptr(a1 + 56))))
			v3 = float64(*(*float32)(unsafe.Pointer(uintptr(a2 + 60))) - *(*float32)(unsafe.Pointer(uintptr(a1 + 60))))
			v4 = v3*v3 + v2*v2
			if v4 < float64(*mem_getFloatPtr(0x5D4594, 2488672)) {
				*mem_getFloatPtr(0x5D4594, 2488672) = float32(v4)
				*memmap.PtrUint32(0x5D4594, 2488668) = uint32(a1)
			}
		}
	}
}
func nox_xxx_updateMagicMissile_53BDA0(a1 int32) int32 {
	var (
		v1     *int32
		v2     int32
		v3     int32
		v4     int32
		v5     int32
		v6     int32
		v7     int16
		v8     int16
		result int32
		v10    float64
	)
	v1 = *(**int32)(unsafe.Pointer(uintptr(a1 + 748)))
	v2 = *v1
	if (int32(*(*uint8)(unsafe.Pointer(uintptr(*v1 + 16))))&0x20) != 0 || gameFrame()-*(*uint32)(unsafe.Pointer(uintptr(a1 + 128))) > (gameFPS()*3) {
		return (*(*func(int32, uint32, uint32) int32)(unsafe.Pointer(uintptr(a1 + 696))))(a1, 0, 0)
	}
	v3 = *(*int32)(unsafe.Add(unsafe.Pointer(v1), 4*1))
	if v3 != 0 {
		v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 16))))
		if (v4&0x20) != 0 || (v4&0x8000) != 0 {
			*(*int32)(unsafe.Add(unsafe.Pointer(v1), 4*1)) = 0
		}
	}
	if *(*int32)(unsafe.Add(unsafe.Pointer(v1), 4*1)) == 0 && (int32(uint8(gameFrame()))&7) == 0 {
		v5 = int32(uintptr(unsafe.Pointer(nox_xxx_spellFlySearchTarget_540610(nil, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 32, 600.0, 0, (*nox_object_t)(unsafe.Pointer(uintptr(v2)))))))
		*(*int32)(unsafe.Add(unsafe.Pointer(v1), 4*1)) = v5
		if uint32(v5) == *(*uint32)(unsafe.Pointer(uintptr(a1 + 508))) {
			*(*int32)(unsafe.Add(unsafe.Pointer(v1), 4*1)) = 0
		}
	}
	v6 = *(*int32)(unsafe.Add(unsafe.Pointer(v1), 4*1))
	if v6 != 0 {
		v7 = int16(*(*uint16)(unsafe.Pointer(uintptr(a1 + 126))))
		if float64((*(*float32)(unsafe.Pointer(uintptr(v6 + 60)))-*(*float32)(unsafe.Pointer(uintptr(a1 + 60))))**mem_getFloatPtr(0x587000, uint32(int32(v7)*8)+194136)-(*(*float32)(unsafe.Pointer(uintptr(v6 + 56)))-*(*float32)(unsafe.Pointer(uintptr(a1 + 56))))**mem_getFloatPtr(0x587000, uint32(int32(v7)*8)+194140)) >= 0.0 {
			*(*uint16)(unsafe.Pointer(uintptr(a1 + 126))) = uint16(int16(int32(v7) + 42))
			if int32(int16(int32(v7)+42)) >= 256 {
				for {
					*(*uint16)(unsafe.Pointer(uintptr(a1 + 126))) -= 256
					if int32(*(*uint16)(unsafe.Pointer(uintptr(a1 + 126)))) < 256 {
						break
					}
				}
			}
		} else {
			v8 = int16(int32(v7) - 42)
			*(*uint16)(unsafe.Pointer(uintptr(a1 + 126))) = uint16(int16(int32(v7) - 42))
			if int32(int16(int32(v7)-42)) < 0 {
				for {
					v8 += 256
					if int32(v8) >= 0 {
						break
					}
				}
				*(*uint16)(unsafe.Pointer(uintptr(a1 + 126))) = uint16(v8)
			}
		}
	}
	result = int32(*(*int16)(unsafe.Pointer(uintptr(a1 + 126)))) * 8
	*(*float32)(unsafe.Pointer(uintptr(a1 + 88))) = *mem_getFloatPtr(0x587000, uint32(result)+194136) * *(*float32)(unsafe.Pointer(uintptr(a1 + 544)))
	v10 = float64(*mem_getFloatPtr(0x587000, uint32(result)+194140) * *(*float32)(unsafe.Pointer(uintptr(a1 + 544))))
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 112))) = 1061997773
	*(*float32)(unsafe.Pointer(uintptr(a1 + 92))) = float32(v10)
	return result
}
func nox_xxx_updateTeleportPentagram_53BEF0(a1 int32) int32 {
	var (
		v1     *uint8
		v2     uint8
		v3     uint8
		result int32
		v5     int32
		v6     *uint8
	)
	_ = v6
	var a1a float4
	v1 = *(**uint8)(unsafe.Pointer(uintptr(a1 + 748)))
	if int32(*v1) != 0 {
		if int32(*v1) <= 2 {
			v2 = *(*uint8)(unsafe.Add(unsafe.Pointer(v1), 9))
			if int32(v2) == int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v1), 8))) {
				*(*uint8)(unsafe.Add(unsafe.Pointer(v1), 20))++
				nox_xxx_unitNeedSync_4E44F0((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
				if int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v1), 20))) == 9 {
					v3 = *(*uint8)(unsafe.Add(unsafe.Pointer(v1), 8))
					*(*uint8)(unsafe.Add(unsafe.Pointer(v1), 20)) = 1
					*(*uint8)(unsafe.Add(unsafe.Pointer(v1), 8)) = uint8(int8(int32(v3) + 1))
				}
				*(*uint8)(unsafe.Add(unsafe.Pointer(v1), 9)) = 0
			} else {
				*(*uint8)(unsafe.Add(unsafe.Pointer(v1), 9)) = uint8(int8(int32(v2) + 1))
			}
		}
	} else {
		if int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v1), 20))) != 0 {
			nox_xxx_unitNeedSync_4E44F0((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
		}
		*(*uint8)(unsafe.Add(unsafe.Pointer(v1), 20)) = 0
	}
	result = int32(*v1)
	if int32(*v1) != 0 {
		v5 = result - 1
		if v5 != 0 {
			result = v5 - 1
			if result == 0 && int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v1), 8))) >= 4 {
				*(*uint8)(unsafe.Add(unsafe.Pointer(v1), 20)) = 0
				*v1 = 0
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), 4*1))) = 0
				return result
			}
		} else {
			result = int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v1), 8)))
			if int32(uint8(int8(result))) != 0 {
				if int32(uint8(int8(result))) == 4 {
					*v1 = 0
					*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), 4*1))) = 0
					return result
				}
			} else if int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v1), 20))) == 8 {
				result = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), 4*3))))
				if result != 0 {
					a1a.field_0 = *(*float32)(unsafe.Pointer(uintptr(a1 + 56))) - *(*float32)(unsafe.Pointer(uintptr(a1 + 176)))
					a1a.field_4 = *(*float32)(unsafe.Pointer(uintptr(a1 + 60))) - *(*float32)(unsafe.Pointer(uintptr(a1 + 176)))
					a1a.field_8 = *(*float32)(unsafe.Pointer(uintptr(a1 + 56))) + *(*float32)(unsafe.Pointer(uintptr(a1 + 176)))
					a1a.field_C = *(*float32)(unsafe.Pointer(uintptr(a1 + 60))) + *(*float32)(unsafe.Pointer(uintptr(a1 + 176)))
					nox_xxx_getUnitsInRect_517C10(&a1a, unsafe.Pointer(funAddr(nox_xxx_fnPentagramTeleport_53C060)), unsafe.Pointer(uintptr(result+56)))
					*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), 4*1))) = 0
					return result
				}
			}
		}
	} else if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), 4*1))) != 0 {
		result = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), 4*3))))
		if result != 0 {
			if *(*uint32)(unsafe.Pointer(uintptr(a1 + 16)))&0x1000000 != 0 {
				v6 = *(**uint8)(unsafe.Pointer(uintptr(result + 748)))
				*v1 = 1
				*(*uint8)(unsafe.Add(unsafe.Pointer(v1), 8)) = 0
				*(*uint8)(unsafe.Add(unsafe.Pointer(v1), 9)) = 0
				*(*uint32)(unsafe.Pointer(uintptr(a1 + 136))) = gameFrame()
				*v6 = 2
				*(*uint8)(unsafe.Add(unsafe.Pointer(v6), 8)) = 0
				*(*uint8)(unsafe.Add(unsafe.Pointer(v6), 9)) = 0
				result = int32(gameFrame())
				*(*uint32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), 4*3))) + 136))) = gameFrame()
			}
		}
	}
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), 4*1))) = 0
	return result
}
func nox_xxx_fnPentagramTeleport_53C060(a1 *float32, a2 int32) {
	if (uint32(*(*float32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(float32(0))*2))) & 0x420000) == 0 {
		nox_xxx_netSendPointFx_522FF0(-119, (*float2)(unsafe.Add(unsafe.Pointer((*float2)(unsafe.Pointer(a1))), unsafe.Sizeof(float2{})*7)))
		nox_xxx_aud_501960(147, (*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), 0, 0)
		nox_xxx_teleportToMB_4E7190((*uint8)(unsafe.Pointer(a1)), (*float32)(unsafe.Pointer(uintptr(a2))))
		nox_xxx_netSendPointFx_522FF0(-119, (*float2)(unsafe.Add(unsafe.Pointer((*float2)(unsafe.Pointer(a1))), unsafe.Sizeof(float2{})*7)))
		nox_xxx_aud_501960(147, (*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), 0, 0)
	}
}
func nox_xxx_updateInvisiblePentagram_53C0C0(a1 int32) int32 {
	var (
		result int32
		v2     int32
		v3     int32
		a1a    float4
	)
	result = a1
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	if *(*uint32)(unsafe.Pointer(uintptr(v2 + 4))) != 0 {
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 12))))
		if v3 != 0 {
			if *(*uint32)(unsafe.Pointer(uintptr(a1 + 16)))&0x1000000 != 0 {
				a1a.field_0 = *(*float32)(unsafe.Pointer(uintptr(a1 + 56))) - *(*float32)(unsafe.Pointer(uintptr(a1 + 176)))
				a1a.field_4 = *(*float32)(unsafe.Pointer(uintptr(a1 + 60))) - *(*float32)(unsafe.Pointer(uintptr(a1 + 176)))
				a1a.field_8 = *(*float32)(unsafe.Pointer(uintptr(a1 + 176))) + *(*float32)(unsafe.Pointer(uintptr(a1 + 56)))
				a1a.field_C = *(*float32)(unsafe.Pointer(uintptr(a1 + 60))) + *(*float32)(unsafe.Pointer(uintptr(a1 + 176)))
				nox_xxx_getUnitsInRect_517C10(&a1a, unsafe.Pointer(funAddr(sub_53C140)), unsafe.Pointer(uintptr(v3+56)))
			}
		}
	}
	*(*uint32)(unsafe.Pointer(uintptr(v2 + 4))) = 0
	return result
}
func sub_53C140(a1 *float32, a2 int32) {
	if (uint32(*(*float32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(float32(0))*2))) & 0x420000) == 0 {
		nox_xxx_teleportToMB_4E7190((*uint8)(unsafe.Pointer(a1)), (*float32)(unsafe.Pointer(uintptr(a2))))
	}
}
func nox_xxx_updateBlow_53C160(a3 int32) {
	var (
		v1 float64
		v2 float32
		v3 float32
		v4 float64
		v5 float32
		v6 float64
		v7 float32
		a2 int2
		a1 float4
	)
	if *(*uint32)(unsafe.Pointer(uintptr(a3 + 16)))&0x1000000 != 0 {
		nox_xxx_xferIndexedDirection_509E20(int32(*(*int16)(unsafe.Pointer(uintptr(a3 + 124)))), &a2)
		if a2.field_0 >= 0 {
			if a2.field_4 < 0 {
				v4 = float64(*(*float32)(unsafe.Pointer(uintptr(a3 + 60)))) - 400.0
				v5 = *(*float32)(unsafe.Pointer(uintptr(a3 + 60)))
				a1.field_0 = *(*float32)(unsafe.Pointer(uintptr(a3 + 56)))
				a1.field_C = v5
				a1.field_4 = float32(v4)
				a1.field_8 = float32(float64(*(*float32)(unsafe.Pointer(uintptr(a3 + 56)))) + 400.0)
				nox_xxx_getUnitsInRectAdv_517ED0(&a1, unsafe.Pointer(funAddr(sub_53C240)), unsafe.Pointer(uintptr(a3)))
				return
			}
			v6 = float64(*(*float32)(unsafe.Pointer(uintptr(a3 + 56)))) + 400.0
			v7 = *(*float32)(unsafe.Pointer(uintptr(a3 + 60)))
			a1.field_0 = *(*float32)(unsafe.Pointer(uintptr(a3 + 56)))
			a1.field_4 = v7
			a1.field_8 = float32(v6)
		} else {
			a1.field_0 = float32(float64(*(*float32)(unsafe.Pointer(uintptr(a3 + 56)))) - 400.0)
			if a2.field_4 < 0 {
				v1 = float64(*(*float32)(unsafe.Pointer(uintptr(a3 + 60)))) - 400.0
				v2 = *(*float32)(unsafe.Pointer(uintptr(a3 + 60)))
				a1.field_8 = *(*float32)(unsafe.Pointer(uintptr(a3 + 56)))
				a1.field_C = v2
				a1.field_4 = float32(v1)
				nox_xxx_getUnitsInRectAdv_517ED0(&a1, unsafe.Pointer(funAddr(sub_53C240)), unsafe.Pointer(uintptr(a3)))
				return
			}
			v3 = *(*float32)(unsafe.Pointer(uintptr(a3 + 56)))
			a1.field_4 = *(*float32)(unsafe.Pointer(uintptr(a3 + 60)))
			a1.field_8 = v3
		}
		a1.field_C = float32(float64(*(*float32)(unsafe.Pointer(uintptr(a3 + 60)))) + 400.0)
		nox_xxx_getUnitsInRectAdv_517ED0(&a1, unsafe.Pointer(funAddr(sub_53C240)), unsafe.Pointer(uintptr(a3)))
		return
	}
}
func sub_53C240(a1 *float32, arg4 int32) {
	var (
		v2  *float32
		v3  int32
		v4  float64
		v5  float64
		v6  float64
		v7  float64
		v8  float64
		v9  float64
		v10 float32
		v11 float32
		a2  int2
		v13 float32
		v14 float32
		v15 float32
		v16 float32
	)
	v2 = a1
	if (int32(uint8(*(*float32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(float32(0))*4))))&0x20) == 0 && (uint32(*(*float32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(float32(0))*2)))&0x400000) == 0 {
		v3 = arg4
		v15 = *(*float32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(float32(0))*14)) - *(*float32)(unsafe.Pointer(uintptr(arg4 + 56)))
		v4 = float64(*(*float32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(float32(0))*15)) - *(*float32)(unsafe.Pointer(uintptr(v3 + 60))))
		v13 = float32(v4)
		v5 = math.Sqrt(v4*float64(v13)+float64(v15*v15)) + 0.1
		v11 = float32(v5)
		if v5 < 400.0 {
			nox_xxx_xferIndexedDirection_509E20(int32(*(*int16)(unsafe.Pointer(uintptr(v3 + 124)))), &a2)
			v6 = float64(v15)
			if float64(v15) >= 0.0 {
				v10 = v15
			} else {
				v10 = float32(-v6)
			}
			if float64(v13) >= 0.0 {
				v16 = v13
			} else {
				v16 = -v13
			}
			switch a2.field_4 + a2.field_0 + a2.field_4*2 + 4 {
			case 0:
				if v6 < 0.0 {
					if float64(v13) < 0.0 {
						v7 = float64(v16 / v10)
						if v7 >= 0.57730001 && v7 <= 0.1732 {
							break
						}
					}
				}
				return
			case 1:
				if float64(v13) < 0.0 {
					if float64(v16/v10) <= 0.3732 {
						break
					}
				}
				return
			case 2:
				if v6 > 0.0 {
					if float64(v13) < 0.0 {
						v7 = float64(v16 / v10)
						if v7 >= 0.57730001 && v7 <= 0.1732 {
							break
						}
					}
				}
				return
			case 3:
				if v6 < 0.0 {
					if float64(v16/v10) <= 0.26789999 {
						break
					}
				}
				return
			case 5:
				if v6 > 0.0 {
					if float64(v16/v10) <= 0.26789999 {
						break
					}
				}
				return
			case 6:
				if v6 < 0.0 && float64(v13) > 0.0 {
					v7 = float64(v16 / v10)
					if v7 >= 0.57730001 && v7 <= 0.1732 {
						break
					}
				}
				return
			case 7:
				if float64(v13) > 0.0 {
					if float64(v16/v10) <= 0.3732 {
						break
					}
				}
				return
			case 8:
				if v6 > 0.0 && float64(v13) > 0.0 {
					v8 = float64(v16 / v10)
					if v8 >= 0.57730001 && v8 <= 0.1732 {
						break
					}
				}
				return
			default:
			}
			if nox_xxx_unitCanInteractWith_5370E0((*nox_object_t)(unsafe.Pointer(uintptr(v3))), (*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v2)))))), 0) != 0 {
				v14 = float32((400.0 - float64(v11)) * (400.0 - float64(v11)) * (400.0 - float64(v11)) * 5e-07)
				v9 = float64(v14) / nox_xxx_objectGetMass_4E4A70(int32(uintptr(unsafe.Pointer(v2))))
				*(*float32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(float32(0))*22)) = float32(v9*float64(*mem_getFloatPtr(0x587000, uint32(int32(*(*int16)(unsafe.Pointer(uintptr(v3 + 124))))*8)+194136)) + float64(*(*float32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(float32(0))*22))))
				*(*float32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(float32(0))*23)) = float32(v9*float64(*mem_getFloatPtr(0x587000, uint32(int32(*(*int16)(unsafe.Pointer(uintptr(v3 + 124))))*8)+194140)) + float64(*(*float32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(float32(0))*23))))
			}
			return
		}
	}
}
func nox_xxx_rechargeItem_53C520(a1 int32, a2 int32) int32 {
	var (
		v2 int32
		v3 int32
		v4 int32
		v5 int32
	)
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 736))))
	if v2 == 0 {
		return 0
	}
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 112))))
	if v3 >= 100 {
		return 0
	}
	v4 = a2 + v3
	if v4 < 100 {
		*(*uint32)(unsafe.Pointer(uintptr(v2 + 112))) = uint32(v4)
	} else {
		*(*uint32)(unsafe.Pointer(uintptr(v2 + 112))) = 100
	}
	v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 112))) * uint32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 109)))) / 100)
	if v5 == int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 108)))) {
		return 0
	}
	*(*uint8)(unsafe.Pointer(uintptr(v2 + 108))) = uint8(int8(v5))
	return 1
}
func nox_xxx_updateObelisk_53C580(a1 int32) int32 {
	var (
		v1     int32
		v2     *int32
		v3     int32
		v4     int32
		v5     int32
		v6     float64
		v7     float64
		v8     int32
		v9     *uint32
		v10    int32
		v11    int32
		v12    int16
		v13    int8
		v14    int16
		v15    int32
		v16    uint16
		v17    int32
		result int32
		v19    float32
		v20    float32
		v21    int32
		v22    int32
		v23    int32
		v24    *int32
	)
	v1 = a1
	v23 = 1
	v2 = *(**int32)(unsafe.Pointer(uintptr(a1 + 748)))
	v24 = *(**int32)(unsafe.Pointer(uintptr(a1 + 748)))
	v3 = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0())))
	v21 = v3
	if v3 == 0 {
		goto LABEL_50
	}
	for {
		v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 16))))
		if (v4 & 0x8000) != 0 {
			goto LABEL_47
		}
		v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 748))))
		if nox_xxx_servObjectHasTeam_419130(v1+48) != 0 {
			if nox_xxx_servCompareTeams_419150(v1+48, v3+48) == 0 {
				goto LABEL_47
			}
		}
		v6 = float64(*(*float32)(unsafe.Pointer(uintptr(v1 + 56))) - *(*float32)(unsafe.Pointer(uintptr(v3 + 56))))
		v7 = float64(*(*float32)(unsafe.Pointer(uintptr(v1 + 60))) - *(*float32)(unsafe.Pointer(uintptr(v3 + 60))))
		v22 = 0
		if v7*v7+v6*v6 >= 2500.0 || nox_xxx_mapCheck_537110((*nox_object_t)(unsafe.Pointer(uintptr(v1))), (*nox_object_t)(unsafe.Pointer(uintptr(v3)))) == 0 {
			goto LABEL_47
		}
		v23 = 0
		if *v24 >= 1 {
			v8 = nox_xxx_getRechargeRate_53C940(*(**uint32)(unsafe.Pointer(uintptr(v5 + 104))))
			if !nox_common_gameFlags_check_40A5C0(0x2000) || nox_common_gameFlags_check_40A5C0(4096) {
				if v8 == 0 {
					v3 = v21
					goto LABEL_25
				}
			} else {
				v8 = 1
			}
			v9 = *(**uint32)(unsafe.Pointer(uintptr(v5 + 104)))
			if v9 != nil {
				v10 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v9), 4*2)))
				if v10&0x1000 != 0 {
					if *(*uint32)(unsafe.Add(unsafe.Pointer(v9), 4*3))&0x47F0000 != 0 {
						v11 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v9), 4*184)))
						if *(*int32)(unsafe.Pointer(uintptr(v11 + 112))) < 100 {
							if nox_common_gameFlags_check_40A5C0(4096) {
								if (gameFrame() % (gameFPS() >> 1)) == 0 {
									nox_xxx_aud_501960(230, (*nox_object_t)(unsafe.Pointer(uintptr(v1))), 0, 0)
								}
							} else {
								v22 = 1
								*v24--
							}
							if nox_xxx_rechargeItem_53C520(int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 104)))), v8) != 0 {
								nox_xxx_netReportCharges_4D82B0(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v5 + 276))) + 2064)))), (*nox_object_t)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(uintptr(v5 + 104))))), int8(*(*uint8)(unsafe.Pointer(uintptr(v11 + 108)))), int8(*(*uint8)(unsafe.Pointer(uintptr(v11 + 109)))))
							}
						}
					}
				}
			}
			v3 = v21
			goto LABEL_25
		}
	LABEL_25:
		if *v24 < 1 || int32(*(*uint16)(unsafe.Pointer(uintptr(v5 + 4)))) >= int32(*(*uint16)(unsafe.Pointer(uintptr(v5 + 8)))) || *v24 <= 0 {
			if v22 == 0 {
				goto LABEL_47
			}
			goto LABEL_43
		}
		v12 = 1
		if nox_common_gameFlags_check_40A5C0(4096) {
			v13 = int8(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v5 + 276))) + 2251))))
			if int32(v13) != 0 {
				if int32(v13) == 1 {
					v14 = int16(nox_float2int(nox_xxx_wizardMaximumMana_587000_312820))
				} else {
					if int32(v13) != 2 {
						goto LABEL_36
					}
					v14 = int16(nox_float2int(nox_xxx_conjurerMaxMana_587000_312804))
				}
			} else {
				v14 = int16(nox_float2int(nox_xxx_warriorMaxMana_587000_312788))
			}
			v12 = v14
		}
	LABEL_36:
		v15 = int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 276))))
		*(*uint16)(unsafe.Pointer(uintptr(v5 + 4))) += uint16(v12)
		nox_xxx_protectMana_56F9E0(int32(*(*uint32)(unsafe.Pointer(uintptr(v15 + 4596)))), v12)
		v16 = *(*uint16)(unsafe.Pointer(uintptr(v5 + 8)))
		if int32(*(*uint16)(unsafe.Pointer(uintptr(v5 + 4)))) > int32(v16) {
			v17 = int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 276))))
			*(*uint16)(unsafe.Pointer(uintptr(v5 + 4))) = v16
			nox_xxx_protectPlayerHPMana_56F870(int32(*(*uint32)(unsafe.Pointer(uintptr(v17 + 4596)))), v16)
		}
		if nox_common_gameFlags_check_40A5C0(4096) {
			if (gameFrame() % (gameFPS() >> 1)) == 0 {
				nox_xxx_aud_501960(230, (*nox_object_t)(unsafe.Pointer(uintptr(v1))), 0, 0)
			}
			if v22 == 0 {
				goto LABEL_47
			}
			goto LABEL_43
		}
		*v24--
	LABEL_43:
		if (*v24 % 8) == 0 {
			v19 = float32(float64(*v24 * 80 / 50))
			nullsub_35(uint32(v1), *(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v19))), 4*0)))
			nox_xxx_unitNeedSync_4E44F0((*nox_object_t)(unsafe.Pointer(uintptr(v1))))
		}
		if (gameFrame() - *(*uint32)(unsafe.Pointer(uintptr(v1 + 136)))) > uint32(int32(gameFPS())>>1) {
			nox_xxx_aud_501960(230, (*nox_object_t)(unsafe.Pointer(uintptr(v1))), 0, 0)
			*(*uint32)(unsafe.Pointer(uintptr(v1 + 136))) = gameFrame()
		}
	LABEL_47:
		v21 = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(v3)))))))
		if v21 == 0 {
			break
		}
		v3 = v21
	}
	result = v23
	if v23 == 0 {
		return result
	}
	v2 = v24
LABEL_50:
	result = int32(gameFrame() / (gameFPS() >> 1))
	if (gameFrame() % (gameFPS() >> 1)) == 0 {
		result = *v2
		if *v2 < 50 {
			if (result % 8) == 0 {
				v20 = float32(float64(result * 80 / 50))
				nullsub_35(uint32(v1), *(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v20))), 4*0)))
				nox_xxx_unitNeedSync_4E44F0((*nox_object_t)(unsafe.Pointer(uintptr(v1))))
			}
			*v2++
		}
	}
	return result
}
func nox_xxx_getRechargeRate_53C940(a1 *uint32) int32 {
	var (
		v1     int32
		result int32
		v3     int32
		i      *int32
		v5     int32
		v6     float32
	)
	if a1 == nil {
		return 0
	}
	v1 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*2)))
	if v1&0x1000 != 0 && *(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*3))&0x4000000 != 0 {
		v6 = float32(nox_xxx_gamedataGetFloat_419D40(internCStr("OblivionStaffRechargeRate")))
		result = nox_float2int(v6)
	} else {
		v3 = 2
		for i = (*int32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*173)) + 8))); ; i = (*int32)(unsafe.Add(unsafe.Pointer(i), 4*1)) {
			v5 = *i
			if *i != 0 {
				if funAddr(*(*func() int32)(unsafe.Pointer(uintptr(v5 + 40)))) == funAddr(nullsub_36) {
					break
				}
			}
			if func() int32 {
				p := &v3
				*p++
				return *p
			}() >= 4 {
				return 0
			}
		}
		result = int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 48))))
	}
	return result
}
func nox_xxx_updateBlackPowderBarrel_53C9A0(a1 *float32) {
	var (
		v1  *float32
		v2  int32
		v3  int32
		v4  *float32
		v5  float32
		v6  int32
		v7  int32
		v8  int32
		v9  *uint32
		v10 *uint32
		v11 int32
		v12 float4
		v13 float32
	)
	v1 = a1
	v2 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), 4*34))))
	v3 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), 4*32))))
	if v2 == v3 {
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), 4*34))) = gameFrame() + uint32(nox_common_randomInt_415FA0(1, 5))
	} else if uint32(v2) == gameFrame() {
		if *memmap.PtrUint32(0x5D4594, 2488684) == 0 {
			*memmap.PtrUint32(0x5D4594, 2488684) = uint32(nox_xxx_getNameId_4E3AA0(internCStr("SmallFlame")))
			*memmap.PtrUint32(0x5D4594, 2488688) = uint32(nox_xxx_getNameId_4E3AA0(internCStr("MediumFlame")))
		}
		v4 = (*float32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(float32(0))*14))
		nox_xxx_mapDamageUnitsAround_4E25B0((*float32)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer((*float32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(float32(0))*14)))))))), 100.0, 30.0, 30, 7, (*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), nil)
		nox_xxx_mapPushUnitsAround_52E040(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer((*float32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(float32(0))*14))))))), 100.0, 30.0, 60.0, (*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), 0, 0)
		v5 = *(*float32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(float32(0))*15))
		v12.field_0 = *(*float32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(float32(0))*14))
		v12.field_4 = v5
		v6 = 4
		for {
			v7 = nox_common_randomInt_415FA0(0, 1)
			v13 = float32(nox_common_randomFloat_416030(0.0, 15.0) + 10.0)
			v8 = nox_common_randomInt_415FA0(0, math.MaxUint8)
			v12.field_8 = v13**mem_getFloatPtr(0x587000, uint32(v8*8)+194136) + *v4
			v12.field_C = v13**mem_getFloatPtr(0x587000, uint32(v8*8)+194140) + *(*float32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(float32(0))*15))
			if nox_xxx_mapTraceRay_535250(&v12, nil, nil, 1) != 0 {
				v9 = (*uint32)(unsafe.Pointer(nox_xxx_newObjectWithTypeInd_4E3450(int32(*memmap.PtrUint32(0x5D4594, uint32(v7*4)+2488684)))))
				v10 = v9
				if v9 != nil {
					nox_xxx_createAt_4DAA50((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v9)))))), nil, v12.field_8, v12.field_C)
					v11 = nox_common_randomInt_415FA0(5, 20)
					nox_xxx_unitSetDecayTime_511660((*nox_object_t)(unsafe.Pointer(v10)), int32(gameFPS()*uint32(v11)))
				}
			}
			v6--
			if v6 == 0 {
				break
			}
		}
	} else if (gameFrame() - uint32(v3)) >= uint32(int32(gameFPS())) {
		nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))))
	}
}
func nox_xxx_updateOneSecondDie_53CB60(a1 int32) {
	if (gameFrame() - *(*uint32)(unsafe.Pointer(uintptr(a1 + 128)))) >= uint32(int32(gameFPS())) {
		nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
	}
}
func nox_xxx_updateWaterBarrel_53CB90(a1 int32) {
	var (
		v1  uint32
		a1a float4
	)
	v1 = gameFrame() - *(*uint32)(unsafe.Pointer(uintptr(a1 + 128)))
	if v1 == 8 {
		*memmap.PtrUint32(0x5D4594, 2488664) = 0
		a1a.field_0 = float32(float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 56)))) - 40.0)
		a1a.field_4 = float32(float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 60)))) - 40.0)
		a1a.field_8 = float32(float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 56)))) + 40.0)
		a1a.field_C = float32(float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 60)))) + 40.0)
		nox_xxx_getUnitsInRect_517C10(&a1a, unsafe.Pointer(funAddr(nox_xxx_waterBarrel_53CC30)), unsafe.Pointer(uintptr(a1+56)))
		if *memmap.PtrUint32(0x5D4594, 2488664) != 0 {
			nox_xxx_aud_501960(283, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
		}
	} else if v1 >= 0x1E {
		nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
	}
}
func nox_xxx_waterBarrel_53CC30(a1 *float32, a2 int32) {
	var (
		v2 int32
		v3 float64
		v4 float64
	)
	v2 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), 4*2))))
	if v2&0x2000 != 0 {
		v3 = float64(*(*float32)(unsafe.Pointer(uintptr(a2))) - *(*float32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(float32(0))*14)))
		v4 = float64(*(*float32)(unsafe.Pointer(uintptr(a2 + 4))) - *(*float32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(float32(0))*15)))
		if math.Sqrt(v4*v4+v3*v3)-float64(*(*float32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(float32(0))*44))) <= 40.0 {
			nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))))
			*memmap.PtrUint32(0x5D4594, 2488664) = 1
		}
	}
}
func nox_xxx_updateSelfDestruct_53CC90(a1 int32) {
	if (gameFrame() - *(*uint32)(unsafe.Pointer(uintptr(a1 + 128)))) > 2 {
		nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
	}
}
func nox_xxx_updateBlackPowderBurn_53CCB0(a1 int32) {
	var (
		v1 int32
		v2 int32
	)
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 136))))
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 128))))
	if v1 == v2 {
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 136))) = gameFrame() + 3
	} else if uint32(v1) == gameFrame() {
		nox_xxx_mapDamageUnitsAround_4E25B0((*float32)(unsafe.Pointer(uintptr(a1+56))), 15.0, 15.0, 1, 1, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), nil)
	} else if gameFrame()-uint32(v2) >= (gameFPS() * 2) {
		nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
	}
}
func nox_xxx_updateDeathBall_53D080(a1 int32) {
	var (
		v1 int32
		v2 int32
		v3 int32
		v4 float32
		v5 float32
		v6 float32
		v7 float32
	)
	v1 = a1
	v2 = int32(gameFrame() - *(*uint32)(unsafe.Pointer(uintptr(a1 + 128))))
	if gameFrame()%(gameFPS()/3) != 0 {
		*memmap.PtrUint32(0x5D4594, 2488700) = 0
		v4 = float32(nox_xxx_gamedataGetFloat_419D40(internCStr("DeathBallCancelRange")))
		nox_xxx_getMissilesInCircle_518170((*float2)(unsafe.Pointer(uintptr(a1+56))), v4, unsafe.Pointer(funAddr(sub_53D170)), (*nox_object_t)(unsafe.Pointer(uintptr(a1))))
		if *memmap.PtrUint32(0x5D4594, 2488700) == 1 {
			nox_xxx_sMakeScorch_537AF0((*float32)(unsafe.Pointer(uintptr(a1+56))), 1)
			nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
		}
	}
	if v2 > 10 {
		v6 = float32(nox_xxx_gamedataGetFloat_419D40(internCStr("DeathBallOutRadius")))
		v7 = float32(nox_xxx_gamedataGetFloat_419D40(internCStr("DeathBallInRadius")))
		v5 = float32(nox_xxx_gamedataGetFloat_419D40(internCStr("DeathBallNearbyDamage")))
		v3 = nox_float2int(v5)
		nox_xxx_mapDamageUnitsAround_4E25B0((*float32)(unsafe.Pointer(uintptr(v1+56))), v6, v7, v3, 2, (*nox_object_t)(unsafe.Pointer(uintptr(v1))), nil)
	}
	if uint32(v2) > gameFPS()*4 {
		nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(v1))))
	}
}
func sub_53D170(a1 int32, a2 int32) {
	var (
		v2  int32
		v3  int32
		v4  int32
		v5  float32
		v6  float32
		a1a int4
	)
	v2 = int32(*memmap.PtrUint32(0x5D4594, 2488704))
	if *memmap.PtrUint32(0x5D4594, 2488704) == 0 {
		v2 = nox_xxx_getNameId_4E3AA0(internCStr("DeathBall"))
		*memmap.PtrUint32(0x5D4594, 2488704) = uint32(v2)
	}
	if a1 != a2 && int32(*(*uint16)(unsafe.Pointer(uintptr(a1 + 4)))) == v2 && nox_xxx_mapCheck_537110((*nox_object_t)(unsafe.Pointer(uintptr(a2))), (*nox_object_t)(unsafe.Pointer(uintptr(a1)))) != 0 {
		*memmap.PtrUint32(0x5D4594, 2488700) = 1
		if (int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 16)))) & 0x20) == 0 {
			a1a.field_0 = nox_float2int(*(*float32)(unsafe.Pointer(uintptr(a2 + 56))))
			v3 = nox_float2int(*(*float32)(unsafe.Pointer(uintptr(a2 + 60))))
			v6 = *(*float32)(unsafe.Pointer(uintptr(a1 + 56)))
			a1a.field_4 = v3
			v4 = nox_float2int(v6)
			v5 = *(*float32)(unsafe.Pointer(uintptr(a1 + 60)))
			a1a.field_8 = v4
			a1a.field_C = nox_float2int(v5)
			nox_xxx_netSendFxGreenBolt_523790(&a1a, 10)
			nox_xxx_sMakeScorch_537AF0((*float32)(unsafe.Pointer(uintptr(a1+56))), 1)
			nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
		}
	}
}
func nox_xxx_updateDeathBallFragment_53D220(a1 int32) {
	var v1 int32
	v1 = int32(gameFrame() - *(*uint32)(unsafe.Pointer(uintptr(a1 + 128))))
	if v1 > 10 {
		nox_xxx_mapDamageUnitsAround_4E25B0((*float32)(unsafe.Pointer(uintptr(a1+56))), 25.0, 0.0, 20, 2, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), nil)
	}
	if uint32(v1) > gameFPS()*2 {
		nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
	}
}
func nox_xxx_updateMoonglow_53D270(a1 int32) {
	var (
		v1 int32
		v2 int32
		v3 float2
	)
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 508))))
	if v1 != 0 {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 8))))&4 != 0 {
			v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 748))))
			if int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 16))))&0x20 != 0 || gameFrame()-*(*uint32)(unsafe.Pointer(uintptr(a1 + 128))) > (gameFPS()*300) {
				nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
				nox_xxx_spellBuffOff_4FF5B0((*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 508)))))), 1)
			} else {
				v3.field_0 = float32(float64(*(*int32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 2284)))))
				v3.field_4 = float32(float64(*(*int32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 2288)))))
				if sub_517590(v3.field_0, v3.field_4) != 0 {
					nox_xxx_unitMove_4E7010((*nox_object_t)(unsafe.Pointer(uintptr(a1))), &v3)
				}
			}
		}
	} else {
		nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
	}
}
func nox_xxx_updateTelekinesis_53D330(a1 int32) {
	var (
		v1 int32
		v2 int32
		v3 int32
		v4 int32
		v5 float4
	)
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 508))))
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 748))))
	if *(*uint32)(unsafe.Pointer(uintptr(v1 + 16)))&0x8020 != 0 || gameFrame()-*(*uint32)(unsafe.Pointer(uintptr(a1 + 128))) > (gameFPS()*20) {
		v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 508))))
		v3 = nox_xxx_spellGetAud44_424800(math.MaxInt8, 2)
		nox_xxx_aud_501960(v3, (*nox_object_t)(unsafe.Pointer(uintptr(v4))), 0, 0)
		nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
		nox_xxx_spellBuffOff_4FF5B0((*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 508)))))), 24)
	} else {
		v5.field_8 = float32(float64(*(*int32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 2284)))))
		v5.field_C = float32(float64(*(*int32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 2288)))))
		v5.field_0 = *(*float32)(unsafe.Pointer(uintptr(v1 + 56)))
		v5.field_4 = *(*float32)(unsafe.Pointer(uintptr(v1 + 60)))
		if nox_xxx_mapTraceRay_535250(&v5, nil, nil, 5) != 0 {
			nox_xxx_unitMove_4E7010((*nox_object_t)(unsafe.Pointer(uintptr(a1))), (*float2)(unsafe.Pointer(&v5.field_8)))
		}
	}
}
func nox_xxx_updateFist_53D400(a1 int32) {
	var (
		v1 float64
		v2 float64
		v3 float64
		a2 float2
	)
	if float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 104)))) <= 0.0 && int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16)))) >= 0 {
		nox_xxx_aud_501960(48, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
		nox_xxx_sMakeScorch_537AF0((*float32)(unsafe.Pointer(uintptr(a1+56))), 2)
		v1 = float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 176))) + *(*float32)(unsafe.Pointer(uintptr(a1 + 56))))
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))) |= 0x80000000
		a2.field_4 = *(*float32)(unsafe.Pointer(uintptr(a1 + 60)))
		a2.field_0 = float32(v1)
		nox_xxx_netSendPointFx_522FF0(-118, &a2)
		v2 = float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 56))) - *(*float32)(unsafe.Pointer(uintptr(a1 + 176))))
		a2.field_4 = *(*float32)(unsafe.Pointer(uintptr(a1 + 60)))
		a2.field_0 = float32(v2)
		nox_xxx_netSendPointFx_522FF0(-118, &a2)
		v3 = float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 176))) + *(*float32)(unsafe.Pointer(uintptr(a1 + 60))))
		a2.field_0 = *(*float32)(unsafe.Pointer(uintptr(a1 + 56)))
		a2.field_4 = float32(v3)
		nox_xxx_netSendPointFx_522FF0(-118, &a2)
		nox_xxx_earthquakeSend_4D9110((*float32)(unsafe.Pointer(uintptr(a1+56))), 30)
	}
	if float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 104)))) >= 200.0 && *(*int32)(unsafe.Pointer(uintptr(a1 + 16))) < 0 {
		nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
	}
	if gameFrame()-*(*uint32)(unsafe.Pointer(uintptr(a1 + 128))) > (gameFPS() * 3) {
		nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
	}
}
func nox_xxx_updateFlameCleanse_53D510(a1 int32) {
	var (
		v1 float32
		v2 float32
		v3 float32
		v4 float4
	)
	if gameFrame() >= uint32(*(*int32)(unsafe.Pointer(uintptr(a1 + 136)))) || (func() bool {
		v1 = *(*float32)(unsafe.Pointer(uintptr(a1 + 156)))
		v2 = *(*float32)(unsafe.Pointer(uintptr(a1 + 56)))
		v4.field_4 = *(*float32)(unsafe.Pointer(uintptr(a1 + 160)))
		v4.field_0 = v1
		v3 = *(*float32)(unsafe.Pointer(uintptr(a1 + 60)))
		v4.field_8 = v2
		v4.field_C = v3
		return nox_xxx_mapTraceRay_535250(&v4, nil, nil, 65) == 0
	}()) || (gameFrame()-*(*uint32)(unsafe.Pointer(uintptr(a1 + 128)))) > 3 && *(*float32)(unsafe.Pointer(uintptr(a1 + 72))) == *(*float32)(unsafe.Pointer(uintptr(a1 + 56))) && *(*float32)(unsafe.Pointer(uintptr(a1 + 76))) == *(*float32)(unsafe.Pointer(uintptr(a1 + 60))) {
		nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
	}
}
func nox_xxx_updateMeteorShower_53D5A0(a2 *float32) {
	var (
		v1 *uint32
		v2 int32
		v3 float64
		v4 float64
		v5 float32
		v6 float64
		v7 float64
		v8 *uint32
		v9 int32
		a3 float32
		a4 float32
		a1 float4
	)
	if gameFrame()-*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a2))), 4*32))) < (gameFPS() * 5) {
		if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a2))), 4*34))) <= gameFrame() {
			v1 = (*uint32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a2))), 4*187))))))
			v2 = nox_common_randomInt_415FA0(0, math.MaxUint8)
			v3 = nox_common_randomFloat_416030(4.0, 12.0)
			v4 = v3 * v3
			v5 = *(*float32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(float32(0))*15))
			a1.field_0 = *(*float32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(float32(0))*14))
			v6 = v4 * float64(*mem_getFloatPtr(0x587000, uint32(v2*8)+194136))
			a1.field_4 = v5
			a3 = float32(v6 + float64(*(*float32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(float32(0))*14))))
			v7 = v4 * float64(*mem_getFloatPtr(0x587000, uint32(v2*8)+194140))
			a1.field_8 = a3
			a4 = float32(v7 + float64(*(*float32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(float32(0))*15))))
			a1.field_C = a4
			if int32(uint8(int8(nox_xxx_traceRay_5374B0(&a1)))) != 0 {
				v8 = (*uint32)(unsafe.Pointer(nox_xxx_newObjectByTypeID_4E3810(internCStr("Meteor"))))
				v9 = int32(uintptr(unsafe.Pointer(v8)))
				if v8 != nil {
					*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v8), 4*187))))) = *v1
					nox_xxx_createAt_4DAA50((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v8)))))), (*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a2)))))), a3, a4)
					*(*uint32)(unsafe.Pointer(uintptr(v9 + 20))) |= 0x20
					nox_xxx_unitRaise_4E46F0((*nox_object_t)(unsafe.Pointer(uintptr(v9))), 255.0)
					*(*uint32)(unsafe.Pointer(uintptr(v9 + 108))) = 3238002688
				}
			}
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a2))), 4*34))) = gameFrame() + uint32(nox_common_randomInt_415FA0(4, 8))
		}
	} else {
		nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a2)))))))
	}
}
func nox_xxx_meteorExplode_53D6E0(a6 int32) {
	var (
		v1 *int32
		v2 *float32
		v3 *uint32
		v4 int32
		v5 float32
		v6 float32
		v7 float32
		v8 float32
		a1 int4
	)
	if float64(*(*float32)(unsafe.Pointer(uintptr(a6 + 104)))) <= 0.0 {
		v1 = *(**int32)(unsafe.Pointer(uintptr(a6 + 748)))
		nox_xxx_aud_501960(87, (*nox_object_t)(unsafe.Pointer(uintptr(a6))), 0, 0)
		v2 = (*float32)(unsafe.Pointer(uintptr(a6 + 56)))
		nox_xxx_sMakeScorch_537AF0((*float32)(unsafe.Pointer(uintptr(a6+56))), 2)
		nox_xxx_earthquakeSend_4D9110((*float32)(unsafe.Pointer(uintptr(a6+56))), 10)
		v3 = (*uint32)(unsafe.Pointer(nox_xxx_newObjectByTypeID_4E3810(internCStr("MeteorExplode"))))
		if v3 != nil {
			nox_xxx_createAt_4DAA50((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v3)))))), nil, *v2, *(*float32)(unsafe.Pointer(uintptr(a6 + 60))))
		}
		v4 = int32(uintptr(unsafe.Pointer(nox_xxx_findParentChainPlayer_4EC580((*nox_object_t)(unsafe.Pointer(uintptr(a6)))))))
		nox_xxx_mapDamageUnitsAround_4E25B0((*float32)(unsafe.Pointer(uintptr(a6+56))), 80.0, 30.0, *v1, 7, (*nox_object_t)(unsafe.Pointer(uintptr(v4))), nil)
		v5 = float32(float64(*v2) - 80.0)
		v6 = float32(float64(*(*float32)(unsafe.Pointer(uintptr(a6 + 60)))) - 80.0)
		v7 = float32(float64(*v2) + 80.0)
		v8 = float32(float64(*(*float32)(unsafe.Pointer(uintptr(a6 + 60)))) + 80.0)
		a1.field_0 = nox_float2int(v5) / 23
		a1.field_4 = nox_float2int(v6) / 23
		a1.field_8 = nox_float2int(v7) / 23
		a1.field_C = nox_float2int(v8) / 23
		nox_xxx_mapDamageToWalls_534FC0(&a1, unsafe.Pointer(uintptr(a6+56)), 80.0, *v1, 7, unsafe.Pointer(uintptr(a6)))
		nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(a6))))
	}
}
func nox_xxx_updateToxicCloud_53D850(a1 int32) {
	var v1 *uint32
	v1 = *(**uint32)(unsafe.Pointer(uintptr(a1 + 748)))
	if *(*uint32)(unsafe.Pointer(uintptr(a1 + 136))) < gameFrame() {
		nox_xxx_unitsGetInCircle_517F90((*float2)(unsafe.Pointer(uintptr(a1+56))), 75.0, unsafe.Pointer(funAddr(sub_53D8C0)), unsafe.Pointer(uintptr(a1)))
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 136))) = gameFrame() + uint32(nox_common_randomInt_415FA0(5, 10))
	}
	if *v1 != 0 {
		*v1--
	}
	if *v1 == 0 {
		nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
	}
}
func sub_53D8C0(a1 int32, a2 int32) {
	var (
		v2 float32
		v3 float32
		v4 float32
		v5 int32
		v6 int32
		v7 int32
		v8 float4
	)
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&6 != 0 {
		v2 = *(*float32)(unsafe.Pointer(uintptr(a1 + 56)))
		v3 = *(*float32)(unsafe.Pointer(uintptr(a2 + 56)))
		v8.field_4 = *(*float32)(unsafe.Pointer(uintptr(a2 + 60)))
		v8.field_0 = v3
		v4 = *(*float32)(unsafe.Pointer(uintptr(a1 + 60)))
		v8.field_8 = v2
		v8.field_C = v4
		if nox_xxx_mapTraceRay_535250(&v8, nil, nil, 5) != 0 {
			v7 = nox_common_randomInt_415FA0(3, 10)
			v5 = int32(uintptr(unsafe.Pointer(nox_xxx_findParentChainPlayer_4EC580((*nox_object_t)(unsafe.Pointer(uintptr(a2)))))))
			(*(*func(int32, int32, int32, int32, int32))(unsafe.Pointer(uintptr(a1 + 716))))(a1, v5, a2, v7, 5)
			v6 = int32(uintptr(unsafe.Pointer(nox_xxx_findParentChainPlayer_4EC580((*nox_object_t)(unsafe.Pointer(uintptr(a2)))))))
			if nox_xxx_unitIsEnemyTo_5330C0((*nox_object_t)(unsafe.Pointer(uintptr(v6))), (*nox_object_t)(unsafe.Pointer(uintptr(a1)))) != 0 {
				nox_xxx_activatePoison_4EE7E0(a1, 1, 1)
			}
		}
	}
}
func nox_xxx_updateSmallToxicCloud_53D960(a1 int32) {
	var v1 *uint32
	v1 = *(**uint32)(unsafe.Pointer(uintptr(a1 + 748)))
	if *(*uint32)(unsafe.Pointer(uintptr(a1 + 136))) < gameFrame() {
		nox_xxx_unitsGetInCircle_517F90((*float2)(unsafe.Pointer(uintptr(a1+56))), 35.0, unsafe.Pointer(funAddr(nox_xxx_toxicCloudPoison_53D9D0)), unsafe.Pointer(uintptr(a1)))
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 136))) = gameFrame() + uint32(nox_common_randomInt_415FA0(5, 10))
	}
	if *v1 != 0 {
		*v1--
	}
	if *v1 == 0 {
		nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
	}
}
func nox_xxx_toxicCloudPoison_53D9D0(a1 int32, a2 int32) {
	var (
		v2 float32
		v3 float32
		v4 float32
		v5 int32
		v6 int32
		v7 float4
	)
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&6 != 0 {
		v2 = *(*float32)(unsafe.Pointer(uintptr(a1 + 56)))
		v3 = *(*float32)(unsafe.Pointer(uintptr(a2 + 56)))
		v7.field_4 = *(*float32)(unsafe.Pointer(uintptr(a2 + 60)))
		v7.field_0 = v3
		v4 = *(*float32)(unsafe.Pointer(uintptr(a1 + 60)))
		v7.field_8 = v2
		v7.field_C = v4
		if nox_xxx_mapTraceRay_535250(&v7, nil, nil, 5) != 0 {
			v5 = int32(uintptr(unsafe.Pointer(nox_xxx_findParentChainPlayer_4EC580((*nox_object_t)(unsafe.Pointer(uintptr(a2)))))))
			(*(*func(int32, int32, int32, uint32, int32))(unsafe.Pointer(uintptr(a1 + 716))))(a1, v5, a2, 0, 5)
			v6 = int32(uintptr(unsafe.Pointer(nox_xxx_findParentChainPlayer_4EC580((*nox_object_t)(unsafe.Pointer(uintptr(a2)))))))
			if nox_xxx_unitIsEnemyTo_5330C0((*nox_object_t)(unsafe.Pointer(uintptr(v6))), (*nox_object_t)(unsafe.Pointer(uintptr(a1)))) != 0 {
				nox_xxx_activatePoison_4EE7E0(a1, 1, 1)
			}
		}
	}
}
func nox_xxx_updateArachnaphobia_53DA60(a1 *int32) {
	var (
		v1 int32
		v2 int32
		v3 *uint32
	)
	v1 = int32(gameFrame())
	if uint32(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*34))) < gameFrame() {
		v2 = int32(*memmap.PtrUint32(0x5D4594, 2488708))
		if *memmap.PtrUint32(0x5D4594, 2488708) == 0 {
			v2 = nox_xxx_getNameId_4E3AA0(internCStr("SmallSpider"))
			*memmap.PtrUint32(0x5D4594, 2488708) = uint32(v2)
		}
		v3 = (*uint32)(unsafe.Pointer(nox_xxx_newObjectWithTypeInd_4E3450(v2)))
		if v3 != nil {
			nox_xxx_createAt_4DAA50((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v3)))))), (*nox_object_t)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*math.MaxInt8))))), *((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(a1))), unsafe.Sizeof(float32(0))*14))), *((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(a1))), unsafe.Sizeof(float32(0))*15))))
		}
		*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*34)) = int32(gameFrame() + uint32(nox_common_randomInt_415FA0(1, 5)))
		v1 = int32(gameFrame())
	}
	if uint32(v1-*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*32))) > (gameFPS() * 3) {
		nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))))
	}
}
func nox_xxx_updateExpire_53DB00(a1 int32) {
	if *(*uint32)(unsafe.Pointer(uintptr(a1 + 128))) > gameFrame() || *(*uint32)(unsafe.Pointer(uintptr(a1 + 136))) < gameFrame() {
		nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
	}
}
func nox_xxx_updateBreak_53DB30(a1 *uint32) *int32 {
	var result *int32
	result = (*int32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*5)))))
	if int32(uint8(uintptr(unsafe.Pointer(result))))&2 != 0 {
		result = (*int32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*4)))))
		if int32(*(*int8)(unsafe.Add(unsafe.Pointer((*int8)(unsafe.Pointer(&result))), 1))) < 0 {
			nox_xxx_unitUnsetXStatus_4E4780((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), 2)
			nox_xxx_unitSetXStatus_4E4800((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), int32(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(uintptr(4)))))))
			result = (*int32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*4)))))
			*((*uint8)(unsafe.Pointer(&result))) = uint8(int8(int32(uint8(uintptr(unsafe.Pointer(result)))) | 0x40))
			*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*34)) = gameFrame() + gameFPS()*2
			*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*4)) = uint32(uintptr(unsafe.Pointer(result)))
		}
	} else if int32(uint8(uintptr(unsafe.Pointer(result))))&4 != 0 {
		result = (*int32)(unsafe.Pointer(uintptr(gameFrame())))
		if gameFrame() > *(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*34)) {
			nox_xxx_unitUnsetXStatus_4E4780((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), 4)
			result = (*int32)(unsafe.Pointer(uintptr(nox_xxx_unitSetXStatus_4E4800((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), int32(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(uintptr(8))))))))))
		}
	} else if int32(uint8(uintptr(unsafe.Pointer(result))))&8 != 0 {
		nox_xxx_unitRemoveFromUpdatable_4DA920((*nox_object_t)(unsafe.Pointer(a1)))
	}
	return result
}
func nox_xxx_updateOpen_53DBB0(a1 *uint32) *int32 {
	var result *int32
	result = (*int32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*5)))))
	if int32(uint8(uintptr(unsafe.Pointer(result))))&2 != 0 {
		result = (*int32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*4)))))
		if int32(*(*int8)(unsafe.Add(unsafe.Pointer((*int8)(unsafe.Pointer(&result))), 1))) < 0 {
			nox_xxx_unitUnsetXStatus_4E4780((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), 2)
			nox_xxx_unitSetXStatus_4E4800((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), int32(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(uintptr(4)))))))
			result = (*int32)(unsafe.Pointer(uintptr(gameFrame())))
			*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*34)) = gameFrame() + gameFPS()*2
		}
	} else if int32(uint8(uintptr(unsafe.Pointer(result))))&4 != 0 {
		result = (*int32)(unsafe.Pointer(uintptr(gameFrame())))
		if gameFrame() > *(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*34)) {
			nox_xxx_unitUnsetXStatus_4E4780((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), 4)
			result = (*int32)(unsafe.Pointer(uintptr(nox_xxx_unitSetXStatus_4E4800((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), int32(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(uintptr(8))))))))))
		}
	} else if int32(uint8(uintptr(unsafe.Pointer(result))))&8 != 0 {
		nox_xxx_unitRemoveFromUpdatable_4DA920((*nox_object_t)(unsafe.Pointer(a1)))
	}
	return result
}
func nox_xxx_updateBreakAndRemove_53DC30(a1 *uint32) {
	var (
		v1 int32
		v2 int32
	)
	switch *(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*5)) {
	case 2:
		v1 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*4)))
		if (v1 & 0x8000) != 0 {
			nox_xxx_unitUnsetXStatus_4E4780((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), 2)
			nox_xxx_unitSetXStatus_4E4800((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), int32(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(uintptr(4)))))))
			*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*34)) = gameFrame() + gameFPS()*2
			v2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*4)))
			*((*uint8)(unsafe.Pointer(&v2))) = uint8(int8(v2 | 0x40))
			*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*4)) = uint32(v2)
		}
	case 4:
		if gameFrame() > *(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*34)) {
			nox_xxx_unitUnsetXStatus_4E4780((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), 4)
			nox_xxx_unitSetXStatus_4E4800((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), int32(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(uintptr(8)))))))
		}
	case 8:
		nox_xxx_unitRemoveFromUpdatable_4DA920((*nox_object_t)(unsafe.Pointer(a1)))
		nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))))
	}
}
func nox_xxx_updateChakramInMotion_53DCC0(a1 int32) {
	var (
		v1  int32
		v2  int32
		v3  int32
		v4  int32
		v5  int32
		v6  int32
		v7  int32
		v8  float64
		v9  float64
		v10 float32
		v11 float32
	)
	v1 = a1
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	v3 = nox_xxx_inventoryGetFirst_4E7980(a1)
	if v3 == 0 || int32(*(*uint8)(unsafe.Pointer(uintptr(v3 + 16))))&0x20 != 0 {
		nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
		return
	}
	v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 12))))
	if v4 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(v4 + 16))))&0x20 != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(v2 + 12))) = 0
	}
	v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 508))))
	if v5 == 0 || int32(*(*uint8)(unsafe.Pointer(uintptr(v5 + 16))))&0x20 != 0 {
		*(*uint8)(unsafe.Pointer(uintptr(v2 + 24))) = 1
		*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = 0
	} else {
		*(*uint32)(unsafe.Pointer(uintptr(v2 + 16))) = *(*uint32)(unsafe.Pointer(uintptr(v5 + 56)))
		*(*uint32)(unsafe.Pointer(uintptr(v2 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(v5 + 60)))
		if nox_xxx_mapCheck_537110((*nox_object_t)(unsafe.Pointer(uintptr(a1))), (*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 508))))))) == 0 {
			*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = 0
		} else if int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 24)))) == 0 {
			*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 508)))
		} else {
			goto LABEL_19
		}
		if int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 24)))) == 0 {
			v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))))
			if v6 != 0 && ((func() int32 {
				v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(v6 + 16))))
				return v7 & 0x20
			}()) != 0 || (v7&0x8000) != 0) {
				*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = 0
				*(*uint8)(unsafe.Pointer(uintptr(v2 + 24))) = 1
			} else {
				v8 = float64(*(*float32)(unsafe.Pointer(uintptr(v2 + 16))) - *(*float32)(unsafe.Pointer(uintptr(a1 + 56))))
				v9 = float64(*(*float32)(unsafe.Pointer(uintptr(v2 + 20))) - *(*float32)(unsafe.Pointer(uintptr(a1 + 60))))
				v11 = float32(v9)
				v10 = float32(math.Sqrt(v9*float64(v11)+v8*v8) + *(*float64)(unsafe.Pointer(&qword_581450_10176)))
				*(*float32)(unsafe.Pointer(uintptr(v1 + 80))) = float32(v8 * float64(*(*float32)(unsafe.Pointer(uintptr(v1 + 544)))) / float64(v10))
				*(*float32)(unsafe.Pointer(uintptr(v1 + 84))) = v11 * *(*float32)(unsafe.Pointer(uintptr(v1 + 544))) / v10
			}
		}
	}
LABEL_19:
	if gameFrame()-*(*uint32)(unsafe.Pointer(uintptr(v1 + 128))) > (gameFPS() * 5) {
		*(*uint8)(unsafe.Pointer(uintptr(v2 + 24))) = 1
		*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = 0
	}
}
func nox_xxx_updateFlag_53DDF0(a1 int32) int32 {
	var (
		v1     int32
		v2     int32
		result int32
		v4     int32
		v5     int32
		v6     int8
		v7     uint8
	)
	v1 = a1
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	result = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))))
	if result != 0 {
		v4 = sub_4ECBD0(a1)
		v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))))
		a1 = v4
		v7 = *(*uint8)(unsafe.Pointer(uintptr(v1 + 52)))
		result = int32(gameFPS() * 3)
		if gameFrame()-uint32(v5) > (gameFPS() * 30) {
			nox_xxx_aud_501960(305, (*nox_object_t)(unsafe.Pointer(uintptr(v1))), 0, 0)
			v6 = int8(a1)
			*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = 0
			sub_4E82C0(v7, 0, v6, 0)
			nox_xxx_unitMove_4E7010((*nox_object_t)(unsafe.Pointer(uintptr(v1))), (*float2)(unsafe.Pointer(uintptr(v2))))
			result = nox_xxx_netInformTextMsg2_4DA180(8, (*uint8)(unsafe.Pointer(&a1)))
		}
	}
	return result
}
func nox_xxx_updateTrapDoor_53DE80(a1 *uint32) *int32 {
	var (
		v1     int32
		result *int32
	)
	v1 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*175)))
	if *(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*4))&0x1000000 != 0 {
		result = (*int32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*5)))))
		if int32(uint8(uintptr(unsafe.Pointer(result))))&2 != 0 {
			nox_xxx_unitUnsetXStatus_4E4780((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), 2)
			result = (*int32)(unsafe.Pointer(uintptr(nox_xxx_unitSetXStatus_4E4800((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), int32(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(uintptr(8))))))))))
		} else if int32(uint8(uintptr(unsafe.Pointer(result))))&4 != 0 {
			result = *(**int32)(unsafe.Pointer(uintptr(v1 + 16)))
			if gameFrame() >= uint32(uintptr(unsafe.Pointer(result))) {
				nox_xxx_unitUnsetXStatus_4E4780((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), 4)
				result = (*int32)(unsafe.Pointer(uintptr(nox_xxx_unitSetXStatus_4E4800((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), int32(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(uintptr(8))))))))))
			}
		} else {
			*(*uint32)(unsafe.Pointer(uintptr(v1 + 24))) = 0
		}
	} else {
		result = *(**int32)(unsafe.Pointer(uintptr(v1 + 16)))
		if result != nil {
			if gameFrame() >= uint32(uintptr(unsafe.Pointer(result))) {
				nox_xxx_unitSetOnOff_4E4670(int32(uintptr(unsafe.Pointer(a1))), 1)
				nox_xxx_unitUnsetXStatus_4E4780((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), 2)
				nox_xxx_unitSetXStatus_4E4800((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), int32(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(uintptr(4)))))))
				*(*uint32)(unsafe.Pointer(uintptr(v1 + 16))) += gameFPS() * 5
				nox_xxx_aud_501960(874, (*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), 0, 0)
			}
		}
	}
	return result
}
func nox_xxx_updateGameBall_53DF40(a3 int32) {
	var (
		v1  int32
		v2  uint64
		v3  int32
		v4  int32
		v5  int32
		v6  float64
		v7  float64
		v8  float64
		v9  float64
		v10 int32
		v11 float4
	)
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a3 + 748))))
	*(*uint32)(unsafe.Pointer(uintptr(a3 + 112))) = 1008981770
	if *(*uint32)(unsafe.Pointer(uintptr(v1))) != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1))) + 16))))&0x20 != 0 {
		sub_4EB9B0(a3, 0)
		nox_xxx_netChangeTeamMb_419570(unsafe.Pointer(uintptr(a3+48)), int32(*(*uint32)(unsafe.Pointer(uintptr(a3 + 36)))))
		sub_4E8290(1, 0)
	}
	v2 = uint64(nox_platform_get_ticks()) - *(*uint64)(unsafe.Pointer(uintptr(v1 + 8)))
	v11.field_4 = *((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(&v2))), unsafe.Sizeof(float32(0))*1)))
	if v2 <= 20000 {
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a3 + 508))))
		if v3 != 0 {
			if uint32(v3) != *(*uint32)(unsafe.Pointer(uintptr(v1))) || (gameFrame()-*(*uint32)(unsafe.Pointer(uintptr(v1 + 16)))) <= uint32(*(*int32)(unsafe.Pointer(uintptr(v1 + 20)))) {
				*(*uint32)(unsafe.Pointer(uintptr(a3 + 16))) |= 0x40
				*(*uint64)(unsafe.Pointer(uintptr(v1 + 8))) = uint64(nox_platform_get_ticks())
				v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a3 + 508))))
				if *(*uint32)(unsafe.Pointer(uintptr(v5 + 16)))&0x8020 != 0 {
					nox_xxx_unitClearOwner_4EC300((*nox_object_t)(unsafe.Pointer(uintptr(a3))))
					sub_4EB9B0(a3, 0)
					nox_xxx_netChangeTeamMb_419570(unsafe.Pointer(uintptr(a3+48)), int32(*(*uint32)(unsafe.Pointer(uintptr(a3 + 36)))))
					sub_4E8290(1, 0)
				} else {
					v6 = float64(*(*float32)(unsafe.Pointer(uintptr(v5 + 176))) + *(*float32)(unsafe.Pointer(uintptr(a3 + 176))))
					v11.field_0 = *(*float32)(unsafe.Pointer(uintptr(v5 + 56)))
					v11.field_4 = *(*float32)(unsafe.Pointer(uintptr(v5 + 60)))
					v7 = v6 + 10.0
					v11.field_8 = float32(v7*float64(*mem_getFloatPtr(0x587000, uint32(int32(*(*int16)(unsafe.Pointer(uintptr(v5 + 124))))*8)+194136)) + float64(*(*float32)(unsafe.Pointer(uintptr(v5 + 56)))))
					v11.field_C = float32(v7*float64(*mem_getFloatPtr(0x587000, uint32(int32(*(*int16)(unsafe.Pointer(uintptr(v5 + 124))))*8)+194140)) + float64(*(*float32)(unsafe.Pointer(uintptr(v5 + 60)))))
					if nox_xxx_mapTraceRay_535250(&v11, nil, nil, 5) != 0 {
						nox_xxx_unitMove_4E7010((*nox_object_t)(unsafe.Pointer(uintptr(a3))), (*float2)(unsafe.Pointer(&v11.field_8)))
					}
				}
			} else {
				*(*uint32)(unsafe.Pointer(uintptr(a3 + 16))) &= 0xFFFFFFBF
				*(*uint32)(unsafe.Pointer(uintptr(a3 + 520))) = 0
				v4 = int32(*(*int16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a3 + 508))) + 124)))) + nox_common_randomInt_415FA0(-32, 32)
				if v4 < 0 {
					v4 += int32(uint32(math.MaxUint8-v4) >> 8 << 8)
				}
				if v4 >= 256 {
					v4 += int32((uint32(v4) >> 8) * 4294967040)
				}
				v11.field_0 = float32(float64(*(*float32)(unsafe.Pointer(uintptr(a3 + 56)))) - float64(*mem_getFloatPtr(0x587000, uint32(v4*8)+194136))*20.0)
				v11.field_4 = float32(float64(*(*float32)(unsafe.Pointer(uintptr(a3 + 60)))) - float64(*mem_getFloatPtr(0x587000, uint32(v4*8)+194140))*20.0)
				nox_xxx_objectApplyForce_52DF80((*float32)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(&v11)))))), (*nox_object_t)(unsafe.Pointer(uintptr(a3))), 30.0)
				nox_xxx_unitClearOwner_4EC300((*nox_object_t)(unsafe.Pointer(uintptr(a3))))
				sub_4E8290(1, 0)
				nox_xxx_aud_501960(926, (*nox_object_t)(unsafe.Pointer(uintptr(a3))), 0, 0)
			}
		} else {
			v8 = float64(*(*float32)(unsafe.Pointer(uintptr(a3 + 84))))
			v9 = float64(*(*float32)(unsafe.Pointer(uintptr(a3 + 80))))
			v10 = int32(*(*uint32)(unsafe.Pointer(uintptr(a3 + 16))))
			*((*uint8)(unsafe.Pointer(&v10))) = uint8(int8(v10 & 0xBF))
			*(*uint32)(unsafe.Pointer(uintptr(a3 + 16))) = uint32(v10)
			if float64(*(*float32)(unsafe.Pointer(uintptr(v1 + 24)))) > math.Sqrt(v9*v9+v8*v8) {
				*(*uint32)(unsafe.Pointer(uintptr(v1))) = 0
			}
		}
	} else {
		sub_417F50(a3)
	}
}
func nox_xxx_updateUndeadKiller_53E190(a1 int32) {
	var v1 int32
	v1 = int32(**(**uint32)(unsafe.Pointer(uintptr(a1 + 700))))
	if v1 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 88))))&1 != 0 {
		nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
	} else if (gameFrame() - *(*uint32)(unsafe.Pointer(uintptr(a1 + 136)))) > 0x46 {
		nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
	}
}
func nox_xxx_updateCrown_53E1D0(a1 int32) {
	var (
		v1 *uint32
		v2 int32
		v3 int32
		v4 float64
		v5 float64
		v6 float4
	)
	v1 = *(**uint32)(unsafe.Pointer(uintptr(a1 + 748)))
	v2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*1)))
	if v2 == 0 || *(*uint32)(unsafe.Pointer(uintptr(v2 + 16)))&0x8020 != 0 {
		if *v1 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(*v1 + 16))))&0x20 != 0 {
			*v1 = 0
		}
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 508))))
		if v3 != 0 {
			if *(*uint32)(unsafe.Pointer(uintptr(v3 + 16)))&0x8020 != 0 {
				nox_xxx_unitClearOwner_4EC300((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
			} else {
				v4 = float64(*(*float32)(unsafe.Pointer(uintptr(v3 + 176))) + *(*float32)(unsafe.Pointer(uintptr(a1 + 176))))
				v6.field_0 = *(*float32)(unsafe.Pointer(uintptr(v3 + 56)))
				v6.field_4 = *(*float32)(unsafe.Pointer(uintptr(v3 + 60)))
				v5 = v4 + 10.0
				v6.field_8 = float32(v5*float64(*mem_getFloatPtr(0x587000, uint32(int32(*(*int16)(unsafe.Pointer(uintptr(v3 + 124))))*8)+194136)) + float64(*(*float32)(unsafe.Pointer(uintptr(v3 + 56)))))
				v6.field_C = float32(v5*float64(*mem_getFloatPtr(0x587000, uint32(int32(*(*int16)(unsafe.Pointer(uintptr(v3 + 124))))*8)+194140)) + float64(*(*float32)(unsafe.Pointer(uintptr(v3 + 60)))))
				if nox_xxx_mapTraceRay_535250(&v6, nil, nil, 5) != 0 {
					nox_xxx_unitMove_4E7010((*nox_object_t)(unsafe.Pointer(uintptr(a1))), (*float2)(unsafe.Pointer(&v6.field_8)))
				}
			}
		}
	} else {
		sub_4F3400(v2, a1, 1)
	}
}
func sub_53E2D0(a1 int32) int32 {
	var result int32
	if *(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))&0x2000000 != 0 {
		result = bool2int32((nox_xxx_unitArmorInventoryEquipFlags_415C70((*nox_object_t)(unsafe.Pointer(uintptr(a1)))) & 0xC0D) == 0)
	} else {
		result = 1
	}
	return result
}
func nox_xxx_recalculateArmorVal_53E300(a1 *uint32) int32 {
	var (
		i      *uint32
		v2     int32
		result int32
		v4     float32
	)
	v4 = 0.0
	for i = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*126))))); i != nil; i = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(i), 4*124))))) {
		if *(*uint32)(unsafe.Add(unsafe.Pointer(i), 4*2))&0x2000000 != 0 {
			v2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(i), 4*4)))
			if v2&0x100 != 0 {
				v4 = float32(nox_xxx_itemApplyDefendEffect_415C00(int32(uintptr(unsafe.Pointer(i)))) + float64(v4))
			}
		}
	}
	if float64(v4) > *(*float64)(unsafe.Pointer(&qword_581450_9512)) {
		v4 = 1.0
	}
	result = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*2)))
	if result&4 != 0 {
		result = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*187)))
		*(*float32)(unsafe.Pointer(uintptr(result + 228))) = v4
	} else if result&2 != 0 {
		result = int32(*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v4))), 4*0)))
		*(*float32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*187)) + 2072))) = v4
	}
	return result
}
func sub_53E3A0(a1 int32, object *nox_object_t) int32 {
	var (
		result int32
		v3     int32
	)
	if (object.obj_class & 0x2000000) == 0 {
		return 0
	}
	v3 = int32(object.obj_flags)
	if (v3 & 0x100) == 0 {
		return 0
	}
	result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 504))))
	if result == 0 {
		return 0
	}
	for unsafe.Pointer(uintptr(result)) != unsafe.Pointer(object) {
		result = int32(*(*uint32)(unsafe.Pointer(uintptr(result + 496))))
		if result == 0 {
			return result
		}
	}
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v3))), 1)) &= 0xFE
	object.obj_flags = uint32(v3)
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 12))))&0x10 != 0 {
		nox_xxx_npcSetItemEquipFlags_4E4B20(a1, object, 0)
	}
	object.obj_flags &= 0xEFFFFFFF
	nox_xxx_recalculateArmorVal_53E300((*uint32)(unsafe.Pointer(uintptr(a1))))
	nox_xxx_itemApplyDisengageEffect_4F3030(object, a1)
	return 1
}
func sub_53E430(a1 *uint32, object *nox_object_t, a3 int32, a4 int32) int32 {
	var (
		v4 int32
		v5 int32
		v7 int32
		v8 int32
		v9 *uint32
	)
	_ = v9
	var v10 int8
	if (object.obj_class & 0x2000000) == 0 {
		return 0
	}
	v4 = int32(object.obj_flags)
	if (v4 & 0x100) == 0 {
		return 0
	}
	v5 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*2)))
	if v5&2 != 0 {
		return sub_53E3A0(int32(uintptr(unsafe.Pointer(a1))), object)
	}
	if (v5 & 4) == 0 {
		return 0
	}
	v7 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*126)))
	v8 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*187)))
	if v7 == 0 {
		return 0
	}
	for unsafe.Pointer(uintptr(v7)) != unsafe.Pointer(object) {
		v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(v7 + 496))))
		if v7 == 0 {
			return 0
		}
	}
	object.obj_flags = uint32(v4) & 0xEFFFFEFF
	v9 = *(**uint32)(unsafe.Pointer(uintptr(v8 + 276)))
	*v9 &= uint32(^nox_xxx_unitArmorInventoryEquipFlags_415C70(object))
	if a3 != 0 {
		nox_xxx_netReportDequip_4D8590(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v8 + 276))) + 2064)))), object)
	}
	if a4 != 0 {
		nox_xxx_netReportDequip_4D84C0(math.MaxUint8, object)
	}
	nox_xxx_recalculateArmorVal_53E300(a1)
	nox_xxx_itemApplyDisengageEffect_4F3030(object, int32(uintptr(unsafe.Pointer(a1))))
	if int32(*(*uint8)(unsafe.Pointer(&object.field_12)))&2 != 0 {
		v10 = int8(*(*uint8)(unsafe.Pointer(uintptr(v8 + 88))))
		if int32(v10) == 15 || int32(v10) == 16 || int32(v10) == 17 {
			nox_xxx_playerSetState_4FA020((*nox_object_t)(unsafe.Pointer(a1)), 13)
		}
	}
	return 1
}
func nox_xxx_NPCEquipArmor_53E520(a1 int32, a2 *uint32) int32 {
	var (
		v3 *uint32
		v4 *uint32
	)
	if (*(*uint32)(unsafe.Add(unsafe.Pointer(a2), 4*2)) & 0x2000000) == 0 {
		return 0
	}
	if *(*uint32)(unsafe.Add(unsafe.Pointer(a2), 4*4))&0x100 != 0 {
		return 0
	}
	v3 = *(**uint32)(unsafe.Pointer(uintptr(a1 + 504)))
	v4 = *(**uint32)(unsafe.Pointer(uintptr(a1 + 504)))
	if v4 == nil {
		return 0
	}
	for v4 != a2 {
		v4 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v4), 4*124)))))
		if v4 == nil {
			return 0
		}
	}
	if v3 != nil {
		for (*(*uint32)(unsafe.Add(unsafe.Pointer(v3), 4*2))&0x2000000) == 0 || (*(*uint32)(unsafe.Add(unsafe.Pointer(v3), 4*4))&0x100) == 0 || *(*uint32)(unsafe.Add(unsafe.Pointer(v3), 4*3)) != *(*uint32)(unsafe.Add(unsafe.Pointer(a2), 4*3)) {
			v3 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), 4*124)))))
			if v3 == nil {
				goto LABEL_18
			}
		}
		sub_53E3A0(a1, (*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v3)))))))
	}
LABEL_18:
	*(*uint32)(unsafe.Add(unsafe.Pointer(a2), 4*4)) |= 0x100
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 12))))&0x10 != 0 {
		nox_xxx_npcSetItemEquipFlags_4E4B20(a1, (*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a2)))))), 1)
	}
	if sub_53E2D0(int32(uintptr(unsafe.Pointer(a2)))) == 0 {
		*(*uint32)(unsafe.Add(unsafe.Pointer(a2), 4*4)) |= 0x10000000
	}
	nox_xxx_recalculateArmorVal_53E300((*uint32)(unsafe.Pointer(uintptr(a1))))
	nox_xxx_itemApplyEngageEffect_4F2FF0((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a2)))))), a1)
	if *(*uint32)(unsafe.Add(unsafe.Pointer(a2), 4*3))&2 != 0 {
		sub_53E600((*uint32)(unsafe.Pointer(uintptr(a1))))
	}
	return 1
}
func sub_53E600(a1 *uint32) {
	var (
		i  *uint32
		v2 int32
	)
	if a1 != nil {
		for i = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*126))))); i != nil; i = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(i), 4*124))))) {
			v2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(i), 4*4)))
			if v2&0x100 != 0 && *(*uint32)(unsafe.Add(unsafe.Pointer(i), 4*2))&0x1001000 != 0 {
				if uint32(nox_xxx_weaponInventoryEquipFlags_415820((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(i))))))))&0x7FFE40C != 0 {
					nox_xxx_playerDequipWeapon_53A140(a1, (*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(i)))))), 1, 1)
				}
			}
		}
	}
}
func nox_xxx_playerEquipArmor_53E650(a1 *uint32, item *nox_object_t, a3 int32, a4 int32) int32 {
	var (
		a2  int32 = int32(uintptr(unsafe.Pointer(item)))
		v4  int32
		v5  int32
		v7  int32
		v8  *uint32
		v9  int32
		v10 int32
		v11 *uint32
	)
	_ = v11
	if (item.obj_class & 0x2000000) == 0 {
		return 0
	}
	v4 = int32(item.obj_flags)
	if v4&0x100 != 0 {
		return 0
	}
	v5 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*2)))
	if v5&2 != 0 {
		return nox_xxx_NPCEquipArmor_53E520(int32(uintptr(unsafe.Pointer(a1))), (*uint32)(unsafe.Pointer(uintptr(a2))))
	}
	if (v5 & 4) == 0 {
		return 0
	}
	v7 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*187)))
	v8 = nox_xxx_armorHaveSameSubclass_53E7B0(int32(uintptr(unsafe.Pointer(a1))), a2)
	if nox_xxx_playerClassCanUseItem_57B3D0(item, int8(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v7 + 276))) + 2251))))) == 0 {
		nox_xxx_netPriMsgToPlayer_4DA2C0((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), internCStr("armor.c:ArmorEquipClassFail"), 0)
		if a4 != 0 {
			nox_xxx_aud_501960(925, (*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), 2, int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*9))))
		}
		return 0
	}
	*((*uint8)(unsafe.Pointer(&v9))) = uint8(int8(bool2int32(nox_xxx_playerCheckStrength_4F3180((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), item))))
	if v9 == 0 {
		nox_xxx_netPriMsgToPlayer_4DA2C0((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), internCStr("armor.c:ArmorEquipStrengthFail"), 0)
		if a4 != 0 {
			nox_xxx_aud_501960(925, (*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), 2, int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*9))))
		}
		return 0
	}
	if v8 != nil {
		sub_53E430(a1, (*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v8)))))), 1, 1)
	}
	v10 = int32(item.obj_flags)
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v10))), 1)) |= 1
	item.obj_flags = uint32(v10)
	v11 = *(**uint32)(unsafe.Pointer(uintptr(v7 + 276)))
	*v11 |= uint32(nox_xxx_unitArmorInventoryEquipFlags_415C70(item))
	nox_xxx_netReportEquip_4D8540(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v7 + 276))) + 2064)))), (*uint32)(unsafe.Pointer(uintptr(a2))), a3)
	if sub_53E2D0(a2) == 0 {
		item.obj_flags |= 0x10000000
	}
	nox_xxx_recalculateArmorVal_53E300(a1)
	nox_xxx_itemApplyEngageEffect_4F2FF0(item, int32(uintptr(unsafe.Pointer(a1))))
	if item.obj_subclass&0x2 != 0 {
		sub_53E600(a1)
	}
	return 1
}
func nox_xxx_armorHaveSameSubclass_53E7B0(a1 int32, a2 int32) *uint32 {
	var result *uint32
	result = *(**uint32)(unsafe.Pointer(uintptr(a1 + 504)))
	if result == nil {
		return nil
	}
	for (*(*uint32)(unsafe.Add(unsafe.Pointer(result), 4*4))&0x100) == 0 || (*(*uint32)(unsafe.Add(unsafe.Pointer(result), 4*2))&0x2000000) == 0 || *(*uint32)(unsafe.Add(unsafe.Pointer(result), 4*3)) != *(*uint32)(unsafe.Pointer(uintptr(a2 + 12))) {
		result = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(result), 4*124)))))
		if result == nil {
			return nil
		}
	}
	return result
}
func nox_xxx_pickupArmor_53E7F0(a1 int32, a2 int32, a3 int32, a4 int32) int32 {
	var (
		v5  *uint32
		v6  int32
		v7  *uint32
		v8  int32
		v9  *uint32
		i   int32
		v11 int32
		v12 *uint32
		v13 int32
		m   int32
		v15 int16
		v16 int32
		j   int32
		v18 *uint32
		k   int32
		l   int32
	)
	if *memmap.PtrUint32(0x5D4594, 2488712) == 0 {
		*memmap.PtrUint32(0x5D4594, 2488712) = uint32(nox_xxx_getNameId_4E3AA0(internCStr("StreetSneakers")))
		*memmap.PtrUint32(0x5D4594, 2488716) = uint32(nox_xxx_getNameId_4E3AA0(internCStr("WizardRobe")))
		dword_5d4594_2488720 = uint32(nox_xxx_getNameId_4E3AA0(internCStr("WoodenShield")))
		dword_5d4594_2488724 = uint32(nox_xxx_getNameId_4E3AA0(internCStr("SteelShield")))
	}
	if !nox_common_gameFlags_check_40A5C0(2048) && !nox_common_gameFlags_check_40A5C0(4096) && sub_409F40(2) != 0 && sub_4E7EC0(a1, (*nox_object_t)(unsafe.Pointer(uintptr(a2)))) != 0 {
		nox_xxx_netPriMsgToPlayer_4DA2C0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), internCStr("armor.c:CannotPickupDuplicateArmor"), 0)
		nox_xxx_aud_501960(925, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 2, int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 36)))))
		return 0
	}
	if !nox_common_gameFlags_check_40A5C0(2048) && !nox_common_gameFlags_check_40A5C0(4096) && int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4 != 0 && nox_xxx_playerClassCanUseItem_57B3D0((*nox_object_t)(unsafe.Pointer(uintptr(a2))), int8(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + 276))) + 2251))))) == 0 {
		nox_xxx_netPriMsgToPlayer_4DA2C0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), internCStr("armor.c:ArmorEquipClassFail"), 0)
		nox_xxx_aud_501960(925, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 2, int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 36)))))
		return 0
	}
	if nox_xxx_pickupDefault_4F31E0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), (*nox_object_t)(unsafe.Pointer(uintptr(a2))), a3) != 1 {
		return 0
	}
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4 != 0 {
		v5 = nox_xxx_armorHaveSameSubclass_53E7B0(a1, a2)
		v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
		v7 = v5
		v8 = sub_419E60((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
		if v8 == 0 {
			if v7 != nil {
				*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v8))), unsafe.Sizeof(uint16(0))*0)) = *((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v7))), unsafe.Sizeof(uint16(0))*2)))
				if uint32(v8) != dword_5d4594_2488720 && uint32(v8) != dword_5d4594_2488724 && uint32(v8) != *memmap.PtrUint32(0x5D4594, 2488712) && uint32(v8) != *memmap.PtrUint32(0x5D4594, 2488716) {
					goto LABEL_40
				}
				if uint32(v8) == *memmap.PtrUint32(0x5D4594, 2488716) {
					v9 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v7), 4*173)))))
					for i = 0; i < 4; i++ {
						if *v9 != 0 {
							break
						}
						v9 = (*uint32)(unsafe.Add(unsafe.Pointer(v9), 4*1))
					}
					if i == 4 {
						nox_xxx_playerEquipArmor_53E650((*uint32)(unsafe.Pointer(uintptr(a1))), (*nox_object_t)(unsafe.Pointer(uintptr(a2))), a4, 0)
					}
					goto LABEL_40
				}
			}
			if (int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 12)))) & 2) == 0 {
				nox_xxx_playerEquipArmor_53E650((*uint32)(unsafe.Pointer(uintptr(a1))), (*nox_object_t)(unsafe.Pointer(uintptr(a2))), a4, 0)
			} else {
				if v7 == nil {
					v11 = int32(*(*uint32)(unsafe.Pointer(uintptr(v6 + 104))))
					if v11 == 0 || (*(*uint32)(unsafe.Pointer(uintptr(v11 + 12)))&0x7FFE40C) == 0 {
						nox_xxx_playerEquipArmor_53E650((*uint32)(unsafe.Pointer(uintptr(a1))), (*nox_object_t)(unsafe.Pointer(uintptr(a2))), a4, 0)
					}
					goto LABEL_40
				}
				v12 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v7), 4*173)))))
				v13 = int32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v7))), unsafe.Sizeof(uint16(0))*2))))
				if uint32(uint16(int16(v13))) != dword_5d4594_2488720 {
					if uint32(v13) == dword_5d4594_2488724 {
						v16 = int32(*(*uint16)(unsafe.Pointer(uintptr(a2 + 4))))
						if uint32(uint16(int16(v16))) == dword_5d4594_2488720 {
							for j = 0; j < 4; j++ {
								if *v12 != 0 {
									break
								}
								v12 = (*uint32)(unsafe.Add(unsafe.Pointer(v12), 4*1))
							}
							if j == 4 {
								v18 = *(**uint32)(unsafe.Pointer(uintptr(a2 + 692)))
								for k = 0; k < 4; k++ {
									if *v18 != 0 {
										break
									}
									v18 = (*uint32)(unsafe.Add(unsafe.Pointer(v18), 4*1))
								}
								if k != 4 {
									nox_xxx_playerEquipArmor_53E650((*uint32)(unsafe.Pointer(uintptr(a1))), (*nox_object_t)(unsafe.Pointer(uintptr(a2))), a4, 0)
								}
							}
						} else if uint32(v16) == dword_5d4594_2488724 {
							for l = 0; l < 4; l++ {
								if *v12 != 0 {
									break
								}
								v12 = (*uint32)(unsafe.Add(unsafe.Pointer(v12), 4*1))
							}
							if l == 4 {
								nox_xxx_playerEquipArmor_53E650((*uint32)(unsafe.Pointer(uintptr(a1))), (*nox_object_t)(unsafe.Pointer(uintptr(a2))), a4, 0)
							}
						}
					}
					goto LABEL_40
				}
				for m = 0; m < 4; m++ {
					if *v12 != 0 {
						break
					}
					v12 = (*uint32)(unsafe.Add(unsafe.Pointer(v12), 4*1))
				}
				if m == 4 {
					nox_xxx_playerEquipArmor_53E650((*uint32)(unsafe.Pointer(uintptr(a1))), (*nox_object_t)(unsafe.Pointer(uintptr(a2))), a4, 0)
				}
			}
		}
	}
LABEL_40:
	v15 = int16(*(*uint16)(unsafe.Pointer(uintptr(a2 + 24))))
	if int32(v15)&0x10 != 0 {
		nox_xxx_aud_501960(804, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
	} else if int32(v15)&8 != 0 {
		nox_xxx_aud_501960(810, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
	} else if int32(v15)&4 != 0 {
		nox_xxx_aud_501960(807, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
	} else if int32(v15)&2 != 0 {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 12))))&0x20 != 0 {
			nox_xxx_aud_501960(816, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
		} else {
			nox_xxx_aud_501960(813, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
		}
	}
	nox_xxx_decay_5116F0((*nox_object_t)(unsafe.Pointer(uintptr(a2))))
	return 1
}
func sub_53EAE0(a1 int32) {
	var v2 int16
	if a1 != 0 {
		v2 = int16(*(*uint16)(unsafe.Pointer(uintptr(a1 + 24))))
		if int32(v2)&0x10 != 0 {
			nox_xxx_aud_501960(805, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
		} else if int32(v2)&8 != 0 {
			nox_xxx_aud_501960(811, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
		} else if int32(v2)&4 != 0 {
			nox_xxx_aud_501960(808, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
		} else if int32(v2)&2 != 0 {
			if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 12))))&0x20 != 0 {
				nox_xxx_aud_501960(817, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
			} else {
				nox_xxx_aud_501960(814, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
			}
		}
	}
}
func nox_xxx_dropArmor_53EB70(a1 int32, a2 *uint32, a3 *int32) int32 {
	if nox_xxx_dropDefault_4ED290((*nox_object_t)(unsafe.Pointer(uintptr(a1))), (*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a2)))))), (*float2)(unsafe.Pointer(a3))) != 1 {
		return 0
	}
	sub_53EAE0(int32(uintptr(unsafe.Pointer(a2))))
	if !nox_common_gameFlags_check_40A5C0(2048) && !nox_common_gameFlags_check_40A5C0(4096) {
		if sub_409F40(2) != 0 {
			nox_xxx_unitSetDecayTime_511660((*nox_object_t)(unsafe.Pointer(a2)), int32(gameFPS()*25))
		}
	}
	return 1
}
func nox_xxx_ItemIsDroppable_53EBF0(a1 int32) int32 {
	var (
		i  *uint8
		v2 int32
	)
	if a1 == 0 {
		return 0
	}
	if dword_5d4594_2488728 == 0 {
		sub_53EC40()
	}
	if *memmap.PtrUint32(0x587000, 279432) == 0 {
		return 0
	}
	for i = (*uint8)(memmap.PtrOff(0x587000, 279432)); *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), 4*1))) != uint32(*(*uint16)(unsafe.Pointer(uintptr(a1 + 4)))); i = (*uint8)(unsafe.Add(unsafe.Pointer(i), 12)) {
		v2 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), 4*3))))
		if v2 == 0 {
			return 0
		}
	}
	return 1
}
func sub_53EC40() *byte {
	var (
		result *byte
		v1     *uint8
	)
	result = *(**byte)(memmap.PtrOff(0x587000, 279432))
	if *memmap.PtrUint32(0x587000, 279432) != 0 {
		v1 = (*uint8)(memmap.PtrOff(0x587000, 279432))
		for {
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), 4*1))) = uint32(nox_xxx_getNameId_4E3AA0(result))
			result = (*byte)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), 4*3))))))
			v1 = (*uint8)(unsafe.Add(unsafe.Pointer(v1), 12))
			if result == nil {
				break
			}
		}
		dword_5d4594_2488728 = 1
	} else {
		dword_5d4594_2488728 = 1
	}
	return result
}
func sub_53EC80(a1 int32, a2 int32) int32 {
	var (
		v2 int32
		i  *uint8
		v4 int32
	)
	if a1 == 0 {
		return 0
	}
	if dword_5d4594_2488728 == 0 {
		sub_53EC40()
	}
	v2 = 0
	if *memmap.PtrUint32(0x587000, 279432) == 0 {
		return 0
	}
	for i = (*uint8)(memmap.PtrOff(0x587000, 279432)); *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), 4*1))) != uint32(*(*uint16)(unsafe.Pointer(uintptr(a1 + 4)))); i = (*uint8)(unsafe.Add(unsafe.Pointer(i), 12)) {
		v4 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), 4*3))))
		v2++
		if v4 == 0 {
			return 0
		}
	}
	return bool2int32((uint32(a2) & *memmap.PtrUint32(0x587000, uint32(v2*12)+279440)) != 0)
}
func nox_xxx_useMushroom_53ECE0(a1 int32, a2 int32) int32 {
	var v2 int32
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 540)))) != 0 {
		nox_xxx_removePoison_4EE9D0(a1)
		nox_xxx_netPriMsgToPlayer_4DA2C0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), internCStr("Use.c:MushroomClean"), 0)
		v2 = nox_xxx_spellGetAud44_424800(14, 1)
		nox_xxx_aud_501960(v2, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
	} else {
		nox_xxx_netPriMsgToPlayer_4DA2C0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), internCStr("Use.c:MushroomConfuse"), 0)
	}
	nox_xxx_buffApplyTo_4FF380((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 3, 300, 5)
	nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(a2))))
	return 1
}
func nox_xxx_useEnchant_53ED60(a1 int32, a2 int32) int32 {
	nox_xxx_buffApplyTo_4FF380((*nox_object_t)(unsafe.Pointer(uintptr(a1))), int32(**(**uint32)(unsafe.Pointer(uintptr(a2 + 736)))), int16(uint16(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 736))) + 4))))), 5)
	nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(a2))))
	return 1
}
func nox_xxx_useCast_53ED90(a1 int32, a2 *uint32) int32 {
	var (
		v2 *int32
		v3 int32
		v4 int32
		v6 [3]int32
	)
	v6[0] = a1
	v2 = (*int32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), 4*184)))))
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4 != 0 {
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
		*(*float32)(unsafe.Pointer(&v6[1])) = float32(float64(*(*int32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 276))) + 2284)))))
		*(*float32)(unsafe.Pointer(&v6[2])) = float32(float64(*(*int32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 276))) + 2288)))))
	} else {
		v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 60))))
		v6[1] = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 56))))
		v6[2] = v4
	}
	nox_xxx_spellAccept_4FD400(*v2, (*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a2)))))), (*nox_object_t)(unsafe.Pointer(a2)), (*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a2)))))), unsafe.Pointer(&v6[0]), 4)
	nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a2)))))))
	return 1
}
func nox_xxx_useConsume_53EE10(a1 int32, a2 int32) int32 {
	var (
		v2 *uint16
		v3 int32
	)
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 8))))&0x10 != 0 {
		if (int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 12)))) & 8) == 0 {
			v2 = *(**uint16)(unsafe.Pointer(uintptr(a1 + 556)))
			if v2 != nil {
				if int32(*v2) < int32(*(*uint16)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint16(0))*2))) {
					nox_xxx_unitAdjustHP_4EE460((*nox_object_t)(unsafe.Pointer(uintptr(a1))), int32(**(**uint32)(unsafe.Pointer(uintptr(a2 + 736)))))
					if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4 != 0 {
						v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 12))))
						if int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + 276))) + 2252)))) != 0 {
							if v3&1 != 0 {
								nox_xxx_aud_501960(324, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
							} else if v3&2 != 0 {
								nox_xxx_aud_501960(325, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
							} else if v3&4 != 0 {
								nox_xxx_aud_501960(326, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
							}
						} else if v3&1 != 0 {
							nox_xxx_aud_501960(314, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
						} else if v3&2 != 0 {
							nox_xxx_aud_501960(315, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
						} else if v3&4 != 0 {
							nox_xxx_aud_501960(316, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
						}
					} else {
						nox_xxx_aud_501960(334, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
					}
					nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(a2))))
				}
			}
		}
	}
	return 1
}
func nox_xxx_useCiderConfuse_53EF00(a1 int32, a2 int32) int32 {
	var v2 int32
	if a1 == 0 || a2 == 0 || *(*uint32)(unsafe.Pointer(uintptr(a1 + 556))) == 0 {
		return 1
	}
	nox_xxx_buffApplyTo_4FF380((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 3, int16(int32(uint16(gameFPS()))*5), 4)
	nox_xxx_netPriMsgToPlayer_4DA2C0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), internCStr("Use.c:CiderConfuse"), 0)
	v2 = nox_xxx_useConsume_53EE10(a1, a2)
	if v2 != 0 {
		nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(a2))))
	}
	return v2
}
func nox_xxx_usePotion_53EF70(a1 int32, a2 int32) int32 {
	var (
		v2     int32
		v3     int32
		result int32
		v5     int32
		v6     *uint16
		v7     int8
		v8     float64
		v9     int32
		v10    int8
		v11    float64
		v12    int32
		v13    int32
		v14    int32
		v15    int32
		v16    int32
		v17    int32
		v18    int32
		v19    int32
		v20    int32
		v21    int32
		v22    int32
		v23    float32
		v24    float32
		v25    int32
		v26    [3]int32
	)
	v2 = a2
	v3 = int32(**(**uint32)(unsafe.Pointer(uintptr(a2 + 736))))
	result = 0
	v25 = int32(**(**uint32)(unsafe.Pointer(uintptr(a2 + 736))))
	if (*(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))&4) != 0 && !(func() bool {
		v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))))
		return (v5 & 0x8000) == 0
	}()) {
		return result
	}
	if (int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 12))))&0x10) == 0 || (func() *uint16 {
		v6 = *(**uint16)(unsafe.Pointer(uintptr(a1 + 556)))
		return v6
	}()) == nil {
		goto LABEL_16
	}
	if int32(*v6) >= int32(*(*uint16)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint16(0))*2))) {
		v2 = a2
		goto LABEL_16
	}
	if *(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))&4 != 0 {
		v7 = int8(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + 276))) + 2251))))
		switch v7 {
		case 0:
			v8 = float64(v25) * float64(nox_xxx_warriorMaxHealth_587000_312784)
			v23 = float32(v8)
			v3 = nox_float2int(v23)
			v25 = v3
		case 1:
			v8 = float64(v25) * float64(nox_xxx_wizardMaxHealth_587000_312816)
			v23 = float32(v8)
			v3 = nox_float2int(v23)
			v25 = v3
		case 2:
			v8 = float64(v25) * float64(nox_xxx_conjurerMaxHealth_587000_312800)
			v23 = float32(v8)
			v3 = nox_float2int(v23)
			v25 = v3
		}
	}
	nox_xxx_unitAdjustHP_4EE460((*nox_object_t)(unsafe.Pointer(uintptr(a1))), v3)
	nox_xxx_aud_501960(754, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
	result = 1
	v2 = a2
LABEL_16:
	if (int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 12))))&0x20) == 0 || (int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4) == 0 || (func() bool {
		v9 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
		return int32(*(*uint16)(unsafe.Pointer(uintptr(v9 + 4)))) >= int32(*(*uint16)(unsafe.Pointer(uintptr(v9 + 8))))
	}()) {
		goto LABEL_27
	}
	v10 = int8(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v9 + 276))) + 2251))))
	if int32(v10) != 0 {
		if int32(v10) == 1 {
			v11 = float64(v25) * float64(nox_xxx_wizardMaximumMana_587000_312820)
		} else {
			if int32(v10) != 2 {
				goto LABEL_26
			}
			v11 = float64(v25) * float64(nox_xxx_conjurerMaxMana_587000_312804)
		}
	} else {
		v11 = float64(v25) * float64(nox_xxx_warriorMaxMana_587000_312788)
	}
	v24 = float32(v11)
	v3 = nox_float2int(v24)
LABEL_26:
	nox_xxx_playerManaAdd_4EEB80((*nox_object_t)(unsafe.Pointer(uintptr(a1))), int16(v3))
	nox_xxx_aud_501960(755, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
	result = 1
LABEL_27:
	if int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 12))))&0x40 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 540)))) != 0 {
		nox_xxx_removePoison_4EE9D0(a1)
		v12 = nox_xxx_spellGetAud44_424800(14, 1)
		nox_xxx_aud_501960(v12, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
		result = 1
	}
	v13 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 12))))
	if v13&0x100 != 0 {
		nox_xxx_buffApplyTo_4FF380((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 9, int16(int32(uint16(gameFPS()))*120), 3)
		result = 1
	}
	v14 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 12))))
	if v14&0x200 != 0 {
		nox_xxx_buffApplyTo_4FF380((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, int16(int32(uint16(gameFPS()))*120), 3)
		result = 1
	}
	v15 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 12))))
	if v15&0x400 != 0 {
		v16 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 56))))
		v17 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 60))))
		v26[0] = a1
		v26[1] = v16
		v26[2] = v17
		nox_xxx_spellAccept_4FD400(51, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), (*nox_object_t)(unsafe.Pointer(uintptr(a1))), (*nox_object_t)(unsafe.Pointer(uintptr(a1))), unsafe.Pointer(&v26[0]), v3)
		result = 1
	}
	v18 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 12))))
	if v18&0x800 != 0 {
		nox_xxx_buffApplyTo_4FF380((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 17, int16(int32(uint16(gameFPS()))*120), 3)
		result = 1
	}
	v19 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 12))))
	if v19&0x1000 != 0 {
		nox_xxx_buffApplyTo_4FF380((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 20, int16(int32(uint16(gameFPS()))*120), 3)
		result = 1
	}
	v20 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 12))))
	if v20&0x2000 != 0 {
		nox_xxx_buffApplyTo_4FF380((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 18, int16(int32(uint16(gameFPS()))*120), 3)
		result = 1
	}
	v21 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 12))))
	if v21&0x4000 != 0 {
		nox_xxx_buffApplyTo_4FF380((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 23, int16(int32(uint16(gameFPS()))*120), 3)
		result = 1
	}
	v22 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 12))))
	if (v22 & 0x8000) != 0 {
		nox_xxx_buffApplyTo_4FF380((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 21, int16(int32(uint16(gameFPS()))*120), 3)
		result = 1
	}
	if *(*uint32)(unsafe.Pointer(uintptr(v2 + 12)))&0x10000 != 0 {
		nox_xxx_buffApplyTo_4FF380((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 13, int16(int32(uint16(gameFPS()))*120), 3)
	} else if result == 0 {
		return 1
	}
	nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(v2))))
	return 1
}
func nox_xxx_useLesserFireballStaff_53F290(a1 int32, a2 *uint32) int32 {
	var (
		v2     int32
		result int32
		v4     float64
		v5     int32
		v6     float64
		v7     float32
		v8     float64
		v9     int32
		v10    *uint32
		v11    int32
		v12    uint8
		v13    float4
	)
	v2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), 4*184)))
	if v2 == 0 {
		return 1
	}
	if *(*uint32)(unsafe.Pointer(uintptr(v2 + 84))) == 0 {
		*(*uint32)(unsafe.Pointer(uintptr(v2 + 84))) = uint32(nox_xxx_getNameId_4E3AA0((*byte)(unsafe.Pointer(uintptr(v2 + 4)))))
	}
	if int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 109)))) == 0 || int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 108)))) != 0 {
		if gameFrame()-*(*uint32)(unsafe.Pointer(uintptr(v2 + 104))) >= (*(*uint32)(unsafe.Pointer(uintptr(v2 + 100))) - uint32(nox_xxx_itemCheckReadinessEffect_4E0960(int32(uintptr(unsafe.Pointer(a2)))))) {
			v4 = float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 176)))) + 4.0
			v5 = int32(*(*int16)(unsafe.Pointer(uintptr(a1 + 124)))) * 8
			v6 = v4 * float64(*mem_getFloatPtr(0x587000, uint32(v5)+194136))
			v7 = *(*float32)(unsafe.Pointer(uintptr(a1 + 60)))
			v13.field_0 = *(*float32)(unsafe.Pointer(uintptr(a1 + 56)))
			v13.field_4 = v7
			v13.field_8 = float32(v6 + float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 56)))))
			v8 = v4*float64(*mem_getFloatPtr(0x587000, uint32(v5)+194140)) + float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 60))))
			v13.field_8 = v13.field_8 + *(*float32)(unsafe.Pointer(uintptr(a1 + 80)))
			v13.field_C = float32(v8 + float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 84)))))
			if nox_xxx_mapTraceRay_535250(&v13, nil, nil, 5) == 0 {
				v13.field_8 = v13.field_0
				v13.field_C = v13.field_4
			}
			nox_xxx_wandShot_53F480(a1, int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 84)))), (*int32)(unsafe.Pointer(&v13.field_8)), (*uint32)(unsafe.Pointer(uintptr(*(*int16)(unsafe.Pointer(uintptr(a1 + 124)))))))
			if int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 96))))&1 != 0 {
				v9 = int32(*(*int16)(unsafe.Pointer(uintptr(a1 + 124)))) + 8
				if v9 >= 256 {
					v9 += int32((uint32(v9) >> 8) * 4294967040)
				}
				nox_xxx_wandShot_53F480(a1, int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 84)))), (*int32)(unsafe.Pointer(&v13.field_8)), (*uint32)(unsafe.Pointer(uintptr(v9))))
				v10 = (*uint32)(unsafe.Pointer(uintptr(int32(*(*int16)(unsafe.Pointer(uintptr(a1 + 124)))) - 8)))
				if int32(uintptr(unsafe.Pointer(v10))) < 0 {
					v10 = (*uint32)(unsafe.Add(unsafe.Pointer(v10), 4*uintptr(((math.MaxUint8-uint32(uintptr(unsafe.Pointer(v10))))>>8)*64)))
				}
				nox_xxx_wandShot_53F480(a1, int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 84)))), (*int32)(unsafe.Pointer(&v13.field_8)), v10)
			}
			nox_xxx_aud_501960(int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 88)))), (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
			if int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 109)))) != 0 {
				v11 = int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 109))))
				v12 = uint8(int8(int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 108)))) - 1))
				*(*uint8)(unsafe.Pointer(uintptr(v2 + 108))) = v12
				*(*uint32)(unsafe.Pointer(uintptr(v2 + 112))) = uint32(int32(v12) * 100 / v11)
				if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4 != 0 {
					nox_xxx_netReportCharges_4D82B0(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + 276))) + 2064)))), (*nox_object_t)(unsafe.Pointer(a2)), int8(v12), int8(v11))
				}
			}
			*(*uint32)(unsafe.Pointer(uintptr(v2 + 104))) = gameFrame()
			result = 1
		} else {
			result = 0
		}
	} else {
		nox_xxx_aud_501960(222, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
		result = 0
	}
	return result
}
func nox_xxx_wandShot_53F480(a1 int32, a2 int32, a3 *int32, a4 *uint32) *uint32 {
	var (
		result *uint32
		v5     *uint32
	)
	result = (*uint32)(unsafe.Pointer(nox_xxx_newObjectWithTypeInd_4E3450(a2)))
	v5 = result
	if result != nil {
		nox_xxx_createAt_4DAA50((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(result)))))), (*nox_object_t)(unsafe.Pointer(uintptr(a1))), *(*float32)(unsafe.Pointer(a3)), *((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(a3))), unsafe.Sizeof(float32(0))*1))))
		result = a4
		*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v5))), unsafe.Sizeof(uint16(0))*62))) = uint16(uintptr(unsafe.Pointer(a4)))
		*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v5))), unsafe.Sizeof(uint16(0))*63))) = uint16(uintptr(unsafe.Pointer(a4)))
		*((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v5))), unsafe.Sizeof(float32(0))*20))) = *mem_getFloatPtr(0x587000, uint32(uintptr(unsafe.Pointer(a4)))*8+194136) * *((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v5))), unsafe.Sizeof(float32(0))*136)))
		*((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v5))), unsafe.Sizeof(float32(0))*21))) = *mem_getFloatPtr(0x587000, uint32(uintptr(unsafe.Pointer(a4)))*8+194140) * *((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v5))), unsafe.Sizeof(float32(0))*136)))
		*((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v5))), unsafe.Sizeof(float32(0))*20))) = *(*float32)(unsafe.Pointer(uintptr(a1 + 80))) + *((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v5))), unsafe.Sizeof(float32(0))*20)))
		*((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v5))), unsafe.Sizeof(float32(0))*21))) = *(*float32)(unsafe.Pointer(uintptr(a1 + 84))) + *((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v5))), unsafe.Sizeof(float32(0))*21)))
	}
	return result
}
func nox_xxx_useWandCastSpell_53F4F0(a1 int32, a2 *uint32) int32 {
	var (
		v2  int32
		v4  int32
		v5  int32
		v6  int32
		v7  int32
		v8  int32
		v9  uint8
		v10 uint8
		v11 [3]int32
	)
	v2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), 4*184)))
	if *memmap.PtrUint32(0x5D4594, 2488732) == 0 {
		*memmap.PtrUint32(0x5D4594, 2488732) = uint32(nox_xxx_getNameId_4E3AA0(internCStr("ForceWand")))
	}
	if int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 109)))) != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 108)))) == 0 {
		nox_xxx_aud_501960(222, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
		return 0
	}
	if gameFrame()-*(*uint32)(unsafe.Pointer(uintptr(v2 + 104))) < (*(*uint32)(unsafe.Pointer(uintptr(v2 + 100))) - uint32(nox_xxx_itemCheckReadinessEffect_4E0960(int32(uintptr(unsafe.Pointer(a2))))*2)) {
		return 0
	}
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4 != 0 {
		v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
		*(*float32)(unsafe.Pointer(&v11[1])) = float32(float64(*(*int32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 276))) + 2284)))))
		*(*float32)(unsafe.Pointer(&v11[2])) = float32(float64(*(*int32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 276))) + 2288)))))
		v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 288))))
		if v5 != 0 {
			v11[0] = v5
			goto LABEL_13
		}
	} else {
		v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 60))))
		v11[1] = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 56))))
		v11[2] = v6
	}
	v11[0] = a1
LABEL_13:
	v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 92))))
	*(*uint32)(unsafe.Pointer(uintptr(v2 + 96))) |= 4
	if nox_xxx_spellAccept_4FD400(v7, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), (*nox_object_t)(unsafe.Pointer(uintptr(a1))), (*nox_object_t)(unsafe.Pointer(uintptr(a1))), unsafe.Pointer(&v11[0]), 4) != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(v2 + 104))) = gameFrame()
		v8 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), 4*2)))
		if (v8&0x1000) == 0 || (*(*uint32)(unsafe.Add(unsafe.Pointer(a2), 4*3))&0x4040000) == 0 {
			v9 = *(*uint8)(unsafe.Pointer(uintptr(v2 + 109)))
			if int32(v9) != 0 {
				v10 = uint8(int8(int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 108)))) - 1))
				*(*uint8)(unsafe.Pointer(uintptr(v2 + 108))) = v10
				*(*uint32)(unsafe.Pointer(uintptr(v2 + 112))) = uint32(int32(v10) * 100 / int32(v9))
				if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4 != 0 {
					nox_xxx_netReportCharges_4D82B0(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + 276))) + 2064)))), (*nox_object_t)(unsafe.Pointer(a2)), int8(v10), int8(v9))
				}
			}
		}
	}
	return 1
}
func nox_xxx_useFireWand_53F670(a1 int32, a2 int32) int32 {
	var (
		v2  int32
		v3  int32
		v4  float64
		v6  float32
		v7  float32
		v8  float32
		v9  float32
		v10 float32
		v11 float32
		v12 float32
		v13 float32
	)
	v2 = a1
	v8 = 0.0
	v3 = int32(*(*int16)(unsafe.Pointer(uintptr(a1 + 124)))) * 8
	v11 = float32(float64(*mem_getFloatPtr(0x587000, uint32(v3)+194136)**(*float32)(unsafe.Pointer(uintptr(a1 + 176)))) * 1.5)
	v10 = float32(float64(*mem_getFloatPtr(0x587000, uint32(v3)+194140)**(*float32)(unsafe.Pointer(uintptr(a1 + 176)))) * 1.5)
	for {
		v12 = float32(nox_common_randomFloat_416030(12.0, 25.0))
		v9 = float32(nox_common_randomFloat_416030(-2.0, 2.0) + float64(v12**mem_getFloatPtr(0x587000, uint32(int32(*(*int16)(unsafe.Pointer(uintptr(v2 + 124))))*8)+194136)))
		v13 = float32(nox_common_randomFloat_416030(-2.0, 2.0) + float64(v12**mem_getFloatPtr(0x587000, uint32(int32(*(*int16)(unsafe.Pointer(uintptr(v2 + 124))))*8)+194140)))
		v7 = v10 + *(*float32)(unsafe.Pointer(uintptr(v2 + 60)))
		v6 = v11 + *(*float32)(unsafe.Pointer(uintptr(v2 + 56)))
		nox_xxx_createSpark_54FD80(v6, v7, 1, 20, v9, v13, 0.0, 0)
		v4 = float64(v8) + 1.0
		v8 = float32(v4)
		if v4 >= 1.0 {
			break
		}
	}
	if (gameFrame() - *(*uint32)(unsafe.Pointer(uintptr(a2 + 136)))) > uint32(int32(gameFPS())) {
		nox_xxx_aud_501960(9, (*nox_object_t)(unsafe.Pointer(uintptr(v2))), 0, 0)
		*(*uint32)(unsafe.Pointer(uintptr(a2 + 136))) = gameFrame()
	}
	return 0
}
func nox_xxx_useRead_53F7C0(a1 int32, a2 int32) int32 {
	var (
		v2 int32
		v3 int32
	)
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4 != 0 {
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 736))))
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 256))))
		if (gameFrame()-uint32(v3) > (gameFPS()*3) || v3 == 0) && nox_xxx_mapCheck_537110((*nox_object_t)(unsafe.Pointer(uintptr(a1))), (*nox_object_t)(unsafe.Pointer(uintptr(a2)))) == 1 {
			nox_xxx_netPriMsgToPlayer_4DA2C0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), (*byte)(unsafe.Pointer(uintptr(v2))), 1)
			*(*uint32)(unsafe.Pointer(uintptr(v2 + 256))) = gameFrame()
		}
	}
	return 1
}
func sub_53F830(a1 int32, a2 int32) int32 {
	var (
		v2 int32
		v3 int32
		v4 int32
		v5 int32
		v6 int32
		v7 int32
		v8 int32
	)
	v2 = a1
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4 != 0 {
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
		v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 736))))
		v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 256))))
		if (gameFrame()-uint32(v5) > (gameFPS()*3) || v5 == 0) && nox_xxx_mapCheck_537110((*nox_object_t)(unsafe.Pointer(uintptr(a1))), (*nox_object_t)(unsafe.Pointer(uintptr(a2)))) == 1 {
			if sub_4D75E0() != 0 {
				v6 = nox_game_getQuestStage_4E3CC0()
				v7 = nox_server_questNextStageThreshold_4D74F0(v6)
				v8 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 276))))
				a1 = v7
				nox_xxx_netInformTextMsg_4DA0F0(int32(*(*uint8)(unsafe.Pointer(uintptr(v8 + 2064)))), 21, &a1)
			} else {
				nox_xxx_netPriMsgToPlayer_4DA2C0((*nox_object_t)(unsafe.Pointer(uintptr(v2))), internCStr("GeneralPrint:WarpClosed"), 1)
			}
			*(*uint32)(unsafe.Pointer(uintptr(v4 + 256))) = gameFrame()
		}
	}
	return 1
}
func nox_xxx_useByNetCode_53F8E0(a1 int32, a2 int32) int32 {
	var result int32
	if a2 == 0 {
		return 1
	}
	if *(*uint32)(unsafe.Pointer(uintptr(a2 + 732))) == 0 {
		return 1
	}
	result = sub_419E60((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
	if result != 1 {
		result = (*(*func(int32, int32) int32)(unsafe.Pointer(uintptr(a2 + 732))))(a1, a2)
	}
	return result
}
func sub_53F930(a1 int32, a2 int32) int32 {
	var (
		v2     int32
		v3     int32
		result int32
	)
	if (int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8)))) & 4) == 0 {
		return 0
	}
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	v3 = nox_xxx_guide_427010(*(**byte)(unsafe.Pointer(uintptr(a2 + 736))))
	if nox_common_gameFlags_check_40A5C0(4096) && int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 2251)))) != 2 {
		nox_xxx_netPriMsgToPlayer_4DA2C0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), internCStr("pickup.c:ObjectEquipClassFail"), 0)
		return 0
	}
	if *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + uint32(v3*4) + 4244))) != 0 {
		nox_xxx_netPriMsgToPlayer_4DA2C0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), internCStr("objcoll.c:AlreadyHaveGuide"), 0)
		result = 0
	} else {
		nox_xxx_awardBeastGuide_4FAE80_magic_plyrgide(a1, v3, 1)
		nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(a2))))
		result = 1
	}
	return result
}
func nox_xxx_useSpellReward_53F9E0(a1 int32, a2 int32) int32 {
	var (
		v2 *uint8
		v3 int32
		v4 int32
		v5 int32
		v6 int8
	)
	v2 = *(**uint8)(unsafe.Pointer(uintptr(a2 + 736)))
	v3 = 0
	v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	if (int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8)))) & 4) == 0 {
		return 0
	}
	v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 276))))
	v6 = int8(*(*uint8)(unsafe.Pointer(uintptr(v5 + 2251))))
	if int32(v6) != 1 && int32(v6) != 2 {
		nox_xxx_netPriMsgToPlayer_4DA2C0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), internCStr("use.c:SpellRewardClassFail"), 0)
		nox_xxx_aud_501960(925, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 2, int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 36)))))
		return 0
	}
	if nox_xxx_playerCheckSpellClass_57AEA0(int32(*(*uint8)(unsafe.Pointer(uintptr(v5 + 2251)))), int32(*v2)) != 0 {
		nox_xxx_netPriMsgToPlayer_4DA2C0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), internCStr("use.c:SpellRewardClassFail"), 0)
		nox_xxx_aud_501960(925, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 2, int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 36)))))
		return 0
	}
	if nox_common_gameFlags_check_40A5C0(6144) && *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 276))) + uint32(int32(*v2)*4) + 3696))) == 0 {
		v3 = 1
	}
	if nox_xxx_spellGrantToPlayer_4FB550((*nox_object_t)(unsafe.Pointer(uintptr(a1))), int32(*v2), 1, v3, 0) != 0 {
		nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(a2))))
	} else {
		nox_xxx_aud_501960(925, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 2, int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 36)))))
	}
	return 1
}
func nox_xxx_useAbilityReward_53FAE0(a1 int32, a2 int32) int32 {
	var (
		v2     *uint8
		v3     int32
		v4     int32
		result int32
	)
	v2 = *(**uint8)(unsafe.Pointer(uintptr(a2 + 736)))
	v3 = 0
	v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	if (int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8)))) & 4) == 0 {
		return 0
	}
	if int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 276))) + 2251)))) != 0 {
		nox_xxx_netPriMsgToPlayer_4DA2C0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), internCStr("pickup.c:ObjectEquipClassFail"), 0)
		nox_xxx_aud_501960(925, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 2, int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 36)))))
		result = 0
	} else {
		if nox_common_gameFlags_check_40A5C0(6144) && *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 276))) + uint32(int32(*v2)*4) + 3696))) == 0 {
			v3 = 1
		}
		if nox_xxx_abilityRewardServ_4FB9C0_ability(a1, int32(*v2), v3) != 0 {
			nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(a2))))
		} else {
			nox_xxx_aud_501960(925, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 2, int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 36)))))
		}
		result = 1
	}
	return result
}
func nox_xxx_respawnPlayerImpl_53FBC0(a1 *float32, a2 int32) *uint32 {
	var (
		v2     int32
		v3     int32
		v4     *float32
		result *uint32
		v6     *uint32
		v7     int32
		v8     int32
		v9     float32
		v10    float32
		v11    *int32
	)
	if *memmap.PtrUint32(0x5D4594, 2488736) == 0 {
		nox_xxx_createCorpse_53FCA0()
	}
	v2 = nox_xxx_math_509EA0(a2)
	v3 = 0
	v4 = mem_getFloatPtr(0x587000, uint32(v2*88)+280376)
	v11 = mem_getI32Ptr(0x5D4594, uint32(v2*44)+2488740)
	for {
		result = (*uint32)(unsafe.Pointer(nox_xxx_newObjectWithTypeInd_4E3450(*v11)))
		v6 = result
		if result == nil {
			break
		}
		if dword_5d4594_2650652 != 0 {
			if nox_common_gameFlags_check_40A5C0(0x2000) {
				v7 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), 4*4)))
				*((*uint8)(unsafe.Pointer(&v7))) = uint8(int8(v7 | 0x40))
				*(*uint32)(unsafe.Add(unsafe.Pointer(v6), 4*4)) = uint32(v7)
			}
		}
		v10 = *(*float32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(float32(0))*1)) + *(*float32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(float32(0))*1))
		v9 = *v4 + *a1
		nox_xxx_createAt_4DAA50((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v6)))))), nil, v9, v10)
		v8 = nox_common_randomInt_415FA0(10, 20)
		result = (*uint32)(unsafe.Pointer(uintptr(nox_xxx_unitSetDecayTime_511660((*nox_object_t)(unsafe.Pointer(v6)), int32(gameFPS()*uint32(v8))))))
		v3++
		v4 = (*float32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(float32(0))*2))
		v11 = (*int32)(unsafe.Add(unsafe.Pointer(v11), 4*1))
		if v3 >= 11 {
			break
		}
	}
	return result
}
func nox_xxx_createCorpse_53FCA0() {
	var (
		i  int32
		v1 *uint8
		v2 [32]byte
	)
	for i = 0; i < 9; i++ {
		switch i {
		case 0:
			v1 = (*uint8)(memmap.PtrOff(0x587000, 281216))
		case 1:
			v1 = (*uint8)(memmap.PtrOff(0x587000, 281220))
		case 2:
			v1 = (*uint8)(memmap.PtrOff(0x587000, 281224))
		case 3:
			v1 = (*uint8)(memmap.PtrOff(0x587000, 281228))
		case 5:
			v1 = (*uint8)(memmap.PtrOff(0x587000, 281232))
		case 6:
			v1 = (*uint8)(memmap.PtrOff(0x587000, 281236))
		case 7:
			v1 = (*uint8)(memmap.PtrOff(0x587000, 281240))
		case 8:
			v1 = (*uint8)(memmap.PtrOff(0x587000, 281244))
		default:
			continue
		}
		nox_sprintf(&v2[0], internCStr("CorpseSkull%s"), v1)
		*memmap.PtrUint32(0x5D4594, uint32(i*44)+2488740) = uint32(nox_xxx_getNameId_4E3AA0(&v2[0]))
		nox_sprintf(&v2[0], internCStr("CorpseRibCage%s"), v1)
		*memmap.PtrUint32(0x5D4594, uint32(i*44)+2488744) = uint32(nox_xxx_getNameId_4E3AA0(&v2[0]))
		nox_sprintf(&v2[0], internCStr("CorpsePelvis%s"), v1)
		*memmap.PtrUint32(0x5D4594, uint32(i*44)+2488748) = uint32(nox_xxx_getNameId_4E3AA0(&v2[0]))
		nox_sprintf(&v2[0], internCStr("CorpseLeftLowerLeg%s"), v1)
		*memmap.PtrUint32(0x5D4594, uint32(i*44)+2488752) = uint32(nox_xxx_getNameId_4E3AA0(&v2[0]))
		nox_sprintf(&v2[0], internCStr("CorpseLeftUpperLeg%s"), v1)
		*memmap.PtrUint32(0x5D4594, uint32(i*44)+2488756) = uint32(nox_xxx_getNameId_4E3AA0(&v2[0]))
		nox_sprintf(&v2[0], internCStr("CorpseLeftLowerArm%s"), v1)
		*memmap.PtrUint32(0x5D4594, uint32(i*44)+2488760) = uint32(nox_xxx_getNameId_4E3AA0(&v2[0]))
		nox_sprintf(&v2[0], internCStr("CorpseLeftUpperArm%s"), v1)
		*memmap.PtrUint32(0x5D4594, uint32(i*44)+2488764) = uint32(nox_xxx_getNameId_4E3AA0(&v2[0]))
		nox_sprintf(&v2[0], internCStr("CorpseRightLowerLeg%s"), v1)
		*memmap.PtrUint32(0x5D4594, uint32(i*44)+2488768) = uint32(nox_xxx_getNameId_4E3AA0(&v2[0]))
		nox_sprintf(&v2[0], internCStr("CorpseRightUpperLeg%s"), v1)
		*memmap.PtrUint32(0x5D4594, uint32(i*44)+2488772) = uint32(nox_xxx_getNameId_4E3AA0(&v2[0]))
		nox_sprintf(&v2[0], internCStr("CorpseRightLowerArm%s"), v1)
		*memmap.PtrUint32(0x5D4594, uint32(i*44)+2488776) = uint32(nox_xxx_getNameId_4E3AA0(&v2[0]))
		nox_sprintf(&v2[0], internCStr("CorpseRightUpperArm%s"), v1)
		*memmap.PtrUint32(0x5D4594, uint32(i*44)+2488780) = uint32(nox_xxx_getNameId_4E3AA0(&v2[0]))
	}
	*memmap.PtrUint32(0x5D4594, 2488736) = 1
}
func nox_xxx_castPixies_540440(a1 int32, a2 int32, a3 int32, a4 int32, a5 int32, a6 int32) int32 {
	var (
		v6  int32
		v7  int32
		v8  int32
		v9  int32
		v10 int32
		v11 float32
		v12 float64
		v13 *uint32
		v14 *int32
		v15 int32
		v17 float4
		v18 float32
		v19 int32
	)
	v6 = a4
	v7 = int32(*memmap.PtrUint32(0x5D4594, 2489140))
	v18 = float32(float64(*(*float32)(unsafe.Pointer(uintptr(a4 + 176)))) + 4.0)
	if *memmap.PtrUint32(0x5D4594, 2489140) == 0 {
		v7 = nox_xxx_getNameId_4E3AA0(internCStr("Pixie"))
		*memmap.PtrUint32(0x5D4594, 2489140) = uint32(v7)
	}
	v8 = nox_xxx_unitIsUnitTT_4E7C80((*nox_object_t)(unsafe.Pointer(uintptr(a3))), v7)
	if v8 < int32(int64(nox_xxx_gamedataGetFloatTable_419D70(internCStr("PixieCount"), a6-1))) {
		v9 = int32(uint32(uint64(int64(nox_xxx_gamedataGetFloatTable_419D70(internCStr("PixieCount"), a6-1))) - uint64(v8)))
		if v9 > 0 {
			v19 = v9
			for {
				v10 = nox_common_randomInt_415FA0(0, math.MaxUint8)
				v11 = *(*float32)(unsafe.Pointer(uintptr(v6 + 60)))
				v12 = float64(v18**mem_getFloatPtr(0x587000, uint32(v10*8)+194136) + *(*float32)(unsafe.Pointer(uintptr(v6 + 56))))
				v17.field_0 = *(*float32)(unsafe.Pointer(uintptr(v6 + 56)))
				v17.field_4 = v11
				v17.field_8 = float32(v12)
				v17.field_C = v18**mem_getFloatPtr(0x587000, uint32(v10*8)+194140) + *(*float32)(unsafe.Pointer(uintptr(v6 + 60)))
				if nox_xxx_mapTraceRay_535250(&v17, nil, nil, 5) != 0 {
					v13 = (*uint32)(unsafe.Pointer(nox_xxx_newObjectByTypeID_4E3810(internCStr("Pixie"))))
					if v13 != nil {
						v14 = (*int32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v13), 4*187)))))
						nox_xxx_createAt_4DAA50((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v13)))))), (*nox_object_t)(unsafe.Pointer(uintptr(a3))), v17.field_8, v17.field_C)
						*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v13))), unsafe.Sizeof(uint16(0))*63))) = uint16(int16(v10))
						*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v13))), unsafe.Sizeof(uint16(0))*62))) = uint16(int16(v10))
						*(*uint32)(unsafe.Add(unsafe.Pointer(v13), 4*20)) = 0
						*(*uint32)(unsafe.Add(unsafe.Pointer(v13), 4*21)) = 0
						*(*int32)(unsafe.Add(unsafe.Pointer(v14), 4*1)) = int32(uintptr(unsafe.Pointer(nox_xxx_spellFlySearchTarget_540610(nil, (*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v13)))))), 32, 600.0, 0, (*nox_object_t)(unsafe.Pointer(uintptr(a3)))))))
						*v14 = a3
						*(*int32)(unsafe.Add(unsafe.Pointer(v14), 4*3)) = a1
						*(*uint32)(unsafe.Add(unsafe.Pointer(v13), 4*39)) = *(*uint32)(unsafe.Pointer(uintptr(v6 + 56)))
						*(*uint32)(unsafe.Add(unsafe.Pointer(v13), 4*40)) = *(*uint32)(unsafe.Pointer(uintptr(v6 + 60)))
						*(*int32)(unsafe.Add(unsafe.Pointer(v14), 4*5)) = int32(gameFrame() + gameFPS()*uint32(nox_common_randomInt_415FA0(30, 90)))
						*(*int32)(unsafe.Add(unsafe.Pointer(v14), 4*6)) = int32(gameFrame())
					}
				}
				v19--
				if v19 == 0 {
					break
				}
			}
		}
		v15 = nox_xxx_spellGetAud44_424800(a1, 0)
		nox_xxx_aud_501960(v15, (*nox_object_t)(unsafe.Pointer(uintptr(v6))), 0, 0)
	}
	return 1
}
func sub_5408A0(a1 int32) int32 {
	var v1 int32
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + uint32((*(*byte)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + 544)))+23)*24)))))
	return bool2int32(v1 >= 18 && v1 <= 20)
}
func nox_xxx_mobCastInversion_5408D0(a1 int32) int32 {
	var (
		v1 int32
		v2 int32
		v3 *int32
		v4 *uint32
		v5 int32
		v7 float32
		v8 int32
		v9 [137]int32
	)
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	v8 = 0
	if (*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))) & 0x1000000) == 0 {
		return 0
	}
	if nox_xxx_monsterCanCast_534300((*nox_object_t)(unsafe.Pointer(uintptr(a1)))) == 0 {
		return 0
	}
	if gameFrame() < uint32(*(*int32)(unsafe.Pointer(uintptr(v1 + 1452)))) {
		return 0
	}
	if sub_5408A0(a1) != 0 {
		return 0
	}
	*memmap.PtrUint32(0x5D4594, 2489156) = 0
	v7 = float32(nox_xxx_gamedataGetFloat_419D40(internCStr("InversionRange")) * 0.5)
	nox_xxx_getMissilesInCircle_518170((*float2)(unsafe.Pointer(uintptr(a1+56))), v7, unsafe.Pointer(funAddr(nox_xxx_unitIsMagicMissile_540B60)), (*nox_object_t)(unsafe.Pointer(uintptr(a1))))
	if *memmap.PtrUint32(0x5D4594, 2489156) == 0 {
		return 0
	}
	v2 = 1
	v3 = &v9[0]
	v4 = (*uint32)(unsafe.Pointer(uintptr(v1 + 1492)))
	for {
		if *v4&0x8000000 != 0 {
			if nox_xxx_spellHasFlags_424A50(v2, 16) {
				*v3 = v2
				v3 = (*int32)(unsafe.Add(unsafe.Pointer(v3), 4*1))
				v8++
			}
		}
		v2++
		v4 = (*uint32)(unsafe.Add(unsafe.Pointer(v4), 4*1))
		if v2 >= 137 {
			break
		}
	}
	if v8 == 0 {
		return 0
	}
	v5 = nox_common_randomInt_415FA0(0, v8-1)
	nox_xxx_monsterCast_540A30((*nox_object_t)(unsafe.Pointer(uintptr(a1))), v9[v5], (*nox_object_t)(unsafe.Pointer(uintptr(a1))))
	*(*uint32)(unsafe.Pointer(uintptr(v1 + 1452))) = gameFrame() + uint32(nox_common_randomInt_415FA0(int32(*(*uint16)(unsafe.Pointer(uintptr(v1 + 1448)))), int32(*(*uint16)(unsafe.Pointer(uintptr(v1 + 1450))))))
	return 1
}
func nox_xxx_unitIsMagicMissile_540B60(a1 int32, a2 int32) int32 {
	var result int32
	result = a1
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&1 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 12))))&2 != 0 {
		result = a2
		if *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + 4))) == uint32(a2) {
			*memmap.PtrUint32(0x5D4594, 2489156) = 1
		}
	}
	return result
}
func nox_xxx_monsterBuffSelf_540B90(a1 int32) int32 {
	var (
		v1  int32
		v2  int32
		v3  int32
		v4  *int32
		v5  *uint32
		v6  int32
		v7  *int32
		v8  int32
		v9  int32
		v11 int32
		v12 [137]int32
	)
	v1 = 0
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	v11 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	if (*(*uint32)(unsafe.Pointer(uintptr(a1 + 16)))&0x1000000) == 0 || nox_xxx_monsterCanCast_534300((*nox_object_t)(unsafe.Pointer(uintptr(a1)))) == 0 || gameFrame() < uint32(*(*int32)(unsafe.Pointer(uintptr(v2 + 1460)))) || sub_5408A0(a1) != 0 {
		return 0
	}
	v3 = 1
	v4 = &v12[0]
	v5 = (*uint32)(unsafe.Pointer(uintptr(v2 + 1492)))
	for {
		if *v5&0x10000000 != 0 {
			if nox_xxx_spellHasFlags_424A50(v3, 16) {
				*v4 = v3
				v1++
				v4 = (*int32)(unsafe.Add(unsafe.Pointer(v4), 4*1))
			}
		}
		v3++
		v5 = (*uint32)(unsafe.Add(unsafe.Pointer(v5), 4*1))
		if v3 >= 137 {
			break
		}
	}
	if v1 == 0 {
		return 0
	}
	v6 = 0
	if v1 > 0 {
		v7 = &v12[0]
		for {
			v8 = a1
			if sub_540CE0(a1, *v7) != 0 {
				return 0
			}
			v6++
			v7 = (*int32)(unsafe.Add(unsafe.Pointer(v7), 4*1))
			if v6 >= v1 {
				goto LABEL_17
			}
		}
	}
	v8 = a1
LABEL_17:
	v9 = nox_common_randomInt_415FA0(0, v1-1)
	nox_xxx_monsterCast_540A30((*nox_object_t)(unsafe.Pointer(uintptr(v8))), v12[v9], (*nox_object_t)(unsafe.Pointer(uintptr(v8))))
	*(*uint32)(unsafe.Pointer(uintptr(v11 + 1460))) = gameFrame() + uint32(nox_common_randomInt_415FA0(int32(*(*uint16)(unsafe.Pointer(uintptr(v11 + 1456)))), int32(*(*uint16)(unsafe.Pointer(uintptr(v11 + 1458))))))
	return 1
}
func sub_540CE0(a1 int32, a2 int32) int32 {
	var (
		i      int32
		result int32
	)
	for i = 0; i < 32; i++ {
		if nox_xxx_getEnchantSpell_424920(i) == a2 {
			break
		}
	}
	if i == 32 {
		result = 0
	} else {
		result = nox_xxx_testUnitBuffs_4FF350((*nox_object_t)(unsafe.Pointer(uintptr(a1))), int8(i))
	}
	return result
}
func sub_540D20(a1 int32) int32 {
	return bool2int32(a1 >= 75 && a1 <= 114)
}
func sub_540D40(a1 int32) int32 {
	var v1 int32
	v1 = int32(uintptr(nox_xxx_spellCastedFirst_4FE930()))
	if v1 == 0 {
		return 0
	}
	for sub_540D20(int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 4))))) == 0 || *(*uint32)(unsafe.Pointer(uintptr(v1 + 16))) != uint32(a1) {
		v1 = int32(uintptr(nox_xxx_spellCastedNext_4FE940(unsafe.Pointer(uintptr(v1)))))
		if v1 == 0 {
			return 0
		}
	}
	return 1
}
func nox_xxx_mobCastRelated2_540D90(a1 int32, a2 int32) int32 {
	var (
		v2  int32
		v3  int32
		v4  int32
		v5  *int32
		v6  *uint32
		v7  int32
		v8  int32
		v9  uint8
		v11 int32
		v12 int32
		v13 [137]int32
	)
	v2 = 0
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	v12 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	if *(*uint32)(unsafe.Pointer(uintptr(a1 + 16)))&0x1000000 != 0 && nox_xxx_monsterCanCast_534300((*nox_object_t)(unsafe.Pointer(uintptr(a1)))) != 0 && gameFrame() >= uint32(*(*int32)(unsafe.Pointer(uintptr(v3 + 1476)))) && sub_5408A0(a1) == 0 {
		v4 = 1
		v5 = &v13[0]
		v11 = 1
		v6 = (*uint32)(unsafe.Pointer(uintptr(v3 + 1492)))
		for {
			if *v6&0x40000000 != 0 && nox_xxx_spellHasFlags_424A50(v4, 16) {
				*v5 = v4
				v2++
				v5 = (*int32)(unsafe.Add(unsafe.Pointer(v5), 4*1))
				if sub_540D20(v4) == 0 {
					v11 = 0
				}
			}
			v4++
			v6 = (*uint32)(unsafe.Add(unsafe.Pointer(v6), 4*1))
			if v4 >= 137 {
				break
			}
		}
		if v2 != 0 {
			for {
				for {
					v7 = nox_common_randomInt_415FA0(0, v2-1)
					v8 = v13[v7]
					if sub_540D20(v13[v7]) != 0 {
						if sub_540D40(a1) != 0 {
							if v11 != 0 {
								return 0
							}
						} else {
							v9 = uint8(int8(bool2int32(nox_xxx_checkSummonedCreaturesLimit_500D70((*nox_object_t)(unsafe.Pointer(uintptr(a1))), v8-74))))
							if int32(v9) != 0 {
								goto LABEL_19
							}
							if v11 != 0 {
								return 0
							}
						}
						break
					}
				LABEL_19:
					if v7 >= 0 {
						nox_xxx_monsterCast_540A30((*nox_object_t)(unsafe.Pointer(uintptr(a1))), v13[v7], (*nox_object_t)(unsafe.Pointer(uintptr(a2))))
						*(*uint32)(unsafe.Pointer(uintptr(v12 + 1476))) = gameFrame() + uint32(nox_common_randomInt_415FA0(int32(*(*uint16)(unsafe.Pointer(uintptr(v12 + 1472)))), int32(*(*uint16)(unsafe.Pointer(uintptr(v12 + 1474))))))
						return 1
					}
				}
			}
		}
	}
	return 0
}
func nox_xxx_monsterCastOffensive_540F20(a1 int32, a2 int32) int32 {
	var (
		v2 int32
		v3 int32
		v4 *int32
		v5 *uint32
		v6 int32
		v8 int32
		v9 [137]int32
	)
	v8 = 0
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	if (*(*uint32)(unsafe.Pointer(uintptr(a1 + 16)))&0x1000000) == 0 || nox_xxx_monsterCanCast_534300((*nox_object_t)(unsafe.Pointer(uintptr(a1)))) == 0 || gameFrame() < uint32(*(*int32)(unsafe.Pointer(uintptr(v2 + 1468)))) || sub_5408A0(a1) != 0 {
		return 0
	}
	v3 = 1
	v4 = &v9[0]
	v5 = (*uint32)(unsafe.Pointer(uintptr(v2 + 1492)))
	for {
		if *v5&0x20000000 != 0 {
			if nox_xxx_spellHasFlags_424A50(v3, 16) {
				*v4 = v3
				v4 = (*int32)(unsafe.Add(unsafe.Pointer(v4), 4*1))
				v8++
			}
		}
		v3++
		v5 = (*uint32)(unsafe.Add(unsafe.Pointer(v5), 4*1))
		if v3 >= 137 {
			break
		}
	}
	if v8 == 0 {
		return 0
	}
	v6 = nox_common_randomInt_415FA0(0, v8-1)
	nox_xxx_monsterCast_540A30((*nox_object_t)(unsafe.Pointer(uintptr(a1))), v9[v6], (*nox_object_t)(unsafe.Pointer(uintptr(a2))))
	*(*uint32)(unsafe.Pointer(uintptr(v2 + 1468))) = gameFrame() + uint32(nox_common_randomInt_415FA0(int32(*(*uint16)(unsafe.Pointer(uintptr(v2 + 1464)))), int32(*(*uint16)(unsafe.Pointer(uintptr(v2 + 1466))))))
	return 1
}
func nox_xxx_mobCastRelated_541050(a1 int32) int32 {
	var (
		v1  int32
		v2  int32
		v3  int32
		v4  *int32
		v5  *uint32
		v6  int32
		v7  *int32
		v8  int32
		v9  int32
		v11 int32
		v12 [137]int32
	)
	v1 = 0
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	v11 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	if (*(*uint32)(unsafe.Pointer(uintptr(a1 + 16)))&0x1000000) == 0 || nox_xxx_monsterCanCast_534300((*nox_object_t)(unsafe.Pointer(uintptr(a1)))) == 0 || gameFrame() < uint32(*(*int32)(unsafe.Pointer(uintptr(v2 + 1484)))) || sub_5408A0(a1) != 0 {
		return 0
	}
	v3 = 1
	v4 = &v12[0]
	v5 = (*uint32)(unsafe.Pointer(uintptr(v2 + 1492)))
	for {
		if int32(*v5) < 0 {
			if nox_xxx_spellHasFlags_424A50(v3, 16) {
				*v4 = v3
				v1++
				v4 = (*int32)(unsafe.Add(unsafe.Pointer(v4), 4*1))
			}
		}
		v3++
		v5 = (*uint32)(unsafe.Add(unsafe.Pointer(v5), 4*1))
		if v3 >= 137 {
			break
		}
	}
	if v1 == 0 {
		return 0
	}
	v6 = 0
	if v1 > 0 {
		v7 = &v12[0]
		for {
			v8 = a1
			if sub_540CE0(a1, *v7) != 0 {
				return 0
			}
			v6++
			v7 = (*int32)(unsafe.Add(unsafe.Pointer(v7), 4*1))
			if v6 >= v1 {
				goto LABEL_17
			}
		}
	}
	v8 = a1
LABEL_17:
	v9 = nox_common_randomInt_415FA0(0, v1-1)
	nox_xxx_monsterCast_540A30((*nox_object_t)(unsafe.Pointer(uintptr(v8))), v12[v9], (*nox_object_t)(unsafe.Pointer(uintptr(v8))))
	*(*uint32)(unsafe.Pointer(uintptr(v11 + 1484))) = gameFrame() + uint32(nox_common_randomInt_415FA0(int32(*(*uint16)(unsafe.Pointer(uintptr(v11 + 1480)))), int32(*(*uint16)(unsafe.Pointer(uintptr(v11 + 1482))))))
	return 1
}
func nox_xxx_mobHealSomeone_5411A0(a3p *nox_object_t) int32 {
	var (
		a3 int32 = int32(uintptr(unsafe.Pointer(a3p)))
		v1 int32
		v2 int32
		v3 *uint16
		v4 uint16
		v6 int32
		v7 float64
		a1 float4
	)
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a3 + 748))))
	if (*(*uint32)(unsafe.Pointer(uintptr(a3 + 16)))&0x1000000) == 0 || nox_xxx_monsterCanCast_534300((*nox_object_t)(unsafe.Pointer(uintptr(a3)))) == 0 || int32(uint8(gameFrame()))&0x1F != 0 || sub_5408A0(a3) != 0 {
		return 0
	}
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 1440))))
	if v2&0x800 != 0 {
		v3 = *(**uint16)(unsafe.Pointer(uintptr(a3 + 556)))
		v4 = *(*uint16)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint16(0))*2))
		if int32(v4) != 0 {
			if uint32(*v3) < uint32(int32(v4)>>1) {
				nox_xxx_monsterCast_540A30((*nox_object_t)(unsafe.Pointer(uintptr(a3))), 41, (*nox_object_t)(unsafe.Pointer(uintptr(a3))))
				return 0
			}
		}
	}
	if (v2 & 0x1000) == 0 {
		return 0
	}
	v6 = bool2int32(nox_common_gameFlags_check_40A5C0(4096))
	dword_5d4594_2489160 = 0
	if v6 != 0 {
		v7 = 640
	} else {
		v7 = 250
	}
	a1.field_0 = float32(float64(*(*float32)(unsafe.Pointer(uintptr(a3 + 56)))) - v7)
	a1.field_4 = float32(float64(*(*float32)(unsafe.Pointer(uintptr(a3 + 60)))) - v7)
	a1.field_8 = float32(v7 + float64(*(*float32)(unsafe.Pointer(uintptr(a3 + 56)))))
	a1.field_C = float32(v7 + float64(*(*float32)(unsafe.Pointer(uintptr(a3 + 60)))))
	nox_xxx_getUnitsInRect_517C10(&a1, unsafe.Pointer(funAddr(nox_xxx_mobMayHealThis_5412A0)), unsafe.Pointer(uintptr(a3)))
	if dword_5d4594_2489160 == 0 {
		return 0
	}
	nox_xxx_monsterCast_540A30((*nox_object_t)(unsafe.Pointer(uintptr(a3))), 41, (*nox_object_t)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_2489160))))))
	return 1
}
func nox_xxx_mobMayHealThis_5412A0(a1 *float32, a2 int32) {
	var v2 int32
	if (*float32)(unsafe.Pointer(uintptr(a2))) != a1 && nox_xxx_unitIsEnemyTo_5330C0((*nox_object_t)(unsafe.Pointer(uintptr(a2))), (*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1))))))) == 0 {
		if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), 4*139))) != 0 {
			v2 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), 4*4))))
			if (v2 & 0x8000) == 0 {
				if nox_xxx_unitCanInteractWith_5370E0((*nox_object_t)(unsafe.Pointer(uintptr(a2))), (*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), 0) != 0 {
					if uint32(**((**uint16)(unsafe.Add(unsafe.Pointer((**uint16)(unsafe.Pointer(a1))), unsafe.Sizeof((*uint16)(nil))*139)))) < uint32(int32(*(*uint16)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), 4*139))) + 4))))>>1) {
						dword_5d4594_2489160 = uint32(uintptr(unsafe.Pointer(a1)))
					}
				}
			}
		}
	}
}
func nox_xxx_mobCast_541300(a1 int32, a2 *uint32, a3 int32) int8 {
	var (
		v3 int32
		v4 int32
	)
	v3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), 4*187)))
	if *(*uint32)(unsafe.Pointer(uintptr(v3 + 1440)))&0x20000 != 0 {
		nox_xxx_mobMorphToPlayer_4FAAF0(a2)
	}
	nox_xxx_mobCalcDir_533CC0(int32(uintptr(unsafe.Pointer(a2))), (*float32)(unsafe.Pointer(uintptr(a3+4))))
	nox_xxx_castSpellByUser_4FDD20(a1, (*nox_object_t)(unsafe.Pointer(a2)), unsafe.Pointer(uintptr(a3)))
	v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 1440))))
	if uint32(v4)&0x20000 != 0 {
		*((*uint8)(unsafe.Pointer(&v4))) = uint8(nox_xxx_mobMorphFromPlayer_4FAAC0(a2))
	}
	return int8(v4)
}
func nox_xxx_mobActionCast_5413B0(a1p *nox_object_t, a2 int32) int8 {
	var (
		a1 *uint32 = (*uint32)(unsafe.Pointer(a1p))
		v2 int32
		v3 *int32
		v4 int32
		v5 int32
		v6 int32
		v8 [3]int32
	)
	v2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*187)))
	v3 = (*int32)(unsafe.Pointer(uintptr(v2 + int32((*(*byte)(unsafe.Pointer(uintptr(v2 + 544)))+23)*24))))
	*((*uint8)(unsafe.Pointer(&v4))) = *(*uint8)(unsafe.Pointer(uintptr(v2 + 482)))
	if int32(uint8(int8(v4))) == 0 {
		*((*uint8)(unsafe.Pointer(&v4))) = *(*uint8)(unsafe.Pointer(uintptr(v2 + 481)))
		if uint32(uint8(int8(v4))) == *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 484))) + 216))) {
			if a2 != 0 {
				*((*uint8)(unsafe.Pointer(&v4))) = uint8(int8(a2 - 1))
				if a2 != 1 {
					return int8(v4)
				}
				v5 = *(*int32)(unsafe.Add(unsafe.Pointer(v3), 4*3))
				v6 = *(*int32)(unsafe.Add(unsafe.Pointer(v3), 4*4))
				v8[0] = 0
				v8[1] = v5
				v8[2] = v6
			} else {
				v8[0] = *(*int32)(unsafe.Add(unsafe.Pointer(v3), 4*3))
				nox_xxx_mobCastRandomRecoil_541490(int32(uintptr(unsafe.Pointer(a1))), (*float32)(unsafe.Pointer(uintptr(v8[0]))), (*float2)(unsafe.Pointer(&v8[1])))
			}
			v4 = nox_xxx_testUnitBuffs_4FF350((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), 29)
			if v4 == 0 {
				*((*uint8)(unsafe.Pointer(&v4))) = uint8(nox_xxx_mobCast_541300(*(*int32)(unsafe.Add(unsafe.Pointer(v3), 4*1)), a1, int32(uintptr(unsafe.Pointer(&v8[0])))))
			}
		} else if int32(uint8(int8(v4))) == 1 {
			v4 = int32(uintptr(nox_xxx_monsterGetSoundSet_424300((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))))))
			if v4 != 0 {
				nox_xxx_aud_501960(int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 56)))), (*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), 0, 0)
			}
		}
	}
	return int8(v4)
}
func nox_xxx_mobCastRandomRecoil_541490(a1 int32, a2 *float32, a3 *float2) {
	var (
		v3 float64
		v4 float32
		v5 float32
		v6 float32
	)
	v5 = float32(1.0 - float64(*(*float32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + 1320)))))
	*a3 = *((*float2)(unsafe.Add(unsafe.Pointer((*float2)(unsafe.Pointer(a2))), unsafe.Sizeof(float2{})*7)))
	v4 = float32(float64(v5) + 1.0)
	v3 = nox_common_randomFloat_416030(v5, v4)
	a3.field_0 = float32(float64(a3.field_0) - v3*float64(*(*float32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(float32(0))*20)))*6.0)
	a3.field_4 = float32(float64(a3.field_4) - v3*float64(*(*float32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(float32(0))*21)))*6.0)
	v6 = float32(float64(v5)*0.80000001 + 0.2)
	a3.field_0 = float32(nox_common_randomFloat_416030(-60.0, 60.0)*float64(v6) + float64(a3.field_0))
	a3.field_4 = float32(nox_common_randomFloat_416030(-60.0, 60.0)*float64(v6) + float64(a3.field_4))
}
func sub_541630(a1 int32, a2 int32) int8 {
	var (
		v2 *uint32
		v3 int32
	)
	v2 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + uint32((*(*byte)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + 544)))+23)*24))))
	if *v2 == 20 {
		v3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), 4*1)))
		*((*uint8)(unsafe.Pointer(&v2))) = uint8(int8(a2))
		if v3 == a2 {
			*((*uint8)(unsafe.Pointer(&v2))) = uint8(int8(nox_xxx_monsterPopAction_50A160((*nox_object_t)(unsafe.Pointer(uintptr(a1))))))
		}
	}
	return int8(uintptr(unsafe.Pointer(v2)))
}
func sub_542BF0(a1 int32, a2 int32, a3 int32) *byte {
	var (
		v3     **int32
		v4     *int32
		v5     *byte
		v6     *byte
		v7     *byte
		v8     int32
		v9     *byte
		v10    *byte
		v11    func(*int32) int32
		v12    *byte
		v13    *byte
		v14    *byte
		v15    *byte
		v16    *byte
		v17    *byte
		v18    *byte
		v19    *byte
		v20    *byte
		v21    *byte
		v22    *byte
		v23    *byte
		v24    *byte
		v25    *byte
		v26    *byte
		v27    *byte
		v28    *byte
		v29    *byte
		v30    *byte
		v31    *byte
		v32    *byte
		v33    *byte
		v34    *byte
		v35    *byte
		v36    *byte
		v37    *byte
		v38    *byte
		v39    *byte
		v40    *byte
		v41    *byte
		result *byte
		i      *byte
		v44    *byte
		v45    *byte
		v46    *byte
		v47    *byte
		v48    int32
	)
	v3 = (**int32)(sub_5049D0())
	v48 = int32(uintptr(unsafe.Pointer(v3)))
	if v3 != nil {
		for {
			v4 = *v3
			if (uint32(*(*int32)(unsafe.Add(unsafe.Pointer(*v3), 4*4))) & 0x80000000) == 0x80000000 {
				if *v4 != 0 {
					v5 = sub_543620(*v4, a1)
					v6 = (*byte)(libc.Realloc(unsafe.Pointer(uintptr(*v4)), libc.StrLen(v5)+1))
					*v4 = int32(uintptr(unsafe.Pointer(v6)))
					libc.StrCpy(v6, v5)
				}
				v7 = nox_script_objCallbackName_508CB0((*nox_object_t)(unsafe.Pointer(v4)), 14)
				if v7 != nil {
					v8 = a3
					if libc.StrLen(v7) != 0 {
						v9 = sub_5435C0(int32(uintptr(unsafe.Pointer(v7))), a1, a2, a3)
						sub_509120((*uint32)(unsafe.Pointer(v4)), 14, v9)
					}
				} else {
					v8 = a3
				}
				v10 = nox_xxx_getUnitName_4E39D0((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v4)))))))
				v11 = asFunc(nox_objectTypeGetXfer(v10), (*func(*int32) int32)(nil)).(func(*int32) int32)
				if funAddr(v11) == funAddr(nox_xxx_unitTriggerXfer_4F4E50) {
					v12 = nox_script_objCallbackName_508CB0((*nox_object_t)(unsafe.Pointer(v4)), 1)
					if libc.StrLen(v12) != 0 {
						v13 = sub_5435C0(int32(uintptr(unsafe.Pointer(v12))), a1, a2, v8)
						sub_509120((*uint32)(unsafe.Pointer(v4)), 1, v13)
					}
					v14 = nox_script_objCallbackName_508CB0((*nox_object_t)(unsafe.Pointer(v4)), 2)
					if libc.StrLen(v14) != 0 {
						v15 = sub_5435C0(int32(uintptr(unsafe.Pointer(v14))), a1, a2, v8)
						sub_509120((*uint32)(unsafe.Pointer(v4)), 2, v15)
					}
					v16 = nox_script_objCallbackName_508CB0((*nox_object_t)(unsafe.Pointer(v4)), 0)
					if libc.StrLen(v16) != 0 {
						v44 = sub_5435C0(int32(uintptr(unsafe.Pointer(v16))), a1, a2, v8)
						sub_509120((*uint32)(unsafe.Pointer(v4)), 0, v44)
					}
				} else if funAddr(v11) == funAddr(nox_xxx_XFerMonster_528DB0) {
					v17 = nox_script_objCallbackName_508CB0((*nox_object_t)(unsafe.Pointer(v4)), 3)
					if libc.StrLen(v17) != 0 {
						v18 = sub_5435C0(int32(uintptr(unsafe.Pointer(v17))), a1, a2, v8)
						sub_509120((*uint32)(unsafe.Pointer(v4)), 3, v18)
					}
					v19 = nox_script_objCallbackName_508CB0((*nox_object_t)(unsafe.Pointer(v4)), 5)
					if libc.StrLen(v19) != 0 {
						v20 = sub_5435C0(int32(uintptr(unsafe.Pointer(v19))), a1, a2, v8)
						sub_509120((*uint32)(unsafe.Pointer(v4)), 5, v20)
					}
					v21 = nox_script_objCallbackName_508CB0((*nox_object_t)(unsafe.Pointer(v4)), 4)
					if libc.StrLen(v21) != 0 {
						v22 = sub_5435C0(int32(uintptr(unsafe.Pointer(v21))), a1, a2, v8)
						sub_509120((*uint32)(unsafe.Pointer(v4)), 4, v22)
					}
					v23 = nox_script_objCallbackName_508CB0((*nox_object_t)(unsafe.Pointer(v4)), 6)
					if libc.StrLen(v23) != 0 {
						v24 = sub_5435C0(int32(uintptr(unsafe.Pointer(v23))), a1, a2, v8)
						sub_509120((*uint32)(unsafe.Pointer(v4)), 6, v24)
					}
					v25 = nox_script_objCallbackName_508CB0((*nox_object_t)(unsafe.Pointer(v4)), 7)
					if libc.StrLen(v25) != 0 {
						v26 = sub_5435C0(int32(uintptr(unsafe.Pointer(v25))), a1, a2, v8)
						sub_509120((*uint32)(unsafe.Pointer(v4)), 7, v26)
					}
					v27 = nox_script_objCallbackName_508CB0((*nox_object_t)(unsafe.Pointer(v4)), 8)
					if libc.StrLen(v27) != 0 {
						v28 = sub_5435C0(int32(uintptr(unsafe.Pointer(v27))), a1, a2, v8)
						sub_509120((*uint32)(unsafe.Pointer(v4)), 8, v28)
					}
					v29 = nox_script_objCallbackName_508CB0((*nox_object_t)(unsafe.Pointer(v4)), 9)
					if libc.StrLen(v29) != 0 {
						v30 = sub_5435C0(int32(uintptr(unsafe.Pointer(v29))), a1, a2, v8)
						sub_509120((*uint32)(unsafe.Pointer(v4)), 9, v30)
					}
					v31 = nox_script_objCallbackName_508CB0((*nox_object_t)(unsafe.Pointer(v4)), 10)
					if libc.StrLen(v31) != 0 {
						v32 = sub_5435C0(int32(uintptr(unsafe.Pointer(v31))), a1, a2, v8)
						sub_509120((*uint32)(unsafe.Pointer(v4)), 10, v32)
					}
					v33 = nox_script_objCallbackName_508CB0((*nox_object_t)(unsafe.Pointer(v4)), 11)
					if libc.StrLen(v33) != 0 {
						v45 = sub_5435C0(int32(uintptr(unsafe.Pointer(v33))), a1, a2, v8)
						sub_509120((*uint32)(unsafe.Pointer(v4)), 11, v45)
					}
				} else if funAddr(v11) == funAddr(nox_xxx_XFerHole_4F51D0) {
					v34 = nox_script_objCallbackName_508CB0((*nox_object_t)(unsafe.Pointer(v4)), 12)
					if libc.StrLen(v34) != 0 {
						v46 = sub_5435C0(int32(uintptr(unsafe.Pointer(v34))), a1, a2, v8)
						sub_509120((*uint32)(unsafe.Pointer(v4)), 12, v46)
					}
				} else if funAddr(v11) == funAddr(nox_xxx_XFerMonsterGen_4F7130) {
					v35 = nox_script_objCallbackName_508CB0((*nox_object_t)(unsafe.Pointer(v4)), 15)
					if v35 != nil && libc.StrLen(v35) != 0 {
						v36 = sub_5435C0(int32(uintptr(unsafe.Pointer(v35))), a1, a2, v8)
						sub_509120((*uint32)(unsafe.Pointer(v4)), 15, v36)
					}
					v37 = nox_script_objCallbackName_508CB0((*nox_object_t)(unsafe.Pointer(v4)), 16)
					if v37 != nil && libc.StrLen(v37) != 0 {
						v38 = sub_5435C0(int32(uintptr(unsafe.Pointer(v37))), a1, a2, v8)
						sub_509120((*uint32)(unsafe.Pointer(v4)), 16, v38)
					}
					v39 = nox_script_objCallbackName_508CB0((*nox_object_t)(unsafe.Pointer(v4)), 18)
					if v39 != nil && libc.StrLen(v39) != 0 {
						v40 = sub_5435C0(int32(uintptr(unsafe.Pointer(v39))), a1, a2, v8)
						sub_509120((*uint32)(unsafe.Pointer(v4)), 18, v40)
					}
					v41 = nox_script_objCallbackName_508CB0((*nox_object_t)(unsafe.Pointer(v4)), 17)
					if v41 != nil && libc.StrLen(v41) != 0 {
						v47 = sub_5435C0(int32(uintptr(unsafe.Pointer(v41))), a1, a2, v8)
						sub_509120((*uint32)(unsafe.Pointer(v4)), 17, v47)
					}
				}
				*(*int32)(unsafe.Add(unsafe.Pointer(v4), 4*4)) &= math.MaxInt32
				v3 = (**int32)(unsafe.Pointer(uintptr(v48)))
			}
			v48 = sub_5049E0(int32(uintptr(unsafe.Pointer(v3))))
			if v48 == 0 {
				break
			}
			v3 = (**int32)(unsafe.Pointer(uintptr(v48)))
		}
	}
	result = (*byte)(nox_xxx_waypointGetList_579860())
	for i = result; result != nil; i = result {
		if *((*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(i))), 4*120))) < 0 {
			if int32(libc.StrLen((*byte)(unsafe.Add(unsafe.Pointer(i), 16)))) > 0 {
				libc.StrCpy((*byte)(unsafe.Add(unsafe.Pointer(i), 16)), sub_543620(int32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(i), 16))))), a1))
			}
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), 4*120))) &= math.MaxInt32
		}
		result = (*byte)(unsafe.Pointer(uintptr(nox_xxx_waypointNext_579870(int32(uintptr(unsafe.Pointer(i)))))))
	}
	return result
}
func sub_5435C0(a1 int32, a2 int32, a3 int32, a4 int32) *byte {
	nox_sprintf((*byte)(memmap.PtrOff(0x5D4594, 2489164)), internCStr("%s%%%d%%%d%%%d"), a1, a2, a3, a4)
	libc.StrLen((*byte)(memmap.PtrOff(0x5D4594, 2489164)))
	libc.StrCpy((*byte)(memmap.PtrOff(0x5D4594, 2489164)), internCStr("ERROR_NAME_TOO_LONG!"))
	return (*byte)(memmap.PtrOff(0x5D4594, 2489164))
}
func sub_543620(a1 int32, a2 int32) *byte {
	nox_sprintf((*byte)(memmap.PtrOff(0x5D4594, 2489164)), internCStr("%s%%%d"), a1, a2)
	libc.StrLen((*byte)(memmap.PtrOff(0x5D4594, 2489164)))
	libc.StrCpy((*byte)(memmap.PtrOff(0x5D4594, 2489164)), internCStr("ERROR_NAME_TOO_LONG!"))
	return (*byte)(memmap.PtrOff(0x5D4594, 2489164))
}
func sub_543680(a1 *float32) int32 {
	var (
		result int32
		a2     float2
		v3     int2
		v4     [8]int32
	)
	if dword_5d4594_3835356 == math.MaxUint8 {
		return 1
	}
	result = nox_xxx_mapGenFixCoords_4D3D90((*float2)(unsafe.Pointer(a1)), &a2)
	if result != 0 {
		v4[0] = int32(*memmap.PtrUint32(0x973F18, 35912))
		dword_5d4594_3835352 = 1
		v4[1] = int32(dword_5d4594_3835348)
		v4[6] = int32(dword_5d4594_3835356)
		v4[7] = int32(dword_5d4594_3835360)
		v3.field_0 = int32(int64(a2.field_0))
		v3.field_4 = int32(int64(a2.field_4))
		result = sub_5437E0(&v3.field_0, int32(uintptr(unsafe.Pointer(&v4[0]))), 46)
		dword_5d4594_3835352 = 0
	}
	return result
}
func sub_5437E0(a1 *int32, a2 int32, a3 int32) int32 {
	var (
		v3  float64
		v4  float64
		v5  int32
		v6  float64
		v7  int32
		v8  int32
		v9  int32
		v10 int32
		v11 int32
		v12 int32
		v13 int32
		v14 int32
		v15 int32
		v16 int32
		v17 *uint8
		v18 int32
		v19 int32
		v20 *int32
		v21 int32
		v23 int32
		v24 float32
	)
	v3 = float64(*a1) + 11.5
	v4 = float64(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*1)))
	v5 = int32(int64(v3 * 0.021739131))
	v23 = int32(int64(v3 * 0.021739131))
	v6 = v4 + 11.5
	v24 = float32(v6)
	v7 = int32(int64(v6 * 0.021739131))
	a1 = (*int32)(unsafe.Pointer(uintptr(int64(v6 * 0.021739131))))
	v8 = int32(int64(v3)) % 46
	*memmap.PtrUint32(0x973F18, 22200) = 0
	v9 = int32(int64(v24) % 46)
	dword_5d4594_2487248 = 0
	if v5 <= 0 || v5 >= math.MaxInt8 || v7 <= 0 || v7 >= math.MaxInt8 {
		v11 = a3
	} else if v8 <= v9 {
		if a3-v8 <= v9 {
			v15 = int32(uintptr(unsafe.Pointer(*(**obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(ptr_5D4594_2650668), unsafe.Sizeof((*obj_5D4594_2650668_t)(nil))*uintptr(v5))))))
			v11 = int32(*(*uint32)(unsafe.Pointer(uintptr(v15 + v7*44 + 24))))
			sub_51DD50(v5, v7, 2, int32(*(*uint32)(unsafe.Pointer(uintptr(v15 + v7*44 + 24)))))
		} else {
			v13 = v5 - 1
			v14 = int32(uintptr(unsafe.Pointer(*(**obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(ptr_5D4594_2650668), unsafe.Sizeof((*obj_5D4594_2650668_t)(nil))*uintptr(v13))))))
			v11 = int32(*(*uint32)(unsafe.Pointer(uintptr(v14 + v7*44 + 4))))
			sub_51DD50(v13, v7, 1, int32(*(*uint32)(unsafe.Pointer(uintptr(v14 + v7*44 + 4)))))
		}
	} else if a3-v8 <= v9 {
		v12 = int32(uintptr(unsafe.Pointer(*(**obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(ptr_5D4594_2650668), unsafe.Sizeof((*obj_5D4594_2650668_t)(nil))*uintptr(v5))))))
		v11 = int32(*(*uint32)(unsafe.Pointer(uintptr(v12 + v7*44 + 4))))
		sub_51DD50(v5, v7, 1, int32(*(*uint32)(unsafe.Pointer(uintptr(v12 + v7*44 + 4)))))
	} else {
		v10 = int32(uintptr(unsafe.Pointer(*(**obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(ptr_5D4594_2650668), unsafe.Sizeof((*obj_5D4594_2650668_t)(nil))*uintptr(v5))))))
		v11 = int32(*(*uint32)(unsafe.Pointer(uintptr(v10 + v7*44 - 20))))
		sub_51DD50(v5, v7-1, 2, int32(*(*uint32)(unsafe.Pointer(uintptr(v10 + v7*44 - 20)))))
	}
	v16 = 0
	if dword_5d4594_2487248 > 0 {
		v17 = (*uint8)(memmap.PtrOff(0x973F18, 16204))
		for {
			v18 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v17))), 4*1))))
			v19 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v17))), -int(4*1)))))
			v20 = *(**int32)(unsafe.Pointer(v17))
			v23 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v17))), -int(4*1)))))
			a1 = v20
			a3 = v18
			if v18&2 != 0 {
				sub_51DD50(v19, int32(uintptr(unsafe.Pointer(v20))), 1, v11)
				sub_51DD50(v23, int32(uintptr(unsafe.Pointer(a1)))+1, 1, v11)
				sub_51DD50(v23-1, int32(uintptr(unsafe.Pointer(a1))), 1, v11)
				sub_51DD50(v23-1, int32(uintptr(unsafe.Pointer(a1)))+1, 1, v11)
			} else if v18&1 != 0 {
				sub_51DD50(v19+1, int32(uintptr(unsafe.Pointer(v20))), 2, v11)
				sub_51DD50(v23+1, int32(uintptr(unsafe.Pointer(a1)))-1, 2, v11)
				sub_51DD50(v23, int32(uintptr(unsafe.Pointer(a1))), 2, v11)
				sub_51DD50(v23, int32(uintptr(unsafe.Pointer(a1)))-1, 2, v11)
			}
			v16++
			v17 = (*uint8)(unsafe.Add(unsafe.Pointer(v17), 12))
			if v16 >= *(*int32)(unsafe.Pointer(&dword_5d4594_2487248)) {
				break
			}
		}
	}
	if sub_51DE30((*uint32)(unsafe.Pointer(&v23)), (*uint32)(unsafe.Pointer(&a1)), (*uint32)(unsafe.Pointer(&a3))) != 0 {
		v21 = a2
		for {
			if a3&2 != 0 {
				sub_543BC0(v23, int32(uintptr(unsafe.Pointer(a1))), 1, v11, v21, 1)
				sub_543BC0(v23, int32(uintptr(unsafe.Pointer(a1)))+1, 1, v11, v21, 4)
				sub_543BC0(v23-1, int32(uintptr(unsafe.Pointer(a1))), 1, v11, v21, 3)
				sub_543BC0(v23-1, int32(uintptr(unsafe.Pointer(a1)))+1, 1, v11, v21, 6)
				sub_543BC0(v23, int32(uintptr(unsafe.Pointer(a1)))-1, 2, v11, v21, 0)
				sub_543BC0(v23+1, int32(uintptr(unsafe.Pointer(a1))), 2, v11, v21, 2)
				sub_543BC0(v23-1, int32(uintptr(unsafe.Pointer(a1))), 2, v11, v21, 5)
				sub_543BC0(v23, int32(uintptr(unsafe.Pointer(a1)))+1, 2, v11, v21, 7)
			} else if a3&1 != 0 {
				sub_543BC0(v23+1, int32(uintptr(unsafe.Pointer(a1))), 2, v11, v21, 4)
				sub_543BC0(v23+1, int32(uintptr(unsafe.Pointer(a1)))-1, 2, v11, v21, 1)
				sub_543BC0(v23, int32(uintptr(unsafe.Pointer(a1))), 2, v11, v21, 6)
				sub_543BC0(v23, int32(uintptr(unsafe.Pointer(a1)))-1, 2, v11, v21, 3)
				sub_543BC0(v23, int32(uintptr(unsafe.Pointer(a1)))-1, 1, v11, v21, 0)
				sub_543BC0(v23+1, int32(uintptr(unsafe.Pointer(a1))), 1, v11, v21, 2)
				sub_543BC0(v23-1, int32(uintptr(unsafe.Pointer(a1))), 1, v11, v21, 5)
				sub_543BC0(v23, int32(uintptr(unsafe.Pointer(a1)))+1, 1, v11, v21, 7)
			}
			if sub_51DE30((*uint32)(unsafe.Pointer(&v23)), (*uint32)(unsafe.Pointer(&a1)), (*uint32)(unsafe.Pointer(&a3))) == 0 {
				break
			}
		}
	}
	return bool2int32(*memmap.PtrUint32(0x973F18, 22200) == 0)
}
func sub_543BC0(a1 int32, a2 int32, a3 int32, a4 int32, a5 int32, a6 int32) {
	var v7 int32
	if a1 > 0 && a1 < math.MaxInt8 {
		if a2 > 0 && a2 < math.MaxInt8 {
			v7 = a3
			if ((a3&1) == 0 || a2 != 1) && ((a3&2) == 0 || a1 != 1) {
				if a3&2 != 0 {
					if uint32(a4) == *(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(ptr_5D4594_2650668), unsafe.Sizeof((*obj_5D4594_2650668_t)(nil))*uintptr(a1)))))) + uint32(a2*44) + 24))) {
						return
					}
					v7 = a3
				}
				if (a3&1) == 0 || uint32(a4) != *(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(ptr_5D4594_2650668), unsafe.Sizeof((*obj_5D4594_2650668_t)(nil))*uintptr(a1)))))) + uint32(a2*44) + 4))) {
					if a5 != 0 {
						*(*uint32)(unsafe.Pointer(uintptr(a5 + 28))) = uint32(a6)
						sub_51DA70(a1, a2, a5, v7, 1)
					}
				}
			}
		}
	}
}
func nox_xxx_tile_543C50(a1 *uint32, a2 int32, a3 int32, a4 int32, a5 int32, a6 int32) int32 {
	var (
		v6     *uint32
		v7     *uint32
		v8     **uint32
		result int32
	)
	_ = result
	var v10 **uint32
	var v11 int32
	var v12 int32
	var v13 *uint32
	var v14 **uint32
	var j *uint32
	var v16 *uint32
	var v17 int32
	var v18 int32
	var v19 *uint32
	var v20 int32
	var i int32
	var v22 int32
	if a2 == math.MaxUint8 || a4 == math.MaxUint8 {
		v19 = a1
		v20 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*4)))
		if v20 != 0 {
			for i = int32(*(*uint32)(unsafe.Pointer(uintptr(v20 + 16)))); i != 0; i = int32(*(*uint32)(unsafe.Pointer(uintptr(i + 16)))) {
				v19 = (*uint32)(unsafe.Pointer(uintptr(v20)))
				v20 = i
			}
			nox_xxx_tileFreeTileOne_4221E0(unsafe.Pointer(uintptr(v20)))
			*(*uint32)(unsafe.Add(unsafe.Pointer(v19), 4*4)) = 0
		}
		return 1
	}
	v6 = a1
	if *a1 == math.MaxUint8 {
		return 1
	}
	v7 = a1
	if a6 != 0 {
		v10 = (**uint32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*4))))
		v11 = 0
		if *(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*4)) == 0 {
			v12 = nox_xxx_mapGenEdge_543EB0(a4, a5)
			*(*uint32)(unsafe.Add(unsafe.Pointer(v7), 4*4)) = uint32(uintptr(unsafe.Pointer(nox_xxx_tileListAddNewSubtile_422160(a2, a3, a4, v12))))
			return 1
		}
		for {
			v7 = *v10
			if **v10 == uint32(a2) && *(*uint32)(unsafe.Add(unsafe.Pointer(v7), 4*1)) == uint32(a3) && *(*uint32)(unsafe.Add(unsafe.Pointer(v7), 4*2)) == uint32(a4) && sub_543E60(int32(uintptr(unsafe.Pointer(v7))), a5) != 0 {
				v11 = 1
			}
			v10 = (**uint32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(v7), 4*4))))
			if *(*uint32)(unsafe.Add(unsafe.Pointer(v7), 4*4)) == 0 {
				break
			}
		}
		if v11 == 0 {
			v12 = nox_xxx_mapGenEdge_543EB0(a4, a5)
			*(*uint32)(unsafe.Add(unsafe.Pointer(v7), 4*4)) = uint32(uintptr(unsafe.Pointer(nox_xxx_tileListAddNewSubtile_422160(a2, a3, a4, v12))))
			return 1
		}
		for {
			v13 = v6
			v22 = 0
			if *(*uint32)(unsafe.Add(unsafe.Pointer(v6), 4*4)) == 0 {
				break
			}
			for {
				v13 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v13), 4*4)))))
				if *v13 == uint32(a2) && *(*uint32)(unsafe.Add(unsafe.Pointer(v13), 4*1)) == uint32(a3) && *(*uint32)(unsafe.Add(unsafe.Pointer(v13), 4*2)) == uint32(a4) {
					v14 = (**uint32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(v13), 4*4))))
					for j = v13; *(*uint32)(unsafe.Add(unsafe.Pointer(j), 4*4)) != 0; v14 = (**uint32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(j), 4*4)))) {
						v16 = *v14
						if **v14 == uint32(a2) && *(*uint32)(unsafe.Add(unsafe.Pointer(v16), 4*1)) == uint32(a3) && (func() bool {
							v17 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v16), 4*2)))
							return v17 == a4
						}()) && (func() int32 {
							v18 = sub_411490(v17, int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v16), 4*3))))
							return sub_543E60(int32(uintptr(unsafe.Pointer(v13))), v18)
						}()) != 0 {
							v22 = 1
							*v14 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v16), 4*4)))))
							nox_xxx_tileFreeTileOne_4221E0(unsafe.Pointer(v16))
						} else {
							j = v16
						}
					}
					v6 = a1
				}
				if *(*uint32)(unsafe.Add(unsafe.Pointer(v13), 4*4)) == 0 {
					break
				}
			}
			if v22 == 0 {
				return 1
			}
		}
		return 1
	}
	v8 = (**uint32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*4))))
	if *(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*4)) != 0 {
		for {
			v7 = *v8
			if **v8 == uint32(a2) && *(*uint32)(unsafe.Add(unsafe.Pointer(v7), 4*1)) == uint32(a3) && *(*uint32)(unsafe.Add(unsafe.Pointer(v7), 4*2)) == uint32(a4) && *(*uint32)(unsafe.Add(unsafe.Pointer(v7), 4*3)) == uint32(a5) {
				break
			}
			v8 = (**uint32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(v7), 4*4))))
			if *(*uint32)(unsafe.Add(unsafe.Pointer(v7), 4*4)) == 0 {
				*(*uint32)(unsafe.Add(unsafe.Pointer(v7), 4*4)) = uint32(uintptr(unsafe.Pointer(nox_xxx_tileListAddNewSubtile_422160(a2, a3, a4, a5))))
				return 1
			}
		}
		return 0
	} else {
		*(*uint32)(unsafe.Add(unsafe.Pointer(v7), 4*4)) = uint32(uintptr(unsafe.Pointer(nox_xxx_tileListAddNewSubtile_422160(a2, a3, a4, a5))))
		return 1
	}
}
func sub_543E60(a1 int32, a2 int32) int32 {
	var (
		v2 int32
		v4 int32
	)
	v2 = a2 + sub_411490(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))), int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12)))))*12
	if *memmap.PtrUint32(0x587000, uint32(v2*4)+282736) == math.MaxUint8 {
		return 0
	}
	v4 = nox_xxx_mapGenEdge_543EB0(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))), int32(*memmap.PtrUint32(0x587000, uint32(v2*4)+282736)))
	if uint32(v4) != *(*uint32)(unsafe.Pointer(uintptr(a1 + 12))) {
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 12))) = uint32(v4)
	}
	return 1
}
func nox_xxx_mapGenEdge_543EB0(a1 int32, a2 int32) int32 {
	var (
		v2     int32
		v3     int32
		result int32
	)
	v2 = int32(*memmap.PtrUint8(0x85B3FC, uint32(a1*60+28696)))
	v3 = int32(*memmap.PtrUint8(0x85B3FC, uint32(a1*60+28697)))
	if v2 == 3 && v3 == int32(*memmap.PtrUint8(0x85B3FC, uint32(a1*60+28696))) {
		return a2
	}
	result = a2
	if a2 != 0 {
		switch a2 {
		case 1:
			result = nox_common_randomInt_415FA0(1, v2-2)
		case 2:
			result = v2 - 1
		case 3:
			result = v2 + nox_common_randomInt_415FA0(0, v3-3)*2
		case 4:
			result = v2 + nox_common_randomInt_415FA0(0, v3-3)*2 + 1
		case 5:
			result = v2 + v3*2 - 4
		case 6:
			result = nox_common_randomInt_415FA0(1, v2-2) + v2 + v3*2 - 4
		case 7:
			result = (v2+v3)*2 - 5
		default:
			result = a2 + (v2+v3)*2 - 12
		}
	}
	return result
}
func sub_543FB0(a1 *byte) int32 {
	var (
		v1 int32
		i  *byte
	)
	if a1 == nil {
		return -1
	}
	v1 = 0
	if *(*int32)(unsafe.Pointer(&dword_5d4594_251572)) <= 0 {
		return -1
	}
	for i = (*byte)(memmap.PtrOff(0x85B3FC, 28644)); libc.StrCmp(i, a1) != 0; i = (*byte)(unsafe.Add(unsafe.Pointer(i), 60)) {
		if func() int32 {
			p := &v1
			*p++
			return *p
		}() >= *(*int32)(unsafe.Pointer(&dword_5d4594_251572)) {
			return -1
		}
	}
	return v1
}
func sub_544020(a1 *byte) int32 {
	var (
		v1     int32
		result int32
	)
	dword_5d4594_2489436 = 0
	if nox_strcmpi(internCStr("NONE"), a1) != 0 {
		v1 = sub_543FB0(a1)
		result = nox_xxx_tileCheckByte3_544070(v1)
	} else {
		dword_5d4594_3835356 = math.MaxUint8
		result = 1
	}
	return result
}
func nox_xxx_tileCheckByte3_544070(a1 int32) int32 {
	if a1 < 0 || a1 >= *(*int32)(unsafe.Pointer(&dword_5d4594_251572)) {
		return 0
	}
	dword_5d4594_3835356 = uint32(a1)
	dword_5d4594_2489436 = 1
	return 1
}
func nox_xxx_tileCheckByte4_5440A0(a1 int32) int32 {
	if dword_5d4594_2489436 == 0 {
		return 1
	}
	if *(*int32)(unsafe.Pointer(&dword_5d4594_3835356)) == -1 || a1 < 0 || a1 >= int32(*memmap.PtrUint16(0x85B3FC, uint32(a1*60+28688))) {
		return 0
	}
	dword_5d4594_3835360 = uint32(a1)
	return 1
}
func nox_xxx_tileSubtile_544310(a1 *float2) int32 {
	var (
		v1     float64
		v2     float64
		v3     int32
		v4     int32
		v5     int32
		v6     int32
		result int32
		v8     [8]int32
		v9     float32
	)
	dword_5d4594_3835352 = 1
	v1 = float64(a1.field_0) + 11.5
	v8[1] = int32(dword_5d4594_3835348)
	v8[0] = int32(*memmap.PtrUint32(0x973F18, 35912))
	v8[2] = 0
	*((*uint8)(unsafe.Pointer(&v8[5]))) = 0
	v8[3] = -1
	v8[4] = -1
	v8[6] = int32(dword_5d4594_3835356)
	v8[7] = int32(dword_5d4594_3835360)
	v2 = float64(a1.field_4) + 11.5
	v3 = int32(int64(v1 * 0.021739131))
	v9 = float32(v2)
	v4 = int32(int64(v2 * 0.021739131))
	v5 = int32(int64(v1)) % 46
	v6 = int32(int64(v9) % 46)
	if *memmap.PtrUint32(0x973F18, 35912) == math.MaxUint8 {
		result = sub_51D9C0(v3, v4, v5, v6, 0)
	} else {
		result = sub_51D9C0(v3, v4, v5, v6, int32(uintptr(unsafe.Pointer(&v8[0]))))
	}
	dword_5d4594_3835352 = 0
	return result
}
func nox_xxx_mobActionMoveTo_5443F0(a1 int32) int8 {
	var (
		v1  int32
		v2  int32
		v4  int32
		v5  float64
		v6  float64
		v7  float64
		v8  float64
		v9  float64
		v10 int8
		v11 int32
		v12 *int32
		v13 *int32
		v14 float32
	)
	v1 = a1
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	if nox_xxx_monsterIsMoveing_534320(a1) == 0 {
		return int8(nox_xxx_monsterPopAction_50A160((*nox_object_t)(unsafe.Pointer(uintptr(a1)))))
	}
	v4 = v2 + int32((*(*byte)(unsafe.Pointer(uintptr(v2 + 544)))+23)*24)
	if sub_50A040(a1) == 3 {
		v5 = float64(*(*float32)(unsafe.Pointer(uintptr(v2 + 1316)))) * 3.0
		v14 = float32(v5)
		v6 = v5 + 30.0
		v7 = float64(*(*float32)(unsafe.Pointer(uintptr(v4 + 4))) - *(*float32)(unsafe.Pointer(uintptr(v1 + 56))))
		v8 = float64(*(*float32)(unsafe.Pointer(uintptr(v4 + 8))) - *(*float32)(unsafe.Pointer(uintptr(v1 + 60))))
		v9 = v8*v8 + v7*v7
		if v9 >= float64(v14*v14) {
			if v9 > v6*v6 {
				sub_534750(v1)
			}
		} else {
			sub_534780(v1)
		}
	}
	if nox_xxx_creatureSetMovePath_50D5A0(v1) == 1 {
		v10 = int8(*(*uint8)(unsafe.Pointer(uintptr(v2 + 284))))
		v11 = bool2int32(int32(v10) == 2 || int32(v10) == 1 && gameFrame()-*(*uint32)(unsafe.Pointer(uintptr(v2 + 540))) < (gameFPS()*5))
		if int32(v10) == 1 {
			*(*uint32)(unsafe.Pointer(uintptr(v2 + 540))) = gameFrame()
		}
		if int32(v10) == 0 && sub_547F10() == 0 && *(*uint32)(unsafe.Pointer(uintptr(v4 + 12))) == 0 {
			nox_xxx_mobCalcDir_533CC0(v1, (*float32)(unsafe.Pointer(uintptr(v4+4))))
			nox_xxx_monsterPopAction_50A160((*nox_object_t)(unsafe.Pointer(uintptr(v1))))
		}
		if v11 != 0 {
			v12 = (*int32)(nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(v1))), 41, internCStr(__FILE__), __LINE__))
			if v12 != nil {
				*(*int32)(unsafe.Add(unsafe.Pointer(v12), 4*1)) = int32(gameFrame() + uint32(nox_common_randomInt_415FA0(int32(gameFPS()*2), int32(gameFPS()*4))))
			}
			nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(v1))), 29, internCStr(__FILE__), __LINE__)
			*(*uint32)(unsafe.Pointer(uintptr(v2 + 1440))) |= 0x200000
		}
		if int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 284)))) != 0 {
			v13 = (*int32)(nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(v1))), 1, internCStr(__FILE__), __LINE__))
			if v13 != nil {
				*(*int32)(unsafe.Add(unsafe.Pointer(v13), 4*1)) = int32(gameFrame() + uint32(nox_common_randomInt_415FA0(int32(gameFPS()>>1), int32(gameFPS()))))
			}
		}
	}
	return nox_xxx_monsterMoveAudio_534030(v1)
}
func nox_xxx_mobActionMoveToFar_5445C0(a1 *int32) int8 {
	var (
		v1 int32
		v2 int32
		v3 *int32
		v4 int32
	)
	v1 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*187))
	if sub_5343C0(int32(uintptr(unsafe.Pointer(a1)))) == 0 && nox_xxx_monsterCanAttackAtWill_534390((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1))))))) == 0 || (func() int32 {
		v2 = sub_545E60((*nox_object_t)(unsafe.Pointer(a1)))
		return v2
	}()) == 0 {
		if nox_xxx_monsterCanAttackAtWill_534390((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1))))))) != 0 {
			if *(*uint32)(unsafe.Pointer(uintptr(v1 + 1196))) != 0 {
				v3 = (*int32)(nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), 15, internCStr(__FILE__), __LINE__))
				if v3 != nil {
					v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 1196))))
					*(*int32)(unsafe.Add(unsafe.Pointer(v3), 4*1)) = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 56))))
					*(*int32)(unsafe.Add(unsafe.Pointer(v3), 4*2)) = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 60))))
					*(*int32)(unsafe.Add(unsafe.Pointer(v3), 4*3)) = int32(gameFrame())
				}
			}
		}
		*((*uint8)(unsafe.Pointer(&v2))) = uint8(nox_xxx_mobActionMoveTo_5443F0(int32(uintptr(unsafe.Pointer(a1)))))
	}
	return int8(v2)
}
func nox_xxx_mobActionDodge_544640(a1 int32) {
	var (
		v1 *float32
		v2 int32
		v3 float64
		v4 float64
		v5 float64
		v6 float64
		v7 float32
		v8 float32
	)
	v1 = (*float32)(unsafe.Pointer(uintptr(a1)))
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	if nox_xxx_monsterIsMoveing_534320(a1) != 0 {
		if nox_xxx_testUnitBuffs_4FF350((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 3) == 0 && nox_xxx_testUnitBuffs_4FF350((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 5) == 0 && nox_xxx_testUnitBuffs_4FF350((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 28) == 0 {
			v3 = float64(*(*float32)(unsafe.Pointer(uintptr(v2 + int32((*(*byte)(unsafe.Pointer(uintptr(v2 + 544)))+23)*24) + 4))) - *(*float32)(unsafe.Pointer(uintptr(a1 + 56))))
			v4 = float64(*(*float32)(unsafe.Pointer(uintptr(v2 + int32((*(*byte)(unsafe.Pointer(uintptr(v2 + 544)))+23)*24) + 8))) - *(*float32)(unsafe.Pointer(uintptr(a1 + 60))))
			v7 = float32(v4)
			v5 = math.Sqrt(v4*float64(v7)+v3*v3) + 9.9999997e-05
			v8 = float32(v5)
			if v5 >= 8.0 {
				v6 = float64(*(*float32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 484))) + 96))) * *(*float32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(float32(0))*136)))
				*(*float32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(float32(0))*136)) = float32(v6)
				*(*float32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(float32(0))*22)) = float32(v6 * v3 / float64(v8))
				*(*float32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(float32(0))*23)) = v7 * *(*float32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(float32(0))*136)) / v8
			} else {
				nox_xxx_monsterPopAction_50A160((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v1)))))))
			}
		}
	} else {
		nox_xxx_monsterPopAction_50A160((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
	}
}
func sub_544740(a1 int32) int32 {
	return sub_534750(a1)
}
func sub_544750(a1 int32) int32 {
	return sub_534780(a1)
}
func nox_xxx_mobActionFlee_544760(a1 int32) int8 {
	var (
		v1  int32
		v2  int32
		v3  int32
		v4  int32
		v5  int32
		v6  float64
		v7  float64
		v8  int32
		v9  int32
		v10 int32
		v11 int32
	)
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	if nox_xxx_monsterIsMoveing_534320(a1) != 0 {
		v3 = v1 + int32((*(*byte)(unsafe.Pointer(uintptr(v1 + 544)))+23)*24)
		v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 1196))))
		if v4 != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(v3 + 4))) = *(*uint32)(unsafe.Pointer(uintptr(v4 + 56)))
			*(*uint32)(unsafe.Pointer(uintptr(v3 + 8))) = *(*uint32)(unsafe.Pointer(uintptr(v4 + 60)))
			v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 1196))))
			v6 = float64(*(*float32)(unsafe.Pointer(uintptr(v5 + 56))) - *(*float32)(unsafe.Pointer(uintptr(a1 + 56))))
			v7 = float64(*(*float32)(unsafe.Pointer(uintptr(v5 + 60))) - *(*float32)(unsafe.Pointer(uintptr(a1 + 60))))
			if float64(*(*float32)(unsafe.Pointer(uintptr(v1 + 1356)))**(*float32)(unsafe.Pointer(uintptr(v1 + 1356)))) > v7*v7+v6*v6 && (gameFrame()-*(*uint32)(unsafe.Pointer(uintptr(v1 + 280)))) > gameFPS()>>1 {
				*(*uint32)(unsafe.Pointer(uintptr(v1 + 8))) = 0
			}
			if nox_xxx_monsterCanCast_534300((*nox_object_t)(unsafe.Pointer(uintptr(a1)))) != 0 && nox_xxx_testUnitBuffs_4FF350((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 29) == 0 && sub_534400(a1) == 0 && sub_534440(a1) == 0 {
				if nox_xxx_checkMobAction_50A0D0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 6) != 0 {
					v8 = nox_xxx_mobCastRelated_541050(a1)
				} else {
					v8 = nox_xxx_monsterBuffSelf_540B90(a1)
				}
				if v8 == 0 && nox_xxx_unitCanInteractWith_5370E0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), (*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 1196)))))), 0) != 0 {
					nox_xxx_monsterCastOffensive_540F20(a1, int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 1196)))))
				}
			}
		}
		v9 = int32(gameFrame())
		if *(*uint32)(unsafe.Pointer(uintptr(v1 + 8))) != 0 && gameFrame()-*(*uint32)(unsafe.Pointer(uintptr(v1 + 280))) > (gameFPS()*2) {
			*(*uint32)(unsafe.Pointer(uintptr(v1 + 8))) = 0
			v9 = int32(gameFrame())
		}
		if *(*uint32)(unsafe.Pointer(uintptr(v1 + 8))) != 0 || (uint32(v9)-*(*uint32)(unsafe.Pointer(uintptr(v1 + 280)))) <= 0xA || (func() bool {
			v10 = v3 + 4
			v11 = nox_xxx_generateRetreatPath_50CA00(v1+12, 32, a1, (*float32)(unsafe.Pointer(uintptr(v10))))
			*(*uint32)(unsafe.Pointer(uintptr(v1 + 8))) = uint32(v11)
			*(*uint32)(unsafe.Pointer(uintptr(v1 + 280))) = gameFrame()
			*(*uint32)(unsafe.Pointer(uintptr(v1 + 268))) = 0
			return v11 > 1
		}()) {
			if nox_xxx_creatureActuallyMove_50D3B0((*float32)(unsafe.Pointer(uintptr(a1)))) != 0 {
				*(*uint32)(unsafe.Pointer(uintptr(v1 + 8))) = 0
			}
			*((*uint8)(unsafe.Pointer(&v2))) = uint8(nox_xxx_monsterMoveAudio_534030(a1))
		} else {
			*(*uint32)(unsafe.Pointer(uintptr(v10))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 56)))
			v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 60))))
			*(*uint32)(unsafe.Pointer(uintptr(v10 + 4))) = uint32(v2)
		}
	} else {
		*((*uint8)(unsafe.Pointer(&v2))) = uint8(int8(nox_xxx_monsterPopAction_50A160((*nox_object_t)(unsafe.Pointer(uintptr(a1))))))
	}
	return int8(v2)
}
func nox_xxx_mobActionReturnToHome_544920(a1 int32) int32 {
	return sub_534750(a1)
}
func sub_544930(a1 int32) int32 {
	return sub_534780(a1)
}
func sub_544940(a1 int32) int32 {
	return sub_534780(a1)
}
func sub_544950(a1 int32) int8 {
	return nox_xxx_mobActionMoveTo_5443F0(a1)
}
func nox_xxx_mobActionHunt_5449D0(a1 int32) *int32 {
	var result *int32
	*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + 1304))) = 1062501089
	result = (*int32)(nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 10, internCStr(__FILE__), __LINE__))
	if result != nil {
		*(*int32)(unsafe.Add(unsafe.Pointer(result), 4*1)) = 0
		*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(result))), 12))) = 128
	}
	return result
}
func nox_xxx_mobSearchEdible_544A00(a1 *nox_object_t, a2 float32) int32 {
	*memmap.PtrUint32(0x5D4594, 2489452) = 0
	*memmap.PtrUint32(0x5D4594, 2489444) = 1259902592
	nox_xxx_unitsGetInCircle_517F90((*float2)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))+56))), a2, unsafe.Pointer(funAddr(nox_xxx_mobSearchEdible2_544A40)), unsafe.Pointer(a1))
	return int32(*memmap.PtrUint32(0x5D4594, 2489452))
}
func nox_xxx_mobSearchEdible2_544A40(a1 int32, a2 int32) {
	var (
		v2 int32
		v3 int32
		v4 float64
		v5 float64
		v6 float64
	)
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&0x10 != 0 {
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12))))
		if (v2&4) == 0 && ((v2&0x80) == 0 || nox_xxx_isNotPoisoned_5347F0(a2) != 0) {
			v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12))))
			if (v3&8) == 0 || v3&0x10 != 0 && nox_common_gameFlags_check_40A5C0(0x2000) && int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 12))))&0x10 != 0 {
				if nox_xxx_unitCanInteractWith_5370E0((*nox_object_t)(unsafe.Pointer(uintptr(a2))), (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0) != 0 {
					v4 = float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 56))) - *(*float32)(unsafe.Pointer(uintptr(a2 + 56))))
					v5 = float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 60))) - *(*float32)(unsafe.Pointer(uintptr(a2 + 60))))
					v6 = v5*v5 + v4*v4
					if v6 < float64(*mem_getFloatPtr(0x5D4594, 2489444)) {
						*mem_getFloatPtr(0x5D4594, 2489444) = float32(v6)
						*memmap.PtrUint32(0x5D4594, 2489452) = uint32(a1)
					}
				}
			}
		}
	}
}
func sub_544AE0(a1 int32, a2 float32) int32 {
	*memmap.PtrUint32(0x5D4594, 2489440) = 0
	*memmap.PtrUint32(0x5D4594, 2489448) = 1259902592
	nox_xxx_unitsGetInCircle_517F90((*float2)(unsafe.Pointer(uintptr(a1+56))), a2, unsafe.Pointer(funAddr(sub_544B20)), unsafe.Pointer(uintptr(a1)))
	return int32(*memmap.PtrUint32(0x5D4594, 2489440))
}
func sub_544B20(a1 int32, a2 int32) {
	var (
		v2 float64
		v3 float64
		v4 float64
	)
	if *(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))&0x1000000 != 0 && nox_xxx_playerClassCanUseItem_57B3D0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0) != 0 && nox_xxx_unitCanInteractWith_5370E0((*nox_object_t)(unsafe.Pointer(uintptr(a2))), (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0) != 0 {
		v2 = float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 56))) - *(*float32)(unsafe.Pointer(uintptr(a2 + 56))))
		v3 = float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 60))) - *(*float32)(unsafe.Pointer(uintptr(a2 + 60))))
		v4 = v3*v3 + v2*v2
		if v4 < float64(*mem_getFloatPtr(0x5D4594, 2489448)) {
			*mem_getFloatPtr(0x5D4594, 2489448) = float32(v4)
			*memmap.PtrUint32(0x5D4594, 2489440) = uint32(a1)
		}
	}
}
func nox_xxx_mobActionPickupObject_544B90(a1 int32) int8 {
	var (
		v1 int32
		v2 int32
		v3 int32
		v4 float64
		v5 float64
		v6 int32
	)
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	v2 = v1 + int32((*(*byte)(unsafe.Pointer(uintptr(v1 + 544)))+23)*24)
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + int32((*(*byte)(unsafe.Pointer(uintptr(v1 + 544)))+23)*24) + 4))))
	if v3 != 0 {
		v4 = float64(*(*float32)(unsafe.Pointer(uintptr(v3 + 56))) - *(*float32)(unsafe.Pointer(uintptr(a1 + 56))))
		v5 = float64(*(*float32)(unsafe.Pointer(uintptr(v3 + 60))) - *(*float32)(unsafe.Pointer(uintptr(a1 + 60))))
		if v5*v5+v4*v4 < 5625.0 {
			if nox_xxx_unitCanInteractWith_5370E0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), (*nox_object_t)(unsafe.Pointer(uintptr(v3))), 0) != 0 {
				nox_xxx_inventoryServPlace_4F36F0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), (*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 4)))))), 1, 1)
				v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 4))))
				if int32(*(*uint8)(unsafe.Pointer(uintptr(v6 + 12))))&0x10 != 0 {
					nox_xxx_useByNetCode_53F8E0(a1, v6)
				}
			}
		}
	}
	return int8(nox_xxx_monsterPopAction_50A160((*nox_object_t)(unsafe.Pointer(uintptr(a1)))))
}
func nox_xxx_mobGenericDeath_544C40(a1 int32) int32 {
	var (
		v1      int32
		v2      int32
		result  int32 = 0
		result2 func(int32) int32
	)
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	v2 = int32(uintptr(nox_xxx_monsterGetSoundSet_424300((*nox_object_t)(unsafe.Pointer(uintptr(a1))))))
	if v2 != 0 {
		nox_xxx_aud_501960(int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 60)))), (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
	}
	nox_xxx_scriptCallByEventBlock_502490(unsafe.Pointer(uintptr(v1+1264)), nil, unsafe.Pointer(uintptr(a1)), 7)
	result2 = *(*func(int32) int32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 484))) + 228)))
	if result2 != nil {
		result = result2(a1)
	}
	return result
}
func nox_xxx_zombieBurnDeleteCheck_544CA0(a1 *uint32) {
	var v1 int32
	v1 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*187)))
	if nox_xxx_unitIsZombie_534A40(int32(uintptr(unsafe.Pointer(a1)))) != 0 {
		if *(*uint32)(unsafe.Pointer(uintptr(v1 + 1440)))&0x80000 != 0 {
			nox_xxx_zombieBurnDelete_544CE0(a1)
		}
	}
}
func nox_xxx_zombieBurnDelete_544CE0(a1 *uint32) {
	var (
		v1 int32
		v2 int32
		v3 int32
	)
	v1 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*math.MaxInt8)))
	if v1 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 8))))&4 != 0 {
		v2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*3)))
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 748))))
		*((*uint8)(unsafe.Pointer(&v2))) = uint8(int8(v2 & math.MaxInt8))
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*3)) = uint32(v2)
		nox_xxx_netFxShield_0_4D9200(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 276))) + 2064)))), int32(uintptr(unsafe.Pointer(a1))))
		nox_xxx_netUnmarkMinimapObj_417300(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 276))) + 2064)))), (*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), 1)
	}
	nox_xxx_soloMonsterKillReward_4EE500_obj_health(int32(uintptr(unsafe.Pointer(a1))))
	nox_xxx_sMakeScorch_537AF0((*float32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*14)))), 1)
	nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))))
}
func sub_544D60(a1 int32) int8 {
	var result int8
	result = int8(a1)
	if int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + 483)))) != 0 {
		result = int8(nox_xxx_monsterPopAction_50A160((*nox_object_t)(unsafe.Pointer(uintptr(a1)))))
	}
	return result
}
func nox_xxx_mobActionDead1_544D80(a1 *uint32) int8 {
	var (
		v1 int32
		v2 func(*uint32)
		v3 int32
		v4 int32
		v5 int32
		v6 bool
		v8 float32
		v9 float32
	)
	v1 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*187)))
	if *(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*131)) == 14 && *(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*3))&0x10000 != 0 {
		nox_xxx_createReleasedSoul_544E60(int32(uintptr(unsafe.Pointer(a1))))
	}
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*21)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*20)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*23)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*22)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*25)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*24)) = 0
	v2 = *(*func(*uint32))(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 484))) + 232)))
	if v2 != nil {
		v2(a1)
	}
	if nox_xxx_unitIsZombie_534A40(int32(uintptr(unsafe.Pointer(a1)))) != 0 {
		v8 = float32(nox_xxx_gamedataGetFloatTable_419D70(internCStr("ZombieDeadDuration"), 0))
		v3 = nox_float2int(v8)
		v9 = float32(nox_xxx_gamedataGetFloatTable_419D70(internCStr("ZombieDeadDuration"), 1))
		v4 = nox_float2int(v9)
		*(*uint32)(unsafe.Pointer(uintptr(v1 + 492))) = uint32(nox_common_randomInt_415FA0(v3, v4))
		v5 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*4)))
		*((*uint8)(unsafe.Pointer(&v5))) = uint8(int8(v5 | 0x10))
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*4)) = uint32(v5)
	} else {
		v6 = nox_xxx_unitIsZombie_534A40(int32(uintptr(unsafe.Pointer(a1)))) == 0
		v5 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*4)))
		if v6 {
			*((*uint8)(unsafe.Pointer(&v5))) = uint8(int8(v5 | 0x18))
		} else {
			*((*uint8)(unsafe.Pointer(&v5))) = uint8(int8(v5 | 0x10))
		}
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*4)) = uint32(v5)
	}
	return int8(v5)
}
func nox_xxx_createReleasedSoul_544E60(a1 int32) {
	var (
		v1 int32
		v2 *uint32
		v3 *uint32
		v4 int16
	)
	v1 = int32(*memmap.PtrUint32(0x5D4594, 2489456))
	if *memmap.PtrUint32(0x5D4594, 2489456) == 0 {
		v1 = nox_xxx_getNameId_4E3AA0(internCStr("ReleasedSoul"))
		*memmap.PtrUint32(0x5D4594, 2489456) = uint32(v1)
	}
	v2 = (*uint32)(unsafe.Pointer(nox_xxx_newObjectWithTypeInd_4E3450(v1)))
	v3 = v2
	if v2 != nil {
		nox_xxx_createAt_4DAA50((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v2)))))), nil, *(*float32)(unsafe.Pointer(uintptr(a1 + 56))), *(*float32)(unsafe.Pointer(uintptr(a1 + 60))))
		v4 = int16(*(*uint16)(unsafe.Pointer(uintptr(a1 + 124))))
		*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v3))), unsafe.Sizeof(uint16(0))*63))) = uint16(v4)
		*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v3))), unsafe.Sizeof(uint16(0))*62))) = uint16(v4)
	}
}
func nox_xxx_mobActionDead2_544EC0(a1 int32) {
	var (
		v1 *uint32
		v2 int32
	)
	v1 = *(**uint32)(unsafe.Pointer(uintptr(a1 + 748)))
	if nox_xxx_unitIsZombie_534A40(a1) != 0 {
		v2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*360)))
		if uint32(v2)&0x80000 != 0 {
			nox_xxx_netSparkExplosionFx_5231B0((*float32)(unsafe.Pointer(uintptr(a1+56))), 100)
			nox_xxx_zombieBurnDelete_544CE0((*uint32)(unsafe.Pointer(uintptr(a1))))
		} else if (gameFrame()-*(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*137))) > *(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*123)) && (uint32(v2)&0x100000) == 0 {
			if *(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*299)) != 0 {
				nox_xxx_mobRaiseZombie_534AB0(a1)
			}
		}
	} else {
		nox_xxx_unitRemoveFromUpdatable_4DA920((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 744))) = 0
		sub_544F70(a1)
		if int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*121)) + 92))))&1 != 0 {
			nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
		}
	}
}
func sub_544F70(a1 int32) int32 {
	var (
		v1     int32
		result int32
		v3     int32
		v4     *uint32
		v5     uint8
		v6     int32
		v7     *uint32
	)
	v1 = 0
	result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(result + 296))))
	*(*uint8)(unsafe.Pointer(uintptr(result + 2094))) = 0
	if v3 > 0 {
		v4 = (*uint32)(unsafe.Pointer(uintptr(result + 300)))
		for {
			*v4 = 0
			v1++
			v4 = (*uint32)(unsafe.Add(unsafe.Pointer(v4), 4*1))
			if v1 >= *(*int32)(unsafe.Pointer(uintptr(result + 296))) {
				break
			}
		}
	}
	v5 = *(*uint8)(unsafe.Pointer(uintptr(result + 1129)))
	v6 = 0
	*(*uint32)(unsafe.Pointer(uintptr(result + 296))) = 0
	*(*uint32)(unsafe.Pointer(uintptr(result + 364))) = 0
	if int32(v5) > 0 {
		v7 = (*uint32)(unsafe.Pointer(uintptr(result + 1132)))
		for {
			*v7 = 0
			v6++
			v7 = (*uint32)(unsafe.Add(unsafe.Pointer(v7), 4*1))
			if v6 >= int32(*(*uint8)(unsafe.Pointer(uintptr(result + 1129)))) {
				break
			}
		}
	}
	*(*uint8)(unsafe.Pointer(uintptr(result + 1129))) = 0
	*(*uint32)(unsafe.Pointer(uintptr(result + 1196))) = 0
	return result
}