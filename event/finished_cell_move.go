package event

import (
	"github.com/yohamta/donburi"
	devents "github.com/yohamta/donburi/features/events"
)

type FinishedCellMoveData struct {
	Entry *donburi.Entry
}

var FinishedCellMove = devents.NewEventType[FinishedCellMoveData]()
