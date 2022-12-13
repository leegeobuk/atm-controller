package db

import (
	"reflect"
	"testing"

	"github.com/leegeobuk/atm-controller/bank/account"
	_card "github.com/leegeobuk/atm-controller/bank/card"
	"github.com/leegeobuk/atm-controller/typeutil"
)

func TestSimpleDB_GetCard(t *testing.T) {
	// given
	testDB := NewSimple[int]()

	type tc[T typeutil.Number] struct {
		name      string
		carNumber string
		want      *_card.Card[T]
		ok        bool
	}

	tests := []tc[int]{
		{
			name:      "fail case: unregistered card",
			carNumber: "1234567890123456",
			want:      nil,
			ok:        false,
		},
		{
			name:      "success case: registered card",
			carNumber: "1234123412341234",
			want:      _card.New[int]("1234123412341234", "1234", account.NewSimpleChecking[int](12_341_234)),
			ok:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := testDB.GetCard(tt.carNumber)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCard() got = %v, want %v", got, tt.want)
			}
			if ok != tt.ok {
				t.Errorf("GetCard() ok = %v, want %v", ok, tt.ok)
			}
		})
	}
}

func TestSimpleDB_GetAccount(t *testing.T) {
	// given
	testDB := NewSimple[int]()

	type tc[T typeutil.Number] struct {
		name       string
		cardNumber string
		want       account.BankAccount[T]
		ok         bool
	}

	tests := []tc[int]{
		{
			name:       "fail case: unregistered card number",
			cardNumber: "1234567890123456",
			want:       nil,
			ok:         false,
		},
		{
			name:       "success case: registered card number",
			cardNumber: "1234123412341234",
			want:       account.NewSimpleChecking[int](12_341_234),
			ok:         true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := testDB.GetAccount(tt.cardNumber)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAccount() got = %v, want %v", got, tt.want)
			}
			if ok != tt.ok {
				t.Errorf("GetAccount() ok = %v, want %v", ok, tt.ok)
			}
		})
	}
}
