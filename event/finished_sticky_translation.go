package event

import devents "github.com/yohamta/donburi/features/events"

type FinishedStickyTranslationData struct{}

var FinishedStickyTranslation = devents.NewEventType[FinishedStickyTranslationData]()
