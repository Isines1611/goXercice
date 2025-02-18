package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		ch <- 16
	}()

	fmt.Println(<-ch)
}
