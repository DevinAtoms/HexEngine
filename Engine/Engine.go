package engine

import (
	"fmt"
	"github.com/DevinAtoms/HexEngine/HexMath"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type mouseCursor struct {
	floatrad float32
	intrad   int32
	color    rl.Color
}

var cursor = mouseCursor{
	floatrad: 3.0,
	intrad:   3.0,
	color:    rl.Black,
}

func DrawScreen() {
	rl.ClearBackground(rl.RayWhite)
	rl.BeginDrawing()
	Render3D()
	Render2D()
	rl.EndDrawing()
}

func Render3D() {
	tile := HexMath.HexCoord{}
	tile.Q = 1
	tile.R = -1
	tile.S = 0
	c, _ := HexMath.GetHexCoord(tile)
	rl.BeginMode3D(Camera)

	//drawOriginHex()
	Wireframe(rl.Vector3Zero(), HexMath.Apothem)
	Wireframe(c, HexMath.Apothem)
	rl.DrawSphere(HexMath.OriginHex.Points[0], .05, rl.Red)
	debugShapes()

	rl.EndMode3D()
}

func Render2D() {
	debugText(&Camera)
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
	//Z Axis
	rl.DrawLine3D(rl.NewVector3(0, 0, -10), rl.NewVector3(0, 0, 10), rl.Green)
	//Y Axis
	rl.DrawLine3D(rl.NewVector3(0, -10, 0), rl.NewVector3(0, 10, 0), rl.Red)
	//X Axis
	rl.DrawLine3D(rl.NewVector3(-10, 0, 0), rl.NewVector3(10, 0, 0), rl.Blue)
}

func Wireframe(center rl.Vector3, size float32) {
	corners := HexMath.HexCorner3D(center, size)
	for i := range corners {
		if i < 5 {
			rl.DrawLine3D(corners[i], corners[i+1], rl.Black)
			rl.DrawLine3D(
				rl.NewVector3(corners[i].X, corners[i].Y+.2, corners[i].Z),
				rl.NewVector3(corners[i+1].X, corners[i+1].Y+.2, corners[i+1].Z), rl.Black)
			rl.DrawLine3D(
				rl.NewVector3(corners[i].X, corners[i].Y+.1, corners[i].Z),
				rl.NewVector3(corners[i+1].X, corners[i+1].Y+.1, corners[i+1].Z), rl.Black)
			rl.DrawLine3D(
				corners[i],
				rl.NewVector3(corners[i].X, corners[i].Y+.2, corners[i].Z), rl.Black)
			rl.DrawLine3D(
				corners[i],
				rl.NewVector3(corners[i].X, corners[i].Y+.1, corners[i].Z), rl.Black)
			rl.DrawLine3D(
				corners[i+1],
				rl.NewVector3(corners[i+1].X, corners[i+1].Y+.2, corners[i+1].Z), rl.Black)
			rl.DrawLine3D(
				corners[i+1],
				rl.NewVector3(corners[i+1].X, corners[i+1].Y+.1, corners[i+1].Z), rl.Black)

		} else if i == 5 {
			rl.DrawLine3D(corners[5], corners[0], rl.Black)
			rl.DrawLine3D(
				rl.NewVector3(corners[5].X, corners[5].Y+.2, corners[5].Z),
				rl.NewVector3(corners[0].X, corners[0].Y+.2, corners[0].Z), rl.Black)
			rl.DrawLine3D(
				corners[5],
				rl.NewVector3(corners[5].X, corners[5].Y+.2, corners[5].Z), rl.Black)
		}
	}
}
