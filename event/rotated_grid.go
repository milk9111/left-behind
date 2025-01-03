package event

import devents "github.com/yohamta/donburi/features/events"

type RotatedGridData struct{}

var RotatedGrid = devents.NewEventType[RotatedGridData]()
