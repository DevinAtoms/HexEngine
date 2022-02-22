package main

import (
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
	screenPointVector3, screenPointVector2 = rl.NewVector3(0, 0, 0), rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2))
	points3D                               = hexCorner3D(screenPointVector3, 10)
	points2D, corners2D                    = hexPointCorner2D(screenPointVector2, 100), points2D[:]
	cursor                                 = mouseCursor{
		floatrad: 4.0,
		intrad:   4.0,
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
		X:      screen.N,
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
)

func Init() rl.Camera {
	var camera = rl.Camera{}
	rl.InitWindow(screenWidth, screenHeight, "Window")
	rl.SetTargetFPS(60)

	camera = rl.Camera{
		Position: rl.NewVector3(0.0, 0.0, 0.0),
		Target:   rl.NewVector3(0.0, 0.0, 0.0),
		Up:       rl.NewVector3(0.0, 0.0, 0.0),
		Fovy:     75}

	return camera
}

func main() {
	camera := Init()
	rl.SetCameraMode(camera, rl.CameraCustom)
	rl.SetCameraPanControl(5)
	for !rl.WindowShouldClose() {

		rl.UpdateCamera(&camera)
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.BeginMode3D(camera)
		rl.DrawSphere(rl.NewVector3(10, 10, 0), 1, rl.Black)
		drawHex3D(points3D)

		//rl.DrawModel(rl.LoadModelFromMesh(rl.GenMeshPoly(6, 10)), screenPointVector3, .25, rl.Green)

		//rl.DrawText("Camera X: "+fmt.Sprintf("%.2f", camera.Position.X), 10, 10, 20, rl.Gray)
		//rl.DrawText("Camera Y: "+fmt.Sprintf("%.2f", camera.Position.Y), 10, 30, 20, rl.Gray)
		//rl.DrawText("Camera Z: "+fmt.Sprintf("%.2f", camera.Position.Z), 10, 50, 20, rl.Gray)
		//rl.DrawText("Mouse X: "+fmt.Sprintf("%d", rl.GetMouseX()), 10, 70, 20, rl.Gray)
		//rl.DrawText("Mouse Y: "+fmt.Sprintf("%d", rl.GetMouseY()), 10, 90, 20, rl.Gray)
		//rl.DrawText("FPS: "+fmt.Sprintf("%.2f", rl.GetFPS()), int32(screen.E)-5-rl.MeasureText("FPS: "+fmt.Sprintf("%.2f", rl.GetFPS()), 20), 10, 20, rl.Gray)
		//rl.DrawLineStrip(corners2D, int32(len(corners2D)), rl.NewColor(0, 0, 0, 0))
		//	rl.DrawCircleV(rl.GetMousePosition(), cursor.floatrad, cursor.color)
		//	rl.DisableCursor()
		rl.EndMode3D()
		rl.EndDrawing()

		//	if rl.GetMouseX()+cursor.intrad > int32(rl.GetScreenWidth()) {
		//		rl.SetMousePosition(rl.GetScreenWidth()-int(cursor.intrad), int(rl.GetMouseY()))
		//	} else if rl.GetMouseX()-cursor.intrad <= 0 {
		//		rl.SetMousePosition(0+int(cursor.intrad), int(rl.GetMouseY()))
		//	}
		//
		//	if rl.GetMouseY()+cursor.intrad > int32(rl.GetScreenHeight()) {
		//		rl.SetMousePosition(int(rl.GetMouseX()), rl.GetScreenHeight()-int(cursor.intrad))
		//	} else if rl.GetMouseY()-cursor.intrad < 0 {
		//		rl.SetMousePosition(int(rl.GetMouseX()), 0+int(cursor.intrad))
		//	}
		//
		//	if rl.CheckCollisionPointRec(rl.GetMousePosition(), right) {
		//		camera.Position.X += panSpeed
		//		camera.Target.X += panSpeed
		//	} else if rl.CheckCollisionPointRec(rl.GetMousePosition(), left) {
		//		camera.Position.X -= panSpeed
		//		camera.Target.X -= panSpeed
		//	} else if rl.CheckCollisionPointRec(rl.GetMousePosition(), bottom) {
		//		camera.Target.Y -= panSpeed
		//		camera.Position.Y -= panSpeed
		//	} else if rl.CheckCollisionPointRec(rl.GetMousePosition(), top) {
		//		camera.Target.Y += panSpeed
		//	} else if rl.CheckCollisionPointRec(rl.GetMousePosition(), bottomRight) {
		//		camera.Position.X += panSpeed
		//		camera.Target.X += panSpeed
		//		camera.Target.Y -= panSpeed
		//	} else if rl.CheckCollisionPointRec(rl.GetMousePosition(), topLeft) {
		//		camera.Position.X -= panSpeed
		//		camera.Target.X -= panSpeed
		//		camera.Target.Y += panSpeed
		//	} else if rl.CheckCollisionPointRec(rl.GetMousePosition(), topRight) {
		//		camera.Position.X += panSpeed
		//		camera.Target.X += panSpeed
		//		camera.Target.Y += panSpeed
		//	} else if rl.CheckCollisionPointRec(rl.GetMousePosition(), bottomLeft) {
		//		camera.Position.X -= panSpeed
		//		camera.Target.X -= panSpeed
		//		camera.Target.Y -= panSpeed
		//	}
	}

	rl.CloseWindow()
}
