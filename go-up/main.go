package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ScreenWidth  = 1920
	ScreenHeight = 1080
	Gravity      = 9800
)

func main() {
	g := NewGame()

	rl.InitWindow(ScreenWidth, ScreenHeight, "GO UP!! NOT DOWN! UP! UP!!!")
	defer rl.CloseWindow()

	rl.SetTargetFPS(240)

	for !rl.WindowShouldClose() {
		g.Update()

		g.Draw()
	}
}
