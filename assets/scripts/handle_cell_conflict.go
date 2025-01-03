package scripts

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/milk9111/left-behind/component"
	"github.com/milk9111/left-behind/engine/tween"
	"github.com/milk9111/left-behind/event"
	"github.com/yohamta/donburi"
	dmath "github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

type HandleCellConflict struct {
	globalTweenF64Queue *component.TweenFloat64QueueData
}

func NewHandleCellConflict() *HandleCellConflict {
	return &HandleCellConflict{}
}

func (h *HandleCellConflict) Start(w donburi.World) {
	h.globalTweenF64Queue = component.MustFindTweenFloat64Queue(w)
}

var outlineFillColor = color.NRGBA{220, 20, 60, 0}

func (h *HandleCellConflict) OnConflictedOnCell(w donburi.World, evt event.ConflictedOnCellData) {
	outline, ok := transform.FindChildWithComponent(evt.Entry, component.TagOutline)
	if !ok {
		sprite := component.Sprite.Get(evt.Entry)

		outline = w.Entry(w.Create(
			transform.Transform,
			component.Sprite,
			component.TagOutline,
		))

		outlineImg := ebiten.NewImage(sprite.Image.Bounds().Dx()+2, sprite.Image.Bounds().Dy()+2)
		outlineImg.Fill(outlineFillColor)
		component.Sprite.SetValue(outline, component.SpriteData{
			Image: outlineImg,
			Layer: component.SpriteLayerBackground,
		})

		transform.ChangeParent(outline, evt.Entry, false)
		transform.Transform.Get(outline).LocalPosition = dmath.NewVec2(-1, -1)
	}

	sprite := component.Sprite.Get(outline)

	t := tween.NewFloat64(
		500*time.Millisecond,
		0,
		1,
		tween.EaseInSine,
		tween.WithFloat64UpdateCallback(func(t float64) {
			c := outlineFillColor
			c.A = uint8(t * 255)

			sprite.Image.Fill(c)
		}),
		tween.WithFloat64FinishedCallback(func() {
			h.globalTweenF64Queue.Enqueue(tween.NewFloat64(
				250*time.Millisecond,
				1,
				0,
				tween.EaseInSine,
				tween.WithFloat64UpdateCallback(func(t float64) {
					c := outlineFillColor
					c.A = uint8(t * 255)

					sprite.Image.Fill(c)
				}),
			))
		}),
	)

	h.globalTweenF64Queue.Enqueue(t)
}
