package ui

import (
	"image/color"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/milk9111/left-behind/assets"
	"golang.org/x/image/colornames"
)

var (
	NormalTextColor    = RGB(0x9dd793)
	CaretColor         = RGB(0xe7c34b)
	disabledCaretColor = RGB(0x766326)
)

type Widget = widget.PreferredSizeLocateableWidget

type Resources struct {
	Button    *ButtonResource
	Panel     *PanelResource
	DarkPanel *PanelResource

	Font1 text.Face
	Font2 text.Face
	Font3 text.Face
}

type TextInputResource struct {
	Image      *widget.TextInputImage
	Padding    widget.Insets
	TextColors *widget.TextInputColor
}

type PanelResource struct {
	Image   *image.NineSlice
	Padding widget.Insets
}

type ButtonResource struct {
	Image      *widget.ButtonImage
	Padding    widget.Insets
	TextColors *widget.ButtonTextColor
}

type ToggleButtonResource struct {
	Image    *widget.ButtonImage
	AltImage *widget.ButtonImage
	Padding  widget.Insets
	Color    color.Color
	AltColor color.Color
}

type OptionButtonResource struct {
	Image      *widget.ButtonImage
	Padding    widget.Insets
	TextColors *widget.ButtonTextColor
	FontFace   text.Face
	Arrow      *widget.ButtonImageImage
}

type ListResources struct {
	Image        *widget.ScrollContainerImage
	Track        *widget.SliderTrackImage
	TrackPadding widget.Insets
	Handle       *widget.ButtonImage
	HandleSize   int
	FontFace     text.Face
	Entry        *widget.ListEntryColor
	EntryPadding widget.Insets
}

func NewAnchorContainer(opts ...widget.AnchorLayoutOpt) *widget.Container {
	return widget.NewContainer(
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(
				widget.AnchorLayoutData{
					StretchHorizontal: true,
				},
			),
		),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout(opts...)),
		// DebugContainerColor(colornames.Green),
	)
}

func NewCenteredAnchorContainer(padding int) *widget.Container {
	return widget.NewContainer(
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(
				widget.AnchorLayoutData{
					HorizontalPosition: widget.AnchorLayoutPositionCenter,
					VerticalPosition:   widget.AnchorLayoutPositionCenter,
				},
			),
		),
		widget.ContainerOpts.Layout(
			widget.NewAnchorLayout(
				widget.AnchorLayoutOpts.Padding(widget.NewInsetsSimple(padding)),
			),
		),
		// DebugContainerColor(colornames.Green),
	)
}

func NewPageContentContainer() *widget.Container {
	return widget.NewContainer(
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			StretchHorizontal: true,
		})),
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Spacing(10),
		)))
}

func NewGridContainer(columns int, opts ...widget.GridLayoutOpt) *widget.Container {
	containerOpts := []widget.GridLayoutOpt{
		widget.GridLayoutOpts.Columns(columns),
	}
	containerOpts = append(containerOpts, opts...)
	return widget.NewContainer(
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(
				widget.AnchorLayoutData{
					StretchHorizontal: true,
					StretchVertical:   true,
				},
			),
		),
		widget.ContainerOpts.Layout(widget.NewGridLayout(containerOpts...)),
		// DebugContainerColor(colornames.Blue),
	)
}

func NewHorizontalContainer() *widget.Container {
	return widget.NewContainer(
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(
				widget.AnchorLayoutData{
					StretchHorizontal: true,
				},
			),
		),
		widget.ContainerOpts.WidgetOpts(
			// instruct the container's anchor layout to center the button both horizontally and vertically
			widget.WidgetOpts.LayoutData(
				widget.AnchorLayoutData{
					HorizontalPosition: widget.AnchorLayoutPositionCenter,
					VerticalPosition:   widget.AnchorLayoutPositionCenter,
				},
			),
		),
		widget.ContainerOpts.Layout(
			widget.NewRowLayout(
				widget.RowLayoutOpts.Direction(widget.DirectionHorizontal),
				widget.RowLayoutOpts.Spacing(10),
				widget.RowLayoutOpts.Padding(widget.Insets{Left: 32, Right: 32, Top: 32}),
			),
		),
		// DebugContainerColor(colornames.Red),
	)
}

func DebugContainerColor(color color.Color) widget.ContainerOpt {
	return widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color))
}

func NewRowLayoutContainerWithMinWidth(minWidth, spacing int, rowscale []bool) *widget.Container {
	return widget.NewContainer(
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			StretchHorizontal: true,
			StretchVertical:   true,
		})),
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.MinSize(minWidth, 0)),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
			}),
		),
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(1),
			widget.GridLayoutOpts.Stretch([]bool{true}, rowscale),
			widget.GridLayoutOpts.Spacing(spacing, spacing),
		)),
	)
}

