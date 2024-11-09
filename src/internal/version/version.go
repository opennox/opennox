package version

import (
	"log/slog"

	"github.com/noxworld-dev/opennox-lib/log"
)

var Log = log.New("version")

func LogVersion() {
	slog.Info("version", "vers", Version(), "commit", Commit())
}

const (
	DefVersion = "v1.9.0-dev"
)

const (
	devCommit = "<dev>"
)

var (
	version = DefVersion
	commit  = devCommit
)

func Version() string {
	return version
}

func Commit() string {
	return commit
}

func ClientVersion() string {
	if IsDev() {
		return version + " (" + commit + ")"
	}
	return version
}

func IsDev() bool {
	return commit == devCommit || semverIsDev(version)
}
