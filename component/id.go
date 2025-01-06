package component

import (
	"github.com/milk9111/left-behind/engine"
	"github.com/yohamta/donburi"
)

type IDData struct {
	ID   engine.ID
	Type string
}

var ID = donburi.NewComponentType[IDData]()
