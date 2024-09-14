package scene

import (
	"encoding/json"
	"log"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Level struct {
	TilesY     int     `json:"height"`
	TilesX     int     `json:"width"`
	TileHeight float32 `json:"tileheight"`
	TileWidth  float32 `json:"tilewidth"`
	Layers     []Layer `json:"layers"`
	Tiles      []Tile
}

type Layer struct {
	Id    int     `json:"id"`
	Image string  `json:"image"`
	X     float32 `json:"x"`
	Y     float32 `json:"y"`
	Tiles []int   `json:"data"`
}

type TileData struct {
	Id          int    `json:"id"`
	Image       string `json:"image"`
	ImageHeight int    `json:"imageheight"`
	ImageWidth  int    `json:"imagewidth"`
}

type TileMap struct {
	Tiles []TileData `json:"tiles"`
}

type Tile struct {
	Tex   rl.Texture2D
	RecIn rl.Rectangle
	Rec   rl.Rectangle
}

func GenerateLevel() (level *Level) {
	levelData, err := os.Open("./scene/Village.json")
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}
	defer levelData.Close()

	decoder := json.NewDecoder(levelData)
	err = decoder.Decode(&level)
	if err != nil {
		log.Fatalf("Failed to decode JSON: %v", err)
	}

	level.Tiles = GenerateTiles(*level)

	return level
}

func GenerateTiles(level Level) (levelTiles []Tile) {
	jsontiles, err := os.Open("./scene/Ground.json")
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}
	defer jsontiles.Close()

	var tiles TileMap

	decoder := json.NewDecoder(jsontiles)
	err = decoder.Decode(&tiles)
	if err != nil {
		log.Fatalf("Failed to decode JSON: %v", err)
	}

	tileData := level.Layers[len(level.Layers)-1].Tiles

	rows, cols := level.TilesY, level.TilesX

	tileArr2d := make([][]int, rows)
	for i := range tileArr2d {
		tileArr2d[i] = make([]int, cols)
	}

	for i := range rows {
		for j := range cols {
			index := i*cols + j
			if index < len(tileData) {
				tileArr2d[i][j] = tileData[index]
			}
		}
	}

	var tiledTextureGrid []rl.Texture2D
	for i := range tiles.Tiles {
		tiledTextureGrid = append(tiledTextureGrid, rl.LoadTexture(tiles.Tiles[i].Image))
	}

	for i := range tileArr2d {
		for j := range tileArr2d[i] {
			if tileArr2d[i][j] > 0 {
				tile := Tile{
					Tex: tiledTextureGrid[tileArr2d[i][j]-1],
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
				}
				levelTiles = append(levelTiles, tile)
			}
		}
	}

	return levelTiles
}

func GenerateBackgroundFromLevel(l *Level) (background []rl.Texture2D) {
	for i := range l.Layers {
		b := rl.LoadTexture(l.Layers[i].Image)
		background = append(background, b)
	}
	return background
}

func DrawLevel(btex rl.Texture2D) {
	rl.DrawTexture(btex, 0, 0, rl.White)
}
