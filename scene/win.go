package scene

import (
	"image/color"
	"time"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/milk9111/left-behind/assets"
	"github.com/milk9111/left-behind/component"
	"github.com/milk9111/left-behind/engine"
	"golang.org/x/image/colornames"
)

type Win struct {
	ui   *ebitenui.UI
	game *component.GameData

	timer *engine.Timer
}

func NewWin(game *component.GameData) *Win {
	rootContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(assets.ColorCozyGreen)),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout(
			widget.AnchorLayoutOpts.Padding(widget.NewInsetsSimple(30)),
		)),
	)

	rowContainer := widget.NewContainer(
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
				StretchHorizontal:  true,
				StretchVertical:    false,
			}),
			widget.WidgetOpts.MinSize(100, 100),
		),
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Padding(widget.NewInsetsSimple(25)),
			widget.RowLayoutOpts.Spacing(10),
		)),
	)
	rootContainer.AddChild(rowContainer)

	headerLabel := widget.NewText(
		widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			HorizontalPosition: widget.AnchorLayoutPositionCenter,
			VerticalPosition:   widget.AnchorLayoutPositionCenter,
			StretchHorizontal:  false,
			StretchVertical:    false,
		})),
		widget.TextOpts.Text(
			"You Win!",
			&text.GoTextFace{
				Source: assets.FontRoboto,
				Size:   80,
			},
			color.White,
		),
		widget.TextOpts.Position(widget.TextPositionCenter, widget.TextPositionCenter),
	)

	rowContainer.AddChild(headerLabel)

	subheaderLabel := widget.NewText(
		widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			HorizontalPosition: widget.AnchorLayoutPositionCenter,
			VerticalPosition:   widget.AnchorLayoutPositionStart,
		})),
		widget.TextOpts.Text(
			"Thanks for playing! Taking you back to the Main Menu now.",
			&text.GoTextFace{
				Source: assets.FontRoboto,
				Size:   20,
			},
			colornames.Lightgray,
		),
		widget.TextOpts.Position(widget.TextPositionCenter, widget.TextPositionCenter),
	)

	rowContainer.AddChild(subheaderLabel)

	return &Win{
		game:  game,
		timer: engine.NewTimer(5 * time.Second),
		ui: &ebitenui.UI{
			Container: rootContainer,
		},
	}
}

func (w *Win) Init() {
	w.timer.Reset()
}

func (w *Win) Update() Scene {
	w.timer.Update()
	if w.timer.IsReady() {
		return SceneMainMenu
	}

	w.ui.Update()

	return SceneWin
}

func (w *Win) Draw(screen *ebiten.Image) {
	w.ui.Draw(screen)
}
