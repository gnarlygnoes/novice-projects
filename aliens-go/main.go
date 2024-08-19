package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	NumStars            = 900
	NumBullets          = 50
	ScreenWidth         = 1000
	ScreenHeight        = 1200
	AnimTime            = 0.5
	EnemyWidth          = 75
	EnemyHeight         = 75
	EnemyGridY          = 5
	EnemyGridX          = 10
	NumDefences         = 4
	DefenceHeight       = 70
	DefenceWidth        = 140
	DefencePositionY    = ScreenHeight - 200 - DefenceHeight
	BulletDisplacement  = -9000
	EnemyDisplacement   = -6000
	EBulletDisplacement = -12000
)

func main() {
	rl.InitWindow(ScreenWidth, ScreenHeight, "Aliens Go Home!")

	// g := Game{}
	// game := g.
	g := NewGame()

	defer rl.CloseWindow()

	// rl.SetTargetFPS(120)

	for !rl.WindowShouldClose() {
		g.Update()

		g.Draw()
	}
}
