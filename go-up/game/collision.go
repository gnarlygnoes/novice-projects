package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func CheckCollisionY(p *Player, t []Tile) (onPlatform bool) {
	playerHeight := p.Rec.Height
	playerBottom := p.Rec.Y + playerHeight
	playerTop := p.Rec.Y

	for _, plat := range t {
		if rl.CheckCollisionRecs(p.Rec, plat.Rec) {
			if playerBottom >= plat.Rec.Y && !(playerBottom > plat.Rec.Y+30) {
				p.Rec.Y = plat.Rec.Y - playerHeight + 1
				onPlatform = true
			} else if playerTop <= plat.Rec.Y+plat.Rec.Height {
				p.Rec.Y = plat.Rec.Y + plat.Rec.Height
				onPlatform = false
			} else {
				onPlatform = false
			}
		}
	}

	return onPlatform
}

func (g *Game) MoveAndCollideX(dt float32) {
	g.player.Rec.X += g.player.Speed * g.player.Direction * dt

	playerWidth := g.player.Rec.Width
	playerLeft := g.player.Rec.X
	playerRight := playerLeft + playerWidth
	playerBottom := g.player.Rec.Y + g.player.Rec.Height

	for _, plat := range g.platformTiles {
		if rl.CheckCollisionRecs(g.player.Rec, plat.Rec) {
			if (playerLeft < plat.Rec.X+plat.Rec.Width) &&
				(playerBottom > plat.Rec.Y+10) && (playerRight > plat.Rec.X+plat.Rec.Width-10) {
				g.player.Rec.X = plat.Rec.X + plat.Rec.Width
			} else if (playerRight > plat.Rec.X) && (playerBottom > plat.Rec.Y+10) {
				g.player.Rec.X = plat.Rec.X - g.player.Rec.Width
			}
		}
	}

	g.BulletCollision()
}

func (g *Game) BulletCollision() {
	for i := range g.player.Bullets {
		for j := range g.platformTiles {
			if rl.CheckCollisionRecs(g.player.Bullets[i].Rec, g.platformTiles[j].Rec) {
				g.player.Bullets[i].Active = false
			}
		}
		for j := range g.groundTiles {
			if rl.CheckCollisionRecs(g.player.Bullets[i].Rec, g.groundTiles[j].Rec) {
				g.player.Bullets[i].Active = false
			}
		}
	}
}

func (g *Game) NPCCollision() {

}
