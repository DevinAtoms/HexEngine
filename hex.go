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

func hexPointCorner3D(center rl.Vector3, size float32) [8]rl.Vector3 {
	var corners [8]rl.Vector3
	corners[7] = center
	for i := range corners {
		angleDeg := 60.0 * float64(i)
		angleRad := rl.Deg2rad * angleDeg
		corners[i] = rl.NewVector3(
			center.X+size*float32(math.Cos(angleRad)), 0,
			center.Y+size*float32(math.Sin(angleRad)))
	}
	return corners
}

func hexCorner3D(center rl.Vector3, size float32) [8]rl.Vector3 {
	var corners [8]rl.Vector3
	corners[7] = center
	corners[0] = rl.Vector3Add(center, rl.NewVector3(size, 0, 0))
	for i := range corners[1:6] {
		angleDeg := 60.0 * float64(i)
		angleRad := rl.Deg2rad * angleDeg
		corners[i] = rl.NewVector3(
			center.X+size*float32(math.Cos(angleRad)), 0,
			center.Y+size*float32(math.Sin(angleRad)))
	}
	return corners
}

func hexPointCorner2D(center rl.Vector2, size float32) [8]rl.Vector2 {
	var corners [8]rl.Vector2
	corners[0] = center
	for i := range corners[1:6] {
		angleDeg := 60.0 * float64(i)
		angleRad := rl.Deg2rad * angleDeg
		corners[i] = rl.NewVector2(
			center.X+size*float32(math.Cos(angleRad)),
			center.Y+size*float32(math.Sin(angleRad)))
	}
	return corners
}

func Vector2to3(vector2 rl.Vector2) rl.Vector3 {
	X := vector2.X
	Y := vector2.Y
	Z := -X - Y
	return rl.Vector3{X: X, Y: Y, Z: Z}
}

func Vector3to2(vector3 rl.Vector3) rl.Vector2 {
	X := vector3.X
	Y := vector3.Y
	return rl.Vector2{X: X, Y: Y}
}

func drawHex3D(hex [8]rl.Vector3) {
	for i := range hex[0:6] {
		rl.DrawLine3D(hex[i], hex[i+1], rl.Black)
		rl.DrawLine3D(hex[i], hex[7], rl.Red)
	}
}
