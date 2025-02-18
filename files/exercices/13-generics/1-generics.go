package main

import "fmt"

// TODO: create the needed struct to use generics of the sumNb which returns the sum of all the numbers in the given array

func main() {
	nb1 := []int{1, 2, 3, 4, 5}
	nb2 := []int32{1, 2, 3, 4, 5}
	nb3 := []int64{1, 2, 3, 4, 5}
	nb4 := []float32{1.0, 2.0, 3.0, 4.0, 5.0}
	nb5 := []float64{1.0, 2.0, 3.0, 4.0, 5.0}

	fmt.Println(sumNb(nb1))
	fmt.Println(sumNb(nb2))
	fmt.Println(sumNb(nb3))
	fmt.Println(sumNb(nb4))
	fmt.Println(sumNb(nb5))
}

func sumNb() {
	// TODO: complete this function	definition and body
}
