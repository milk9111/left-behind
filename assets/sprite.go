package assets

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed sprites/sticky_block.png
	stickyBlock_png []byte
	//go:embed sprites/floating_block.png
	floatingBlock_png []byte
	//go:embed sprites/trixie_v2.png
	trixie_png []byte
	//go:embed sprites/tree_v2.png
	tree_png []byte
	//go:embed sprites/grass.png
	grass_png []byte
	//go:embed sprites/grass_tile.png
	grassTile_png []byte
	//go:embed sprites/rock.png
	rock_png []byte
	//go:embed sprites/bunny.png
	bunny_png []byte
	//go:embed sprites/truffles.png
	truffles_png []byte
)

var (
	SpriteStickyBlock   *ebiten.Image
	SpriteFloatingBlock *ebiten.Image
	SpriteTrixie        *ebiten.Image
	SpriteTree          *ebiten.Image
	SpriteGrass         *ebiten.Image
	SpriteGrassTile     *ebiten.Image
	SpriteRock          *ebiten.Image
	SpriteBunny         *ebiten.Image
	SpriteTruffles      *ebiten.Image
)

func init() {
	SpriteStickyBlock = mustImage(stickyBlock_png)
	SpriteFloatingBlock = mustImage(floatingBlock_png)
	SpriteTrixie = mustImage(trixie_png)
	SpriteTree = mustImage(tree_png)
	SpriteGrass = mustImage(grass_png)
	SpriteGrassTile = mustImage(grassTile_png)
	SpriteRock = mustImage(rock_png)
	SpriteBunny = mustImage(bunny_png)
	SpriteTruffles = mustImage(truffles_png)
}

func mustImage(b []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewBuffer(b))
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}
