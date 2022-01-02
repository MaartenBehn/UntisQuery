package panel

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Stroby241/UntisQuery/event"
)

type addTeacherPanel struct {
	*panel
	firstname  *widget.Entry
	lastname   *widget.Entry
	errorLabel *widget.Label
}

func newAddTeacherPanel() *addTeacherPanel {
	p := &addTeacherPanel{
		panel:      newPanel(),
		firstname:  widget.NewEntry(),
		lastname:   widget.NewEntry(),
		errorLabel: widget.NewLabel(""),
	}

	p.content = container.NewVBox(
		p.loadingBar,
		widget.NewLabel("Add Teacher:"),
		container.NewBorder(nil, nil, widget.NewLabel("Username: "), nil, p.firstname),
		container.NewBorder(nil, nil, widget.NewLabel("Password: "), nil, p.lastname),
		container.NewCenter(container.NewHBox(
			widget.NewButton("Add Teacher", func() {
				event.Go(event.EventAddTeacher, [2]string{p.firstname.Text, p.lastname.Text})
			}),
			widget.NewButton("Back", func() {
				event.Go(event.EventSetPanel, PanelQuery)
			}),
			widget.NewButton("Fill [Debug]", func() {
				p.firstname.SetText("Sven")
				p.lastname.SetText("Adiek")
			}),
		)),
		p.errorLabel,
	)

	p.onShow = func() {
		p.errorLabel.SetText("")
		event.Go(event.EventLoading, 0.0)
		p.firstname.SetText("")
		p.lastname.SetText("")
	}

	event.On(event.EventAddTeacherResult, func(data interface{}) {
		event.Go(event.EventLoading, 0.0)
		if data != nil {
			p.errorLabel.SetText(data.(error).Error())
		} else {
			event.Go(event.EventUpdateTeacherList, nil)
			event.Go(event.EventSetPanel, PanelQuery)
		}
	})

	return p
}
