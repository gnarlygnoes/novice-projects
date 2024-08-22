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
	// onSurface bool
	direction float32
	moving    bool
	jumping   bool
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
	// if p.rec.Y >= ScreenHeight-p.rec.Height || p.onSurface {
	// 	p.rec.Y = ScreenHeight - p.rec.Height
	// 	p.vVel = 0
	// } else {
	// 	p.vVel += Gravity * dt
	// }

	PlayerInputs(p, dt)

	g.MoveAndCollideY(dt)
	g.MoveAndCollideX(dt)

	// p.rec.Y += p.vVel * dt
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

	// if rl.CheckCollisionRecs(g.player.rec, g.platform.rec) {
	// 	if playerLeft <= g.platform.rec.X+g.platform.rec.Width {
	// 		g.player.rec.X = g.platform.rec.X + g.platform.rec.Width
	// 	}
	// }
}

func (g *Game) MoveAndCollideY(dt float32) {
	playerHeight := g.player.rec.Height
	playerTop := g.player.rec.Y
	playerBottom := playerTop + playerHeight

	// platformWidth := g.platform.rec.Width
	// platformLeft := g.platform.rec.X
	// platformRight := g.platform.rec.X + platformWidth
	// platformHeight := g.platform.rec.Y + g.platform.rec.Height

	// currentPlayerX := g.player.rec.X
	// var modY float32
	if !g.player.jumping {
		if rl.CheckCollisionRecs(g.player.rec, g.platform.rec) {
			if playerBottom >= g.platform.rec.Y {
				g.player.rec.Y = g.platform.rec.Y - playerHeight
				g.player.vVel = 0
			}
		} else if playerBottom >= ScreenHeight {
			g.player.rec.Y = ScreenHeight - playerHeight
			g.player.vVel = 0
			// } else if playerBottom >= ScreenHeight-modY {
			// 	g.player.rec.Y = ScreenHeight
			// }
		} else {
			g.player.vVel += Gravity * dt
		}
	} else {
		g.player.vVel = -g.player.jumpSpeed
		g.player.jumping = false
	}
	// if rl.CheckCollisionRecs(g.player.rec, g.platform.rec) {
	// 	if playerBottom >= g.platform.rec.Y {
	// 		g.player.rec.Y = g.platform.rec.Y - playerHeight
	// 		g.player.vVel = 0
	// 	}
	// }

	// if g.player.jumping {
	// 	g.player.vVel = -g.player.jumpSpeed
	// 	g.player.jumping = false
	// }

	g.player.rec.Y += g.player.vVel * dt
}

// func (g *Game) MoveAndCollide(dt float32) {
// 	g.player.rec.X += g.player.speed * g.player.direction * dt

// 	playerWidth := g.player.rec.Width
// 	playerHeight := g.player.rec.Height

// 	playerLeft := g.player.rec.X
// 	playerTop := g.player.rec.Y
// 	playerRight := g.player.rec.X + g.player.rec.Width
// 	playerBottom := playerTop + playerHeight

// 	platformWidth := g.platform.rec.Width
// 	// platformHeight := g.platform.rec.Height

// 	platformLeft := g.platform.rec.X
// 	platformTop := g.platform.rec.Y
// 	platformRight := platformLeft + platformWidth
// 	// platformBottom := platformTop + platformHeight

// 	if playerLeft <= 0 {
// 		g.player.rec.X = 0
// 	}
// 	if playerRight >= ScreenWidth {
// 		g.player.rec.X = ScreenWidth - playerWidth
// 	}

// 	// if rl.CheckCollisionRecs(g.player.rec, g.platform.rec) {
// 	// 	// if playerBottom >= platformTop {
// 	// 	// 	g.player.rec.Y = platformTop - playerHeight - 100
// 	// 	// }
// 	// 	if playerBottom > platformTop && playerLeft < platformRight {
// 	// 		g.player.rec.X = platformRight
// 	// 	}
// 	// }

// 	if rl.CheckCollisionRecs(g.player.rec, g.platform.rec) {
// 		if playerBottom > platformTop {
// 			// g.player.rec.Y = platformTop - playerHeight
// 			g.player.onSurface = true
// 		}
// 		if playerLeft <= platformRight {
// 			g.player.rec.X = platformRight
// 		}
// 	} else {
// 		g.player.onSurface = false
// 	}
// }
