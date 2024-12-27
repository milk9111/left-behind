package tween

import (
	"time"

	"github.com/milk9111/left-behind/engine"
	dmath "github.com/yohamta/donburi/features/math"
)

type Vec2 struct {
	timer  *engine.Timer
	start  dmath.Vec2
	end    dmath.Vec2
	easing func(t float64) float64
}

func NewVec2(
	duration time.Duration,
	start, end dmath.Vec2,
	easing func(t float64) float64,
) *Vec2 {
	return &Vec2{
		timer:  engine.NewTimer(duration),
		start:  start,
		end:    end,
		easing: easing,
	}
}

func (t *Vec2) Update() dmath.Vec2 {
	if t.timer.IsReady() {
		return t.end
	}

	nextX := engine.Lerp(t.start.X, t.end.X, t.easing(t.timer.PercentDone()))
	nextY := engine.Lerp(t.start.Y, t.end.Y, t.easing(t.timer.PercentDone()))
	t.timer.Update()

	return dmath.NewVec2(nextX, nextY)
}

func (t *Vec2) Done() bool {
	return t.timer.IsReady()
}

func (t *Vec2) End() dmath.Vec2 {
	return t.end
}
