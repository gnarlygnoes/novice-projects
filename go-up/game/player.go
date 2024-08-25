package game

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	rec       rl.Rectangle
	colour    rl.Color
	speed     float32
	jumpSpeed float32
	vVel      float32
	onSurface bool
	direction float32
	moving    bool
	jump      bool
	resetPos  bool
}

func NewPlayer() *Player {
	return &Player{
		rec: rl.Rectangle{
			Width:  50,
			Height: 100,
			X:      ScreenWidth / 2,
			Y:      0},
		colour:    rl.Color{R: 150, G: 70, B: 50, A: 255},
		speed:     1000,
		jumpSpeed: 1500,
		vVel:      0,
	}
}

func (p *Player) Update(g *Game, dt float32) {
	fmt.Println(p.onSurface)

	if p.rec.Y >= ScreenHeight-ScreenHeight/10-p.rec.Height {
		p.rec.Y = ScreenHeight - ScreenHeight/10 - p.rec.Height
		p.onSurface = true
	} else if PlatformCollisionY(g) {
		p.onSurface = true
	} else {
		p.onSurface = false
	}

	if p.onSurface {
		g.player.vVel = 0
	} else {
		p.vVel += Gravity * dt
	}

	PlayerInputs(p, dt)
	g.MoveAndCollideX(dt)

	if p.jump {
		p.vVel = -p.jumpSpeed
		p.jump = false
	}

	if p.resetPos {
		p.rec.X = ScreenWidth / 2
		p.rec.Y = 0
		p.resetPos = false
	}

	p.rec.Y += p.vVel * dt
}
