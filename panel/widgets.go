package panel

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type Entry struct {
	plusMinSize fyne.Size
	widget.Entry
}

func newEntry() *Entry {
	e := &Entry{}
	e.plusMinSize = fyne.NewSize(0, 0)
	e.ExtendBaseWidget(e)
	return e
}

func (e *Entry) MinSize() fyne.Size {
	return e.Entry.MinSize().Add(e.plusMinSize)
}
