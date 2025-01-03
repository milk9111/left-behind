package tween

import (
	"time"

	"github.com/milk9111/left-behind/engine"
	"github.com/yohamta/donburi"
)

type FinishedEventData struct{}

type Float64 struct {
	timer            *engine.Timer
	start            float64
	end              float64
	easing           func(t float64) float64
	finishedCallback func()
	updateCallback   func(t float64)
}

type float64Options struct {
	finishedCallback func()
	updateCallback   func(t float64)
}

type Float64Option func(opts *float64Options)

func WithFloat64FinishedCallback(callback func()) Float64Option {
	return func(opts *float64Options) {
		opts.finishedCallback = callback
	}
}

func WithFloat64UpdateCallback(callback func(t float64)) Float64Option {
	return func(opts *float64Options) {
		opts.updateCallback = callback
	}
}

func NewFloat64(
	duration time.Duration,
	start, end float64,
	easing func(t float64) float64,
	options ...Float64Option,
) *Float64 {
	opts := &float64Options{}
	for _, opt := range options {
		opt(opts)
	}

	return &Float64{
		timer:            engine.NewTimer(duration),
		start:            start,
		end:              end,
		easing:           easing,
		finishedCallback: opts.finishedCallback,
		updateCallback:   opts.updateCallback,
	}
}

func (t *Float64) Update(w donburi.World) float64 {
	if t.timer.IsReady() {
		return t.end
	}

	next := engine.Lerp(t.start, t.end, t.easing(t.timer.PercentDone()))
	t.timer.Update()

	if t.timer.IsReady() {
		if t.finishedCallback != nil {
			t.finishedCallback()
		}
		next = t.end
	}

	if t.updateCallback != nil {
		t.updateCallback(next)
	}

	return next
}

func (t *Float64) Done() bool {
	return t.timer.IsReady()
}

func (t *Float64) End() float64 {
	return t.end
}
