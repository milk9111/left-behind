package component

import "github.com/yohamta/donburi"

type GameData struct {
	WorldWidth  int
	WorldHeight int
	TileSize    int
}

var Game = donburi.NewComponentType[GameData]()
