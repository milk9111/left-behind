package component

import (
	"github.com/milk9111/left-behind/engine"
	"github.com/yohamta/donburi"
)

type AudioQueueEntry struct {
	Clip   []byte
	Volume float64
}

func NewAudioQueueEntry(clip []byte, volume float64) *AudioQueueEntry {
	return &AudioQueueEntry{
		Clip:   clip,
		Volume: volume,
	}
}

type AudioQueueData struct {
	*engine.Queue[*AudioQueueEntry]
}

var AudioQueue = donburi.NewComponentType[AudioQueueData]()
