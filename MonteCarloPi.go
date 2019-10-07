package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/schollz/progressbar/v2"
	"math"
	"time"
)

func Pi(iterations int, c *gin.Context) {
	bar := progressbar.NewOptions(100, progressbar.OptionEnableColorCodes(true), progressbar.OptionSetWriter(c.Writer))
	k := 0
	var e float64
	t0 := time.Now()
	for n := 1; n < iterations; n++ {
		x := randFloat(-1.0, 1.0)
		y := randFloat(-1.0, 1.0)
		if distanceTo0(x, y) <= 1.0 {
			k += 1
		}
		e = float64(k) / float64(n) * 4.0
		if n%(iterations/100) == 0 {
			bar.Describe(fmt.Sprintf("The Estimate: %32.32f The Difference: %32.32f After %s", e, math.Pi-e, time.Since(t0).String()))
			_ = bar.Add(1)
		}
	}
	fmt.Fprintf(c.Writer, "\nThe Final Estimate: %v The Final Difference: %v Done after %s\n", e, math.Pi-e, time.Since(t0).String())
}