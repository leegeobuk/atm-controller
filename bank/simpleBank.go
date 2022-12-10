package bank

// SimpleBank is a simple bank system
// that implements Bank interface.
type SimpleBank struct {
}

// NewSimple returns new SimpleBank
func NewSimple() *SimpleBank {
	return &SimpleBank{}
}
