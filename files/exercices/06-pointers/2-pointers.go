package main

import "fmt"

var x int = 10

func getPX() **int {
	ptr := &x
	return &ptr
}

func main() {
	ptr := getPX()

	// TODO: using 'ptr', modify the value of 'x' to '200' and print it.

	fmt.Println((**ptr))
}
