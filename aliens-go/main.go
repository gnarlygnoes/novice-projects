package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	BulletSpeed = 30
	NumStars    = 9000
	EnemySpeed  = 2
)

type Game struct {
	ScreenWidth  int32
	ScreenHeight int32
	GameActive   bool
	PlayerScore  int

	Player  Player
	Bullets [32]Bullet
	Enemies [13][13]Enemy
	stars   [NumStars]Star
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
	Colour rl.Color
	Alive  bool
}

type Star struct {
	x, y   int32
	w, h   float32
	Colour rl.Color
}

// type RingBuffer struct {
// 	data       []*Data
// 	size       int
// 	lastInsert int
// 	nextRead   int
// }

// type Data struct {
// 	Value string
// }

// func NewRingBuffer(size int) *RingBuffer {
// 	return &RingBuffer{
// 		data:       make([]*Data, size),
// 		size:       size,
// 		lastInsert: -1,
// 	}
// }

func main() {
	game := Game{}
	game.InitGame()

	defer rl.CloseWindow()

	rl.SetTargetFPS(120)

	rl.InitWindow(game.ScreenWidth, game.ScreenHeight, "Aliens of Golang")
	for !rl.WindowShouldClose() {
		game.Update()

		game.Draw()
	}
}

func (g *Game) InitGame() {
	g.ScreenWidth = 1600
	g.ScreenHeight = 1200
	g.GameActive = true
	g.PlayerScore = 0

	// Initialise player
	g.Player.Rec.Width = 60
	g.Player.Rec.Height = 80
	g.Player.Rec.X = float32(g.ScreenWidth/2 - int32(g.Player.Rec.Width/2))
	g.Player.Rec.Y = float32(g.ScreenHeight - int32(g.Player.Rec.Height))
	g.Player.Speed = 10
	g.Player.Health = 1
	g.Player.Colour = rl.Red

	// Initialise stars
	for i := range g.stars {
		g.stars[i] = GenerateStars()
	}

	// Initialise bullets
	// b := Bullet{}
	// b.Rec.Width = 5
	// b.Rec.Height = 20
	// b.Active = false
	for i := range g.Bullets {
		g.Bullets[i].Rec.Width = 5
		g.Bullets[i].Rec.Height = 20
		g.Bullets[i].Active = false
		g.Bullets[i].Rec.X = -1000
		g.Bullets[i].Rec.Y = 0
	}

	// Initialise enemies
	e := Enemy{}
	e.Rec.Width = 60
	e.Rec.Height = 80
	var enemiesVertical int = 5
	var enemiesHorizontal int32 = g.ScreenWidth / (2 * int32(e.Rec.Width))

	for row := range enemiesVertical {
		for i := range enemiesHorizontal {
			g.Enemies[row][i].Alive = true
			g.Enemies[row][i].Rec.Width = e.Rec.Width
			g.Enemies[row][i].Rec.Height = e.Rec.Height
			g.Enemies[row][i].Rec.X = g.Enemies[row][i].Rec.Width + 2*g.Enemies[row][i].Rec.Width*float32(i)
			g.Enemies[row][i].Colour = rl.Blue
			g.Enemies[row][i].Rec.Y = float32(row) * g.Enemies[row][i].Rec.Height * 2
		}
	}
}

func (g *Game) HandleInputs() {
	if rl.IsKeyDown(rl.KeyRight) && g.Player.Rec.X < float32(g.ScreenWidth)-g.Player.Rec.Width {
		g.Player.Rec.X += g.Player.Speed
	}
	if rl.IsKeyDown(rl.KeyLeft) && g.Player.Rec.X > 0.0 {
		g.Player.Rec.X -= g.Player.Speed
	}
	if rl.IsKeyDown(rl.KeyDown) && g.Player.Rec.Y < float32(g.ScreenHeight)-g.Player.Rec.Height {
		g.Player.Rec.Y += g.Player.Speed
	}
	if rl.IsKeyDown(rl.KeyUp) && g.Player.Rec.Y > 0.0 {
		g.Player.Rec.Y -= g.Player.Speed
	}
	if rl.IsKeyDown(rl.KeySpace) {
		for i := range g.Bullets {
			g.Bullets[i].Active = true
			g.Bullets[i].Rec.X = g.Player.Rec.X + g.Player.Rec.Width/2 - g.Bullets[i].Rec.Width/2
			g.Bullets[i].Rec.Y = g.Player.Rec.Y - g.Bullets[i].Rec.Height
		}
	}
}

