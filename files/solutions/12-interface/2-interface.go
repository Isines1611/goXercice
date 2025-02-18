package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) String() string { // Solution: use String interface
	return fmt.Sprintf("Name: %v and age: %v", p.name, p.age)
}

func main() {
	alice := person{"alice", 25}
	bob := person{"bob", 27}

	fmt.Println(alice, bob)
}
