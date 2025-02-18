package main

import (
	"fmt"
	"time"
)

func printMessage(message string) {
	time.Sleep(2 * time.Second)
	fmt.Println(message)
}

func main() {
	go printMessage("Hello from the goroutine")

	// Do not change the line below
	time.Sleep(3 * time.Second)
}
