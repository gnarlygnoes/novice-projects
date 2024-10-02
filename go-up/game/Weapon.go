package game

import (
	"goup/engine"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type RangedWeap struct {
	// Active bool
	Id          engine.CId
	Rec         rl.Rectangle
	Colour      rl.Color
	Speed       float32
	ReloadSpeed float32
	Damage      int
}

func (p *Player) Shoot() {
	b := RangedWeap{
		Id: engine.NextId(),
		Rec: rl.Rectangle{
			X:      p.Rec.X + p.Rec.Width/2,
			Y:      p.Rec.Y + 25,
			Width:  20,
			Height: 5,
		},
		Colour:      rl.Orange,
		Speed:       2000 * p.Facing,
		ReloadSpeed: 0.7,
		Damage:      1,
	}

	p.Bullets[b.Id] = b

	p.canShoot = false

	p.BulletTimer.StartTime = rl.GetTime()
	// return b
}

func (p *Player) BulletsUpdate(g *Game, dt float32) {
	for id, b := range p.Bullets {
		// if p.Bullets[id].Active {
		b.Rec.X += b.Speed * dt
		// p.Bullets[id].Rec.X += p.Bullets[id].Speed * dt
		// }
		if b.Rec.X > b.Rec.X+ScreenWidth || b.Rec.X < b.Rec.X-ScreenWidth {
			delete(p.Bullets, id)
		}
		// if !p.Bullets[id].Active {
		// 	p.Bullets[id].Rec = rl.Rectangle{}
		// 	p.Bullets[id].Speed = 0
		// }
		p.Bullets[id] = b
	}
	p.BulletTimer.LifeTime = rl.GetTime() - p.BulletTimer.StartTime
	if p.BulletTimer.LifeTime >= 0.7 {
		p.canShoot = true
	}
}

// func (p *Player) ShootTimer() {
// 	if
// }

func (g *Game) BulletCollision() {
	for id := range g.player.Bullets {
		for _, tile := range g.LevelData.Tiles {
			if rl.CheckCollisionRecs(g.player.Bullets[id].Rec, tile.Rec) {
				// g.player.Bullets[id].Active = false
				delete(g.player.Bullets, id)
			}
		}
		for j, npc := range g.npcs {
			if rl.CheckCollisionRecs(g.player.Bullets[id].Rec, npc.Rec) {
				// g.player.Bullets[id].Active = false
				// g.enemies[i] =
				delete(g.player.Bullets, id)
				// delete(g.npcs, j)
				// npc.Health -= g.player.Bullets[id].Damage
				npc.Health--
			}
			g.npcs[j] = npc
			if g.npcs[j].Health <= 0 {
				delete(g.npcs, j)
			}
			// fmt.Println(g.npcs[j].Health)
		}
	}
}

type MeleeWeapon struct {
}
