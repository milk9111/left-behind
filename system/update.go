package system

import (
	"github.com/milk9111/left-behind/component"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type Update struct {
	query *donburi.Query
}

func NewUpdate() *Update {
	return &Update{
		query: donburi.NewQuery(
			filter.Contains(component.Update),
		),
	}
}

func (u *Update) Update(w donburi.World) {
	u.query.Each(w, func(e *donburi.Entry) {
		component.Update.Get(e).Handler.Update(w)
	})
}
