package wallet

import (
	"errors"
	"fmt"
)

type Stringer interface {
	String() string
}

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC,", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
	// In Go, when you call a function or a method the arguments are copied.
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
	// In Go, when you call a function or a method the arguments are copied.
}

func (w *Wallet) Balance() Bitcoin {
	//fmt.Printf("address of balance in test is %p \n", &w.balance)
	return w.balance
}
