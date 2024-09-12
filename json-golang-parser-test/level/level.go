package level

import rl "github.com/gen2brain/raylib-go/raylib"

type Tile struct {
	Rec    rl.Rectangle
	Colour rl.Color
}

type Level struct {
	Tiles []Tile
}
