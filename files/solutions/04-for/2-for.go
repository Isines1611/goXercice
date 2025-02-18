package main

import "fmt"

func main() {
	printNumber()
}

func printNumber() {
	i := 1 // Solution: declare i outside of for
	for i < 11 {
		fmt.Println(i)
		i++ // Solution: increment i outside of for
	}
}
