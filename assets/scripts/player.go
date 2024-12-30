package scripts

import (
	"time"

	"github.com/milk9111/left-behind/assets"
	"github.com/milk9111/left-behind/component"
	"github.com/milk9111/left-behind/engine/tween"
	"github.com/milk9111/left-behind/event"
	"github.com/yohamta/donburi"
	dmath "github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

type Player struct {
	e          *donburi.Entry
	t          *transform.TransformData
	cell       *component.CellData
	audioQueue *component.AudioQueueData

	grid                 *Grid
	goal                 *transform.TransformData
	globalTweenVec2Queue *component.TweenVec2QueueData

	goalReached bool
}

func NewPlayer(e *donburi.Entry) *Player {
	return &Player{
		e:          e,
		t:          transform.Transform.Get(e),
		cell:       component.Cell.Get(e),
		audioQueue: component.AudioQueue.Get(e),
	}
}

func (p *Player) Start(w donburi.World) {
	p.grid = GridComponent.Get(MustFindEntry(w, GridComponent))
	p.goal = transform.Transform.Get(MustFindEntry(w, component.TagGoal))
	p.globalTweenVec2Queue = component.MustFindTweenVec2Queue(w)

	p.t.LocalPosition = p.grid.t.LocalPosition.Add(p.cell.Position)
}

func (p *Player) Update(w donburi.World) {
	if !p.goalReached && p.HasReachedGoal() {
		p.goalReached = true
		event.ReachedGoal.Publish(w, event.ReachedGoalData{})
	}
}

func (p *Player) OnInput(w donburi.World, inputEventType component.InputEventType) {
	var nextPos dmath.Vec2
	if inputEventType == component.InputEventTypeMoveLeft {
		nextPos = p.cell.Position.Add(dmath.NewVec2(-32, 0))
		if !p.grid.CanMove(nextPos) {
			p.audioQueue.Enqueue(assets.SFXBadMove)
			return
		}
	} else if inputEventType == component.InputEventTypeMoveBehind {
		nextPos = p.cell.Position.Add(dmath.NewVec2(0, 32))
		if !p.grid.CanMove(nextPos) {
			p.audioQueue.Enqueue(assets.SFXBadMove)
			return
		}
	} else {
		return // exit early because it's not input player cares about
	}

	vec2Tween := tween.NewVec2(
		w,
		250*time.Millisecond,
		p.cell.Position,
		nextPos,
		tween.EaseInOutCubic,
		tween.WithVec2UpdateCallback(func(t dmath.Vec2) {
			p.cell.Position = t
			p.t.LocalPosition = p.grid.t.LocalPosition.Add(p.cell.Position)
		}),
	)

	vec2Tween.FinishedEvent.Subscribe(w, func(w donburi.World, _ tween.FinishedEventData) {
		event.FinishedPlayerMove.Publish(w, event.FinishedPlayerMoveData{})
	})

	p.globalTweenVec2Queue.Enqueue(vec2Tween)
	p.grid.Move(p.cell.Position, nextPos)
	event.StartedPlayerMove.Publish(w, event.StartedPlayerMoveData{})
}

func (p *Player) HasReachedGoal() bool {
	return p.t.LocalPosition.Equal(p.goal.LocalPosition)
}

func (p *Player) OnReachedGoal(_ donburi.World, _ event.ReachedGoalData) {
	p.audioQueue.Enqueue(assets.SFXGoalReached)
}

var PlayerComponent = donburi.NewComponentType[Player]()
