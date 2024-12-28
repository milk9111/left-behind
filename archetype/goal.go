package archetype

import (
	"github.com/milk9111/left-behind/assets"
	"github.com/milk9111/left-behind/assets/scripts"
	"github.com/milk9111/left-behind/component"
	"github.com/yohamta/donburi"
	dmath "github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

// TODO - Sticky blocks and non-sticky blocks
func NewGoal(w donburi.World, position dmath.Vec2) *donburi.Entry {
	e := w.Entry(w.Create(
		transform.Transform,
		component.Sprite,
		component.Cell,
		component.Start,
		component.TagGoal,
		scripts.StaticComponent,
	))

	transform.Transform.Get(e).LocalPosition = position

	component.Cell.SetValue(e, component.CellData{
		Position: position,
		IsSticky: false,
		Type:     component.CellTypeGoal,
	})

	component.Sprite.SetValue(e, component.SpriteData{
		Image: assets.SpriteGoal,
		Layer: component.SpriteLayerEntity,
	})

	static := scripts.NewStatic(e)
	component.Start.SetValue(e, component.StartData{
		Handler: static,
	})

	scripts.StaticComponent.Set(e, static)

	return e
}
