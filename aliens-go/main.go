package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	NumStars            = 900
	NumBullets          = 50
	ScreenWidth         = 1000
	ScreenHeight        = 1200
	AnimTime            = 0.5
	EnemyWidth          = 75
	EnemyHeight         = 75
	EnemyGridY          = 5
	EnemyGridX          = 10
	NumDefences         = 4
	DefenceHeight       = 70
	DefenceWidth        = 140
	DefencePositionY    = ScreenHeight - 200 - DefenceHeight
	BulletDisplacement  = -9000
	EnemyDisplacement   = -6000
	EBulletDisplacement = -12000
)

type Game struct {
	gameActive   bool
	movingRight  bool
	enemyHBox    rl.Vector2
	enemiesAlive int
	playerWon    bool
	playerScore  int
	bulletTimer  int32
	enemyXLen    int
	enemyXMin    int
	texSegmentH  int32
	texSegmentV  int32
	dt           float32
	runningTime  float32
	gamePaused   bool

	enemySpeed  float32
	playerSpeed float32
	bulletSpeed float32

	gameTexture rl.Texture2D
	clouds      rl.Texture2D
	// buildingTexture rl.Texture2D

	frame int32

	Player      Player
	Bullets     [NumBullets]Bullet
	Enemies     [EnemyGridY][EnemyGridX]Enemy
	EnemyBullet EnemyBullet
	Stars       [NumStars]Star
	Defence     [NumDefences]Defence
}

type Player struct {
	RecIn  rl.Rectangle
	Rec    rl.Rectangle
	Pos    rl.Vector2
	Colour rl.Color
	Health int32
}

type Bullet struct {
	Rec    rl.Rectangle
	Active bool
}

type Enemy struct {
	RecIn     rl.Rectangle
	Rec       rl.Rectangle
	Pos       rl.Vector2
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
	RecIn  rl.Rectangle
	Rec    rl.Rectangle
	Pos    rl.Vector2
	Colour rl.Color
	Health int32
	Active bool
}

func main() {
	// game := Game{
	// 	playerSpeed: 1000,
	// 	enemySpeed:  20,
	// 	bulletSpeed: 2000,
	// 	runningTime: 0,
	// 	playerScore: 0,
	// 	gameActive:  true,
	// 	bulletTimer: 5,
	// }

	rl.InitWindow(ScreenWidth, ScreenHeight, "Aliens Go Home!")

	game := InitGame()

	defer rl.CloseWindow()

	rl.SetTargetFPS(120)

	for !rl.WindowShouldClose() {
		game.Update()

		game.Draw()
	}
}

func InitGame() *Game {
	g := &Game{}
	g.gamePaused = false
	g.playerSpeed = 1000
	g.enemySpeed = 20
	g.bulletSpeed = 2000
	g.runningTime = 0
	g.playerScore = 0
	g.gameActive = true
	g.bulletTimer = 5

	// Initialise stars
	for i := range g.Stars {
		g.Stars[i] = GenerateStars()
	}

	// Initialise clouds and buildings
	// g.clouds = GenerateTexture(ScreenWidth, ScreenHeight, 5.5)
	// g.buildingTexture = rl.LoadTexture("img/SpaceInvaders_BackgroundBuildings.png")

	g.frame = 0

	g.gameTexture = rl.LoadTexture("img/SpaceInvaders.png")
	g.texSegmentH = g.gameTexture.Width / 7
	g.texSegmentV = g.gameTexture.Height / 5

	g.Player.RecIn.Width = float32(g.gameTexture.Width) / 7
	g.Player.RecIn.Height = float32(g.gameTexture.Height) / 5
	g.Player.RecIn.X = float32(g.texSegmentH) * 4
	g.Player.RecIn.Y = 0
	g.Player.Rec.Width = 80
	g.Player.Rec.Height = 80
	g.Player.Rec.X = float32(ScreenWidth/2 - int32(g.Player.Rec.Width/2))
	g.Player.Rec.Y = float32(ScreenHeight - int32(g.Player.Rec.Height) - 40)
	g.Player.Pos.X = 0
	g.Player.Pos.Y = 0

	// g.playerSpeed = 10
	g.Player.Health = 3
	g.Player.Colour = rl.Red

	// Initialise player bullets
	for i := range g.Bullets {
		g.Bullets[i].Rec.Width = 5
		g.Bullets[i].Rec.Height = 20
		g.Bullets[i].Active = false
		g.Bullets[i].Rec.X = -1000
		g.Bullets[i].Rec.Y = 0
	}

	// Initialise enemies
	g.enemiesAlive = 0
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
			g.enemiesAlive++
		}
	}
	g.enemyHBox.X = 11
	g.enemyHBox.Y = (EnemyWidth + 15) * EnemyGridX
	g.enemyXLen = EnemyGridX - 1
	g.enemyXMin = 0

	// Initialise enemy bullets
	g.EnemyBullet.Shoot = false
	g.EnemyBullet.Rec.Width = 5
	g.EnemyBullet.Rec.Height = 20
	g.EnemyBullet.Rec.X = -2000
	g.EnemyBullet.Rec.Y = -1000

	// Initialise defences
	for i := range g.Defence {
		g.Defence[i].Rec.Width = DefenceWidth
		g.Defence[i].Rec.Height = DefenceHeight
		g.Defence[i].Rec.X = 10 + ScreenWidth/(NumDefences*6) + float32(ScreenWidth/4)*float32(i)
		g.Defence[i].Rec.Y = DefencePositionY
		g.Defence[i].Health = 25
		g.Defence[i].Active = true
	}
	return g
}

