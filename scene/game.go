package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/milk9111/left-behind/archetype"
	"github.com/milk9111/left-behind/component"
	"github.com/milk9111/left-behind/system"
	"github.com/yohamta/donburi"
)

type System interface {
	Update(w donburi.World)
}

type Drawable interface {
	Draw(w donburi.World, s *ebiten.Image)
}

type Debuggable interface {
	DebugDraw(w donburi.World, s *ebiten.Image)
}

type Game struct {
	game *component.GameData

	world       donburi.World
	systems     []System
	drawables   []Drawable
	debuggables []Debuggable
}

func NewGame(game *component.GameData) *Game {
	g := &Game{
		game: game,
	}

	g.loadLevel()

	return g
}

func (g *Game) loadLevel() {
	render := system.NewRender(g.game.WorldWidth, g.game.WorldHeight)

	g.systems = []System{
		render,
	}

	g.drawables = []Drawable{
		render,
	}

	g.world = g.createWorld()
}

func (g *Game) createWorld() donburi.World {
	w := donburi.NewWorld()

	game := w.Entry(w.Create(component.Game))
	component.Game.Set(game, g.game)

	archetype.NewGrid(w, g.game, 5, 5)

	return w
}

func (g *Game) Update() {
	for _, s := range g.systems {
		s.Update(g.world)
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Clear()
	for _, d := range g.drawables {
		d.Draw(g.world, screen)
	}

	// TODO - run DebugDraw
}
