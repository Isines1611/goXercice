package main

import "fmt"

func main() {
	if bigger(5, 10) == 10 && bigger(-105, 16) == 16 && bigger(16, 16) == 16 {
		fmt.Println("All good")
	}
}

func bigger(x, y int) int {
	// TODO: complete this function to return the biggest number
	// If both are equal, return one of them
	// Do not use:
	// - other function
	// - other imports
	// - other variables than x and y
}
