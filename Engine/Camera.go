package engine

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	panSpeed = .1
)

var (
	rightclick    bool
	scrollwheel   float64
	verticalAngle float64
	rotateAngle   float64
	rotated       float32
	v1            rl.Vector3
	v2            rl.Vector3
	matrix        rl.Matrix
	panning       bool
	rotating      bool
	distance      = 5.0
	Camera        = rl.Camera{
		// -Z Forward / +Z Backwards
		// -X Left / +X Right
		// -Y Down / +Y Up
		Position: rl.NewVector3(0.0, 2, -10),
		Target:   rl.NewVector3(0.0, 0.0, 0.0),
		Up:       rl.NewVector3(0.0, 1.0, 0.0),
		Fovy:     75}
)

func RotateCamera(camera *rl.Camera3D, lock bool) {

	rightclick = rl.IsMouseButtonDown(1)
	scrollwheel = float64(rl.GetMouseWheelMove())

	verticalAngle = float64(rl.Clamp(float32(verticalAngle), -1.5, 1.5))

	camera.Position.X = -float32(math.Sin(rotateAngle)*distance*math.Cos(verticalAngle) - float64(camera.Target.X))
	camera.Position.Y = -float32(math.Sin(verticalAngle)*distance - float64(camera.Target.Y))
	camera.Position.Z = -float32(math.Cos(rotateAngle)*distance*math.Cos(verticalAngle) - float64(camera.Target.Z))

	distance -= .1 * scrollwheel

	if lock {
		rotateAngle = float64(rl.Clamp(float32(rotateAngle), -2.5, 2.5))
		if !rotating {
			if rotateAngle > 0.1 {
				rotateAngle -= .075
			} else if rotateAngle < -0.1 {
				rotateAngle += .075
			} else {
				rotateAngle = 0
			}
		}
	}

}

func CameraControl(camera *rl.Camera) {
	RotateCamera(camera, false)

	look, _ := LocalMatrix(camera)
	v1 = rl.Vector3Transform(camera.Position, look)
	v2 = rl.Vector3Transform(camera.Target, look)

	matrix = rl.MatrixInvert(look)

	if !rotating {
		if rl.IsKeyDown(rl.KeyW) {
			panning = true
			v2.Z -= panSpeed
			v1.Z -= panSpeed
			camera.Position = rl.Vector3Transform(v1, matrix)
			camera.Target = rl.Vector3Transform(v2, matrix)
		} else {
			panning = false
		}
		if rl.IsKeyDown(rl.KeyS) {
			panning = true
			v2.Z += panSpeed
			v1.Z += panSpeed
			camera.Position = rl.Vector3Transform(v1, matrix)
			camera.Target = rl.Vector3Transform(v2, matrix)
		} else {
			panning = false
		}
		if rl.IsKeyDown(rl.KeyD) {
			panning = true
			v2.X += panSpeed
			v1.X += panSpeed
			camera.Position = rl.Vector3Transform(v1, matrix)
			camera.Target = rl.Vector3Transform(v2, matrix)
		} else {
			panning = false
		}
		if rl.IsKeyDown(rl.KeyA) {
			panning = true
			v2.X -= panSpeed
			v1.X -= panSpeed
			camera.Position = rl.Vector3Transform(v1, matrix)
			camera.Target = rl.Vector3Transform(v2, matrix)
		} else {
			panning = false
		}
		if rl.IsKeyDown(rl.KeySpace) {
			panning = true
			v2.Y += panSpeed
			v1.Y += panSpeed
			camera.Position = rl.Vector3Transform(v1, matrix)
			camera.Target = rl.Vector3Transform(v2, matrix)
		}
		if rl.IsKeyDown(rl.KeyLeftControl) {
			panning = true
			v2.Y -= panSpeed
			v1.Y -= panSpeed
			camera.Position = rl.Vector3Transform(v1, matrix)
			camera.Target = rl.Vector3Transform(v2, matrix)
		} else {
			panning = false
		}

	}

	if !panning {
		if rl.IsKeyDown(rl.KeyLeft) {
			rotating = true
			rotateAngle -= .1
			rotated -= .1
		}
		if rl.IsKeyDown(rl.KeyRight) {
			rotating = true
			rotateAngle += .1
		}
		if rightclick {
			rotating = true
			rotateAngle += float64(rl.GetMouseDelta().X * -.01)
			verticalAngle += float64(rl.GetMouseDelta().Y * -.01)
		}
		if !rl.IsKeyDown(rl.KeyLeft) && !rl.IsKeyDown(rl.KeyRight) && !rightclick {
			rotating = false
		}
	}

}

func LocalMatrix(d *rl.Camera3D) (rl.Matrix, rl.Matrix) {

	target := rl.NewVector3(d.Target.X, d.Position.Y, d.Target.Z)

	zoomZ := rl.Vector3Normalize(rl.Vector3Subtract(d.Target, d.Position))
	zoomX := rl.Vector3Normalize(rl.Vector3CrossProduct(zoomZ, d.Up))
	zoomY := rl.Vector3CrossProduct(zoomX, zoomZ)
	zoomZ = rl.Vector3Negate(zoomZ)

	zaxis := rl.Vector3Normalize(rl.Vector3Subtract(target, d.Position))
	xaxis := rl.Vector3Normalize(rl.Vector3CrossProduct(zaxis, d.Up))
	yaxis := rl.Vector3CrossProduct(xaxis, zaxis)
	zaxis = rl.Vector3Negate(zaxis)

	zoomxdot := rl.Vector3DotProduct(zoomX, Camera.Position)
	zoomydot := rl.Vector3DotProduct(zoomY, Camera.Position)
	zoomzdot := rl.Vector3DotProduct(zoomZ, Camera.Position)

	xdot := rl.Vector3DotProduct(xaxis, Camera.Position)
	ydot := rl.Vector3DotProduct(yaxis, Camera.Position)
	zdot := rl.Vector3DotProduct(zaxis, Camera.Position)

	view := rl.Matrix{
		M0:  xaxis.X,
		M4:  xaxis.Y,
		M8:  xaxis.Z,
		M12: -xdot,
		M1:  yaxis.X,
		M5:  yaxis.Y,
		M9:  yaxis.Z,
		M13: -ydot,
		M2:  zaxis.X,
		M6:  zaxis.Y,
		M10: zaxis.Z,
		M14: -zdot,
		M3:  0,
		M7:  0,
		M11: 0,
		M15: 1,
	}
	zoom := rl.Matrix{
		M0:  zoomX.X,
		M4:  zoomX.Y,
		M8:  zoomX.Z,
		M12: -zoomxdot,
		M1:  zoomY.X,
		M5:  zoomY.Y,
		M9:  zoomY.Z,
		M13: -zoomydot,
		M2:  zoomZ.X,
		M6:  zoomZ.Y,
		M10: zoomZ.Z,
		M14: -zoomzdot,
		M3:  0,
		M7:  0,
		M11: 0,
		M15: 1,
	}
	return view, zoom
}
