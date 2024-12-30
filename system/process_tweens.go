package system

import (
	"github.com/milk9111/left-behind/component"
	"github.com/yohamta/donburi"
)

type ProcessTweens struct {
	float64Queue *component.TweenFloat64QueueData
	vec2Queue    *component.TweenVec2QueueData
}

func NewProcessTweens() *ProcessTweens {
	return &ProcessTweens{}
}

func (p *ProcessTweens) Update(w donburi.World) {
	if p.float64Queue == nil {
		p.float64Queue = component.MustFindTweenFloat64Queue(w)
	}

	if p.vec2Queue == nil {
		p.vec2Queue = component.MustFindTweenVec2Queue(w)
	}

	f64QueueLen := p.float64Queue.Len()
	for i := 0; i < f64QueueLen; i++ {
		tweenF64 := p.float64Queue.Dequeue()
		tweenF64.Update(w)

		if !tweenF64.Done() {
			p.float64Queue.Enqueue(tweenF64)
		}
	}

	vec2QueueLen := p.vec2Queue.Len()
	for i := 0; i < vec2QueueLen; i++ {
		tweenVec2 := p.vec2Queue.Dequeue()
		tweenVec2.Update(w)

		if !tweenVec2.Done() {
			p.vec2Queue.Enqueue(tweenVec2)
		}
	}
}
