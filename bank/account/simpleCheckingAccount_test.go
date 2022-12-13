package account

import (
	"errors"
	"testing"

	"github.com/leegeobuk/atm-controller/typeutil"
)

func TestSimpleCheckingAccount_Balance(t *testing.T) {
	// given
	testAccount := NewSimpleChecking[int](1000)

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

func TestSimpleCheckingAccount_Deposit(t *testing.T) {
	// given
	testAccount := NewSimpleChecking[int](1000)

	type tc[T typeutil.Number] struct {
		name   string
		amount T
		want   T
	}

	tests := []tc[int]{
		{
			name:   "success case",
			amount: 100,
			want:   1100,
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

func TestSimpleCheckingAccount_Withdraw(t *testing.T) {
	// given
	testAccount := NewSimpleChecking[int](1000)

	type tc[T typeutil.Number] struct {
		name    string
		amount  T
		want    T
		wantErr error
	}

	tests := []tc[int]{
		{
			name:    "fail case: amount > balance",
			amount:  10000,
			want:    1000,
			wantErr: ErrWithdrawAmount,
		},
		{
			name:    "success case: amount <= balance",
			amount:  100,
			want:    900,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := testAccount.Withdraw(tt.amount); !errors.Is(err, tt.wantErr) {
				t.Errorf("Withdraw() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got := testAccount.Balance(); got != tt.want {
				t.Errorf("Withdraw() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimpleCheckingAccount_Limit(t *testing.T) {
	// given
	testAccount := NewSimpleChecking[int](1000)

	tests := []struct {
		name string
		want int
	}{
		{
			name: "success case",
			want: -1,
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

func TestSimpleCheckingAccount_Type(t *testing.T) {
	// given
	testAccount := NewSimpleChecking[int](1000)

	tests := []struct {
		name string
		want string
	}{
		{
			name: "success case",
			want: "Checking account",
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
