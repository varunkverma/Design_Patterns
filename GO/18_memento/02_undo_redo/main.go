package main

import "fmt"

type Memento struct {
	Balance int
}

type BankAccount struct {
	balance int
	changes []*Memento
	current int // current Memento's index
}

func (ba *BankAccount) String() string {
	return fmt.Sprintf("Balance = $%d, current = %d", ba.balance, ba.current)
}

func (ba *BankAccount) Deposit(amount int) *Memento {
	ba.balance += amount

	m := &Memento{
		Balance: ba.balance,
	}

	ba.changes = append(ba.changes, m)
	ba.current++

	fmt.Printf("Deposited: $%d, current balance: $%d\n", amount, ba.balance)

	return m
}

func (ba *BankAccount) Restore(m *Memento) {
	if m != nil {
		ba.balance = m.Balance
		ba.changes = append(ba.changes, m)
		ba.current = len(ba.changes) - 1
	}
}

func (ba *BankAccount) Undo() *Memento {
	if ba.current > 0 { // if you are at the initial state
		ba.current--
		m := ba.changes[ba.current]
		ba.balance = m.Balance
		return m
	}
	return nil
}

func (ba *BankAccount) Redo() *Memento {
	if ba.current+1 < len(ba.changes) { // if you are at the latest state
		ba.current++
		m := ba.changes[ba.current]
		ba.balance = m.Balance
		return m
	}
	return nil
}

func NewBankAccount(balance int) *BankAccount {
	ba := &BankAccount{
		balance: balance,
		changes: []*Memento{},
	}
	ba.changes = append(ba.changes, &Memento{Balance: balance})
	return ba
}

func main() {
	ba := NewBankAccount(100)

	fmt.Println(ba.String())

	ba.Deposit(50)
	ba.Deposit(25)

	fmt.Println(ba.String())

	ba.Undo()
	fmt.Println("Undo 1,", ba.String())
	ba.Undo()
	fmt.Println("Undo 2,", ba.String())
	ba.Redo()
	fmt.Println("Redo 1,", ba.String())

}
