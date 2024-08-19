package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
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

	dt          float32
	runningTime float32

	enemySpeed  float32
	playerSpeed float32
	bulletSpeed float32

	gameTexture rl.Texture2D
	texSegmentH int32
	texSegmentV int32

	frame int32

	Player      Player
	Bullets     [NumBullets]Bullet
	Enemies     [EnemyGridY][EnemyGridX]Enemy
	EnemyBullet EnemyBullet
	Stars       [NumStars]Star
	Defence     [NumDefences]Defence
}

func NewGame() *Game {
	gameTexture := rl.LoadTexture("img/SpaceInvaders.png")
	texSegmentH := gameTexture.Width / 7
	texSegmentV := gameTexture.Height / 5

	player := NewPlayer(gameTexture, texSegmentH, texSegmentV)
	// Player
	g := &Game{
		gameActive:  true,
		playerScore: 0,
		enemySpeed:  20,
		bulletSpeed: 2000,
		runningTime: 0,
		bulletTimer: 5,

		Player: *player,
	}

	// Initialise stars
	for i := range g.Stars {
		g.Stars[i] = GenerateStars()
	}

	g.frame = 0

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

func (g *Game) Update() {
	g.dt = rl.GetFrameTime()

	g.runningTime += g.dt

	g.HandleInputs()
	if g.gameActive {
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
	// rl.DrawTexture(g.clouds, 0, 0, rl.Gray)
	// rl.DrawTexturePro(g.buildingTexture)

	/* Start Screen:
	You is like a Spartan or something and you is being attacked and must use the defences
	to protecc yourself rofl.

	Or: You is being attacked by aliens. You has like a space barricade or something and can
	use it to defend yourself from alien goo-bullets.
	*/

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
			NewGame()
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
			NewGame()
		}
	}

	rl.DrawFPS(20, 20)

	rl.EndDrawing()
}
