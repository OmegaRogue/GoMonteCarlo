package main

import (
	"fmt"
	"time"
)

func Integration(f func(float64) float64, a, b float64, N int) float64 {

	ba := b - a

	var e float64
	t0 := time.Now()

	c := make(chan float64, N)
	for i := 0; i < N; i++ {
		go integrate(f, a, b, c)
	}

	e = ba * (1.0 / float64(N)) * <-c

	fmt.Println("The Final Estimate:", e, "Done after", time.Since(t0).String())

	return e
}

func integrate(f func(float64) float64, a, b float64, c chan float64) {
	c <- f(randFloat(a, b)) + <-c
}
