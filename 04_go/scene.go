package main

import "github.com/hajimehoshi/ebiten/v2"

type Scene struct {
	Id        int32      `json:"id"`
	Gravity   float32    `json:"gravity"`
	Player    Player     `json:"player"`
	Platforms []Platform `json:"platforms"`
}

func (scene *Scene) Update() int {
	if err := scene.Player.Update(scene); err != nil {
		return -2
	}

	return -1
}

func (scene *Scene) Draw(screen *ebiten.Image) {
	scene.Player.Draw(screen)
}
