package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"math"
)

type hex struct {
	center rl.Vector3
}

func hexDirection(direction int32) rl.Vector3 {
	xUp := rl.NewVector3(0, -1, 1)
	xDown := rl.NewVector3(0, 1, -1)
	yUp := rl.NewVector3(-1, 0, 1)
	yDown := rl.NewVector3(1, 0, -1)
	zUp := rl.NewVector3(1, -1, 0)
	zDown := rl.NewVector3(-1, 1, 0)
	dir := []rl.Vector3{xUp, zUp, yDown, xDown, zDown, yUp}
	return dir[direction]
}

func hexNeighbor(h hex, direction int32) hex {
	neighborHex := hex{center: rl.Vector3Add(h.center, hexDirection(direction))}
	return neighborHex
}

func hexToPixel(h hex, size float32) rl.Vector2 {
	x := size * (3 / 2 * h.center.X)
	y := size * (float32(math.Sqrt(3))/2*h.center.X + float32(math.Sqrt(3))*h.center.Y)
	point := rl.NewVector2(x, y)
	return point
}

func pixelToHex(point rl.Vector2, size float32) {
	x := (2 / 3 * point.X) / size
	y := (-1/3*point.X + float32(math.Sqrt(3)/3)*point.Y) / size
	rl.Vector3Normalize()
}

func hexPointCorner3D(center rl.Vector3, size float32) []rl.Vector3 {
	var corners []rl.Vector3
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
