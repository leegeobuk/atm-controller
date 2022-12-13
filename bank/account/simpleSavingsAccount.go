package account

import "github.com/leegeobuk/atm-controller/typeutil"

const savingsWithdrawLimit = 10

// SimpleSavingsAccount is a savings account
// that implements BankAccount.
type SimpleSavingsAccount[T typeutil.Number] struct {
	balance T
	limit   int
}

// NewSimpleSavings returns SimpleSavingsAccount with given balance.
func NewSimpleSavings[T typeutil.Number](balance T) *SimpleSavingsAccount[T] {
	return &SimpleSavingsAccount[T]{balance: balance, limit: savingsWithdrawLimit}
}

// Balance returns current balance.
func (acc *SimpleSavingsAccount[T]) Balance() T {
	return acc.balance
}

// Deposit adds 105% of the amount to balance.
func (acc *SimpleSavingsAccount[T]) Deposit(amount T) {
	amount += amount / 20
	acc.balance += amount
}

// Withdraw withdraws amount from balance and
// decrease limit by 1.
// If amount > balance, balance is not withdrawn.
// If limit reaches 0, balance cannot be withdrawn anymore.
func (acc *SimpleSavingsAccount[T]) Withdraw(amount T) error {
	if amount > acc.balance {
		return ErrWithdrawAmount
	}
	if acc.limit == 0 {
		return ErrWithdrawLimit
	}

	acc.balance -= amount
	acc.limit--

	return nil
}

// Type returns type of the bank account.
func (acc *SimpleSavingsAccount[T]) Type() string {
	return "Savings account"
}
