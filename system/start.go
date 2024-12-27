package system

import (
	"github.com/milk9111/left-behind/component"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type Start struct {
	query *donburi.Query
}

func NewStart() *Start {
	return &Start{
		query: donburi.NewQuery(
			filter.Contains(component.Start),
		),
	}
}

func (u *Start) Update(w donburi.World) {
	u.query.Each(w, func(e *donburi.Entry) {
		component.Start.Get(e).Handler.Start(w)
	})
}
