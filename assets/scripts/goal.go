package scripts

import (
	"github.com/milk9111/left-behind/component"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
)

type Goal struct {
	e *donburi.Entry
	t *transform.TransformData
}

func NewGoal(e *donburi.Entry) *Goal {
	return &Goal{
		e: e,
		t: transform.Transform.Get(e),
	}
}

func (g *Goal) Start(w donburi.World) {
	g.t.LocalPosition = transform.Transform.Get(MustFindEntry(w, GridComponent)).LocalPosition.Add(component.Sticky.Get(g.e).Position)
}

var GoalComponent = donburi.NewComponentType[Goal]()
