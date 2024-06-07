package legacy_test

import (
	"testing"
	"time"

	"github.com/noxworld-dev/opennox-lib/platform"
	"github.com/noxworld-dev/opennox/v1/legacy"
	"github.com/stretchr/testify/require"
)

type TestPlatform struct {
	t time.Duration
}

func (p *TestPlatform) Ticks() time.Duration {
	return p.t
}

func (p *TestPlatform) Sleep(dt time.Duration) {
	p.t = dt
}

func (p *TestPlatform) TimeSeed() int64 {
	return 0
}

func (p *TestPlatform) RandInt() int {
	return 0
}

func (p *TestPlatform) RandSeed(seed int64) {
}

func (p *TestPlatform) RandSeedTime() {
}

func platformTicks() uint64 {
	return uint64(platform.Ticks() / time.Millisecond)
}

func TestInterpolation(t *testing.T) {
	legacy.PlatformTicks = platformTicks
	p := TestPlatform{}
	platform.Set(&p)

	timer := legacy.TimerNew()
	legacy.TimerInit(timer, 1)
	require.Equal(t, legacy.Timer{
		Flags:          3,
		Current:        65536,
		Target:         65536,
		Delta_per_tick: 1073741,
		Max_tick_delta: 1000,
		Last_updated:   0,
	}, legacy.TimerInspect(timer))

	platform.Sleep(10100000)
	require.Equal(t, uint64(10), platformTicks())
	legacy.TimerUpdate(timer)
	require.Equal(t, legacy.Timer{
		Flags:          3,
		Current:        65536,
		Target:         65536,
		Delta_per_tick: 1073741,
		Max_tick_delta: 1000,
		Last_updated:   0,
	}, legacy.TimerInspect(timer))

	legacy.TimerSetInterp(timer, 0x2000)
	require.Equal(t, legacy.Timer{
		Flags:          2,
		Current:        65536,
		Target:         536870912,
		Delta_per_tick: 1073741,
		Max_tick_delta: 1000,
		Last_updated:   10,
	}, legacy.TimerInspect(timer))
	legacy.TimerUpdate(timer)
	require.Equal(t, legacy.Timer{
		Flags:          2,
		Current:        65536,
		Target:         536870912,
		Delta_per_tick: 1073741,
		Max_tick_delta: 1000,
		Last_updated:   10,
	}, legacy.TimerInspect(timer))

	platform.Sleep(11000000)
	require.Equal(t, uint64(11), platformTicks())

	require.Equal(t, legacy.Timer{
		Flags:          2,
		Current:        65536,
		Target:         536870912,
		Delta_per_tick: 1073741,
		Max_tick_delta: 1000,
		Last_updated:   10,
	}, legacy.TimerInspect(timer))
	legacy.TimerUpdate(timer)
	require.Equal(t, legacy.Timer{
		Flags:          2,
		Current:        1139277,
		Target:         536870912,
		Delta_per_tick: 1073741,
		Max_tick_delta: 1000,
		Last_updated:   11,
	}, legacy.TimerInspect(timer))
	legacy.TimerUpdate(timer)
	require.Equal(t, legacy.Timer{
		Flags:          2,
		Current:        1139277,
		Target:         536870912,
		Delta_per_tick: 1073741,
		Max_tick_delta: 1000,
		Last_updated:   11,
	}, legacy.TimerInspect(timer))

	platform.Sleep(13000000)
	require.Equal(t, uint64(13), platformTicks())

	legacy.TimerUpdate(timer)
	require.Equal(t, legacy.Timer{
		Flags:          2,
		Current:        3286759,
		Target:         536870912,
		Delta_per_tick: 1073741,
		Max_tick_delta: 1000,
		Last_updated:   13,
	}, legacy.TimerInspect(timer))
}

func TestRawSet(t *testing.T) {
	legacy.PlatformTicks = platformTicks
	p := TestPlatform{}
	platform.Set(&p)

	timer := legacy.TimerNew()
	legacy.TimerInit(timer, 1)
	require.Equal(t, legacy.Timer{
		Flags:          3,
		Current:        65536,
		Target:         65536,
		Delta_per_tick: 1073741,
		Max_tick_delta: 1000,
		Last_updated:   0,
	}, legacy.TimerInspect(timer))

	platform.Sleep(10100000)
	require.Equal(t, uint64(10), platformTicks())
	legacy.TimerUpdate(timer)
	require.Equal(t, legacy.Timer{
		Flags:          3,
		Current:        65536,
		Target:         65536,
		Delta_per_tick: 1073741,
		Max_tick_delta: 1000,
		Last_updated:   0,
	}, legacy.TimerInspect(timer))

	legacy.TimerSetRaw(timer, 0x2000)
	require.Equal(t, legacy.Timer{
		Flags:          3,
		Current:        0x10000,
		Target:         0x20000000,
		Delta_per_tick: 1073741,
		Max_tick_delta: 1000,
		Last_updated:   0,
	}, legacy.TimerInspect(timer))

	legacy.TimerUpdate(timer)
	require.Equal(t, legacy.Timer{
		Flags:          3,
		Current:        0x20000000,
		Target:         0x20000000,
		Delta_per_tick: 1073741,
		Max_tick_delta: 1000,
		Last_updated:   0,
	}, legacy.TimerInspect(timer))
}

func TestSetParams(t *testing.T) {
	legacy.PlatformTicks = platformTicks
	p := TestPlatform{}
	platform.Set(&p)

	timer := legacy.TimerNew()
	legacy.TimerInit(timer, 1)
	// So Current will increase 0x10000(65536) per 1000 tick
	// Due to integer division, it will increase 65 per tick
	legacy.TimerSetParams(timer, 1000, 1)
	require.Equal(t, legacy.Timer{
		Flags:          3,
		Current:        65536,
		Target:         65536,
		Delta_per_tick: 65,
		Max_tick_delta: 1000,
		Last_updated:   0,
	}, legacy.TimerInspect(timer))
	legacy.TimerSetInterp(timer, 0x2000)
	require.Equal(t, legacy.Timer{
		Flags:          2,
		Current:        0x10000,
		Target:         0x20000000,
		Delta_per_tick: 65,
		Max_tick_delta: 1000,
		Last_updated:   0,
	}, legacy.TimerInspect(timer))

	platform.Sleep(1100000)
	require.Equal(t, uint64(1), platformTicks())
	legacy.TimerUpdate(timer)
	require.Equal(t, legacy.Timer{
		Flags:          2,
		Current:        0x10000 + 65,
		Target:         0x20000000,
		Delta_per_tick: 65,
		Max_tick_delta: 1000,
		Last_updated:   1,
	}, legacy.TimerInspect(timer))

	platform.Sleep(10100000)
	require.Equal(t, uint64(10), platformTicks())
	legacy.TimerUpdate(timer)
	require.Equal(t, legacy.Timer{
		Flags:          2,
		Current:        0x10000 + 65*10,
		Target:         0x20000000,
		Delta_per_tick: 65,
		Max_tick_delta: 1000,
		Last_updated:   10,
	}, legacy.TimerInspect(timer))

	platform.Sleep(1000100000)
	require.Equal(t, uint64(1000), platformTicks())
	legacy.TimerUpdate(timer)
	require.Equal(t, legacy.Timer{
		Flags:          2,
		Current:        0x10000 + 65*1000,
		Target:         0x20000000,
		Delta_per_tick: 65,
		Max_tick_delta: 1000,
		Last_updated:   1000,
	}, legacy.TimerInspect(timer))
}
