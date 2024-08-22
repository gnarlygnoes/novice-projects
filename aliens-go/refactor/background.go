package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Star struct {
	x, y   int32
	w, h   float32
	Colour rl.Color
}

func GenerateStars() Star {
	rVal := rl.GetRandomValue(100, 255)
	gVal := rl.GetRandomValue(100, 255)
	bVal := rl.GetRandomValue(100, 255)

	var c rl.Color
	c.R = uint8(rVal)
	c.G = uint8(gVal)
	c.B = uint8(bVal)
	c.A = 255

	var star Star
	star.x = rl.GetRandomValue(0, int32(rl.GetScreenWidth()))
	star.y = rl.GetRandomValue(0, int32(rl.GetScreenHeight()))
	star.w = float32(rl.GetRandomValue(1, 5)) / 1.3
	star.h = star.w
	star.Colour = c

	return star
}

// func GenerateTexture(width, height int, scale float32) rl.Texture2D {
// 	bImage := rl.GenImagePerlinNoise(ScreenWidth, ScreenHeight, 0, 0, scale)
// 	bTexture := rl.LoadTextureFromImage(bImage)
// 	rl.UnloadImage(bImage)

// 	return bTexture
// }
