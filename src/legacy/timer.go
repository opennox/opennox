package legacy

/*
#include "defs.h"
#include "GAME2_2.h"
*/
import "C"

type Timer struct {
	Flags          uint32 // bit 1: raw or interp, bit 2: marked when updated
	Current        uint32 // current value
	Target         uint32 // target value
	Delta_per_tick uint32 // (interp only) the amount to add per tick
	Max_tick_delta uint64 // (interp only) the maximum delta per tick
	Last_updated   uint64 // the last tick the value was updated
}

func TimerInspect(self *C.timer) Timer {
	return Timer{
		Flags:          uint32(self.flags),
		Current:        uint32(self.current),
		Target:         uint32(self.target),
		Delta_per_tick: uint32(self.delta_per_tick),
		Max_tick_delta: uint64(self.max_tick_delta),
		Last_updated:   uint64(self.last_updated),
	}
}

// Create a new timer in heap for testing purpose
func TimerNew() *C.timer {
	ret := C.timer{}
	return &ret
}

func TimerInit(self *C.timer, target int) {
	C.sub_4862E0(self, C.int(target))
}

func TimerSetParams(self *C.timer, max_delta uint, interp_velocity int) {
	C.sub_486380(self, C.uint(max_delta), 0, C.int(interp_velocity))
}

func TimerSetRaw(self *C.timer, target int) {
	C.sub_486320(self, C.int(target))
}

func TimerSetInterp(self *C.timer, target int) {
	C.sub_486350(self, C.int(target))
}

// Based on self's mode(raw or interp), it will update the value to be current tick
func TimerUpdate(self *C.timer) {
	C.sub_4863B0(self)
}

type TimerGroup struct {
	// z[0]: usually internal value, ranged from 0 to 0x4000
	// z[1]: For UI display? ranged from 0 to 100
	// z[2]: Not quite sure. Ranged from 0 to 0x2000
	z [3]Timer
}

func TimerGroupInit(self *C.timerGroup) {
	C.sub_4864A0(self)
}

func TimerGroupUpdate(self *C.timerGroup) {
	C.sub_486520(self)
}

// Return true if any of self has been updated
func TimerGroupIsUpdated(self *C.timerGroup) bool {
	return C.sub_486550(self) != 0
}

// Looks like merge two timerGroups, but z[2] is bit unusual.
func TimerGroupMix(self *C.timerGroup, other *C.timerGroup) {
	C.sub_486570(self, other)
}

// Remove clear bit of all timers in self
func TimerGroupClearUpdated(self *C.timerGroup) {
	C.sub_486620(self)
}
