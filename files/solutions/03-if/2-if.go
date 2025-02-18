package main

import "fmt"

func main() {
	if abs(5) == 5 && abs(-105) == -105 && abs(16) == 16 {
		fmt.Println("All good")
	}
}

func abs(x int) int {
	if x > 0 { // Solution: if it's a negative number, we should multiply by '-1'
		return -x
	}
	return x
}
