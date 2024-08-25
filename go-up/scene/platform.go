package scene

import rl "github.com/gen2brain/raylib-go/raylib"

type Platform struct {
	rec    rl.Rectangle
	colour rl.Color
}

func MakePlatform(X, Y, w, h float32, colour rl.Color) *Platform {
	p := Platform{
		rec: rl.Rectangle{
			X:      X,
			Y:      Y,
			Width:  w,
			Height: h,
		},
		colour: colour,
	}

	return &p
}
