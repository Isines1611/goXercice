package main

import (
	"fmt"
)

// TODO: We want to create a simple banking system
// Each account has a name, balance and a transaction history and can deposit, withdraw and transfer money.
// For the whole exercice, when asked to print balance, only print 2 decimals (example: 2.15779 => 2.15)
// What is asked is:
// - Create the banking interface
// - Create the account struct
// - Create a function to create accounts
//		->  you should print "Account create with: $INITIALE-BALANCE" if everything went well, otherwise nothing is printed
// - Implemenent the 'deposit', 'withdraw' and 'transfer' such as:
// 		* 'deposit': if negative amount, print: "Deposit should be greater than zero", otherwise add the amount and add this to the transaction: "Deposite of: $BALANCE"
// 		* 'withdraw': if negative amount, print "Withdraw should be greater than zero",
// 			if not enough money, return error: "Error funds", if all went well then add to the transaction: "Withdrew: $BALANCE"
// 		* 'transfer': if not enough money, return error: "Error funds".
// 			if all went well then add to the transaction to the corresponding accounts: "Transferred: $BALANCE to ACCOUNT-NAME", "Received: $BALANCE from ACCOUNT-NAME"
// - Finally, implement 'printTransaction'. If no account is given then print "Unknown account type".
// 		Otherwise, print the transactions as: "\nTransaction history for ACCOUNT-NAME:\n" and each transaction

// GOOD LUCK !

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
