package main

import "fmt"

func main() {
	printNumber()
}

func printNumber() {
	for i := 1; i < 11; i++ { // Solution: Careful on for declaration to include 10
		fmt.Println(i)
	}
}
