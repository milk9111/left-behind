package scripts

import (
	"github.com/milk9111/left-behind/component"
	"github.com/milk9111/left-behind/event"
	"github.com/yohamta/donburi"
)

type Reverse struct {
	e *donburi.Entry

	cell *component.CellData
}

func NewReverse(e *donburi.Entry) *Reverse {
	return &Reverse{
		e:    e,
		cell: component.Cell.Get(e),
	}
}

func (r *Reverse) OnFinishedCellMove(_ donburi.World, evt event.FinishedCellMoveData) {
	cell := component.Cell.Get(evt.Entry)
	if !cell.Position.Equal(r.cell.Position) {
		return
	}

	sticky := component.Sticky.Get(evt.Entry)
	sticky.Disabled = !sticky.Disabled
}
