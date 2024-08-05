package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	bulletSpeed = 20
)

type Game struct {
	ScreenWidth  int32
	ScreenHeight int32
	GameActive   bool

	Player  Player
	Bullets []Bullet
	Enemies [50]Enemy
}

type Player struct {
	Rec    rl.Rectangle
	Speed  float32
	Colour rl.Color
	Health int32
}

type Bullet struct {
	Rec    rl.Rectangle
	Active bool
}

type Enemy struct {
	Rec    rl.Rectangle
	Speed  float32
	Colour rl.Color
	Alive  bool
}

func main() {
	game := Game{}

	defer rl.CloseWindow()

	rl.SetTargetFPS(120)
	game.Init()
	rl.InitWindow(game.ScreenWidth, game.ScreenHeight, "Aliens of Go")
	for !rl.WindowShouldClose() {
		game.Update()

		game.Draw()
	}
}

func (g *Game) Init() {
	g.ScreenWidth = 800
	g.ScreenHeight = 1200
	g.GameActive = true

	g.Player.Rec.Width = 60
	g.Player.Rec.Height = 80
	g.Player.Rec.X = float32(g.ScreenWidth/2 - int32(g.Player.Rec.Width/2))
	g.Player.Rec.Y = float32(g.ScreenHeight - int32(g.Player.Rec.Height))
	g.Player.Speed = 10
	g.Player.Health = 1
	g.Player.Colour = rl.Red

	for i := range g.Enemies {
		g.Enemies[i].Rec.Width = 60
		g.Enemies[i].Rec.Height = 80
		g.Enemies[i].Rec.X = float32(rl.GetRandomValue(0, g.ScreenWidth-g.Enemies[i].Rec.ToInt32().Width))
		g.Enemies[i].Rec.Y = float32(rl.GetRandomValue(-g.ScreenHeight*4, 0))
		g.Enemies[i].Colour = rl.Blue
		g.Enemies[i].Alive = true
		g.Enemies[i].Speed = 2
	}

	// Generate bullets here using for loop and fixed array size.
}

func (g *Game) Update() {
	if rl.IsKeyDown(rl.KeyRight) && g.Player.Rec.X < float32(g.ScreenWidth)-g.Player.Rec.Width {
		g.Player.Rec.X += g.Player.Speed
	}
	if rl.IsKeyDown(rl.KeyLeft) && g.Player.Rec.X > 0.0 {
		g.Player.Rec.X -= g.Player.Speed
	}
	if rl.IsKeyDown(rl.KeyDown) && g.Player.Rec.Y < float32(g.ScreenHeight)-g.Player.Rec.Width {
		g.Player.Rec.Y += g.Player.Speed
	}
	if rl.IsKeyDown(rl.KeyUp) && g.Player.Rec.Y > 0.0 {
		g.Player.Rec.Y -= g.Player.Speed
	}
	if rl.IsKeyPressed(rl.KeySpace) {
		b := Bullet{}
		b.Rec.X = g.Player.Rec.X + g.Player.Rec.Width/2 - b.Rec.Width/2
		b.Rec.Y = g.Player.Rec.Y - b.Rec.Height
		b.Rec.Width = 5
		b.Rec.Height = 20
		b.Active = true
		g.Bullets = append(g.Bullets, b)
	}

	// GAME LOGIC
	// Bullet
	for i := range g.Bullets {
		if g.Bullets[i].Active {
			g.Bullets[i].Rec.Y -= bulletSpeed
			// if g.Bullets[i].Rec.Y < 0 {
			// 	g.Bullets = append(g.Bullets[:i], g.Bullets[i+1:]...)
			// }
			if g.Bullets[i].Rec.Y < 0 {
				g.Bullets[i].Active = false
			}
		}
	}

	for i := range g.Enemies {
		if rl.CheckCollisionRecs(g.Player.Rec, g.Enemies[i].Rec) {
			g.Player.Health -= 1
		}
	}
	for i := range g.Bullets {
		for j := range g.Enemies {
			if rl.CheckCollisionRecs(g.Bullets[i].Rec, g.Enemies[j].Rec) {
				// g.Bullets[i].Active = false
				g.Bullets[i].Rec.X = -1000
				g.Enemies[j].Rec.X = float32(rl.GetRandomValue(0, g.ScreenWidth-g.Enemies[i].Rec.ToInt32().Width))
				g.Enemies[j].Rec.Y = float32(rl.GetRandomValue(-g.ScreenHeight*3, -g.ScreenHeight))
			}
		}
	}

	// Endgame scenaria.
	if g.Player.Health <= 0 {
		g.GameActive = false
	}
}

func (g *Game) Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)
	rl.DrawFPS(20, 20)
	// rl.GetFrameTime()

	if g.GameActive {
		rl.DrawRectangleRec(g.Player.Rec, g.Player.Colour)

		for i := range g.Enemies {
			if g.Enemies[i].Alive {
				rl.DrawRectangleRec(g.Enemies[i].Rec, rl.Blue)
				g.Enemies[i].Rec.Y += g.Enemies[i].Speed
			}
		}
		// Draw Bullets
		for _, b := range g.Bullets {
			if b.Active {
				rl.DrawRectangleRec(b.Rec, rl.Orange)
			}
		}
	} else {
		rl.ClearBackground(rl.Black)
		loseText := "You are lose"
		rl.DrawText(loseText, g.ScreenWidth/2-100, 200, 20, rl.Red)
		if rl.IsKeyPressed(rl.KeyEnter) {
			g.Init()
		}
	}
	rl.EndDrawing()
}
