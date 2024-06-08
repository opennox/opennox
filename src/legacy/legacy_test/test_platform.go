package legacy_test

// DO NOT MERGE: This file should be replaced by https://github.com/noxworld-dev/opennox-lib/pull/6

import (
	"time"

	"github.com/noxworld-dev/opennox-lib/platform"
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
