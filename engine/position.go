package engine

import (
	"github.com/yohamta/donburi"
	dmath "github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

func ScaledPosition(e *donburi.Entry) dmath.Vec2 {
	return transform.WorldPosition(e).Mul(transform.WorldScale(e))
}

func IndexToVec2(col, row int) dmath.Vec2 {
	return dmath.NewVec2(float64(col*32), float64(row*32))
}

func Vec2ToIndex(vec2 dmath.Vec2) (int, int) {
	return int(vec2.X / 32), int(vec2.Y / 32)
}

func Vec2InverseScalar(vec2 dmath.Vec2, scalar float64) dmath.Vec2 {
	return dmath.NewVec2(scalar/vec2.X, scalar/vec2.Y)
}
