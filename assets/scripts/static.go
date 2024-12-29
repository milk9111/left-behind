package scripts

import (
	"github.com/milk9111/left-behind/component"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
)

// Static is a hacky way to sync up the local position with the Grid. This should really be its own component instead.
type Static struct {
	e *donburi.Entry
	t *transform.TransformData
}

func NewStatic(e *donburi.Entry) *Static {
	return &Static{
		e: e,
		t: transform.Transform.Get(e),
	}
}

func (s *Static) Start(w donburi.World) {
	s.t.LocalPosition = transform.Transform.Get(MustFindEntry(w, GridComponent)).LocalPosition.Add(component.Cell.Get(s.e).Position)
}

var StaticComponent = donburi.NewComponentType[Static]()
