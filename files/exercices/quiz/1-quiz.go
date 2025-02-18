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
	// TODO: complete the 'calculatePrice' function, it must return a price based on the amount of apple and fidelity points:
	// - if there's 1 fidelity points, then one apple is 'free'
	// - if there's 2 fidelity points, then 5 apples are 'free'
	// - if there's 3 fidelity points, then half apples are 'free'
	// the price price of an apple is 2
	// Notice, if having 4 apples and 2 points, then 5 are free, so the price is 0
}
