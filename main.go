package main

import (
	"flag"
	"fmt"
	"image"
	"log"

	"github.com/milk9111/left-behind/assets"
	"github.com/milk9111/left-behind/component"
	"github.com/milk9111/left-behind/scene"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

func init() {
	audio.NewContext(44100)
}

func main() {
	forceMonitorFlag := flag.Bool("q", false, "force game window to third monitor (for testing)")
	startingLevelFlag := flag.String("L", "A", "start game at specific level")
	flag.Parse()

	// couldn't get the screen size to dynamically resize so decided to hardcode it here
	config := Config{
		WorldWidth:   1024 / 2,
		WorldHeight:  768 / 2,
		ScreenWidth:  1024,
		ScreenHeight: 768,
	}

	// setup has 3 monitors and this forces the game window to open up on the 3rd one
	monitors := ebiten.AppendMonitors(nil)
	if forceMonitorFlag != nil && *forceMonitorFlag && len(monitors) == 3 {
		ebiten.SetMonitor(monitors[2])
	}

	ebiten.SetWindowTitle("Trixie the Truffler")
	ebiten.SetWindowIcon([]image.Image{
		assets.SpriteIcon16x16,
		assets.SpriteTruffles,
		assets.SpriteIcon48x48,
	})
	ebiten.SetWindowSize(config.ScreenWidth, config.ScreenHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.MaximizeWindow()

	var startingLevel *assets.Level
	if startingLevelFlag != nil {
		startingLevel = assets.Levels[*startingLevelFlag]
	} else {
		startingLevel = assets.StartingLevel()
	}

	musicLoop := audio.NewInfiniteLoopWithIntroF32(assets.Music, 0, assets.Music.Length())

	musicPlayer, err := audio.CurrentContext().NewPlayerF32(musicLoop)
	if err != nil {
		log.Fatal(err)
	}

	musicPlayer.SetVolume(0.35)
	musicPlayer.Play()

	err = ebiten.RunGame(NewGame(startingLevel, config))
	if err != nil {
		log.Fatal(err)
	}
}

type Scene interface {
	Update() scene.Scene
	Init()
	Draw(screen *ebiten.Image)
}

type Game struct {
	scenes       map[scene.Scene]Scene
	currentScene scene.Scene
	worldWidth   int
	worldHeight  int
	screenWidth  int
	screenHeight int
}

type Config struct {
	Quick        bool
	WorldWidth   int
	WorldHeight  int
	ScreenWidth  int
	ScreenHeight int
}

func NewGame(startingLevel *assets.Level, config Config) *Game {
	gameData := &component.GameData{
		WorldWidth:  config.WorldWidth,
		WorldHeight: config.WorldHeight,
		TileSize:    32,
	}

	return &Game{
		worldWidth:   config.WorldWidth,
		worldHeight:  config.WorldHeight,
		screenWidth:  config.ScreenWidth,
		screenHeight: config.ScreenHeight,
		scenes: map[scene.Scene]Scene{
			scene.SceneGame: scene.NewGame(
				gameData,
				startingLevel,
			),
			scene.SceneMainMenu: scene.NewMainMenu(
				gameData,
			),
			scene.SceneWin: scene.NewWin(
				gameData,
			),
		},
		currentScene: scene.SceneMainMenu,
	}
}

func (g *Game) Update() error {
	sc, ok := g.scenes[g.currentScene]
	if !ok {
		panic(fmt.Errorf("invalid scene %d", g.currentScene))
	}

	scene := sc.Update()
	if scene != g.currentScene {
		g.scenes[scene].Init()
		g.currentScene = scene
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	sc, ok := g.scenes[g.currentScene]
	if !ok {
		panic(fmt.Errorf("invalid scene %d", g.currentScene))
	}

	sc.Draw(screen)
}

func (g *Game) Layout(_, _ int) (int, int) {
	return g.screenWidth, g.screenHeight
}
