package archetype

import (
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/milk9111/left-behind/assets/scripts"
	"github.com/milk9111/left-behind/component"
	"github.com/milk9111/left-behind/engine/ui"
	"github.com/yohamta/donburi"
)

func NewUI(w donburi.World, game *component.GameData, levelName string) *donburi.Entry {
	e := w.Entry(w.Create(component.UI, component.Start, component.Update))

	res := ui.LoadResources(&ui.Resources{})

	rootContainer := ui.NewRowLayoutContainer(25, []bool{false, true})

	rowContainer := ui.NewCenteredAnchorContainer(0)
	rootContainer.AddChild(rowContainer)

	levelLabelContainer := ui.NewCenteredAnchorContainer(20)
	rowContainer.AddChild(levelLabelContainer)

	levelLabel := ui.NewCenteredLabel(
		levelName,
		res.Font2,
	)
	levelLabelContainer.AddChild(levelLabel)

	pauseMenuRootContainer := ui.NewCenteredAnchorContainer(20)
	rootContainer.AddChild(pauseMenuRootContainer)
	pauseMenuContainer := ui.NewDarkPanel(
		res,
		20,
		40,
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
			}),
			widget.WidgetOpts.MinSize(20, 40),
		),
	)
	pauseMenuRootContainer.AddChild(pauseMenuContainer)

	pauseMenuRowContainer := ui.NewRowLayoutContainer(10, []bool{false, true, false})
	pauseMenuContainer.AddChild(pauseMenuRowContainer)

	pauseMenuLabel := ui.NewCenteredLabel(
		"Paused",
		res.Font2,
	)
	pauseMenuRowContainer.AddChild(pauseMenuLabel)

	pauseMenuRowContainer.AddChild(
		ui.NewSeparator(
			widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
			},
		),
	)

	pauseMenuScript := scripts.NewPauseMenu(pauseMenuRootContainer.GetWidget())

	pauseMenuExitButton := ui.NewButton(
		res,
		"Exit",
		func() {
			pauseMenuScript.OnClick(w)
		},
	)

	pauseMenuRowContainer.AddChild(pauseMenuExitButton)

	rootUI := &ebitenui.UI{
		Container: rootContainer,
	}

	component.UI.SetValue(e, component.UIData{
		Container: rootUI,
	})

	component.Start.SetValue(e, component.StartData{
		Handler: pauseMenuScript,
	})

	component.Update.SetValue(e, component.UpdateData{
		Handler: pauseMenuScript,
	})

	return e
}

/*
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
				widget.GridLayoutOpts.Stretch([]bool{true}, []bool{false, false}),
				// Padding defines how much space to put around the outside of the grid.
				widget.GridLayoutOpts.Padding(widget.Insets{
					Top:    20,
					Bottom: 60,
					Left:   game.WorldWidth / 2,
					Right:  game.WorldWidth / 2,
				}),
				// Spacing defines how much space to put between each column and row
				widget.GridLayoutOpts.Spacing(0, 20),
			),
		),
	)

	headerContainer := widget.NewContainer(
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				VerticalPosition:   widget.AnchorLayoutPositionStart,
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				StretchHorizontal:  true,
				StretchVertical:    true,
			}),
		),
		widget.ContainerOpts.Layout(
			widget.NewAnchorLayout(
				widget.AnchorLayoutOpts.Padding(
					widget.NewInsetsSimple(20),
				),
			),
		),
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
				StretchHorizontal:  true,
				StretchVertical:    true,
			}),
		),
	)

	headerContainer.AddChild(label)
	rootContainer.AddChild(headerContainer)

	pauseMenuContainer := widget.NewContainer(
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(
				widget.AnchorLayoutData{
					HorizontalPosition: widget.AnchorLayoutPositionCenter,
					VerticalPosition:   widget.AnchorLayoutPositionCenter,
					StretchHorizontal:  false,
					StretchVertical:    false,
				},
			),
		),
		widget.ContainerOpts.BackgroundImage(ui.NewNineSliceColor(colornames.Peru)),
		widget.ContainerOpts.Layout(
			widget.NewAnchorLayout(
				widget.AnchorLayoutOpts.Padding(
					widget.NewInsetsSimple(20),
				),
			),
		),
	)

	pauseMenuLabel := widget.NewText(
		widget.TextOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(
				widget.RowLayoutData{
					Position: widget.RowLayoutPositionCenter,
				},
			),
		),
		widget.TextOpts.Text(
			"Paused",
			&text.GoTextFace{
				Source: assets.FontGoregular,
				Size:   40,
			},
			colornames.White,
		),
		widget.TextOpts.Position(
			widget.TextPositionCenter,
			widget.TextPositionCenter,
		),
	)

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
			"Main Menu",
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
			fmt.Println("pressed main menu button")
		}),
	)

	buttonContainer.AddChild(button)

	pauseMenuContainer.AddChild(pauseMenuLabel)
	pauseMenuContainer.AddChild(buttonContainer)

	rootContainer.AddChild(pauseMenuContainer)
*/
