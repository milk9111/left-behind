package event

import devents "github.com/yohamta/donburi/features/events"

type StartedPlayerMoveData struct{}

var StartedPlayerMove = devents.NewEventType[StartedPlayerMoveData]()
