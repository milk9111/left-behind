package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	dmath "github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

type SpriteLayer int

const (
	SpriteLayerBackground SpriteLayer = iota
	SpriteLayerEntity
	SpriteLayerUI
)

type SpriteData struct {
	Image  *ebiten.Image
	Layer  SpriteLayer
	Hidden bool

	pivot *dmath.Vec2
}

func (s *SpriteData) PivotPoint() dmath.Vec2 {
	if s.pivot == nil {
		s.pivot = vec2Ptr(dmath.NewVec2(float64(s.Image.Bounds().Dx())/2, float64(s.Image.Bounds().Dy())/2))
	}

	return *s.pivot
}

var Sprite = donburi.NewComponentType[SpriteData]()

func WorldHidden(entry *donburi.Entry) bool {
	s := Sprite.Get(entry)

	p, ok := transform.GetParent(entry)
	if !ok {
		return s.Hidden
	}

	hidden := WorldHidden(p) || s.Hidden
	return hidden
}

func vec2Ptr(v dmath.Vec2) *dmath.Vec2 {
	return &v
}
