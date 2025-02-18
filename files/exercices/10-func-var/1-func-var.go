package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	// TODO: define pow as the function power => x^n
	// pow :=

	if compute(pow) == compute(math.Pow) {
		fmt.Println("All good")
	}
}
