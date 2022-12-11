package account

// UserAccount represents user account registered in bank system.
type UserAccount struct {
	cardNumber string
	pin        string
}

func NewUser(cardNumber, pin string) *UserAccount {
	return &UserAccount{
		cardNumber: cardNumber,
		pin:        pin,
	}
}

// BankAccounts returns all bank accounts of the user account
func (a *UserAccount) BankAccounts() []BankAccount {
	return nil
}
