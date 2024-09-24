package game

import (
	"goup/engine"
	"goup/scene"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Rec           rl.Rectangle
	Colour        rl.Color
	Speed         float32
	JumpSpeed     float32
	VertVel       float32
	OnSurface     bool
	Direction     float32
	Moving        bool
	Jump          bool
	ResetPos      bool
	Crouched      bool
	Bullets       map[engine.CId]RangedWeap
	Shooting      bool
	Facing        float32
	maxHealth     int
	currentHealth int
	canShoot      bool
	BulletTimer   engine.Timer
}

func NewPlayer(health int, pos rl.Vector2) *Player {
	return &Player{
		Rec: rl.Rectangle{
			Width:  50,
			Height: 100,
			X:      pos.X,
			Y:      pos.Y},
		Colour:        rl.Color{R: 150, G: 70, B: 50, A: 255},
		Speed:         700,
		JumpSpeed:     2000,
		VertVel:       0,
		Shooting:      false,
		Bullets:       map[engine.CId]RangedWeap{},
		maxHealth:     health,
		currentHealth: health,
		Facing:        1,
		canShoot:      true,
	}
}

func (p *Player) Update(g *Game, dt float32) {
	if CheckCollisionY(&p.Rec, g.LevelData.Tiles) {
		p.OnSurface = true
	} else if p.Rec.Y+p.Rec.Height > g.LevelData.TileHeight*float32(g.LevelData.TilesY) {
		p.Rec.Y = g.LevelData.TileHeight*float32(g.LevelData.TilesY) - p.Rec.Height + 1
		p.OnSurface = true
	} else {
		p.OnSurface = false
	}

	if p.OnSurface {
		g.player.VertVel = 0
	} else {
		p.VertVel += Gravity * dt
	}

	PlayerInputs(p, dt)
	g.MoveAndCollideX(dt)

	if p.Jump {
		p.VertVel = -p.JumpSpeed
		p.Jump = false
	}

	if p.Crouched {
		p.Rec.Height = 50
	} else {
		p.Rec.Height = 100
	}

	if p.ResetPos {
		p.Rec.X = ScreenWidth / 2
		p.Rec.Y = 0
		p.ResetPos = false
	}

	p.BulletsUpdate(g, dt)

	p.Rec.Y += p.VertVel * dt
}

func CheckCollisionY(Rec *rl.Rectangle, t []scene.Tile) (onTile bool) {
	recHeight := Rec.Height
	recBottom := Rec.Y + recHeight
	recTop := Rec.Y
	recPosX := Rec.X + (Rec.Width / 2)

	// for _, tile := range t {
	// 	if rl.CheckCollisionRecs(*Rec, tile.Rec) {
	// 		if playerBottom >= tile.Rec.Y && !(playerBottom > tile.Rec.Y+30) {
	// 			Rec.Y = tile.Rec.Y - playerHeight + 1
	// 			onTile = true
	// 		} else if playerTop <= tile.Rec.Y+tile.Rec.Height {
	// 			Rec.Y = tile.Rec.Y + tile.Rec.Height
	// 			onTile = false
	// 		} else {
	// 			onTile = false
	// 		}
	// 	}
	// }

	for _, tile := range t {
		if rl.CheckCollisionRecs(*Rec, tile.Rec) {
			if recBottom >= tile.Rec.Y {
				for i := 1; i < len(tile.CollisionLines)-1; i++ {
					if recPosX > tile.CollisionLines[i-1].X && recPosX < tile.CollisionLines[i].X {
						gradient := (tile.CollisionLines[i].Y - tile.CollisionLines[i-1].Y) / (tile.CollisionLines[i].X - tile.CollisionLines[i-1].X)
						posYOnLine := tile.CollisionLines[i-1].Y + (recPosX-tile.CollisionLines[i-1].X)*gradient
						// fmt.Println("Gradient: ", gradient, "PosX: ", recPosX, "PosYOnLine: ", posYOnLine)
						if Rec.Y+Rec.Height >= posYOnLine {
							Rec.Y = posYOnLine - recHeight + 1
							onTile = true
						}
					}
				}
				// Rec.Y = tile.Rec.Y - playerHeight + 1
				// onTile = true
			} else if recTop <= tile.Rec.Y+tile.Rec.Height {
				Rec.Y = tile.Rec.Y + tile.Rec.Height
				onTile = false
			} else {
				onTile = false
			}
		}
	}

	return onTile
}

func (g *Game) MoveAndCollideX(dt float32) {
	g.player.Rec.X += g.player.Speed * g.player.Direction * dt

	playerWidth := g.player.Rec.Width
	playerLeft := g.player.Rec.X
	playerRight := playerLeft + playerWidth
	playerBottom := g.player.Rec.Y + g.player.Rec.Height

	// gradient :=

	// playerPosX := playerLeft + (playerRight / 2)

	// for _, plat := range g.LevelData.Tiles {
	// 	if rl.CheckCollisionRecs(g.player.Rec, plat.Rec) {
	// 		if (playerLeft < plat.Rec.X+plat.Rec.Width) &&
	// 			(playerBottom > plat.Rec.Y+10) && (playerRight > plat.Rec.X+plat.Rec.Width-10) {
	// 			g.player.Rec.X = plat.Rec.X + plat.Rec.Width
	// 		} else if (playerRight > plat.Rec.X) && (playerBottom > plat.Rec.Y+10) {
	// 			g.player.Rec.X = plat.Rec.X - playerWidth
	// 		}
	// 	}
	// 	// }
	// }

	// for _, plat := range g.LevelData.Tiles {
	// 	if rl.CheckCollisionRecs(g.player.Rec, plat.Rec) {
	// 		if (playerLeft < plat.Rec.X+plat.Rec.Width) &&
	// 			(playerBottom > plat.Rec.Y+10) && (playerRight > plat.Rec.X+plat.Rec.Width-10) {
	// 			g.player.Rec.X = plat.Rec.X + plat.Rec.Width
	// 		} else if (playerRight > plat.Rec.X) && (playerBottom > plat.Rec.Y+10) {
	// 			g.player.Rec.X = plat.Rec.X - playerWidth
	// 		}
	// 	}
	// }

	for _, item := range g.items {
		if rl.CheckCollisionRecs(g.player.Rec, item.Rec) {
			if item.ItemType == "health +1" && g.player.currentHealth < g.player.maxHealth {
				g.player.currentHealth++
				delete(g.items, item.Id)
			}
		}
	}

	for _, npc := range g.npcs {
		if rl.CheckCollisionRecs(g.player.Rec, npc.Rec) {
			if (playerLeft < npc.Rec.X+npc.Rec.Width) && (playerBottom > npc.Rec.Y+10) &&
				playerRight > npc.Rec.X+npc.Rec.Width-10 {
				g.player.Rec.X = npc.Rec.X + npc.Rec.Width
			} else if (playerRight > npc.Rec.X) && (playerBottom > npc.Rec.Y+10) {
				g.player.Rec.X = npc.Rec.X - playerWidth
			}
		}
	}

	g.BulletCollision()
}
