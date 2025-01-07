package tween

import (
	"math"
	"time"

	"github.com/milk9111/left-behind/engine"
	"github.com/yohamta/donburi"
	dmath "github.com/yohamta/donburi/features/math"
)

type Vec2 struct {
	timer            *engine.Timer
	start            dmath.Vec2
	end              dmath.Vec2
	easing           func(t float64) float64
	lerping          func() dmath.Vec2
	finishedCallback func()
	updateCallback   func(t dmath.Vec2)
}

type vec2Options struct {
	finishedCallback func()
	updateCallback   func(t dmath.Vec2)
}

type Vec2Option func(opts *vec2Options)

func WithVec2UpdateCallback(callback func(t dmath.Vec2)) Vec2Option {
	return func(opts *vec2Options) {
		opts.updateCallback = callback
	}
}

func WithVec2FinishedCallback(callback func()) Vec2Option {
	return func(opts *vec2Options) {
		opts.finishedCallback = callback
	}
}

func NewVec2(
	duration time.Duration,
	start, end dmath.Vec2,
	easing func(t float64) float64,
	options ...Vec2Option,
) *Vec2 {
	opts := &vec2Options{}

	for _, opt := range options {
		opt(opts)
	}

	v := &Vec2{
		timer:            engine.NewTimer(duration),
		start:            start,
		end:              end,
		easing:           easing,
		finishedCallback: opts.finishedCallback,
		updateCallback:   opts.updateCallback,
	}

	v.lerping = v.Lerp

	return v
}

func NewArcVec2(
	duration time.Duration,
	start, end dmath.Vec2,
	easing func(t float64) float64,
	options ...Vec2Option,
) *Vec2 {
	opts := &vec2Options{}

	for _, opt := range options {
		opt(opts)
	}

	v := &Vec2{
		timer:            engine.NewTimer(duration),
		start:            start,
		end:              end,
		easing:           easing,
		finishedCallback: opts.finishedCallback,
		updateCallback:   opts.updateCallback,
	}

	v.lerping = v.ArcLerp

	return v
}

func (v *Vec2) Update(w donburi.World) dmath.Vec2 {
	if v.timer.IsReady() {
		return v.end
	}

	next := v.lerping()

	v.timer.Update()

	if v.timer.IsReady() {
		if v.finishedCallback != nil {
			v.finishedCallback()
		}

		next = v.end
	}

	if v.updateCallback != nil {
		v.updateCallback(next)
	}

	return next
}

func (v *Vec2) Done() bool {
	return v.timer.IsReady()
}

func (v *Vec2) End() dmath.Vec2 {
	return v.end
}

func (v *Vec2) Lerp() dmath.Vec2 {
	nextX := engine.Lerp(v.start.X, v.end.X, v.easing(v.timer.PercentDone()))
	nextY := engine.Lerp(v.start.Y, v.end.Y, v.easing(v.timer.PercentDone()))
	return dmath.NewVec2(nextX, nextY)
}

func (v *Vec2) ArcLerp() dmath.Vec2 {
	const height = 8

	dx := v.end.Y - v.start.Y
	dy := v.start.X - v.end.Y
	magnitude := math.Sqrt(dx*dx + dy*dy)
	perpendicularX := dx / magnitude
	perpendicularY := dy / magnitude

	halfLength := math.Sqrt(math.Pow(v.end.X-v.start.X, 2)+math.Pow(v.end.Y-v.start.Y, 2)) / 2
	radius := math.Sqrt(halfLength*halfLength + height*height)
	centerX := perpendicularX * (radius - height)
	centerY := perpendicularY * (radius - height)

	startAngle := math.Atan2(v.start.Y-centerY, v.start.X-centerX)
	endAngle := math.Atan2(v.end.Y-centerY, v.end.X-centerX)

	angle := engine.Lerp(startAngle, endAngle, v.timer.PercentDone())

	if endAngle < startAngle {
		endAngle += 2 * math.Pi
	}

	return dmath.NewVec2(centerX+radius*math.Cos(angle), centerY+radius*math.Sin(angle))
}

func CalculateMidpoint(start, end dmath.Vec2, height float64) dmath.Vec2 {
	mid := dmath.Vec2{
		X: (start.X + end.X) / 2,
		Y: (start.Y + end.Y) / 2,
	}

	dx := end.X - start.X
	dy := end.Y - start.Y
	d := math.Sqrt(dx*dx + dy*dy)

	mid.X += height * (dy / d)
	mid.Y -= height * (dx / d)

	return mid
}
