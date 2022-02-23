package main

import (
	"fmt"
	"github.com/gen2brain/raylib-go/raylib"
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
	borderwidth  = int32(100)
	panSpeed     = .75
	screenWidth  = int32(1600)
	screenHeight = int32(900)
)

var (
	screenPointVector3 = rl.NewVector3(0, 0, 0)
	points3D           = hexCorner3D(screenPointVector3, 10)
	//screenPointVector2 = rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2))
	//points2D = hexPointCorner2D(screenPointVector2, 100)
	//corners2D = points2D[:]
	cursor = mouseCursor{
		floatrad: 3.0,
		intrad:   3.0,
		color:    rl.Black,
	}
	screen = window{
		N: 0,
		W: 0,
		S: float32(screenHeight),
		E: float32(screenWidth),
	}
	top = rl.Rectangle{X: screen.W + float32(borderwidth),
		Y:      screen.N,
		Width:  screen.E - (float32(borderwidth) * 2),
		Height: float32(borderwidth)}
	bottom = rl.Rectangle{
		X:      screen.W + float32(borderwidth),
		Y:      screen.S - float32(borderwidth),
		Width:  screen.E - (float32(borderwidth) * 2),
		Height: float32(borderwidth),
	}
	left = rl.Rectangle{
		X:      screen.W,
		Y:      screen.N + float32(borderwidth),
		Width:  float32(borderwidth),
		Height: screen.S - (float32(borderwidth) * 2),
	}
	right = rl.Rectangle{
		X:      screen.E - float32(borderwidth),
		Y:      screen.N + float32(borderwidth),
		Width:  float32(borderwidth),
		Height: screen.S - (float32(borderwidth) * 2),
	}
	topLeft = rl.Rectangle{
		X:      screen.W,
		Y:      screen.N,
		Width:  float32(borderwidth),
		Height: float32(borderwidth),
	}
	topRight = rl.Rectangle{
		X:      screen.E - float32(borderwidth),
		Y:      screen.N,
		Width:  float32(borderwidth),
		Height: float32(borderwidth),
	}
	bottomLeft = rl.Rectangle{
		X:      screen.W,
		Y:      screen.S - float32(borderwidth),
		Width:  float32(borderwidth),
		Height: float32(borderwidth),
	}
	bottomRight = rl.Rectangle{
		X:      screen.E - float32(borderwidth),
		Y:      screen.S - float32(borderwidth),
		Width:  float32(borderwidth),
		Height: float32(borderwidth),
	}
	camera = rl.Camera{
		Position: rl.NewVector3(0.0, -50.0, -25.0),
		Target:   rl.NewVector3(0.0, 0.0, 0.0),
		Up:       rl.NewVector3(0.0, 0.0, 1.0),
		Fovy:     75}
)

func main() {
	rl.InitWindow(screenWidth, screenHeight, "Window")
	rl.SetTargetFPS(60)

	hexMesh := rl.GenMeshPoly(6, 10)
	hexModel := rl.LoadModelFromMesh(hexMesh)

	for !rl.WindowShouldClose() {
		cameraControl(&camera)

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.BeginMode3D(camera)
		//drawHex3D(points3D)
		rl.DrawGrid(10, 10)
		rl.DrawModelWires(hexModel, rl.NewVector3(2, -2, -5), 2, rl.Green)
		rl.EndMode3D()
		debugText(&camera)
		rl.DrawPoly(rl.NewVector2(float32(rl.GetMouseX())-3, float32(rl.GetMouseY())), 2, cursor.floatrad, 135, cursor.color)
		rl.EndDrawing()
		rl.DisableCursor()

		if rl.GetMouseX()+cursor.intrad > int32(rl.GetScreenWidth()) {
			rl.SetMousePosition(rl.GetScreenWidth()-int(cursor.intrad), int(rl.GetMouseY()))
		} else if rl.GetMouseX()-cursor.intrad <= 0 {
			rl.SetMousePosition(0+int(cursor.intrad), int(rl.GetMouseY()))
		}
		if rl.GetMouseY()+cursor.intrad > int32(rl.GetScreenHeight()) {
			rl.SetMousePosition(int(rl.GetMouseX()), rl.GetScreenHeight()-int(cursor.intrad))
		} else if rl.GetMouseY()-cursor.intrad < 0 {
			rl.SetMousePosition(int(rl.GetMouseX()), 0+int(cursor.intrad))
		}

	}
	rl.CloseWindow()
}

func debugText(camera *rl.Camera) {
	rl.DrawText("Pos X: "+fmt.Sprintf("%.2f", camera.Position.X)+", Y: "+fmt.Sprintf("%.2f", camera.Position.Y)+", Z: "+fmt.Sprintf("%.2f", camera.Position.Z), 10, 10, 20, rl.Gray)
	rl.DrawText("Target X: "+fmt.Sprintf("%.2f", camera.Target.X)+", Y: "+fmt.Sprintf("%.2f", camera.Target.Y)+", Z: "+fmt.Sprintf("%.2f", camera.Target.Z), 10, 30, 20, rl.Gray)
	rl.DrawText("Up X: "+fmt.Sprintf("%.2f", camera.Up.X)+", Y: "+fmt.Sprintf("%.2f", camera.Up.Y)+", Z: "+fmt.Sprintf("%.2f", camera.Up.Z), 10, 50, 20, rl.Gray)
	rl.DrawText("FPS: "+fmt.Sprintf("%.2f", rl.GetFPS()), int32(screen.E)-5-rl.MeasureText("FPS: "+fmt.Sprintf("%.2f", rl.GetFPS()), 20), 10, 20, rl.Gray)
}
