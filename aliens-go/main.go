package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	BulletSpeed   = 30
	NumStars      = 900
	NumBullets    = 50
	EnemySpeed    = 0.5
	ScreenWidth   = 1600
	ScreenHeight  = 1200
	EnemiesWidth  = 60
	EnemiesHeight = 80
	EnemiesNumY   = 5
	EnemiesNumX   = (ScreenWidth / (EnemiesWidth * 2)) - 1
)

type Game struct {
	GameActive  bool
	PlayerScore int
	MovingRight bool
	EnemyBox    rl.Vector2

	Player  Player
	Bullets [NumBullets]Bullet
	Enemies [EnemiesNumY][EnemiesNumX]Enemy
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
	Rec       rl.Rectangle
	Colour    rl.Color
	Alive     bool
	HitPoints int
}

type Star struct {
	x, y   int32
	w, h   float32
	Colour rl.Color
}

func main() {
	game := Game{}
	rl.InitWindow(ScreenWidth, ScreenHeight, "Aliens of Golang")

	game.InitGame()

	// defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	// rl.GetFPS()

	// var dT float32 = rl.GetFrameTime()
	// gamespeed := dT * float32(rl.GetFPS())

	for !rl.WindowShouldClose() {
		game.Update()

		game.Draw()
	}
}

func (g *Game) InitGame() {
	g.GameActive = true
	g.PlayerScore = 0

	// Initialise player
	g.Player.Rec.Width = 60
	g.Player.Rec.Height = 80
	g.Player.Rec.X = float32(ScreenWidth/2 - int32(g.Player.Rec.Width/2))
	g.Player.Rec.Y = float32(ScreenHeight - int32(g.Player.Rec.Height))
	g.Player.Speed = 10
	g.Player.Health = 1
	g.Player.Colour = rl.Red

	// Initialise stars
	for i := range g.stars {
		g.stars[i] = GenerateStars()
	}

	// Initialise bullets
	for i := range g.Bullets {
		g.Bullets[i].Rec.Width = 5
		g.Bullets[i].Rec.Height = 20
		g.Bullets[i].Active = false
		g.Bullets[i].Rec.X = -1000
		g.Bullets[i].Rec.Y = 0
	}

	for row := range EnemiesNumY {
		for i := range EnemiesNumX {
			g.Enemies[row][i].HitPoints = EnemiesNumY - row
			g.Enemies[row][i].Alive = true
			g.Enemies[row][i].Rec.Width = EnemiesWidth
			g.Enemies[row][i].Rec.Height = EnemiesHeight
			g.Enemies[row][i].Rec.X = 2 * EnemiesWidth * float32(i)
			g.Enemies[row][i].Colour = rl.Green
			g.Enemies[row][i].Rec.Y = float32(row) * EnemiesHeight * 2
		}
	}

	g.EnemyBox.X = 0
	g.EnemyBox.Y = 2 * EnemiesWidth * EnemiesNumX
}

func (g *Game) HandleInputs() {
	if rl.IsKeyDown(rl.KeyRight) && g.Player.Rec.X < float32(ScreenWidth)-g.Player.Rec.Width {
		g.Player.Rec.X += g.Player.Speed
	}
	if rl.IsKeyDown(rl.KeyLeft) && g.Player.Rec.X > 0.0 {
		g.Player.Rec.X -= g.Player.Speed
	}
	// if rl.IsKeyDown(rl.KeyDown) && g.Player.Rec.Y < float32(ScreenHeight)-g.Player.Rec.Height {
	// 	g.Player.Rec.Y += g.Player.Speed
	// }
	// if rl.IsKeyDown(rl.KeyUp) && g.Player.Rec.Y > 0.0 {
	// 	g.Player.Rec.Y -= g.Player.Speed
	// }
	if rl.IsKeyPressed(rl.KeySpace) {
		g.Shoot()
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
			g.Bullets[i].Rec.Y -= BulletSpeed
			if g.Bullets[i].Rec.Y < 0 {
				g.Bullets[i].Active = false
			}
		}
		// i %= NumBullets
	}
}

func (g *Game) HandleCollisions() {
	for i := range EnemiesNumY {
		for j := range EnemiesNumX {
			if rl.CheckCollisionRecs(g.Player.Rec, g.Enemies[i][j].Rec) {
				g.Player.Health -= 1
				g.Enemies[i][j].HitPoints--
			}
		}
	}
	for i := range g.Bullets {
		for j := range g.Enemies {
			for k := range g.Enemies[j] {
				if rl.CheckCollisionRecs(g.Bullets[i].Rec, g.Enemies[j][k].Rec) {
					g.Bullets[i].Active = false
					g.Enemies[j][k].HitPoints--
					g.Bullets[i].Rec.Y = -1000
					// g.Enemies[j][k].Rec.X = -2000
					g.PlayerScore++
				}
			}
		}
	}
}

func (g *Game) EnemyBehaviour() {
	if g.EnemyBox.Y >= ScreenWidth {
		g.MovingRight = false
	}
	// if g.Enemies[i][j].Rec.X <= 50 {
	if g.EnemyBox.X <= 50 {
		g.MovingRight = true
	}

	for i := range g.Enemies {
		for j := range g.Enemies[i] {
			g.Enemies[i][j].Rec.Y += 0
		}
	}
	if g.MovingRight {
		for i := range g.Enemies {
			for j := range g.Enemies[i] {
				g.Enemies[i][j].Rec.X += EnemySpeed
			}
		}
		g.EnemyBox.X += EnemySpeed
		g.EnemyBox.Y += EnemySpeed
	}
	if !g.MovingRight {
		for i := range g.Enemies {
			for j := range g.Enemies[i] {
				g.Enemies[i][j].Rec.X -= EnemySpeed
			}
		}
		g.EnemyBox.X -= EnemySpeed
		g.EnemyBox.Y -= EnemySpeed
	}

	for i := range g.Enemies {
		for j := range g.Enemies[i] {
			if g.Enemies[i][j].HitPoints <= 0 {
				g.Enemies[i][j].Alive = false
				g.Enemies[i][j].Rec.X = -1000
			}
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
	// if g.PlayerScore >= 100 {
	// 	g.GameActive = false
	// }
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

		// Draw Enemies
		for row := range g.Enemies {
			for i := range g.Enemies[row] {
				if g.Enemies[row][i].Alive {
					rl.DrawRectangleRec(g.Enemies[row][i].Rec, g.Enemies[row][i].Colour)
				}
			}
		}
	} else {
		rl.ClearBackground(rl.Black)
		// if g.PlayerScore >= 10 {
		// 	text := "YOUR WINNER!!! OMG!"
		// 	rl.DrawText(text, ScreenWidth/2-100, 200, 20, rl.Green)
		// } else {
		text := "You are lose. Hit enter to start again rofl."
		rl.DrawText(text, ScreenWidth/2-250, 200, 20, rl.Red)
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
