package main

import "fmt"

func main() {
	var x float32 = 10.0
	var y int = 10

	// TODO: Change the line below to print message "x changed type"
	tmpX := int(x)

	if tmpX == y {
		fmt.Println("x changed type")
	} else {
		fmt.Println("x and tmpX are different")
	}
}
