package main

import (
	"fmt"
	"time"
)

type safeCounter struct {
	// TODO: complete the struct with a mutex
}

func (c *safeCounter) inc() {
	// TODO: Complete the method to add 1 to the counter SAFELY
}

// Value returns the current value of the counter.
func (c *safeCounter) value() int {
	// TODO: Complete the method to return the value of the counter SAFELY
}

func main() {
	c := &safeCounter{}

	for i := 0; i < 100; i++ {
		go c.inc()
	}

	time.Sleep(time.Second)
	fmt.Println("Counter:", c.value())
}
