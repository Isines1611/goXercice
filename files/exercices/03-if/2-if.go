package main

import "fmt"

func main() {
	if abs(5) == 5 && abs(-105) == -105 && abs(16) == 16 {
		fmt.Println("All good")
	}
}

func abs(x int) int {
	// TODO: complete this if statement to return the absolute value
	// Only change the three dots (...) 
	// Do not use:
	// - other function
	// - other imports

	if ... {
		return -x
	}
	return x
}
