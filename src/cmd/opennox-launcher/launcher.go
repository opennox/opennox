package main

import (
	"errors"
	"fmt"
	_ "image/png"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/noxworld-dev/opennox-lib/env"

	"github.com/noxworld-dev/opennox/v1/common/config"
	"github.com/noxworld-dev/opennox/v1/common/serial"
	"github.com/noxworld-dev/opennox/v1/internal/version"
)

type launcherUI struct {
	root       fyne.CanvasObject
	confReload []func()
}

func (a *App) newConfigEntry(key string) *widget.Entry {
	e := widget.NewEntry()
	e.OnChanged = func(v string) {
		a.conf.Set(key, v)
	}
	a.ui.confReload = append(a.ui.confReload, func() {
		e.Text = a.conf.GetString(key)
		e.Refresh()
	})
	return e
}

func (a *App) newConfigPass(key string) *widget.Entry {
	e := a.newConfigEntry(key)
	e.Password = true
	return e
}

func (a *App) newConfigToggle(key string) *widget.Check {
	e := widget.NewCheck("", func(v bool) {
		a.conf.Set(key, v)
	})
	a.ui.confReload = append(a.ui.confReload, func() {
		e.Checked = a.conf.GetBool(key)
		e.Refresh()
	})
	return e
}

func (a *App) newFuncEntry(fnc func(v string)) *widget.Entry {
	e := widget.NewEntry()
	e.OnChanged = fnc
	return e
}

func (a *App) newPtrEntry(p *string) *widget.Entry {
	e := widget.NewEntry()
	e.OnChanged = func(v string) {
		*p = v
	}
	return e
}

func (a *App) newPtrPass(p *string) *widget.Entry {
	e := a.newPtrEntry(p)
	e.Password = true
	return e
}

func (a *App) newPtrToggle(p *bool) *widget.Check {
	return widget.NewCheck("", func(v bool) {
		*p = v
	})
}

func (a *App) reloadConfigUI() {
	for _, fnc := range a.ui.confReload {
		fnc()
	}
}

func (a *App) initUI() {
	logo := canvas.NewImageFromResource(Logo())
	logo.FillMode = canvas.ImageFillContain
	const height = 64
	logo.SetMinSize(fyne.NewSize(height, height))

	headerCur := canvas.NewText("OpenNox "+version.Version(), theme.ForegroundColor())
	headerCur.TextSize = theme.TextHeadingSize()

	headerNew := canvas.NewText("", theme.ForegroundColor())
	headerCur.TextSize = theme.TextSubHeadingSize()
	if !version.IsDev() {
		go func() {
			if !version.IsLatest() {
				latest := version.Latest()
				headerNew.Color = theme.SuccessColor()
				headerNew.Text = fmt.Sprintf("New version is available: %s", latest)
				headerNew.Refresh()
			}
		}()
	} else {
		headerNew.Color = theme.WarningColor()
		headerNew.Text = "Development version"
		headerNew.Refresh()
	}
	runLegacy := widget.NewButtonWithIcon("Play", theme.MediaPlayIcon(), func() {
		go a.runGame(false)
	})
	runHD := widget.NewButtonWithIcon("Play HD", theme.MediaPlayIcon(), func() {
		go a.runGame(true)
	})
	tabs := []*container.TabItem{
		{
			Text:    "General",
			Icon:    theme.DocumentIcon(),
			Content: a.generalUI(),
		},
		{
			Text:    "Online",
			Icon:    theme.LoginIcon(),
			Content: a.onlineUI(),
		},
		{
			Text:    "Server",
			Icon:    theme.ComputerIcon(),
			Content: a.serverUI(),
		},
		{
			Text:    "Tools",
			Icon:    theme.SettingsIcon(),
			Content: a.toolsUI(),
		},
	}
	if env.IsDevMode() {
		tabs = append(tabs, []*container.TabItem{
			{
				Text:    "Dev",
				Icon:    theme.SettingsIcon(),
				Content: a.devUI(),
			},
		}...)
	}
	tabs = append(tabs, []*container.TabItem{
		{
			Text:    "Help",
			Icon:    theme.HelpIcon(),
			Content: a.helpUI(),
		},
	}...)
	appTabs := container.NewAppTabs(tabs...)
	a.ui.root = container.NewBorder(
		container.NewBorder(
			nil, nil,
			container.NewHBox(logo, container.NewPadded(container.NewVBox(headerCur, headerNew))),
			container.NewVBox(container.NewHBox(runLegacy, runHD)),
		),
		nil, nil, nil,
		appTabs,
	)
	a.setRoot(a.ui.root)
	a.reloadConfigUI()
}

func (a *App) generalUI() fyne.CanvasObject {
	return container.NewPadded(widget.NewForm([]*widget.FormItem{
		{
			Text: "Nox data",
			Widget: container.NewBorder(
				nil, nil, nil,
				container.NewHBox(
					widget.NewButtonWithIcon("", theme.FolderIcon(), func() {
						go func() {
							a.setDataDirDialog()
							a.reloadConfigUI()
						}()
					}),
					widget.NewButtonWithIcon("", theme.ViewRefreshIcon(), func() {
						go func() {
							if !a.discoverDataDir() {
								a.ErrorMsg(errors.New("Cannot find Nox installation!"))
							}
							a.reloadConfigUI()
						}()
					}),
				),
				a.newConfigEntry(config.KeyNoxDataPath),
			),
			HintText: "Nox installation directory",
		},
		{
			Text: "Serial",
			Widget: container.NewBorder(
				nil, nil, nil,
				widget.NewButtonWithIcon("", theme.ViewRefreshIcon(), func() {
					go func() {
						a.conf.Set(config.KeyNoxSerial, serial.Generate())
						a.reloadConfigUI()
					}()
				}),
				a.newConfigPass(config.KeyNoxSerial),
			),
			HintText: "Nox serial number (used in multiplayer)",
		},
	}...))
}

