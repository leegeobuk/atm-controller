package card

import (
	"github.com/leegeobuk/atm-controller/bank/account"
	"github.com/leegeobuk/atm-controller/typeutil"
)

// Card gets verified and returns linked bank account.
type Card[T typeutil.Number] struct {
	cardNumber, pin string
	bankAccount     account.BankAccount[T]
}

// New returns new card.
func New[T typeutil.Number](cardNumber, pin string, bankAccount account.BankAccount[T]) *Card[T] {
	return &Card[T]{
		cardNumber:  cardNumber,
		pin:         pin,
		bankAccount: bankAccount,
	}
}

// CardNumber returns card number.
func (c *Card[T]) CardNumber() string {
	return c.cardNumber
}

// PIN returns PIN of card.
func (c *Card[T]) PIN() string {
	return c.pin
}

// BankAccount returns bank account linked to card.
func (c *Card[T]) BankAccount() account.BankAccount[T] {
	return c.bankAccount
}
