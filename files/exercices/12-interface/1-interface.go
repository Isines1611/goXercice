package main

import "fmt"

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

type triangle struct {
	base, height float64
}

// TODO: Define the 'areaCalculator' interface and make the 3 struct (rectange, circle, circle) implement the interface with 'area' method

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
