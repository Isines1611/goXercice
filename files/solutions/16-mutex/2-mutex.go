package main

import (
	"fmt"
	"sync"
	"time"
)

type bankAccount struct {
	mu      sync.Mutex
	balance int
}

func (a *bankAccount) deposit(amount int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.balance += amount
}

func (a *bankAccount) withdraw(amount int) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.balance >= amount {
		a.balance -= amount
	}
}

func (a *bankAccount) getBalance() int {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.balance
}

func main() {
	account := &bankAccount{balance: 1000}

	go account.deposit(500)
	go account.withdraw(200)

	time.Sleep(time.Second)

	fmt.Println(account.getBalance())
}
