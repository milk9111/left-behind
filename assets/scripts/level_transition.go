package scripts

import (
	"time"

	"github.com/milk9111/left-behind/engine"
	"github.com/milk9111/left-behind/event"
	"github.com/yohamta/donburi"
)

type LevelTransition struct {
	timer     *engine.Timer
	isStarted bool
}

func NewLevelTransition(waitDuration time.Duration) *LevelTransition {
	return &LevelTransition{
		timer:     engine.NewTimer(waitDuration),
		isStarted: false,
	}
}

func (l *LevelTransition) Update(w donburi.World) {
	if !l.isStarted {
		return
	}

	l.timer.Update()

	if l.timer.IsReady() {
		event.FinishedLevelTransition.Publish(w, event.FinishedLevelTransitionData{})
		l.isStarted = false
	}
}

func (l *LevelTransition) OnReachedGoal(_ donburi.World, _ event.ReachedGoalData) {
	l.isStarted = true
}
