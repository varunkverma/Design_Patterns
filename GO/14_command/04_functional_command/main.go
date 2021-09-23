package main

import "fmt"

type BankAccount struct {
	balance int
}

func Deposit(b *BankAccount, amount int) {
	b.balance += amount
	fmt.Printf("Deposited %d, current balance is %d\n", amount, b.balance)
}

func Withdraw(b *BankAccount, amount int) {
	if b.balance >= amount {
		b.balance -= amount
		fmt.Printf("Withdrew %d, current balance is %d\n", amount, b.balance)
	}
}

func main() {
	ba := &BankAccount{0}

	var commands []func()

	commands = append(commands, func() {
		Deposit(ba, 100)
	})
	commands = append(commands, func() {
		Withdraw(ba, 25)
	})

	for _, cmd := range commands {
		cmd()
	}
}
