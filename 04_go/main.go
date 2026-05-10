package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// global game object id count, has the value of the next game object id to be created
var gameObjectId int32 = 0

var WindowWidth = 1280

var WindowHeight = 720

func main() {
	game := &Game{}
	scene := Scene{
		Gravity: 1.00001,
		Player:  MakePlayer(50, 50),
	}

	for i := 0; i < 3; i++ {
		scene.Platforms = append(
			scene.Platforms,
			MakePlatform(
				float32((i*80)+100),
				float32((i*80)+400),
				150,
				25,
				color.RGBA{0, 150, 150, 255},
			),
		)
	}

	game.Scene = scene

	ebiten.SetFullscreen(true)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	Scene Scene
}

func (g *Game) Update() error {
	if nextScene := g.Scene.Update(); nextScene >= 0 {
		g.LoadScene(nextScene)
	} else if nextScene == -2 {
		return ebiten.Termination
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Scene.Draw(screen)

	for _, p := range g.Scene.Platforms {
		p.Draw(screen)
	}
}

func (g *Game) Layout(outsideWith, outsideHeight int) (screenWidth, screenHeight int) {
	return WindowWidth, WindowHeight
}

func (g *Game) LoadScene(nextScene int) {
	// TODO
}
