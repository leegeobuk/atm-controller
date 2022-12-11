package account

import "github.com/leegeobuk/atm-controller/typeutil"

// UserAccount represents user account registered in bank system.
type UserAccount[T typeutil.Number] struct {
	cardNumber string
	pin        string
}

func NewUser[T typeutil.Number](cardNumber, pin string) *UserAccount[T] {
	return &UserAccount[T]{
		cardNumber: cardNumber,
		pin:        pin,
	}
}

// BankAccounts returns all bank accounts of the user account
func (a *UserAccount[T]) BankAccounts() []BankAccount[T] {
	return nil
}
