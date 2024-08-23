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
	onGround  bool
	direction float32
	moving    bool
	jump      bool
	// isJumping bool
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

	// return &p
}

func (p *Player) Update(g *Game, dt float32) {
	if p.rec.Y >= ScreenHeight-ScreenHeight/10-p.rec.Height {
		p.rec.Y = ScreenHeight - ScreenHeight/10 - p.rec.Height
		p.onGround = true
	} else {

		p.onGround = false
	}

	PlayerInputs(p, dt)

	// g.GroundCollision(dt)
	g.MoveAndCollideX(dt)

	if p.onGround {
		g.player.vVel = 0
	} else {
		p.vVel += Gravity * dt
	}

	if p.jump {
		p.vVel = -p.jumpSpeed
		p.jump = false
	}

	p.rec.Y += p.vVel * dt
}

func (g *Game) MoveAndCollideX(dt float32) {
	g.player.rec.X += g.player.speed * g.player.direction * dt

	playerWidth := g.player.rec.Width
	playerLeft := g.player.rec.X
	playerRight := playerLeft + playerWidth

	if playerLeft <= 0 {
		g.player.rec.X = 0
	}
	if playerRight >= ScreenWidth {
		g.player.rec.X = ScreenWidth - playerWidth
	}
}

// func (g *Game) GroundCollision(dt float32) {
// 	playerHeight := g.player.rec.Height
// 	playerTop := g.player.rec.Y
// 	playerBottom := playerTop + playerHeight

// 	if playerBottom >= ScreenHeight-ScreenHeight/10 {
// 		if playerBottom > ScreenHeight-ScreenHeight/10 {
// 			g.player.rec.Y = ScreenHeight - ScreenHeight/10 - playerHeight
// 		}
// 		g.player.onGround = true
// 	} else {
// 		g.player.onGround = false
// 	}

// 	if !g.player.onGround {
// 		g.player.vVel += Gravity * dt
// 	} else {
// 		// g.player.rec.Y = ScreenHeight - playerHeight
// 		g.player.vVel = 0
// 	}

// 	if g.player.jump && g.player.onGround {
// 		g.player.vVel = -g.player.jumpSpeed
// 		g.player.isJumping = true
// 		g.player.jump = false
// 	} else {
// 		g.player.isJumping = false
// 	}

// 	g.player.rec.Y += g.player.vVel * dt
// }

// func (g *Game) ObjectCollision() {

// }
