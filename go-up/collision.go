package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) PlatformCollision() bool {
	if rl.CheckCollisionRecs(g.player.rec, g.platform.rec) {
		return true
	} else {
		return false
	}
}

func (g *Game) HandleCollisions(dt float32) {
	if g.player.moving {
		g.player.rec.X += g.player.speed * g.player.direction * dt
		if g.PlatformCollision() {

		}
	}

	// if rl.CheckCollisionRecs(g.player.rec, g.platform.rec) {
	// 	if g.player.rec.Y+g.player.rec.Height >= g.platform.rec.Y {
	// 		g.player.onSurface = true
	// 	}
	// 	if g.player.rec.X <= g.platform.rec.X+g.platform.rec.Width {
	// 		// g.player.canMoveLeft = false
	// 		// g.player.speed = 0
	// 		g.player.direction = 0
	// 	}
	// } else {
	// 	g.player.onSurface = false
	// 	// g.player.canMoveLeft = true
	// }
}
