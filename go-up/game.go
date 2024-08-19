package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	player Player
}

func NewGame() *Game {
	g := &Game{
		player: *NewPlayer(),
	}
	return g
}

func (g *Game) Update() {
	PlayerInputs(g)
}

func (g *Game) Draw() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.Blue)

	rl.DrawRectangleRec(g.player.rec, g.player.colour)

	rl.EndDrawing()
}
