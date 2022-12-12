package bank

import (
	"github.com/leegeobuk/atm-controller/bank/account"
	"github.com/leegeobuk/atm-controller/typeutil"
)

// Bank processes requests from ATM
type Bank[T typeutil.Number] interface {
	VerifyCardNumber(cardNumber string) bool
	VerifyPIN(pin string) bool
	GetBankAccount(cardNumber, pin string) account.BankAccount[T]
	Balance(bankAccount account.BankAccount[T]) T
	Deposit(bankAccount account.BankAccount[T], amount T)
	Withdraw(bankAccount account.BankAccount[T], amount T) error
}
