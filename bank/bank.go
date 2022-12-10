package bank

// Bank processes requests from ATM
type Bank interface {
	VerifyCardNumber(cardNumber string) bool
	VerifyPIN(pin string) bool
}
