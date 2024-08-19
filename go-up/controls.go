package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func PlayerInputs(g *Game) {
	if rl.IsKeyDown(rl.KeyLeft) && g.player.rec.X > 0 {
		g.player.rec.X -= 10
	}
	if rl.IsKeyDown(rl.KeyRight) && g.player.rec.X < ScreenWidth-g.player.rec.Width {
		g.player.rec.X += 10
	}
	if rl.IsKeyPressed(rl.KeySpace) {
		// g.player.rec.Y -= 100
	}
}
