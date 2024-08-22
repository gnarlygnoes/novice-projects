package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Player struct {
	RecIn  rl.Rectangle
	Rec    rl.Rectangle
	Pos    rl.Vector2
	Colour rl.Color
	Health int32
	Speed  int32
}

type Bullet struct {
	Rec    rl.Rectangle
	Active bool
}

func NewPlayer(gt rl.Texture2D, w, h int32) *Player {
	p := &Player{
		RecIn: rl.Rectangle{
			Width:  float32(w),
			Height: float32(h),
			Y:      0,
			X:      float32(w) * 4,
		},
		Rec: rl.Rectangle{
			X:      float32(ScreenWidth/2 - int32(40)),
			Y:      float32(ScreenHeight - int32(40)),
			Width:  80,
			Height: 80,
		},
		Pos:    rl.Vector2{X: 0, Y: 0},
		Colour: rl.Red,
		Health: 5,
		Speed:  2000,
	}
	// return &Player{
	// 	RecIn:  p.RecIn,
	// 	Rec:    p.Rec,
	// 	Pos:    p.Pos,
	// 	Colour: p.Colour,
	// 	Health: p.Health,
	// 	Speed:  p.Speed,
	// }
	return p
}

func (g *Game) HandleInputs() {
	if rl.IsKeyDown(rl.KeyRight) && g.Player.Rec.X < float32(ScreenWidth)-g.Player.Rec.Width {
		g.Player.Rec.X += g.playerSpeed * g.dt
	}
	if rl.IsKeyDown(rl.KeyLeft) && g.Player.Rec.X > 0.0 {
		g.Player.Rec.X -= g.playerSpeed * g.dt
	}
	if rl.IsKeyPressed(rl.KeySpace) {
		g.Shoot()
	}
	if rl.IsKeyPressed(rl.KeyF2) {
		rl.SetTargetFPS(60)
	}
	if rl.IsKeyPressed(rl.KeyF3) {
		rl.SetTargetFPS(120)
	}
	if rl.IsKeyPressed(rl.KeyF4) {
		rl.SetTargetFPS(0)
	}
}

func (g *Game) Shoot() {
	for i := range g.Bullets {
		if !g.Bullets[i].Active {
			g.Bullets[i].Active = true
			g.Bullets[i].Rec.X = g.Player.Rec.X + g.Player.Rec.Width/2 - g.Bullets[i].Rec.Width/2
			g.Bullets[i].Rec.Y = g.Player.Rec.Y - g.Bullets[i].Rec.Height

			break
		}
	}
}

func (g *Game) BulletLogic() {
	for i := range g.Bullets {
		if g.Bullets[i].Active {
			g.Bullets[i].Rec.Y -= g.bulletSpeed * g.dt
			if g.Bullets[i].Rec.Y < 0 {
				g.Bullets[i].Active = false
			}
		}
	}
}
