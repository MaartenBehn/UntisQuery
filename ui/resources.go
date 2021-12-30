package ui

import (
	"github.com/blizzy78/ebitenui/image"
	"github.com/blizzy78/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/font"
	"image/color"
	"strconv"
)

const (
	backgroundColor = "131a22"

	textIdleColor     = "dff4ff"
	textDisabledColor = "5a7a91"

	labelIdleColor     = textIdleColor
	labelDisabledColor = textDisabledColor

	buttonIdleColor     = textIdleColor
	buttonDisabledColor = labelDisabledColor

	listSelectedBackground         = "4b687a"
	listDisabledSelectedBackground = "2a3944"

	headerColor = textIdleColor

	textInputCaretColor         = "e7c34b"
	textInputDisabledCaretColor = "766326"

	toolTipColor = backgroundColor

	separatorColor = listDisabledSelectedBackground
)

type uiResources struct {
	fonts *Font

	background *image.NineSlice

	separatorColor color.Color

	text        *textResources
	button      *buttonResources
	label       *labelResources
	checkbox    *checkboxResources
	comboButton *comboButtonResources
	list        *listResources
	slider      *sliderResources
	panel       *panelResources
	tabBook     *tabBookResources
	header      *headerResources
	textInput   *textInputResources
	toolTip     *toolTipResources
}

type textResources struct {
	idleColor     color.Color
	disabledColor color.Color
	face          font.Face
	titleFace     font.Face
	bigTitleFace  font.Face
	smallFace     font.Face
}

type buttonResources struct {
	image   *widget.ButtonImage
	text    *widget.ButtonTextColor
	face    font.Face
	padding widget.Insets
}

type checkboxResources struct {
	image   *widget.ButtonImage
	graphic *widget.CheckboxGraphicImage
	spacing int
}

type labelResources struct {
	text *widget.LabelColor
	face font.Face
}

type comboButtonResources struct {
	image   *widget.ButtonImage
	text    *widget.ButtonTextColor
	face    font.Face
	graphic *widget.ButtonImageImage
	padding widget.Insets
}

type listResources struct {
	image        *widget.ScrollContainerImage
	track        *widget.SliderTrackImage
	trackPadding widget.Insets
	handle       *widget.ButtonImage
	handleSize   int
	face         font.Face
	entry        *widget.ListEntryColor
	entryPadding widget.Insets
}

type sliderResources struct {
	trackImage *widget.SliderTrackImage
	handle     *widget.ButtonImage
	handleSize int
}

type panelResources struct {
	image   *image.NineSlice
	padding widget.Insets
}

type tabBookResources struct {
	idleButton     *widget.ButtonImage
	selectedButton *widget.ButtonImage
	buttonFace     font.Face
	buttonText     *widget.ButtonTextColor
	buttonPadding  widget.Insets
}

type headerResources struct {
	background *image.NineSlice
	padding    widget.Insets
	face       font.Face
	color      color.Color
}

type textInputResources struct {
	image   *widget.TextInputImage
	padding widget.Insets
	face    font.Face
	color   *widget.TextInputColor
}

type toolTipResources struct {
	background *image.NineSlice
	padding    widget.Insets
	face       font.Face
	color      color.Color
}

func newUIResources() (*uiResources, error) {
	background := image.NewNineSliceColor(hexToColor(backgroundColor))

	fonts, err := LoadFonts()
	if err != nil {
		return nil, err
	}

	button, err := newButtonResources(fonts)
	if err != nil {
		return nil, err
	}

	checkbox, err := newCheckboxResources()
	if err != nil {
		return nil, err
	}

	comboButton, err := newComboButtonResources(fonts)
	if err != nil {
		return nil, err
	}

	list, err := newListResources(fonts)
	if err != nil {
		return nil, err
	}

	slider, err := newSliderResources()
	if err != nil {
		return nil, err
	}

	panel, err := newPanelResources()
	if err != nil {
		return nil, err
	}

	tabBook, err := newTabBookResources(fonts)
	if err != nil {
		return nil, err
	}

	header, err := newHeaderResources(fonts)
	if err != nil {
		return nil, err
	}

	textInput, err := newTextInputResources(fonts)
	if err != nil {
		return nil, err
	}

	toolTip, err := newToolTipResources(fonts)
	if err != nil {
		return nil, err
	}

	return &uiResources{
		fonts: fonts,

		background: background,

		separatorColor: hexToColor(separatorColor),

		text: &textResources{
			idleColor:     hexToColor(textIdleColor),
			disabledColor: hexToColor(textDisabledColor),
			face:          fonts.Face,
			titleFace:     fonts.TitleFace,
			bigTitleFace:  fonts.BigTitleFace,
			smallFace:     fonts.ToolTipFace,
		},

		button:      button,
		label:       newLabelResources(fonts),
		checkbox:    checkbox,
		comboButton: comboButton,
		list:        list,
		slider:      slider,
		panel:       panel,
		tabBook:     tabBook,
		header:      header,
		textInput:   textInput,
		toolTip:     toolTip,
	}, nil
}

