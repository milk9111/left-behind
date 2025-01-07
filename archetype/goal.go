package archetype

import (
	"github.com/milk9111/left-behind/assets"
	"github.com/milk9111/left-behind/assets/scripts"
	"github.com/milk9111/left-behind/component"
	"github.com/yohamta/donburi"
	dmath "github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

func NewGoal(w donburi.World, position dmath.Vec2) *donburi.Entry {
	e := w.Entry(w.Create(
		transform.Transform,
		component.Sprite,
		component.Cell,
		component.Start,
		component.TagGoal,
		component.Sticky,
		scripts.StaticComponent,
	))

	transform.Transform.Get(e).LocalPosition = position

	component.Cell.SetValue(e, component.CellData{
		Position: position,
		Type:     component.CellTypeGoal,
	})

	component.Sprite.SetValue(e, component.SpriteData{
		Image: assets.SpriteTruffles,
		Layer: component.SpriteLayerEntity,
	})

	component.Sticky.SetValue(e, component.StickyData{
		Disabled: true,
	})

	static := scripts.NewStatic(e)
	component.Start.SetValue(e, component.StartData{
		Handler: static,
	})

	scripts.StaticComponent.Set(e, static)

	return e
}
