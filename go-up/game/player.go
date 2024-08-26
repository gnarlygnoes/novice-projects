package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Rec       rl.Rectangle
	Colour    rl.Color
	Speed     float32
	JumpSpeed float32
	VertVel   float32
	OnSurface bool
	Direction float32
	Moving    bool
	Jump      bool
	ResetPos  bool
	Crouched  bool
	// Bullets   [50]Bullet
	Bullets  [50]Bullet
	Shooting bool
	Facing   float32
}

func NewPlayer() *Player {
	return &Player{
		Rec: rl.Rectangle{
			Width:  50,
			Height: 100,
			X:      ScreenWidth / 2,
			Y:      0},
		Colour:    rl.Color{R: 150, G: 70, B: 50, A: 255},
		Speed:     700,
		JumpSpeed: 1500,
		VertVel:   0,
		Shooting:  false,
		Bullets:   RangedProjectilesInit(),
	}
}

func (p *Player) Update(g *Game, dt float32) {
	if CheckCollisionY(p, g.groundTiles) {
		p.OnSurface = true
	} else if CheckCollisionY(p, g.platformTiles) {
		p.OnSurface = true
	} else {
		p.OnSurface = false
	}

	if p.OnSurface {
		g.player.VertVel = 0
	} else {
		p.VertVel += Gravity * dt
	}

	PlayerInputs(p, dt)
	g.MoveAndCollideX(dt)

	if p.Jump {
		p.VertVel = -p.JumpSpeed
		p.Jump = false
	}

	if p.Crouched {
		p.Rec.Height = 50
	} else {
		p.Rec.Height = 100
	}

	if p.ResetPos {
		p.Rec.X = ScreenWidth / 2
		p.Rec.Y = 0
		p.ResetPos = false
	}

	p.BulletsUpdate(g, dt)

	p.Rec.Y += p.VertVel * dt
}
