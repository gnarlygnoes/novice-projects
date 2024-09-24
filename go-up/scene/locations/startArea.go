package locations

import (
	"fmt"
	"goup/engine"
)

// "goup/engine"
// "goup/game"

type GameLevel struct {
	LevelName string
	TileSet   string
	NpcData   map[engine.CId]NpcData
}

type NpcData struct {
	Id         engine.CId
	PosX, PosY float32
	// Height, Width float32
	// Tex           rl.Texture2D
	NpcType string
	// IsHostile bool
}

func FirstLevel() GameLevel {
	npcs := map[engine.CId]NpcData{}
	e := NpcData{
		Id:      engine.NextId(),
		PosX:    2000,
		PosY:    1000,
		NpcType: "Green Square",
		// IsHostile: true,
	}
	e1 := NpcData{
		Id:      engine.NextId(),
		PosX:    4000,
		PosY:    1000,
		NpcType: "Red Rectangle",
	}
	e2 := NpcData{
		Id:      engine.NextId(),
		PosX:    6000,
		PosY:    1000,
		NpcType: "string",
	}

	npcs[e.Id] = e
	npcs[e1.Id] = e1
	npcs[e2.Id] = e2

	fmt.Println(npcs[0])
	return GameLevel{
		LevelName: "./scene/Village.json",
		TileSet:   "./scene/GroundRevamped.json",
		NpcData:   npcs,
	}
}
