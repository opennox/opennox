package legacy

/*
#include "defs.h"
#include "GAME2_2.h"
*/
import "C"
import "unsafe"

type Timer struct {
	Flags        uint32 // bit 1: raw(set) or interp(unset), bit 2: set when updated
	Current      uint32 // current value
	Target       uint32 // target value
	DeltaPerTick uint32 // (interp only) the amount to add per tick
	MaxTickDelta uint64 // (interp only) the maximum delta per tick
	LastUpdated  uint64 // the last tick the value was updated
}

var _ = [1]struct{}{}[32-unsafe.Sizeof(Timer{})]

func (self *Timer) C() *C.timer {
	return (*C.timer)(unsafe.Pointer(self))
}

func TimerInit(self *Timer, target int) {
	C.sub_4862E0(self.C(), C.int(target))
}

func TimerSetParams(self *Timer, max_delta uint, interp_velocity int) {
	C.sub_486380(self.C(), C.uint(max_delta), 0, C.int(interp_velocity))
}

func TimerSetRaw(self *Timer, target int) {
	C.sub_486320(self.C(), C.int(target))
}

func TimerSetInterp(self *Timer, target int) {
	C.sub_486350(self.C(), C.int(target))
}

// Based on self's mode(raw or interp), it will update the value to be current tick
func TimerUpdate(self *Timer) {
	C.sub_4863B0(self.C())
}

type TimerGroup struct {
	// Timers[0]: usually internal value, ranged from 0 to 0x4000
	// Timers[1]: For UI display? ranged from 0 to 100
	// Timers[2]: Not quite sure. Ranged from 0 to 0x2000
	Timers [3]Timer
}

var _ = [1]struct{}{}[96-unsafe.Sizeof(TimerGroup{})]

func (self *TimerGroup) C() *C.timerGroup {
	return (*C.timerGroup)(unsafe.Pointer(self))
}

func TimerGroupInit(self *TimerGroup) {
	C.sub_4864A0(self.C())
}

func TimerGroupUpdate(self *TimerGroup) {
	C.sub_486520(self.C())
}

// Return true if any of self has been updated
func TimerGroupIsUpdated(self *TimerGroup) bool {
	return C.sub_486550(self.C()) != 0
}

// Looks like merge two timerGroups, but Timers[2] is bit unusual.
func TimerGroupMix(self *TimerGroup, other *TimerGroup) {
	C.sub_486570(self.C(), other.C())
}

// Remove clear bit of all timers in self
func TimerGroupClearUpdated(self *TimerGroup) {
	C.sub_486620(self.C())
}
