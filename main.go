package main

import (
	"fmt"
	"github.com/gen2brain/raylib-go/raylib"
	hex "github.com/hautenessa/hexagolang"
)

type mouseCursor struct {
	floatrad float32
	intrad   int32
	color    rl.Color
}

type window struct {
	N float32
	S float32
	E float32
	W float32
}

const (
	borderwidth  = int32(100)
	panSpeed     = .75
	screenWidth  = int32(1600)
	screenHeight = int32(900)
)

func main() {
	//rl.SetConfigFlags(rl.FlagFullscreenMode)
	rl.InitWindow(screenWidth, screenHeight, "Window")

	camera := rl.Camera{
		Position:   rl.NewVector3(0.0, 100.0, 1.0),
		Target:     rl.NewVector3(0.0, 0.0, 0.0),
		Up:         rl.NewVector3(0.0, 1.0, 0.0),
		Fovy:       75,
		Projection: rl.CameraOrthographic,
	}

	screenOrigin := hex.F{X: 0, Y: 0}
	hexRad := hex.F{X: 1, Y: 1}
	layout := hex.MakeLayout(hexRad, screenOrigin, hex.OrientationPointy)
	myFirstHex := hex.H{}                       // Uses axial coordinates.
	screenPoint := layout.CenterFor(myFirstHex) // convert the hexagon center into screen coordinates.

	cursor := mouseCursor{
		floatrad: 4.0,
		intrad:   4.0,
		color:    rl.Black,
	}
	screen := window{
		N: 0,
		W: 0,
		S: float32(screenHeight),
		E: float32(screenWidth),
	}

	top := rl.Rectangle{X: screen.W + float32(borderwidth), Y: screen.N, Width: screen.E - (float32(borderwidth) * 2), Height: float32(borderwidth)}
	bottom := rl.Rectangle{X: screen.W + float32(borderwidth), Y: screen.S - float32(borderwidth), Width: screen.E - (float32(borderwidth) * 2), Height: float32(borderwidth)}
	left := rl.Rectangle{X: screen.W, Y: screen.N + float32(borderwidth), Width: float32(borderwidth), Height: screen.S - (float32(borderwidth) * 2)}
	right := rl.Rectangle{X: screen.E - float32(borderwidth), Y: screen.N + float32(borderwidth), Width: float32(borderwidth), Height: screen.S - (float32(borderwidth) * 2)}
	topLeft := rl.Rectangle{X: screen.N, Y: screen.N, Width: float32(borderwidth), Height: float32(borderwidth)}
	topRight := rl.Rectangle{X: screen.E - float32(borderwidth), Y: screen.N, Width: float32(borderwidth), Height: float32(borderwidth)}
	bottomLeft := rl.Rectangle{X: screen.W, Y: screen.S - float32(borderwidth), Width: float32(borderwidth), Height: float32(borderwidth)}
	bottomRight := rl.Rectangle{X: screen.E - float32(borderwidth), Y: screen.S - float32(borderwidth), Width: float32(borderwidth), Height: float32(borderwidth)}

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		rl.UpdateCamera(&camera)
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		// Debug Borders
		//rl.DrawRectangleRec(top, rl.Yellow)
		//rl.DrawRectangleRec(left, rl.Green)
		//rl.DrawRectangleRec(right, rl.Blue)
		//rl.DrawRectangleRec(bottom, rl.Orange)
		//rl.DrawRectangleRec(topLeft, rl.Red)
		//rl.DrawRectangleRec(topRight, rl.Red)
		//rl.DrawRectangleRec(bottomLeft, rl.Red)
		//rl.DrawRectangleRec(bottomRight, rl.Red)

		rl.BeginMode3D(camera)
		rl.DrawGrid(50, 1.0)
		rl.DrawModel(rl.LoadModelFromMesh(rl.GenMeshPoly(6, 10)), rl.NewVector3(float32(screenPoint.X), float32(screenPoint.Y), 0), .25, rl.Green)
		rl.EndMode3D()

		rl.DrawText("Camera X: "+fmt.Sprintf("%.2f", camera.Position.X), 10, 10, 20, rl.Gray)
		rl.DrawText("Camera Y: "+fmt.Sprintf("%.2f", camera.Position.Y), 10, 30, 20, rl.Gray)
		rl.DrawText("Camera Z: "+fmt.Sprintf("%.2f", camera.Position.Z), 10, 50, 20, rl.Gray)
		rl.DrawText("Mouse X: "+fmt.Sprintf("%d", rl.GetMouseX()), 10, 70, 20, rl.Gray)
		rl.DrawText("Mouse Y: "+fmt.Sprintf("%d", rl.GetMouseY()), 10, 90, 20, rl.Gray)
		rl.DrawText("FPS: "+fmt.Sprintf("%.2f", rl.GetFPS()), int32(screen.E)-5-rl.MeasureText("FPS: "+fmt.Sprintf("%.2f", rl.GetFPS()), 20), 10, 20, rl.Gray)
		rl.DrawCircleV(rl.GetMousePosition(), cursor.floatrad, cursor.color)
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

		if rl.CheckCollisionPointRec(rl.GetMousePosition(), right) {
			camera.Position.X += panSpeed
			camera.Target.X += panSpeed
		} else if rl.CheckCollisionPointRec(rl.GetMousePosition(), left) {
			camera.Position.X -= panSpeed
			camera.Target.X -= panSpeed
		} else if rl.CheckCollisionPointRec(rl.GetMousePosition(), bottom) {
			camera.Position.Z += panSpeed
			camera.Target.Z += panSpeed
		} else if rl.CheckCollisionPointRec(rl.GetMousePosition(), top) {
			camera.Position.Z -= panSpeed
			camera.Target.Z -= panSpeed
		} else if rl.CheckCollisionPointRec(rl.GetMousePosition(), bottomRight) {
			camera.Position.X += panSpeed
			camera.Target.X += panSpeed
			camera.Position.Z += panSpeed
			camera.Target.Z += panSpeed
		} else if rl.CheckCollisionPointRec(rl.GetMousePosition(), topLeft) {
			camera.Position.X -= panSpeed
			camera.Target.X -= panSpeed
			camera.Position.Z -= panSpeed
			camera.Target.Z -= panSpeed
		} else if rl.CheckCollisionPointRec(rl.GetMousePosition(), topRight) {
			camera.Position.X += panSpeed
			camera.Target.X += panSpeed
			camera.Position.Z -= panSpeed
			camera.Target.Z -= panSpeed
		} else if rl.CheckCollisionPointRec(rl.GetMousePosition(), bottomLeft) {
			camera.Position.X -= panSpeed
			camera.Target.X -= panSpeed
			camera.Position.Z += panSpeed
			camera.Target.Z += panSpeed
		}
	}

	rl.CloseWindow()
}
