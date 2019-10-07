package main

import (
	"fmt"
	"github.com/schollz/progressbar/v2"
	"log"
	"time"
)

//func Integration1D(f func(float64) float64, a, b float64, N int) float64 {
//
//
//	var t0 int64
//	var tEnd int64
//	var e float64
//	t0 = time.Now().UnixNano()
//	ba := b - a
//
//	R := arrayRandFloat(a, b, N)
//	fR := make([]float64, N)
//
//	for i, v := range R {
//		fR[i] = f(v)
//	}
//	sumFR := floatArraySum(fR, 0, N-1)
//
//	e = ba * (1.0 / float64(N)) * sumFR
//
//	tEnd = time.Now().UnixNano()
//
//	fmt.Println("The Final Estimate:", e, "Done after", tEnd-t0, "nanoseconds")
//
//	return e
//}

func Integration(f func(float64) float64, a, b float64, N int) float64 {

	bar := progressbar.NewOptions(100, progressbar.OptionEnableColorCodes(true), progressbar.OptionSetWriter(log.Writer()))
	ba := b - a

	var e float64
	t0 := time.Now()

	var k float64

	for n := 0; n < N; n++ {
		k += f(randFloat(a, b))

		if n%(N/100) == 0 {
			e = ba * (1.0 / float64(n)) * k
			bar.Describe(fmt.Sprintf("The Estimate: %32.32f", e))
			_ = bar.Add(1)
		}
	}

	e = ba * (1.0 / float64(N)) * k

	_, _ = fmt.Fprintln(log.Writer(), "\nThe Final Estimate:", e, "Done after", time.Since(t0).String())

	return e
}
