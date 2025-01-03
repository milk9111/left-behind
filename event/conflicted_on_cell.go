package event

import (
	"github.com/yohamta/donburi"
	devents "github.com/yohamta/donburi/features/events"
)

type ConflictedOnCellData struct {
	Entry *donburi.Entry
}

var ConflictedOnCell = devents.NewEventType[ConflictedOnCellData]()
