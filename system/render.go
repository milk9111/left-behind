package system

import (
	"slices"
	"sort"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/milk9111/left-behind/component"
	"github.com/milk9111/left-behind/engine"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
	"golang.org/x/exp/shiny/materialdesign/colornames"
)

type Render struct {
	query *donburi.Query
	world *ebiten.Image
}

func NewRender(width, height int) *Render {
	return &Render{
		query: donburi.NewQuery(
			filter.Contains(
				component.Sprite,
			),
		),
		world: ebiten.NewImage(width, height),
	}
}

func (r *Render) Update(w donburi.World) {

}

func (r *Render) Draw(w donburi.World, screen *ebiten.Image) {
	r.world.Clear()

	r.world.Fill(colornames.Red400)

	var entries []*donburi.Entry
	r.query.Each(w, func(e *donburi.Entry) {
		entries = append(entries, e)
	})

	byLayer := make(map[int][]*donburi.Entry)
	for _, entry := range entries {
		layer := int(component.Sprite.Get(entry).Layer)
		if _, ok := byLayer[layer]; !ok {
			byLayer[layer] = []*donburi.Entry{}
		}

		byLayer[layer] = append(byLayer[layer], entry)
	}

	layers := make([]int, len(byLayer))
	i := 0
	for layer := range byLayer {
		layers[i] = layer
		i++
	}

	sort.Ints(layers)

	for _, layer := range layers {
		layerEntries := byLayer[layer]
		slices.SortFunc(layerEntries, func(a, b *donburi.Entry) int {
			aSprite := component.Sprite.Get(a)
			bSprite := component.Sprite.Get(b)

			if aSprite.Hidden || bSprite.Hidden {
				return -1
			}

			aPos := transform.WorldPosition(a)
			bPos := transform.WorldPosition(b)

			return int(aPos.Y - bPos.Y)
		})

		for _, e := range byLayer[layer] {
			sprite := component.Sprite.Get(e)
			if sprite.Hidden {
				continue
			}

			position := transform.WorldPosition(e)
			rotation := transform.WorldRotation(e)

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(-float64(sprite.Image.Bounds().Dx())/2, -float64(sprite.Image.Bounds().Dy())/2)
			op.GeoM.Rotate(engine.Deg2Rad(rotation))
			op.GeoM.Translate(position.X+float64(sprite.Image.Bounds().Dx())/2, position.Y+float64(sprite.Image.Bounds().Dy())/2)
			r.world.DrawImage(sprite.Image, op)
		}
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(2, 2)
	screen.DrawImage(r.world, op)
}
