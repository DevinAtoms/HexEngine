package HexMath

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
	Coord  rl.Vector2
	Points [7]rl.Vector3
}

type HexCoord struct {
	Q float32
	R float32
	S float32
}

var OriginHex HexTile

func HexCorner3D(center rl.Vector3, size float32) [7]rl.Vector3 {
	var corners [7]rl.Vector3
	points := corners[0:6]
	for i := range points {
		angleDeg := 60.0*float64(i) + 30
		angleRadCos := float32(math.Cos(rl.Deg2rad * angleDeg))
		angleRadSin := float32(math.Sin(rl.Deg2rad * angleDeg))
		corners[i] = rl.NewVector3(
			center.X+size*angleRadCos, 0,
			center.Z-size*angleRadSin)
	}
	return corners
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
