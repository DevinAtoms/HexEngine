package main

import (
	"errors"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

const (
	Apothem = 0.5
	HexRad  = 0.5774
)

type HexTile struct {
	Model  rl.Model
	Coords HexCoord
	Points [7]rl.Vector3
	HexMat rl.Matrix
}

type HexCoord struct {
	Q float32
	R float32
	S float32
}

var OriginHex HexTile

func LoadHex() *HexTile {
	tile := HexTile{
		Model:  rl.LoadModel("assets/grass.obj"),
		HexMat: rl.Matrix{},
	}
	return &tile
}

func HexCorner3D(center rl.Vector3, size float32) [7]rl.Vector3 {
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

func DrawOriginHex(origin *HexTile) {
	origin.Points = HexCorner3D(rl.Vector3Zero(), HexRad)
	rl.DrawModelEx(origin.Model, rl.Vector3Zero(), rl.NewVector3(0, 1, 0), 0, rl.NewVector3(1, 1, 1), rl.Gray)
}

func Wireframe(center rl.Vector3, size float32) {
	corners := HexCorner3D(center, size)
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

func DrawHex(loc rl.Vector3, tile HexTile) {
	tile.Points = HexCorner3D(loc, HexRad)
	rl.DrawModelEx(tile.Model, loc, rl.NewVector3(0, 1, 0), 0, rl.NewVector3(1, 1, 1), rl.Gray)
}

func GetHexCoord(h HexCoord) (rl.Vector3, error) {
	var dir int32
	v := rl.Vector3{}

	if h.Q+h.R+h.S == 0 {
		if h.Q == 0 {
			if h.R < 0 {
				dir = 1
			} else {
				dir = 4
			}
		} else if h.R == 0 {
			if h.S < 0 {
				dir = 6
			} else {
				dir = 3
			}
		} else {
			if h.Q < 0 {
				dir = 5
			} else {
				dir = 2
			}
		}
	} else {
		return rl.Vector3Zero(), errors.New("Q+R+S !=0")
	}

	angleDeg := float64(dir * 60)
	angleRadCos := float32(math.Cos(rl.Deg2rad * angleDeg))
	angleRadSin := float32(math.Sin(rl.Deg2rad * angleDeg))
	loc := rl.Vector3Multiply(rl.NewVector3(v.X+.433*angleRadCos, 0, v.Z+.433*angleRadSin), 2.0)
	return loc, nil
}

func NewCoords(Q float32, R float32, S float32) HexCoord {
	return HexCoord{Q: Q, R: R, S: S}
}
