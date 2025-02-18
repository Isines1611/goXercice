package main

import "fmt"

func modifyPtr(ptr *int) int {
	*ptr = 10
	return *ptr
}

func main() {
	i := 16
	fmt.Println(modifyPtr(&i)) // Solution: use '&' to create pointer
}
