package mobile

import (
	"github.com/Stroby241/UntisQuerry/core"
	"github.com/hajimehoshi/ebiten/v2/mobile"
)

func init() {
	game, err := core.NewGame()
	if err != nil {
		panic(err)
	}
	mobile.SetGame(game)
}

// Dummy is a dummy exported function.
//
// gomobile doesn't compile a package that doesn't include any exported function.
// Dummy forces gomobile to compile this package.
func Dummy() {}
