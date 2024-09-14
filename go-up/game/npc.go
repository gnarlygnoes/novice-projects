package game

import (
	"goup/engine"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type NPC struct {
	ID           CId
	Rec          rl.Rectangle
	Colour       rl.Color
	Health       int
	isEnemy      bool
	OnSurface    bool
	VertVel      float32
	hasWeight    bool
	shooter      bool
	facing       float32
	AIBullets    map[CId]RangedWeap
	canShoot     bool
	timer        engine.Timer
	ReloadSpeed  float64
	BulletColour rl.Color
	// canPatrol bool
}

func NewNPC(xpos, ypos float32, isEnemy bool, npcType string) NPC { // take type as string and return
	bullets := map[CId]RangedWeap{}
	if npcType == "Red Rectangle" {
		return NPC{
			ID: NextId(),
			Rec: rl.Rectangle{
				X:      xpos,
				Y:      ypos,
				Width:  50,
				Height: 80,
			},
			Colour:    rl.Red,
			Health:    2,
			isEnemy:   isEnemy,
			VertVel:   0,
			hasWeight: true,
			shooter:   false,
			AIBullets: bullets,
		}
	}
	if npcType == "Green Square" {
		rs := float64(rl.GetRandomValue(10, 20) / 10)
		c := rl.Color{
			R: 50,
			G: 200,
			B: 50,
			A: 255,
		}

		return NPC{
			ID: NextId(),
			Rec: rl.Rectangle{
				X:      xpos,
				Y:      ypos,
				Width:  60,
				Height: 60,
			},
			Colour:       c,
			Health:       3,
			isEnemy:      isEnemy,
			VertVel:      0,
			hasWeight:    true,
			shooter:      true,
			AIBullets:    bullets,
			canShoot:     true,
			ReloadSpeed:  rs,
			BulletColour: c,
		}
	}
	rs := 2.7
	return NPC{
		ID: NextId(),
		Rec: rl.Rectangle{
			X:      xpos,
			Y:      ypos,
			Width:  50,
			Height: 80,
		},
		Colour:       rl.White,
		Health:       2,
		isEnemy:      isEnemy,
		VertVel:      0,
		hasWeight:    true,
		shooter:      true,
		AIBullets:    bullets,
		BulletColour: rl.White,
		ReloadSpeed:  rs,
		canShoot:     true,
	}
}

func (g *Game) UpdateNPC(dt float32) {
	for id, npc := range g.npcs {
		if g.npcs[id].hasWeight {
			if CheckCollisionY(&npc.Rec, g.LevelData.Tiles) {
				npc.OnSurface = true
			} else {
				npc.OnSurface = false
			}

			if g.npcs[id].OnSurface {
				npc.VertVel = 0
			} else {
				npc.VertVel += Gravity * dt
			}

			npc.Rec.Y += g.npcs[id].VertVel * dt
		}

		// Shooting logic
		if npc.shooter && npc.canShoot {
			if g.player.Rec.X < npc.Rec.X {
				npc.facing = -1
				if npc.Rec.X-g.player.Rec.X <= 800 && npc.Rec.X-g.player.Rec.X > 0 {
					npc.Shoot()
				}
			} else {
				npc.facing = 1
				if g.player.Rec.X-npc.Rec.X <= 800 && g.player.Rec.X-npc.Rec.X > 0 {
					npc.Shoot()
				}
			}
		}
		npc.BulletsUpdate(g, dt)

		if !npc.canShoot {
			npc.timer.LifeTime = rl.GetTime() - npc.timer.StartTime
			// npcReloadTime := float64(rl.GetRandomValue(10, 20) / 10)
			if npc.timer.LifeTime >= npc.ReloadSpeed {
				npc.canShoot = true
			}
		}
		g.npcs[id] = npc

		// if g.npcs[id].Health <= 0 {
		// 	delete(g.npcs, id)
		// }
	}
}

func (npc *NPC) Shoot() {
	b := RangedWeap{
		Id: NextId(),
		Rec: rl.Rectangle{
			X:      npc.Rec.X + npc.Rec.Width/2,
			Y:      npc.Rec.Y + npc.Rec.Height/2,
			Width:  10,
			Height: 5,
		},
		Colour: npc.BulletColour,
		Speed:  1000 * npc.facing,
	}
	npc.AIBullets[b.Id] = b

	npc.canShoot = false
	npc.timer.StartTime = rl.GetTime()
}

func (n *NPC) BulletsUpdate(g *Game, dt float32) {
	for id, b := range n.AIBullets {
		b.Rec.X += b.Speed * dt

		n.AIBullets[id] = b

		if b.Rec.X > b.Rec.X+ScreenWidth || b.Rec.X < b.Rec.X-ScreenWidth {
			delete(n.AIBullets, id)
		}

		for _, tile := range g.levelTiles {
			if rl.CheckCollisionRecs(b.Rec, tile.Rec) {
				delete(n.AIBullets, id)
			}
		}

		if rl.CheckCollisionRecs(b.Rec, g.player.Rec) {
			delete(n.AIBullets, id)
			g.player.currentHealth--
		}
	}
}
