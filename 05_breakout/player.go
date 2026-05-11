package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Direction int

const (
	DirUp Direction = iota
	DirDown
	DirLeft
	DirRight
)

type Player struct {
	Id       int32         `json:"id"`
	Health   int32         `json:"health"`
	Jumps    int32         `json:"jumps"`
	MaxJumps int32         `json:"max_jumps"`
	Dir      Direction     `json:"direction"`
	Pos      Point         `json:"pos"`
	PrevPos  Point         `json:"prev_pos"`
	Dx       float32       `json:"dx"`
	Dy       float32       `json:"dy"`
	Img      *ebiten.Image `json:"-"`
}

// TODO: MakePlayer should be updated for a real game
func MakePlayer(x, y float32) Player {
	player := Player{
		Id:       gameObjectId,
		Health:   100,
		MaxJumps: 1,
		Pos:      Point{x, y},
		PrevPos:  Point{x, y},
		Img:      ebiten.NewImage(50, 50),
	}
	player.Img.Fill(color.RGBA{255, 0, 0, 255}) // red

	gameObjectId++

	return player
}

func (p *Player) Update(scene *Scene) error {
	// previous position is now the "current" position
	p.PrevPos = p.Pos

	// player input
	// --------------------------------------------------------------------------
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.Pos.X -= 4
		p.Dir = DirLeft
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.Pos.X += 4
		p.Dir = DirRight
	}

	// TODO: change quit for a real game
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return ebiten.Termination
	}

	// jump
	if ebiten.IsKeyPressed(ebiten.KeyZ) && p.Jumps < p.MaxJumps {
		p.Jumps++
		p.Dy -= 20
	}

	// dash
	if p.Dx == 0 && inpututil.IsKeyJustPressed(ebiten.KeyC) {
		if p.Dir == DirLeft {
			p.Dx = -15
		} else if p.Dir == DirRight {
			p.Dx = 15
		}
	}

	// player physics
	// --------------------------------------------------------------------------
	p.Pos.Y += p.Dy
	p.Pos.X += p.Dx

	// gravity
	if p.Pos.Y >= (float32(WindowHeight) - 50) {
		p.Dy = 0
		p.Pos.Y = (float32(WindowHeight) - 50)
		p.Jumps = 0
	} else {
		p.Dy += scene.Gravity
	}

	// slow down dx
	if p.Dx < 0 {
		p.Dx += 1
	} else if p.Dx > 0 {
		p.Dx -= 1
	}

	// player collisions
	for _, platform := range scene.Platforms {
		dim := platform.CollisionBox()
		collType := Collision(p, platform)

		switch collType {
		case None: // do nothing
		case Above:
			p.Dy = 0
			p.Pos.Y = dim.Y - 50 - 1
			p.Jumps = 0
		case Below:
			p.Dy = 0
			p.Pos.Y = dim.Y + dim.H + 1
		case Left:
			p.Dx = 0
			p.Pos.X = dim.X - 50 - 1
		case Right:
			p.Dx = 0
			p.Pos.X = dim.X + dim.W + 1
		}
	}

	return nil
}

func (p *Player) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(p.Pos.X), float64(p.Pos.Y))

	screen.DrawImage(p.Img, opts)
}

func (p *Player) ColliderType() CollisionType {
	return All
}

func (p *Player) CollisionBox() Rect { // aka hitbox
	return Rect{
		p.Pos.X,
		p.Pos.Y,
		50,
		50,
	}
}

func (p *Player) PrevCollisionBox() Rect {
	return Rect{
		p.PrevPos.X,
		p.PrevPos.Y,
		50,
		50,
	}
}
