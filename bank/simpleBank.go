package bank

import (
	"errors"

	"github.com/leegeobuk/atm-controller/bank/account"
	_card "github.com/leegeobuk/atm-controller/bank/card"
	"github.com/leegeobuk/atm-controller/db"
	"github.com/leegeobuk/atm-controller/typeutil"
)

var (
	ErrCardNumber = errors.New("card number doesn't exist")
	ErrPIN        = errors.New("PIN doesn't match")
)

// SimpleBank is a simple bank system
// that implements Bank interface.
type SimpleBank[T typeutil.Number] struct {
	db db.DB[T]
}

// NewSimple returns new SimpleBank
func NewSimple[T typeutil.Number](database db.DB[T]) *SimpleBank[T] {
	return &SimpleBank[T]{db: database}
}

// ValidateCardNumber validates cardNumber.
func (sb *SimpleBank[T]) ValidateCardNumber(cardNumber string) bool {
	return len(cardNumber) == 16
}

// ValidatePIN validates PIN.
func (sb *SimpleBank[T]) ValidatePIN(pin string) bool {
	return len(pin) == 4
}

// VerifyCard verifies card.
func (sb *SimpleBank[T]) VerifyCard(cardNumber, pin string) (*_card.Card[T], error) {
	card, ok := sb.db.GetCard(cardNumber)
	if !ok {
		return nil, ErrCardNumber
	}
	if card.PIN() != pin {
		return nil, ErrPIN
	}

	return card, nil
}

// GetBankAccount returns bank account linked to cardNumber.
func (sb *SimpleBank[T]) GetBankAccount(cardNumber string) (account.BankAccount[T], bool) {
	return sb.db.GetAccount(cardNumber)
}

func (sb *SimpleBank[T]) Balance(bankAccount account.BankAccount[T]) T {
	return bankAccount.Balance()
}

func (sb *SimpleBank[T]) Deposit(bankAccount account.BankAccount[T], amount T) {
	bankAccount.Deposit(amount)
	sb.db.UpdateAccount(bankAccount)
}

func (sb *SimpleBank[T]) Withdraw(bankAccount account.BankAccount[T], amount T) error {
	if err := bankAccount.Withdraw(amount); err != nil {
		return err
	}

	sb.db.UpdateAccount(bankAccount)
	return nil
}
