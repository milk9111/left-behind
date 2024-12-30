package event

import devents "github.com/yohamta/donburi/features/events"

type StartedStickyTranslationData struct {
	IsRotatingBehind bool
}

var StartedStickyTranslation = devents.NewEventType[StartedStickyTranslationData]()
