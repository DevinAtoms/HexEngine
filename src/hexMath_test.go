package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"testing"
)

// I'm trying to learn better testing habits so here's an attempt to do some testing on my hexMath

func testCoords(t *testing.T) {
	coords := NewCoords(1, -1, 0)
	hex, err := GetHexCoord(coords)
	if hex == rl.Vector3Zero() || err == nil {
		t.Fatalf("%v", err)
	}
}
