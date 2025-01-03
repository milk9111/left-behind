package event

import devents "github.com/yohamta/donburi/features/events"

type UnpausedGameData struct{}

var UnpausedGame = devents.NewEventType[UnpausedGameData]()
