package component

import "github.com/yohamta/donburi"

type InputEventType int

const (
	InputEventTypeUnknown InputEventType = iota
	InputEventTypeMoveLeft
	InputEventTypeMoveBehind
	InputEventTypeRotateLeft
	InputEventTypeRotateBehind
)

type inputHandler interface {
	OnInput(InputEventType)
}

type InputHandlerData struct {
	Handler inputHandler
}

var InputHandler = donburi.NewComponentType[InputHandlerData]()
