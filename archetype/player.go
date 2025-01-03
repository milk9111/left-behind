package archetype

import (
	"github.com/milk9111/left-behind/assets"
	"github.com/milk9111/left-behind/assets/scripts"
	"github.com/milk9111/left-behind/component"
	"github.com/milk9111/left-behind/engine"
	"github.com/milk9111/left-behind/event"
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
		component.Cell,
		component.AudioQueue,
		component.Sticky,
	))

	transform.Transform.Get(e).LocalPosition = position

	transform.Transform.Get(e).LocalScale = dmath.NewVec2(0.9, 0.9)

	component.Sprite.SetValue(e, component.SpriteData{
		Image: assets.SpriteTrixie,
		Layer: component.SpriteLayerEntity,
		Pivot: &dmath.Vec2{
			X: 16,
			Y: 32,
		},
	})

	component.AudioQueue.SetValue(e, component.AudioQueueData{
		Queue: engine.NewQueue[[]byte](),
	})

	component.Sticky.SetValue(e, component.StickyData{
		Disabled: false,
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

	component.Cell.SetValue(e, component.CellData{
		Position: position,
		Type:     component.CellTypePlayer,
	})

	scripts.PlayerComponent.Set(e, player)

	event.ReachedGoal.Subscribe(w, player.OnReachedGoal)
	event.FinishedPlayerMove.Subscribe(w, player.OnFinishedPlayerMove)

	return e
}
