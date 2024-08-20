package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	rec       rl.Rectangle
	colour    rl.Color
	speed     float32
	jumpSpeed float32
	vVel      float32
}

func NewPlayer() *Player {
	return &Player{
		rec: rl.Rectangle{Width: 50,
			Height: 100,
			X:      ScreenWidth / 2,
			Y:      ScreenHeight - 100},
		colour:    rl.Color{R: 150, G: 70, B: 50, A: 255},
		speed:     1000,
		jumpSpeed: 1500,
		vVel:      0,
	}

	// return &p
}

func (p *Player) Update(dt float32) {
	if p.rec.Y >= ScreenHeight-p.rec.Height {
		p.vVel = 0
	} else {
		p.vVel += Gravity * dt
	}

	PlayerInputs(p, dt)

	p.rec.Y += p.vVel * dt
}
