package assets

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed sprites/test_cell.png
	testCell_png []byte
	//go:embed sprites/test_cell_2.png
	testCell2_png []byte
	//go:embed sprites/test_player.png
	testPlayer_png []byte
	//go:embed sprites/goal.png
	goal_png []byte
	//go:embed sprites/sticky_block.png
	stickyBlock_png []byte
)

var (
	SpriteTestCell    *ebiten.Image
	SpriteTestCell2   *ebiten.Image
	SpriteTestPlayer  *ebiten.Image
	SpriteGoal        *ebiten.Image
	SpriteStickyBlock *ebiten.Image
)

func init() {
	SpriteTestCell = mustImage(testCell_png)
	SpriteTestCell2 = mustImage(testCell2_png)
	SpriteTestPlayer = mustImage(testPlayer_png)
	SpriteGoal = mustImage(goal_png)
	SpriteStickyBlock = mustImage(stickyBlock_png)
}

func mustImage(b []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewBuffer(b))
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}
