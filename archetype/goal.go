package archetype

import (
	"github.com/milk9111/left-behind/assets"
	"github.com/milk9111/left-behind/component"
	"github.com/yohamta/donburi"
	dmath "github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

// TODO - Level file and parsing
// TODO - Sticky blocks and non-sticky blocks
func NewGoal(w donburi.World, position dmath.Vec2) *donburi.Entry {
	e := w.Entry(w.Create(
		transform.Transform,
		component.Sprite,
		component.Sticky,
		component.TagGoal,
	))

	transform.Transform.Get(e).LocalPosition = position

	component.Sticky.SetValue(e, component.StickyData{
		Disabled: true,
	})

	component.Sprite.SetValue(e, component.SpriteData{
		Image: assets.SpriteGoal,
		Layer: component.SpriteLayerEntity,
	})

	return e
}
