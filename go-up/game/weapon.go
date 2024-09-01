package game

import rl "github.com/gen2brain/raylib-go/raylib"

type Bullet struct {
	Active bool
	Rec    rl.Rectangle
	Colour rl.Color
	Speed  float32
}

func PlayerProjectilesInit() (projectiles [50]Bullet) {
	for i := range projectiles {
		projectiles[i] = Bullet{
			Active: false,
			Rec: rl.Rectangle{
				X:      0,
				Y:      0,
				Width:  0,
				Height: 0,
			},
			Colour: rl.Orange,
			Speed:  0,
		}
	}

	return projectiles
}

func (p *Player) Shoot() {
	// if g.player.Shooting {
	for i := range p.Bullets {
		if !p.Bullets[i].Active {
			p.Bullets[i].Active = true
			p.Bullets[i].Rec.X = p.Rec.X + p.Rec.Width/2
			p.Bullets[i].Rec.Y = p.Rec.Y + 25
			p.Bullets[i].Rec.Width = 20
			p.Bullets[i].Rec.Height = 5
			if p.Facing == 0 {
				p.Facing = 1
			}
			p.Bullets[i].Speed = 2000 * p.Facing
			break
		}
	}
}

func (p *Player) BulletsUpdate(g *Game, dt float32) {
	for i := range p.Bullets {
		if p.Bullets[i].Active {
			p.Bullets[i].Rec.X += p.Bullets[i].Speed * dt
		}
		if p.Bullets[i].Rec.X > p.Rec.X+ScreenWidth || p.Bullets[i].Rec.X < p.Rec.X-ScreenWidth {
			p.Bullets[i].Active = false
		}
		if !p.Bullets[i].Active {
			p.Bullets[i].Rec = rl.Rectangle{}
			p.Bullets[i].Speed = 0
		}
	}
}

type MeleeWeapon struct {
}

// func EnemyProjectilesInit() (projectiles [50]Bullet) {
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
