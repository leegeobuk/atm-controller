package account

import "github.com/leegeobuk/atm-controller/typeutil"

// BankAccount represents bank account of a bank
type BankAccount[T typeutil.Number] interface {
	Balance() T
	Deposit(amount T)
	Withdraw(amount T) error
	Limit() int
	Type() string
}
