package main

import (
	"math"
	"math/rand"
)

func distanceTo0(x, y float64) float64 {
	x2 := x * x
	y2 := y * y
	d := math.Sqrt(x2 + y2)
	return d
}

func randFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
