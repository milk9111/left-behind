package event

import devents "github.com/yohamta/donburi/features/events"

type FinishedPlayerMoveData struct{}

var FinishedPlayerMove = devents.NewEventType[FinishedPlayerMoveData]()
