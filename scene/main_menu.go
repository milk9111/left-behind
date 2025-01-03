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
	"golang.org/x/image/colornames"
)

type MainMenu struct {
	ui   *ebitenui.UI
	game *component.GameData

	nextScene Scene

	backgroundImage *ebiten.Image
	time            int
}

func NewMainMenu(game *component.GameData) *MainMenu {
	m := &MainMenu{
		game:      game,
		nextScene: SceneMainMenu,
	}

	backgroundImage := ebiten.NewImage(game.WorldWidth, game.WorldHeight)
	widthPiece := game.WorldWidth / 3
	heightPiece := game.WorldHeight / 3

	//This creates the root container for this UI.
	rootContainer := widget.NewContainer(
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
		widget.ContainerOpts.BackgroundImage(
			image.NewNineSlice(backgroundImage,
				[3]int{widthPiece, widthPiece, widthPiece},
				[3]int{heightPiece, heightPiece, heightPiece},
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
			"Left Behind",
			&text.GoTextFace{
				Source: assets.FontGoregular,
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
			widget.WidgetOpts.TrackHover(false),
		),
		widget.ContainerOpts.Layout(
			widget.NewAnchorLayout(),
		),
	)

	buttonImg := &widget.ButtonImage{
		Idle:         ui.NewNineSliceColor(colornames.Peru),
		Hover:        ui.NewNineSliceColor(colornames.Sienna),
		Pressed:      ui.NewNineSliceColor(colornames.Saddlebrown),
		PressedHover: ui.NewNineSliceColor(colornames.Chocolate),
		Disabled:     ui.NewNineSliceColor(colornames.Lightgray),
	}

	buttonTextColor := &widget.ButtonTextColor{
		Idle:     colornames.White,
		Hover:    colornames.White,
		Pressed:  colornames.White,
		Disabled: colornames.Black,
	}

	button := widget.NewButton(
		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(
				widget.AnchorLayoutData{
					HorizontalPosition: widget.AnchorLayoutPositionCenter,
					VerticalPosition:   widget.AnchorLayoutPositionStart,
					StretchHorizontal:  false,
					StretchVertical:    false,
				},
			),
		),
		widget.ButtonOpts.Image(buttonImg),
		widget.ButtonOpts.Text(
			"Trixie the Truffler",
			&text.GoTextFace{
				Source: assets.FontGoregular,
				Size:   20,
			},
			buttonTextColor,
		),
		widget.ButtonOpts.TextPadding(widget.Insets{
			Left:   25,
			Right:  25,
			Top:    10,
			Bottom: 10,
		}),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			m.nextScene = SceneGame
		}),
	)

	buttonContainer.AddChild(button)

	rootContainer.AddChild(buttonContainer)

	m.backgroundImage = backgroundImage
	m.ui = &ebitenui.UI{
		Container: rootContainer,
	}

	return m
}

func (m *MainMenu) Init() {
	m.nextScene = SceneMainMenu
}

func (m *MainMenu) Update() Scene {
	m.time++
	m.ui.Update()

	return m.nextScene
}

func (m *MainMenu) Draw(screen *ebiten.Image) {
	m.backgroundImage.Clear()

	rsop := &ebiten.DrawRectShaderOptions{}
	rsop.Uniforms = map[string]any{
		"Time":       float32(m.time) / 60,
		"Resolution": []float32{float32(m.backgroundImage.Bounds().Dx()), float32(m.backgroundImage.Bounds().Dy())},
	}
	m.backgroundImage.DrawRectShader(m.backgroundImage.Bounds().Dx(), m.backgroundImage.Bounds().Dy(), assets.ShaderRainbowFlows, rsop)

	m.ui.Draw(screen)
}
