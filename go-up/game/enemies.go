package game

import rl "github.com/gen2brain/raylib-go/raylib"

type Enemy struct {
	Rec     rl.Rectangle
	Colour  rl.Color
	Health  int
	isEnemy bool
	// hasWeight bool
}

func MakeEnemy(xpos, ypos float32, isEnemy bool) Enemy {
	return Enemy{
		Rec: rl.Rectangle{
			X:      xpos,
			Y:      ypos,
			Width:  50,
			Height: 80,
		},
		Colour:  rl.Red,
		Health:  50,
		isEnemy: isEnemy,
	}
}

// func (g *Game) Enemies(enemy Enemy) []Enemy {
// 	g.enemies = append(g.enemies, enemy)

// 	return g.enemies
// }
