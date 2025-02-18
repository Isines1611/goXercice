package main

import "fmt"

type Number interface {
	int | int32 | int64 | float32 | float64
}

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

func sumNb[T Number](nb []T) T {
	var res T
	for i := range nb {
		res += nb[i]
	}
	return res
}
