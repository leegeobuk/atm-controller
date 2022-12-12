package atm

import (
	"io"
	"strings"
	"testing"
)

func TestATM_verifyCardNumber(t *testing.T) {
	//given
	testATM, _, largeInput := setup[int]()

	tests := []struct {
		name           string
		input          string
		r              io.Reader
		iter           int
		wantCardNumber string
		wantIsValid    bool
	}{
		{
			name:           "fail case: input=1",
			input:          "1\n",
			r:              nil,
			iter:           3,
			wantCardNumber: "",
			wantIsValid:    false,
		},
		{
			name:           "scanner error case: large input",
			input:          largeInput,
			r:              nil,
			iter:           3,
			wantCardNumber: "",
			wantIsValid:    false,
		},
		{
			name:           "success case: input=1234123412341234",
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
			tt.r = strings.NewReader(strings.Repeat(tt.input, tt.iter))

			cardNumber, isValid := testATM.verifyCardNumber(tt.r, tt.iter)
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
	testATM, _, largeInput := setup[int]()

	tests := []struct {
		name        string
		input       string
		r           io.Reader
		iter        int
		wantPIN     string
		wantIsValid bool
	}{
		{
			name:        "fail case: input=1",
			input:       "1\n",
			r:           nil,
			iter:        3,
			wantPIN:     "",
			wantIsValid: false,
		},
		{
			name:        "scanner error case: large input",
			input:       largeInput,
			r:           nil,
			iter:        3,
			wantPIN:     "",
			wantIsValid: false,
		},
		{
			name:        "success case: input=1111",
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
			tt.r = strings.NewReader(strings.Repeat(tt.input, tt.iter))

			pin, isValid := testATM.verifyPIN(tt.r, tt.iter)
			if pin != tt.wantPIN {
				t.Errorf("verifyPIN() pin = %v, wantOption %v", pin, tt.wantPIN)
			}
			if isValid != tt.wantIsValid {
				t.Errorf("verifyPIN() isValid = %v, wantOption %v", isValid, tt.wantIsValid)
			}
		})
	}
}
