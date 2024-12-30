package tween

import (
	"time"

	"github.com/milk9111/left-behind/engine"
	"github.com/yohamta/donburi"
	devents "github.com/yohamta/donburi/features/events"
)

type FinishedEventData struct{}

type Float64 struct {
	FinishedEvent *devents.EventType[FinishedEventData]

	timer          *engine.Timer
	start          float64
	end            float64
	easing         func(t float64) float64
	updateCallback func(t float64)
}

type float64Options struct {
	updateCallback func(t float64)
}

type Float64Option func(opts *float64Options)

func WithFloat64UpdateCallback(callback func(t float64)) Float64Option {
	return func(opts *float64Options) {
		opts.updateCallback = callback
	}
}

func NewFloat64(
	w donburi.World,
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
		timer:          engine.NewTimer(duration),
		start:          start,
		end:            end,
		easing:         easing,
		updateCallback: opts.updateCallback,
		FinishedEvent:  devents.NewEventType[FinishedEventData](),
	}
}

func (t *Float64) Update(w donburi.World) float64 {
	if t.timer.IsReady() {
		return t.end
	}

	next := engine.Lerp(t.start, t.end, t.easing(t.timer.PercentDone()))
	t.timer.Update()

	if t.timer.IsReady() {
		t.FinishedEvent.Publish(w, FinishedEventData{})
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
