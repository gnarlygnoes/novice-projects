package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	BulletSpeed     = 30
	NumStars        = 900
	NumBullets      = 50
	MaxEnemyBullets = 50
	ScreenWidth     = 1000
	ScreenHeight    = 1200
	EnemyWidth      = 60
	EnemyHeight     = 80
	EnemyGridY      = 5
	EnemyGridX      = 10
	NumDefences     = 6
)

type Game struct {
	GameActive       bool
	MovingRight      bool
	EnemyHBox        rl.Vector2
	EnemiesAlive     int
	PlayerWon        bool
	playerScore      int
	ToggleFrameLimit bool
	bulletTimer      int32
	EnemyShoot       bool
	enemySpeed       float32

	Player      Player
	Bullets     [NumBullets]Bullet
	Enemies     [EnemyGridY][EnemyGridX]Enemy
	EnemyBullet EnemyBullet
	stars       [NumStars]Star
	Defence     [NumDefences]Defence
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
	Shooting  bool
}

type EnemyBullet struct {
	Rec   rl.Rectangle
	Shoot bool
}

type Star struct {
	x, y   int32
	w, h   float32
	Colour rl.Color
}

type Defence struct {
	Rec    rl.Rectangle
	Colour rl.Color
	Health int32
	Active bool
}

func main() {
	game := Game{}
	rl.InitWindow(ScreenWidth, ScreenHeight, "Aliens of Golang")

	game.InitGame()

	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	// rl.GetFPS()

	for !rl.WindowShouldClose() {
		game.Update()

		game.Draw()
	}
}

func (g *Game) InitGame() {
	g.GameActive = true
	g.playerScore = 0
	g.bulletTimer = 4 //rl.GetRandomValue(4, 6)
	g.EnemyShoot = false
	g.enemySpeed = 0.5

	// Initialise player
	g.Player.Rec.Width = 60
	g.Player.Rec.Height = 80
	g.Player.Rec.X = float32(ScreenWidth/2 - int32(g.Player.Rec.Width/2))
	g.Player.Rec.Y = float32(ScreenHeight - int32(g.Player.Rec.Height))
	g.Player.Speed = 10
	g.Player.Health = 3
	g.Player.Colour = rl.Red

	// Initialise stars
	for i := range g.stars {
		g.stars[i] = GenerateStars()
	}

	// Initialise player bullets
	for i := range g.Bullets {
		g.Bullets[i].Rec.Width = 5
		g.Bullets[i].Rec.Height = 20
		g.Bullets[i].Active = false
		g.Bullets[i].Rec.X = -1000
		g.Bullets[i].Rec.Y = 0
	}

	// Initialise enemies
	g.EnemiesAlive = 0
	for i := range g.Enemies {
		for j := range g.Enemies[i] {
			g.Enemies[i][j].HitPoints = EnemyGridY - i
			g.Enemies[i][j].Alive = true
			g.Enemies[i][j].Rec.Width = EnemyWidth
			g.Enemies[i][j].Rec.Height = EnemyHeight
			g.Enemies[i][j].Rec.X = 11 + (EnemyWidth+15)*float32(j)
			g.Enemies[i][j].Colour = rl.Green
			g.Enemies[i][j].Rec.Y = float32(i) * (EnemyHeight + 30)
			g.Enemies[i][j].Shooting = true
			g.EnemiesAlive++
		}
	}
	g.EnemyHBox.X = 11
	g.EnemyHBox.Y = (EnemyWidth + 15) * EnemyGridX

	// Initialise enemy bullets
	g.EnemyBullet.Shoot = false
	g.EnemyBullet.Rec.Width = 5
	g.EnemyBullet.Rec.Height = 20
	g.EnemyBullet.Rec.X = -2000
	g.EnemyBullet.Rec.Y = -1000

	// Initialise defences
	for i := range g.Defence {
		g.Defence[i].Rec.Width = ScreenWidth / 12
		g.Defence[i].Rec.Height = 100
		g.Defence[i].Rec.X = 50 + ScreenWidth/(6/float32(i)) + float32(5*i)
		g.Defence[i].Rec.Y = ScreenHeight - 200 - g.Defence[i].Rec.Height
		g.Defence[i].Health = 10
		g.Defence[i].Active = true
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
	// if rl.IsKeyDown(rl.KeyDown) && g.Player.Rec.Y < float32(ScreenHeight)-g.Player.Rec.Height {
	// 	g.Player.Rec.Y += g.Player.Speed
	// }
	// if rl.IsKeyDown(rl.KeyUp) && g.Player.Rec.Y > 0.0 {
	// 	g.Player.Rec.Y -= g.Player.Speed
	// }
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
			g.Bullets[i].Rec.Y -= BulletSpeed
			if g.Bullets[i].Rec.Y < 0 {
				g.Bullets[i].Active = false
			}
		}
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
				}
			}
		}

		for j := range g.Defence {
			if rl.CheckCollisionRecs(g.Bullets[i].Rec, g.Defence[j].Rec) {
				g.Bullets[i].Active = false
				g.Bullets[i].Rec.Y = -1000
				g.Defence[j].Health--
			}
		}

		if rl.CheckCollisionRecs(g.Bullets[i].Rec, g.EnemyBullet.Rec) {
			g.Bullets[i].Active = false
			g.Bullets[i].Rec.X = -1000
			g.EnemyBullet.Rec.X = -3000
			// g.EnemyBullet.Shoot = false
		}
	}

	if rl.CheckCollisionRecs(g.EnemyBullet.Rec, g.Player.Rec) {
		g.EnemyBullet.Shoot = false
		g.EnemyBullet.Rec.Y = -2000
		g.Player.Health--
	}
	for i := range g.Defence {
		if rl.CheckCollisionRecs(g.EnemyBullet.Rec, g.Defence[i].Rec) {
			g.EnemyBullet.Shoot = false
			g.EnemyBullet.Rec.Y = -2000
			g.Defence[i].Health--
		}
	}
}

