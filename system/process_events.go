package system

import (
	"github.com/yohamta/donburi"
	devents "github.com/yohamta/donburi/features/events"
)

type ProcessEvents struct{}

func NewProcessEvents() *ProcessEvents {
	return &ProcessEvents{}
}

func (p *ProcessEvents) Update(w donburi.World) {
	devents.ProcessAllEvents(w)
}
