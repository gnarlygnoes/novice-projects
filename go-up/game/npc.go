package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type NPC struct {
	ID        CId
	Rec       rl.Rectangle
	Colour    rl.Color
	Health    int
	isEnemy   bool
	OnSurface bool
	VertVel   float32
	hasWeight bool
}

func NewNPC(xpos, ypos float32, isEnemy bool, npcType string) NPC { // take type as string and return
	if npcType == "redRectangle" {
		return NPC{
			ID: NextId(),
			Rec: rl.Rectangle{
				X:      xpos,
				Y:      ypos,
				Width:  50,
				Height: 80,
			},
			Colour:    rl.Red,
			Health:    2,
			isEnemy:   isEnemy,
			VertVel:   0,
			hasWeight: true,
		}
	}
	return NPC{
		ID: NextId(),
		Rec: rl.Rectangle{
			X:      xpos,
			Y:      ypos,
			Width:  50,
			Height: 80,
		},
		Colour:    rl.White,
		Health:    2,
		isEnemy:   isEnemy,
		VertVel:   0,
		hasWeight: true,
	}
}

func (g *Game) UpdateNPC(dt float32) {
	for id, npc := range g.enemies {
		// fmt.Println(g.enemies[i].ID)
		// fmt.Println(g.enemies[id].hasWeight)
		if g.enemies[id].hasWeight {
			if CheckCollisionY(&npc.Rec, g.levelTiles) {
				npc.OnSurface = true
				// g.enemies[id] = npc
			} else {
				npc.OnSurface = false
				// g.enemies[id] = npc
			}

			if g.enemies[id].OnSurface {
				npc.VertVel = 0
				// g.enemies[id] = npc
			} else {
				npc.VertVel += Gravity * dt
				// g.enemies[id] = npc
			}

			npc.Rec.Y += g.enemies[id].VertVel * dt
			g.enemies[id] = npc
		}
	}
}
