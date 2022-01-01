package ui

import (
	"fmt"
	"github.com/Stroby241/UntisQuerry/event"
	"github.com/Stroby241/UntisQuerry/state"
	"github.com/blizzy78/ebitenui/widget"
	"time"
)

func createLoadingPage(res *uiResources) widget.PreferredSizeLocateableWidget {
	c := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Padding(widget.NewInsetsSimple(5)),
		)),
	)

	text := widget.NewLabel(widget.LabelOpts.Text("", res.text.face, res.label.text))
	c.AddChild(text)

	var name string

	startTime := time.Now()

	event.On(event.EventStartLoading, func(data interface{}) {
		name = data.(string)
		text.Label = fmt.Sprintf("Loading: %s\n", name)

		startTime = time.Now()
		event.Go(event.EventSetPage, state.PageLoading)
	})

	event.On(event.EventUpdateLoading, func(data interface{}) {
		percent := data.(float64)
		remainig := time.Since(startTime).Seconds() / percent * (100 - percent)

		text.Label = fmt.Sprintf("Loading: %s\n %.5f Percent \n %f sec remaining", name, percent, remainig)
	})

	return c
}
