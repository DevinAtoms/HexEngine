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

func hexCorner3D(center rl.Vector3, size float32) [7]rl.Vector3 {
	var corners [7]rl.Vector3
	points := corners[0:6]
	for i := range points {
		angleDeg := 60.0 * float64(i)
		angleRadCos := float32(math.Cos(rl.Deg2rad * angleDeg))
		angleRadSin := float32(math.Sin(rl.Deg2rad * angleDeg))
		corners[i] = rl.NewVector3(
			center.X+size*angleRadCos, center.Y+size*angleRadSin,
			0)
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

func drawHex3D(hex [7]rl.Vector3) {
	center := hex[6]
	for i := range hex {
		if i < 5 {
			rl.DrawLine3D(hex[i], hex[i+1], rl.Black)
			rl.DrawLine3D(hex[i], center, rl.Red)
			rl.DrawText("test", int32(Vector3to2(hex[i]).X), int32(Vector3to2(hex[i]).Y), 8, rl.Black)
		} else if i == 5 {
			rl.DrawLine3D(hex[5], hex[0], rl.Black)
			rl.DrawLine3D(hex[i], center, rl.Red)
		}
	}
	//rl.DrawSphere(hex[0], .5, rl.Red)
	//rl.DrawSphere(hex[1], .5, rl.Blue)
	//rl.DrawSphere(hex[2], .5, rl.Green)
	//rl.DrawSphere(hex[3], .5, rl.Yellow)
	//rl.DrawSphere(hex[4], .5, rl.Orange)
	//rl.DrawSphere(hex[5], .5, rl.Brown)
	//rl.DrawSphere(center, .5, rl.White)
}
