package account

// SimpleCheckingAccount is a checking account
// that simply notates balance in integer.
type SimpleCheckingAccount struct {
	balance int
}

func (s *SimpleCheckingAccount) Balance() int {
	//TODO implement me
	panic("implement me")
}

func (s *SimpleCheckingAccount) Deposit(amount int) int {
	//TODO implement me
	panic("implement me")
}

func (s *SimpleCheckingAccount) Withdraw(amount int) (int, error) {
	//TODO implement me
	panic("implement me")
}

// Name returns name of the account.
func (s *SimpleCheckingAccount) Name() string {
	return "Checking account"
}
