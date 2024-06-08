package legacy

/*
#include "defs.h"
#include "GAME2_2.h"
*/
import "C"
import "unsafe"

//export timer
type Timer struct {
	Flags        uint32 // bit 1: raw(set) or interp(unset), bit 2: set when updated
	Current      uint32 // current value
	Target       uint32 // target value
	DeltaPerTick uint32 // (interp only) the amount to add per tick
	MaxTickDelta uint64 // (interp only) the maximum delta per tick
	LastUpdated  uint64 // the last tick the value was updated
}

var _ = [1]struct{}{}[32-unsafe.Sizeof(Timer{})]

func TimerInit(self *Timer, target int) {
	C.sub_4862E0((*C.timer)(unsafe.Pointer(self)), C.int(target))
}

func TimerSetParams(self *Timer, max_delta uint, interp_velocity int) {
	C.sub_486380((*C.timer)(unsafe.Pointer(self)), C.uint(max_delta), 0, C.int(interp_velocity))
}

func TimerSetRaw(self *Timer, target int) {
	C.sub_486320((*C.timer)(unsafe.Pointer(self)), C.int(target))
}

func TimerSetInterp(self *Timer, target int) {
	C.sub_486350((*C.timer)(unsafe.Pointer(self)), C.int(target))
}

// Based on self's mode(raw or interp), it will update the value to be current tick
func TimerUpdate(self *Timer) {
	C.sub_4863B0((*C.timer)(unsafe.Pointer(self)))
}

//export timerGroup
type TimerGroup struct {
	// Timers[0]: usually internal value, ranged from 0 to 0x4000
	// Timers[1]: For UI display? ranged from 0 to 100
	// Timers[2]: Not quite sure. Ranged from 0 to 0x2000
	Timers [3]Timer
}

var _ = [1]struct{}{}[96-unsafe.Sizeof(TimerGroup{})]

func TimerGroupInit(self *TimerGroup) {
	C.sub_4864A0((*C.timerGroup)(unsafe.Pointer(self)))
}

func TimerGroupUpdate(self *TimerGroup) {
	C.sub_486520((*C.timerGroup)(unsafe.Pointer(self)))
}

// Return true if any of self has been updated
func TimerGroupIsUpdated(self *TimerGroup) bool {
	return C.sub_486550((*C.timerGroup)(unsafe.Pointer(self))) != 0
}

// Looks like merge two timerGroups, but z[2] is bit unusual.
func TimerGroupMix(self *TimerGroup, other *TimerGroup) {
	C.sub_486570((*C.timerGroup)(unsafe.Pointer(self)), (*C.timerGroup)(unsafe.Pointer(other)))
}

// Remove clear bit of all timers in self
func TimerGroupClearUpdated(self *TimerGroup) {
	C.sub_486620((*C.timerGroup)(unsafe.Pointer(self)))
}
