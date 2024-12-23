package archetype

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/milk9111/left-behind/assets"
	"github.com/milk9111/left-behind/component"
	"github.com/yohamta/donburi"
	dmath "github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

func NewGrid(w donburi.World, game *component.GameData, cols, rows int) *donburi.Entry {
	e := w.Entry(
		w.Create(
			transform.Transform,
			component.Sprite,
		),
	)

	transform.Transform.Get(e).LocalPosition = dmath.NewVec2(
		float64(game.WorldWidth)/2-(float64(cols)*float64(game.TileSize))/2,
		float64(game.WorldHeight)/2-(float64(rows)*float64(game.TileSize))/2,
	)

	op := &ebiten.DrawImageOptions{}
	gridImg := ebiten.NewImage(cols*game.TileSize, rows*game.TileSize)
	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			op.GeoM.Reset()
			spriteImg := assets.SpriteTestCell
			if (i+j)%2 == 1 {
				spriteImg = assets.SpriteTestCell2
			}

			op.GeoM.Translate(
				float64(i*spriteImg.Bounds().Dx()),
				float64(j*spriteImg.Bounds().Dy()),
			)
			gridImg.DrawImage(spriteImg, op)
		}
	}

	component.Sprite.SetValue(e, component.SpriteData{
		Image: gridImg,
		Layer: component.SpriteLayerBackground,
	})

	return e
}
