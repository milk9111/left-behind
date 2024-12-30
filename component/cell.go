package component

import (
	"github.com/yohamta/donburi"
	dmath "github.com/yohamta/donburi/features/math"
)

const (
	CellTypeEmpty         = ""
	CellTypePlayer        = "P"
	CellTypeGoal          = "G"
	CellTypeStickyBlock   = "s"
	CellTypeFloatingBlock = "f"
)

type CellData struct {
	Position dmath.Vec2 // sticky position in the grid, not actual position in world space
	Type     string
}

var Cell = donburi.NewComponentType[CellData]()
