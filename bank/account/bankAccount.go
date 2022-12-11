package account

// BankAccount represents bank account of a bank
type BankAccount interface {
	Balance() int
	Deposit(amount int) int
	Withdraw(amount int) (int, error)
	Name() string
}
