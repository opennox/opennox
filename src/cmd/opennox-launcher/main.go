package main

import (
	"errors"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"

	"github.com/noxworld-dev/opennox/v1/common/config"
)

const Title = "OpenNox Launcher"

func main() {
	NewApp().Run()
}

func NewApp() *App {
	a := &App{a: app.New(), conf: config.New("")}
	a.a.SetIcon(Logo())
	a.w = a.a.NewWindow(Title)
	a.w.Resize(fyne.NewSize(1024, 600))
	a.initUI()
	return a
}

type App struct {
	a    fyne.App
	w    fyne.Window
	dir  string
	conf *config.Config
	run  runConfig
	ui   launcherUI
}

func (a *App) setRoot(root fyne.CanvasObject) {
	a.w.SetContent(root)
}

func (a *App) InfoMsg(title, text string) {
	dia := dialog.NewInformation(title, text, a.w)
	done := make(chan struct{})
	dia.SetOnClosed(func() {
		close(done)
	})
	dia.Show()
	<-done
}

func (a *App) ErrorMsg(err error) {
	dia := dialog.NewError(err, a.w)
	done := make(chan struct{})
	dia.SetOnClosed(func() {
		close(done)
	})
	dia.Show()
	<-done
}

func (a *App) YesNoDialog(title, text string) bool {
	done := make(chan bool, 1)
	dialog.NewConfirm(title, text, func(ok bool) {
		done <- ok
	}, a.w).Show()
	return <-done
}

func (a *App) FileDialog(save bool, def string, ext []string) (string, error) {
	var (
		path string
		gerr error
		done = make(chan struct{})
	)
	var dia *dialog.FileDialog
	if save {
		dia = dialog.NewFileSave(func(rc fyne.URIWriteCloser, err error) {
			gerr = err
			if rc != nil {
				path = rc.URI().Path()
			} else if err == nil {
				gerr = errors.New("not selected")
			}
			close(done)
		}, a.w)
	} else {
		dia = dialog.NewFileOpen(func(rc fyne.URIReadCloser, err error) {
			gerr = err
			if rc != nil {
				path = rc.URI().Path()
			} else if err == nil {
				gerr = errors.New("not selected")
			}
			close(done)
		}, a.w)
	}
	if def != "" {
		dia.SetFileName(def)
	}
	if len(ext) != 0 {
		dia.SetFilter(&storage.ExtensionFileFilter{Extensions: ext})
	}
	dia.Show()
	<-done
	return path, gerr
}

func (a *App) FolderDialog() (string, error) {
	var (
		path string
		gerr error
		done = make(chan struct{})
	)
	dia := dialog.NewFolderOpen(func(rc fyne.ListableURI, err error) {
		gerr = err
		if rc != nil {
			path = rc.Path()
		} else if err == nil {
			gerr = errors.New("not selected")
		}
		close(done)
	}, a.w)
	dia.Show()
	<-done
	return path, gerr
}

func (a *App) Run() {
	a.dir, _ = os.Getwd()
	go a.main()
	a.w.ShowAndRun()
	_ = a.conf.WriteIfChanged()
}

func (a *App) main() {
	if !a.setDataDirAndLoadConfig() {
		a.w.Close()
		return
	}
	a.reloadConfigUI()
}
