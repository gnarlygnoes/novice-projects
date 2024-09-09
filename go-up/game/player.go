package game

import (
	"goup/engine"

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
	Bullets       map[CId]RangedWeap
	Shooting      bool
	Facing        float32
	maxHealth     int
	currentHealth int
	canShoot      bool
	BulletTimer   engine.Timer
}

func NewPlayer(health int) *Player {
	// b := map[CId]
	return &Player{
		Rec: rl.Rectangle{
			Width:  50,
			Height: 100,
			X:      ScreenWidth / 2,
			Y:      0},
		Colour:        rl.Color{R: 150, G: 70, B: 50, A: 255},
		Speed:         700,
		JumpSpeed:     1500,
		VertVel:       0,
		Shooting:      false,
		Bullets:       map[CId]RangedWeap{},
		maxHealth:     health,
		currentHealth: health,
		Facing:        1,
		canShoot:      true,
	}
}

func (p *Player) Update(g *Game, dt float32) {
	if CheckCollisionY(&p.Rec, g.levelTiles) {
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

func CheckCollisionY(Rec *rl.Rectangle, t []Tile) (onPlatform bool) {
	playerHeight := Rec.Height
	playerBottom := Rec.Y + playerHeight
	playerTop := Rec.Y

	for _, plat := range t {
		if rl.CheckCollisionRecs(*Rec, plat.Rec) {
			if playerBottom >= plat.Rec.Y && !(playerBottom > plat.Rec.Y+30) {
				Rec.Y = plat.Rec.Y - playerHeight + 1
				onPlatform = true
			} else if playerTop <= plat.Rec.Y+plat.Rec.Height {
				Rec.Y = plat.Rec.Y + plat.Rec.Height
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

	for _, plat := range g.levelTiles {
		if rl.CheckCollisionRecs(g.player.Rec, plat.Rec) {
			if (playerLeft < plat.Rec.X+plat.Rec.Width) &&
				(playerBottom > plat.Rec.Y+10) && (playerRight > plat.Rec.X+plat.Rec.Width-10) {
				g.player.Rec.X = plat.Rec.X + plat.Rec.Width
			} else if (playerRight > plat.Rec.X) && (playerBottom > plat.Rec.Y+10) {
				g.player.Rec.X = plat.Rec.X - g.player.Rec.Width
			}
		}
	}

	for _, item := range g.items {
		if rl.CheckCollisionRecs(g.player.Rec, item.Rec) {
			if item.ItemType == "health +1" && g.player.currentHealth < g.player.maxHealth {
				g.player.currentHealth++
				delete(g.items, item.Id)
			}
		}
	}

	g.BulletCollision()
}
