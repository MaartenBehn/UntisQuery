package ui

import (
	"fmt"
	"github.com/Stroby241/UntisQuerry/event"
	"github.com/Stroby241/UntisQuerry/state"
	"github.com/blizzy78/ebitenui/widget"
	"strconv"
	"time"
)

func createQuerryPage(res *uiResources) widget.PreferredSizeLocateableWidget {
	c := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Padding(widget.NewInsetsSimple(5)),
		)),
	)

	c.AddChild(widget.NewLabel(
		widget.LabelOpts.Text("Querry", res.text.face, res.label.text)),
	)

	enties := []interface{}{}
	if len(state.Teachers) > 0 {
		for _, teacher := range state.Teachers {
			enties = append(enties, teacher)
		}
	} else {
		enties = append(enties, &state.Teacher{
			Firstname: "None",
			Lastname:  "",
			Id:        -1,
		})
	}

	teacherButton := newListComboButton(
		enties,
		func(e interface{}) string {
			t := e.(*state.Teacher)
			return t.Firstname + " " + t.Lastname
		},
		func(e interface{}) string {
			t := e.(*state.Teacher)
			return t.Firstname + " " + t.Lastname
		},
		func(args *widget.ListComboButtonEntrySelectedEventArgs) {
			c.RequestRelayout()
		},
		res,
	)
	c.AddChild(teacherButton)

	c.AddChild(widget.NewButton(
		widget.ButtonOpts.Image(res.button.image),
		widget.ButtonOpts.Text("Add Teacher", res.button.face, res.button.text),
		widget.ButtonOpts.TextPadding(res.button.padding),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			event.Go(event.EventSetPage, state.PageAddTeacher)
		}),
	))

	now := time.Now()

	baseInputOpts := []widget.TextInputOpt{
		widget.TextInputOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
			Stretch: true,
		})),
		widget.TextInputOpts.Image(res.textInput.image),
		widget.TextInputOpts.Color(res.textInput.color),
		widget.TextInputOpts.Padding(widget.Insets{
			Left:   13,
			Right:  13,
			Top:    7,
			Bottom: 7,
		}),
		widget.TextInputOpts.Face(res.textInput.face),
		widget.TextInputOpts.CaretOpts(
			widget.CaretOpts.Size(res.textInput.face, 2),
		),
	}

	timeContainer := widget.NewContainer(widget.ContainerOpts.Layout(widget.NewRowLayout(
		widget.RowLayoutOpts.Direction(widget.DirectionHorizontal),
		widget.RowLayoutOpts.Padding(widget.NewInsetsSimple(5)),
	)))

	dayInput := widget.NewTextInput(append(baseInputOpts,
		widget.TextInputOpts.Placeholder("From Day"),
	)...)
	dayInput.InputText = fmt.Sprint(now.Day())
	timeContainer.AddChild(dayInput)

	monthInput := widget.NewTextInput(append(baseInputOpts,
		widget.TextInputOpts.Placeholder("From Month"),
	)...)
	monthInput.InputText = fmt.Sprint(int(now.Month()))
	timeContainer.AddChild(monthInput)

	yearInput := widget.NewTextInput(append(baseInputOpts,
		widget.TextInputOpts.Placeholder("From Month"),
	)...)
	yearInput.InputText = fmt.Sprint(now.Year())
	timeContainer.AddChild(yearInput)

	c.AddChild(timeContainer)

	c.AddChild(widget.NewButton(
		widget.ButtonOpts.Image(res.button.image),
		widget.ButtonOpts.Text("Query", res.button.face, res.button.text),
		widget.ButtonOpts.TextPadding(res.button.padding),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {

			teacher := teacherButton.SelectedEntry().(*state.Teacher)
			if teacher.Id < 0 {
				return
			}

			day, err := strconv.Atoi(dayInput.InputText)
			if err != nil {
				return
			}
			month, err := strconv.Atoi(monthInput.InputText)
			if err != nil {
				return
			}
			year, err := strconv.Atoi(yearInput.InputText)
			if err != nil {
				return
			}

			fromDate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, now.Location())

			go func() {
				event.Go(event.EventLoadTimeTable, fromDate)
				event.Go(event.EventQuerryTaecher, teacher)
			}()
		}),
	))

	c.AddChild(widget.NewButton(
		widget.ButtonOpts.Image(res.button.image),
		widget.ButtonOpts.Text("Logout", res.button.face, res.button.text),
		widget.ButtonOpts.TextPadding(res.button.padding),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			event.Go(event.EventLogout, nil)
		}),
	))

	c.AddChild(widget.NewButton(
		widget.ButtonOpts.Image(res.button.image),
		widget.ButtonOpts.Text("Fill [Debug]", res.button.face, res.button.text),
		widget.ButtonOpts.TextPadding(res.button.padding),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			dayInput.InputText = "11"
			monthInput.InputText = "1"
			yearInput.InputText = "2022"
		}),
	))

	periodText := widget.NewLabel(widget.LabelOpts.Text("", res.text.face, res.label.text))
	c.AddChild(periodText)

	event.On(event.EventUpdateQuerryText, func(data interface{}) {
		periodText.Label = data.(string)
	})

	return c
}
