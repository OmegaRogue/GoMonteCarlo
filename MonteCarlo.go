package main

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math"
	math_rand "math/rand"
	"os"
	"strconv"
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

	N, _:= strconv.ParseInt(os.Getenv("N"),10,64)
	a, _:= strconv.ParseFloat(os.Getenv("min"),64)
	b, _:= strconv.ParseFloat(os.Getenv("max"),64)

	switch os.Getenv("Mode") {
	case "Pi":
		Pi(int(N))
	case "Integral":
		e := Integration(f, a, b, int(N))
		fmt.Println(e, "Diff:", 2.0-e)
	default:
		Pi(int(N))
	}



}

func f(x float64) (y float64) {
	y = math.Sin(x)
	return
}
