package main

import "fmt"

var x int = 10

func getPX() **int {
	ptr := &x
	return &ptr
}

func main() {
	ptr := getPX()

	**ptr = 200 // // Solution: use 2 '*' because the had a pointer to a pointer of int

	fmt.Println((**ptr))
}
