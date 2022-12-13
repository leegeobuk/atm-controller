package bank

import (
	"github.com/leegeobuk/atm-controller/bank/account"
	"github.com/leegeobuk/atm-controller/db"
	"github.com/leegeobuk/atm-controller/typeutil"
)

// SimpleBank is a simple bank system
// that implements Bank interface.
type SimpleBank[T typeutil.Number] struct {
	database db.DB[T]
}

// NewSimple returns new SimpleBank
func NewSimple[T typeutil.Number](database db.DB[T]) *SimpleBank[T] {
	return &SimpleBank[T]{database: database}
}

// VerifyCardNumber verifies if cardNumber is valid.
func (sb *SimpleBank[T]) VerifyCardNumber(cardNumber string) bool {
	return len(cardNumber) == 16
}

// VerifyPIN verifies if PIN is valid.
func (sb *SimpleBank[T]) VerifyPIN(pin string) bool {
	return len(pin) == 4
}

// GetBankAccount returns bank account connected to cardNumber.
func (sb *SimpleBank[T]) GetBankAccount(cardNumber, pin string) account.BankAccount[T] {
	return sb.database.GetAccount(cardNumber)
}

func (sb *SimpleBank[T]) Balance(bankAccount account.BankAccount[T]) T {
	return bankAccount.Balance()
}

func (sb *SimpleBank[T]) Deposit(bankAccount account.BankAccount[T], amount T) {
	bankAccount.Deposit(amount)
}

func (sb *SimpleBank[T]) Withdraw(bankAccount account.BankAccount[T], amount T) error {
	return bankAccount.Withdraw(amount)
}
