package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/milk9111/left-behind/archetype"
	"github.com/milk9111/left-behind/assets"
	"github.com/milk9111/left-behind/component"
	"github.com/milk9111/left-behind/event"
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
	game  *component.GameData
	level *assets.Level

	world       donburi.World
	systems     []System
	drawables   []Drawable
	debuggables []Debuggable
}

func NewGame(game *component.GameData, level *assets.Level) *Game {
	g := &Game{
		game:  game,
		level: level,
	}

	g.loadLevel()

	return g
}

func (g *Game) loadLevel() {
	render := system.NewRender(g.game.WorldWidth, g.game.WorldHeight)

	g.systems = []System{
		system.NewInput(),
		system.NewUpdate(),
		system.NewSticky(),
		system.NewProcessEvents(),
		system.NewAudio(),
		render,
	}

	g.drawables = []Drawable{
		render,
	}

	g.world = g.createWorld()

	system.NewStart().Update(g.world)
}

func (g *Game) createWorld() donburi.World {
	w := donburi.NewWorld()

	game := w.Entry(w.Create(component.Game))
	component.Game.Set(game, g.game)

	archetype.NewGrid(w, g.game, g.level.Cols, g.level.Rows)

	archetype.NewPlayer(w, g.level.PlayerPosition())

	archetype.NewGoal(w, g.level.GoalPosition())

	archetype.NewLevelTransition(w)

	for _, pos := range g.level.StickyBlockPositions() {
		archetype.NewStickyBlock(w, pos)
	}

	for _, pos := range g.level.FloatingBlockPositions() {
		archetype.NewFloatingBlock(w, pos)
	}

	event.FinishedLevelTransition.Subscribe(w, g.OnFinishedLevelTransition)

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

func (g *Game) OnFinishedLevelTransition(_ donburi.World, _ event.FinishedLevelTransitionData) {
	g.level = assets.NextLevel(g.level.Name)
	g.loadLevel()
}
