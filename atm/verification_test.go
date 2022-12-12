package atm

import (
	"io"
	"strings"
	"testing"

	"github.com/leegeobuk/atm-controller/bank"
	"github.com/leegeobuk/atm-controller/cashbin"
)

func TestATM_verifyCardNumber(t *testing.T) {
	//given
	newBank, cashBin := bank.NewSimple[int](), cashbin.NewSimple()
	newATM := New[int](newBank, cashBin)
	sb := strings.Builder{}

	tests := []struct {
		name           string
		input          string
		r              io.Reader
		iter           int
		wantCardNumber string
		wantIsValid    bool
	}{
		{
			name:           "input=1",
			input:          "1\n",
			r:              nil,
			iter:           3,
			wantCardNumber: "",
			wantIsValid:    false,
		},
		{
			name:           "input=1234123412341234",
			input:          "1234123412341234\n",
			r:              nil,
			iter:           3,
			wantCardNumber: "1234123412341234",
			wantIsValid:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// simulate failures for iter times
			for i := 0; i < tt.iter; i++ {
				sb.WriteString(tt.input)
			}
			tt.r = strings.NewReader(sb.String())
			sb.Reset()

			cardNumber, isValid := newATM.verifyCardNumber(tt.r, tt.iter)
			if cardNumber != tt.wantCardNumber {
				t.Errorf("verifyCardNumber() cardNumber = %v, wantCardNumber %v", cardNumber, tt.wantCardNumber)
			}
			if isValid != tt.wantIsValid {
				t.Errorf("verifyCardNumber() isValid = %v, wantCardNumber %v", isValid, tt.wantIsValid)
			}
		})
	}
}

func TestATM_verifyPIN(t *testing.T) {
	//given
	newBank, cashBin := bank.NewSimple[int](), cashbin.NewSimple()
	newATM := New[int](newBank, cashBin)
	sb := strings.Builder{}

	tests := []struct {
		name        string
		input       string
		r           io.Reader
		iter        int
		wantPIN     string
		wantIsValid bool
	}{
		{
			name:        "input=1",
			input:       "1\n",
			r:           nil,
			iter:        3,
			wantPIN:     "",
			wantIsValid: false,
		},
		{
			name:        "input=1111",
			input:       "1111\n",
			r:           nil,
			iter:        3,
			wantPIN:     "1111",
			wantIsValid: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// simulate failures for iter times
			for i := 0; i < tt.iter; i++ {
				sb.WriteString(tt.input)
			}
			tt.r = strings.NewReader(sb.String())
			sb.Reset()

			pin, isValid := newATM.verifyPIN(tt.r, tt.iter)
			if pin != tt.wantPIN {
				t.Errorf("verifyPIN() pin = %v, wantOption %v", pin, tt.wantPIN)
			}
			if isValid != tt.wantIsValid {
				t.Errorf("verifyPIN() isValid = %v, wantOption %v", isValid, tt.wantIsValid)
			}
		})
	}
}
