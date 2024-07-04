package timer

import (
	"testing"
	"time"

	"github.com/noxworld-dev/opennox-lib/noxtest"
	"github.com/noxworld-dev/opennox-lib/platform"
	"github.com/stretchr/testify/require"
)

func platformTicks() uint64 {
	return uint64(platform.Ticks() / time.Millisecond)
}

func TestInterpolation(t *testing.T) {
	PlatformTicks = platformTicks
	p := noxtest.MockPlatform{}
	platform.Set(&p)

	timer := &Timer{}
	timer.Init(1)
	require.Equal(t, Timer{
		Flags:        3,
		Current:      0x10000,
		Target:       0x10000,
		DeltaPerTick: 1073741,
		MaxTickDelta: 1000,
		LastUpdated:  0,
	}, *timer)

	platform.Sleep(10100000)
	require.Equal(t, uint64(10), platformTicks())
	timer.Update()
	require.Equal(t, Timer{
		Flags:        3,
		Current:      0x10000,
		Target:       0x10000,
		DeltaPerTick: 1073741,
		MaxTickDelta: 1000,
		LastUpdated:  0,
	}, *timer)

	timer.SetInterp(0x2000)
	require.Equal(t, Timer{
		Flags:        2,
		Current:      0x10000,
		Target:       0x20000000,
		DeltaPerTick: 1073741,
		MaxTickDelta: 1000,
		LastUpdated:  10,
	}, *timer)
	timer.Update()
	require.Equal(t, Timer{
		Flags:        2,
		Current:      0x10000,
		Target:       0x20000000,
		DeltaPerTick: 1073741,
		MaxTickDelta: 1000,
		LastUpdated:  10,
	}, *timer)

	platform.Sleep(900000)
	require.Equal(t, uint64(11), platformTicks())

	timer.Update()
	require.Equal(t, Timer{
		Flags:        2,
		Current:      0x10000 + 1073741,
		Target:       0x20000000,
		DeltaPerTick: 1073741,
		MaxTickDelta: 1000,
		LastUpdated:  11,
	}, *timer)

	timer.Update()
	require.Equal(t, Timer{
		Flags:        2,
		Current:      0x10000 + 1073741,
		Target:       0x20000000,
		DeltaPerTick: 1073741,
		MaxTickDelta: 1000,
		LastUpdated:  11,
	}, *timer)

	platform.Sleep(2000000)
	require.Equal(t, uint64(13), platformTicks())

	timer.Update()
	require.Equal(t, Timer{
		Flags:        2,
		Current:      0x10000 + 1073741*3,
		Target:       0x20000000,
		DeltaPerTick: 1073741,
		MaxTickDelta: 1000,
		LastUpdated:  13,
	}, *timer)
}

func TestInterpolation2(t *testing.T) {
	PlatformTicks = platformTicks
	p := noxtest.MockPlatform{}
	platform.Set(&p)

	timer := &Timer{}
	timer.Init(0x20)
	require.Equal(t, Timer{
		Flags:        3,
		Current:      0x200000,
		Target:       0x200000,
		DeltaPerTick: 1073741,
		MaxTickDelta: 1000,
		LastUpdated:  0,
	}, *timer)

	platform.Sleep(10100000)
	require.Equal(t, uint64(10), platformTicks())
	timer.Update()
	require.Equal(t, Timer{
		Flags:        3,
		Current:      0x200000,
		Target:       0x200000,
		DeltaPerTick: 1073741,
		MaxTickDelta: 1000,
		LastUpdated:  0,
	}, *timer)

	timer.SetInterp(5)
	require.Equal(t, Timer{
		Flags:        2,
		Current:      0x200000,
		Target:       0x50000,
		DeltaPerTick: 1073741,
		MaxTickDelta: 1000,
		LastUpdated:  10,
	}, *timer)
	timer.Update()
	require.Equal(t, Timer{
		Flags:        2,
		Current:      0x200000,
		Target:       0x50000,
		DeltaPerTick: 1073741,
		MaxTickDelta: 1000,
		LastUpdated:  10,
	}, *timer)

	platform.Sleep(900000)
	require.Equal(t, uint64(11), platformTicks())

	timer.Update()
	require.Equal(t, Timer{
		Flags:        2,
		Current:      0x200000 - 1073741,
		Target:       0x50000,
		DeltaPerTick: 1073741,
		MaxTickDelta: 1000,
		LastUpdated:  11,
	}, *timer)

	timer.Update()
	require.Equal(t, Timer{
		Flags:        2,
		Current:      0x200000 - 1073741,
		Target:       0x50000,
		DeltaPerTick: 1073741,
		MaxTickDelta: 1000,
		LastUpdated:  11,
	}, *timer)

	platform.Sleep(2000000)
	require.Equal(t, uint64(13), platformTicks())

	timer.Update()
	require.Equal(t, Timer{
		Flags:        2,
		Current:      0x50000,
		Target:       0x50000,
		DeltaPerTick: 1073741,
		MaxTickDelta: 1000,
		LastUpdated:  13,
	}, *timer)
}

