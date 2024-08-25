package main

import (
	"goup/game"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	g := game.NewGame()

	rl.InitWindow(game.ScreenWidth, game.ScreenHeight, "GO UP!! NOT DOWN! UP! UP!!!")
	defer rl.CloseWindow()

	rl.SetTargetFPS(240)

	for !rl.WindowShouldClose() {
		// rl.DrawFPS(20, 30)
		g.Update()

		g.Draw()
	}
}
