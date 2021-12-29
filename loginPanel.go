package main

/*
import (
	"UntisAPI"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type loginPanel struct {
	*panel
	name       *widget.Entry
	password   *widget.Entry
	school     *widget.Entry
	server     *widget.Entry
	errorLable *widget.Label
}

func newLoginPanel() *loginPanel {
	p := &loginPanel{panel: newPanel()}

	p.name = widget.NewEntry()
	p.password = widget.NewEntry()
	p.school = widget.NewEntry()
	p.server = widget.NewEntry()
	p.errorLable = widget.NewLabel("")

	p.content[layoutMobile] = container.NewVBox(
		widget.NewLabel("Login:"),
		container.NewBorder(nil, nil, widget.NewLabel("Username: "), nil, p.name),
		container.NewBorder(nil, nil, widget.NewLabel("Password: "), nil, p.password),
		container.NewBorder(nil, nil, widget.NewLabel("School: "), nil, p.school),
		container.NewBorder(nil, nil, widget.NewLabel("Server: "), nil, p.server),
		container.NewCenter(widget.NewButton("Login", p.onloginClick)),
		p.errorLable,
	)

	// Presetting my user
	p.name.SetText("maarten8")
	p.password.SetText("behn500")
	p.school.SetText("TBZ Mitte Bremen")
	p.server.SetText("https://tipo.webuntis.com")

	p.content[layoutDesktop] = p.content[layoutMobile]

	return p
}

func (p *loginPanel) onloginClick() {
	user = UntisAPI.NewUser(p.name.Text, p.password.Text, p.school.Text, p.server.Text)

	defer func() {
		if err := recover(); err != nil {
			p.errorLable.SetText("Wrong input!")
		}
	}()
	user.Login()

	if user.LoggedIn {
		currentPanel = panelIdQuerry
	} else {
		p.errorLable.SetText("Wrong input!")
	}
}


*/
