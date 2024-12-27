package assets

import (
	_ "embed"
	"encoding/json"
)

var (
	//go:embed levels/level_1.json
	level1_json []byte
)

var (
	Level1 *Level
)

type Level struct {
	Name string   `json:"name"`
	Cols int      `json:"cols"`
	Rows int      `json:"rows"`
	Data []string `json:"data"`
}

func init() {
	Level1 = mustLevel(level1_json)
}

func mustLevel(b []byte) *Level {
	var level Level
	if err := json.Unmarshal(b, &level); err != nil {
		panic(err)
	}

	return &level
}
