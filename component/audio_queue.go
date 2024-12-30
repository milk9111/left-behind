package component

import (
	"github.com/milk9111/left-behind/engine"
	"github.com/yohamta/donburi"
)

type AudioQueueData struct {
	*engine.Queue[[]byte]
}

var AudioQueue = donburi.NewComponentType[AudioQueueData]()
