package account

import (
	"errors"
	"testing"

	"github.com/leegeobuk/atm-controller/typeutil"
)

func TestSimpleSavingsAccount_Balance(t *testing.T) {
	// given
	testAccount := NewSimpleSavings[int](1000)

	type tc[T typeutil.Number] struct {
		name string
		want T
	}

	tests := []tc[int]{
		{
			name: "success case",
			want: 1000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testAccount.Balance(); got != tt.want {
				t.Errorf("Balance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimpleSavingsAccount_Deposit(t *testing.T) {
	// given
	testAccount := NewSimpleSavings[int](1000)

	type tc[T typeutil.Number] struct {
		name   string
		amount T
		want   T
	}

	tests := []tc[int]{
		{
			name:   "success case",
			amount: 100,
			want:   1105,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testAccount.Deposit(tt.amount)
			if got := testAccount.Balance(); got != tt.want {
				t.Errorf("Deposit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimpleSavingsAccount_Withdraw(t *testing.T) {
	type tc[T typeutil.Number] struct {
		name       string
		account    BankAccount[T]
		amount     T
		wantAmount T
		wantLimit  int
		wantErr    error
	}

	tests := []tc[int]{
		{
			name:       "fail case: amount > balance",
			account:    NewSimpleSavings[int](1000),
			amount:     10000,
			wantAmount: 1000,
			wantLimit:  savingsWithdrawalLimit,
			wantErr:    ErrWithdrawAmount,
		},
		{
			name:       "fail case: limit reached",
			account:    &SimpleSavingsAccount[int]{balance: 1000, limit: 0},
			amount:     100,
			wantAmount: 1000,
			wantLimit:  0,
			wantErr:    ErrWithdrawLimit,
		},
		{
			name:       "success case: amount <= balance",
			account:    NewSimpleSavings[int](1000),
			amount:     100,
			wantAmount: 900,
			wantLimit:  savingsWithdrawalLimit - 1,
			wantErr:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.account.Withdraw(tt.amount); !errors.Is(err, tt.wantErr) {
				t.Errorf("Withdraw() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got := tt.account.Balance(); got != tt.wantAmount {
				t.Errorf("Withdraw() = %v, wantAmount %v", got, tt.wantAmount)
			}
			if limit := tt.account.Limit(); limit != tt.wantLimit {
				t.Errorf("Withdraw() = %v, want %v", limit, tt.wantLimit)
			}
		})
	}
}

func TestSimpleSavingsAccount_Limit(t *testing.T) {
	// given
	testAccount := NewSimpleSavings[int](1000)

	tests := []struct {
		name string
		want int
	}{
		{
			name: "success case",
			want: savingsWithdrawalLimit,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testAccount.Limit(); got != tt.want {
				t.Errorf("Limit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimpleSavingsAccount_Type(t *testing.T) {
	// given
	testAccount := NewSimpleSavings[int](1000)

	tests := []struct {
		name string
		want string
	}{
		{
			name: "success case",
			want: "Savings account",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testAccount.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}
