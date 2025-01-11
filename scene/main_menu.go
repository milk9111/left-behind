package scene

import (
	"image/color"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/milk9111/left-behind/assets"
	"github.com/milk9111/left-behind/component"
	"github.com/milk9111/left-behind/engine/ui"
)

type MainMenu struct {
	ui   *ebitenui.UI
	game *component.GameData

	nextScene Scene
}

func NewMainMenu(game *component.GameData) *MainMenu {
	m := &MainMenu{
		game:      game,
		nextScene: SceneMainMenu,
	}

	//This creates the root container for this UI.
	rootContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(assets.ColorCozyGreen)),
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.TrackHover(false)),
		widget.ContainerOpts.Layout(
			widget.NewGridLayout(
				// It is using a GridLayout with a single column
				widget.GridLayoutOpts.Columns(1),
				// It uses the Stretch parameter to define how the rows will be layed out.
				// - a fixed sized header
				// - a content row that stretches to fill all remaining space
				// - a fixed sized footer
				widget.GridLayoutOpts.Stretch([]bool{true}, []bool{false, true, false}),
				// Padding defines how much space to put around the outside of the grid.
				widget.GridLayoutOpts.Padding(widget.Insets{
					Top:    20,
					Bottom: 20,
				}),
				// Spacing defines how much space to put between each column and row
				widget.GridLayoutOpts.Spacing(0, 20),
			),
		),
	)

	headerContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(
			widget.NewRowLayout(
				widget.RowLayoutOpts.Direction(widget.DirectionVertical),
				widget.RowLayoutOpts.Spacing(15),
				widget.RowLayoutOpts.Padding(
					widget.NewInsetsSimple(100),
				),
			),
		),
	)

	titleLabelContainer := widget.NewContainer(
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Stretch: true,
			}),
			widget.WidgetOpts.TrackHover(false),
		),
		widget.ContainerOpts.Layout(
			widget.NewAnchorLayout(
				widget.AnchorLayoutOpts.Padding(widget.Insets{
					Right:  25,
					Left:   25,
					Top:    4,
					Bottom: 4,
				}),
			),
		),
	)

	titleLabel := widget.NewText(
		widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			HorizontalPosition: widget.AnchorLayoutPositionCenter,
			VerticalPosition:   widget.AnchorLayoutPositionStart,
		})),
		widget.TextOpts.Text(
			"Trixie the Truffler",
			&text.GoTextFace{
				Source: assets.FontRoboto,
				Size:   80,
			},
			color.White,
		),
		widget.TextOpts.Position(widget.TextPositionCenter, widget.TextPositionCenter),
	)

	titleLabelContainer.AddChild(titleLabel)

	headerContainer.AddChild(titleLabelContainer)

	rootContainer.AddChild(headerContainer)

	buttonContainer := widget.NewContainer(
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Stretch: true,
			}),
		),
		widget.ContainerOpts.Layout(
			widget.NewAnchorLayout(),
		),
	)

	res := &ui.Resources{}
	res = ui.LoadResources(res)

	button := ui.NewButton(
		res,
		"Start Game",
		func() {
			m.nextScene = SceneGame
		},
	)

	buttonContainer.AddChild(button)

	rootContainer.AddChild(buttonContainer)

	m.ui = &ebitenui.UI{
		Container: rootContainer,
	}

	return m
}

func (m *MainMenu) Init() {
	m.nextScene = SceneMainMenu
}

func (m *MainMenu) Update() Scene {
	m.ui.Update()

	return m.nextScene
}

func (m *MainMenu) Draw(screen *ebiten.Image) {
	m.ui.Draw(screen)
}
