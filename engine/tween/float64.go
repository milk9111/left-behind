package tween

import (
	"time"

	"github.com/milk9111/left-behind/engine"
)

type Float64 struct {
	timer  *engine.Timer
	start  float64
	end    float64
	easing func(t float64) float64
}

func NewFloat64(
	duration time.Duration,
	start, end float64,
	easing func(t float64) float64,
) *Float64 {
	return &Float64{
		timer:  engine.NewTimer(duration),
		start:  start,
		end:    end,
		easing: easing,
	}
}

func (t *Float64) Update() float64 {
	next := engine.Lerp(t.start, t.end, t.easing(t.timer.PercentDone()))
	t.timer.Update()

	if t.timer.IsReady() {
		return t.end
	}

	return next
}

func (t *Float64) Done() bool {
	return t.timer.IsReady()
}

func (t *Float64) End() float64 {
	return t.end
}
