package event

import devents "github.com/yohamta/donburi/features/events"

type FinishedLevelTransitionData struct{}

var FinishedLevelTransition = devents.NewEventType[FinishedLevelTransitionData]()
