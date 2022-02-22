package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"math"
)

type hexDim func(int32, int32) int32
type hexCoord func(int32, int32) float32
type hexPoints func(rl.Vector3, float32) [7]rl.Vector3

type hex struct {
	width  hexDim
	height hexDim
	points hexPoints
}

func hexPointCorner3D(center rl.Vector3, size float32) [7]rl.Vector3 {
	var corners [7]rl.Vector3
	for i := range corners {
		angleDeg := 60.0 * float64(i)
		angleRad := math.Pi / 180 * angleDeg
		corners[i] = rl.NewVector3(
			center.X+size*float32(math.Cos(angleRad)), 0,
			center.Y+size*float32(math.Sin(angleRad)))
	}
	return corners
}

func hexPointCorner2D(center rl.Vector2, size float32) [7]rl.Vector2 {
	var corners [7]rl.Vector2
	for i := range corners {
		angleDeg := 60.0 * float64(i)
		angleRad := math.Pi / 180 * angleDeg
		corners[i] = rl.NewVector2(
			center.X+size*float32(math.Cos(angleRad)),
			center.Y+size*float32(math.Sin(angleRad)))
	}
	corners[6] = corners[0]
	return corners
}
