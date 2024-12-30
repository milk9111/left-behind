package tween

import (
	"time"

	"github.com/milk9111/left-behind/engine"
	"github.com/yohamta/donburi"
	devents "github.com/yohamta/donburi/features/events"
	dmath "github.com/yohamta/donburi/features/math"
)

type Vec2 struct {
	FinishedEvent *devents.EventType[FinishedEventData]

	timer          *engine.Timer
	start          dmath.Vec2
	end            dmath.Vec2
	easing         func(t float64) float64
	updateCallback func(t dmath.Vec2)
}

type vec2Options struct {
	updateCallback func(t dmath.Vec2)
}

type Vec2Option func(opts *vec2Options)

func WithVec2UpdateCallback(callback func(t dmath.Vec2)) Vec2Option {
	return func(opts *vec2Options) {
		opts.updateCallback = callback
	}
}

func NewVec2(
	w donburi.World,
	duration time.Duration,
	start, end dmath.Vec2,
	easing func(t float64) float64,
	options ...Vec2Option,
) *Vec2 {
	opts := &vec2Options{}

	for _, opt := range options {
		opt(opts)
	}

	return &Vec2{
		timer:          engine.NewTimer(duration),
		start:          start,
		end:            end,
		easing:         easing,
		updateCallback: opts.updateCallback,
		FinishedEvent:  devents.NewEventType[FinishedEventData](),
	}
}

func (t *Vec2) Update(w donburi.World) dmath.Vec2 {
	if t.timer.IsReady() {
		return t.end
	}

	nextX := engine.Lerp(t.start.X, t.end.X, t.easing(t.timer.PercentDone()))
	nextY := engine.Lerp(t.start.Y, t.end.Y, t.easing(t.timer.PercentDone()))
	next := dmath.NewVec2(nextX, nextY)

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

func (t *Vec2) Done() bool {
	return t.timer.IsReady()
}

func (t *Vec2) End() dmath.Vec2 {
	return t.end
}
