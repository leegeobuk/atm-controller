package account

import (
	"errors"

	"github.com/leegeobuk/atm-controller/typeutil"
)

var (
	errWithdrawAmount = errors.New("withdraw larger than balance")
	errWithdrawLimit  = errors.New("withdraw limit reached")
)

// SimpleCheckingAccount is a checking account
// that implements BankAccount.
type SimpleCheckingAccount[T typeutil.Number] struct {
	balance T
}

// NewSimpleChecking return SimpleCheckingAccount wit given balance.
func NewSimpleChecking[T typeutil.Number](balance T) *SimpleCheckingAccount[T] {
	return &SimpleCheckingAccount[T]{
		balance: balance,
	}
}

// Balance returns current balance.
func (acc *SimpleCheckingAccount[T]) Balance() T {
	return acc.balance
}

// Deposit adds amount to balance.
func (acc *SimpleCheckingAccount[T]) Deposit(amount T) {
	acc.balance += amount
}

// Withdraw withdraws money from balance and returns updated balance.
// error is returned if amount > balance.
func (acc *SimpleCheckingAccount[T]) Withdraw(amount T) error {
	if amount > acc.balance {
		return errWithdrawAmount
	}

	acc.balance -= amount
	return nil
}

// Type returns type of the bank account
func (acc *SimpleCheckingAccount[T]) Type() string {
	return "Checking account"
}
