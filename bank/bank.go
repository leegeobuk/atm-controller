package bank

import (
	"github.com/leegeobuk/atm-controller/bank/account"
	"github.com/leegeobuk/atm-controller/bank/card"
	"github.com/leegeobuk/atm-controller/typeutil"
)

// Bank processes requests from ATM
type Bank[T typeutil.Number] interface {
	ValidateCardNumber(cardNumber string) bool
	ValidatePIN(pin string) bool
	VerifyCard(cardNumber, pin string) (*card.Card[T], error)
	GetBankAccount(cardNumber string) (account.BankAccount[T], bool)
	Balance(bankAccount account.BankAccount[T]) T
	Deposit(bankAccount account.BankAccount[T], amount T)
	Withdraw(bankAccount account.BankAccount[T], amount T) error
}
