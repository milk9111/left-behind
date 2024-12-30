package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/milk9111/left-behind/component"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

var errNoUIFound = newSystemError("no UI found")

type UI struct {
	query *donburi.Query
}

func NewUI() *UI {
	return &UI{
		query: donburi.NewQuery(filter.Contains(component.UI)),
	}
}

func (u *UI) Update(w donburi.World) {
	e, ok := u.query.First(w)
	if !ok {
		panic(errNoUIFound)
	}

	ui := component.UI.Get(e)

	ui.Container.Update()
}

func (u *UI) Draw(w donburi.World, screen *ebiten.Image) {
	e, ok := u.query.First(w)
	if !ok {
		panic(errNoUIFound)
	}

	ui := component.UI.Get(e)

	ui.Container.Draw(screen)
}
