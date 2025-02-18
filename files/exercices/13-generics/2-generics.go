// HARDER EXERICE
package main

import "fmt"

// TODO: complete the 'stack' struct
type stack struct {
}

// TODO: make 'stack' implement method:
// - 'push': that adds an item to the stack
// - 'pop': that removes and returns the last element of the stack and true
// if 'pop' while the stack is empty then return the zero value and false

func main() {
	s := stack[int]{}

	s.push(10)
	s.push(20)

	val, ok := s.pop()
	fmt.Println(val, ok)

	val, ok = s.pop()
	fmt.Println(val, ok)

	val, ok = s.pop()
	fmt.Println(val, ok)

	s2 := stack[string]{}

	s2.push("Hello")
	s2.push("World")

	val, ok = s2.pop()
	fmt.Println(val, ok)

	val, ok = s2.pop()
	fmt.Println(val, ok)

	val, ok = s2.pop()
	fmt.Println(val, ok)
}
