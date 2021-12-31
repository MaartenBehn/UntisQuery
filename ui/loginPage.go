package ui

import (
	"github.com/Stroby241/UntisQuerry/event"
	"github.com/Stroby241/UntisQuerry/state"
	"github.com/blizzy78/ebitenui/widget"
)

func createLoginPage(res *uiResources) widget.PreferredSizeLocateableWidget {
	c := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Padding(widget.NewInsetsSimple(5)),
		)),
	)

	c.AddChild(widget.NewLabel(
		widget.LabelOpts.Text("Login", res.text.face, res.label.text)),
	)

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

	username := widget.NewTextInput(append(baseInputOpts,
		widget.TextInputOpts.Placeholder("Username"),
	)...)
	c.AddChild(username)

	password := widget.NewTextInput(append(baseInputOpts,
		widget.TextInputOpts.Placeholder("Password"),
		widget.TextInputOpts.Secure(true),
	)...)
	c.AddChild(password)

	school := widget.NewTextInput(append(baseInputOpts,
		widget.TextInputOpts.Placeholder("Schoolname"),
	)...)
	c.AddChild(school)

	server := widget.NewTextInput(append(baseInputOpts,
		widget.TextInputOpts.Placeholder("Server"),
	)...)
	c.AddChild(server)

	c.AddChild(widget.NewButton(
		widget.ButtonOpts.Image(res.button.image),
		widget.ButtonOpts.Text("Login", res.button.face, res.button.text),
		widget.ButtonOpts.TextPadding(res.button.padding),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			event.Go(event.EventLogin, [4]string{username.InputText, password.InputText, school.InputText, server.InputText})
		}),
	))

	c.AddChild(widget.NewButton(
		widget.ButtonOpts.Image(res.button.image),
		widget.ButtonOpts.Text("Back", res.button.face, res.button.text),
		widget.ButtonOpts.TextPadding(res.button.padding),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			event.Go(event.EventSetPage, state.PageStart)
		}),
	))

	c.AddChild(widget.NewButton(
		widget.ButtonOpts.Image(res.button.image),
		widget.ButtonOpts.Text("Fill[Debug]", res.button.face, res.button.text),
		widget.ButtonOpts.TextPadding(res.button.padding),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			username.InputText = "maarten8"
			password.InputText = "behn500"
			school.InputText = "TBZ Mitte Bremen"
			server.InputText = "https://tipo.webuntis.com"
		}),
	))

	event.On(event.EventLoginResult, func(data interface{}) {
		if data.(bool) {
			event.Go(event.EventSetPage, state.PageQuerry)
		} else {
			event.Go(event.EventSetPage, state.PageStart)
		}
	})

	return c
}
