package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	BulletSpeed     = 30
	NumStars        = 900
	NumBullets      = 50
	MaxEnemyBullets = 500
	EnemySpeed      = 0.5
	ScreenWidth     = 1600
	ScreenHeight    = 1200
	EnemyWidth      = 60
	EnemyHeight     = 80
	EnemyGridY      = 5
	EnemyGridX      = (ScreenWidth / (EnemyWidth * 2)) - 1
	NumDefences     = 6
)

type Game struct {
	GameActive   bool
	PlayerScore  int
	MovingRight  bool
	EnemyBox     rl.Vector2
	EnemiesAlive int
	PlayerWon    bool

	Player       Player
	Bullets      [NumBullets]Bullet
	Enemies      [EnemyGridY][EnemyGridX]Enemy
	EnemyBullets [MaxEnemyBullets]EnemyBullet
	stars        [NumStars]Star
	Defence      [NumDefences]Defence
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
	Rec         rl.Rectangle
	Colour      rl.Color
	Alive       bool
	HitPoints   int
	Shooting    bool
	BulletTimer int32
}

type EnemyBullet struct {
	Rec    rl.Rectangle
	Active bool
}

type Star struct {
	x, y   int32
	w, h   float32
	Colour rl.Color
}

type Defence struct {
	Rec    rl.Rectangle
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
	// var active bool = true
	for !rl.WindowShouldClose() {
		game.Update()
		// x := rl.GetRandomValue(4, 8)
		// if int32(rl.GetTime())%x == 0 && active {

		// 	active = false
		// }
		// if int32(rl.GetTime())%x == 1 {
		// 	active = true
		// }
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

	// Initialise enemies
	g.EnemiesAlive = 0
	for row := range EnemyGridY {
		for i := range EnemyGridX {
			g.Enemies[row][i].HitPoints = EnemyGridY - row
			g.Enemies[row][i].Alive = true
			g.Enemies[row][i].Rec.Width = EnemyWidth
			g.Enemies[row][i].Rec.Height = EnemyHeight
			g.Enemies[row][i].Rec.X = 2 * EnemyWidth * float32(i)
			g.Enemies[row][i].Colour = rl.Green
			g.Enemies[row][i].Rec.Y = float32(row) * EnemyHeight * 2
			g.Enemies[row][i].Shooting = true
			g.EnemiesAlive++
		}
	}

	g.EnemyBox.X = 0
	g.EnemyBox.Y = 2 * EnemyWidth * EnemyGridX

	// Initialise their shoot rate
	// var timer int32
	for i := range g.Enemies {
		for j := range g.Enemies[i] {
			if g.Enemies[i][j].Shooting {
				g.Enemies[i][j].BulletTimer = rl.GetRandomValue(1, 8)
			}
		}
	}

	// Initialize enemy bullets as inactive
	for i := range g.EnemyBullets {
		g.EnemyBullets[i].Active = false
		g.EnemyBullets[i].Rec.Width = 5
		g.EnemyBullets[i].Rec.Height = 20
	}

	// Initialise defences
	for i := range g.Defence {
		g.Defence[i].Rec.Width = ScreenWidth / 12
		g.Defence[i].Rec.Height = 100
		g.Defence[i].Rec.X = 50 + ScreenWidth/(6/float32(i)) + float32(5*i)
		g.Defence[i].Rec.Y = ScreenHeight - 200 - g.Defence[i].Rec.Height
		// d.Colour = rl.Gray
	}
}

func (g *Game) HandleInputs() {
	if rl.IsKeyDown(rl.KeyRight) && g.Player.Rec.X < float32(ScreenWidth)-g.Player.Rec.Width {
		g.Player.Rec.X += g.Player.Speed
	}
	if rl.IsKeyDown(rl.KeyLeft) && g.Player.Rec.X > 0.0 {
		g.Player.Rec.X -= g.Player.Speed
	}
	if rl.IsKeyDown(rl.KeyDown) && g.Player.Rec.Y < float32(ScreenHeight)-g.Player.Rec.Height {
		g.Player.Rec.Y += g.Player.Speed
	}
	if rl.IsKeyDown(rl.KeyUp) && g.Player.Rec.Y > 0.0 {
		g.Player.Rec.Y -= g.Player.Speed
	}
	if rl.IsKeyDown(rl.KeySpace) {
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
	for i := range EnemyGridY {
		for j := range EnemyGridX {
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
					g.Bullets[i].Rec.X = -1000
					// g.Enemies[j][k].Rec.X = -2000
					g.PlayerScore++
				}
			}
		}
	}
	for i := range g.Bullets {
		for j := range g.Defence {
			if rl.CheckCollisionRecs(g.Bullets[i].Rec, g.Defence[j].Rec) {
				g.Bullets[i].Active = false
				g.Bullets[i].Rec.Y = -1000
			}
		}
	}
	for i := range g.Defence {
		if rl.CheckCollisionRecs(g.Player.Rec, g.Defence[i].Rec) {
		}
	}
}

