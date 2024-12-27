package engine

import (
	"time"
)

type Timer struct {
	currentTick int
	targetTick  int
}

func NewTimer(d time.Duration) *Timer {
	return &Timer{
		currentTick: 0,
		targetTick:  int(d.Milliseconds()) * 60 / 1000,
	}
}

func (t *Timer) Update() {
	if t.currentTick < t.targetTick {
		t.currentTick++
	}
}

func (t *Timer) IsReady() bool {
	return t.currentTick >= t.targetTick
}

func (t *Timer) Reset() {
	t.currentTick = 0
}

func (t *Timer) PercentDone() float64 {
	return float64(t.currentTick) / float64(t.targetTick)
}

func (t *Timer) OverridePercentDone(percentDone float64) {
	t.currentTick = int(percentDone * float64(t.targetTick))
}
