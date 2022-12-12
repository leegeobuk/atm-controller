package bank

import (
	"github.com/leegeobuk/atm-controller/bank/account"
	"github.com/leegeobuk/atm-controller/typeutil"
)

// SimpleBank is a simple bank system
// that implements Bank interface.
type SimpleBank[T typeutil.Number] struct {
}

// NewSimple returns new SimpleBank
func NewSimple[T typeutil.Number]() *SimpleBank[T] {
	return &SimpleBank[T]{}
}

// VerifyCardNumber verifies if cardNumber is valid.
func (s *SimpleBank[T]) VerifyCardNumber(cardNumber string) bool {
	return len(cardNumber) == 16
}

// VerifyPIN verifies if PIN is valid.
func (s *SimpleBank[T]) VerifyPIN(pin string) bool {
	return len(pin) == 4
}

// GetBankAccount returns all bank accounts of user account.
func (s *SimpleBank[T]) GetBankAccount(cardNumber, pin string) account.BankAccount[T] {
	// TODO: implement retrieval from db
	return &account.SimpleCheckingAccount[T]{}
}

func (s *SimpleBank[T]) Balance(bankAccount account.BankAccount[T]) T {
	return bankAccount.Balance()
}

func (s *SimpleBank[T]) Deposit(bankAccount account.BankAccount[T], amount T) {
	bankAccount.Deposit(amount)
}

func (s *SimpleBank[T]) Withdraw(bankAccount account.BankAccount[T], amount T) error {
	return bankAccount.Withdraw(amount)
}
