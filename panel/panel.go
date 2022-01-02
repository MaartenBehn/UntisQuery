package panel

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/Stroby241/UntisQuery/event"
)

const (
	PanelStart      = 0
	PanelQuery      = 1
	PanelAddTeacher = 2
	panelMax        = 3
)

type panel struct {
	content    fyne.CanvasObject
	onShow     func()
	loadingBar *widget.ProgressBar
}

var panels []*panel

func newPanel() *panel {
	p := &panel{
		loadingBar: widget.NewProgressBar(),
	}

	event.On(event.EventLoading, func(data interface{}) {
		p.loadingBar.SetValue(data.(float64))
	})

	return p
}

func Init(window *fyne.Window) {
	panels = make([]*panel, panelMax)
	panels[PanelStart] = newStartPanel().panel
	panels[PanelQuery] = newQueryPanel().panel
	panels[PanelAddTeacher] = newAddTeacherPanel().panel

	event.On(event.EventSetPanel, func(data interface{}) {
		panel := panels[data.(int)]
		if panel.onShow != nil {
			panel.onShow()
		}
		(*window).SetContent(panel.content)
		(*window).Canvas().Content().Refresh()
		(*window).Show()
	})
}
