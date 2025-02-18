package main

import "fmt"

func main() {
	countdown(5)
}

func countdown(x int) {
	for x > 0 {
		fmt.Println(x)
		x-- // Solution: decrease x
	}
}
