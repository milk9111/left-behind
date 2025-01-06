package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/milk9111/left-behind/archetype"
	"github.com/milk9111/left-behind/assets"
	"github.com/milk9111/left-behind/assets/scripts"
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
	Draw(w donburi.World, screen *ebiten.Image)
}

type Debugable interface {
	DebugDraw(w donburi.World, screen *ebiten.Image)
}

type Game struct {
	game  *component.GameData
	level *assets.Level

	world      donburi.World
	systems    []System
	drawables  []Drawable
	debugables []Debugable

	inputSystem *system.Input

	nextScene   Scene
	paused      bool
	debugPaused bool // same as paused except doesn't disable input
	isStepping  bool
	step        int
	count       int
}

func NewGame(game *component.GameData, level *assets.Level) *Game {
	g := &Game{
		game:  game,
		level: level,
	}

	g.loadLevel()

	return g
}

func (g *Game) Init() {
	g.loadLevel()
}

func (g *Game) loadLevel() {
	g.nextScene = SceneGame

	render := system.NewRender(g.game.WorldWidth, g.game.WorldHeight)
	debug := system.NewDebug(g.nextStep, g.debugPause)
	ui := system.NewUI()
	g.inputSystem = system.NewInput()

	g.systems = []System{
		g.inputSystem,
		system.NewUpdate(),
		system.NewProcessTweens(),
		system.NewProcessEvents(),
		system.NewAudio(),
		render,
		ui, // doesn't matter where ui is in this order
		debug,
	}

	g.drawables = []Drawable{
		render,
		ui, // ui needs to draw after render
	}

	g.debugables = []Debugable{
		debug,
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

	archetype.NewUI(w, g.game, g.level.Name)

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

	handleCellConflictEntry := w.Entry(w.Create(component.Start))
	handleCellConflict := scripts.NewHandleCellConflict()

	component.Start.SetValue(handleCellConflictEntry, component.StartData{
		Handler: handleCellConflict,
	})

	event.StartedPlayerMove.Subscribe(w, g.OnStartedPlayerMove)
	event.FinishedPlayerMove.Subscribe(w, g.OnFinishedPlayerMove)
	event.StartedStickyTranslation.Subscribe(w, g.OnStartedStickyTranslation)
	event.FinishedStickyTranslation.Subscribe(w, g.OnFinishedStickyTranslation)
	event.FinishedLevelTransition.Subscribe(w, g.OnFinishedLevelTransition)
	event.ConflictedOnCell.Subscribe(w, handleCellConflict.OnConflictedOnCell)
	event.PausedGame.Subscribe(w, g.OnPausedGame)
	event.UnpausedGame.Subscribe(w, g.OnUnpausedGame)
	event.ExitedGame.Subscribe(w, g.OnExitedGame)

	return w
}

func (g *Game) nextStep() {
	if !g.debugPaused {
		return
	}

	g.isStepping = true
	g.step = g.count + 1
}

func (g *Game) debugPause() {
	g.debugPaused = !g.debugPaused
}

func (g *Game) Update() Scene {
	for _, s := range g.systems {
		if g.debugPaused && !g.isStepping {
			if _, ok := s.(*system.Debug); !ok {
				continue
			}
		}

		s.Update(g.world)
	}

	if g.isStepping && g.step == g.count {
		g.isStepping = false
	}

	g.count++

	return g.nextScene
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Clear()
	for _, d := range g.drawables {
		d.Draw(g.world, screen)
	}

	for _, d := range g.debugables {
		d.DebugDraw(g.world, screen)
	}
}

func (g *Game) OnStartedPlayerMove(_ donburi.World, _ event.StartedPlayerMoveData) {
	if g.paused {
		return
	}
	g.inputSystem.SetDisabled(true)
}

func (g *Game) OnFinishedPlayerMove(_ donburi.World, _ event.FinishedPlayerMoveData) {
	if g.paused {
		return
	}
	g.inputSystem.SetDisabled(false)
}

func (g *Game) OnStartedStickyTranslation(_ donburi.World, _ event.StartedStickyTranslationData) {
	if g.paused {
		return
	}
	g.inputSystem.SetDisabled(true)
}

func (g *Game) OnFinishedStickyTranslation(_ donburi.World, _ event.FinishedStickyTranslationData) {
	if g.paused {
		return
	}
	g.inputSystem.SetDisabled(false)
}

func (g *Game) OnPausedGame(_ donburi.World, _ event.PausedGameData) {
	g.paused = true
	g.inputSystem.SetDisabled(true)
}

func (g *Game) OnUnpausedGame(_ donburi.World, _ event.UnpausedGameData) {
	g.paused = false
	g.inputSystem.SetDisabled(false)
}

func (g *Game) OnFinishedLevelTransition(_ donburi.World, _ event.FinishedLevelTransitionData) {
	g.level = assets.NextLevel(g.level.Name)
	g.loadLevel()
}

func (g *Game) OnExitedGame(_ donburi.World, _ event.ExitedGameData) {
	g.nextScene = SceneMainMenu
}
