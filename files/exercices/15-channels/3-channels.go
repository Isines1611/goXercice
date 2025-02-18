package main

import (
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		ch1 <- "first"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		ch2 <- "second"
	}()

	for i := 0; i < 2; i++ {
		// TODO: in the for loop, using 'select', print both message, keeping order: "first", "second"
	}
}
