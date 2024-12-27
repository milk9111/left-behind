package component

import (
	"github.com/milk9111/left-behind/engine/tween"
	"github.com/yohamta/donburi"
	dmath "github.com/yohamta/donburi/features/math"
)

type StickyData struct {
	Disabled       bool
	QueuedPosition *dmath.Vec2
	Tween          *tween.Vec2
}

var Sticky = donburi.NewComponentType[StickyData]()
