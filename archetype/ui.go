package archetype

import (
	"fmt"
	"image/color"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/milk9111/left-behind/assets"
	"github.com/milk9111/left-behind/assets/scripts"
	"github.com/milk9111/left-behind/component"
	"github.com/milk9111/left-behind/engine/ui"
	"github.com/yohamta/donburi"
)

func NewUI(w donburi.World, game *component.GameData, levelIndex int) *donburi.Entry {
	e := w.Entry(w.Create(component.UI, component.Start, component.Update))

	res := ui.LoadResources(&ui.Resources{})

	rootContainer := ui.NewRowLayoutContainer(25, []bool{false, true})

	levelLabelContainer := ui.NewTopCenteredAnchorContainer(10)
	rootContainer.AddChild(levelLabelContainer)

	levelLabel := ui.NewColoredLabel(
		fmt.Sprintf("%d / %d", levelIndex, len(assets.Levels)),
		res.Font2,
		color.White,
		widget.TextOpts.Position(widget.TextPositionStart, widget.TextPositionStart),
	)
	levelLabelContainer.AddChild(levelLabel)

	pauseMenuRootContainer := ui.NewCenteredAnchorContainer(50)
	rootContainer.AddChild(pauseMenuRootContainer)
	pauseMenuContainer := ui.NewDarkPanel(
		res,
		20,
		40,
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionStart,
			}),
			widget.WidgetOpts.MinSize(300, 400),
		),
	)
	pauseMenuRootContainer.AddChild(pauseMenuContainer)

	pauseMenuRowContainer := ui.NewRowLayoutContainer(10, []bool{false, false, true})
	// pauseMenuRowContainer.BackgroundImage = ui.NewNineSliceColor(colornames.Red)
	pauseMenuContainer.AddChild(pauseMenuRowContainer)

	pauseMenuLabel := ui.NewTopCenteredLabel(
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

	pauseMenuButtonRowContainer := ui.NewRowLayoutContainer(10, []bool{false})
	pauseMenuRowContainer.AddChild(pauseMenuButtonRowContainer)

	pauseMenuExitButton := ui.NewButton(
		res,
		"Exit",
		func() {
			pauseMenuScript.OnClick(w)
		},
	)
	pauseMenuButtonRowContainer.AddChild(pauseMenuExitButton)

	controlsText := ui.NewCenteredAnchorContainer(0)
	controlsText.BackgroundImage = ui.NewNineSliceImage(assets.SpriteControlsText, 0, 0)
	pauseMenuButtonRowContainer.AddChild(controlsText)

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
