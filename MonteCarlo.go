package main

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math"
	math_rand "math/rand"
)

func init() {

	var b [8]byte
	_, err := crypto_rand.Read(b[:])
	if err != nil {
		panic("cannot seed math/rand package with cryptographically secure random number generator")
	}
	math_rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
}

func main() {
	//Pi(10000000000)

	e := Integration(f, 0, math.Pi, 10000000)
	fmt.Println(e, "Diff:", 2.0-e)

}

func f(x float64) (y float64) {
	y = math.Sin(x)
	return
}
