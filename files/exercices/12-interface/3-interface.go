package main

import "fmt"

// TODO: Create a custom error struct then make this struct implement a 'Error' method
// The 'Error' method should return the string: "Cannot divide by zero"

func safeDivision(a, b int) (int, error) {
	// TODO: if division possible return the result with no error
	// otherwise return 0 with the error
}

func main() {
	res, err := safeDivision(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", res)
	}

	res, err = safeDivision(5, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", res)
	}
}
