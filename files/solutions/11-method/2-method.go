package main

import "fmt"

type bank struct {
	owner   string
	balance float64
}

func (b bank) deposit(amount int) {
	b.balance += float64(amount) // Solution: remember type conversion
	fmt.Println(b.balance)
}

func (b bank) withdraw(amount int) {
	if b.balance-float64(amount) < 0 {
		fmt.Println("Insufficient balance")
		return
	}

	b.balance -= float64(amount)
	fmt.Println(b.balance)
}

func main() {
	acc := bank{"Alice", 1000.0}

	acc.deposit(500)
	acc.withdraw(200)
	acc.withdraw(1500)

}
