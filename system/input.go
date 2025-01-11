package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/milk9111/left-behind/component"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type Input struct {
	query    *donburi.Query
	disabled bool

	restartLevelCallback func()
}

func NewInput(restartLevelCallback func()) *Input {
	return &Input{
		query: donburi.NewQuery(filter.Contains(
			component.InputHandler,
		)),
		restartLevelCallback: restartLevelCallback,
	}
}

func (i *Input) Update(w donburi.World) {
	if i.disabled {
		return
	}

	var inputEventType component.InputEventType
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		inputEventType = component.InputEventTypeMoveLeft
	} else if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		inputEventType = component.InputEventTypeMoveBehind
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		inputEventType = component.InputEventTypeRotateLeft
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		inputEventType = component.InputEventTypeRotateBehind
		// } else if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		// 	i.restartLevelCallback()
	} else {
		return
	}

	i.query.Each(w, func(e *donburi.Entry) {
		component.InputHandler.Get(e).Handler.OnInput(w, inputEventType)
	})
}

func (i *Input) SetDisabled(disabled bool) {
	i.disabled = disabled
}
