package scene

import rl "github.com/gen2brain/raylib-go/raylib"

func LoadImage(img rl.Image) rl.Texture2D {
	tex := rl.LoadTextureFromImage(&img)

	return tex
}

// func makeSceneFromImage(img rl.Image) {
// 	tex := rl.LoadTextureFromImage(&img)

// }
