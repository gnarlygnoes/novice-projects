package game

import (
	"goup/levels"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Tile struct {
	Rec    rl.Rectangle
	Colour rl.Color
}

func GenerateTileMap() (mapTiles []Tile, npcs map[CId]NPC, items map[CId]Item, endpoint float32) {

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

	npcs = map[CId]NPC{}
	items = map[CId]Item{}

	e1 := NewNPC(1800, 500, true, "Red Rectangle")
	e2 := NewNPC(2500, 700, true, "Cat")
	e3 := NewNPC(5000, 800, true, "Green Square")

	npcs[e1.ID] = e1
	npcs[e2.ID] = e2
	npcs[e3.ID] = e3

	healingSalve := MakeItem("health +1", 3200, float32(ScreenHeight)-float32(tileHeight*4)-30)
	items[healingSalve.Id] = healingSalve

	endpoint = float32(tileWidth * 33)

	return mapTiles, npcs, items, endpoint
}
