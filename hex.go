package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"math"
)

type hexTile struct {
	hexPoly rl.Mesh
	center  rl.Vector2
	corners [5]rl.Vector2
}

func hexPointCorner(center rl.Vector3, size float32) [5]rl.Vector3 {
	var corners [5]rl.Vector3
	for i := range corners {
		angleDeg := 60.0 * float64(i)
		angleRad := math.Pi / 180 * angleDeg
		corners[i] = rl.NewVector3(
			center.X+size*float32(math.Cos(angleRad)),
			center.Y+size*float32(math.Sin(angleRad)),
			0)
	}
	return corners
}
