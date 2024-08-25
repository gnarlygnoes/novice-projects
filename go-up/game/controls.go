package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func PlayerInputs(p *Player, dt float32) {
	p.moving = false
	p.direction = 0
	if rl.IsKeyDown(rl.KeyRight) {
		p.moving = true
		p.direction = 1
	}
	if rl.IsKeyDown(rl.KeyLeft) {
		p.moving = true
		p.direction = -1
	}
	if rl.IsKeyPressed(rl.KeySpace) && p.onSurface {
		p.jump = true
	}
	if rl.IsKeyPressed(rl.KeyR) {
		p.resetPos = true
	}
}
