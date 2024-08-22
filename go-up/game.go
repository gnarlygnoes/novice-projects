package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	player   Player
	platform Platform
}

func NewGame() *Game {
	// backgroundTexture :=

	g := &Game{
		player:   *NewPlayer(),
		platform: *MakePlatform(0, ScreenHeight-50, 1000, 50, rl.Green),
	}
	return g
}

// func (g *Game) SetGameMode() {}

func (g *Game) Update() {
	dt := rl.GetFrameTime()

	g.player.Update(g, dt)
	// g.MoveAndCollideX(dt)

	// g.HandleCollisions(dt)
}

func (g *Game) Draw() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.Blue)

	rl.DrawRectangleRec(g.player.rec, g.player.colour)
	rl.DrawRectangleRec(g.platform.rec, g.platform.colour)

	rl.EndDrawing()
}
