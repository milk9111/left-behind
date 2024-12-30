package system

// import (
// 	"time"

// 	"github.com/milk9111/left-behind/assets/scripts"
// 	"github.com/milk9111/left-behind/component"
// 	"github.com/milk9111/left-behind/engine/tween"
// 	"github.com/yohamta/donburi"
// 	dmath "github.com/yohamta/donburi/features/math"
// 	"github.com/yohamta/donburi/features/transform"
// 	"github.com/yohamta/donburi/filter"
// )

// type Sticky struct {
// 	query *donburi.Query

// 	gridPos *dmath.Vec2
// }

// func NewSticky() *Sticky {
// 	return &Sticky{
// 		query: donburi.NewQuery(filter.Contains(transform.Transform, component.Cell)),
// 	}
// }

// func (s *Sticky) Update(w donburi.World) {
// 	if s.gridPos == nil {
// 		e := scripts.MustFindEntry(w, scripts.GridComponent)
// 		s.gridPos = &transform.Transform.Get(e).LocalPosition
// 	}

// 	s.query.Each(w, func(e *donburi.Entry) {
// 		t := transform.Transform.Get(e)
// 		cell := component.Cell.Get(e)

// 		if !cell.IsSticky || cell.QueuedPosition == nil {
// 			return
// 		}

// 		vec2Tween := tween.NewVec2(
// 			1000*time.Millisecond,
// 			cell.Position,
// 			*cell.QueuedPosition,
// 			tween.EaseInOutCubic,
// 			tween.WithVec2UpdateCallback(func(v dmath.Vec2) {
// 				cell.Position = v
// 				t.LocalPosition = (*s.gridPos).Add(cell.Position)
// 			}),
// 		)

// 		vec2Tween.FinishedEvent.Subscribe(w, func(_ donburi.World, _ tween.FinishedEventData) {
// 			cell.QueuedPosition = nil
// 		})

// 		cell.Position = cell.Tween.Update()
// 		t.LocalPosition = (*s.gridPos).Add(cell.Position)
// 		if cell.Tween.Done() {
// 			cell.Position = *cell.QueuedPosition
// 			t.LocalPosition = (*s.gridPos).Add(cell.Position)
// 			cell.QueuedPosition = nil
// 			cell.Tween = nil
// 			return
// 		}
// 	})
// }
