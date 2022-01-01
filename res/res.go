package res

import (
	"bytes"
	"embed"
	"image"
	"image/png"
)

//go:embed fonts/NotoSans-Bold.ttf
//go:embed fonts/NotoSans-Regular.ttf
//go:embed graphics/arrow-down-disabled.png
//go:embed graphics/arrow-down-idle.png
//go:embed graphics/button-disabled.png
//go:embed graphics/button-hover.png
//go:embed graphics/button-idle.png
//go:embed graphics/button-pressed.png
//go:embed graphics/button-selected-disabled.png
//go:embed graphics/button-selected-hover.png
//go:embed graphics/button-selected-idle.png
//go:embed graphics/button-selected-pressed.png
//go:embed graphics/checkbox-checked-disabled.png
//go:embed graphics/checkbox-checked-idle.png
//go:embed graphics/checkbox-disabled.png
//go:embed graphics/checkbox-greyed-disabled.png
//go:embed graphics/checkbox-greyed-idle.png
//go:embed graphics/checkbox-hover.png
//go:embed graphics/checkbox-idle.png
//go:embed graphics/checkbox-unchecked-disabled.png
//go:embed graphics/checkbox-unchecked-idle.png
//go:embed graphics/combo-button-disabled.png
//go:embed graphics/combo-button-hover.png
//go:embed graphics/combo-button-idle.png
//go:embed graphics/combo-button-pressed.png
//go:embed graphics/graphics.svg
//go:embed graphics/header.png
//go:embed graphics/list-disabled.png
//go:embed graphics/list-idle.png
//go:embed graphics/list-mask.png
//go:embed graphics/list-track-disabled.png
//go:embed graphics/list-track-idle.png
//go:embed graphics/panel-idle.png
//go:embed graphics/slider-handle-disabled.png
//go:embed graphics/slider-handle-hover.png
//go:embed graphics/slider-handle-idle.png
//go:embed graphics/slider-track-disabled.png
//go:embed graphics/slider-track-idle.png
//go:embed graphics/text-input-disabled.png
//go:embed graphics/text-input-hover.png
//go:embed graphics/text-input-idle.png
//go:embed graphics/tool-tip.png
var content embed.FS

func LoadByte(name string) ([]byte, error) {
	file, err := content.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func LoadImage(name string) (image.Image, error) {
	data, err := LoadByte(name)
	if err != nil {
		return nil, err
	}

	img, err := png.Decode(bytes.NewReader(data))
	return img, nil
}
