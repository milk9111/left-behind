package component

import (
	"github.com/milk9111/left-behind/engine"
	"github.com/milk9111/left-behind/engine/tween"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

var errNoTweenFloat64QueueFound = newComponentError("no tween float64 queue found")

type TweenFloat64QueueData struct {
	*engine.Queue[*tween.Float64]
}

var TweenFloat64Queue = donburi.NewComponentType[TweenFloat64QueueData]()

func MustFindTweenFloat64Queue(w donburi.World) *TweenFloat64QueueData {
	t, ok := donburi.NewQuery(filter.Contains(TweenFloat64Queue)).First(w)
	if !ok {
		panic(errNoTweenFloat64QueueFound)
	}

	return TweenFloat64Queue.Get(t)
}
