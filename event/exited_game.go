package event

import devents "github.com/yohamta/donburi/features/events"

type ExitedGameData struct{}

var ExitedGame = devents.NewEventType[ExitedGameData]()
