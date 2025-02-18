package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	pow := func(x, y float64) float64 {
		return math.Pow(x, y) // Solution: use math.Pow to be easier
	}

	if compute(pow) == compute(math.Pow) {
		fmt.Println("All good")
	}
}
