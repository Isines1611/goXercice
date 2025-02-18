package main

import (
	"errors"
	"fmt"
)

var errFund = errors.New("Error funds.")

type bank interface {
	deposit(amount float64)
	withdraw(amount float64) error
	transfer(to *account, amount float64) error
}

type account struct {
	owner       string
	balance     float64
	transaction []string
}

func newAccount(owner string, initiale float64) *account {
	return &account{
		owner:       owner,
		balance:     initiale,
		transaction: []string{fmt.Sprintf("Account create with: $%.2f", initiale)},
	}
}

func (a *account) deposit(amount float64) {
	if amount <= 0 {
		fmt.Println("Deposit should be greater than zero")
		return
	}

	a.balance += amount
	a.transaction = append(a.transaction, fmt.Sprintf("Deposite of: $%.2f", amount))
}

func (a *account) withdraw(amount float64) error {
	if amount <= 0 {
		fmt.Println("Withdraw should be greater than zero")
		return errFund
	}

	if amount > a.balance {
		return errFund
	}

	a.balance -= amount
	a.transaction = append(a.transaction, fmt.Sprintf("Withdrew: $%.2f", amount))
	return nil
}

func (a *account) transfer(to *account, amount float64) error {
	err := a.withdraw(amount)
	if err != nil {
		return err
	}

	to.deposit(amount)
	a.transaction = append(a.transaction, fmt.Sprintf("Transferred: $%.2f to %s", amount, to.owner))
	to.transaction = append(to.transaction, fmt.Sprintf("Received: $%.2f from %s", amount, a.owner))
	return nil
}

func printTransaction(a bank) {
	if acc, ok := a.(*account); ok {
		fmt.Printf("\nTransaction history for %s:\n", acc.owner)
		for _, transaction := range acc.transaction {
			fmt.Println(transaction)
		}
	} else {
		fmt.Println("❌ Unknown account type")
	}
}

func main() {
	acc1 := newAccount("Alice", 1000)
	acc2 := newAccount("Bob", 500)

	// Perform transactions
	acc1.deposit(200)
	acc1.withdraw(50)
	err := acc1.transfer(acc2, 300)
	if err != nil {
		fmt.Println("Error during transfer:", err)
	}

	// Print transaction histories
	printTransaction(acc1)
	printTransaction(acc2)
}
