package main

import (
	"fmt"
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
		select {
		case msg := <-ch1: // Solution: get from the good channel
			fmt.Println("Received from intCh:", msg)
		case msg := <-ch2:
			fmt.Println("Received from strCh:", msg)
		}
	}
}
