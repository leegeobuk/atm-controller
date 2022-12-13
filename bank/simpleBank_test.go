package bank

import (
	"errors"
	"reflect"
	"testing"

	"github.com/leegeobuk/atm-controller/bank/account"
	_card "github.com/leegeobuk/atm-controller/bank/card"
	"github.com/leegeobuk/atm-controller/db"
	"github.com/leegeobuk/atm-controller/typeutil"
)

func TestSimpleBank_ValidateCardNumber(t *testing.T) {
	// given
	testBank := NewSimple[int](db.NewSimple[int]())

	tests := []struct {
		name       string
		cardNumber string
		want       bool
	}{
		{
			name:       "card number length=1",
			cardNumber: "1",
			want:       false,
		},
		{
			name:       "card number length=2",
			cardNumber: "12",
			want:       false,
		},
		{
			name:       "card number length=3",
			cardNumber: "123",
			want:       false,
		},
		{
			name:       "card number length=4",
			cardNumber: "1234",
			want:       false,
		},
		{
			name:       "card number length=5",
			cardNumber: "12345",
			want:       false,
		},
		{
			name:       "card number length=6",
			cardNumber: "123456",
			want:       false,
		},
		{
			name:       "card number length=7",
			cardNumber: "1234567",
			want:       false,
		},
		{
			name:       "card number length=8",
			cardNumber: "12345678",
			want:       false,
		},
		{
			name:       "card number length=9",
			cardNumber: "123456789",
			want:       false,
		},
		{
			name:       "card number length=10",
			cardNumber: "1234123412",
			want:       false,
		},
		{
			name:       "card number length=11",
			cardNumber: "12341234123",
			want:       false,
		},
		{
			name:       "card number length=12",
			cardNumber: "123412341234",
			want:       false,
		},
		{
			name:       "card number length=13",
			cardNumber: "1234123412341",
			want:       false,
		},
		{
			name:       "card number length=14",
			cardNumber: "12341234123412",
			want:       false,
		},
		{
			name:       "card number length=15",
			cardNumber: "123412341234123",
			want:       false,
		},
		{
			name:       "card number length=16",
			cardNumber: "1234123412341234",
			want:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testBank.ValidateCardNumber(tt.cardNumber); got != tt.want {
				t.Errorf("ValidateCardNumber(%s) = %v, want %v", tt.cardNumber, got, tt.want)
			}
		})
	}
}

func TestSimpleBank_ValidatePIN(t *testing.T) {
	// given
	testBank := NewSimple[int](db.NewSimple[int]())

	tests := []struct {
		name string
		pin  string
		want bool
	}{
		{
			name: "PIN length=1",
			pin:  "1",
			want: false,
		},
		{
			name: "PIN length=2",
			pin:  "12",
			want: false,
		},
		{
			name: "PIN length=3",
			pin:  "123",
			want: false,
		},
		{
			name: "PIN length=4",
			pin:  "1234",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testBank.ValidatePIN(tt.pin); got != tt.want {
				t.Errorf("ValidatePIN(%s) = %v, want %v", tt.pin, got, tt.want)
			}
		})
	}
}

