package db

import (
	"github.com/leegeobuk/atm-controller/bank/account"
	"github.com/leegeobuk/atm-controller/bank/card"
	"github.com/leegeobuk/atm-controller/typeutil"
)

// DB interface to communicate with Bank.
type DB[T typeutil.Number] interface {
	GetCard(carNumber string) (*card.Card[T], bool)
	GetAccount(cardNumber string) (account.BankAccount[T], bool)
	UpdateAccount(bankAccount account.BankAccount[T])
}
