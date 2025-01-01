package ui

import (
	"image/color"

	"github.com/ebitenui/ebitenui/image"
	"github.com/hajimehoshi/ebiten/v2"
)

func NewNineSliceImage(img *ebiten.Image, centerWidth, centerHeight int) *image.NineSlice {
	w := img.Bounds().Dx()
	h := img.Bounds().Dy()

	return image.NewNineSlice(img,
		[3]int{(w - centerWidth) / 2, centerWidth, w - (w-centerWidth)/2 - centerWidth},
		[3]int{(h - centerHeight) / 2, centerHeight, h - (h-centerHeight)/2 - centerHeight},
	)
}

func NewNineSliceColor(c color.Color) *image.NineSlice {
	return image.NewNineSliceColor(c)
}

// RGB returns a color.RGBA created from the bits of rgb value.
// RGB(0xAABBCC) is identical to color.RGBA{R: 0xAA, G: 0xBB, B: 0xCC, A: 0xFF}
func RGB(rgb uint64) color.RGBA {
	return color.RGBA{
		R: uint8((rgb & (0xFF << (8 * 2))) >> (8 * 2)),
		G: uint8((rgb & (0xFF << (8 * 1))) >> (8 * 1)),
		B: uint8((rgb & (0xFF << (8 * 0))) >> (8 * 0)),
		A: 0xFF,
	}
}
