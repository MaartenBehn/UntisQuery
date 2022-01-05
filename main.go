package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/Stroby241/UntisQuery/event"
	"github.com/Stroby241/UntisQuery/panel"
	"github.com/Stroby241/UntisQuery/untis"
	"time"
)

func main() {
	a := app.New()
	window := a.NewWindow("Untis Querry")
	window.SetContent(widget.NewLabel("Hello World!"))

	event.On(event.EventHandleError, func(data interface{}) {
		if data != nil {
			dialog.ShowError(data.(error), window)
		}

		if r := recover(); r != nil {
			dialog.ShowError(r.(error), window)
		}
	})

	defer event.Go(event.EventHandleError, nil)

	untis.Init()
	panel.Init(&window)

	event.Go(event.EventSetPanel, panel.PanelStart)

	go updateLoop(&window)
	window.ShowAndRun()
}

var running bool
var fps float64

const maxFPS = 30

func updateLoop(window *fyne.Window) {
	defer event.Go(event.EventHandleError, nil)

	startTime := time.Now()
	var startDuration time.Duration
	wait := time.Duration(1000000000 / int(maxFPS))
	running = true
	for running {
		startDuration = time.Since(startTime)
		// All update Calls

		event.Go(event.EventUpdate, nil)

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
