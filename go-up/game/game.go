package game

import (
	"fmt"
	"goup/engine"
	"goup/scene"
	"goup/scene/locations"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ScreenWidth  = 1920
	ScreenHeight = 1080
	Gravity      = 9800
)

type Game struct {
	Background []rl.Texture2D
	paused     bool
	GameMode   int
	LevelData  scene.Level

	Camera *Camera
	player Player

	LevelNum int

	levelTiles []scene.Tile

	npcs       map[engine.CId]NPC
	items      map[engine.CId]Item
	startpoint rl.Vector2
	endpoint   float32

	displayCollisions bool
}

func NewGame() *Game {

	// img := rl.LoadImage("./img/GrassyField.png")
	// backgroundTex := rl.LoadTextureFromImage(img)
	// rl.UnloadImage(img)
	// tex := rl.LoadTexture("./img/Mossy Tileset/Mossy - Tileset.png")
	// l := 1
	// t, e, i, sp, ep := GenerateLevel(l)

	startLevel := locations.FirstLevel()
	// npcData := startLevel.NpcData
	npcs := MapNpcs(startLevel.NpcData)

	level := scene.GenerateLevel(startLevel.LevelName, startLevel.TileSet)
	background := scene.GenerateBackgroundFromLevel(level)
	// var bTex []rl.Texture2D
	// for i := range level.Layers {
	// 	if level.Layers[i].Id > 1 {
	// 		bTex[i] = rl.LoadTexture(level.Layers[i].Image)
	// 	}
	// }
	// fmt.Println("BBBBBBBBBBBBBBBBBBBBBBBB: ", level.Layers)
	// make(NPC, 0)

	// g :=
	// levellyLevel := locations.ReturnLevel()
	// jsonfile := startLevel.LevelName

	return &Game{
		// Background: backgroundTex,
		// (rl.Image{"./img/GrassyField.png"}),
		LevelData:  *level,
		Background: background,

		player:   *NewPlayer(3, rl.Vector2{X: ScreenWidth / 2, Y: 0}),
		Camera:   NewCamera(ScreenWidth, ScreenHeight),
		GameMode: 1,

		// levelTiles: t,
		npcs: npcs,
		// items:      i,
		// startpoint: sp,
		// endpoint:   ep,
		// LevelNum:   l,
	}
}

func (g *Game) SetGameMode() {
	switch g.LevelNum {
	case 1:
		g.LevelNum = 2
		// g.levelTiles, g.npcs, g.items, g.startpoint, g.endpoint = GenerateLevel(g.LevelNum)
		g.player.Rec.X = g.startpoint.X
		g.player.Rec.Y = g.startpoint.Y - g.player.Rec.Height - 500
	case 2:
		g.GameMode = 2
	}
}

func (g *Game) Update() {
	// fmt.Println(g.player.Bullets)
	dt := rl.GetFrameTime()

	GameInputs(g)

	if !g.paused {

		if g.player.Rec.X >= g.endpoint {
			g.SetGameMode()
		}

		// for i := range g.player.Bullets {
		// fmt.Println(len(g.player.Bullets))
		// }

		g.player.Update(g, dt)
		g.Camera.Update(&g.player, g)
		g.UpdateNPC(dt)
	} else {
		gamePaused := "GAME PAUSED"
		rl.DrawText(gamePaused, ScreenWidth/2-int32(len(gamePaused)*24)/2, ScreenHeight/2, 36, rl.Red)
	}
}

func (g *Game) Draw() {
	rl.BeginDrawing()

	if g.GameMode == 1 {
		// rl.ClearBackground(rl.Blue)
		for i := range g.Background {
			scene.DrawLevel(g.Background[i])
		}
		// rl.DrawTexture(l, 0, 0, rl.White)
		// for i := range g.Background {
		// 	// if l.Id > 1 {
		// 	// bTex := rl.LoadTexture(l.Image)
		// 	rl.DrawTexture(g.Background[i], 0, 0, rl.White)
		// 	// }
		// }
		// fmt.Println("AAAAAAAAAAAAAAAAAAAAAAAA: ", g.LevelData)

		// rl.DrawTexture(g.Background, 0, 0, rl.White)

		rl.BeginMode2D(g.Camera.Camera2D)

		// userInterface.DrawInterface(g)

		// for i := range g.levelTiles {
		// 	rl.DrawRectangleRec(g.levelTiles[i].Rec, g.levelTiles[i].Colour)
		// }

		// for i := range g.levelTiles {
		// 	rl.DrawRectangleRec(g.levelTiles[i].Rec, g.levelTiles[i].Colour)
		// }
		// rl.DrawPoly()

		for i := range g.LevelData.Tiles {
			rl.DrawTexturePro(g.LevelData.Tiles[i].Tex, g.LevelData.Tiles[i].RecIn,
				g.LevelData.Tiles[i].Rec, rl.Vector2{X: 0, Y: 0}, 0, rl.White)

			if g.displayCollisions {
				for j := 1; j < len(g.LevelData.Tiles[i].CollisionLines)-1; j++ {
					// rl.DrawLine(int32(g.LevelData.Tiles[i].Rec.X)+int32(g.LevelData.Tiles[i].CollisionLines[j-1].X),
					// 	int32(g.LevelData.Tiles[i].Rec.Y)+int32(g.LevelData.Tiles[i].CollisionLines[j-1].Y),
					// 	int32(g.LevelData.Tiles[i].Rec.X)+int32(g.LevelData.Tiles[i].CollisionLines[j].X),
					// 	int32(g.LevelData.Tiles[i].Rec.Y)+int32(g.LevelData.Tiles[i].CollisionLines[j].Y),
					// 	rl.Red)
					rl.DrawLine(int32(g.LevelData.Tiles[i].CollisionLines[j-1].X), int32(g.LevelData.Tiles[i].CollisionLines[j-1].Y),
						int32(g.LevelData.Tiles[i].CollisionLines[j].X), int32(g.LevelData.Tiles[i].CollisionLines[j].Y),
						rl.Red)
				}
			}
		}

		for i := range g.npcs {
			rl.DrawRectangleRec(g.npcs[i].Rec, g.npcs[i].Colour)
		}

		for _, b := range g.player.Bullets {
			rl.DrawRectangleRec(b.Rec, b.Colour)
		}

		for i := range g.npcs {
			for _, b := range g.npcs[i].AIBullets {
				rl.DrawRectangleRec(b.Rec, b.Colour)
			}
		}

		for _, i := range g.items {
			rl.DrawRectangleRec(i.Rec, i.Colour)
		}

		rl.DrawRectangleRec(g.player.Rec, g.player.Colour)

		rl.EndMode2D()

		// Draw Onscreen UI
		rl.DrawText(fmt.Sprint("Health: ", g.player.currentHealth), 50, ScreenHeight-50, 36, rl.White)
	}
	if g.GameMode == 2 {
		// rl.ClearBackground(rl.Blue)
		// rl.DrawTexture(g.Background, 0, 0, rl.White)
		// rl.BeginMode2D(g.Camera.Camera2D)

		rl.DrawText("Wow, you are really good at this game. So proud of u.\n\nPress Enter to Continue",
			ScreenWidth/2-200, ScreenHeight/2, 36, rl.Red)

		rl.EndMode2D()
	}
	rl.DrawFPS(20, 30)
	rl.EndDrawing()
}
