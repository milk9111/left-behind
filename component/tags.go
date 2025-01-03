package component

import "github.com/yohamta/donburi"

type TagGoalData struct{}

var TagGoal = donburi.NewComponentType[TagGoalData]()

type TagOutlineData struct{}

var TagOutline = donburi.NewComponentType[TagOutlineData]()
