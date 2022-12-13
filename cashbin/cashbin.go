package cashbin

import "github.com/leegeobuk/atm-controller/typeutil"

// CashBin is a cash pool of an ATM
type CashBin[T typeutil.Number] interface {
	Cash() T
	Deposit(amount T)
	Withdraw(amount T) error
}
