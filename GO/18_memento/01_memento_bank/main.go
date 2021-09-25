package main

import "fmt"

type Memento struct {
	Balance int
}

type BankAccount struct {
	balance int
}

func NewBankAccount(amount int) (*BankAccount, *Memento) {
	return &BankAccount{
			balance: amount,
		},
		&Memento{
			Balance: amount,
		}
}

func (ba *BankAccount) Deposit(amount int) *Memento {
	ba.balance += amount
	return &Memento{
		Balance: ba.balance,
	}
}

func (ba *BankAccount) Restore(m *Memento) {
	ba.balance = m.Balance
}

func main() {
	ba, m0 := NewBankAccount(100)

	m1 := ba.Deposit(50)
	m2 := ba.Deposit(20)
	fmt.Println(ba)

	ba.Restore(m1)
	fmt.Println(ba)

	ba.Restore(m2)
	fmt.Println(ba)

	ba.Restore(m0)
	fmt.Println(ba)
}
