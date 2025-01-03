package scripts

import (
	"time"

	"github.com/milk9111/left-behind/assets"
	"github.com/milk9111/left-behind/component"
	"github.com/milk9111/left-behind/engine"
	"github.com/milk9111/left-behind/engine/tween"
	"github.com/milk9111/left-behind/event"
	"github.com/yohamta/donburi"
	dmath "github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
)

type StickyTranslation struct {
	e          *donburi.Entry
	query      *donburi.Query
	audioQueue *component.AudioQueueData

	grid                 *Grid
	player               *Player
	globalTweenVec2Queue *component.TweenVec2QueueData
}

func NewStickyTranslation(e *donburi.Entry) *StickyTranslation {
	return &StickyTranslation{
		e:          e,
		query:      donburi.NewQuery(filter.Contains(component.Sticky, component.Cell, transform.Transform)),
		audioQueue: component.AudioQueue.Get(e),
	}
}

func (s *StickyTranslation) Start(w donburi.World) {
	s.grid = GridComponent.Get(MustFindEntry(w, GridComponent))
	s.player = PlayerComponent.Get(MustFindEntry(w, PlayerComponent))
	s.globalTweenVec2Queue = component.MustFindTweenVec2Queue(w)
}

func (s *StickyTranslation) OnInput(w donburi.World, inputEventType component.InputEventType) {
	if inputEventType != component.InputEventTypeRotateBehind &&
		inputEventType != component.InputEventTypeRotateLeft {
		return
	}

	stickyCells := make(map[*donburi.Entry]dmath.Vec2)
	conflictingCells := make(map[*donburi.Entry]struct{})
	s.query.Each(w, func(e *donburi.Entry) {
		sticky := component.Sticky.Get(e)
		if sticky.Disabled {
			return
		}

		cell := component.Cell.Get(e)
		col, row := engine.Vec2ToIndex(cell.Position)

		nextCol := row
		nextRow := s.grid.cols - 1 - col
		if inputEventType == component.InputEventTypeRotateBehind {
			nextCol = s.grid.cols - 1 - col
			nextRow = s.grid.rows - 1 - row
		}

		stickyCells[e] = engine.IndexToVec2(nextCol, nextRow)

		nextCellEntry := s.grid.Cell(nextCol, nextRow)
		if nextCellEntry == nil { // empty cell so no conflict
			return
		}

		var nextSticky *component.StickyData
		if nextCellEntry.HasComponent(component.Sticky) {
			nextSticky = component.Sticky.Get(nextCellEntry)
		}

		if nextSticky != nil && nextSticky.Disabled == sticky.Disabled { // has same sticky property as current so no conflict
			return
		}

		nextCell := component.Cell.Get(nextCellEntry)
		if (cell.Type == component.CellTypeGoal || cell.Type == component.CellTypePlayer) &&
			(nextCell.Type == component.CellTypeGoal || nextCell.Type == component.CellTypePlayer) { // player and goal are colliding so no conflict
			// TODO - trigger win event here?
			return
		}

		conflictingCells[e] = struct{}{}
		conflictingCells[nextCellEntry] = struct{}{}
	})

	if len(conflictingCells) > 0 {
		s.audioQueue.Enqueue(assets.SFXBadMove)
		for e := range conflictingCells {
			event.ConflictedOnCell.Publish(w, event.ConflictedOnCellData{Entry: e})
		}
		return
	}

	event.StartedStickyTranslation.Publish(w, event.StartedStickyTranslationData{
		IsRotatingBehind: inputEventType == component.InputEventTypeRotateBehind,
	})

	count := len(stickyCells)
	for e, pos := range stickyCells {
		cell := component.Cell.Get(e)

		s.grid.Move(cell.Position, pos)

		t := transform.Transform.Get(e)
		vec2Tween := tween.NewVec2(
			1000*time.Millisecond,
			cell.Position,
			pos,
			tween.EaseInOutCubic,
			tween.WithVec2UpdateCallback(func(v dmath.Vec2) {
				cell.Position = v
				t.LocalPosition = s.grid.t.LocalPosition.Add(cell.Position)
			}),
			tween.WithVec2FinishedCallback(func() {
				event.FinishedCellMove.Publish(w, event.FinishedCellMoveData{
					Entry: e,
				})

				if cell.Type == component.CellTypePlayer {
					event.FinishedPlayerMove.Publish(w, event.FinishedPlayerMoveData{})
				}

				count--
				if count <= 0 {
					event.FinishedStickyTranslation.Publish(w, event.FinishedStickyTranslationData{})
				}
			}),
		)

		s.globalTweenVec2Queue.Enqueue(vec2Tween)
	}
}
