package main

import rl "github.com/gen2brain/raylib-go/raylib"

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
