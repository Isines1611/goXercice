// HARDER EXERICE
package main

import "fmt"

type stack[T any] struct {
	items []T
}

func (s *stack[T]) push(item T) {
	s.items = append(s.items, item)
}

func (s *stack[T]) pop() (T, bool) {
	if len(s.items) <= 0 {
		var zeroVal T
		return zeroVal, false
	}

	last := len(s.items) - 1
	item := s.items[last]
	s.items = s.items[:last]
	return item, true
}

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

	val2, ok2 := s2.pop()
	fmt.Println(val2, ok2)

	val2, ok2 = s2.pop()
	fmt.Println(val2, ok2)

	val2, ok2 = s2.pop()
	fmt.Println(val2, ok2)
}