func (g *Game) EnemyBehaviour() {
	if g.EnemyBox.Y >= ScreenWidth+EnemyWidth {
		g.MovingRight = false
	}
	if g.EnemyBox.X <= 10 {
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
			if g.Enemies[i][j].Alive {
				if g.Enemies[i][j].HitPoints <= 0 {
					g.Enemies[i][j].Alive = false
					g.EnemiesAlive--
					// fmt.Println(("Enemy got rekt"))
					g.Enemies[i][j].Rec.X = -1000
				}
			}
		}
	}
}

func (g *Game) EnemyGoBoom() {
	for i := range g.Enemies {
		for j := range g.Enemies[i] {
			if g.Enemies[i][j].Shooting {
				if rl.GetTime() >= 1 {
					if int32(rl.GetTime())%g.Enemies[i][j].BulletTimer == 0 {
						for r := range g.EnemyBullets {
							g.EnemyBullets[r].Active = true
							// g.EnemyBullets[i].Rec.Y -= 10
							g.EnemyBullets[r].Rec.X = g.Enemies[i][j].Rec.X + (g.Enemies[i][j].Rec.Width / 2)
							g.EnemyBullets[r].Rec.Y = g.Enemies[i][j].Rec.Y + g.Enemies[i][j].Rec.Height
						}
					}
				}
			}
		}
	}
	for i := range g.EnemyBullets {
		g.EnemyBullets[i].Rec.Y += 10
		if g.EnemyBullets[i].Rec.Y >= ScreenHeight {
			g.EnemyBullets[i].Active = false
		}
	}
}

func (g *Game) Update() {
	g.HandleInputs()
	g.BulletLogic()
	g.HandleCollisions()
	g.EnemyBehaviour()
	g.EnemyGoBoom()

	// Endgame scenaria.
	if g.Player.Health <= 0 {
		g.GameActive = false
	}
	if g.EnemiesAlive <= 0 {
		g.GameActive = false
		g.PlayerWon = true
	}
	// if g.PlayerScore >= 100 {
	// 	g.GameActive = false
	// }
}

func (g *Game) Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)
	// rl.GetFrameTime()

	// Draw a beautiful starrry canopy
	for i := range g.stars {
		rl.DrawRectangle(g.stars[i].x,
			g.stars[i].y,
			int32(g.stars[i].w),
			int32(g.stars[i].h),
			g.stars[i].Colour)
	}

	/* Start Screen:
	You is like a Spartan or something and you is being attacked and must use the defences
	to protecc yourself rofl.

	Or: You is being attacked by aliens. You has like a space barricade or something and can
	use it to defend yourself from alien goo-bullets.
	*/

	// var timeTaken int32
	// if g.GameActive {
	// 	timeTaken = int32(rl.GetTime())
	// }

	if g.GameActive {
		rl.DrawRectangleRec(g.Player.Rec, g.Player.Colour)

		//Draw Defences
		for i := range g.Defence {
			rl.DrawRectangleRec(g.Defence[i].Rec, rl.Gray)
		}

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

		// Draw enemy bullets
		for i := range g.EnemyBullets {
			if g.EnemyBullets[i].Active {
				rl.DrawRectangleRec(g.EnemyBullets[i].Rec, rl.Blue)
			}
		}
	} else if !g.GameActive && g.PlayerWon {
		rl.ClearBackground(rl.Black)
		// if g.EnemiesAlive <= 0 {
		text := "YOU'RE WINNER ! OMG! \n\n\tSo proud of u."
		// var text1 int32 = int32(timeTaken)
		// text2 := " seconds. So proud of u."
		rl.DrawText(text, ScreenWidth/2-100, 200, 20, rl.Green)
	} else {
		text := "You are lose. Hit enter to start again rofl."
		rl.DrawText(text, ScreenWidth/2-450, 200, 40, rl.Red)
		// rl.DrawText(text, 50, 100, 40, rl.Red)
	}
	if rl.IsKeyPressed(rl.KeyEnter) {
		g.InitGame()
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
