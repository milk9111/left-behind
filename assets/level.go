package assets

import (
	_ "embed"
	"encoding/json"

	"github.com/milk9111/left-behind/component"
	"github.com/milk9111/left-behind/engine"
	dmath "github.com/yohamta/donburi/features/math"
)

var (
	//go:embed levels/level_1.json
	level1_json []byte
	//go:embed levels/level_2.json
	level2_json []byte
)

var (
	Level1 *Level
	Level2 *Level
)

func init() {
	Level1 = mustLevel(level1_json)
	Level2 = mustLevel(level2_json)
}

type Level struct {
	Name string   `json:"name"`
	Cols int      `json:"cols"`
	Rows int      `json:"rows"`
	Data []string `json:"data"`
}

func mustLevel(b []byte) *Level {
	var level Level
	if err := json.Unmarshal(b, &level); err != nil {
		panic(err)
	}

	return &level
}

func (l *Level) PlayerPosition() dmath.Vec2 {
	for i, c := range l.Data {
		if c != component.CellTypePlayer {
			continue
		}

		return engine.IndexToVec2(i%l.Cols, i/l.Rows)
	}

	panic("no player position found")
}

func (l *Level) GoalPosition() dmath.Vec2 {
	for i, c := range l.Data {
		if c != component.CellTypeGoal {
			continue
		}

		return engine.IndexToVec2(i%l.Cols, i/l.Rows)
	}

	panic("no goal position found")
}

func (l *Level) StickyBlockPositions() []dmath.Vec2 {
	var stickyBlocks []dmath.Vec2
	for i, c := range l.Data {
		if c != component.CellTypeStickyBlock {
			continue
		}

		stickyBlocks = append(stickyBlocks, engine.IndexToVec2(i%l.Cols, i/l.Rows))
	}

	return stickyBlocks
}
