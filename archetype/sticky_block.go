package archetype

import (
	"github.com/milk9111/left-behind/assets"
	"github.com/milk9111/left-behind/assets/scripts"
	"github.com/milk9111/left-behind/component"
	"github.com/milk9111/left-behind/engine"
	"github.com/yohamta/donburi"
	dmath "github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

func NewStickyBlock(w donburi.World, position dmath.Vec2) *donburi.Entry {
	e := w.Entry(w.Create(
		transform.Transform,
		component.Sprite,
		component.Cell,
		scripts.StaticComponent,
		component.Start,
		component.Sticky,
		component.ID,
	))

	component.ID.SetValue(e, component.IDData{
		ID: engine.NewID(),
	})

	component.Sprite.SetValue(e, component.SpriteData{
		Image: assets.SpriteBunny,
		Layer: component.SpriteLayerEntity,
		Pivot: &dmath.Vec2{
			X: 16,
			Y: 28,
		},
	})

	component.Cell.SetValue(e, component.CellData{
		Position: position,
		Type:     component.CellTypeStickyBlock,
	})

	component.Sticky.SetValue(e, component.StickyData{
		Disabled: false,
	})

	static := scripts.NewStatic(e)

	component.Start.SetValue(e, component.StartData{
		Handler: static,
	})

	scripts.StaticComponent.Set(e, static)

	return e
}