func TestSimpleBank_VerifyCard(t *testing.T) {
	// given
	testBank := setup[int]()

	type tc[T typeutil.Number] struct {
		name       string
		cardNumber string
		pin        string
		want       *_card.Card[T]
		wantErr    error
	}

	tests := []tc[int]{
		{
			name:       "fail case: unregistered card number",
			cardNumber: "1234567890123456",
			pin:        "0000",
			want:       nil,
			wantErr:    ErrCardNumber,
		},
		{
			name:       "fail case: wrong PIN",
			cardNumber: "1234123412341234",
			pin:        "0000",
			want:       nil,
			wantErr:    ErrPIN,
		},
		{
			name:       "success case: registered SimpleCheckingAccount",
			cardNumber: "1234123412341234",
			pin:        "1234",
			want:       _card.New[int]("1234123412341234", "1234", account.NewSimpleChecking[int](12_341_234)),
			wantErr:    nil,
		},
		{
			name:       "success case: registered SimpleSavingsAccount",
			cardNumber: "4321432143214321",
			pin:        "4321",
			want:       _card.New[int]("4321432143214321", "4321", account.NewSimpleSavings[int](43_214_321)),
			wantErr:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := testBank.VerifyCard(tt.cardNumber, tt.pin)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("VerifyCard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VerifyCard() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimpleBank_GetBankAccount(t *testing.T) {
	// given
	testBank := setup[int]()

	type tc[T typeutil.Number] struct {
		name       string
		cardNumber string
		want       account.BankAccount[T]
		wantErr    error
	}

	tests := []tc[int]{
		{
			name:       "fail case: unregistered bank account",
			cardNumber: "1234567890123456",
			want:       nil,
			wantErr:    ErrCardNumber,
		},
		{
			name:       "success case: registered SimpleCheckingAccount",
			cardNumber: "1234123412341234",
			want:       account.NewSimpleChecking[int](12_341_234),
			wantErr:    nil,
		},
		{
			name:       "success case: registered SimpleSavingsAccount",
			cardNumber: "1000100010001000",
			want:       account.NewSimpleSavings[int](10_001_000),
			wantErr:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := testBank.GetBankAccount(tt.cardNumber)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("GetBankAccount() err = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBankAccount() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimpleBank_Balance(t *testing.T) {
	// given
	testBank := setup[int]()

	type tc[T typeutil.Number] struct {
		name        string
		bankAccount account.BankAccount[T]
		want        T
	}

	tests := []tc[int]{
		{
			name:        "success case: right balance",
			bankAccount: account.NewSimpleChecking(12_341_234),
			want:        12_341_234,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testBank.Balance(tt.bankAccount); got != tt.want {
				t.Errorf("Balance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimpleBank_Deposit(t *testing.T) {
	// given
	testBank := setup[int]()

	type tc[T typeutil.Number] struct {
		name        string
		bankAccount account.BankAccount[T]
		amount      T
		want        T
	}

	tests := []tc[int]{
		{
			name:        "success case: right amount added for SimpleCheckingAccount",
			bankAccount: account.NewSimpleChecking[int](1_000_000),
			amount:      100,
			want:        1_000_100,
		},
		{
			name:        "success case: right amount added for SimpleSavingsAccount",
			bankAccount: account.NewSimpleSavings[int](1_000_000),
			amount:      100,
			want:        1_000_105,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testBank.Deposit(tt.bankAccount, tt.amount)
			if tt.bankAccount.Balance() != tt.want {
				t.Errorf("Balance() = %v, want %v", tt.bankAccount.Balance(), tt.want)
			}
		})
	}
}

func TestSimpleBank_Withdraw(t *testing.T) {
	// given
	testBank := setup[int]()

	type tc[T typeutil.Number] struct {
		name        string
		bankAccount account.BankAccount[T]
		amount      T
		wantErr     error
	}

	tests := []tc[int]{
		{
			name:        "fail case: amount > balance",
			bankAccount: account.NewSimpleChecking[int](100),
			amount:      200,
			wantErr:     account.ErrWithdrawAmount,
		},
		{
			name:        "fail case: withdrawal limit reached",
			bankAccount: account.NewSimpleSavings[int](100),
			amount:      2,
			wantErr:     account.ErrWithdrawLimit,
		},
		{
			name:        "fail case: amount > balance while limit > 0",
			bankAccount: account.NewSimpleSavings[int](100),
			amount:      20,
			wantErr:     account.ErrWithdrawAmount,
		},
		{
			name:        "success case: amount <= balance",
			bankAccount: account.NewSimpleChecking[int](100),
			amount:      1,
			wantErr:     nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			limit := tt.bankAccount.Limit()
			for i := 0; i < limit; i++ {
				testBank.Withdraw(tt.bankAccount, tt.amount)
			}

			if err := testBank.Withdraw(tt.bankAccount, tt.amount); !errors.Is(err, tt.wantErr) {
				t.Errorf("Withdraw() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func setup[T typeutil.Number]() (testBank *SimpleBank[T]) {
	testDB := db.NewSimple[T]()
	return NewSimple[T](testDB)
}