func newButtonResources(fonts *Font) (*buttonResources, error) {
	idle, err := loadImageNineSlice("res/graphics/button-idle.png", 12, 0)
	if err != nil {
		return nil, err
	}

	hover, err := loadImageNineSlice("res/graphics/button-hover.png", 12, 0)
	if err != nil {
		return nil, err
	}

	pressed, err := loadImageNineSlice("res/graphics/button-pressed.png", 12, 0)
	if err != nil {
		return nil, err
	}

	disabled, err := loadImageNineSlice("res/graphics/button-disabled.png", 12, 0)
	if err != nil {
		return nil, err
	}

	i := &widget.ButtonImage{
		Idle:     idle,
		Hover:    hover,
		Pressed:  pressed,
		Disabled: disabled,
	}

	return &buttonResources{
		image: i,

		text: &widget.ButtonTextColor{
			Idle:     hexToColor(buttonIdleColor),
			Disabled: hexToColor(buttonDisabledColor),
		},

		face: fonts.Face,

		padding: widget.Insets{
			Left:  30,
			Right: 30,
		},
	}, nil
}

func newCheckboxResources() (*checkboxResources, error) {
	idle, err := loadImageNineSlice("res/graphics/checkbox-idle.png", 20, 0)
	if err != nil {
		return nil, err
	}

	hover, err := loadImageNineSlice("res/graphics/checkbox-hover.png", 20, 0)
	if err != nil {
		return nil, err
	}

	disabled, err := loadImageNineSlice("res/graphics/checkbox-disabled.png", 20, 0)
	if err != nil {
		return nil, err
	}

	checked, err := loadGraphicImages("res/graphics/checkbox-checked-idle.png", "res/graphics/checkbox-checked-disabled.png")
	if err != nil {
		return nil, err
	}

	unchecked, err := loadGraphicImages("res/graphics/checkbox-unchecked-idle.png", "res/graphics/checkbox-unchecked-disabled.png")
	if err != nil {
		return nil, err
	}

	greyed, err := loadGraphicImages("res/graphics/checkbox-greyed-idle.png", "res/graphics/checkbox-greyed-disabled.png")
	if err != nil {
		return nil, err
	}

	return &checkboxResources{
		image: &widget.ButtonImage{
			Idle:     idle,
			Hover:    hover,
			Pressed:  hover,
			Disabled: disabled,
		},

		graphic: &widget.CheckboxGraphicImage{
			Checked:   checked,
			Unchecked: unchecked,
			Greyed:    greyed,
		},

		spacing: 10,
	}, nil
}

func newLabelResources(fonts *Font) *labelResources {
	return &labelResources{
		text: &widget.LabelColor{
			Idle:     hexToColor(labelIdleColor),
			Disabled: hexToColor(labelDisabledColor),
		},

		face: fonts.Face,
	}
}

func newComboButtonResources(fonts *Font) (*comboButtonResources, error) {
	idle, err := loadImageNineSlice("res/graphics/combo-button-idle.png", 12, 0)
	if err != nil {
		return nil, err
	}

	hover, err := loadImageNineSlice("res/graphics/combo-button-hover.png", 12, 0)
	if err != nil {
		return nil, err
	}

	pressed, err := loadImageNineSlice("res/graphics/combo-button-pressed.png", 12, 0)
	if err != nil {
		return nil, err
	}

	disabled, err := loadImageNineSlice("res/graphics/combo-button-disabled.png", 12, 0)
	if err != nil {
		return nil, err
	}

	i := &widget.ButtonImage{
		Idle:     idle,
		Hover:    hover,
		Pressed:  pressed,
		Disabled: disabled,
	}

	arrowDown, err := loadGraphicImages("res/graphics/arrow-down-idle.png", "res/graphics/arrow-down-disabled.png")
	if err != nil {
		return nil, err
	}

	return &comboButtonResources{
		image: i,

		text: &widget.ButtonTextColor{
			Idle:     hexToColor(buttonIdleColor),
			Disabled: hexToColor(buttonDisabledColor),
		},

		face:    fonts.Face,
		graphic: arrowDown,

		padding: widget.Insets{
			Left:  30,
			Right: 30,
		},
	}, nil
}