func (g *Game) EnemyBehaviour() {
	if g.EnemyHBox.Y >= ScreenWidth {
		g.MovingRight = false
		for i := range g.Enemies {
			for j := range g.Enemies[i] {
				g.Enemies[i][j].Rec.Y += 40
			}
		}
	}
	if g.EnemyHBox.X <= 10 {
		g.MovingRight = true
		for i := range g.Enemies {
			for j := range g.Enemies[i] {
				g.Enemies[i][j].Rec.Y += 40
			}
		}
	}

	if g.MovingRight {
		for i := range g.Enemies {
			for j := range g.Enemies[i] {
				g.Enemies[i][j].Rec.X += float32(g.enemySpeed)
			}
		}
		g.EnemyHBox.X += g.enemySpeed
		g.EnemyHBox.Y += g.enemySpeed
	}
	if !g.MovingRight {
		for i := range g.Enemies {
			for j := range g.Enemies[i] {
				g.Enemies[i][j].Rec.X -= g.enemySpeed
			}
		}
		g.EnemyHBox.X -= g.enemySpeed
		g.EnemyHBox.Y -= g.enemySpeed
	}

	for i := range g.Enemies {
		for j := range g.Enemies[i] {
			if g.Enemies[i][j].Alive {
				if g.Enemies[i][j].HitPoints <= 0 {
					g.Enemies[i][j].Alive = false
					g.EnemiesAlive--
					g.Enemies[i][j].Rec.X = -1000
					g.playerScore += 100
				}
			}
		}
	}
}

func (g *Game) EnemyGoBoom() {
	if rl.GetTime() >= 3 {
		if int32(rl.GetTime())%g.bulletTimer == 0 && !g.EnemyBullet.Shoot {
			shooter := rl.GetRandomValue(0, EnemyGridX-1)
			for i := EnemyGridY - 1; i >= 0; i-- {
				if g.Enemies[i][shooter].Alive {
					g.EnemyBullet.Shoot = true
					g.EnemyBullet.Rec.X = g.Enemies[i][shooter].Rec.X + g.Enemies[i][shooter].Rec.Width/2
					g.EnemyBullet.Rec.Y = g.Enemies[i][shooter].Rec.Y + g.Enemies[i][shooter].Rec.Height
					break
				}
			}
		}
	}

	if g.EnemyBullet.Shoot {
		g.EnemyBullet.Rec.Y += BulletSpeed
		if g.EnemyBullet.Rec.Y > ScreenHeight {
			g.EnemyBullet.Shoot = false
		}
	}
}

func (g *Game) DefenceBehaviour() {
	for i := range g.Defence {
		if g.Defence[i].Health <= 0 {
			g.Defence[i].Active = false
			g.Defence[i].Rec.X = -3000
		}
	}
}

func (g *Game) Update() {
	g.HandleInputs()
	g.BulletLogic()
	g.HandleCollisions()
	g.EnemyBehaviour()
	g.EnemyGoBoom()
	g.DefenceBehaviour()

	switch g.playerScore {
	case 1000:
		g.enemySpeed = .75
	case 2000:
		g.enemySpeed = 1
		g.bulletTimer = 3
	case 3000:
		g.enemySpeed = 1.5
		g.bulletTimer = 2
	case 4000:
		g.enemySpeed = 3
		g.bulletTimer = 1
	case 4900:
		g.enemySpeed = 10
	}
	// Endgame scenaria.
	if g.Player.Health <= 0 {
		g.GameActive = false
	}
	if g.EnemiesAlive <= 0 {
		g.GameActive = false
		g.PlayerWon = true
	}
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

	if g.GameActive {

		rl.DrawRectangleRec(g.Player.Rec, g.Player.Colour)

		//Draw defences
		for _, d := range g.Defence {
			if d.Active {
				rl.DrawRectangleRec(d.Rec, rl.Gray)
			}
		}

		// Draw bullets
		for _, b := range g.Bullets {
			if b.Active {
				rl.DrawRectangleRec(b.Rec, rl.Orange)
			}
		}

		// Draw enemies
		for i := range g.Enemies {
			for _, e := range g.Enemies[i] {
				if e.Alive {
					rl.DrawRectangleRec(e.Rec, e.Colour)
				}
			}
		}

		// Draw enemy bullets
		if g.EnemyBullet.Shoot {
			rl.DrawRectangleRec(g.EnemyBullet.Rec, rl.Blue)
		}

		rl.DrawText(fmt.Sprint("Health: ", g.Player.Health), 20, ScreenHeight-40, 30, rl.White)
		rl.DrawText(fmt.Sprint("Score: ", g.playerScore), ScreenWidth-200, ScreenHeight-40, 30, rl.White)
	} else if !g.GameActive && g.PlayerWon {
		rl.ClearBackground(rl.Black)
		text := "YOU'RE WINNER ! OMG! \n\n\tSo proud of u."
		rl.DrawText(text, ScreenWidth/2-100, 200, 20, rl.Green)
		if rl.IsKeyPressed(rl.KeyEnter) {
			g.InitGame()
		}
	} else {
		text := "You are lose. Hit enter to start again rofl."
		rl.DrawText(text, ScreenWidth/2-450, 200, 40, rl.Red)
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
	c.A = 255

	var star Star
	star.x = rl.GetRandomValue(0, int32(rl.GetScreenWidth()))
	star.y = rl.GetRandomValue(0, int32(rl.GetScreenHeight()))
	star.w = float32(rl.GetRandomValue(1, 5)) / 1.3
	star.h = star.w
	star.Colour = c

	return star
}
