package system

import (
	"github.com/milk9111/left-behind/component"
	"github.com/milk9111/left-behind/engine/audio"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type Audio struct {
	query *donburi.Query
}

func NewAudio() *Audio {
	return &Audio{
		query: donburi.NewQuery(
			filter.Contains(
				component.AudioQueue,
			),
		),
	}
}

func (a *Audio) Update(w donburi.World) {
	a.query.Each(w, func(e *donburi.Entry) {
		audioQueue := component.AudioQueue.Get(e)
		for !audioQueue.Empty() {
			entry := audioQueue.Dequeue()
			if len(entry.Clip) == 0 {
				continue
			}

			err := audio.PlayOneShotWav(entry.Clip, audio.WithVolume(entry.Volume))
			if err != nil {
				panic(err)
			}
		}

		audioQueue.Reset()
	})
}
