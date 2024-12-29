package event

import devents "github.com/yohamta/donburi/features/events"

type ReachedGoalData struct{}

var ReachedGoal = devents.NewEventType[ReachedGoalData]()
