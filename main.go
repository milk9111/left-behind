package main

import (
	"flag"
	"fmt"
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
	flag.Parse()

	// couldn't get the screen size to dynamically resize so decided to hardcode it here
	config := Config{
		WorldWidth:   1024 / 2,
		WorldHeight:  768 / 2,
		ScreenWidth:  1024,
		ScreenHeight: 768,
	}

	// TODO - figure out how to scale sprites properly while keeping this screen size

	// setup has 3 monitors and this forces the game window to open up on the 3rd one
	monitors := ebiten.AppendMonitors(nil)
	if forceMonitorFlag != nil && *forceMonitorFlag && len(monitors) == 3 {
		ebiten.SetMonitor(monitors[2])
	}

	// icon16x16, _, err := image.Decode(bytes.NewReader(assets.Icon16x16_png))
	// if err != nil {
	// 	panic(err)
	// }

	// icon32x32, _, err := image.Decode(bytes.NewReader(assets.Icon32x32_png))
	// if err != nil {
	// 	panic(err)
	// }

	// icon48x48, _, err := image.Decode(bytes.NewReader(assets.Icon48x48_png))
	// if err != nil {
	// 	panic(err)
	// }

	ebiten.SetWindowTitle("Trixie the Truffler")
	// ebiten.SetWindowIcon([]image.Image{
	// 	icon16x16,
	// 	icon32x32,
	// 	icon48x48,
	// })
	ebiten.SetWindowSize(config.ScreenWidth, config.ScreenHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.MaximizeWindow()

	err := ebiten.RunGame(NewGame(config))
	if err != nil {
		log.Fatal(err)
	}
}

type Scene interface {
	Update() scene.Scene
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

func NewGame(config Config) *Game {
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
				assets.StartingLevel(),
			),
			scene.SceneMainMenu: scene.NewMainMenu(
				gameData,
			),
		},
		currentScene: scene.SceneGame,
	}
}

func (g *Game) Update() error {
	sc, ok := g.scenes[g.currentScene]
	if !ok {
		panic(fmt.Errorf("invalid scene %d", g.currentScene))
	}

	g.currentScene = sc.Update()

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