func (a *App) onlineUI() fyne.CanvasObject {
	items := []*widget.FormItem{
		{
			Text:     "Use lobby",
			Widget:   a.newConfigToggle(config.KeyNetRegister),
			HintText: "Register multiplayer games in global lobby",
		},
		{
			Text:     "Port forward",
			Widget:   a.newConfigToggle(config.KeyNetPortForward),
			HintText: "Forward port automatically (requires UPnP on the router)",
		},
		{
			Text:     "Chat map",
			Widget:   a.newConfigEntry(config.KeyDefChatMap),
			HintText: "Default chat map when hosting a game",
		},
	}
	if env.IsDevMode() {
		items = append(items, []*widget.FormItem{
			{
				Text:     "Lobby URL",
				Widget:   a.newConfigEntry(config.KeyNetLobbyAddr),
				HintText: "Lobby server API URL",
			},
			{
				Text:     "Log packets",
				Widget:   a.newConfigToggle(config.KeyNetDebug),
				HintText: "Log all UDP packets",
			},
			{
				Text:     "XOR encoding",
				Widget:   a.newConfigToggle(config.KeyNetXor),
				HintText: "XOR-encode network packets (WARNING: do not disable unless you know what you are doing!)",
			},
		}...)
	}
	return container.NewPadded(widget.NewForm(items...))
}

const sshPort = "18522"

func (a *App) serverUI() fyne.CanvasObject {
	items := []*widget.FormItem{
		{
			Text:     "Enable SSH",
			Widget:   a.newPtrToggle(&a.run.rcon),
			HintText: "Enable SSH server on port " + sshPort,
		},
		{
			Text:     "SSH password",
			Widget:   a.newPtrPass(&a.run.rconPass),
			HintText: "Password for SSH server",
		},
		{
			Text:     "API token",
			Widget:   a.newConfigPass(config.KeyServerAPIToken),
			HintText: "Server HTTP API token for remote control",
		},
		{
			Text:     "API commands",
			Widget:   a.newConfigToggle(config.KeyServerAPICmds),
			HintText: "Allow sending console commands via server HTTP API",
		},
		{
			Text:     "API map change",
			Widget:   a.newConfigToggle(config.KeyServerAPIMapChange),
			HintText: "Allow changing map via server HTTP API",
		},
	}
	return container.NewPadded(widget.NewForm(items...))
}

func (a *App) toolsUI() fyne.CanvasObject {
	var updateRecording func()
	recordFile := a.newFuncEntry(func(v string) {
		a.run.record = v
		a.run.replay = ""
		updateRecording()
	})
	replayFile := a.newFuncEntry(func(v string) {
		a.run.replay = v
		a.run.record = ""
		updateRecording()
	})
	updateRecording = func() {
		recordFile.Text = a.run.record
		recordFile.Refresh()
		replayFile.Text = a.run.replay
		replayFile.Refresh()
	}
	return container.NewAppTabs([]*container.TabItem{
		{
			Text: "Record",
			Content: widget.NewForm([]*widget.FormItem{
				{
					Text: "Record to file",
					Widget: container.NewBorder(
						nil, nil, nil,
						widget.NewButtonWithIcon("", theme.FileIcon(), func() {
							go func() {
								if path, err := a.FileDialog(true, "recording.yaml", []string{".yaml"}); err == nil {
									a.run.record = path
									a.run.replay = ""
									updateRecording()
								}
							}()
						}),
						recordFile,
					),
					HintText: "Record game input to a file so that it can be replayed later",
				},
				{
					Text: "Replay file",
					Widget: container.NewBorder(
						nil, nil, nil,
						widget.NewButtonWithIcon("", theme.FileIcon(), func() {
							go func() {
								if path, err := a.FileDialog(false, a.run.record, []string{".yaml"}); err == nil {
									a.run.replay = path
									a.run.record = ""
									updateRecording()
								}
							}()
						}),
						replayFile,
					),
					HintText: "Replay existing game input recording",
				},
			}...),
		},
	}...)
}

func (a *App) devUI() fyne.CanvasObject {
	return container.NewAppTabs([]*container.TabItem{
		{
			Text: "HTTP API",
			Content: widget.NewForm([]*widget.FormItem{
				{
					Text: "Dev APIs",
					Widget: container.NewHBox(
						widget.NewHyperlink("Profiler", parseURL("http://localhost:6060/debug/pprof")),
						widget.NewHyperlink("Players", parseURL("http://localhost:6060/debug/nox/players")),
						widget.NewHyperlink("Objects", parseURL("http://localhost:6060/debug/nox/objects")),
					),
				},
			}...),
		},
	}...)
}

func (a *App) helpUI() fyne.CanvasObject {
	return container.NewPadded(container.NewVBox(
		widget.NewHyperlink("Documentation", parseURL("https://noxworld-dev.github.io/opennox-docs/")),
		widget.NewHyperlink("Report an issue", parseURL("https://github.com/noxworld-dev/opennox/issues")),
		widget.NewHyperlink("Discord", parseURL("https://discord.gg/HgDUeXhAyW")),
	))
}

func parseURL(s string) *url.URL {
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	return u
}