func newListResources(fonts *Font) (*listResources, error) {
	idle, _, err := ebitenutil.NewImageFromFile("res/graphics/list-idle.png")
	if err != nil {
		return nil, err
	}

	disabled, _, err := ebitenutil.NewImageFromFile("res/graphics/list-disabled.png")
	if err != nil {
		return nil, err
	}

	mask, _, err := ebitenutil.NewImageFromFile("res/graphics/list-mask.png")
	if err != nil {
		return nil, err
	}

	trackIdle, _, err := ebitenutil.NewImageFromFile("res/graphics/list-track-idle.png")
	if err != nil {
		return nil, err
	}

	trackDisabled, _, err := ebitenutil.NewImageFromFile("res/graphics/list-track-disabled.png")
	if err != nil {
		return nil, err
	}

	handleIdle, _, err := ebitenutil.NewImageFromFile("res/graphics/slider-handle-idle.png")
	if err != nil {
		return nil, err
	}

	handleHover, _, err := ebitenutil.NewImageFromFile("res/graphics/slider-handle-hover.png")
	if err != nil {
		return nil, err
	}

	return &listResources{
		image: &widget.ScrollContainerImage{
			Idle:     image.NewNineSlice(idle, [3]int{25, 12, 22}, [3]int{25, 12, 25}),
			Disabled: image.NewNineSlice(disabled, [3]int{25, 12, 22}, [3]int{25, 12, 25}),
			Mask:     image.NewNineSlice(mask, [3]int{26, 10, 23}, [3]int{26, 10, 26}),
		},

		track: &widget.SliderTrackImage{
			Idle:     image.NewNineSlice(trackIdle, [3]int{5, 0, 0}, [3]int{25, 12, 25}),
			Hover:    image.NewNineSlice(trackIdle, [3]int{5, 0, 0}, [3]int{25, 12, 25}),
			Disabled: image.NewNineSlice(trackDisabled, [3]int{0, 5, 0}, [3]int{25, 12, 25}),
		},

		trackPadding: widget.Insets{
			Top:    5,
			Bottom: 24,
		},

		handle: &widget.ButtonImage{
			Idle:     image.NewNineSliceSimple(handleIdle, 0, 5),
			Hover:    image.NewNineSliceSimple(handleHover, 0, 5),
			Pressed:  image.NewNineSliceSimple(handleHover, 0, 5),
			Disabled: image.NewNineSliceSimple(handleIdle, 0, 5),
		},

		handleSize: 5,
		face:       fonts.Face,

		entry: &widget.ListEntryColor{
			Unselected:         hexToColor(textIdleColor),
			DisabledUnselected: hexToColor(textDisabledColor),

			Selected:         hexToColor(textIdleColor),
			DisabledSelected: hexToColor(textDisabledColor),

			SelectedBackground:         hexToColor(listSelectedBackground),
			DisabledSelectedBackground: hexToColor(listDisabledSelectedBackground),
		},

		entryPadding: widget.Insets{
			Left:   30,
			Right:  30,
			Top:    2,
			Bottom: 2,
		},
	}, nil
}

func newSliderResources() (*sliderResources, error) {
	idle, _, err := ebitenutil.NewImageFromFile("res/graphics/slider-track-idle.png")
	if err != nil {
		return nil, err
	}

	disabled, _, err := ebitenutil.NewImageFromFile("res/graphics/slider-track-disabled.png")
	if err != nil {
		return nil, err
	}

	handleIdle, _, err := ebitenutil.NewImageFromFile("res/graphics/slider-handle-idle.png")
	if err != nil {
		return nil, err
	}

	handleHover, _, err := ebitenutil.NewImageFromFile("res/graphics/slider-handle-hover.png")
	if err != nil {
		return nil, err
	}

	handleDisabled, _, err := ebitenutil.NewImageFromFile("res/graphics/slider-handle-disabled.png")
	if err != nil {
		return nil, err
	}

	return &sliderResources{
		trackImage: &widget.SliderTrackImage{
			Idle:     image.NewNineSlice(idle, [3]int{0, 19, 0}, [3]int{6, 0, 0}),
			Hover:    image.NewNineSlice(idle, [3]int{0, 19, 0}, [3]int{6, 0, 0}),
			Disabled: image.NewNineSlice(disabled, [3]int{0, 19, 0}, [3]int{6, 0, 0}),
		},

		handle: &widget.ButtonImage{
			Idle:     image.NewNineSliceSimple(handleIdle, 0, 5),
			Hover:    image.NewNineSliceSimple(handleHover, 0, 5),
			Pressed:  image.NewNineSliceSimple(handleHover, 0, 5),
			Disabled: image.NewNineSliceSimple(handleDisabled, 0, 5),
		},

		handleSize: 6,
	}, nil
}

func newPanelResources() (*panelResources, error) {
	i, err := loadImageNineSlice("res/graphics/panel-idle.png", 10, 10)
	if err != nil {
		return nil, err
	}

	return &panelResources{
		image: i,
		padding: widget.Insets{
			Left:   30,
			Right:  30,
			Top:    20,
			Bottom: 20,
		},
	}, nil
}

