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

	cells := MustFindEntries(w, component.Cell)

	for _, e := range cells {
		cell := component.Cell.Get(e)
		col, row := engine.Vec2ToIndex(cell.Position)
		g.grid[col][row] = e
	}
}

// func (g *Grid) Update(_ donburi.World) {
// 	if g.tween == nil {
// 		return
// 	}

// 	nextRotation := g.tween.Update()
// 	g.t.LocalRotation = nextRotation
// 	if g.tween.Done() {
// 		g.t.LocalRotation = g.tween.End()
// 		g.player.inputDisabled = false
// 		g.inputDisabled = false
// 		g.tween = nil
// 	}
// }

// func (g *Grid) OnInput(inputEventType component.InputEventType) {
// 	if g.inputDisabled {
// 		return
// 	}

// 	// TODO - move this into OnStartedStickyTranslation
// 	// TODO - wire everything up to use the new sticky_translation.go script
// 	var nextTween *tween.Float64
// 	if inputEventType == component.InputEventTypeRotateLeft {
// 		nextTween = tween.NewFloat64(1000*time.Millisecond, g.t.LocalRotation, g.t.LocalRotation-90, tween.EaseInOutCubic)
// 	} else if inputEventType == component.InputEventTypeRotateBehind {
// 		nextTween = tween.NewFloat64(1000*time.Millisecond, g.t.LocalRotation, g.t.LocalRotation-180, tween.EaseInOutCubic)
// 	} else {
// 		return // exit early because it's not input grid cares about
// 	}

// 	// instantiate next grid
// 	nextGrid := make([][]*component.CellData, g.cols)
// 	for i := 0; i < g.cols; i++ {
// 		nextGrid[i] = make([]*component.CellData, g.rows)
// 		for j := 0; j < g.rows; j++ {
// 			s := g.grid[i][j]
// 			if s == nil || s.IsSticky {
// 				continue
// 			}

// 			nextGrid[i][j] = s
// 		}
// 	}

// 	// fill up next grid with translated cells
// 	for i := 0; i < g.cols; i++ {
// 		for j := 0; j < g.rows; j++ {
// 			s := g.grid[i][j]
// 			if s == nil || !s.IsSticky {
// 				continue
// 			}

// 			x := j
// 			y := g.cols - 1 - i
// 			if inputEventType == component.InputEventTypeRotateBehind {
// 				x = g.cols - 1 - i
// 				y = g.rows - 1 - j
// 			}
// 			nextGrid[x][y] = s
// 		}
// 	}

// 	// compare next grid with current grid for invalid conflicts
// 	hasConflict := false
// 	for i := 0; i < g.cols; i++ {
// 		for j := 0; j < g.rows; j++ {
// 			curr := g.grid[i][j]
// 			next := nextGrid[i][j]

// 			if curr == nil || next == nil || curr == next || curr.IsSticky == next.IsSticky {
// 				continue
// 			}

// 			if !(curr.Type == component.CellTypeGoal && next.Type == component.CellTypePlayer) {
// 				hasConflict = true
// 				break
// 			}
// 		}

// 		if hasConflict {
// 			break
// 		}
// 	}

// 	if !hasConflict {
// 		// apply translation
// 		for i := 0; i < g.cols; i++ {
// 			for j := 0; j < g.rows; j++ {
// 				s := nextGrid[i][j]
// 				if s == nil {
// 					continue
// 				}

// 				pos := engine.IndexToVec2(i, j)
// 				s.QueuedPosition = &pos
// 			}
// 		}

// 		g.grid = nextGrid
// 		g.tween = nextTween
// 		g.player.inputDisabled = true
// 		g.inputDisabled = true
// 	} else {
// 		g.audioQueue.Enqueue(assets.SFXBadMove)
// 	}
// }

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
