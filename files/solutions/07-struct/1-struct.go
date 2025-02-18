package main

import "fmt"

type color struct {
	red, blue, green int // Solution: declare as int
}

func main() {
	red := color{255, 0, 0} // Solution: declare them with the RGB values
	blue := color{0, 255, 0}
	green := color{0, 0, 255}

	if red.red == 255 && red.blue == 0 && red.green == 0 &&
		blue.red == 0 && blue.blue == 255 && blue.green == 0 &&
		green.red == 0 && green.blue == 0 && green.green == 255 {
		fmt.Println("All good.")
	}
}
