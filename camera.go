package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var verticalAngle float64
var horizontalAngle float64
var distance = float64(rl.Vector3Distance(camera.Position, camera.Target))
var horizontalDistance = distance * math.Cos(verticalAngle*rl.Deg2rad)

func cameraControl(camera *rl.Camera) {
	if !panNE() && !panNW() && !panSE() && !panSW() {

		camera.Position.Y -= float32(rl.GetMouseWheelMove() * 2)
		camera.Target.Y -= float32(rl.GetMouseWheelMove() * 2)

		if rl.IsKeyDown(rl.KeySpace) {
			rl.Vector3Transform(camera.Position, rl.MatrixRotate(camera.Position, 45))
			rl.DrawSphere(rl.Vector3Zero(), 1, rl.Black)
		}

		if rl.IsKeyDown(rl.KeyW) || panN() {
			camera.Target.Z += panSpeed
			camera.Position.Z += panSpeed
		}
		if rl.IsKeyDown(rl.KeyS) || panS() {
			camera.Target.Z -= panSpeed
			camera.Position.Z -= panSpeed
		}
		if rl.IsKeyDown(rl.KeyD) || panE() {
			camera.Target.X -= panSpeed
			camera.Position.X -= panSpeed
		}
		if rl.IsKeyDown(rl.KeyA) || panW() {
			camera.Target.X += panSpeed
			camera.Position.X += panSpeed
		}
	}
	if rl.IsKeyDown(rl.KeyUp) || panN() && verticalAngle < 88.0 {
		camera.Position.X = float32(horizontalDistance * math.Cos(horizontalAngle*rl.Deg2rad))
		camera.Position.Z = float32(horizontalDistance * math.Sin(horizontalAngle*rl.Deg2rad))
		camera.Position.Y = float32(distance * math.Sin(verticalAngle*rl.Deg2rad))
		verticalAngle += 1.0
	}
	if rl.IsKeyDown(rl.KeyDown) || panS() && verticalAngle > 2.0 {
		camera.Position.X = float32(horizontalDistance * math.Cos(horizontalAngle*rl.Deg2rad))
		camera.Position.Z = float32(horizontalDistance * math.Sin(horizontalAngle*rl.Deg2rad))
		camera.Position.Y = float32(distance * math.Sin(verticalAngle*rl.Deg2rad))
		verticalAngle -= 1.0
	}
	if rl.IsKeyDown(rl.KeyRight) || panE() {

		camera.Position.X = float32(horizontalDistance * math.Cos(horizontalAngle*rl.Deg2rad))
		camera.Position.Z = float32(horizontalDistance * math.Sin(horizontalAngle*rl.Deg2rad))
		horizontalAngle += 1.0
	}
	if rl.IsKeyDown(rl.KeyLeft) || panW() {
		camera.Position.X = float32(horizontalDistance * math.Cos(horizontalAngle*rl.Deg2rad))
		camera.Position.Z = float32(horizontalDistance * math.Sin(horizontalAngle*rl.Deg2rad))
		horizontalAngle -= 1.0
	}
	if panNE() {
		camera.Target.Z += panSpeed
		camera.Position.Z += panSpeed
		camera.Target.X -= panSpeed
		camera.Position.X -= panSpeed
	}
	if panSE() {
		camera.Target.Z -= panSpeed
		camera.Position.Z -= panSpeed
		camera.Target.X -= panSpeed
		camera.Position.X -= panSpeed
	}
	if panNW() {
		camera.Target.Z += panSpeed
		camera.Position.Z += panSpeed
		camera.Target.X += panSpeed
		camera.Position.X += panSpeed
	}
	if panSW() {
		camera.Target.Z -= panSpeed
		camera.Position.Z -= panSpeed
		camera.Target.X += panSpeed
		camera.Position.X += panSpeed
	}
}

func panN() bool {
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), top) {
		return true
	} else {
		return false
	}
}

func panS() bool {
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), bottom) {
		return true
	} else {
		return false
	}
}

func panE() bool {
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), right) {
		return true
	} else {
		return false
	}
}

func panW() bool {
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), left) {
		return true
	} else {
		return false
	}
}

func panNW() bool {
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), topLeft) {
		return true
	} else {
		return false
	}
}

func panSW() bool {
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), bottomLeft) {
		return true
	} else {
		return false
	}
}

func panNE() bool {
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), topRight) {
		return true
	} else {
		return false
	}
}

func panSE() bool {
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), bottomRight) {
		return true
	} else {
		return false
	}
}
