package engine

import (
	"github.com/yohamta/donburi"
	dmath "github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

func ScaledPosition(e *donburi.Entry) dmath.Vec2 {
	return transform.WorldPosition(e).Mul(transform.WorldScale(e))
}
