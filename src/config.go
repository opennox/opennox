package opennox

import (
	"github.com/adrg/xdg"
	"github.com/spf13/viper"
	"os"
	"path/filepath"

	"github.com/noxworld-dev/opennox-lib/env"
	"github.com/noxworld-dev/opennox-lib/log"
)

func init() {
	viper.SetConfigName(configName)
	viper.SetConfigType(configExt)

	// Snap
	if snapCommonDir := env.AppUserDir(); snapCommonDir != "" {
		viper.AddConfigPath(snapCommonDir)
	}

	// Current working directory unless it's in a location only for binaries
	wd, err1 := os.Getwd()
	if err1 != nil {
		wd = "."
	}
	if absd, _ := filepath.Abs(filepath.Dir(wd)); absd != "/usr/bin" && absd != "/bin" {
		viper.AddConfigPath(wd)
	}

	// User Config
	viper.AddConfigPath(filepath.Join(xdg.ConfigHome, "opennox"))

	// System Config
	for _, dir := range xdg.ConfigDirs {
		viper.AddConfigPath(filepath.Join(dir, "opennox"))
	}
}

const (
	configName        = "opennox"
	configExt         = "yml"
	configNoxDataPath = "game.data"
	configNoxSerial   = "game.serial"
)

var (
	configLog      = log.New("config")
	configPath     string
	configDirty    bool
	configReadOnly = isDedicatedServer
	onConfigRead   []func()
)

func registerOnConfigRead(fnc func()) {
	onConfigRead = append(onConfigRead, fnc)
}

func writeConfig() error {
	if configReadOnly || env.IsE2E() {
		return nil
	}
	configLog.Printf("writing to %q", configPath)
	if err := viper.WriteConfig(); err != nil {
		configLog.Printf("cannot save file: %v", err)
		return err
	}
	return nil
}

func maybeWriteConfig() {
	if !configDirty {
		return
	}
	_ = writeConfig()
}

func writeConfigLater() {
	configDirty = true
}

func configStrPtr(key, env string, def string, ptr *string) {
	viper.SetDefault(key, def)
	if env != "" {
		viper.BindEnv(key, env)
	}
	*ptr = viper.GetString(key)
	registerOnConfigRead(func() {
		*ptr = viper.GetString(key)
	})
}

func configBoolPtr(key, env string, def bool, ptr *bool) {
	viper.SetDefault(key, def)
	if env != "" {
		viper.BindEnv(key, env)
	}
	*ptr = viper.GetBool(key)
	registerOnConfigRead(func() {
		*ptr = viper.GetBool(key)
	})
}

func configHiddenBoolPtr(key, env string, ptr *bool) {
	if env != "" {
		viper.BindEnv(key, env)
	}
	*ptr = viper.GetBool(key)
	registerOnConfigRead(func() {
		*ptr = viper.GetBool(key)
	})
}

func readConfig(path string) error {
	defer func() {
		for _, fnc := range onConfigRead {
			fnc()
		}
	}()
	if path != "" {
		if abs, err := filepath.Abs(path); err == nil {
			path = abs
		}
		configPath = path
		viper.SetConfigFile(configPath)
	}
	err := viper.ReadInConfig()
	if err == nil {
		configPath = viper.ConfigFileUsed()
		configDirty = false
		configLog.Printf("using file: %q", configPath)
		return nil
	}
	_, ok := err.(viper.ConfigFileNotFoundError)
	if ok || os.IsNotExist(err) || filepath.Base(viper.ConfigFileUsed()) == configName {
		if snapCommonDir := env.AppUserDir(); snapCommonDir != "" {
			configPath = filepath.Join(snapCommonDir, configName+"."+configExt)
			viper.SetConfigFile(configPath)
		} else {
			configPath = filepath.Join(xdg.ConfigHome, "opennox", configName+"."+configExt)
			viper.SetConfigFile(configPath)
		}
		writeConfigLater()
		configLog.Println("file not found, using defaults")
		return nil
	}
	return err // error parsing config
}
