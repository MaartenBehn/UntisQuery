package main

import (
	"UntisAPI"
	"time"
)

var user *UntisAPI.User
var timetable []map[int]UntisAPI.Period
var rooms map[int]UntisAPI.Room
var classes map[int]UntisAPI.Class

func main() {
	user = UntisAPI.NewUser("maarten8", "behn500", "TBZ Mitte Bremen", "https://tipo.webuntis.com")
	user.Login()

	date := UntisAPI.ToUntisDate(time.Now())
	loadTimetable(date, date)

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

	timetable = []map[int]UntisV2.Period{}
	counter := 0
	for _, room := range rooms {
		fmt.Printf("Loading timetable of room: %d of %d. \r", counter, len(rooms))

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
