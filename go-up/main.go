package main

import (
	"goup/game"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {

	rl.InitWindow(game.ScreenWidth, game.ScreenHeight, "GO UP!! NOT DOWN! UP! UP!!!")
	// level := scene.GenerateLevel()
	defer rl.CloseWindow()

	rl.SetTargetFPS(240)

	g := game.NewGame()

	for !rl.WindowShouldClose() {
		rl.DrawFPS(20, 30)
		g.Update()

		g.Draw()
	}
}
