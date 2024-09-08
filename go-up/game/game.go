package game

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ScreenWidth  = 1920
	ScreenHeight = 1080
	Gravity      = 9800
)

type Game struct {
	// Background rl.Texture2D

	Camera *Camera
	player Player
	// enemy []Enemy

	levelTiles []Tile
	// platformTiles []Tile

	// enemies []NPC
	enemies map[CId]NPC
}

func NewGame() *Game {
	// img := rl.LoadImage("./img/GrassyField.png")
	// backgroundTex := rl.LoadTextureFromImage(img)
	// rl.UnloadImage(img)
	// tex := rl.LoadTexture("./img/Mossy Tileset/Mossy - Tileset.png")
	t, e := GenerateTileMap()
	// make(NPC, 0)
	g := &Game{
		// Background: backgroundTex,
		// (rl.Image{"./img/GrassyField.png"}),

		player: *NewPlayer(3),
		Camera: NewCamera(ScreenWidth, ScreenHeight),

		levelTiles: t,
		enemies:    e,
	}

	return g
}

func (g *Game) SetGameMode() {}

func (g *Game) Update() {
	// fmt.Println(g.player.Bullets)
	dt := rl.GetFrameTime()

	// for i := range g.player.Bullets {
	// fmt.Println(len(g.player.Bullets))
	// }

	g.player.Update(g, dt)
	g.Camera.Update(&g.player)
	g.UpdateNPC(dt)
}

func (g *Game) Draw() {

	rl.BeginDrawing()

	rl.ClearBackground(rl.Blue)
	// rl.DrawTexture(g.Background, 0, 0, rl.White)
	rl.BeginMode2D(g.Camera.Camera2D)

	// userInterface.DrawInterface(g)

	for i := range g.levelTiles {
		rl.DrawRectangleRec(g.levelTiles[i].Rec, g.levelTiles[i].Colour)
	}

	for i := range g.levelTiles {
		rl.DrawRectangleRec(g.levelTiles[i].Rec, g.levelTiles[i].Colour)
	}

	for i := range g.enemies {
		rl.DrawRectangleRec(g.enemies[i].Rec, g.enemies[i].Colour)
	}

	for _, b := range g.player.Bullets {
		rl.DrawRectangleRec(b.Rec, b.Colour)
	}

	rl.DrawRectangleRec(g.player.Rec, g.player.Colour)
	rl.EndMode2D()

	// Draw Onscreen UI
	rl.DrawText(fmt.Sprint("Health: ", g.player.currentHealth), 50, ScreenHeight-50, 36, rl.White)

	rl.EndDrawing()
}
