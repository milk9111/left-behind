package assets

import (
	_ "embed"
	"encoding/json"

	"github.com/milk9111/left-behind/assets/scripts"
	dmath "github.com/yohamta/donburi/features/math"
)

var (
	//go:embed levels/level_1.json
	level1_json []byte
)

var (
	Level1 *Level
)

func init() {
	Level1 = mustLevel(level1_json)
}

const (
	cellTypeEmpty  = ""
	cellTypePlayer = "P"
	cellTypeGoal   = "G"
)

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

func (l *Level) PlayerPos() dmath.Vec2 {
	for i, c := range l.Data {
		if c != cellTypePlayer {
			continue
		}

		return scripts.IndexToVec2(i%l.Cols, i/l.Rows)
	}

	panic("no player position found")
}

func (l *Level) GoalPos() dmath.Vec2 {
	for i, c := range l.Data {
		if c != cellTypeGoal {
			continue
		}

		return scripts.IndexToVec2(i%l.Cols, i/l.Rows)
	}

	panic("no goal position found")
}
