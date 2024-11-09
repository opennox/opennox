package main

import "fmt"

func (a *App) loadConfig() bool {
	if err := a.conf.Read(""); err != nil {
		err = fmt.Errorf("Cannot read config: %w", err)
		a.ErrorMsg(err)
		return false
	}
	return true
}
