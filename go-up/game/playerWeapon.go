package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type RangedWeap struct {
	// Active bool
	Id          CId
	Rec         rl.Rectangle
	Colour      rl.Color
	Speed       float32
	ReloadSpeed float32
	Damage      int
}

// func PlayerProjectilesInit() (projectiles [50]Bullet) {
// 	for i := range projectiles {
// 		projectiles[i] = Bullet{
// 			Active: false,
// 			Rec: rl.Rectangle{
// 				X:      0,
// 				Y:      0,
// 				Width:  0,
// 				Height: 0,
// 			},
// 			Colour: rl.Orange,
// 			Speed:  0,
// 		}
// 	}

// 	return projectiles
// }

func (p *Player) Shoot() {
	// if g.player.Shooting {
	// for i := range p.Bullets {
	// 	if !p.Bullets[i].Active {
	// 		p.Bullets[i].Active = true
	// 		p.Bullets[i].Rec.X = p.Rec.X + p.Rec.Width/2
	// 		p.Bullets[i].Rec.Y = p.Rec.Y + 25
	// 		p.Bullets[i].Rec.Width = 20
	// 		p.Bullets[i].Rec.Height = 5
	// 		if p.Facing == 0 {
	// 			p.Facing = 1
	// 		}
	// 		p.Bullets[i].Speed = 2000 * p.Facing
	// 		break
	// 	}
	// }
	b := RangedWeap{
		Id: NextId(),
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

	// for i := range p.Bullets {
	// 	if i == len(p.Bullets)-1 {
	// 		p.Bullets[i] = b
	// 	}
	// }
	// for id := range p.Bullets {
	// 	p.Bullets[id] = b
	// }
	// p.Bullets = map[Bullet{}
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
		for j := range g.levelTiles {
			if rl.CheckCollisionRecs(g.player.Bullets[id].Rec, g.levelTiles[j].Rec) {
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
