package main

import "fmt"

func modifyPtr(ptr *int) int {
	*ptr = 10
	return *ptr
}

func main() {
	// TODO: create a pointer and pass it to 'modifyPtr'
	fmt.Println(modifyPtr())
}
