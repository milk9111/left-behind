package component

import (
	"github.com/milk9111/left-behind/engine"
	"github.com/milk9111/left-behind/engine/tween"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

var errNoTweenVec2QueueFound = newComponentError("no tween vec2 queue found")

type TweenVec2QueueData struct {
	*engine.Queue[*tween.Vec2]
}

var TweenVec2Queue = donburi.NewComponentType[TweenVec2QueueData]()

func MustFindTweenVec2Queue(w donburi.World) *TweenVec2QueueData {
	t, ok := donburi.NewQuery(filter.Contains(TweenVec2Queue)).First(w)
	if !ok {
		panic(errNoTweenVec2QueueFound)
	}

	return TweenVec2Queue.Get(t)
}
