package main

import "fmt"

func main() {

	var matrix [][]int = make([][]int, 9) // Solution: Initialize matrix

	matrix[0] = []int{0} // Solution: access matrix and set the values
	matrix[1] = []int{1}
	matrix[2] = []int{0}
	matrix[3] = []int{0}
	matrix[4] = []int{0}
	matrix[5] = []int{0}
	matrix[6] = []int{1}
	matrix[7] = []int{0}
	matrix[8] = []int{1}

	fmt.Println(matrix)
}
