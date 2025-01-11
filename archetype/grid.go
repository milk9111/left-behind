package archetype

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/milk9111/left-behind/assets"
	"github.com/milk9111/left-behind/assets/scripts"
	"github.com/milk9111/left-behind/component"
	"github.com/milk9111/left-behind/engine"
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

	t := transform.Transform.Get(e)
	t.LocalPosition = pos

	op := &ebiten.DrawImageOptions{}
	gridImg := ebiten.NewImage(cols*game.TileSize*int(scale.X), rows*game.TileSize*int(scale.Y))
	x := 0.0
	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			op.GeoM.Reset()
			spriteImg := assets.SpriteGrassTile

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

	component.AudioQueue.SetValue(e, component.AudioQueueData{
		Queue: engine.NewQueue[*component.AudioQueueEntry](),
	})

	grid := scripts.NewGrid(e, cols, rows)

	component.Start.SetValue(e, component.StartData{
		Handler: grid,
	})

	scripts.GridComponent.Set(e, grid)

	event.StartedStickyTranslation.Subscribe(w, grid.OnStartedStickyTranslation)
	event.FinishedStickyTranslation.Subscribe(w, grid.OnFinishedStickyTranslation)

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

	outlineImgWorldPos := transform.WorldPosition(outline)
	newBackground(w, game, image.Rect(
		int(outlineImgWorldPos.X),
		int(outlineImgWorldPos.Y),
		int(outlineImgWorldPos.X)+outlineImg.Bounds().Dx(),
		int(outlineImgWorldPos.Y)+outlineImg.Bounds().Dy(),
	))

	return e
}

func newBackground(w donburi.World, game *component.GameData, gridBounds image.Rectangle) *donburi.Entry {
	e := w.Entry(w.Create(transform.Transform, component.Sprite))

	backgroundImage := ebiten.NewImage(game.WorldWidth, game.WorldHeight)
	backgroundImage.Fill(colornames.Darkgreen)

	component.Sprite.SetValue(e, component.SpriteData{
		Image: backgroundImage,
		Layer: component.SpriteLayerBackground,
	})

	for _, pos := range engine.PoissonDiskSampling(float64(game.WorldWidth), float64(game.WorldHeight), 30, 200) {
		x := int(pos.X)
		y := int(pos.Y)

		intersection := gridBounds.Intersect(image.Rect(x, y, x+game.TileSize, y+game.TileSize))
		for !intersection.Eq(image.Rectangle{}) {
			x = engine.RandomRangeInt(0, game.WorldWidth)
			y = engine.RandomRangeInt(0, game.WorldHeight)

			intersection = gridBounds.Intersect(image.Rect(x, y, x+game.TileSize, y+game.TileSize))
		}

		newGrass(w, e, dmath.NewVec2(
			float64(x),
			float64(y),
		))
	}

	treeBounds := assets.SpriteTree.Bounds()
	for _, pos := range engine.PoissonDiskSampling(float64(game.WorldWidth), float64(game.WorldHeight), 60, 50) {
		x := int(pos.X)
		y := int(pos.Y)

		intersection := gridBounds.Intersect(image.Rect(x, y, x+treeBounds.Dx(), y+treeBounds.Dy()))
		for !intersection.Eq(image.Rectangle{}) {
			x = engine.RandomRangeInt(0, game.WorldWidth)
			y = engine.RandomRangeInt(0, game.WorldHeight)

			intersection = gridBounds.Intersect(image.Rect(x, y, x+treeBounds.Dx(), y+treeBounds.Dy()))
		}

		newTree(w, dmath.NewVec2(
			float64(x),
			float64(y),
		))
	}

	return e
}

func newGrass(w donburi.World, parent *donburi.Entry, pos dmath.Vec2) *donburi.Entry {
	e := w.Entry(w.Create(transform.Transform, component.Sprite))

	transform.ChangeParent(e, parent, false)

	transform.Transform.Get(e).LocalPosition = pos

	component.Sprite.SetValue(e, component.SpriteData{
		Image: assets.SpriteGrass,
		Layer: component.SpriteLayerEntity,
	})

	return e
}

func newTree(w donburi.World, pos dmath.Vec2) *donburi.Entry {
	e := w.Entry(w.Create(transform.Transform, component.Sprite))

	transform.Transform.Get(e).LocalPosition = pos

	component.Sprite.SetValue(e, component.SpriteData{
		Image: assets.SpriteTree,
		Layer: component.SpriteLayerEntity,
		Pivot: &dmath.Vec2{
			X: 16,
			Y: 48,
		},
	})

	return e
}
