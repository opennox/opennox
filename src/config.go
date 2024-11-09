package opennox

import (
	"github.com/noxworld-dev/opennox-lib/env"

	"github.com/noxworld-dev/opennox/v1/common/config"
)

var (
	configReadOnly = isDedicatedServer
	onConfigRead   []func()
)

func registerOnConfigRead(fnc func()) {
	onConfigRead = append(onConfigRead, fnc)
}

func maybeWriteConfig() {
	if configReadOnly || env.IsE2E() {
		return
	}
	_ = config.Global.WriteIfChanged()
}

func configStrPtr(key, env string, def string, ptr *string) {
	config.Global.SetDefault(key, def)
	if env != "" {
		config.Global.SetEnv(key, env)
	}
	*ptr = config.Global.GetString(key)
	registerOnConfigRead(func() {
		*ptr = config.Global.GetString(key)
	})
}

func configBoolPtr(key, env string, def bool, ptr *bool) {
	config.Global.SetDefault(key, def)
	if env != "" {
		config.Global.SetEnv(key, env)
	}
	*ptr = config.Global.GetBool(key)
	registerOnConfigRead(func() {
		*ptr = config.Global.GetBool(key)
	})
}

func configHiddenBoolPtr(key, env string, ptr *bool) {
	if env != "" {
		config.Global.SetEnv(key, env)
	}
	*ptr = config.Global.GetBool(key)
	registerOnConfigRead(func() {
		*ptr = config.Global.GetBool(key)
	})
}

func readConfig(path string) error {
	defer func() {
		for _, fnc := range onConfigRead {
			fnc()
		}
	}()
	return config.Global.Read(path)
}
