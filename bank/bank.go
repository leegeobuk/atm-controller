package bank

import "github.com/leegeobuk/atm-controller/bank/account"

// Bank processes requests from ATM
type Bank interface {
	VerifyCardNumber(cardNumber string) bool
	VerifyPIN(pin string) bool
	GetBankAccounts(cardNumber, pin string) []account.BankAccount
}
