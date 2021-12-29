package main

/*
import (
	"UntisAPI"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
	"time"
)

type querryPanel struct {
	*panel

	fromDay   *widget.Entry
	fromMonth *widget.Entry
	fromYear  *widget.Entry

	tillDay   *widget.Entry
	tillMonth *widget.Entry
	tillYear  *widget.Entry

	firstName *widget.Entry
	lastName  *widget.Entry
	result    *widget.Label
}

func newQurreyPanel() *querryPanel {
	p := &querryPanel{panel: newPanel()}

	p.fromDay = widget.NewEntry()
	p.fromMonth = widget.NewEntry()
	p.fromYear = widget.NewEntry()

	p.tillDay = widget.NewEntry()
	p.tillMonth = widget.NewEntry()
	p.tillYear = widget.NewEntry()

	date := time.Now()
	p.fromDay.SetText(strconv.Itoa(date.Day()))
	p.fromMonth.SetText(strconv.Itoa(int(date.Month())))
	p.fromYear.SetText(strconv.Itoa(date.Year()))

	p.tillDay.SetText(strconv.Itoa(date.Day()))
	p.tillMonth.SetText(strconv.Itoa(int(date.Month())))
	p.tillYear.SetText(strconv.Itoa(date.Year()))

	p.firstName = widget.NewEntry()
	p.lastName = widget.NewEntry()
	p.result = widget.NewLabel("")

	p.content[layoutMobile] = container.NewScroll(container.NewVBox(
		container.NewHBox(widget.NewLabel("From Date: "), p.fromDay, p.fromMonth, p.fromYear),
		container.NewHBox(widget.NewLabel("Till Date: "), p.tillDay, p.tillMonth, p.tillYear),

		container.NewCenter(widget.NewButton("Load Data", p.onLoadClick)),
		widget.NewLabel("Querry Teacher:"),
		container.NewBorder(nil, nil, widget.NewLabel("Firstname: "), nil, p.firstName),
		container.NewBorder(nil, nil, widget.NewLabel("Lastname: "), nil, p.lastName),

		container.NewCenter(widget.NewButton("Querry", p.onQuerryClick)),
		p.result,
	))
	p.content[layoutDesktop] = p.content[layoutMobile]

	return p
}

func (p querryPanel) onLoadClick() {

	fromDay, err := strconv.Atoi(p.fromDay.Text)
	if err != nil {
		return
	}
	fromMonth, err := strconv.Atoi(p.fromMonth.Text)
	if err != nil {
		return
	}
	fromYear, err := strconv.Atoi(p.fromYear.Text)
	if err != nil {
		return
	}

	fromDate := time.Date(fromYear, time.Month(fromMonth), fromDay, 0, 0, 0, 0, UntisAPI.TimeZone())

	tillDay, err := strconv.Atoi(p.tillDay.Text)
	if err != nil {
		return
	}
	tillMonth, err := strconv.Atoi(p.tillMonth.Text)
	if err != nil {
		return
	}
	tillYear, err := strconv.Atoi(p.tillYear.Text)
	if err != nil {
		return
	}

	tillDate := time.Date(tillYear, time.Month(tillMonth), tillDay, 0, 0, 0, 0, UntisAPI.TimeZone())

	loadTimetable(UntisAPI.ToUntisDate(fromDate), UntisAPI.ToUntisDate(tillDate))
}

func (p querryPanel) onQuerryClick() {
	p.result.SetText(querryTeacher(p.firstName.Text, p.lastName.Text))
}


*/
