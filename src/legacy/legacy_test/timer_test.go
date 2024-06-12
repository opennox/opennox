package legacy_test

import (
	"testing"
	"time"

	"github.com/noxworld-dev/opennox-lib/noxtest"
	"github.com/noxworld-dev/opennox-lib/platform"
	"github.com/noxworld-dev/opennox/v1/legacy"
	"github.com/stretchr/testify/require"
)

func platformTicks() uint64 {
	return uint64(platform.Ticks() / time.Millisecond)
}

func TestInterpolation(t *testing.T) {
	legacy.PlatformTicks = platformTicks
	p := noxtest.MockPlatform{}
	platform.Set(&p)

	timer := &legacy.Timer{}
	legacy.TimerInit(timer, 1)
	require.Equal(t, legacy.Timer{
		Flags:        3,
		Current:      0x10000,
		Target:       0x10000,
		DeltaPerTick: 1073741,
		MaxTickDelta: 1000,
		LastUpdated:  0,
	}, *timer)

	platform.Sleep(10100000)
	require.Equal(t, uint64(10), platformTicks())
	legacy.TimerUpdate(timer)
	require.Equal(t, legacy.Timer{
		Flags:        3,
		Current:      0x10000,
		Target:       0x10000,
		DeltaPerTick: 1073741,
		MaxTickDelta: 1000,
		LastUpdated:  0,
	}, *timer)

	legacy.TimerSetInterp(timer, 0x2000)
	require.Equal(t, legacy.Timer{
		Flags:        2,
		Current:      0x10000,
		Target:       0x20000000,
		DeltaPerTick: 1073741,
		MaxTickDelta: 1000,
		LastUpdated:  10,
	}, *timer)
	legacy.TimerUpdate(timer)
	require.Equal(t, legacy.Timer{
		Flags:        2,
		Current:      0x10000,
		Target:       0x20000000,
		DeltaPerTick: 1073741,
		MaxTickDelta: 1000,
		LastUpdated:  10,
	}, *timer)

	platform.Sleep(900000)
	require.Equal(t, uint64(11), platformTicks())

	legacy.TimerUpdate(timer)
	require.Equal(t, legacy.Timer{
		Flags:        2,
		Current:      0x10000 + 1073741,
		Target:       0x20000000,
		DeltaPerTick: 1073741,
		MaxTickDelta: 1000,
		LastUpdated:  11,
	}, *timer)

	legacy.TimerUpdate(timer)
	require.Equal(t, legacy.Timer{
		Flags:        2,
		Current:      0x10000 + 1073741,
		Target:       0x20000000,
		DeltaPerTick: 1073741,
		MaxTickDelta: 1000,
		LastUpdated:  11,
	}, *timer)

	platform.Sleep(2000000)
	require.Equal(t, uint64(13), platformTicks())

	legacy.TimerUpdate(timer)
	require.Equal(t, legacy.Timer{
		Flags:        2,
		Current:      0x10000 + 1073741*3,
		Target:       0x20000000,
		DeltaPerTick: 1073741,
		MaxTickDelta: 1000,
		LastUpdated:  13,
	}, *timer)
}

func TestRawSet(t *testing.T) {
	legacy.PlatformTicks = platformTicks
	p := noxtest.MockPlatform{}
	platform.Set(&p)

	timer := &legacy.Timer{}
	legacy.TimerInit(timer, 1)
	require.Equal(t, legacy.Timer{
		Flags:        3,
		Current:      0x10000,
		Target:       0x10000,
		DeltaPerTick: 1073741,
		MaxTickDelta: 1000,
		LastUpdated:  0,
	}, *timer)

	platform.Sleep(10100000)
	require.Equal(t, uint64(10), platformTicks())

	legacy.TimerUpdate(timer)
	require.Equal(t, legacy.Timer{
		Flags:        3,
		Current:      0x10000,
		Target:       0x10000,
		DeltaPerTick: 1073741,
		MaxTickDelta: 1000,
		LastUpdated:  0,
	}, *timer)

	legacy.TimerSetRaw(timer, 0x2000)
	require.Equal(t, legacy.Timer{
		Flags:        3,
		Current:      0x10000,
		Target:       0x20000000,
		DeltaPerTick: 1073741,
		MaxTickDelta: 1000,
		LastUpdated:  0,
	}, *timer)

	legacy.TimerUpdate(timer)
	require.Equal(t, legacy.Timer{
		Flags:        3,
		Current:      0x20000000,
		Target:       0x20000000,
		DeltaPerTick: 1073741,
		MaxTickDelta: 1000,
		LastUpdated:  0,
	}, *timer)
}

func TestSetParams(t *testing.T) {
	legacy.PlatformTicks = platformTicks
	p := noxtest.MockPlatform{}
	platform.Set(&p)

	timer := &legacy.Timer{}
	legacy.TimerInit(timer, 1)
	// So Current will increase 0x10000(65536) per 1000 tick
	// Due to integer division, it will increase 65 per tick
	legacy.TimerSetParams(timer, 1000, 1)
	require.Equal(t, legacy.Timer{
		Flags:        3,
		Current:      0x10000,
		Target:       0x10000,
		DeltaPerTick: 65,
		MaxTickDelta: 1000,
		LastUpdated:  0,
	}, *timer)
	legacy.TimerSetInterp(timer, 0x2000)
	require.Equal(t, legacy.Timer{
		Flags:        2,
		Current:      0x10000,
		Target:       0x20000000,
		DeltaPerTick: 65,
		MaxTickDelta: 1000,
		LastUpdated:  0,
	}, *timer)

	platform.Sleep(1100000)
	require.Equal(t, uint64(1), platformTicks())
	legacy.TimerUpdate(timer)
	require.Equal(t, legacy.Timer{
		Flags:        2,
		Current:      0x10000 + 65,
		Target:       0x20000000,
		DeltaPerTick: 65,
		MaxTickDelta: 1000,
		LastUpdated:  1,
	}, *timer)

	platform.Sleep(9000000)
	require.Equal(t, uint64(10), platformTicks())
	legacy.TimerUpdate(timer)
	require.Equal(t, legacy.Timer{
		Flags:        2,
		Current:      0x10000 + 65*10,
		Target:       0x20000000,
		DeltaPerTick: 65,
		MaxTickDelta: 1000,
		LastUpdated:  10,
	}, *timer)

	platform.Sleep(990100000)
	require.Equal(t, uint64(1000), platformTicks())
	legacy.TimerUpdate(timer)
	require.Equal(t, legacy.Timer{
		Flags:        2,
		Current:      0x10000 + 65*1000,
		Target:       0x20000000,
		DeltaPerTick: 65,
		MaxTickDelta: 1000,
		LastUpdated:  1000,
	}, *timer)
}
