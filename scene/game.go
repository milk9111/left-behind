package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/milk9111/left-behind/archetype"
	"github.com/milk9111/left-behind/assets"
	"github.com/milk9111/left-behind/component"
	"github.com/milk9111/left-behind/engine"
	"github.com/milk9111/left-behind/engine/tween"
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

	inputSystem *system.Input
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
	ui := system.NewUI()
	g.inputSystem = system.NewInput()

	g.systems = []System{
		g.inputSystem,
		system.NewUpdate(),
		// system.NewSticky(),
		system.NewProcessTweens(),
		system.NewProcessEvents(),
		system.NewAudio(),
		render,
		ui, // doesn't matter where ui is in this order
	}

	g.drawables = []Drawable{
		render,
		ui, // ui needs to draw after render
	}

	g.world = g.createWorld()

	system.NewStart().Update(g.world)
}

func (g *Game) createWorld() donburi.World {
	w := donburi.NewWorld()

	game := w.Entry(w.Create(component.Game))
	component.Game.Set(game, g.game)

	f64Queue := w.Entry(w.Create(component.TweenFloat64Queue))
	component.TweenFloat64Queue.SetValue(f64Queue, component.TweenFloat64QueueData{
		Queue: engine.NewQueue[*tween.Float64](),
	})

	vec2Queue := w.Entry(w.Create(component.TweenVec2Queue))
	component.TweenVec2Queue.SetValue(vec2Queue, component.TweenVec2QueueData{
		Queue: engine.NewQueue[*tween.Vec2](),
	})

	archetype.NewUI(w, g.level.Name)

	archetype.NewGrid(w, g.game, g.level.Cols, g.level.Rows)

	archetype.NewPlayer(w, g.level.PlayerPosition())

	archetype.NewGoal(w, g.level.GoalPosition())

	archetype.NewLevelTransition(w)

	archetype.NewStickyTranslation(w)

	for _, pos := range g.level.StickyBlockPositions() {
		archetype.NewStickyBlock(w, pos)
	}

	for _, pos := range g.level.FloatingBlockPositions() {
		archetype.NewFloatingBlock(w, pos)
	}

	event.StartedPlayerMove.Subscribe(w, g.OnStartedPlayerMove)
	event.FinishedPlayerMove.Subscribe(w, g.OnFinishedPlayerMove)
	event.StartedStickyTranslation.Subscribe(w, g.OnStartedStickyTranslation)
	event.FinishedStickyTranslation.Subscribe(w, g.OnFinishedStickyTranslation)
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
}

func (g *Game) OnStartedPlayerMove(_ donburi.World, _ event.StartedPlayerMoveData) {
	g.inputSystem.SetDisabled(true)
}

func (g *Game) OnFinishedPlayerMove(_ donburi.World, _ event.FinishedPlayerMoveData) {
	g.inputSystem.SetDisabled(false)
}

func (g *Game) OnStartedStickyTranslation(_ donburi.World, _ event.StartedStickyTranslationData) {
	g.inputSystem.SetDisabled(true)
}

func (g *Game) OnFinishedStickyTranslation(_ donburi.World, _ event.FinishedStickyTranslationData) {
	g.inputSystem.SetDisabled(false)
}

func (g *Game) OnFinishedLevelTransition(_ donburi.World, _ event.FinishedLevelTransitionData) {
	g.level = assets.NextLevel(g.level.Name)
	g.loadLevel()
}
