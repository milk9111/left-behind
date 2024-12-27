package system

import (
	"time"

	"github.com/milk9111/left-behind/component"
	"github.com/milk9111/left-behind/engine/tween"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
)

type Stick struct {
	query *donburi.Query
}

func NewStick() *Stick {
	return &Stick{
		query: donburi.NewQuery(filter.Contains(transform.Transform, component.Sticky)),
	}
}

func (s *Stick) Update(w donburi.World) {
	s.query.Each(w, func(e *donburi.Entry) {
		t := transform.Transform.Get(e)
		sticky := component.Sticky.Get(e)

		if sticky.Disabled || sticky.QueuedPosition == nil {
			return
		}

		if sticky.QueuedPosition != nil && sticky.Tween == nil {
			sticky.Tween = tween.NewVec2(1000*time.Millisecond, t.LocalPosition, *sticky.QueuedPosition, tween.EaseInOutCubic)
		}

		t.LocalPosition = sticky.Tween.Update()
		if sticky.Tween.Done() {
			t.LocalPosition = *sticky.QueuedPosition
			sticky.QueuedPosition = nil
			sticky.Tween = nil
			return
		}
	})
}
