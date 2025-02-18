package main

import "fmt"

type person struct {
	name string
	age  int
}

// TODO: Implement 'fmt.Stringer' interface for 'person', making them say: "Name: <name> and age: <age>"

func main() {
	alice := person{"alice", 25}
	bob := person{"bob", 27}

	fmt.Println(alice)
	fmt.Println(bob)
}
