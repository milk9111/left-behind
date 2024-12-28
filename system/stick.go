package system

import (
	"time"

	"github.com/milk9111/left-behind/assets/scripts"
	"github.com/milk9111/left-behind/component"
	"github.com/milk9111/left-behind/engine/tween"
	"github.com/yohamta/donburi"
	dmath "github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
)

type Stick struct {
	query *donburi.Query

	gridPos *dmath.Vec2
}

func NewStick() *Stick {
	return &Stick{
		query: donburi.NewQuery(filter.Contains(transform.Transform, component.Sticky)),
	}
}

func (s *Stick) Update(w donburi.World) {
	if s.gridPos == nil {
		e := scripts.MustFindEntry(w, scripts.GridComponent)
		s.gridPos = &transform.Transform.Get(e).LocalPosition
	}

	s.query.Each(w, func(e *donburi.Entry) {
		t := transform.Transform.Get(e)
		sticky := component.Sticky.Get(e)

		if sticky.Disabled || sticky.QueuedPosition == nil {
			return
		}

		if sticky.QueuedPosition != nil && sticky.Tween == nil {
			sticky.Tween = tween.NewVec2(1000*time.Millisecond, sticky.Position, *sticky.QueuedPosition, tween.EaseInOutCubic)
		}

		sticky.Position = sticky.Tween.Update()
		t.LocalPosition = (*s.gridPos).Add(sticky.Position)
		if sticky.Tween.Done() {
			sticky.Position = *sticky.QueuedPosition
			t.LocalPosition = (*s.gridPos).Add(sticky.Position)
			sticky.QueuedPosition = nil
			sticky.Tween = nil
			return
		}
	})
}
