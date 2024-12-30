package archetype

import (
	"image/color"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/milk9111/left-behind/assets"
	"github.com/milk9111/left-behind/component"
	"github.com/yohamta/donburi"
)

func NewUI(w donburi.World, levelName string) *donburi.Entry {
	e := w.Entry(w.Create(component.UI))

	rootContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(
			widget.NewAnchorLayout(
				widget.AnchorLayoutOpts.Padding(
					widget.NewInsetsSimple(5),
				),
			),
		),
	)

	headerContainer := widget.NewContainer(
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				VerticalPosition:   widget.AnchorLayoutPositionStart,
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				StretchHorizontal:  true,
				StretchVertical:    false,
			}),
		),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout(
			widget.AnchorLayoutOpts.Padding(
				widget.NewInsetsSimple(20),
			),
		)),
	)

	label := widget.NewText(
		widget.TextOpts.Text(
			levelName,
			&text.GoTextFace{
				Source: assets.FontGoregular,
				Size:   30,
			},
			color.White,
		),
		widget.TextOpts.Position(widget.TextPositionCenter, widget.TextPositionCenter),
		widget.TextOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionStart,
				StretchHorizontal:  false,
				StretchVertical:    false,
			}),
		),
	)

	headerContainer.AddChild(label)
	rootContainer.AddChild(headerContainer)

	component.UI.SetValue(e, component.UIData{
		Container: &ebitenui.UI{
			Container: rootContainer,
		},
	})

	return e
}
