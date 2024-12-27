package scripts

import (
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
