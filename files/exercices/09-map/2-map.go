package main

import "fmt"

func getMap() map[string]int {
	grades := make(map[string]int)

	grades["Bob"] = 30
	grades["Greg"] = -25
	grades["Alice"] = 80
	grades["Charlie"] = 50
	grades["Yan"] = 1000

	return grades
}

func main() {
	grades := getMap()

	// TODO: delete the unwanted grades from the arrays.
	// Grades must be int between 0 and 100
	// DO not:
	// - change the function 'getMap'

	fmt.Println(grades)
}
