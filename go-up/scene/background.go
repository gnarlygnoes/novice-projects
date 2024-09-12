package scene

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

// func GenerateBackgroundImages() []rl.Texture2D {

// }
