package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	player Player
	// platform Platform
	// tilemap  [10][10]int32
	groundTiles   []GroundTile
	platformTiles []PlatformTile
}

func NewGame() *Game {
	// backgroundTexture :=

	g := &Game{
		player: *NewPlayer(),
		// platform: *MakePlatform(0, ScreenHeight-50, 1000, 50, rl.Green),
		// tilemap:  GenerateTileMap(),
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

	g.groundTiles, g.platformTiles = g.GenerateTileMap()

	for i := range g.groundTiles {
		rl.DrawRectangleRec(g.groundTiles[i].rec, g.groundTiles[i].colour)
	}

	for i := range g.platformTiles {
		rl.DrawRectangleRec(g.platformTiles[i].rec, g.platformTiles[i].colour)
	}

	rl.DrawRectangleRec(g.player.rec, g.player.colour)
	// rl.DrawRectangleRec(g.platform.rec, g.platform.colour)

	rl.EndDrawing()
}
