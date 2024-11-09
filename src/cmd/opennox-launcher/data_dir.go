package main

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/noxworld-dev/opennox-lib/datapath"

	"github.com/noxworld-dev/opennox/v1/common/config"
)

func (a *App) discoverDataDir() bool {
	path := datapath.FindData()
	if path == "" {
		return false
	}
	a.conf.Set(config.KeyNoxDataPath, path)
	datapath.SetData(path)
	a.normalizeDataPath()
	return true
}

func (a *App) setDataDirDialog() bool {
	path, err := a.FolderDialog()
	if err != nil {
		return false
	}
	a.conf.Set(config.KeyNoxDataPath, path)
	datapath.SetData(path)
	a.normalizeDataPath()
	return true
}

func (a *App) normalizeDataPath() {
	if path := a.conf.GetString(config.KeyNoxDataPath); filepath.IsAbs(path) && strings.HasPrefix(path, filepath.Dir(a.conf.Path())) {
		if rel, err := filepath.Rel(filepath.Dir(a.conf.Path()), path); err == nil {
			a.conf.Set(config.KeyNoxDataPath, rel)
		}
	}
}

func (a *App) setDataDirAndLoadConfig() bool {
	if !a.loadConfig() {
		return false
	}
	if a.conf.IsSet(config.KeyNoxDataPath) {
		path := a.conf.GetString(config.KeyNoxDataPath)
		datapath.SetData(path)
	} else if !a.discoverDataDir() {
		a.InfoMsg("Nox not found", "Cannot find Nox installation. Please select Nox folder in the next dialog.")
		if !a.setDataDirDialog() {
			a.ErrorMsg(errors.New("No Nox directory selected!"))
			return false
		}
	}
	a.normalizeDataPath()
	if err := a.conf.WriteIfChanged(); err != nil {
		err = fmt.Errorf("Cannot write config: %w", err)
		a.ErrorMsg(err)
		return false
	}
	return true
}
