package bank

import (
	"github.com/leegeobuk/atm-controller/bank/account"
)

// SimpleBank is a simple bank system
// that implements Bank interface.
type SimpleBank struct {
}

// NewSimple returns new SimpleBank
func NewSimple() *SimpleBank {
	return &SimpleBank{}
}

// VerifyCardNumber verifies if cardNumber is valid.
func (s *SimpleBank) VerifyCardNumber(cardNumber string) bool {
	return len(cardNumber) == 16
}

// VerifyPIN verifies if PIN is valid.
func (s *SimpleBank) VerifyPIN(pin string) bool {
	return len(pin) == 4
}

// GetBankAccounts returns all bank accounts of user account.
func (s *SimpleBank) GetBankAccounts(cardNumber, pin string) []account.BankAccount {
	// TODO: implement retrieval from db
	return []account.BankAccount{&account.SimpleCheckingAccount{}}
}