// func (g *Game) Shoot() Bullet {
// 	b := Bullet{}
// 	b.Rec.X = g.Player.Rec.X + g.Player.Rec.Width/2 - b.Rec.Width/2
// 	b.Rec.Y = g.Player.Rec.Y - b.Rec.Height
// 	b.Rec.Width = 5
// 	b.Rec.Height = 20
// 	b.Active = true
// g.Bullets = append(g.Bullets, b)

// 	// for i := range g.Bullets {
// 	// 	g.Bullets[i].Active = true
// 	// 	g.Bullets[i].Rec.X = g.Player.Rec.X + g.Player.Rec.Width/2 - g.Bullets[i].Rec.Width/2
// 	// 	g.Bullets[i].Rec.Y = g.Player.Rec.Y - g.Bullets[i].Rec.Height
// 	// 	g.Bullets[i].Rec.Width = 5
// 	// 	g.Bullets[i].Rec.Height = 20
// 	// 	g.Bullets[i].Rec.Y -= BulletSpeed
// 	// }
// 	return b
// }

func (g *Game) BulletLogic() {
	for i := range g.Bullets {
		if g.Bullets[i].Active {
			g.Bullets[i].Rec.Y -= BulletSpeed
			if g.Bullets[i].Rec.Y < 0 {
				g.Bullets[i].Active = false
			}
		}
	}
}

func (g *Game) HandleCollisions() {
	for i := range g.Enemies {
		for j := range g.Enemies {
			if rl.CheckCollisionRecs(g.Player.Rec, g.Enemies[i][j].Rec) {
				g.Player.Health -= 1
			}
		}
	}
	for i := range g.Bullets {
		for j := range g.Enemies {
			for k := range g.Enemies[j] {
				if rl.CheckCollisionRecs(g.Bullets[i].Rec, g.Enemies[j][k].Rec) {
					g.Bullets[i].Active = false
					g.Enemies[j][k].Alive = false
					g.Bullets[i].Rec.X = -1000
					g.Enemies[j][k].Rec.X = -1000
					g.PlayerScore++
				}
			}
		}
	}
}

func (g *Game) EnemyBehaviour() {
	for i := range g.Enemies {
		for j := range g.Enemies[i] {
			g.Enemies[i][j].Rec.Y += EnemySpeed
		}
	}
}

func (g *Game) Update() {
	g.HandleInputs()
	g.BulletLogic()
	g.HandleCollisions()
	g.EnemyBehaviour()

	// Endgame scenaria.
	if g.Player.Health <= 0 {
		g.GameActive = false
	}
}

func (g *Game) Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)
	// rl.GetFrameTime()
	for i := range g.stars {
		rl.DrawRectangle(g.stars[i].x,
			g.stars[i].y,
			int32(g.stars[i].w),
			int32(g.stars[i].h),
			g.stars[i].Colour)
	}
	if g.GameActive {
		rl.DrawRectangleRec(g.Player.Rec, g.Player.Colour)

		// Draw a beautiful starrry canopy

		// Draw Bullets
		for _, b := range g.Bullets {
			if b.Active {
				rl.DrawRectangleRec(b.Rec, rl.Orange)
			}
		}
		for row := range g.Enemies {
			for i := range g.Enemies {
				if g.Enemies[row][i].Alive {
					rl.DrawRectangleRec(g.Enemies[row][i].Rec, g.Enemies[row][i].Colour)
				}
			}
		}
	} else {
		// rl.ClearBackground(rl.Black)
		// if g.PlayerScore >= 10 {
		// 	text := "YOUR WINNER!!! OMG!"
		// 	rl.DrawText(text, g.ScreenWidth/2-100, 200, 20, rl.Green)
		// } else {
		text := "You are lose. Hit enter to start again rofl."
		rl.DrawText(text, g.ScreenWidth/2-250, 200, 20, rl.Red)
		// }
		if rl.IsKeyPressed(rl.KeyEnter) {
			g.InitGame()
		}
	}

	rl.DrawFPS(20, 20)
	rl.EndDrawing()
}

func GenerateStars() Star {
	rVal := rl.GetRandomValue(100, 255)
	gVal := rl.GetRandomValue(100, 255)
	bVal := rl.GetRandomValue(100, 255)

	var c rl.Color
	c.R = uint8(rVal)
	c.G = uint8(gVal)
	c.B = uint8(bVal)
	c.A = 200

	var star Star
	star.x = rl.GetRandomValue(0, int32(rl.GetScreenWidth()))
	star.y = rl.GetRandomValue(0, int32(rl.GetScreenHeight()))
	star.w = float32(rl.GetRandomValue(1, 5)) / 1.3
	star.h = star.w
	star.Colour = c

	return star
}