func (g *Game) HandleInputs() {
	if rl.IsKeyDown(rl.KeyRight) && g.Player.Rec.X < float32(ScreenWidth)-g.Player.Rec.Width && !g.gamePaused {
		g.movingRight = true
		g.Player.Rec.X += g.playerSpeed * g.dt
	}
	if rl.IsKeyDown(rl.KeyLeft) && g.Player.Rec.X > 0.0 && !g.gamePaused {
		g.Player.Rec.X -= g.playerSpeed * g.dt
	}
	// if rl.IsKeyDown(rl.KeyDown) && g.Player.Rec.Y < float32(ScreenHeight)-g.Player.Rec.Height {
	// 	g.Player.Rec.Y += g.Player.Speed
	// }
	// if rl.IsKeyDown(rl.KeyUp) && g.Player.Rec.Y > 0.0 {
	// 	g.Player.Rec.Y -= g.Player.Speed
	// }
	if rl.IsKeyPressed(rl.KeySpace) && !g.gamePaused {
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
	if rl.IsKeyPressed(rl.KeyP) {
		if !g.gamePaused {
			g.gamePaused = true
		} else {
			g.gamePaused = false
		}
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
					g.Bullets[i].Rec.X = BulletDisplacement
				}
			}
		}

		for j := range g.Defence {
			if rl.CheckCollisionRecs(g.Bullets[i].Rec, g.Defence[j].Rec) {
				g.Bullets[i].Active = false
				g.Bullets[i].Rec.X = BulletDisplacement
				g.Defence[j].Health--
			}
			if rl.CheckCollisionRecs(g.EnemyBullet.Rec, g.Defence[j].Rec) {
				g.EnemyBullet.Shoot = false
				g.EnemyBullet.Rec.X = EBulletDisplacement
				g.Defence[i].Health--
			}
		}

		if rl.CheckCollisionRecs(g.Bullets[i].Rec, g.EnemyBullet.Rec) {
			g.Bullets[i].Active = false
			g.Bullets[i].Rec.X = -1000
			g.EnemyBullet.Rec.X = BulletDisplacement
		}
	}

	if rl.CheckCollisionRecs(g.EnemyBullet.Rec, g.Player.Rec) {
		g.EnemyBullet.Shoot = false
		g.EnemyBullet.Rec.X = EBulletDisplacement
		g.Player.Health--
	}
}

