package timer

import (
	"unsafe"
)

type Timer struct {
	Flags        uint32 // 0, bit 1: raw(set) or interp(unset), bit 2: set when updated
	Current      uint32 // 4, current value
	Target       uint32 // 8, target value
	DeltaPerTick uint32 // 12, (interp only) the amount to add per tick
	MaxTickDelta uint64 // 16, (interp only) the maximum delta per tick
	LastUpdated  uint64 // 24, the last tick the value was updated
}

var _ = [1]struct{}{}[32-unsafe.Sizeof(Timer{})]

func (timer *Timer) C() unsafe.Pointer {
	return unsafe.Pointer(timer)
}

type TimerGroup struct {
	// Timers[0]: usually internal value, ranged from 0 to 0x4000
	// Timers[1]: For UI display? ranged from 0 to 100
	// Timers[2]: Not quite sure. Ranged from 0 to 0x2000
	Timers [3]Timer
}

var _ = [1]struct{}{}[96-unsafe.Sizeof(TimerGroup{})]

func (timerGroup *TimerGroup) C() unsafe.Pointer {
	return unsafe.Pointer(timerGroup)
}

func (t *Timer) Init(target int32) bool {
	t.Flags = 0
	t.LastUpdated = uint64(PlatformTicks())
	t.SetParams(uint32(0x3E8), uint32(0x4000))
	t.SetRaw(uint32(target))
	return t.Update()
}

func (t *Timer) SetParams(max_delta uint32, interp_velocity uint32) {
	t.MaxTickDelta = uint64(max_delta)
	t.DeltaPerTick = (uint32(interp_velocity) << 16) / max_delta
}

func (t *Timer) SetRaw(target uint32) {
	t.Flags |= 1
	t.Target = target << 16
}

func (t *Timer) SetInterp(target uint32) {
	t.Flags &= 0xFFFFFFFE
	if t.Current == t.Target {
		ticks := uint64(PlatformTicks())
		t.LastUpdated = ticks
	}
	t.Target = target << 16
}

// Based on self's mode(raw or interp), it will update the value to be current tick
func (t *Timer) Update() bool {
	diff := int32(t.Target) - int32(t.Current)
	if diff == 0 {
		return false
	}
	if (t.Flags & 1) != 0 {
		t.Current = t.Target
		t.Flags = t.Flags | 2
		return true
	}

	currentTick := PlatformTicks()
	tickDelta := currentTick - t.LastUpdated
	t.LastUpdated = currentTick
	if tickDelta > t.MaxTickDelta {
		tickDelta = t.MaxTickDelta
	}
	valueDelta := uint32(tickDelta) * t.DeltaPerTick
	if diff >= 0 {
		if valueDelta < uint32(diff) {
			t.Current += valueDelta
		} else {
			t.Current += uint32(diff)
		}
	} else {
		if valueDelta <= uint32(-diff) {
			t.Current -= valueDelta
		} else {
			t.Current += uint32(diff)
		}
	}
	t.Flags |= 2
	return true
}

func (t *TimerGroup) Init() {
	t.Timers[0].Init(int32(0x4000))
	t.Timers[1].Init(int32(100))
	t.Timers[2].Init(int32(0x2000))
	t.Timers[0].SetParams(uint32(0x3E8), uint32(0x4000))
	t.Timers[1].SetParams(uint32(0x3E8), uint32(10))
	t.Timers[2].SetParams(uint32(0x3E8), uint32(0x4000))
	t.Timers[0].SetRaw(uint32(0x4000))
}

func (t *TimerGroup) Update() {
	t.Timers[0].Update()
	t.Timers[1].Update()
	t.Timers[2].Update()
}

// Return true if any of self has been updated
func (t *TimerGroup) IsUpdated() bool {
	return t.Timers[0].Flags&2 != 0 || t.Timers[1].Flags&2 != 0 || t.Timers[2].Flags&2 != 0
}

// Looks like merge two timerGroups, but Timers[2] is bit unusual.
func (t *TimerGroup) Mix(other *TimerGroup) {
	t.Timers[0].SetRaw(((other.Timers[0].Current) >> 16) * ((t.Timers[0].Current) >> 16) / 0x4000)
	t.Timers[0].Update()
	t.Timers[1].SetRaw(((other.Timers[1].Current) >> 16) * ((t.Timers[1].Current) >> 16) / 100)
	t.Timers[1].Update()
	if other.Timers[2].Current>>16 == 0x2000 {
		t.Timers[2].Update()
		return
	}
	v2 := int32((t.Timers[2].Current >> 16) + (other.Timers[2].Current >> 16) - 0x2000)
	if v2 >= 0 {
		if v2 > 0x4000 {
			v2 = 0x4000
		}
	} else {
		v2 = 0
	}
	t.Timers[2].SetRaw(uint32(v2))
	t.Timers[2].Update()
}

func TimerGroupClearUpdated(self *TimerGroup) {
	self.ClearUpdated()
}

// Remove clear bit of all timers in self
func (t *TimerGroup) ClearUpdated() {
	t.Timers[0].Flags &= 0xFFFFFFFD
	t.Timers[2].Flags &= 0xFFFFFFFD
	t.Timers[1].Flags &= 0xFFFFFFFD
}
