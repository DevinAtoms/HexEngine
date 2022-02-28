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
	panSpeed     = .1
	screenWidth  = int32(1600)
	screenHeight = int32(900)
)

var (
	originHex *hexTile
	cursor    = mouseCursor{
		floatrad: 3.0,
		intrad:   3.0,
		color:    rl.Black,
	}
	camera = rl.Camera{
		// -Z Forward / +Z Backwards
		// -X Left / +X Right
		// -Y Down / +Y Up
		Position: rl.NewVector3(0.0, 2, -10),
		Target:   rl.NewVector3(0.0, 0.0, 0.0),
		Up:       rl.NewVector3(0.0, 1.0, 0.0),
		Fovy:     75}
)

func main() {
	rl.SetConfigFlags(rl.FlagFullscreenMode)
	rl.InitWindow(int32(rl.GetMonitorWidth(1)), int32(rl.GetMonitorHeight(1)), "Window")
	rl.SetTargetFPS(60)
	rotateCamera(&camera, false)

	originHex = loadHex()

	for !rl.WindowShouldClose() {
		cameraControl(&camera)
		rl.ClearBackground(rl.RayWhite)
		rl.BeginDrawing()
		render3D()
		render2D()
		rl.EndDrawing()
		rl.DisableCursor()
	}
	closeApp()
}

func closeApp() {
	rl.UnloadModel(originHex.model)
	rl.CloseWindow()
}

func render3D() {
	rl.BeginMode3D(camera)
	//drawOriginHex()
	wireframe(rl.Vector3Zero(), apothem)
	wireframe(getHexCoord(newCoords(1, -1, 0)), apothem)
	//drawHex(hexCoords(rl.NewVector3(0, 0, 0)))

	debugShapes()

	rl.EndMode3D()
}

func render2D() {
	debugText(&camera)
	rl.DrawPoly(rl.NewVector2(float32(rl.GetMouseX())-3, float32(rl.GetMouseY())), 2, cursor.floatrad, 135, cursor.color)
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

func debugText(camera *rl.Camera) {
	rl.DrawText("Pos X: "+fmt.Sprintf("%.2f", camera.Position.X)+", Y: "+fmt.Sprintf("%.2f", camera.Position.Y)+", Z: "+fmt.Sprintf("%.2f", camera.Position.Z), 10, 10, 20, rl.Gray)
	rl.DrawText("Target X: "+fmt.Sprintf("%.2f", camera.Target.X)+", Y: "+fmt.Sprintf("%.2f", camera.Target.Y)+", Z: "+fmt.Sprintf("%.2f", camera.Target.Z), 10, 30, 20, rl.Gray)
}

func debugShapes() {
	rl.DrawGrid(100, 1)
	//Center Marker
	//rl.DrawCubeWires(rl.NewVector3(0, 0, 0), 2, 2, 1, rl.Black)
	//Camera Target Marker
	//rl.DrawSphere(camera.Target, .1, rl.Black)
	//Z Axis
	rl.DrawLine3D(rl.NewVector3(0, 0, -10), rl.NewVector3(0, 0, 10), rl.Green)
	//Y Axis
	rl.DrawLine3D(rl.NewVector3(0, -10, 0), rl.NewVector3(0, 10, 0), rl.Red)
	//XAxis
	rl.DrawLine3D(rl.NewVector3(-10, 0, 0), rl.NewVector3(10, 0, 0), rl.Blue)
}
