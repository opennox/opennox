package audiofx

import (
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/legacy"
	"github.com/noxworld-dev/opennox/v1/legacy/common/ccall"
	"github.com/noxworld-dev/opennox/v1/legacy/timer"
	"github.com/noxworld-dev/opennox/v1/lists"
)

type AudioFx struct {
	isInitialized uint32

	inst          *structAt155144
	timerGroupPtr **timer.TimerGroup
}

// @param dword_587000_155144 Allocated memory to be used to store AudioFx state
// @param ptr_5d4594_1193340 To write timerGroup pointer location during initalization
func NewAudioFx(dword_587000_155144 unsafe.Pointer, ptr_5d4594_1193340 **timer.TimerGroup) *AudioFx {
	return &AudioFx{
		inst:          (*structAt155144)(dword_587000_155144),
		isInitialized: 0,
		timerGroupPtr: ptr_5d4594_1193340,
	}
}

type Struct88 struct {
	lists.ListItem[Struct88]
	field_12 [19]uint32
}

var _ = [1]struct{}{}[88-unsafe.Sizeof(Struct88{})]

func (s *Struct88) getList() *lists.ListItem[Struct88] {
	return &s.ListItem
}

type Struct264 struct {
	lists.ListItem[Struct264]
	field_12 [63]uint32
}

var _ = [1]struct{}{}[264-unsafe.Sizeof(Struct264{})]

func (s *Struct264) getList() *lists.ListItem[Struct264] {
	return &s.ListItem
}

type structAt155144 struct {
	field_0  lists.ListHead[Struct88, *Struct88]
	field_12 lists.ListHead[Struct264, *Struct264]
	field_24 int32
	field_28 uint32 // unused
	field_32 timer.TimerGroup
}

func (a *AudioFx) Initialize() int {
	a.inst.field_0.Clear()
	a.inst.field_12.Clear()
	a.inst.field_24 = 0
	*a.timerGroupPtr = &a.inst.field_32
	a.inst.field_32.Init()
	a.isInitialized = 1
	return 0
}

func (a *AudioFx) Sub_486EF0() {
	if a.isInitialized == 0 {
		return
	}
	if a.inst.field_24 == 0 {
		listHead := &a.inst.field_12
		for v1 := listHead.Next(); v1 != &listHead.ListItem; v1 = v1.Next() {
			if (*(*int32)(unsafe.Add(unsafe.Pointer(v1), 4*3)) & 2) == 0 {
				ccall.CallVoidPtr(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v1), 4*54)), unsafe.Pointer(v1))
			}
		}
	}
}

func (a *AudioFx) Sub_487050(a1 *Struct88) {
	a.inst.field_0.Append(a1)
}

func (a *AudioFx) Sub_4870E0(a1 **Struct88) *Struct88 {
	result := a.inst.field_0.FirstSafe()
	*a1 = result
	return result
}

func (a *AudioFx) Sub_487310(a1 *Struct264) int32 {
	a.inst.field_24 += 1
	a.inst.field_12.Append(a1)
	result := a.inst.field_24 - 1
	a.inst.field_24 = result
	if result < 0 {
		a.inst.field_24 = 0
	}
	return result
}

func (a *AudioFx) Sub_4875B0(a1 **Struct264) *Struct264 {
	result := a.inst.field_12.FirstSafe()
	*a1 = result
	return result
}

// ----- (004875D0) --------------------------------------------------------
func sub_4875D0(a1 **Struct264) *Struct264 {
	if *a1 != nil {
		*a1 = (*a1).NextSafe().UnsafeGet()
	}
	return *a1
}

func (a *AudioFx) Sub_4875F0() int32 {
	a.inst.field_24 += 1
	var v3 *Struct264
	v0 := a.Sub_4875B0(&v3)
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
	result := a.inst.field_24 - 1
	a.inst.field_24 = result
	if result < 0 {
		a.inst.field_24 = 0
	}
	return result
}

func (a *AudioFx) Sub_4876A0(a1 *Struct264) {
	a.inst.field_24 += 1
	a1.getList().Remove()
	result := a.inst.field_24 - 1
	a.inst.field_24 = result
	if result < 0 {
		a.inst.field_24 = 0
	}
}

func (a *AudioFx) Cleanup() {
	if a.isInitialized != 0 {
		a.Sub_4875F0()
		legacy.Sub_4870A0()
		a.isInitialized = 0
	}
}
