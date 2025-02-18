package main

import "fmt"

func main() {
	if calculatePrice(10, 5) == 20 &&
		calculatePrice(-5, 5) == 0 &&
		calculatePrice(6, 0) == 12 &&
		calculatePrice(8, 1) == 14 &&
		calculatePrice(1, 1) == 0 &&
		calculatePrice(10, 2) == 10 &&
		calculatePrice(4, 2) == 0 &&
		calculatePrice(14, 3) == 14 &&
		calculatePrice(-22, 3) == 0 {
		fmt.Println("All good.")
	}
}

func calculatePrice(apple, fidelity int) int {
	if apple < 0 {
		return 0
	}

	switch fidelity {
	case 1:
		return (apple - 1) * 2
	case 2:
		return (apple - 5) * 2
	case 3:
		return apple
	default:
		return apple * 2
	}
}
