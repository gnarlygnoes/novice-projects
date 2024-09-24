package game

import (
	"goup/engine"
	"goup/levels"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Tile struct {
	Rec    rl.Rectangle
	Colour rl.Color
	End    bool
}

func GenerateLevel(levelNum int) (mapTiles []Tile, npcs map[engine.CId]NPC,
	items map[engine.CId]Item, startpoint rl.Vector2, endpoint float32) {

	x := levels.ReturnLevel(levelNum)

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
			case 3:
				endpoint = float32(tileWidth * i)
			case 4:
				startpoint.X = float32(tileWidth * i)
				startpoint.Y = float32(tileHeight * j)
			}
		}
	}

	npcs = map[engine.CId]NPC{}

	items = map[engine.CId]Item{}

	npcs, items = GenerateEntities(npcs, items, float32(tileWidth), float32(tileHeight), levelNum)

	return mapTiles, npcs, items, startpoint, endpoint
}

func GenerateEntities(npcs map[engine.CId]NPC, items map[engine.CId]Item,
	tw, th float32, level int) (map[engine.CId]NPC, map[engine.CId]Item) {
	switch level {
	case 1:
		e1 := NewNPC(1800, 500, true, "Red Rectangle")
		e2 := NewNPC(2500, 700, true, "Cat")
		e3 := NewNPC(5000, 800, true, "Green Square")

		npcs[e1.ID] = e1
		npcs[e2.ID] = e2
		npcs[e3.ID] = e3

		healingSalve := MakeItem("health +1", 3200, float32(ScreenHeight)-float32(th*4)-30)
		items[healingSalve.Id] = healingSalve
	}

	return npcs, items
}
