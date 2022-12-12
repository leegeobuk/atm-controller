package atm

import (
	"io"
	"strings"
	"testing"
)

func TestATM_selectBankActions(t *testing.T) {
	//given
	testATM, _, largeInput := setup[int]()

	tests := []struct {
		name        string
		input       string
		r           io.Reader
		iter        int
		wantOption  int
		wantIsValid bool
	}{
		{
			name:        "fail case: input=-1",
			input:       "-1\n",
			r:           nil,
			iter:        3,
			wantOption:  -1,
			wantIsValid: false,
		},
		{
			name:        "strconv error case: input=\"\"",
			input:       "\n",
			r:           nil,
			iter:        3,
			wantOption:  -1,
			wantIsValid: false,
		},
		{
			name:        "scanner error case: large input",
			input:       largeInput,
			r:           nil,
			iter:        3,
			wantOption:  -1,
			wantIsValid: false,
		},
		{
			name:        "success case: input=1",
			input:       "1\n",
			r:           nil,
			iter:        3,
			wantOption:  1,
			wantIsValid: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// simulate failures for iter times
			tt.r = strings.NewReader(strings.Repeat(tt.input, tt.iter))

			option, isValid := testATM.selectBankActions(tt.r, tt.iter)
			if option != tt.wantOption {
				t.Errorf("selectBankActions() = %v, wantOption %v", option, tt.wantOption)
			}
			if isValid != tt.wantIsValid {
				t.Errorf("selectBankActions() = %v, wantIsValid %v", isValid, tt.wantIsValid)
			}
		})
	}
}
