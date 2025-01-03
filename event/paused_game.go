package event

import devents "github.com/yohamta/donburi/features/events"

type PausedGameData struct{}

var PausedGame = devents.NewEventType[PausedGameData]()
