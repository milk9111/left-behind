package scripts

import (
	"fmt"
	"time"

	"github.com/milk9111/left-behind/assets"
	"github.com/milk9111/left-behind/component"
	"github.com/milk9111/left-behind/engine"
	"github.com/milk9111/left-behind/engine/tween"
	"github.com/milk9111/left-behind/event"
	"github.com/yohamta/donburi"
	dmath "github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

type Grid struct {
	e          *donburi.Entry
	t          *transform.TransformData
	audioQueue *component.AudioQueueData

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
		e:          e,
		t:          transform.Transform.Get(e),
		audioQueue: component.AudioQueue.Get(e),
		cols:       cols,
		rows:       rows,
		grid:       grid,
	}
}

func (g *Grid) Start(w donburi.World) {
	g.player = MustFindComponent(w, PlayerComponent)
	g.globalTweenF64Queue = component.MustFindTweenFloat64Queue(w)

	g.GatherCells(w)
}

func (g *Grid) GatherCells(w donburi.World) {
	cells := MustFindEntries(w, component.Cell)

	for _, e := range cells {
		cell := component.Cell.Get(e)
		col, row := engine.Vec2ToIndex(cell.Position)
		g.grid[col][row] = e
	}
}

func (g *Grid) OnStartedStickyTranslation(w donburi.World, eventData event.StartedStickyTranslationData) {
	rotationDegrees := 90.0
	audioClip := assets.SFXRotateLeft
	if eventData.IsRotatingBehind {
		rotationDegrees = 180
		audioClip = assets.SFXRotateBehind
	}

	g.audioQueue.Enqueue(audioClip)

	nextTween := tween.NewFloat64(
		1000*time.Millisecond,
		g.t.LocalRotation,
		g.t.LocalRotation-rotationDegrees,
		tween.EaseInOutCubic,
		tween.WithFloat64UpdateCallback(func(t float64) {
			g.t.LocalRotation = t
		}),
		tween.WithFloat64FinishedCallback(func() {
			event.RotatedGrid.Publish(w, event.RotatedGridData{})
		}),
	)

	g.globalTweenF64Queue.Enqueue(nextTween)
}

func (g *Grid) OnFinishedStickyTranslation(w donburi.World, _ event.FinishedStickyTranslationData) {
	g.GatherCells(w)
}

func (g *Grid) CanMove(pos dmath.Vec2) bool {
	col, row := engine.Vec2ToIndex(pos)
	if !g.isValidIndex(col, row) {
		return false
	}

	s := g.grid[col][row]

	return s == nil || component.Cell.Get(s).Type == component.CellTypeGoal
}

func (g *Grid) Move(e *donburi.Entry, currPos, nextPos dmath.Vec2) {
	currCol, currRow := engine.Vec2ToIndex(currPos)
	nextCol, nextRow := engine.Vec2ToIndex(nextPos)

	g.grid[nextCol][nextRow] = e
	g.grid[currCol][currRow] = nil
}

func (g *Grid) SetCell(e *donburi.Entry, pos dmath.Vec2) {
	col, row := engine.Vec2ToIndex(pos)

	g.grid[col][row] = e
}

func (g *Grid) Cell(col, row int) *donburi.Entry {
	if !g.isValidIndex(col, row) {
		return nil
	}

	return g.grid[col][row]
}

func (g *Grid) Cells() [][]*donburi.Entry {
	return g.grid
}

func (g *Grid) isValidIndex(col, row int) bool {
	return 0 <= col && col < len(g.grid) && 0 <= row && row < len(g.grid[col])
}

func (g *Grid) Print() {
	for row := 0; row < g.rows; row++ {
		for col, entries := range g.grid {
			if entries[row] == nil {
				fmt.Printf("|            |")
				continue
			}

			fmt.Printf("| (%d,%d) %s %d |", col, row, component.Cell.Get(entries[row]).Type, entries[row].Id())
		}
		fmt.Println()
	}
}

var GridComponent = donburi.NewComponentType[Grid]()
