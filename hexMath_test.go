package HexEngine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"testing"
)

// I'm trying to learn better testing habits so here's an attempt to do some testing on my hexMath

func TestCoords(t *testing.T) {
	coords := HexMath_Go.NewCoords(1, -1, 0)
	hex, err := HexMath_Go.GetHexCoord(coords)
	if hex == rl.Vector3Zero() || err == nil {
		t.Fatalf("%v", err)
	}
}
