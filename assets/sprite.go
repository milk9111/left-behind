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
	button_png []byte
	//go:embed sprites/floating_block.png
	panel_png []byte
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
	//go:embed sprites/controls_text.png
	controlsText_png []byte
	//go:embed sprites/icon_16x16.png
	icon16x16_png []byte
	//go:embed sprites/icon_48x48.png
	icon48x48_png []byte
)

var (
	SpritePanel        *ebiten.Image
	SpriteTrixie       *ebiten.Image
	SpriteTree         *ebiten.Image
	SpriteGrass        *ebiten.Image
	SpriteGrassTile    *ebiten.Image
	SpriteRock         *ebiten.Image
	SpriteBunny        *ebiten.Image
	SpriteTruffles     *ebiten.Image
	SpriteControlsText *ebiten.Image
	SpriteIcon16x16    *ebiten.Image
	SpriteIcon48x48    *ebiten.Image

	SpriteButtonIdle         *ebiten.Image
	SpriteButtonHover        *ebiten.Image
	SpriteButtonPressed      *ebiten.Image
	SpriteButtonPressedHover *ebiten.Image
	SpriteButtonDisabled     *ebiten.Image
)

func init() {
	SpritePanel = mustImage(panel_png)
	SpriteTrixie = mustImage(trixie_png)
	SpriteTree = mustImage(tree_png)
	SpriteGrass = mustImage(grass_png)
	SpriteGrassTile = mustImage(grassTile_png)
	SpriteRock = mustImage(rock_png)
	SpriteBunny = mustImage(bunny_png)
	SpriteTruffles = mustImage(truffles_png)
	SpriteControlsText = mustImage(controlsText_png)
	SpriteIcon16x16 = mustImage(icon16x16_png)
	SpriteIcon48x48 = mustImage(icon48x48_png)

	spriteSheetButton := mustImage(button_png)
	SpriteButtonIdle = spriteSheetButton.SubImage(image.Rect(0, 0, 32, 32)).(*ebiten.Image)
	SpriteButtonHover = spriteSheetButton.SubImage(image.Rect(32, 0, 64, 32)).(*ebiten.Image)
	SpriteButtonPressed = spriteSheetButton.SubImage(image.Rect(64, 0, 96, 32)).(*ebiten.Image)
	SpriteButtonPressedHover = spriteSheetButton.SubImage(image.Rect(0, 32, 32, 64)).(*ebiten.Image)
	SpriteButtonDisabled = spriteSheetButton.SubImage(image.Rect(32, 32, 64, 64)).(*ebiten.Image)
}

func mustImage(b []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewBuffer(b))
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}
