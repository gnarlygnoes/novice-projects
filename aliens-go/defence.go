package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Defence struct {
	RecIn  rl.Rectangle
	Rec    rl.Rectangle
	Pos    rl.Vector2
	Colour rl.Color
	Health int32
	Active bool
}

func (g *Game) DefenceBehaviour() {
	for i := range g.Defence {
		if g.Defence[i].Health <= 0 {
			g.Defence[i].Active = false
			g.Defence[i].Rec.X = -15000
		}
	}
}
