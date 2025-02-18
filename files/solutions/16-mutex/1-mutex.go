package main

import (
	"fmt"
	"sync"
	"time"
)

type safeCounter struct {
	mu sync.Mutex
	v  int
}

func (c *safeCounter) inc() {
	c.mu.Lock()
	defer c.mu.Unlock() // Solution: use defer to not forget to unlock
	c.v++
}

func (c *safeCounter) value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v
}

func main() {
	c := &safeCounter{}

	for i := 0; i < 100; i++ {
		go c.inc()
	}

	time.Sleep(time.Second)
	fmt.Println("Counter:", c.value())
}
