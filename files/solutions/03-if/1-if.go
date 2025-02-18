package main

import "fmt"

func main() {
	if bigger(5, 10) == 10 && bigger(-105, 16) == 16 && bigger(16, 16) == 16 {
		fmt.Println("All good")
	}
}

func bigger(x, y int) int {
	if x > y { // Solution: return x if it's bigger µ, y otherwise
		return x
	} else {
		return y
	}
}