func NewRowLayoutContainer(spacing int, rowscale []bool) *widget.Container {
	return NewRowLayoutContainerWithMinWidth(0, spacing, rowscale)
}

func NewTransparentSeparator() widget.PreferredSizeLocateableWidget {
	c := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Padding(widget.Insets{
				Top:    4,
				Bottom: 4,
			}))),
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(
			widget.RowLayoutData{Stretch: true},
		)))

	c.AddChild(widget.NewGraphic(
		widget.GraphicOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
			Stretch:   true,
			MaxHeight: 1,
		})),
		widget.GraphicOpts.ImageNineSlice(image.NewNineSliceColor(color.RGBA{})),
	))

	return c
}

func NewSeparator(ld interface{}) widget.PreferredSizeLocateableWidget {
	c := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Padding(widget.Insets{
				Top:    20,
				Bottom: 20,
			}))),
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(ld)))

	c.AddChild(widget.NewGraphic(
		widget.GraphicOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
			Stretch:   true,
			MaxHeight: 2,
		})),
		widget.GraphicOpts.ImageNineSlice(image.NewNineSliceColor(RGB(0x2a3944))),
	))

	return c
}

func NewCenteredLabel(text string, ff text.Face) *widget.Text {
	return widget.NewText(
		widget.TextOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Stretch: true,
			}),
		),
		widget.TextOpts.Position(widget.TextPositionCenter, widget.TextPositionCenter),
		widget.TextOpts.Text(text, ff, NormalTextColor),
	)
}

func NewColoredLabel(text string, ff text.Face, clr color.Color, options ...widget.TextOpt) *widget.Text {
	opts := []widget.TextOpt{
		widget.TextOpts.Text(text, ff, clr),
	}
	if len(options) != 0 {
		opts = append(opts, options...)
	}
	return widget.NewText(opts...)
}

func NewLabel(text string, ff text.Face, options ...widget.TextOpt) *widget.Text {
	return NewColoredLabel(text, ff, NormalTextColor, options...)
}

type RecipeView struct {
	Container *widget.Container
	Icon1     *widget.Graphic
	Icon2     *widget.Graphic
	Separator *widget.Text
}

func NewRecipeView(res *Resources) *RecipeView {
	iconsContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionHorizontal),
			widget.RowLayoutOpts.Spacing(4),
			widget.RowLayoutOpts.Padding(widget.Insets{
				Top: 14,
			}),
		)),
	)

	icon1 := widget.NewGraphic()
	iconsContainer.AddChild(icon1)

	separator := widget.NewText(
		widget.TextOpts.Text("", res.Font2, res.Button.TextColors.Idle),
		widget.TextOpts.Position(widget.TextPositionCenter, widget.TextPositionCenter),
	)
	iconsContainer.AddChild(separator)

	icon2 := widget.NewGraphic()
	iconsContainer.AddChild(icon2)

	n := &RecipeView{
		Container: iconsContainer,
		Icon1:     icon1,
		Icon2:     icon2,
		Separator: separator,
	}
	return n
}

func (r *RecipeView) SetImages(a, b *ebiten.Image) {
	r.Icon1.Image = a
	r.Icon2.Image = b
	if a == nil && b == nil {
		r.Container.GetWidget().Visibility = widget.Visibility_Hide
		r.Separator.Label = ""
	} else {
		r.Container.GetWidget().Visibility = widget.Visibility_Show
		r.Separator.Label = "+"
	}
}

type ItemButton struct {
	Widget widget.PreferredSizeLocateableWidget
	Button *widget.Button
	label  *widget.Text
	state  bool
	res    *ToggleButtonResource
}

func (b *ItemButton) IsToggled() bool {
	return b.state
}

func (b *ItemButton) SetDisabled(disabled bool) {
	b.Button.GetWidget().Disabled = disabled
}

func (b *ItemButton) SetToggled(state bool) {
	if b.state == state {
		return
	}
	b.Toggle()
}

func (b *ItemButton) Toggle() {
	b.state = !b.state
	if b.state {
		b.Button.Image = b.res.AltImage
		if b.label != nil {
			b.label.Color = b.res.AltColor
		}
	} else {
		b.Button.Image = b.res.Image
		if b.label != nil {
			b.label.Color = b.res.Color
		}
	}
}

type ButtonConfig struct {
	Font text.Face

	Text string

	OnPressed func()
	OnHover   func()
}

