package component

import "github.com/yohamta/donburi"

type updater interface {
	Update(w donburi.World)
}

type UpdateData struct {
	Handler updater
}

var Update = donburi.NewComponentType[UpdateData]()
