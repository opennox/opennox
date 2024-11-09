package main

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/noxworld-dev/opennox-lib/env"
)

type runConfig struct {
	rcon     bool
	rconPass string
	record   string
	replay   string
}

func (a *App) logFile(name string) (*os.File, error) {
	dir := filepath.Join(a.dir, "logs")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}
	return os.Create(filepath.Join(dir, name))
}

func (a *App) runGame(hd bool) {
	if err := a.conf.WriteIfChanged(); err != nil {
		err = fmt.Errorf("Cannot write config: %w", err)
		a.ErrorMsg(err)
	}

	exe := "opennox"
	name := "OpenNox"
	if hd {
		exe = "opennox-hd"
		name = "OpenNox HD"
	}
	if runtime.GOOS == "windows" {
		exe += ".exe"
	}
	var args, envs []string
	if env.IsDevMode() || a.run.rcon {
		host := "127.0.0.1" // for dev mode
		if a.run.rcon {
			host = "0.0.0.0" // enabled explicitly
		}
		args = append(args, "--rcon", host+":"+sshPort)
		if a.run.rconPass != "" {
			args = append(args, "--rcon-pass", a.run.rconPass)
		}
	}
	if a.run.replay != "" {
		envs = append(envs, "NOX_E2E="+a.run.replay)
	} else if a.run.record != "" {
		envs = append(envs, "NOX_E2E_RECORD="+a.run.record)
	}
	for {
		if a.runOpenNox(name, filepath.Join(a.dir, exe), args, envs) {
			break // closed properly
		}
		if !a.handleCrash() {
			break
		}
	}
}

func (a *App) runOpenNox(name string, path string, args, envs []string) bool {
	isDev := env.IsDevMode()
	cmd := exec.Command(path, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if len(envs) != 0 {
		slog.Info("running with envs", "env", envs)
		cmd.Env = append(os.Environ(), envs...)
	}
	if l, err := a.logFile("opennox.log"); err == nil {
		cmd.Stdout = io.MultiWriter(cmd.Stdout, l)
		cmd.Stderr = io.MultiWriter(cmd.Stderr, l)
		defer l.Close()
	}
	err := cmd.Start()
	if err != nil {
		slog.Error("cannot start opennox", "path", path, "err", err)
	}
	if errors.Is(err, fs.ErrNotExist) {
		a.ErrorMsg(errors.New("OpenNox not found!"))
		return true // do not restart
	} else if err != nil {
		a.ErrorMsg(fmt.Errorf("Cannot start %s: %w", name, err))
		return true // do not restart
	}
	slog.Info("opennox started", "path", path)
	if !isDev {
		a.w.Hide()
	}
	err = cmd.Wait()
	a.w.Show()
	if err != nil {
		slog.Error("opennox failed", "err", err)
		return false
	}
	slog.Info("opennox stopped")
	return true
}

func (a *App) handleCrash() bool {
	// TODO: Ask to send a bug report.
	return a.YesNoDialog("OpenNox crashed", "OpenNox crashed, would you like to restart?")
}
