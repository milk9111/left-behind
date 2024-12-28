package archetype

import (
	"github.com/milk9111/left-behind/assets"
	"github.com/milk9111/left-behind/assets/scripts"
	"github.com/milk9111/left-behind/component"
	"github.com/yohamta/donburi"
	dmath "github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

func NewPlayer(w donburi.World, position dmath.Vec2) *donburi.Entry {
	e := w.Entry(w.Create(
		transform.Transform,
		component.Sprite,
		scripts.PlayerComponent,
		component.InputHandler,
		component.Start,
		component.Update,
		component.Sticky,
	))

	transform.Transform.Get(e).LocalPosition = position

	component.Sprite.SetValue(e, component.SpriteData{
		Image: assets.SpriteTestPlayer,
		Layer: component.SpriteLayerEntity,
	})

	player := scripts.NewPlayer(e)

	component.InputHandler.SetValue(e, component.InputHandlerData{
		Handler: player,
	})

	component.Start.SetValue(e, component.StartData{
		Handler: player,
	})

	component.Update.SetValue(e, component.UpdateData{
		Handler: player,
	})

	component.Sticky.SetValue(e, component.StickyData{
		Position: position,
		Disabled: false,
	})

	scripts.PlayerComponent.Set(e, player)

	return e
}
