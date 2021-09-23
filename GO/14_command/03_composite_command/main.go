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
	Succeeded() bool
	SetSucceeded(value bool)
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

func (bac *BankAccountCommand) Succeeded() bool {
	return bac.succeeded
}

func (bac *BankAccountCommand) SetSucceeded(value bool) {
	bac.succeeded = value
}

type CompositeBankAccountCommand struct {
	commands []Command
}

func (cbac *CompositeBankAccountCommand) Call() {
	for _, cmd := range cbac.commands {
		(cmd).Call()
	}
}

func (cbac *CompositeBankAccountCommand) Undo() {
	for i := range cbac.commands {
		cmd := cbac.commands[len(cbac.commands)-i-1]
		(cmd).Undo()
	}
}

func (cbac *CompositeBankAccountCommand) Succeeded() bool {
	for _, cmd := range cbac.commands {
		if !(cmd).Succeeded() {
			return false
		}
	}
	return true
}

func (cbac *CompositeBankAccountCommand) SetSucceeded(value bool) {
	for _, cmd := range cbac.commands {
		(cmd).SetSucceeded(value)
	}
}

type MoneyTransferCommand struct {
	CompositeBankAccountCommand
	from, to *BankAccount
	amount   int
}

func NewMoneyTransferCommand(from, to *BankAccount, amount int) *MoneyTransferCommand {
	mtc := &MoneyTransferCommand{
		from:   from,
		to:     to,
		amount: amount,
		CompositeBankAccountCommand: CompositeBankAccountCommand{
			commands: make([]Command, 0),
		},
	}
	mtc.commands = append(mtc.commands, NewBankAccountCommand(from, WITHDRAW, amount))
	mtc.commands = append(mtc.commands, NewBankAccountCommand(to, DEPOSIT, amount))
	return mtc
}

// ensures atomicity
func (mtc *MoneyTransferCommand) Call() {
	ok := true
	for _, cmd := range mtc.commands {
		if ok {
			(cmd).Call()
			ok = cmd.Succeeded()
		} else {
			cmd.SetSucceeded(false)
		}
	}
}

func main() {
	from := &BankAccount{
		balance: 100,
	}
	to := &BankAccount{
		balance: 0,
	}
	mtc := NewMoneyTransferCommand(from, to, 75)
	mtc.Call()
	fmt.Println(from, to)

	mtc.Undo()
	fmt.Println(from, to)
}
