package scripts

import (
	"time"

	"github.com/milk9111/left-behind/component"
	"github.com/milk9111/left-behind/engine/tween"
	"github.com/yohamta/donburi"
	dmath "github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

type Grid struct {
	e    *donburi.Entry
	t    *transform.TransformData
	cols int
	rows int

	grid   [][]*component.StickyData
	player *Player

	tween         *tween.Float64
	inputDisabled bool
}

func NewGrid(
	e *donburi.Entry,
	cols int,
	rows int,
) *Grid {
	grid := make([][]*component.StickyData, cols)
	for i := 0; i < cols; i++ {
		grid[i] = make([]*component.StickyData, rows)
	}

	return &Grid{
		e:    e,
		t:    transform.Transform.Get(e),
		cols: cols,
		rows: rows,
		grid: grid,
	}
}

func (g *Grid) Start(w donburi.World) {
	g.player = MustFindComponent(w, PlayerComponent)

	stickies := MustFindEntries(w, component.Sticky)

	for _, e := range stickies {
		sticky := component.Sticky.Get(e)
		col, row := Vec2ToIndex(sticky.Position)
		g.grid[col][row] = sticky
	}
}

func (g *Grid) Update() {
	if g.tween == nil {
		return
	}

	nextRotation := g.tween.Update()
	g.t.LocalRotation = nextRotation
	if g.tween.Done() {
		g.t.LocalRotation = g.tween.End()
		g.player.inputDisabled = false
		g.inputDisabled = false
		g.tween = nil
	}
}

func (g *Grid) OnInput(inputEventType component.InputEventType) {
	if (g.tween != nil && !g.tween.Done()) || g.inputDisabled {
		return
	}

	if inputEventType == component.InputEventTypeRotateLeft {
		g.tween = tween.NewFloat64(1000*time.Millisecond, g.t.LocalRotation, g.t.LocalRotation-90, tween.EaseInOutCubic)
	} else if inputEventType == component.InputEventTypeRotateBehind {
		g.tween = tween.NewFloat64(1000*time.Millisecond, g.t.LocalRotation, g.t.LocalRotation-180, tween.EaseInOutCubic)
	} else {
		return // exit early because it's not input grid cares about
	}

	nextGrid := make([][]*component.StickyData, g.cols)
	for i := 0; i < g.cols; i++ {
		nextGrid[i] = make([]*component.StickyData, g.rows)
	}

	for i := 0; i < g.cols; i++ {
		for j := 0; j < g.rows; j++ {
			s := g.grid[i][j]
			if s == nil {
				continue
			}

			if s.Disabled {
				nextGrid[i][j] = s
			} else {
				x := j
				y := g.cols - 1 - i
				if inputEventType == component.InputEventTypeRotateBehind {
					x = g.cols - 1 - i
					y = g.rows - 1 - j
				}
				nextGrid[x][y] = s
				pos := IndexToVec2(x, y)
				s.QueuedPosition = &pos
			}

			g.grid[i][j] = nil
		}
	}

	g.grid = nextGrid
	g.player.inputDisabled = true
	g.inputDisabled = true
}

func (g *Grid) CanMove(pos dmath.Vec2) bool {
	col, row := Vec2ToIndex(pos)
	if 0 > col || col >= len(g.grid) {
		return false
	}

	return 0 <= row && row < len(g.grid[col])
}

func (g *Grid) Move(currPos, nextPos dmath.Vec2) {
	currCol, currRow := Vec2ToIndex(currPos)
	nextCol, nextRow := Vec2ToIndex(nextPos)

	s := g.grid[currCol][currRow]
	g.grid[nextCol][nextRow] = s
	g.grid[currCol][currRow] = nil
}

func IndexToVec2(col, row int) dmath.Vec2 {
	return dmath.NewVec2(float64(col*32), float64(row*32))
}

func Vec2ToIndex(vec2 dmath.Vec2) (int, int) {
	return int(vec2.X / 32), int(vec2.Y / 32)
}

var GridComponent = donburi.NewComponentType[Grid]()
