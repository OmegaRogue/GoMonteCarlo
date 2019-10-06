package main

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math"
	math_rand "math/rand"
	"time"
)

var t_0 int64
var t_end int64

func init() {

	var b [8]byte
	_, err := crypto_rand.Read(b[:])
	if err != nil {
		panic("cannot seed math/rand package with cryptographically secure random number generator")
	}
	math_rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
}

func main() {
	k := 0

	var e float64
	t_0 = time.Now().UnixNano()
	for n := 1; n < 10000000000; n++ {
		x := math_rand.Float64()
		y := math_rand.Float64()
		if distanceTo0(x, y) <= 1.0 {
			k += 1
		}
		e = float64(k) / float64(n) * 4.0
		if n%100000000 == 0 {
			fmt.Println("The Estimate:", e, "The Difference:", math.Pi-e, "After", time.Now().UnixNano()-t_0, "nanoseconds")
		}
	}
	t_end = time.Now().UnixNano()
	fmt.Println("The Final Estimate:", e, "The Final Difference:", math.Pi-e, "Done after", t_end-t_0, "nanoseconds")

}

func distanceTo0(x, y float64) float64 {
	x2 := x * x
	y2 := y * y
	d := math.Sqrt(x2 + y2)
	return d
}
