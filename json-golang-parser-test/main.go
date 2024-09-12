package main

import (
	"encoding/json"
	"fmt"
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
	TilesY     int               `json:"height"`
	TilesX     int               `json:"width"`
	TileHeight float32           `json:"tileheight"`
	TileWidth  float32           `json:"tilewidth"`
	Background []BackGroundLayer `json:"layers"`
}

type BackGroundLayer struct {
	Id    int     `json:"id"`
	Image string  `json:"image"`
	X     float32 `json:"x"`
	Y     float32 `json:"y"`
}

type Player struct {
	Rec    rl.Rectangle
	Colour rl.Color
}

func main() {
	levelData, err := os.Open("Village.json")
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}
	defer levelData.Close()

	var level LevelData
	// var levelBackground BackGroundLayer

	decoder := json.NewDecoder(levelData)
	err = decoder.Decode(&level)
	if err != nil {
		log.Fatalf("Failed to decode JSON: %v", err)
	}
	// level.Background =

	// fmt.Println("Width: ", level.TilesX)
	// fmt.Println("Height: ", level.TilesY)
	// fmt.Println("Tile Height: ", level.TileHeight)
	// fmt.Println("Tile Width: ", level.TileWidth)
	fmt.Println("Background: ", level.Background[0])

	rl.InitWindow(ScreenWidth, ScreenHeight, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(240)

	background := rl.LoadTexture(level.Background[0].Image)
	background1 := rl.LoadTexture(level.Background[1].Image)
	background2 := rl.LoadTexture(level.Background[2].Image)
	flora := rl.LoadTexture(level.Background[2].Image)
	// fmt.Println("HHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHH", background)

	// player := Player{
	// 	Rec: rl.Rectangle{
	// 		X:      0,
	// 		Y:      1000,
	// 		Width:  50,
	// 		Height: 80,
	// 	},
	// 	Colour: rl.Purple,
	// }
	player := MakePlayer()
	var playerSpeed float32 = 5

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		// rl.ClearBackground(rl.RayWhite)
		// rl.DrawTextureEx(background, rl.Vector2{0, 0}, 0, 0, rl.White)
		rl.DrawTexture(background, 0, 0, rl.White)
		rl.DrawTexture(background1, 0, 0, rl.White)
		rl.DrawTexture(background2, 0, 0, rl.White)
		rl.DrawTexture(flora, 0, 0, rl.White)

		// tile layer

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
