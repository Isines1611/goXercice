package main

import "fmt"

type vector2 struct {
	x int
	y int
}

var vec vector2 = vector2{1, 2}

func main() {
	vec.x = 5 // Solution: access fields with '.'
	vec.y = 16

	fmt.Println(vec)
}
