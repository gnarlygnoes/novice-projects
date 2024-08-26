package game

import rl "github.com/gen2brain/raylib-go/raylib"

type Camera struct {
	rl.Camera2D
	Width  float32
	Height float32
}

func NewCamera(w, h int) *Camera {
	return &Camera{
		Camera2D: rl.Camera2D{
			Zoom: 1,
		},
		Width:  float32(w),
		Height: float32(h),
	}
}

func (c *Camera) Update(player *Player) {
	c.Target.X = player.Rec.X + (player.Rec.Width / 2) - (ScreenWidth / 2)
}
