package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(squareRoot(4))
	piNumber()
}

func squareRoot(number float64) float64 {
	return math.Sqrt(number)
}

func piNumber() {
	fmt.Println(math.Pi)
}
