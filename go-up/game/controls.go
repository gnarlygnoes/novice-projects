package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func PlayerInputs(p *Player, dt float32) {
	p.Moving = false
	p.Direction = 0
	if rl.IsKeyDown(rl.KeyRight) {
		p.Moving = true
		p.Direction, p.Facing = 1, 1
	}
	if rl.IsKeyDown(rl.KeyLeft) {
		p.Moving = true
		p.Direction, p.Facing = -1, -1
	}
	if rl.IsKeyPressed(rl.KeySpace) && p.OnSurface {
		p.Jump = true
		// p.canMove = false
	}
	if rl.IsKeyPressed(rl.KeyLeftControl) && !p.Crouched {
		p.Crouched = true
	} else if rl.IsKeyPressed(rl.KeyLeftControl) && p.Crouched {
		p.Crouched = false
		p.Rec.Y -= 50
	}
	// if key a, shoot. if key d, perform melee attack -- for now with a spear.
	if rl.IsKeyDown(rl.KeyA) && p.canShoot { // && !p.Shooting {
		p.Shoot()
	}

	if rl.IsKeyPressed(rl.KeyR) {
		p.ResetPos = true
	}
}
