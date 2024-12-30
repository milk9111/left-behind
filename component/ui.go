package component

import (
	"github.com/ebitenui/ebitenui"
	"github.com/yohamta/donburi"
)

type UIData struct {
	Container *ebitenui.UI
}

var UI = donburi.NewComponentType[UIData]()
