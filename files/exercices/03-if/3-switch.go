package main

import "fmt"

func main() {
	if categorizeNumber(-50) == "Negative" && categorizeNumber(-16) == "Negative" && categorizeNumber(0) == "Zero" && categorizeNumber(16) == "Positive" && categorizeNumber(125) == "Positive" {
		fmt.Println("All good")
	}
}

func categorizeNumber(number int) string {
	// TODO: Use a switch statement to return:
	// - "Positive" if n is greater than 0
	// - "Negative" if n is less than 0
	// - "Zero" if n is exactly 0
}
