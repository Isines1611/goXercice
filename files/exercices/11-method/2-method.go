package main

type bank struct {
	owner   string
	balance float64
}

// TODO: create the 'deposit' and 'withdraw' method:
// 'deposit' should add the given amount to the balance and print ONLY the new balance
// 'withdraw' should withdraw the amount given and print ONLY the new balance. If not enough balance, then print: "Insufficient balance"

func main() {
	acc := bank{"Alice", 1000.0}

	acc.deposit(500)
	acc.withdraw(200)
	acc.withdraw(1500)
}
