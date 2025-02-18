package main

import "fmt"

func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	count := counter()

	fmt.Println(count())
	fmt.Println(count())
	fmt.Println(count())
}
