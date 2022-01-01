package ui

import (
	"github.com/Stroby241/UntisQuerry/res"
	"github.com/blizzy78/ebitenui/image"
	"github.com/blizzy78/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
)

func loadGraphicImages(idle string, disabled string) (*widget.ButtonImageImage, error) {

	img, err := res.LoadImage(idle)
	if err != nil {
		return nil, err
	}
	idleImage := ebiten.NewImageFromImage(img)

	var disabledImage *ebiten.Image
	if disabled != "" {

		img, err := res.LoadImage(disabled)
		if err != nil {
			return nil, err
		}
		disabledImage = ebiten.NewImageFromImage(img)

		if err != nil {
			return nil, err
		}
	}

	return &widget.ButtonImageImage{
		Idle:     idleImage,
		Disabled: disabledImage,
	}, nil
}

func loadImage(name string) (*ebiten.Image, error) {
	img, err := res.LoadImage(name)
	if err != nil {
		return nil, err
	}
	ebietnImage := ebiten.NewImageFromImage(img)
	return ebietnImage, nil
}

func loadImageNineSlice(path string, centerWidth int, centerHeight int) (*image.NineSlice, error) {
	img, err := res.LoadImage(path)
	if err != nil {
		return nil, err
	}
	ebitenImage := ebiten.NewImageFromImage(img)

	w, h := ebitenImage.Size()
	return image.NewNineSlice(ebitenImage,
			[3]int{(w - centerWidth) / 2, centerWidth, w - (w-centerWidth)/2 - centerWidth},
			[3]int{(h - centerHeight) / 2, centerHeight, h - (h-centerHeight)/2 - centerHeight}),
		nil
}
