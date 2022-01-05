package panel

import (
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Stroby241/UntisQuery/event"
	"github.com/Stroby241/UntisQuery/state"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type queryPanel struct {
	*panel
	teacherSelecter *widget.Select
	day             *Entry
	month           *Entry
	year            *Entry
	peroidList      *widget.List
	errorLabel      *widget.Label

	needQuery bool
}

func newQueryPanel(window *fyne.Window) *queryPanel {
	p := &queryPanel{
		panel: newPanel(),
		teacherSelecter: widget.NewSelect([]string{}, func(s string) {

		}),
		day:   newEntry(),
		month: newEntry(),
		year:  newEntry(),
		peroidList: widget.NewList(func() int {
			return len(state.Periods)
		}, func() fyne.CanvasObject {
			return widget.NewLabel("")
		}, func(id widget.ListItemID, o fyne.CanvasObject) {
			p := state.Periods[id]

			text := fmt.Sprintf("%02d:%02d > %02d:%02d |",
				p.StartTime.Hour(),
				p.StartTime.Minute(),
				p.EndTime.Hour(),
				p.EndTime.Minute(),
			)
			for _, s := range p.Rooms {
				text += " " + s
			}
			text += " |"
			for _, s := range p.Classes {
				text += " " + s
			}
			text += " |"
			for _, s := range p.Subjects {
				text += " " + s
			}
			o.(*widget.Label).SetText(text)

			// TODO This produces ugly result
			size := (*window).Canvas().Size()
			for o.(*widget.Label).MinSize().Width > size.Width {
				i := len(text) / 2
				text = text[:i] + "\n" + text[i:]
				o.(*widget.Label).SetText(text)
			}
		}),

		errorLabel: widget.NewLabel(""),
	}

	now := time.Now()
	p.day.SetText(fmt.Sprintf("%d", now.Day()))
	p.month.SetText(fmt.Sprintf("%d", now.Month()))
	p.year.SetText(fmt.Sprintf("%d", now.Year()))

	p.day.Validator = func(s string) error {
		if len(s) == 0 {
			return errors.New("day can't be empty")
		}
		if len(s) > 2 {
			return errors.New("day can't be more than two runes")
		}
		for _, r := range s {
			if !unicode.IsDigit(r) {
				return errors.New("day can only contain digits")
			}
		}

		return nil
	}
	p.month.Validator = func(s string) error {
		if len(s) == 0 {
			return errors.New("month can't be empty")
		}
		if len(s) > 2 {
			return errors.New("month can't be more than two runes")
		}
		for _, r := range s {
			if !unicode.IsDigit(r) {
				return errors.New("month can only contain digits")
			}
		}

		return nil
	}
	p.year.Validator = func(s string) error {
		if len(s) != 4 {
			return errors.New("year must contain 4 digtits")
		}
		for _, r := range s {
			if !unicode.IsDigit(r) {
				return errors.New("day can only contain digits")
			}
		}

		return nil
	}
	p.year.plusMinSize = fyne.NewSize(10, 0)

	p.day.OnChanged = p.onDateInputChange
	p.month.OnChanged = p.onDateInputChange
	p.year.OnChanged = p.onDateInputChange

	p.teacherSelecter.OnChanged = func(s string) {
		p.needQuery = true
	}

	p.content = container.NewBorder(
		container.NewVBox(
			p.loadingBar,
			container.NewBorder(nil, nil, nil,
				widget.NewButton("Logout", func() {
					event.Go(event.EventLogout, nil)
					event.Go(event.EventSetPanel, PanelStart)
				}),
				widget.NewLabel("Query:"),
			),
			container.NewBorder(nil, nil, nil,
				widget.NewButton("Add Teacher", func() {
					event.Go(event.EventSetPanel, PanelAddTeacher)
				}),
				p.teacherSelecter,
			),
			container.NewHBox(
				widget.NewLabel("Date: "),
				p.day,
				p.month,
				p.year,
				widget.NewButton("+", func() {
					date, err := p.getDate()
					if err != nil {
						p.errorLabel.SetText(err.Error())
						return
					}
					p.day.SetText(fmt.Sprintf("%d", date.Day()+1))
					p.getDate()
				}),
				widget.NewButton("-", func() {
					date, err := p.getDate()
					if err != nil {
						p.errorLabel.SetText(err.Error())
						return
					}
					p.day.SetText(fmt.Sprintf("%d", date.Day()-1))
					p.getDate()
				}),
			),
		),
		p.errorLabel,
		nil,
		nil,
		p.peroidList,
	)

	p.onShow = func() {
		p.errorLabel.SetText("")
		event.Go(event.EventLoading, 0.0)

		state.LoadTeacher()
		event.Go(event.EventUpdateTeacherList, nil)
	}

	event.On(event.EventUpdateTeacherList, func(data interface{}) {
		var names []string
		for _, teacher := range state.Teachers {
			names = append(names, teacher.Firstname+" "+teacher.Lastname)
		}
		p.teacherSelecter.Options = names

		if len(names) == 1 {
			p.teacherSelecter.SetSelected(names[0])
		}
	})

	event.On(event.EventLoadTimeTableResult, func(data interface{}) {
		event.Go(event.EventLoading, 0.0)
		if data != nil {
			p.errorLabel.SetText(data.(error).Error())
		} else {
			p.loadTeacher()
		}
	})

	event.On(event.EventQuerryTaecherResult, func(data interface{}) {
		event.Go(event.EventLoading, 0.0)
		if data != nil {
			p.errorLabel.SetText(data.(error).Error())
		} else {
			p.peroidList.Refresh()
		}
	})

	event.On(event.EventUpdate, func(data interface{}) {
		if p.needQuery {
			p.needQuery = false
			p.query()
		}
	})

	return p
}

func (p *queryPanel) onDateInputChange(s string) {
	_, err := p.getDate()
	if err != nil {
		p.errorLabel.SetText(err.Error())
	} else {
		p.errorLabel.SetText("")
		p.needQuery = true
	}
}

func (p *queryPanel) getDate() (time.Time, error) {
	day, err := strconv.Atoi(p.day.Text)
	if err != nil {
		return time.Time{}, errors.New("day can't be converted to int")
	}
	month, err := strconv.Atoi(p.month.Text)
	if err != nil {
		return time.Time{}, errors.New("month can't be converted to int")
	}
	year, err := strconv.Atoi(p.year.Text)
	if err != nil {
		return time.Time{}, errors.New("year can't be converted to int")
	}

	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Now().Location())

	p.day.SetText(strconv.Itoa(date.Day()))
	p.month.SetText(strconv.Itoa(int(date.Month())))
	p.year.SetText(strconv.Itoa(date.Year()))

	return date, nil
}

func (p *queryPanel) query() {
	err := p.day.Validate()
	if err != nil {
		p.errorLabel.SetText(err.Error())
		return
	}
	err = p.month.Validate()
	if err != nil {
		p.errorLabel.SetText(err.Error())
		return
	}
	err = p.year.Validate()
	if err != nil {
		p.errorLabel.SetText(err.Error())
		return
	}
	date, err := p.getDate()
	if err != nil {
		p.errorLabel.SetText(err.Error())
		return
	}

	event.Go(event.EventLoadTimeTable, date)
	p.content.Refresh()
}

func (p *queryPanel) loadTeacher() {
	teacherNames := strings.Split(p.teacherSelecter.Selected, " ")

	var teacher *state.Teacher
	for _, t := range state.Teachers {
		if t.Firstname == teacherNames[0] && t.Lastname == teacherNames[1] {
			teacher = t
		}
	}

	if teacher == nil {
		p.errorLabel.SetText("selected Teacher and Teacher list don't match")
		return
	}

	event.Go(event.EventQuerryTaecher, teacher)
}
