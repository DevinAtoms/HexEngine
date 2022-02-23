package main

import rl "github.com/gen2brain/raylib-go/raylib"

func cameraControl(camera *rl.Camera) {
	if !panNE() && !panNW() && !panSE() && !panSW() {
		if rl.IsKeyDown(rl.KeyUp) {
			camera.Target.Y += panSpeed
			camera.Position.Y += panSpeed
		}
		if rl.IsKeyDown(rl.KeyDown) {
			camera.Target.Y -= panSpeed
			camera.Position.Y -= panSpeed
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
			camera.Target.X += panSpeed
			camera.Position.X += panSpeed
		}
		if rl.IsKeyDown(rl.KeyA) || panW() {
			camera.Target.X -= panSpeed
			camera.Position.X -= panSpeed
		}
	}
	if panNE() {
		camera.Target.Z += panSpeed
		camera.Position.Z += panSpeed
		camera.Target.X += panSpeed
		camera.Position.X += panSpeed
	}
	if panSE() {
		camera.Target.Z -= panSpeed
		camera.Position.Z -= panSpeed
		camera.Target.X += panSpeed
		camera.Position.X += panSpeed
	}
	if panNW() {
		camera.Target.Z += panSpeed
		camera.Position.Z += panSpeed
		camera.Target.X -= panSpeed
		camera.Position.X -= panSpeed
	}
	if panSW() {
		camera.Target.Z -= panSpeed
		camera.Position.Z -= panSpeed
		camera.Target.X -= panSpeed
		camera.Position.X -= panSpeed
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
