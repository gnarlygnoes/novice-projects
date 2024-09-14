package scene

import (
	"encoding/json"
	"log"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Level struct {
	TilesY     int               `json:"height"`
	TilesX     int               `json:"width"`
	TileHeight float32           `json:"tileheight"`
	TileWidth  float32           `json:"tilewidth"`
	Layers     []BackGroundLayer `json:"layers"`
	// Tilemap    []TileLayer       `json:"layers"`
}

type BackGroundLayer struct {
	Id    int     `json:"id"`
	Image string  `json:"image"`
	X     float32 `json:"x"`
	Y     float32 `json:"y"`
}

type TileLayer struct {
	Id       int   `json:"id"`
	TilesY   int   `json:"height"`
	TilesX   int   `json:"width"`
	TileData []int `json:"data"`
}

func GenerateLevel() (l *Level) {
	levelData, err := os.Open("./scene/Village.json")
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}
	defer levelData.Close()

	// var l Level

	decoder := json.NewDecoder(levelData)
	err = decoder.Decode(&l)
	if err != nil {
		log.Fatalf("Failed to decode JSON: %v", err)
	}
	// fmt.Println("AAAAAAAAAAAAAAABackground: ", len(l.Layers))
	// for i := range l.Layers {
	// 	if l.Layers[i].Id == 1 {

	// 	}
	// }
	// var bTex []rl.Texture2D
	// var t rl.Texture2D
	// for i := range l.Layers {
	// 	if l.Layers[i].Id > 1 {
	// 		t := rl.LoadTexture(l.Layers[i].Image)
	// 		bTex = append(bTex, t)
	// 	}
	// }
	// bTex = rl.LoadTexture(l.Layers[0].Image)

	return l
}

func GenerateBackground(l *Level) rl.Texture2D {
	background := rl.LoadTexture(l.Layers[0].Image)
	return background
}

func DrawLevel(btex rl.Texture2D) {
	rl.DrawTexture(btex, 0, 0, rl.White)
}
