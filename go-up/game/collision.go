package game

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func PlatformCollisionY(g *Game) (onPlatform bool) {
	playerHeight := g.player.rec.Height
	playerBottom := g.player.rec.Y + playerHeight
	playerTop := g.player.rec.Y

	for _, plat := range g.platformTiles {
		if rl.CheckCollisionRecs(g.player.rec, plat.rec) {
			if playerBottom >= plat.rec.Y && !(playerBottom > plat.rec.Y+30) {
				g.player.rec.Y = plat.rec.Y - playerHeight + 1
				onPlatform = true
			} else if playerTop <= plat.rec.Y+plat.rec.Height {
				g.player.rec.Y = plat.rec.Y + plat.rec.Height
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

	if playerLeft <= 0 {
		g.player.rec.X = 0
	}
	if playerRight >= ScreenWidth {
		g.player.rec.X = ScreenWidth - playerWidth
	}

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
	fmt.Println(g.player.vVel)
}
