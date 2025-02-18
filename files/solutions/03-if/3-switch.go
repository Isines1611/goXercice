package main

import "fmt"

func main() {
	if categorizeNumber(-50) == "Negative" && categorizeNumber(-16) == "Negative" && categorizeNumber(0) == "Zero" && categorizeNumber(16) == "Positive" && categorizeNumber(125) == "Positive" {
		fmt.Println("All good")
	}
}

func categorizeNumber(number int) string {
	switch { // Solution: Don't switch over a given variable
	case number > 0:
		return "Postivie"
	case number < 0:
		return "Negative"
	default:
		return "Zero"
	}
}
