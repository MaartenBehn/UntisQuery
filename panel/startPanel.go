package panel

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Stroby241/UntisQuery/event"
	"github.com/Stroby241/UntisQuery/state"
)

type startPanel struct {
	*panel
	loginSelecter *widget.Select
	name          *widget.Entry
	password      *widget.Entry
	school        *widget.Entry
	server        *widget.Entry
	errorLabel    *widget.Label
}

func newStartPanel() *startPanel {
	p := &startPanel{
		panel:      newPanel(),
		name:       widget.NewEntry(),
		password:   widget.NewEntry(),
		school:     widget.NewEntry(),
		server:     widget.NewEntry(),
		errorLabel: widget.NewLabel(""),

		loginSelecter: widget.NewSelect([]string{}, func(s string) {}),
	}

	p.password.Password = true

	p.content = container.NewVBox(
		p.loadingBar,
		widget.NewLabel("Login:"),
		container.NewBorder(nil, nil, nil,
			container.NewHBox(widget.NewButton("Load", func() {
				var login *state.Login
				for _, l := range state.Logins {
					if l.Username == p.loginSelecter.Selected {
						login = l
					}
				}

				if login != nil {
					p.name.SetText(login.Username)
					p.password.SetText(login.Password)
					p.school.SetText(login.School)
					p.server.SetText(login.Server)
				}
			}),
				widget.NewButton("Remove", func() {
					for i, l := range state.Logins {
						if l.Username == p.loginSelecter.Selected {
							state.Logins = append(state.Logins[:i], state.Logins[i+1:]...)
							break
						}
					}
					state.LoadLogins()
					p.updateLoginList()
				})),
			p.loginSelecter),
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

		state.LoadLogins()
		p.updateLoginList()
	}

	event.On(event.EventLoginResult, func(data interface{}) {
		event.Go(event.EventLoading, 0.0)
		if data != nil {
			p.errorLabel.SetText(data.(error).Error())
		} else {
			event.Go(event.EventSetPanel, PanelQuery)

			found := false
			login := &state.Login{
				Username: p.name.Text,
				Password: p.password.Text,
				School:   p.school.Text,
				Server:   p.server.Text,
			}
			for i, l := range state.Logins {
				if l.Username == p.name.Text {
					state.Logins[i] = login
					found = true
					break
				}
			}
			if !found {
				state.Logins = append(state.Logins, login)
			}
			state.SaveLogins()
		}
	})

	return p
}

func (p *startPanel) updateLoginList() {
	var names []string
	for _, login := range state.Logins {
		names = append(names, login.Username)
	}
	p.loginSelecter.Options = names

	if len(names) == 1 {
		p.loginSelecter.SetSelected(names[0])
	}
}
