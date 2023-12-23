package pointers

import (
	"errors"
	"fmt"
)

// Bitconin represents a number of Bitcoins.
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet stores the number of Bitcoin somene owns.
type Wallet struct {
	balance Bitcoin
}

// Deposit will add some Bitcoin to a wallet.
func (w *Wallet) Deposit(amaunt Bitcoin) {
	w.balance += amaunt
}

// Balance returns the number of Bitcoin a wallet has.
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// ErrInsufficientFunds means a wallet does not have enough Bitcoin to perform a withdraw.
var ErrInsufficientFunds = errors.New("cannnot withdraw, insufficient funds")

// Withdraw substracts some Bitcoin from the wallet, returning an error if it cannnot be performed.
func (w *Wallet) Withdraw(amaunt Bitcoin) error {
	if amaunt > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amaunt
	return nil
}
