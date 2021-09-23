package main

import "fmt"

var overDraftLimit = -500

type BankAccount struct {
	balance int
}

func (b *BankAccount) Deposit(amount int) {
	b.balance += amount
	fmt.Printf("Deposited %d, current balance is %d\n", amount, b.balance)
}

func (b *BankAccount) Withdraw(amount int) bool {
	if b.balance-amount >= overDraftLimit {
		b.balance -= amount
		fmt.Printf("Withdrew %d, current balance is %d\n", amount, b.balance)
		return true
	}
	return false
}

// Command interface
type Command interface {
	Call()
	Undo()
}

type Action int

const (
	DEPOSIT Action = iota
	WITHDRAW
)

// Command structure
type BankAccountCommand struct {
	account   *BankAccount // operations to be performed on
	action    Action       // what operation
	amount    int          // parameters
	succeeded bool         // whether the command succeeded or not
}

func NewBankAccountCommand(account *BankAccount, action Action, amount int) *BankAccountCommand {
	return &BankAccountCommand{
		account: account,
		action:  action,
		amount:  amount,
	}
}

func (bac *BankAccountCommand) Call() {
	switch bac.action {
	case DEPOSIT:
		bac.account.Deposit(bac.amount)
		bac.succeeded = true
	case WITHDRAW:
		bac.succeeded = bac.account.Withdraw(bac.amount)
	}
}

func (bac *BankAccountCommand) Undo() {
	if !bac.succeeded {
		return
	}
	switch bac.action {
	case DEPOSIT:
		bac.account.Withdraw(bac.amount)
	case WITHDRAW:
		bac.account.Deposit(bac.amount)
	}
}

func main() {
	ba := &BankAccount{}
	cmd := NewBankAccountCommand(ba, DEPOSIT, 100)
	cmd.Call()
	cmd2 := NewBankAccountCommand(ba, WITHDRAW, 50)
	cmd2.Call()
	cmd2.Undo()
	fmt.Println(ba)
	cmd.Undo()
	fmt.Println(ba)
}
