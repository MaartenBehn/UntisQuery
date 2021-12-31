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

	fromTime := widget.NewContainer(widget.ContainerOpts.Layout(widget.NewRowLayout(
		widget.RowLayoutOpts.Direction(widget.DirectionHorizontal),
		widget.RowLayoutOpts.Padding(widget.NewInsetsSimple(5)),
	)))

	fromDayInput := widget.NewTextInput(append(baseInputOpts,
		widget.TextInputOpts.Placeholder("From Day"),
	)...)
	fromDayInput.InputText = fmt.Sprint(now.Day())
	fromTime.AddChild(fromDayInput)

	fromMonthInput := widget.NewTextInput(append(baseInputOpts,
		widget.TextInputOpts.Placeholder("From Month"),
	)...)
	fromMonthInput.InputText = fmt.Sprint(int(now.Month()))
	fromTime.AddChild(fromMonthInput)

	fromYearInput := widget.NewTextInput(append(baseInputOpts,
		widget.TextInputOpts.Placeholder("From Month"),
	)...)
	fromYearInput.InputText = fmt.Sprint(now.Year())
	fromTime.AddChild(fromYearInput)

	c.AddChild(fromTime)

	tillTime := widget.NewContainer(widget.ContainerOpts.Layout(widget.NewRowLayout(
		widget.RowLayoutOpts.Direction(widget.DirectionHorizontal),
		widget.RowLayoutOpts.Padding(widget.NewInsetsSimple(5)),
	)))

	tillDayInput := widget.NewTextInput(append(baseInputOpts,
		widget.TextInputOpts.Placeholder("Till Day"),
	)...)
	tillDayInput.InputText = fmt.Sprint(now.Day())
	tillTime.AddChild(tillDayInput)

	tillMonthInput := widget.NewTextInput(append(baseInputOpts,
		widget.TextInputOpts.Placeholder("Till Month"),
	)...)
	tillMonthInput.InputText = fmt.Sprint(int(now.Month()))
	tillTime.AddChild(tillMonthInput)

	tillYearInput := widget.NewTextInput(append(baseInputOpts,
		widget.TextInputOpts.Placeholder("Till Month"),
	)...)
	tillYearInput.InputText = fmt.Sprint(now.Year())
	tillTime.AddChild(tillYearInput)

	c.AddChild(tillTime)

	c.AddChild(widget.NewButton(
		widget.ButtonOpts.Image(res.button.image),
		widget.ButtonOpts.Text("Query", res.button.face, res.button.text),
		widget.ButtonOpts.TextPadding(res.button.padding),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {

			teacher := teacherButton.SelectedEntry().(*state.Teacher)
			if teacher.Id < 0 {
				return
			}

			fromDay, err := strconv.Atoi(fromDayInput.InputText)
			if err != nil {
				return
			}
			fromMonth, err := strconv.Atoi(fromMonthInput.InputText)
			if err != nil {
				return
			}
			fromYear, err := strconv.Atoi(fromYearInput.InputText)
			if err != nil {
				return
			}

			fromDate := time.Date(fromYear, time.Month(fromMonth), fromDay, 0, 0, 0, 0, now.Location())

			tillDay, err := strconv.Atoi(tillDayInput.InputText)
			if err != nil {
				return
			}
			tillMonth, err := strconv.Atoi(tillMonthInput.InputText)
			if err != nil {
				return
			}
			tillYear, err := strconv.Atoi(tillYearInput.InputText)
			if err != nil {
				return
			}

			tillDate := time.Date(tillYear, time.Month(tillMonth), tillDay, 0, 0, 0, 0, now.Location())

			go func() {
				event.Go(event.EventLoadTimeTable, [2]time.Time{fromDate, tillDate})
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

	return c
}
