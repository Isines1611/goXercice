package main

import "fmt"

// TODO: make a closure that increments a counter
func counter() ...

func main() {
	count := counter()

	fmt.Println(count())
	fmt.Println(count())
	fmt.Println(count())
}
