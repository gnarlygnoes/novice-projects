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

func NewNPC(xpos, ypos float32, isEnemy bool) NPC { // take type as string and return
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

func (g *Game) UpdateNPC(dt float32) {
	for i := range g.enemies {
		// fmt.Println(g.enemies[i].ID)
		if g.enemies[i].hasWeight {
			if CheckCollisionY(&g.enemies[i].Rec, g.levelTiles) {
				g.enemies[i].OnSurface = true
			} else {
				g.enemies[i].OnSurface = false
			}

			if g.enemies[i].OnSurface {
				g.enemies[i].VertVel = 0
			} else {
				g.enemies[i].VertVel += Gravity * dt
			}

			g.enemies[i].Rec.Y += g.enemies[i].VertVel * dt
		}
	}
}
