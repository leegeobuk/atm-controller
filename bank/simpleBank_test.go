package bank

import (
	"testing"

	"github.com/leegeobuk/atm-controller/db"
)

func TestSimpleBank_VerifyCardNumber(t *testing.T) {
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

func TestSimpleBank_VerifyPIN(t *testing.T) {
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