func (g *Game) EnemyBehaviour() {
	for i := range g.Enemies {
		for j := range g.Enemies[i] {
			if g.enemyHBox.Y >= ScreenWidth {
				g.movingRight = false

				if g.Enemies[4][j].Rec.Y < DefencePositionY-DefenceHeight-15 {
					g.Enemies[i][j].Rec.Y += 30
				}

			}
			if g.enemyHBox.X <= 10 {
				g.movingRight = true

				if g.Enemies[4][j].Rec.Y < DefencePositionY-DefenceHeight-15 {
					g.Enemies[i][j].Rec.Y += 30
				}
			}
		}
	}

	if g.movingRight {
		for i := range g.Enemies {
			for j := range g.Enemies[i] {
				g.Enemies[i][j].Rec.X += g.enemySpeed * g.dt
			}
		}
		g.enemyHBox.X += g.enemySpeed * g.dt
		g.enemyHBox.Y += g.enemySpeed * g.dt
	}
	if !g.movingRight {
		for i := range g.Enemies {
			for j := range g.Enemies[i] {
				g.Enemies[i][j].Rec.X -= g.enemySpeed * g.dt
			}
		}
		g.enemyHBox.X -= g.enemySpeed * g.dt
		g.enemyHBox.Y -= g.enemySpeed * g.dt
	}

	eMaxAlive := 0
	eMinAlive := 0
	for i := range g.Enemies {
		for j := range g.Enemies[i] {
			if g.Enemies[i][j].Alive {
				if g.Enemies[i][j].HitPoints <= 0 {
					g.Enemies[i][j].Alive = false
					g.enemiesAlive--
					g.Enemies[i][j].Rec.X = -1000
					g.playerScore += 100
				}
			}
		}

		if g.enemyXLen > 0 {
			if g.Enemies[i][g.enemyXLen].Alive {
				eMaxAlive++
			}
			if eMaxAlive == 0 {
				g.enemyXLen--
				g.enemyHBox.Y -= (EnemyWidth + 15)
			}
		}

		if g.enemyXMin < 10 {
			if g.Enemies[i][g.enemyXMin].Alive {
				eMinAlive++
			}
			if eMinAlive == 0 {
				g.enemyXMin++
				g.enemyHBox.X += (EnemyWidth + 15)
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
		g.EnemyBullet.Rec.Y += g.bulletSpeed * g.dt
		if g.EnemyBullet.Rec.Y > ScreenHeight {
			g.EnemyBullet.Shoot = false
		}
	}
}

func (g *Game) DefenceBehaviour() {
	for i := range g.Defence {
		if g.Defence[i].Health <= 0 {
			g.Defence[i].Active = false
			g.Defence[i].Rec.X = -15000
		}
	}
}

func (g *Game) Update() {
	g.dt = rl.GetFrameTime()

	g.runningTime += g.dt

	g.HandleInputs()

	if g.gameActive && !g.gamePaused {
		g.BulletLogic()
		g.HandleCollisions()
		g.EnemyBehaviour()
		g.EnemyGoBoom()
		g.DefenceBehaviour()

		switch g.playerScore {
		case 1000:
			g.enemySpeed = 30
		case 2000:
			g.enemySpeed = 50
			g.bulletTimer = 3
		case 3000:
			g.enemySpeed = 80
			g.bulletTimer = 2
		case 4000:
			g.enemySpeed = 120
			g.bulletTimer = 1
		case 4900:
			g.enemySpeed = 200
		}
	}
	// Endgame scenaria.
	if g.Player.Health <= 0 {
		g.gameActive = false
	}
	if g.enemiesAlive <= 0 {
		g.gameActive = false
		g.playerWon = true
	}
}

func (g *Game) Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	// Draw a beautiful starrry canopy
	for i := range g.Stars {
		rl.DrawRectangle(g.Stars[i].x,
			g.Stars[i].y,
			int32(g.Stars[i].w),
			int32(g.Stars[i].h),
			g.Stars[i].Colour)
	}

	// Draw pretty clouds
	rl.DrawTexture(g.clouds, 0, 0, rl.Gray)
	// rl.DrawTexturePro(g.buildingTexture)

	/* Start Screen:
	You is like a Spartan or something and you is being attacked and must use the defences
	to protecc yourself rofl.

	Or: You is being attacked by aliens. You has like a space barricade or something and can
	use it to defend yourself from alien goo-bullets.
	*/

	if g.gamePaused {
		rl.DrawText("GAME PAUSED", ScreenWidth/2-200, ScreenHeight/2, 50, rl.White)
	}

	if g.gameActive {
		rl.DrawTexturePro(g.gameTexture, g.Player.RecIn, g.Player.Rec, g.Player.Pos, 0, g.Player.Colour)

		//Draw defences
		for _, d := range g.Defence {
			if d.Active {
				d.RecIn.Width = float32(g.texSegmentH) * 2
				d.RecIn.Height = float32(g.texSegmentV)
				d.RecIn.X = 3 * float32(g.texSegmentH)
				if d.Health > 15 {
					d.RecIn.Y = float32(g.texSegmentV)
				} else if d.Health > 8 {
					d.RecIn.Y = 2 * float32(g.texSegmentV)
				} else {
					d.RecIn.Y = 3 * float32(g.texSegmentV)
				}
				rl.DrawTexturePro(g.gameTexture, d.RecIn, d.Rec, d.Pos, 0, rl.Gray)
			}
		}

		// Draw bullets
		for _, b := range g.Bullets {
			if b.Active {
				rl.DrawRectangleRec(b.Rec, rl.Orange)
			}
		}

		// Draw enemies
		if g.runningTime >= AnimTime {
			g.frame++
			if g.frame > 1 {
				g.frame = 0
			}
			g.runningTime = 0
		}
		var enemyColour rl.Color = rl.Red
		for i := range g.Enemies {
			for _, e := range g.Enemies[i] {
				e.RecIn.Width = float32(g.texSegmentH)
				e.RecIn.Height = float32(g.texSegmentV)
				e.RecIn.X = float32(g.frame) * e.RecIn.Width
				if i == 0 {
					e.RecIn.Y = 0
				} else if i == 1 {
					e.RecIn.Y = float32(g.texSegmentV)
				} else if i == 2 {
					e.RecIn.Y = float32(g.texSegmentV) * 2
					enemyColour = rl.Orange
				} else if i == 3 {
					e.RecIn.Y = float32(g.texSegmentV) * 3
					enemyColour = rl.Yellow
				} else {
					e.RecIn.Y = float32(g.texSegmentV) * 4
					enemyColour = rl.Green
				}
				if e.Alive {
					rl.DrawTexturePro(g.gameTexture, e.RecIn, e.Rec, e.Pos, 0, enemyColour)
				}
			}
		}

		// Draw enemy bullets
		if g.EnemyBullet.Shoot {
			rl.DrawRectangleRec(g.EnemyBullet.Rec, rl.Blue)
		}

		rl.DrawText(fmt.Sprint("Health: ", g.Player.Health), 20, ScreenHeight-40, 30, rl.White)
		rl.DrawText(fmt.Sprint("Score: ", g.playerScore), ScreenWidth-200, ScreenHeight-40, 30, rl.White)
	} else if !g.gameActive && g.playerWon {
		rl.ClearBackground(rl.Black)
		for i := range g.Enemies {
			for _, e := range g.Enemies[i] {
				e.Rec.X = EnemyDisplacement
			}
		}
		text := "YOU'RE WINNER ! OMG! \n\n\tSo proud of u."
		rl.DrawText(text, ScreenWidth/2-100, 200, 20, rl.Green)
		if rl.IsKeyPressed(rl.KeyEnter) {
			g.enemySpeed = 20
			g.runningTime = 0
			g.playerScore = 0
			g.gameActive = true
			g.bulletTimer = 5
			InitGame()
		}
	} else {
		text := "You are lose. Hit enter to start again rofl."
		rl.DrawText(text, ScreenWidth/2-450, 200, 40, rl.Red)
		rl.DrawText(fmt.Sprint("Scorus finalis: ", g.playerScore), ScreenWidth/2-450, 400, 40, rl.White)
		// rl.DrawText(fmt.Sprint("Dat means u killed "))
		if rl.IsKeyPressed(rl.KeyEnter) {
			g.enemySpeed = 20
			g.runningTime = 0
			g.playerScore = 0
			g.gameActive = true
			g.bulletTimer = 5
			InitGame()
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

func GenerateTexture(width, height int, scale float32) rl.Texture2D {
	bImage := rl.GenImagePerlinNoise(ScreenWidth, ScreenHeight, 0, 0, scale)
	bTexture := rl.LoadTextureFromImage(bImage)
	rl.UnloadImage(bImage)

	return bTexture
}
