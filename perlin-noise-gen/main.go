package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(1920, 1080, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	// image := rl.GenImagePerlinNoise(800, 450, 0, 0, 10)
	// image := rl.GenImageCellular(800, 450, 20)
	image := rl.LoadImage("./img/GrassyField.png")
	// var tex rl.Texture2D
	tex := rl.LoadTextureFromImage(image)
	// rl.ExportImage(*image, "./img/perlinimage.png")
	rl.UnloadImage(image)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		// rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)
		rl.DrawTexture(tex, 1920/2, -2000, rl.White)
		// 	rl.Color{
		// 	R: 100,
		// 	G: 100,
		// 	B: 0,
		// 	A: 255,
		// }

		rl.EndDrawing()
	}
}
