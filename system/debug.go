package system

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/milk9111/left-behind/assets/scripts"
	"github.com/milk9111/left-behind/component"
	"github.com/milk9111/left-behind/engine"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
	"golang.org/x/image/colornames"
)

type Debug struct {
	query *donburi.Query

	grid              *scripts.Grid
	holdStepTimer     *engine.Timer
	nextStepCallback  func()
	pauseGameCallback func()
}

func NewDebug(nextStepCallback, pauseGameCallback func()) *Debug {
	return &Debug{
		query: donburi.NewQuery(
			filter.Contains(
				transform.Transform,
				component.Cell,
				component.Sprite,
			),
		),
		holdStepTimer:     engine.NewTimer(250 * time.Millisecond),
		nextStepCallback:  nextStepCallback,
		pauseGameCallback: pauseGameCallback,
	}
}

func (d *Debug) Update(w donburi.World) {
	if d.grid == nil {
		d.grid = scripts.MustFindComponent(w, scripts.GridComponent)
	}

	isHoldingStep := ebiten.IsKeyPressed(ebiten.KeyE)
	if isHoldingStep {
		d.holdStepTimer.Update()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyE) || (isHoldingStep && d.holdStepTimer.IsReady()) {
		d.nextStepCallback()
	}

	if inpututil.IsKeyJustReleased(ebiten.KeyE) {
		d.holdStepTimer.Reset()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		d.pauseGameCallback()
	}
}

func (d *Debug) DebugDraw(w donburi.World, screen *ebiten.Image) {
	if d.grid == nil {
		return
	}

	d.query.Each(w, func(e *donburi.Entry) {
		t := transform.Transform.Get(e)
		sprite := component.Sprite.Get(e)
		vector.StrokeRect(
			screen,
			float32(t.LocalPosition.X),
			float32(t.LocalPosition.Y),
			float32(sprite.Image.Bounds().Dx()),
			float32(sprite.Image.Bounds().Dy()),
			2,
			colornames.Red,
			false,
		)
	})

	for col, entries := range d.grid.Cells() {
		for row, e := range entries {
			if e == nil {
				continue
			}

			// t := transform.Transform.Get(e)
			sprite := component.Sprite.Get(e)
			cell := component.Cell.Get(e)

			pos := engine.IndexToVec2(col, row)

			var id string
			if e.HasComponent(component.ID) {
				id = fmt.Sprintf(" %d", e.Id())
			}

			ebitenutil.DebugPrintAt(screen, fmt.Sprintf("(%d, %d)", col, row), int(pos.X), int(pos.Y))
			ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s%s", cell.Type, id), int(pos.X)+sprite.Image.Bounds().Dx()/2, int(pos.Y)+sprite.Image.Bounds().Dy()/2)
		}
	}
}
