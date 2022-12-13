package cashbin

import (
	"errors"
	"testing"

	"github.com/leegeobuk/atm-controller/bank/account"
	"github.com/leegeobuk/atm-controller/typeutil"
)

func TestSimpleCashBin_Cash(t *testing.T) {
	type tc[T typeutil.Number] struct {
		name    string
		cashBin CashBin[T]
		want    T
	}
	tests := []tc[int]{
		{
			name:    "success case",
			cashBin: NewSimple[int](1_000_000),
			want:    1_000_000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cashBin.Cash(); got != tt.want {
				t.Errorf("Cash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimpleCashBin_Deposit(t *testing.T) {
	type tc[T typeutil.Number] struct {
		name    string
		cashBin CashBin[T]
		amount  T
		want    T
	}
	tests := []tc[int]{
		{
			name:    "success case",
			cashBin: NewSimple[int](1_000_000),
			amount:  100,
			want:    1_000_100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cashBin.Deposit(tt.amount)
			if got := tt.cashBin.Cash(); got != tt.want {
				t.Errorf("Deposit() = %v, %v", got, tt.want)
			}
		})
	}
}

func TestSimpleCashBin_Withdraw(t *testing.T) {
	type tc[T typeutil.Number] struct {
		name    string
		cashBin CashBin[T]
		amount  T
		want    T
		wantErr error
	}
	tests := []tc[int]{
		{
			name:    "fail case: amount > cash",
			cashBin: NewSimple[int](1_000_000),
			amount:  1_000_000_000,
			want:    1_000_000,
			wantErr: account.ErrWithdrawAmount,
		},
		{
			name:    "success case: amount <= cash",
			cashBin: NewSimple[int](1_000_000),
			amount:  1_000,
			want:    999_000,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.cashBin.Withdraw(tt.amount); !errors.Is(err, tt.wantErr) {
				t.Errorf("Withdraw() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got := tt.cashBin.Cash(); got != tt.want {
				t.Errorf("Withdraw() = %v, want %v", got, tt.want)
			}
		})
	}
}