func TestInterpolation3(t *testing.T) {
	PlatformTicks = platformTicks
	p := noxtest.MockPlatform{}
	platform.Set(&p)

	timer := &Timer{}
	timer.Init(1)
	timer.SetParams(10, 10)
	require.Equal(t, Timer{
		Flags:        3,
		Current:      0x10000,
		Target:       0x10000,
		DeltaPerTick: 0x10000,
		MaxTickDelta: 10,
		LastUpdated:  0,
	}, *timer)

	platform.Sleep(10100000)
	require.Equal(t, uint64(10), platformTicks())
	timer.Update()
	require.Equal(t, Timer{
		Flags:        3,
		Current:      0x10000,
		Target:       0x10000,
		DeltaPerTick: 0x10000,
		MaxTickDelta: 10,
		LastUpdated:  0,
	}, *timer)

	timer.SetInterp(0x2000)
	require.Equal(t, Timer{
		Flags:        2,
		Current:      0x10000,
		Target:       0x20000000,
		DeltaPerTick: 0x10000,
		MaxTickDelta: 10,
		LastUpdated:  10,
	}, *timer)
	timer.Update()
	require.Equal(t, Timer{
		Flags:        2,
		Current:      0x10000,
		Target:       0x20000000,
		DeltaPerTick: 0x10000,
		MaxTickDelta: 10,
		LastUpdated:  10,
	}, *timer)

	platform.Sleep(20000000000)
	require.Equal(t, uint64(20010), platformTicks())

	timer.Update()
	require.Equal(t, Timer{
		Flags:        2,
		Current:      0x10000 + 10*0x10000,
		Target:       0x20000000,
		DeltaPerTick: 0x10000,
		MaxTickDelta: 10,
		LastUpdated:  20010,
	}, *timer)

	timer.Update()
	require.Equal(t, Timer{
		Flags:        2,
		Current:      0x10000 + 10*0x10000,
		Target:       0x20000000,
		DeltaPerTick: 0x10000,
		MaxTickDelta: 10,
		LastUpdated:  20010,
	}, *timer)

	platform.Sleep(2000000)
	require.Equal(t, uint64(20012), platformTicks())

	timer.Update()
	require.Equal(t, Timer{
		Flags:        2,
		Current:      0x10000 + 12*0x10000,
		Target:       0x20000000,
		DeltaPerTick: 0x10000,
		MaxTickDelta: 10,
		LastUpdated:  20012,
	}, *timer)
}

func TestRawSet(t *testing.T) {
	PlatformTicks = platformTicks
	p := noxtest.MockPlatform{}
	platform.Set(&p)

	timer := &Timer{}
	timer.Init(1)
	require.Equal(t, Timer{
		Flags:        3,
		Current:      0x10000,
		Target:       0x10000,
		DeltaPerTick: 1073741,
		MaxTickDelta: 1000,
		LastUpdated:  0,
	}, *timer)

	platform.Sleep(10100000)
	require.Equal(t, uint64(10), platformTicks())

	timer.Update()
	require.Equal(t, Timer{
		Flags:        3,
		Current:      0x10000,
		Target:       0x10000,
		DeltaPerTick: 1073741,
		MaxTickDelta: 1000,
		LastUpdated:  0,
	}, *timer)

	timer.SetRaw(0x2000)
	require.Equal(t, Timer{
		Flags:        3,
		Current:      0x10000,
		Target:       0x20000000,
		DeltaPerTick: 1073741,
		MaxTickDelta: 1000,
		LastUpdated:  0,
	}, *timer)

	timer.Update()
	require.Equal(t, Timer{
		Flags:        3,
		Current:      0x20000000,
		Target:       0x20000000,
		DeltaPerTick: 1073741,
		MaxTickDelta: 1000,
		LastUpdated:  0,
	}, *timer)
}

func TestSetParams(t *testing.T) {
	PlatformTicks = platformTicks
	p := noxtest.MockPlatform{}
	platform.Set(&p)

	timer := &Timer{}
	timer.Init(1)
	// So Current will increase 0x10000(65536) per 1000 tick
	// Due to integer division, it will increase 65 per tick
	timer.SetParams(1000, 1)
	require.Equal(t, Timer{
		Flags:        3,
		Current:      0x10000,
		Target:       0x10000,
		DeltaPerTick: 65,
		MaxTickDelta: 1000,
		LastUpdated:  0,
	}, *timer)
	timer.SetInterp(0x2000)
	require.Equal(t, Timer{
		Flags:        2,
		Current:      0x10000,
		Target:       0x20000000,
		DeltaPerTick: 65,
		MaxTickDelta: 1000,
		LastUpdated:  0,
	}, *timer)

	platform.Sleep(1100000)
	require.Equal(t, uint64(1), platformTicks())
	timer.Update()
	require.Equal(t, Timer{
		Flags:        2,
		Current:      0x10000 + 65,
		Target:       0x20000000,
		DeltaPerTick: 65,
		MaxTickDelta: 1000,
		LastUpdated:  1,
	}, *timer)

	platform.Sleep(9000000)
	require.Equal(t, uint64(10), platformTicks())
	timer.Update()
	require.Equal(t, Timer{
		Flags:        2,
		Current:      0x10000 + 65*10,
		Target:       0x20000000,
		DeltaPerTick: 65,
		MaxTickDelta: 1000,
		LastUpdated:  10,
	}, *timer)

	platform.Sleep(990100000)
	require.Equal(t, uint64(1000), platformTicks())
	timer.Update()
	require.Equal(t, Timer{
		Flags:        2,
		Current:      0x10000 + 65*1000,
		Target:       0x20000000,
		DeltaPerTick: 65,
		MaxTickDelta: 1000,
		LastUpdated:  1000,
	}, *timer)
}
