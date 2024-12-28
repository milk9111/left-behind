package component

import (
	"github.com/milk9111/left-behind/engine/tween"
	"github.com/yohamta/donburi"
	dmath "github.com/yohamta/donburi/features/math"
)

const (
	CellTypeEmpty       = ""
	CellTypePlayer      = "P"
	CellTypeGoal        = "G"
	CellTypeStickyBlock = "s"
)

type CellData struct {
	IsSticky       bool
	Position       dmath.Vec2 // sticky position in the grid, not actual position in world space
	Type           string
	QueuedPosition *dmath.Vec2
	Tween          *tween.Vec2
}

var Cell = donburi.NewComponentType[CellData]()
