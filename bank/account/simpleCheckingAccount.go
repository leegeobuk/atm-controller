package account

import (
	"errors"

	"github.com/leegeobuk/atm-controller/typeutil"
)

// SimpleCheckingAccount is a checking account
// that simply notates balance in integer.
type SimpleCheckingAccount[T typeutil.Number] struct {
	balance T
}

// Balance returns balance left.
func (s *SimpleCheckingAccount[T]) Balance() T {
	return s.balance
}

// Deposit adds amount to balance and returns updated balance.
func (s *SimpleCheckingAccount[T]) Deposit(amount T) {
	s.balance += amount
}

// Withdraw withdraws money from balance and returns updated balance.
// error is returned if amount > balance.
func (s *SimpleCheckingAccount[T]) Withdraw(amount T) error {
	if amount > s.balance {
		return errors.New("cannot withdraw more than balance")
	}

	s.balance -= amount
	return nil
}

// Name returns type of the bank account
func (s *SimpleCheckingAccount[T]) Name() string {
	return "Checking account"
}
