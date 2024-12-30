package scripts

import (
	"time"

	"github.com/milk9111/left-behind/component"
	"github.com/milk9111/left-behind/engine"
	"github.com/milk9111/left-behind/engine/tween"
	"github.com/milk9111/left-behind/event"
	"github.com/yohamta/donburi"
	dmath "github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

type Grid struct {
	e *donburi.Entry
	t *transform.TransformData

	cols int
	rows int
	grid [][]*donburi.Entry

	player              *Player
	globalTweenF64Queue *component.TweenFloat64QueueData
}

func NewGrid(
	e *donburi.Entry,
	cols int,
	rows int,
) *Grid {
	grid := make([][]*donburi.Entry, cols)
	for i := 0; i < cols; i++ {
		grid[i] = make([]*donburi.Entry, rows)
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
	g.globalTweenF64Queue = component.MustFindTweenFloat64Queue(w)

	cells := MustFindEntries(w, component.Cell)

	for _, e := range cells {
		cell := component.Cell.Get(e)
		col, row := engine.Vec2ToIndex(cell.Position)
		g.grid[col][row] = e
	}
}

func (g *Grid) OnStartedStickyTranslation(w donburi.World, eventData event.StartedStickyTranslationData) {
	rotationDegrees := 90.0
	if eventData.IsRotatingBehind {
		rotationDegrees = 180
	}

	nextTween := tween.NewFloat64(
		w,
		1000*time.Millisecond,
		g.t.LocalRotation,
		g.t.LocalRotation-rotationDegrees,
		tween.EaseInOutCubic,
		tween.WithFloat64UpdateCallback(func(t float64) {
			g.t.LocalRotation = t
		}),
	)

	g.globalTweenF64Queue.Enqueue(nextTween)
}

func (g *Grid) CanMove(pos dmath.Vec2) bool {
	col, row := engine.Vec2ToIndex(pos)
	if !g.isValidIndex(col, row) {
		return false
	}

	s := g.grid[col][row]

	return s == nil || component.Cell.Get(s).Type == component.CellTypeGoal
}

func (g *Grid) Move(currPos, nextPos dmath.Vec2) {
	currCol, currRow := engine.Vec2ToIndex(currPos)
	nextCol, nextRow := engine.Vec2ToIndex(nextPos)

	s := g.grid[currCol][currRow]
	g.grid[nextCol][nextRow] = s
	g.grid[currCol][currRow] = nil
}

func (g *Grid) Cell(col, row int) *donburi.Entry {
	if !g.isValidIndex(col, row) {
		return nil
	}

	return g.grid[col][row]
}

func (g *Grid) isValidIndex(col, row int) bool {
	return 0 <= col && col < len(g.grid) && 0 <= row && row < len(g.grid[col])
}

var GridComponent = donburi.NewComponentType[Grid]()
