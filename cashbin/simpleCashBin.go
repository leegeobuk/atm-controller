package cashbin

import (
	"github.com/leegeobuk/atm-controller/bank/account"
	"github.com/leegeobuk/atm-controller/typeutil"
)

// SimpleCashBin is a simple cash bin
// that implements CashBin.
type SimpleCashBin[T typeutil.Number] struct {
	cash T
}

// NewSimple returns new SimpleCashBin
func NewSimple[T typeutil.Number](cashPool T) *SimpleCashBin[T] {
	return &SimpleCashBin[T]{
		cash: cashPool,
	}
}

// Cash returns the amount of cash.
func (cb *SimpleCashBin[T]) Cash() T {
	return cb.cash
}

// Deposit deposits cash to the cash bin.
func (cb *SimpleCashBin[T]) Deposit(amount T) {
	cb.cash += amount
}

// Withdraw withdraws cash from the cash bin.
// Error is returned when amount > cash.
func (cb *SimpleCashBin[T]) Withdraw(amount T) error {
	if amount > cb.cash {
		return account.ErrWithdrawAmount
	}

	cb.cash -= amount
	return nil
}
