package game

import (
	"goup/engine"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Item struct {
	Id       engine.CId
	Rec      rl.Rectangle
	Colour   rl.Color
	ItemType string
}

func MakeItem(itemType string, x, y float32) (i Item) {
	if itemType == "health +1" {
		i = Item{
			Id: engine.NextId(),
			Rec: rl.Rectangle{
				X:      x,
				Y:      y,
				Width:  30,
				Height: 30,
			},
			Colour:   rl.Red,
			ItemType: itemType,
		}
	}

	return i
}

// func (g *Game) UpdateItems() {
// 	for id, item := range g.items {

// 	}
// }
