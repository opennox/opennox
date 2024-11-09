package main

import (
	_ "embed"

	"fyne.io/fyne/v2"
)

//go:embed res/opennox_logo.png
var logoPNG []byte

func Logo() fyne.Resource {
	return &fyne.StaticResource{StaticName: "logo.png", StaticContent: logoPNG}
}
