package audiofx

import (
	"testing"
	"time"
	"unsafe"

	"github.com/noxworld-dev/opennox-lib/platform"
	"github.com/noxworld-dev/opennox/v1/legacy"
	"github.com/noxworld-dev/opennox/v1/legacy/timer"
	"github.com/shoenig/test/must"
)

func platformTicks() uint64 {
	return uint64(platform.Ticks() / time.Millisecond)
}

func TestAudioFxInitialize(t *testing.T) {
	timer.PlatformTicks = platformTicks
	var timerGroupPtr *timer.TimerGroup = nil
	var memory = structAt155144{}
	audioFx := NewAudioFx(unsafe.Pointer(&memory), &timerGroupPtr)
	must.Eq(t, audioFx.Initialize(), 0)
	must.Eq(t, audioFx.isInitialized, 1)
	must.NotEq(t, audioFx.inst.field_0.First(), nil)
	must.Eq(t, audioFx.inst.field_0.FirstSafe(), nil)
	must.Eq(t, &audioFx.inst.field_32, timerGroupPtr)
	legacy.Sub_4870E0 = func(a1 unsafe.Pointer) unsafe.Pointer {
		return nil
	}
	audioFx.Cleanup()
	must.Eq(t, audioFx.isInitialized, 0)
}
