package archetype

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/milk9111/left-behind/assets"
	"github.com/milk9111/left-behind/assets/scripts"
	"github.com/milk9111/left-behind/component"
	"github.com/milk9111/left-behind/event"
	"github.com/yohamta/donburi"
	dmath "github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
	"golang.org/x/image/colornames"
)

func NewGrid(w donburi.World, game *component.GameData, cols, rows int) *donburi.Entry {
	e := w.Entry(
		w.Create(
			transform.Transform,
			component.Sprite,
			component.InputHandler,
			component.Update,
			component.Start,
			scripts.GridComponent,
			component.AudioQueue,
		),
	)

	// scales the sprite from 32 bit to 64 bit without needing to change the screen size
	// EDIT: didn't need to end up doing this but leaving here for future reference
	scale := dmath.NewVec2(1, 1)
	pos := dmath.NewVec2(
		float64(game.WorldWidth)/2-(float64(cols)*float64(game.TileSize)*scale.X)/2,
		float64(game.WorldHeight)/2-(float64(rows)*float64(game.TileSize)*scale.Y)/2,
	)

	transform.Transform.Get(e).LocalPosition = pos

	op := &ebiten.DrawImageOptions{}
	gridImg := ebiten.NewImage(cols*game.TileSize*int(scale.X), rows*game.TileSize*int(scale.Y))
	x := 0.0
	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			op.GeoM.Reset()
			spriteImg := assets.SpriteTestCell
			if (i+j)%2 == 1 {
				spriteImg = assets.SpriteTestCell2
			}

			x = math.Max(x, float64(i*spriteImg.Bounds().Dx()*int(scale.X)))

			op.GeoM.Scale(scale.X, scale.Y)
			op.GeoM.Translate(
				float64(i*spriteImg.Bounds().Dx()*int(scale.X)),
				float64(j*spriteImg.Bounds().Dy()*int(scale.Y)),
			)
			gridImg.DrawImage(spriteImg, op)
		}
	}

	component.Sprite.SetValue(e, component.SpriteData{
		Image:  gridImg,
		Layer:  component.SpriteLayerBackground,
		Hidden: false,
	})

	grid := scripts.NewGrid(e, cols, rows)

	component.InputHandler.SetValue(e, component.InputHandlerData{
		Handler: grid,
	})

	component.Update.SetValue(e, component.UpdateData{
		Handler: grid,
	})

	component.Start.SetValue(e, component.StartData{
		Handler: grid,
	})

	scripts.GridComponent.Set(e, grid)

	event.ReachedGoal.Subscribe(w, grid.OnReachedGoal)

	outline := w.Entry(w.Create(
		transform.Transform,
		component.Sprite,
	))

	outlineImg := ebiten.NewImage(gridImg.Bounds().Dx()+4, gridImg.Bounds().Dy()+4)
	outlineImg.Fill(colornames.Black)
	component.Sprite.SetValue(outline, component.SpriteData{
		Image: outlineImg,
		Layer: component.SpriteLayerBackground,
	})

	transform.ChangeParent(outline, e, false)
	transform.Transform.Get(outline).LocalPosition = dmath.NewVec2(-2, -2)

	return e
}
