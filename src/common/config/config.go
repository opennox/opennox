package config

import (
	"log/slog"
	"os"
	"path/filepath"
	"runtime"

	"github.com/noxworld-dev/opennox-lib/env"
	"github.com/spf13/viper"
)

var Global = New("")

var workDir string

func init() {
	var err error
	workDir, err = os.Getwd()
	if err != nil {
		workDir = "."
	}
}

const (
	FileName     = FileBaseName + FileExt
	FileExt      = "." + FileBaseExt
	FileBaseName = "opennox"
	FileBaseExt  = "yml"
)

const (
	KeyNoxDataPath        = "game.data"
	KeyNoxSerial          = "game.serial"
	KeyServerAPIToken     = "server.api_token"
	KeyServerAPICmds      = "server.control.allow_cmds"
	KeyServerAPIMapChange = "server.control.allow_map_change"
	KeyNetRegister        = "network.xwis.register"
	KeyNetLobbyAddr       = "network.lobby.address"
	KeyNetPortForward     = "network.port_forward"
	KeyNetDebug           = "network.debug"
	KeyNetXor             = "network.xor"
	KeyDefChatMap         = "server.maps.default_chat"
)

func New(wd string) *Config {
	if wd == "" {
		wd = workDir
	}
	v := viper.New()

	v.SetConfigName(FileBaseName)
	v.AddConfigPath(wd)
	var configPath string
	if sdir := env.AppUserDir(); sdir != "" {
		v.AddConfigPath(sdir)
		configPath = filepath.Join(sdir, FileName)
	} else {
		configPath = filepath.Join(wd, FileName)
	}
	v.AddConfigPath(filepath.Dir(os.Args[0]))
	if runtime.GOOS != "windows" {
		if dir := os.Getenv("XDG_CONFIG_HOME"); dir != "" {
			v.AddConfigPath(filepath.Join(dir, "opennox"))
		}
		if home, err := os.UserHomeDir(); err == nil {
			// Linux Snapcraft installation replaces HOME variable
			if rhome := os.Getenv("SNAP_REAL_HOME"); rhome != "" {
				home = rhome
			}
			v.AddConfigPath(filepath.Join(home, ".config/opennox"))
		}
	}

	return &Config{v: v, path: configPath, Log: slog.Default()}
}

type Config struct {
	v     *viper.Viper
	path  string
	dirty bool
	Log   *slog.Logger
}

func (c *Config) Path() string {
	return c.path
}

func (c *Config) Reset() {
	c.v = viper.New()
	c.dirty = false
}

func (c *Config) Read(path string) error {
	c.dirty = false
	if path != "" {
		if abs, err := filepath.Abs(path); err == nil {
			path = abs
		}
		c.v.SetConfigFile(path)
		c.path = path
	}
	err := c.v.ReadInConfig()
	if err == nil {
		c.path = c.v.ConfigFileUsed()
		c.Log.Info("using config file", "path", c.path)
		return nil
	}
	if _, ok := err.(viper.ConfigFileNotFoundError); ok || os.IsNotExist(err) {
		c.dirty = true
		c.Log.Warn("config file not found, using defaults")
		return nil
	}
	// There's a weird behavior in viper that it may try to read config file without extension.
	// It's a problem for us because we have "opennox" binary on Linux, so viper will read the binary itself and fail.
	// So we check this case and ignore it as if the config was not found.
	if filepath.Base(c.v.ConfigFileUsed()) == FileBaseName {
		c.dirty = true
		c.Log.Warn("config file not found, using defaults")
		return nil
	}
	return err // error parsing config
}

func (c *Config) Write(path string) error {
	if path == "" {
		path = c.path
	}
	c.Log.Info("writing config", "path", c.path)
	if err := c.v.WriteConfigAs(c.path); err != nil {
		c.Log.Error("cannot save config file", "path", c.path, "err", err)
		return err
	}
	c.path = path
	c.dirty = false
	return nil
}

func (c *Config) WriteIfChanged() error {
	if !c.dirty {
		return nil
	}
	return c.Write(c.path)
}

func (c *Config) Get(key string) any {
	return c.v.Get(key)
}

func (c *Config) IsSet(key string) bool {
	return c.v.IsSet(key)
}

func (c *Config) Set(key string, val any) {
	var old any
	if !c.dirty {
		old = c.v.Get(key)
	}
	c.v.Set(key, val)
	if !c.dirty {
		c.dirty = old != val
	}
}

func (c *Config) SetDefault(key string, val any) {
	c.v.SetDefault(key, val)
}

func (c *Config) SetEnv(key string, env string) {
	c.v.BindEnv(key, env)
}

func (c *Config) GetString(key string) string {
	return c.v.GetString(key)
}

func (c *Config) GetInt(key string) int {
	return c.v.GetInt(key)
}

func (c *Config) GetFloat(key string) float64 {
	return c.v.GetFloat64(key)
}

func (c *Config) GetBool(key string) bool {
	return c.v.GetBool(key)
}
