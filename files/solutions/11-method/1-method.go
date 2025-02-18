package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	alice := person{"alice", "smith"}
	bob := person{"bob", "williams"}

	alice.greet()
	bob.greet()
}

func (p person) greet() {
	fmt.Printf("Hello, I am %s %s\n", p.firstName, p.lastName) // Solution: use format to make it easier
}
