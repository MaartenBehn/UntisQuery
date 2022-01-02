package panel

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Stroby241/UntisQuery/event"
)

type startPanel struct {
	*panel
	name       *widget.Entry
	password   *widget.Entry
	school     *widget.Entry
	server     *widget.Entry
	errorLabel *widget.Label
}

func newStartPanel() *startPanel {
	p := &startPanel{
		panel:      newPanel(),
		name:       widget.NewEntry(),
		password:   widget.NewEntry(),
		school:     widget.NewEntry(),
		server:     widget.NewEntry(),
		errorLabel: widget.NewLabel(""),
	}

	p.password.Password = true

	p.content = container.NewVBox(
		p.loadingBar,
		widget.NewLabel("Login:"),
		container.NewBorder(nil, nil, widget.NewLabel("Username: "), nil, p.name),
		container.NewBorder(nil, nil, widget.NewLabel("Password: "), nil, p.password),
		container.NewBorder(nil, nil, widget.NewLabel("School: "), nil, p.school),
		container.NewBorder(nil, nil, widget.NewLabel("Server: "), nil, p.server),
		container.NewCenter(container.NewHBox(
			widget.NewButton("Login", func() {
				event.Go(event.EventLogin, [4]string{p.name.Text, p.password.Text, p.school.Text, p.server.Text})
			}),
			widget.NewButton("Fill [Debug]", func() {
				p.name.SetText("maarten8")
				p.password.SetText("behn500")
				p.school.SetText("TBZ Mitte Bremen")
				p.server.SetText("https://tipo.webuntis.com")
			}),
		)),
		p.errorLabel,
	)

	p.onShow = func() {
		p.errorLabel.SetText("")
		event.Go(event.EventLoading, 0.0)
		p.name.SetText("")
		p.password.SetText("")
		p.school.SetText("")
		p.server.SetText("")
	}

	event.On(event.EventLoginResult, func(data interface{}) {
		event.Go(event.EventLoading, 0.0)
		if data != nil {
			p.errorLabel.SetText(data.(error).Error())
		} else {
			event.Go(event.EventSetPanel, PanelQuery)
		}
	})

	return p
}
