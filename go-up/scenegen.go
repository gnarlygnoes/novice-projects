package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

//	type Tile struct {
//		// rec          rl.Rectangle
//		// colour       rl.Color
//		// hasCollision bool
//		tile int
//	}
// var [][]tile int

// func GenerateTileMap() (x [10][10]int32) {
// 	for i := range x {
// 		for j := range x[i] {
// 			x[i][j] = rl.GetRandomValue(0, 1)
// 		}
// 	}
// 	return x
// }

type GroundTile struct {
	groundCollision bool
	rec             rl.Rectangle
	colour          rl.Color
}

type PlatformTile struct {
	platformCollision bool
	rec               rl.Rectangle
	colour            rl.Color
}

func (g *Game) GenerateTileMap() ([]GroundTile, []PlatformTile) {
	var groundTiles []GroundTile
	var platformTiles []PlatformTile

	// x := GenerateTileMap()
	// fmt.Println(x)
	x := [][]int32{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 2, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 2, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 2, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	}

	// groundTile := Tile{
	// 	groundCollision: true,

	// }
	tileWidth := ScreenWidth / 10
	tileHeight := ScreenHeight / 10

	for i := range x {
		for j, tile := range x[i] {
			// var c rl.Color

			switch tile {
			case 1:
				// c = rl.Green
				// collision type = ground
				tile := GroundTile{
					groundCollision: true,
					rec: rl.Rectangle{
						X:      float32(tileWidth * i),
						Y:      float32(tileHeight * j),
						Width:  float32(tileWidth),
						Height: float32(tileHeight),
					},
					colour: rl.Green,
				}
				groundTiles = append(groundTiles, tile)
			case 2:
				tile := PlatformTile{
					platformCollision: true,
					rec: rl.Rectangle{
						X:      float32(tileWidth * i),
						Y:      float32(tileHeight * j),
						Width:  float32(tileWidth),
						Height: float32(tileHeight),
					},
					colour: rl.Brown,
				}
				// c = rl.Brown
				// g.collisionType = platform
				platformTiles = append(platformTiles, tile)
			}

			// if tile >= 1 {
			// 	// rl.DrawRectangle(int32(i)*int32(tileWidth), int32(j)*int32(tileHeight),
			// 	// 	int32(tileWidth), int32(tileHeight), c)
			// }
		}
	}

	return groundTiles, platformTiles
}
