package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func CheckCollisionY(p *Player, t []Tile) (onPlatform bool) {
	playerHeight := p.rec.Height
	playerBottom := p.rec.Y + playerHeight
	playerTop := p.rec.Y

	for _, plat := range t {
		if rl.CheckCollisionRecs(p.rec, plat.rec) {
			if playerBottom >= plat.rec.Y && !(playerBottom > plat.rec.Y+30) {
				p.rec.Y = plat.rec.Y - playerHeight + 1
				onPlatform = true
			} else if playerTop <= plat.rec.Y+plat.rec.Height {
				p.rec.Y = plat.rec.Y + plat.rec.Height
				onPlatform = false
			} else {
				onPlatform = false
			}
		}
	}

	return onPlatform
}

func (g *Game) MoveAndCollideX(dt float32) {
	g.player.rec.X += g.player.speed * g.player.direction * dt

	playerWidth := g.player.rec.Width
	playerLeft := g.player.rec.X
	playerRight := playerLeft + playerWidth
	playerBottom := g.player.rec.Y + g.player.rec.Height

	for _, plat := range g.platformTiles {
		if rl.CheckCollisionRecs(g.player.rec, plat.rec) {

			if (playerLeft < plat.rec.X+plat.rec.Width) &&
				(playerBottom > plat.rec.Y+10) && (playerRight > plat.rec.X+plat.rec.Width-10) {
				g.player.rec.X = plat.rec.X + plat.rec.Width
			} else if (playerRight > plat.rec.X) && (playerBottom > plat.rec.Y+10) {
				g.player.rec.X = plat.rec.X - g.player.rec.Width
			}
		}
	}
}
