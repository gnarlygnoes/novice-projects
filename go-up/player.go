package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	rec        rl.Rectangle
	colour     rl.Color
	hasPhysics bool
}

func NewPlayer() *Player {
	return &Player{
		rec: rl.Rectangle{Width: 50,
			Height: 100,
			X:      ScreenWidth / 2,
			Y:      ScreenHeight - 100},
		colour:     rl.Color{R: 150, G: 70, B: 50, A: 255},
		hasPhysics: true,
	}

	// return &p
}
