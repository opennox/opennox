package opennox

import (
	"github.com/noxworld-dev/opennox/v1/legacy/timer"
)

func init() {
	timer.PlatformTicks = platformTicks
}
