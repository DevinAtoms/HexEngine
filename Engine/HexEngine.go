package Engine

import (
	"github.com/DevinAtoms/HexEngine/HexMath"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type (
	mouseCursor struct {
		floatrad float32
		intrad   int32
		color    rl.Color
	}
	window struct {
		N float32
		S float32
		E float32
		W float32
	}
)

const (
	//borderwidth = int32(100)
	screenWidth  = int32(1600)
	screenHeight = int32(900)
)

var cursor = mouseCursor{
	floatrad: 3.0,
	intrad:   3.0,
	color:    rl.Black,
}

func main() {
	rl.InitWindow(screenWidth, screenHeight, "Window")

	rl.SetTargetFPS(60)

	rotateCamera(&Camera, false)

	loadAssets()

	for !rl.WindowShouldClose() {
		cameraControl(&Camera)

		drawScreen()

		rl.DisableCursor()
	}
	closeApp()
}

func closeApp() {
	rl.UnloadModel(HexMath.OriginHex.Model)
	rl.CloseWindow()
}

func loadAssets() {
	HexMath.OriginHex = *LoadOriginHex()
}

func LoadOriginHex() *HexMath.HexTile {

	tile := HexMath.HexTile{}
	tile.Model = rl.LoadModel("assets/grass.obj")
	tile.Points = HexMath.HexCorner3D(rl.Vector3Zero(), HexMath.Apothem)
	tile.Coord = rl.NewVector2(0, 0)

	return &tile
}
