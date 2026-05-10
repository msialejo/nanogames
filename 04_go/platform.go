package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Platform struct {
	Dim Rect
	Img *ebiten.Image
}

func MakePlatform(x, y, w, h float32, c color.Color) Platform {
	platform := Platform{
		Dim: Rect{x, y, w, h},
		Img: ebiten.NewImage(int(w), int(h)),
	}
	platform.Img.Fill(c)

	return platform
}

func (p Platform) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(p.Dim.X), float64(p.Dim.Y))

	screen.DrawImage(p.Img, opts)
}

func (p Platform) ColliderType() CollisionType {
	return All
}

func (p Platform) CollisionBox() Rect {
	return p.Dim
}

func (p Platform) PrevCollisionBox() Rect {
	return p.Dim
}