func NewSmallButton(res *Resources, text string, onclick func()) *widget.Button {
	return NewButtonWithConfig(res, ButtonConfig{
		Font:      res.Font1,
		Text:      text,
		OnPressed: onclick,
	})
}

func NewButton(res *Resources, text string, onclick func()) *widget.Button {
	return NewButtonWithConfig(res, ButtonConfig{
		Font:      res.Font3,
		Text:      text,
		OnPressed: onclick,
	})
}

func NewButtonWithConfig(res *Resources, config ButtonConfig) *widget.Button {
	ff := config.Font
	if ff == nil {
		ff = res.Font2
	}

	options := []widget.ButtonOpt{
		widget.ButtonOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
			Stretch: true,
		})),
		widget.ButtonOpts.Image(res.Button.Image),
		widget.ButtonOpts.Text(config.Text, ff, res.Button.TextColors),
		widget.ButtonOpts.TextPadding(res.Button.Padding),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			if config.OnPressed != nil {
				config.OnPressed()
			}
		}),
	}

	if config.OnHover != nil {
		options = append(options, widget.ButtonOpts.CursorEnteredHandler(func(args *widget.ButtonHoverEventArgs) {
			config.OnHover()
		}))
	}

	b := widget.NewButton(options...)
	return b
}

func NewTextPanel(res *Resources, minWidth, minHeight int) *widget.Container {
	return NewDarkPanel(res, minWidth, minHeight,
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Spacing(4),
			widget.RowLayoutOpts.Padding(widget.Insets{
				Top:    16,
				Bottom: 16,
				Left:   20,
				Right:  20,
			}),
		)))
}

func NewDarkPanel(res *Resources, minWidth, minHeight int, opts ...widget.ContainerOpt) *widget.Container {
	return newPanel(res, minWidth, minHeight, true, opts...)
}

func NewPanel(res *Resources, minWidth, minHeight int, opts ...widget.ContainerOpt) *widget.Container {
	return newPanel(res, minWidth, minHeight, false, opts...)
}

func newPanel(res *Resources, minWidth, minHeight int, dark bool, opts ...widget.ContainerOpt) *widget.Container {
	panelRes := res.Panel
	if dark {
		panelRes = res.DarkPanel
	}
	options := []widget.ContainerOpt{
		widget.ContainerOpts.BackgroundImage(panelRes.Image),
		widget.ContainerOpts.Layout(
			widget.NewAnchorLayout(
				widget.AnchorLayoutOpts.Padding(widget.NewInsetsSimple(25)),
			),
		),
		widget.ContainerOpts.WidgetOpts(
			// instruct the container's anchor layout to center the button both horizontally and vertically
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				// StretchHorizontal:  true,
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
			}),
			widget.WidgetOpts.MinSize(minWidth, minHeight),
		),
	}
	options = append(options, opts...)

	return widget.NewContainer(options...)
}

func LoadResources(dst *Resources) *Resources {
	result := dst

	{
		result.Panel = &PanelResource{
			Image: NewNineSliceColor(colornames.Forestgreen),
			Padding: widget.Insets{
				Left:   16,
				Right:  16,
				Top:    10,
				Bottom: 10,
			},
		}
	}

	{
		result.DarkPanel = &PanelResource{
			Image: NewNineSliceColor(colornames.Darkgreen),
			Padding: widget.Insets{
				Left:   16,
				Right:  16,
				Top:    10,
				Bottom: 10,
			},
		}
	}

	{
		idle := NewNineSliceColor(colornames.Peru)
		hover := NewNineSliceColor(colornames.Sienna)
		pressed := NewNineSliceColor(colornames.Saddlebrown)
		pressedHover := NewNineSliceColor(colornames.Chocolate)
		disabled := NewNineSliceColor(colornames.Lightgray)
		buttonPadding := widget.Insets{
			Left:  30,
			Right: 30,
		}
		buttonColors := &widget.ButtonTextColor{
			Idle:     NormalTextColor,
			Disabled: RGB(0x5a7a91),
		}
		result.Button = &ButtonResource{
			Image: &widget.ButtonImage{
				Idle:         idle,
				Hover:        hover,
				Pressed:      pressed,
				PressedHover: pressedHover,
				Disabled:     disabled,
			},
			Padding:    buttonPadding,
			TextColors: buttonColors,
		}
	}

	result.Font1 = &text.GoTextFace{
		Source: assets.FontGoregular,
		Size:   80,
	}

	result.Font2 = &text.GoTextFace{
		Source: assets.FontGoregular,
		Size:   40,
	}

	result.Font3 = &text.GoTextFace{
		Source: assets.FontGoregular,
		Size:   20,
	}

	return result
}
