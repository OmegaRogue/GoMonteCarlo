package main

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"math"
	math_rand "math/rand"
	"os"
	"strconv"

	_ "github.com/heroku/x/hmetrics/onload"
	"log"
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

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	//router := gin.New()
	//router.Use(gin.Logger())
	//
	//router.GET("/pi", func(c *gin.Context) {
	//	N, _ := strconv.ParseInt(c.Query("N"), 10, 64)
	//	Pi(int(N))
	//
	//})
	//router.GET("/integral", func(c *gin.Context) {
	//	var N, _ = strconv.ParseInt(c.Query("N"), 10, 64)
	//	var a, _ = strconv.ParseFloat(c.DefaultQuery("min", "0"), 64)
	//	var b, _ = strconv.ParseFloat(c.Query("max"), 64)
	//
	//	Integration(f, a, b, int(N))
	//
	//})
	//router.Run(":" + port)

	var N, _ = strconv.ParseInt(os.Getenv("N"), 10, 64)
	var a, _ = strconv.ParseFloat(os.Getenv("min"), 64)
	var b, _ = strconv.ParseFloat(os.Getenv("max"), 64)

	switch os.Getenv("Mode") {
	case "pi":
		Pi(int(N))
	case "integral":
		Integration(f, a, b, int(N))

	}

}

func f(x float64) (y float64) {
	y = 3*math.Pow(x, 2) - 8*x + 7
	return
}
