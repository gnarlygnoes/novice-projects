package game

import (
	"goup/levels"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Tile struct {
	Rec    rl.Rectangle
	Colour rl.Color
}

func GenerateTileMap() (mapTiles []Tile, enemies []NPC) {

	x := levels.GenerateGameLevels()

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
				mapTiles = append(mapTiles, tile)
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

				mapTiles = append(mapTiles, tile)
			}
		}
	}

	e1 := NewNPC(1800, 500, true)
	e2 := NewNPC(2500, 700, true)

	enemies = append(enemies, e1, e2)

	// enemies = make([]NPC, )

	return mapTiles, enemies
}
