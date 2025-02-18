package main

import "fmt"

func main() {
	fmt.Println("Start")
	printMessages()
	fmt.Println("End")
}

func printMessages() {
	defer fmt.Println("Thrid") // Solution: defer is push on stack, so do it back order
	defer fmt.Println("Second")
	defer fmt.Println("First")
}
