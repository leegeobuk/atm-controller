package db

import (
	"github.com/leegeobuk/atm-controller/bank/account"
	"github.com/leegeobuk/atm-controller/typeutil"
)

// DB interface to communicate with Bank.
type DB[T typeutil.Number] interface {
	GetAccount(cardNumber string) account.BankAccount[T]
	UpdateAccount(bankAccount account.BankAccount[T])
}
