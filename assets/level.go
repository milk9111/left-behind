package assets

import (
	"embed"
	"encoding/json"
	"fmt"

	"github.com/milk9111/left-behind/component"
	"github.com/milk9111/left-behind/engine"
	dmath "github.com/yohamta/donburi/features/math"
)

var (
	//go:embed all:levels
	levelsFS embed.FS
)

var Levels map[string]*Level
var LevelConfig *levelConfig

func init() {
	Levels = make(map[string]*Level)

	dirEntries, err := levelsFS.ReadDir("levels")
	if err != nil {
		panic(err)
	}

	for _, entry := range dirEntries {
		f, err := levelsFS.ReadFile(fmt.Sprintf("levels/%s", entry.Name()))
		if err != nil {
			panic(err)
		}

		if entry.Name() == "_config.json" {
			var conf levelConfig
			if err := json.Unmarshal(f, &conf); err != nil {
				panic(err)
			}

			LevelConfig = &conf
		} else {
			level := mustLevel(f)
			Levels[level.Name] = level
		}
	}
}

type levelConfig struct {
	LevelOrder    []string `json:"level_order"`
	StartingLevel string   `json:"starting_level"`
}

func StartingLevel() *Level {
	return Levels[LevelConfig.StartingLevel]
}

func NextLevel(name string) *Level {
	for i, lvl := range LevelConfig.LevelOrder {
		if name == lvl && i+1 < len(LevelConfig.LevelOrder) {
			return Levels[LevelConfig.LevelOrder[i+1]]
		}
	}

	return StartingLevel()
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

func (l *Level) FloatingBlockPositions() []dmath.Vec2 {
	var floatingBlocks []dmath.Vec2
	for i, c := range l.Data {
		if c != component.CellTypeFloatingBlock {
			continue
		}

		floatingBlocks = append(floatingBlocks, engine.IndexToVec2(i%l.Cols, i/l.Rows))
	}

	return floatingBlocks
}
