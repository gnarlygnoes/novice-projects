package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Tile struct {
	Rec    rl.Rectangle
	Colour rl.Color
}

func GenerateTileMap(g *Game) ([]Tile, []Tile) {
	var groundTiles []Tile
	var platformTiles []Tile

	x := [][]int32{
		{0, 0, 0, 2, 2, 2, 2, 2, 2, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 2, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 2, 0, 2, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 2, 0, 0, 1},
		{0, 0, 0, 0, 2, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 2, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 2, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 2, 1},
		{0, 0, 0, 0, 0, 0, 0, 2, 2, 1},
		{0, 0, 0, 0, 0, 0, 2, 2, 2, 1},
		{0, 0, 0, 0, 0, 0, 2, 2, 2, 1},
		{0, 0, 0, 0, 0, 0, 0, 2, 2, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 2, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 2, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 2, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 2, 0, 2, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 2, 0, 0, 1},
		{0, 0, 0, 0, 2, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 2, 1},
		{0, 0, 0, 2, 2, 2, 2, 2, 2, 1},
	}

	tileWidth := ScreenWidth / 10
	tileHeight := ScreenHeight / 10

	for i := range x {
		for j, tile := range x[i] {
			switch tile {
			case 1:
				tile := Tile{
					Rec: rl.Rectangle{
						X:      float32(tileWidth * i),
						Y:      float32(tileHeight * j),
						Width:  float32(tileWidth),
						Height: float32(tileHeight),
					},
					Colour: rl.Green,
				}
				groundTiles = append(groundTiles, tile)
			case 2:
				tile := Tile{
					Rec: rl.Rectangle{
						X:      float32(tileWidth * i),
						Y:      float32(tileHeight * j),
						Width:  float32(tileWidth),
						Height: float32(tileHeight),
					},
					Colour: rl.Brown,
				}

				platformTiles = append(platformTiles, tile)
			}
		}
	}

	return groundTiles, platformTiles
}
