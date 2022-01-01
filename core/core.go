package core

import (
	"github.com/Stroby241/UntisQuerry/event"
	"github.com/Stroby241/UntisQuerry/state"
	"github.com/Stroby241/UntisQuerry/ui"
	"github.com/Stroby241/UntisQuerry/untis"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

func (g *Game) Update() error {
	event.Go(event.EventUpdate, nil)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	event.Go(event.EventDraw, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func NewGame() (*Game, error) {
	_, _, err := ui.Init()
	if err != nil {
		return nil, err
	}
	untis.Init()

	event.Go(event.EventSetPage, state.PageStart)

	game := &Game{}
	return game, nil
}
