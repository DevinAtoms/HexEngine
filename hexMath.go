package main

import (
	"math"

	"github.com/gen2brain/raylib-go/raylib"
)

const apothem = 0.433
const hexRad = .5

type hexTile struct {
	model  rl.Model
	coords hexCoord
	points [7]rl.Vector3
	hexMat rl.Matrix
}

type hexCoord struct {
	Q float32
	R float32
	S float32
}

func loadHex() *hexTile {
	tile := hexTile{
		model:  rl.LoadModel("assets/newhex.obj"),
		hexMat: rl.Matrix{},
	}
	return &tile
}

func hexCorner3D(center rl.Vector3, size float32) [7]rl.Vector3 {
	var corners [7]rl.Vector3
	points := corners[0:6]
	for i := range points {
		angleDeg := 60.0*float64(i) + 30
		angleRadCos := float32(math.Cos(rl.Deg2rad * angleDeg))
		angleRadSin := float32(math.Sin(rl.Deg2rad * angleDeg))
		corners[i] = rl.NewVector3(
			center.X+size*angleRadCos, 0,
			center.Z+size*angleRadSin)
	}
	return corners
}

func drawOriginHex() {
	tile := originHex
	tile.points = hexCorner3D(rl.Vector3Zero(), 0.5)
	rl.DrawModelEx(tile.model, rl.Vector3Zero(), rl.NewVector3(0, 1, 0), 0, rl.NewVector3(1, 1, 1), rl.Gray)
}

func wireframe(tile *hexTile) {

	corners := tile.points

	for i := range tile.points {
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

func drawHex(loc rl.Vector3) {
	tile := originHex
	tile.points = hexCorner3D(loc, 0.5)
	wireframe(tile)
	rl.DrawModelEx(tile.model, loc, rl.NewVector3(0, 1, 0), 0, rl.NewVector3(1, 1, 1), rl.Gray)
}

func hexCoords(coord rl.Vector3) rl.Vector3 {

	dist := float32(hexRad - apothem)

	angleDeg := 30.0
	angleRadCos := float32(math.Cos(rl.Deg2rad * angleDeg))
	angleRadSin := float32(math.Sin(rl.Deg2rad * angleDeg))
	R := rl.NewVector3(angleRadSin-dist, 0, angleRadCos-dist-(hexRad/10))

	//cubeQ := apothem*2
	//cubeR :=
	//
	//gridQ := coord.X
	//gridR := coord.Y
	//gridS := coord.Z

	// Q, -R, +S
	rl.NewVector3(apothem, 0, .75)
	// Q, +R, -S
	rl.NewVector3(-apothem, 0, -.75)

	// +Q, R, -S
	rl.NewVector3(apothem*2, 0, 0)
	// -Q, R, +S
	rl.NewVector3(-apothem*2, 0, 0)

	// -Q, +R, S
	rl.NewVector3(apothem, 0, -.75)
	// +Q, -R, S
	rl.NewVector3(-apothem, 0, .75)
	return R
}
