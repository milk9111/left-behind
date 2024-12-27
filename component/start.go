package component

import "github.com/yohamta/donburi"

type starter interface {
	Start(w donburi.World)
}

type StartData struct {
	Handler starter
}

var Start = donburi.NewComponentType[StartData]()
