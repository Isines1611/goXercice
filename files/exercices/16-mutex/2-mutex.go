package main

import (
	"fmt"
	"time"
)

type bankAccount struct {
	// TODO: Complete the struct with a mutex
}

func (a *bankAccount) deposit(amount int) {
	// TODO: complete the method to deposit money SAFELY
}

func (a *bankAccount) withdraw(amount int) {
	// TODO: complete the method to withdraw money SAFELY ONLY if enough money, otherwise do nothing
}

func (a *bankAccount) getBalance() int {
	// TODO: complete the method to return the balance SAFELY
}

func main() {
	account := &bankAccount{1000} // TODO: modify this line to suit your 'bankAccount' implementation

	go account.deposit(500)
	go account.withdraw(200)

	time.Sleep(time.Second)

	fmt.Println(account.getBalance())
}
