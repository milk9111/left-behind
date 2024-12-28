package scripts

import (
	"fmt"
	"time"

	"github.com/milk9111/left-behind/assets"
	"github.com/milk9111/left-behind/component"
	"github.com/milk9111/left-behind/engine/tween"
	"github.com/yohamta/donburi"
	dmath "github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

type Player struct {
	e          *donburi.Entry
	t          *transform.TransformData
	cell       *component.CellData
	audioQueue *component.AudioQueueData

	grid *Grid
	goal *transform.TransformData

	tween         *tween.Vec2
	inputDisabled bool // I might be better off disabling the system instead
	goalReached   bool
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

	p.t.LocalPosition = p.grid.t.LocalPosition.Add(p.cell.Position)
}

func (p *Player) Update() {
	if !p.goalReached && p.HasReachedGoal() {
		p.goalReached = true
		p.GoalReached()
	}

	if p.tween == nil {
		return
	}

	p.cell.Position = p.tween.Update()
	p.t.LocalPosition = p.grid.t.LocalPosition.Add(p.cell.Position)
	if p.tween.Done() {
		p.cell.Position = p.tween.End()
		p.t.LocalPosition = p.grid.t.LocalPosition.Add(p.cell.Position)
		p.grid.inputDisabled = false
		p.inputDisabled = false
		p.tween = nil
	}
}

func (p *Player) OnInput(inputEventType component.InputEventType) {
	if (p.tween != nil && !p.tween.Done()) || p.inputDisabled {
		return
	}

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

	p.tween = tween.NewVec2(250*time.Millisecond, p.cell.Position, nextPos, tween.EaseInOutCubic)
	p.grid.Move(p.cell.Position, nextPos)
	p.grid.inputDisabled = true
	p.inputDisabled = true
}

func (p *Player) HasReachedGoal() bool {
	return p.t.LocalPosition.Equal(p.goal.LocalPosition)
}

func (p *Player) GoalReached() {
	fmt.Println("REACHED GOAL!!!")
}

var PlayerComponent = donburi.NewComponentType[Player]()
