package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ScreenWidth  = 1920
	ScreenHeight = 1080
	Gravity      = 9800
)

type Game struct {
	Camera *Camera
	player Player
	// enemy []Enemy

	groundTiles   []Tile
	platformTiles []Tile
}

func NewGame() *Game {
	g := &Game{
		player: *NewPlayer(),
		Camera: NewCamera(ScreenWidth, ScreenHeight),
	}

	return g
}

// func (g *Game) SetGameMode() {}

func (g *Game) Update() {
	dt := rl.GetFrameTime()

	g.player.Update(g, dt)
	g.Camera.Update(&g.player)
}

func (g *Game) Draw() {

	rl.BeginDrawing()

	rl.ClearBackground(rl.Blue)
	rl.BeginMode2D(g.Camera.Camera2D)

	g.groundTiles, g.platformTiles = GenerateTileMap(g)

	for i := range g.groundTiles {
		rl.DrawRectangleRec(g.groundTiles[i].Rec, g.groundTiles[i].Colour)
	}

	for i := range g.platformTiles {
		rl.DrawRectangleRec(g.platformTiles[i].Rec, g.platformTiles[i].Colour)
	}

	for _, b := range g.player.Bullets {
		rl.DrawRectangleRec(b.Rec, b.Colour)
	}

	rl.DrawRectangleRec(g.player.Rec, g.player.Colour)

	rl.EndDrawing()
}
