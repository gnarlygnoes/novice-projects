package main

import (
	"encoding/json"
	"log"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"os"
// )

// type User struct {
// 	Name  string `json:"name"`
// 	Age   int    `json:"age"`
// 	Email string `json:"email"`
// }

// type Users struct {
// 	Users []User `json:"users"`
// }

// func main() {
// 	userData, err := os.Open("user.json")
// 	if err != nil {
// 		log.Fatalf("Error parsing JSON: %v", err)
// 	}
// 	defer userData.Close()

// 	// Create an instance of Person to store parsed data
// 	// var user User
// 	var users Users

// 	// Parse the JSON into the struct
// 	// err := json.Unmarshal([]byte(jsonData), &person)

// 	decoder := json.NewDecoder(userData)
// 	err = decoder.Decode(&users)
// 	// json.Unmarshal([]byte{}, &users)
// 	if err != nil {
// 		log.Fatalf("Failed to decode JSON: %v", err)
// 	}

// 	// Output the result
// 	for i := range users.Users {
// 		fmt.Printf("Name: %s, Age: %d, Email: %s\n", users.Users[i].Name, users.Users[i].Age, users.Users[i].Email)
// 	}
// }

const (
	ScreenWidth  = 1920
	ScreenHeight = 1080
)

type LevelData struct {
	TilesY     int     `json:"height"`
	TilesX     int     `json:"width"`
	TileHeight float32 `json:"tileheight"`
	TileWidth  float32 `json:"tilewidth"`
	Layer      []Layer `json:"layers"`
	// Tilemap    []TileLayer       `json:"layers"`
}

type Layer struct {
	Id      int     `json:"id"`
	Image   string  `json:"image"`
	X       float32 `json:"x"`
	Y       float32 `json:"y"`
	TileMap []int   `json:"data"`
}

type Tile struct {
	Rec    rl.Rectangle
	Colour rl.Color
}

type Tiles struct {
	Tiles []TileTex `json:"tiles"`
}

type TileTex struct {
	Id          int    `json:"id"`
	Image       string `json:"image"`
	ImageHeight int    `json:"imageheight"`
	ImageWidth  int    `json:"imagewidth"`
}

type TexturedTile struct {
	Tex   rl.Texture2D
	RecIn rl.Rectangle
	Rec   rl.Rectangle
}

type Player struct {
	Rec    rl.Rectangle
	Colour rl.Color
}

type Camera struct {
	rl.Camera2D
	Width  float32
	Height float32
}

func main() {
	levelData, err := os.Open("Village.json")
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}
	defer levelData.Close()

	var level LevelData

	decoder := json.NewDecoder(levelData)
	err = decoder.Decode(&level)
	if err != nil {
		log.Fatalf("Failed to decode JSON: %v", err)
	}

	tileTexData, fail := os.Open("Ground.json")
	if fail != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}
	defer levelData.Close()

	var tiles Tiles

	decoderTile := json.NewDecoder(tileTexData)
	fail = decoderTile.Decode(&tiles)
	if fail != nil {
		log.Fatalf("Failed to decode JSON: %v", fail)
	}

	rl.InitWindow(ScreenWidth, ScreenHeight, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(420)

	background := rl.LoadTexture(level.Layer[0].Image)
	background1 := rl.LoadTexture(level.Layer[1].Image)
	background2 := rl.LoadTexture(level.Layer[2].Image)
	flora := rl.LoadTexture(level.Layer[2].Image)

	var tileRec []TexturedTile

	tileData := level.Layer[len(level.Layer)-1].TileMap

	rows, cols := level.TilesY, level.TilesX

	tileData2d := make([][]int, rows)
	for i := range tileData2d {
		tileData2d[i] = make([]int, cols)
	}

	for i := range rows {
		for j := range cols {
			index := i*cols + j
			if index < len(tileData) {
				tileData2d[i][j] = tileData[index]
			}
		}
	}

	var tileTextureSlice []rl.Texture2D
	for i := range tiles.Tiles {
		tileTextureSlice = append(tileTextureSlice, rl.LoadTexture(tiles.Tiles[i].Image))
	}

	for i := range tileData2d {
		for j := range tileData2d[i] {
			if tileData2d[i][j] > 0 {
				tile := TexturedTile{
					Tex: tileTextureSlice[tileData2d[i][j]-1],
					RecIn: rl.Rectangle{
						X:      0,
						Y:      0,
						Width:  128,
						Height: 128,
					},
					Rec: rl.Rectangle{
						X:      float32(j) * float32(54),
						Y:      float32(i) * float32(54),
						Width:  54,
						Height: 54,
					},
					// Colour: rl.Red,
				}
				tileRec = append(tileRec, tile)
			}
		}
	}

	player := MakePlayer()
	var playerSpeed float32 = 5

	camera := MakeCamera(ScreenWidth, ScreenHeight)
	// c := rl.Camera2D{}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		UpdateCamera(&player, &camera)

		// rl.ClearBackground(rl.RayWhite)
		// rl.DrawTextureEx(background, rl.Vector2{0, 0}, 0, 0, rl.White)
		rl.DrawTexture(background, 0, 0, rl.White)
		rl.DrawTexture(background1, 0, 0, rl.White)
		rl.DrawTexture(background2, 0, 0, rl.White)
		rl.DrawTexture(flora, 0, 0, rl.White)

		// fmt.Println(tileRec[1])

		for i := range tileRec {
			// rl.DrawRectangleRec(tileRec[i].Rec, rl.White)
			rl.DrawTexturePro(tileRec[i].Tex, tileRec[i].RecIn, tileRec[i].Rec, rl.Vector2{X: 0, Y: 0}, 0, rl.White)
		}

		if rl.IsKeyDown(rl.KeyA) {
			player.Rec.X -= playerSpeed
		}
		if rl.IsKeyDown(rl.KeyD) {
			player.Rec.X += playerSpeed
		}
		if rl.IsKeyDown(rl.KeyW) {
			player.Rec.Y -= playerSpeed
		}
		if rl.IsKeyDown(rl.KeyS) {
			player.Rec.Y += playerSpeed
		}

		if player.Rec.X < 0 {
			player.Rec.X = 0
		}
		if player.Rec.X+player.Rec.Width > ScreenWidth {
			player.Rec.X = ScreenWidth - player.Rec.Width
		}
		if player.Rec.Y < 0 {
			player.Rec.Y = 0
		}
		if player.Rec.Y+player.Rec.Height > ScreenHeight {
			player.Rec.Y = ScreenHeight - player.Rec.Height
		}

		rl.DrawRectangleRec(player.Rec, player.Colour)

		rl.DrawFPS(20, 30)
		rl.EndDrawing()
	}
}

func MakePlayer() Player {
	return Player{
		Rec: rl.Rectangle{
			X:      0,
			Y:      1000,
			Width:  50,
			Height: 80,
		},
		Colour: rl.Purple}
}

func MakeCamera(w, h int) Camera {
	return Camera{
		Camera2D: rl.Camera2D{
			Zoom: 1,
		},
		Width:  float32(w),
		Height: float32(h),
	}
}

func UpdateCamera(player *Player, c *Camera) {
	c.Target.X = player.Rec.X + (player.Rec.Width / 2) - (ScreenWidth / 2)
}
