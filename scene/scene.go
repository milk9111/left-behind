package scene

type Scene int

const (
	SceneUnknown Scene = iota
	SceneGame
	SceneMainMenu
	SceneWin
)
