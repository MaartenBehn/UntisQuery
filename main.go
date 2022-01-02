package main

import (
	"fyne.io/fyne/v2/app"
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

	untis.Init()
	panel.Init(&window)

	event.Go(event.EventSetPanel, panel.PanelStart)

	go updateLoop()
	window.ShowAndRun()
}

var running bool
var fps float64

const maxFPS = 30

func updateLoop() {
	startTime := time.Now()
	var startDuration time.Duration
	wait := time.Duration(1000000000 / int(maxFPS))
	running = true
	for running {
		startDuration = time.Since(startTime)
		// All update Calls

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
