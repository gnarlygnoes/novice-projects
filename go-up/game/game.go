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
	player        Player
	groundTiles   []Tile
	platformTiles []Tile
}

func NewGame() *Game {
	g := &Game{
		player: *NewPlayer(),
	}
	return g
}

// func (g *Game) SetGameMode() {}

func (g *Game) Update() {
	dt := rl.GetFrameTime()

	g.player.Update(g, dt)
}

func (g *Game) Draw() {

	rl.BeginDrawing()

	rl.ClearBackground(rl.Blue)

	g.groundTiles, g.platformTiles = GenerateTileMap(g)

	for i := range g.groundTiles {
		rl.DrawRectangleRec(g.groundTiles[i].rec, g.groundTiles[i].colour)
	}

	for i := range g.platformTiles {
		rl.DrawRectangleRec(g.platformTiles[i].rec, g.platformTiles[i].colour)
	}

	rl.DrawRectangleRec(g.player.rec, g.player.colour)

	rl.EndDrawing()
}
