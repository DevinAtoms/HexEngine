package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"math"
)

func hexPointCorner(center rl.Vector2, size float32, c float64) [5]rl.Vector2 {
	var cornerIndex = []float64{0, 60, 120, 180, 240, 300}
	var corners = [5]rl.Vector2{}

	for i, s := range cornerIndex {
		corners[i] = rl.NewVector2(
			float32(math.Cos(rl.Pi/180*(60*s))),
			float32(math.Sin(rl.Pi/180*(60*s))))
	}

	return corners
}
