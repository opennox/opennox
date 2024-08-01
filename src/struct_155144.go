package opennox

import (
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"github.com/noxworld-dev/opennox/v1/legacy"
	"github.com/noxworld-dev/opennox/v1/legacy/common/ccall"
	"github.com/noxworld-dev/opennox/v1/legacy/timer"
)

type Struct88 struct {
	field_0  listItem
	field_12 [19]uint32
}

func (s *Struct88) getList() *listItem {
	return &s.field_0
}

type Struct264 struct {
	field_0  listItem
	field_12 [63]uint32
}

func (s *Struct264) getList() *listItem {
	return &s.field_0
}

type StructAt155144 struct {
	field_0  listHead[Struct88, *Struct88]
	field_12 listHead[Struct264, *Struct264]
	field_24 int32
	field_28 uint32 // unused
	field_32 timer.TimerGroup
}

func inst() *StructAt155144 {
	return (*StructAt155144)(legacy.Get_dword_587000_155144())
}

func sub_486F30() int {
	inst().field_0.Clear()
	inst().field_12.Clear()
	inst().field_24 = 0
	*memmap.PtrT[*timer.TimerGroup](0x5D4594, 1193340) = &inst().field_32
	inst().field_32.Init()
	dword_5d4594_1193336 = 1
	return 0
}

func sub_486EF0() {
	if dword_5d4594_1193336 != 0 {
		if inst().field_24 == 0 {
			listHead := &inst().field_12
			for v1 := listHead.next; v1 != &listHead.listItem; v1 = v1.next {
				if (*(*int32)(unsafe.Add(unsafe.Pointer(v1), 4*3)) & 2) == 0 {
					ccall.CallVoidPtr(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v1), 4*54)), unsafe.Pointer(v1))
				}
			}
		}
	}
}

func sub_487050(a1 *Struct88) {
	inst().field_0.Append(a1)
}

func sub_4870E0(a1 **Struct88) *Struct88 {
	result := inst().field_0.First()
	*a1 = result
	return result
}

func sub_487310(a1 *Struct264) int32 {
	inst().field_24 += 1
	inst().field_12.Append(a1)
	result := inst().field_24 - 1
	inst().field_24 = result
	if result < 0 {
		inst().field_24 = 0
	}
	return result
}

func sub_4875B0(a1 **Struct264) *Struct264 {
	result := inst().field_12.First()
	*a1 = result
	return result
}

// ----- (004875D0) --------------------------------------------------------
func sub_4875D0(a1 **Struct264) *Struct264 {
	if *a1 != nil {
		// FIXME: should be no cast
		*a1 = (*Struct264)(unsafe.Pointer((*a1).getList().Next()))
	}
	return *a1
}

func sub_4875F0() int32 {
	inst().field_24 += 1
	var v3 *Struct264
	v0 := sub_4875B0(&v3)
	if v0 != nil {
		for {
			v1 := sub_4875D0(&v3)
			legacy.Sub_487680(unsafe.Pointer(v0))
			v0 = v1
			if v1 == nil {
				break
			}
		}
	}
	result := inst().field_24 - 1
	inst().field_24 = result
	if result < 0 {
		inst().field_24 = 0
	}
	return result
}

func sub_4876A0(a1 *Struct264) {
	inst().field_24 += 1
	a1.getList().Remove()
	result := inst().field_24 - 1
	inst().field_24 = result
	if result < 0 {
		inst().field_24 = 0
	}
}
