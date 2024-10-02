package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ScreenWidth  = 1920
	ScreenHeight = 1080
	NumBranches  = 7
)

func main() {
	rl.InitWindow(ScreenWidth, ScreenHeight, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	// rl.SetTargetFPS(60)

	centre := rl.Vector2{
		X: ScreenWidth / 2,
		Y: ScreenHeight / 2,
	}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.DrawFPS(30, 20)

		rl.ClearBackground(rl.Black)
		draw_snowflakes(centre, 4, 7, 200, 10)

		rl.EndDrawing()
	}
}

func draw_snowflakes(centre rl.Vector2, levels, branches int, length, thickness float32) {
	branchAngle := 2.0 * math.Pi / float32(branches)
	colour := rl.White
	switch levels {
	case 4:
		colour = rl.RayWhite
	case 3:
		colour = rl.Red
	case 2:
		colour = rl.Yellow
	case 1:
		colour = rl.Orange
	}

	if levels > 0 {
		for i := range branches {
			line := rl.Vector2{
				X: centre.X + float32(math.Cos(float64(branchAngle)*float64(i)))*length,
				Y: centre.Y + float32(math.Sin(float64(branchAngle)*float64(i)))*length,
			}
			rl.DrawLineEx(centre, line, thickness, colour)
			draw_snowflakes(line, levels-1, branches, length/1.5, thickness/2.0)
		}
	}
}
