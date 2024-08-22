package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Enemy struct {
	RecIn       rl.Rectangle
	Rec         rl.Rectangle
	Pos         rl.Vector2
	Colour      rl.Color
	Alive       bool
	HitPoints   int
	Shooting    bool
	MovingRight bool
	HBox        rl.Vector2
}

type EnemyBullet struct {
	Rec   rl.Rectangle
	Shoot bool
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
