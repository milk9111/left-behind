package archetype

import (
	"github.com/milk9111/left-behind/assets"
	"github.com/milk9111/left-behind/assets/scripts"
	"github.com/milk9111/left-behind/component"
	"github.com/yohamta/donburi"
	dmath "github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

func NewFloatingBlock(w donburi.World, position dmath.Vec2) *donburi.Entry {
	e := w.Entry(w.Create(
		transform.Transform,
		component.Sprite,
		component.Cell,
		scripts.StaticComponent,
		component.Start,
		component.Sticky,
	))

	component.Sprite.SetValue(e, component.SpriteData{
		Image: assets.SpriteRock,
		Layer: component.SpriteLayerEntity,
	})

	component.Cell.SetValue(e, component.CellData{
		Position: position,
		Type:     component.CellTypeFloatingBlock,
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
