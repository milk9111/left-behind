package archetype

import (
	"time"

	"github.com/milk9111/left-behind/assets/scripts"
	"github.com/milk9111/left-behind/component"
	"github.com/milk9111/left-behind/event"
	"github.com/yohamta/donburi"
)

func NewLevelTransition(w donburi.World) *donburi.Entry {
	e := w.Entry(w.Create(
		component.Update,
	))

	levelTransition := scripts.NewLevelTransition(1500 * time.Millisecond)

	component.Update.SetValue(e, component.UpdateData{
		Handler: levelTransition,
	})

	event.ReachedGoal.Subscribe(w, levelTransition.OnReachedGoal)

	return e
}
