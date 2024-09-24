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
	// Npcs       map[game.CId]game.NPC
}

type Layer struct {
	Id    int     `json:"id"`
	Image string  `json:"image"`
	X     float32 `json:"x"`
	Y     float32 `json:"y"`
	Tiles []int   `json:"data"`
}

type TileData struct {
	Id          int                `json:"id"`
	Image       string             `json:"image"`
	ImageWidth  int                `json:"imagewidth"`
	ImageHeight int                `json:"imageheight"`
	ObjectGroup CollisionObjectEnd `json:"objectgroup"`
}

type CollisionObjectEnd struct {
	Id      int                    `json:"id"`
	X       float32                `json:"x"`
	Y       float32                `json:"y"`
	Width   float32                `json:"width"`
	Height  float32                `json:"height"`
	Objects []CollisionObjectStart `json:"objects"`
}

type CollisionObjectStart struct {
	Id      int             `json:"id"`
	Width   float32         `json:"width"`
	Height  float32         `json:"height"`
	X       float32         `json:"x"`
	Y       float32         `json:"y"`
	Polygon []CollisionPoly `json:"polygon"`
}

type CollisionPoly struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

type TileMap struct {
	Tiles []TileData `json:"tiles"`
}

type Tile struct {
	Tex            rl.Texture2D
	RecIn          rl.Rectangle
	Rec            rl.Rectangle
	CollisionLines []rl.Vector2
}

func GenerateLevel(levelname, tileset string) (level *Level) {
	levelData, err := os.Open(levelname)
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}
	defer levelData.Close()

	decoder := json.NewDecoder(levelData)
	err = decoder.Decode(&level)
	if err != nil {
		log.Fatalf("Failed to decode JSON: %v", err)
	}

	level.Tiles = GenerateTiles(*level, "./scene/GroundRevamped.json")

	return level
}

// func GenerateEntities(level Level)

func GenerateTiles(level Level, tilesetName string) (levelTiles []Tile) {
	jsontiles, err := os.Open(tilesetName)
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
	w, h := level.TileWidth, level.TileHeight

	for i := range tileArr2d {
		for j := range tileArr2d[i] {
			if tileArr2d[i][j] > 0 {
				tile := Tile{
					Tex: tiledTextureGrid[tileArr2d[i][j]-1],
					RecIn: rl.Rectangle{
						X:      0,
						Y:      0,
						Width:  w,
						Height: h,
					},
					Rec: rl.Rectangle{
						X:      float32(j) * float32(w),
						Y:      float32(i) * float32(h),
						Width:  w,
						Height: h,
					},
				}
				levelTiles = append(levelTiles, tile)
			}
		}
	}

	for i := range levelTiles {
		levelTiles[i].CollisionLines = GenerateTileCollision(tiles.Tiles[levelTiles[i].Tex.ID-3])
		for j := range levelTiles[i].CollisionLines {
			levelTiles[i].CollisionLines[j].X += levelTiles[i].Rec.X
			levelTiles[i].CollisionLines[j].Y += levelTiles[i].Rec.Y
		}
	}

	return levelTiles
}

func GenerateTileCollision(tile TileData) []rl.Vector2 {
	var collisionPoly []rl.Vector2
	for _, co := range tile.ObjectGroup.Objects {
		if co.Width > 0 {
			collisionPoly = append(collisionPoly, rl.Vector2{X: 0, Y: 0})
			collisionPoly = append(collisionPoly, rl.Vector2{X: 128, Y: 0})
			collisionPoly = append(collisionPoly, rl.Vector2{X: 128, Y: 128})
			collisionPoly = append(collisionPoly, rl.Vector2{X: 0, Y: 128})
		} else {
			collisionPoly = append(collisionPoly, rl.Vector2{X: tile.ObjectGroup.X + co.X, Y: tile.ObjectGroup.Y + co.Y})
			for _, p := range co.Polygon {
				collisionPoly = append(collisionPoly, rl.Vector2{X: p.X + co.X, Y: p.Y + co.Y})
			}
		}
		collisionPoly = append(collisionPoly, rl.Vector2{X: tile.ObjectGroup.X + co.X, Y: tile.ObjectGroup.Y + co.Y})
	}
	return collisionPoly
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
