package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func PlayerInputs(p *Player, dt float32) {
	if rl.IsKeyDown(rl.KeyLeft) && p.rec.X > 0 {
		p.rec.X -= p.speed * dt
	}
	if rl.IsKeyDown(rl.KeyRight) && p.rec.X < ScreenWidth-p.rec.Width {
		p.rec.X += p.speed * dt
	}
	if rl.IsKeyPressed(rl.KeySpace) && p.vVel == 0 {
		p.vVel = -p.jumpSpeed
	}
}
