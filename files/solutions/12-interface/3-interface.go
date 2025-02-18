package main

import "fmt"

type DivideZero struct{}

func (e DivideZero) Error() string { // Solution: use Error interface
	return "Cannot divide by zero"
}

func safeDivision(a, b int) (int, error) {
	if b == 0 {
		return 0, DivideZero{} // Solution: use custom error
	}

	return a / b, nil
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
