package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func PlayerInputs(p *Player, dt float32) {
	p.moving = false
	p.direction = 0
	if rl.IsKeyDown(rl.KeyRight) {
		p.moving = true
		p.direction = 1
		// p.rec.X += p.speed * dt
	}
	if rl.IsKeyDown(rl.KeyLeft) {
		p.moving = true
		p.direction = -1
		// p.rec.X -= p.direction * p.speed * dt
	}
	if rl.IsKeyPressed(rl.KeySpace) && p.vVel == 0 {
		p.jump = true
		// p.vVel = -p.jumpSpeed
	}
}