func newTabBookResources(fonts *Font) (*tabBookResources, error) {
	selectedIdle, err := loadImageNineSlice("res/graphics/button-selected-idle.png", 12, 0)
	if err != nil {
		return nil, err
	}

	selectedHover, err := loadImageNineSlice("res/graphics/button-selected-hover.png", 12, 0)
	if err != nil {
		return nil, err
	}

	selectedPressed, err := loadImageNineSlice("res/graphics/button-selected-pressed.png", 12, 0)
	if err != nil {
		return nil, err
	}

	selectedDisabled, err := loadImageNineSlice("res/graphics/button-selected-disabled.png", 12, 0)
	if err != nil {
		return nil, err
	}

	selected := &widget.ButtonImage{
		Idle:     selectedIdle,
		Hover:    selectedHover,
		Pressed:  selectedPressed,
		Disabled: selectedDisabled,
	}

	idle, err := loadImageNineSlice("res/graphics/button-idle.png", 12, 0)
	if err != nil {
		return nil, err
	}

	hover, err := loadImageNineSlice("res/graphics/button-hover.png", 12, 0)
	if err != nil {
		return nil, err
	}

	pressed, err := loadImageNineSlice("res/graphics/button-pressed.png", 12, 0)
	if err != nil {
		return nil, err
	}

	disabled, err := loadImageNineSlice("res/graphics/button-disabled.png", 12, 0)
	if err != nil {
		return nil, err
	}

	unselected := &widget.ButtonImage{
		Idle:     idle,
		Hover:    hover,
		Pressed:  pressed,
		Disabled: disabled,
	}

	return &tabBookResources{
		selectedButton: selected,
		idleButton:     unselected,
		buttonFace:     fonts.Face,

		buttonText: &widget.ButtonTextColor{
			Idle:     hexToColor(buttonIdleColor),
			Disabled: hexToColor(buttonDisabledColor),
		},

		buttonPadding: widget.Insets{
			Left:  30,
			Right: 30,
		},
	}, nil
}

func newHeaderResources(fonts *Font) (*headerResources, error) {
	bg, err := loadImageNineSlice("res/graphics/header.png", 446, 9)
	if err != nil {
		return nil, err
	}

	return &headerResources{
		background: bg,

		padding: widget.Insets{
			Left:   25,
			Right:  25,
			Top:    4,
			Bottom: 4,
		},

		face:  fonts.BigTitleFace,
		color: hexToColor(headerColor),
	}, nil
}

func newTextInputResources(fonts *Font) (*textInputResources, error) {
	idle, _, err := ebitenutil.NewImageFromFile("res/graphics/text-input-idle.png")
	if err != nil {
		return nil, err
	}

	disabled, _, err := ebitenutil.NewImageFromFile("res/graphics/text-input-disabled.png")
	if err != nil {
		return nil, err
	}

	return &textInputResources{
		image: &widget.TextInputImage{
			Idle:     image.NewNineSlice(idle, [3]int{9, 14, 6}, [3]int{9, 14, 6}),
			Disabled: image.NewNineSlice(disabled, [3]int{9, 14, 6}, [3]int{9, 14, 6}),
		},

		padding: widget.Insets{
			Left:   8,
			Right:  8,
			Top:    4,
			Bottom: 4,
		},

		face: fonts.Face,

		color: &widget.TextInputColor{
			Idle:          hexToColor(textIdleColor),
			Disabled:      hexToColor(textDisabledColor),
			Caret:         hexToColor(textInputCaretColor),
			DisabledCaret: hexToColor(textInputDisabledCaretColor),
		},
	}, nil
}

func newToolTipResources(fonts *Font) (*toolTipResources, error) {
	bg, _, err := ebitenutil.NewImageFromFile("res/graphics/tool-tip.png")
	if err != nil {
		return nil, err
	}

	return &toolTipResources{
		background: image.NewNineSlice(bg, [3]int{19, 6, 13}, [3]int{19, 5, 13}),

		padding: widget.Insets{
			Left:   15,
			Right:  15,
			Top:    10,
			Bottom: 10,
		},

		face:  fonts.ToolTipFace,
		color: hexToColor(toolTipColor),
	}, nil
}

func (u *uiResources) close() {
	u.fonts.Close()
}

func hexToColor(h string) color.Color {

	u, err := strconv.ParseUint(h, 16, 0)
	if err != nil {
		panic(err)
	}

	if len(h) == 6 {
		return color.RGBA{
			R: uint8(u & 0xff0000 >> 16),
			G: uint8(u & 0xff00 >> 8),
			B: uint8(u & 0xff),
			A: 255,
		}
	} else if len(h) == 8 {
		return color.RGBA{
			R: uint8(u & 0xff000000 >> 24),
			G: uint8(u & 0xff0000 >> 16),
			B: uint8(u & 0xff00 >> 8),
			A: uint8(u & 0xff),
		}
	}
	return color.RGBA{}
}
