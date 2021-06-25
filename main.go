package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"time"
)

const (
	panelIdLogin  = 1
	panelIdQuerry = 2
	panelIdMax    = 3
	panelIdStart  = 1
)

var window fyne.Window
var panels []*panel
var currentPanel int

func main() {
	a := app.New()

	a.Settings().SetTheme(theme.DarkTheme())

	window = a.NewWindow("Untis Teacher Querry")
	window.Resize(fyne.NewSize(400, 900))

	panels = make([]*panel, panelIdMax)
	panels[panelIdLogin] = newLoginPanel().panel
	panels[panelIdQuerry] = newQurreyPanel().panel

	currentPanel = panelIdStart

	go updateLoop()
	window.ShowAndRun()
	running = false

	if user.LoggedIn {
		user.Logout()
	}
}

var running bool
var fps float64

const maxFPS float64 = 10

func updateLoop() {
	startTime := time.Now()
	var startDuration time.Duration
	wait := time.Duration(1000000000 / int(maxFPS))

	running = true
	for running {
		startDuration = time.Since(startTime)
		// All update Calls

		checkLayout()
		checkContent()

		diff := time.Since(startTime) - startDuration
		if diff > 0 {
			fps = (wait.Seconds() / diff.Seconds()) * maxFPS
		} else {
			fps = 10000
		}
		if diff < wait {
			time.Sleep(wait - diff)
		}
	}
}

func checkLayout() {
	newLayout := getLayout(window.Content().Size())
	if newLayout != currentLayout {
		currentLayout = newLayout

		changeContent(panels[currentPanel].content[currentLayout])
	}
}

var lastCurrentPanel int

func checkContent() {
	if currentPanel != lastCurrentPanel {
		changeContent(panels[currentPanel].content[currentLayout])
	}
	lastCurrentPanel = currentPanel
}

func changeContent(content fyne.CanvasObject) {
	window.SetContent(content)
	window.Canvas().Content().Refresh()
	window.Show()
}
