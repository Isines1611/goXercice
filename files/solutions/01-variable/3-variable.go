package main

import "fmt"

func main() {
	var x float32 = 10.0
	var y int = 10

	tmpX := int(x) // Solution: convert type of x to int

	if tmpX == y {
		fmt.Println("x changed type")
	} else {
		fmt.Println("x and tmpX are different")
	}
}
