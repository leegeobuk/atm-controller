package atm

import (
	"io"
	"strings"
	"testing"
)

func TestATM_promptCardNumber(t *testing.T) {
	//given
	testATM, largeInput := setup[int]()

	tests := []struct {
		name           string
		input          string
		r              io.Reader
		iter           int
		wantCardNumber string
		wantErr        error
	}{
		{
			name:           "fail case: input=1",
			input:          "1\n",
			r:              nil,
			iter:           3,
			wantCardNumber: "",
			wantErr:        errInvalidInput,
		},
		{
			name:           "scanner error case: large input",
			input:          largeInput,
			r:              nil,
			iter:           3,
			wantCardNumber: "",
			wantErr:        errInvalidInput,
		},
		{
			name:           "success case: input=1234123412341234",
			input:          "1234123412341234\n",
			r:              nil,
			iter:           3,
			wantCardNumber: "1234123412341234",
			wantErr:        nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// simulate failures for iter times
			tt.r = strings.NewReader(strings.Repeat(tt.input, tt.iter))

			cardNumber, err := testATM.promptCardNumber(tt.r, tt.iter)
			if err != tt.wantErr {
				t.Errorf("promptCardNumber() err = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if cardNumber != tt.wantCardNumber {
				t.Errorf("promptCardNumber() cardNumber = %v, wantCardNumber %v", cardNumber, tt.wantCardNumber)
			}
		})
	}
}

func TestATM_promptPIN(t *testing.T) {
	//given
	testATM, largeInput := setup[int]()

	tests := []struct {
		name    string
		input   string
		r       io.Reader
		iter    int
		wantPIN string
		wantErr error
	}{
		{
			name:    "fail case: input=1",
			input:   "1\n",
			r:       nil,
			iter:    3,
			wantPIN: "",
			wantErr: errInvalidInput,
		},
		{
			name:    "fail case: large input",
			input:   largeInput,
			r:       nil,
			iter:    3,
			wantPIN: "",
			wantErr: errInvalidInput,
		},
		{
			name:    "fail case: input=1",
			input:   "1\n",
			r:       nil,
			iter:    3,
			wantPIN: "",
			wantErr: errInvalidInput,
		},
		{
			name:    "fail case: input=11",
			input:   "11\n",
			r:       nil,
			iter:    3,
			wantPIN: "",
			wantErr: errInvalidInput,
		},
		{
			name:    "fail case: input=111",
			input:   "111\n",
			r:       nil,
			iter:    3,
			wantPIN: "",
			wantErr: errInvalidInput,
		},
		{
			name:    "success case: input=1111",
			input:   "1111\n",
			r:       nil,
			iter:    3,
			wantPIN: "1111",
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// simulate failures for iter times
			tt.r = strings.NewReader(strings.Repeat(tt.input, tt.iter))

			pin, err := testATM.promptPIN(tt.r, tt.iter)
			if err != tt.wantErr {
				t.Errorf("promptPIN() err = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if pin != tt.wantPIN {
				t.Errorf("promptPIN() pin = %v, wantPIN %v", pin, tt.wantPIN)
			}
		})
	}
}
