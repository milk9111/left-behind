package scripts

import (
	"fmt"
	"time"

	"github.com/milk9111/left-behind/component"
	"github.com/milk9111/left-behind/engine/tween"
	"github.com/yohamta/donburi"
	dmath "github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

type Player struct {
	e      *donburi.Entry
	t      *transform.TransformData
	sticky *component.StickyData

	grid *Grid
	goal *transform.TransformData

	tween         *tween.Vec2
	inputDisabled bool // I might be better off disabling the system instead
	goalReached   bool
}

func NewPlayer(e *donburi.Entry) *Player {
	return &Player{
		e:      e,
		t:      transform.Transform.Get(e),
		sticky: component.Sticky.Get(e),
	}
}

func (p *Player) Start(w donburi.World) {
	p.grid = GridComponent.Get(MustFindEntry(w, GridComponent))
	p.goal = transform.Transform.Get(MustFindEntry(w, component.TagGoal))

	p.t.LocalPosition = p.grid.t.LocalPosition.Add(p.sticky.Position)
}

func (p *Player) Update() {
	if !p.goalReached && p.HasReachedGoal() {
		p.goalReached = true
		p.GoalReached()
	}

	if p.tween == nil {
		return
	}

	p.sticky.Position = p.tween.Update()
	p.t.LocalPosition = p.grid.t.LocalPosition.Add(p.sticky.Position)
	if p.tween.Done() {
		p.sticky.Position = p.tween.End()
		p.t.LocalPosition = p.grid.t.LocalPosition.Add(p.sticky.Position)
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
		nextPos = p.sticky.Position.Add(dmath.NewVec2(-32, 0))
		if !p.grid.CanMove(nextPos) {
			fmt.Println("can't move left")
			return
		}
	} else if inputEventType == component.InputEventTypeMoveBehind {
		nextPos = p.sticky.Position.Add(dmath.NewVec2(0, 32))
		if !p.grid.CanMove(nextPos) {
			fmt.Println("can't move behind")
			return
		}
	} else {
		return // exit early because it's not input player cares about
	}

	p.tween = tween.NewVec2(250*time.Millisecond, p.sticky.Position, nextPos, tween.EaseInOutCubic)
	p.grid.Move(p.sticky.Position, nextPos)
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
