package main

import (
	"fmt"
	"math"
)

type areaCalculator interface { // Solution: Define the interface
	area() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

type triangle struct {
	base, height float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (t triangle) area() float64 {
	return t.base * t.height * 1 / 2
}

func main() {
	var shapes []areaCalculator = []areaCalculator{
		rectangle{width: 5, height: 10},
		circle{radius: 5},
		triangle{5, 8},
	}

	for _, shape := range shapes {
		fmt.Println(shape.area())
	}
}
