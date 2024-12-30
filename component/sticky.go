package component

import "github.com/yohamta/donburi"

type StickyData struct {
	Disabled bool
}

var Sticky = donburi.NewComponentType[StickyData]()
