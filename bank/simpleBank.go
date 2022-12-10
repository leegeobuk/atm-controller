package bank

// SimpleBank is a simple bank system
// that implements Bank interface.
type SimpleBank struct {
}

// NewSimple returns new SimpleBank
func NewSimple() *SimpleBank {
	return &SimpleBank{}
}

// VerifyCardNumber returns whether cardNumber is valid or not.
func (s *SimpleBank) VerifyCardNumber(cardNumber string) bool {
	return len(cardNumber) == 16
}
