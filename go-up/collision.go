package main

// import (
// 	"fmt"

// 	rl "github.com/gen2brain/raylib-go/raylib"
// )

// func (g *Game) PlatformCollision() bool {
// 	for _, pf := range g.platformTiles {
// 		if rl.CheckCollisionRecs(g.player.rec, pf.rec) {
// 			fmt.Println("CRASH!")
// 			return true
// 		} else {
// 			return false
// 		}
// 	}
// 	// return
// }

// func (g *Game) HandleCollisions(dt float32) {
// 	if g.player.moving {
// 		g.player.rec.X += g.player.speed * g.player.direction * dt
// 		if g.PlatformCollision() {

// 		}
// 	}
// }
