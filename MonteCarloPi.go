package main

import (
	"fmt"
	"math"
	"time"
)

func Pi(iterations int) {
	k := 0
	var t0 int64
	var tEnd int64
	var e float64
	t0 = time.Now().UnixNano()
	for n := 1; n < iterations; n++ {
		x := randFloat(-1.0, 1.0)
		y := randFloat(-1.0, 1.0)
		if distanceTo0(x, y) <= 1.0 {
			k += 1
		}
		e = float64(k) / float64(n) * 4.0
		if n%(iterations/100) == 0 {
			fmt.Println("The Estimate:", e, "The Difference:", math.Pi-e, "After", time.Now().UnixNano()-t0, "nanoseconds")
		}
	}
	tEnd = time.Now().UnixNano()
	fmt.Println("The Final Estimate:", e, "The Final Difference:", math.Pi-e, "Done after", tEnd-t0, "nanoseconds")

}
