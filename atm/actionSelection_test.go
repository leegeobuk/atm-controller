package atm

import (
	"errors"
	"io"
	"strconv"
	"strings"
	"testing"

	"github.com/leegeobuk/atm-controller/typeutil"
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

func TestATM_promptAmount(t *testing.T) {
	// given
	testATM, _, largeInput := setup[int]()

	type tc[T typeutil.Number] struct {
		name    string
		input   string
		r       io.Reader
		option  int
		iter    int
		want    T
		wantErr error
	}
	tests := []tc[int]{
		{
			name:    "fail case: input=-1",
			input:   "-1\n",
			r:       nil,
			option:  2,
			iter:    3,
			want:    -1,
			wantErr: errInvalidInput,
		},
		{
			name:    "fail case: large input",
			input:   largeInput,
			r:       nil,
			option:  2,
			iter:    3,
			want:    -1,
			wantErr: errInvalidInput,
		},
		{
			name:    "fail case: parseInt error",
			input:   "s\n",
			r:       nil,
			option:  2,
			iter:    3,
			want:    -1,
			wantErr: errInvalidInput,
		},
		{
			name:    "success case: input=1",
			input:   "1\n",
			r:       nil,
			option:  2,
			iter:    3,
			want:    1,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// simulate failures for iter times
			tt.r = strings.NewReader(strings.Repeat(tt.input, tt.iter))

			got, err := testATM.promptAmount(tt.r, tt.option, tt.iter)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("promptAmount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("promptAmount() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringToNumber(t *testing.T) {
	type tc[T typeutil.Number] struct {
		name    string
		amount  string
		want    T
		wantErr error
	}

	tests := []tc[int]{
		{
			name:    "fail case: input != number",
			amount:  "s",
			want:    -1,
			wantErr: strconv.ErrSyntax,
		},
		{
			name:    "fail case: input too large",
			amount:  strings.Repeat("1", 20),
			want:    -1,
			wantErr: strconv.ErrRange,
		},
		{
			name:    "success case: input=1",
			amount:  "1",
			want:    1,
			wantErr: nil,
		},
		{
			name:    "success case: input=-1",
			amount:  "-1",
			want:    -1,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := stringToNumber[int](tt.amount)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("stringToNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("stringToNumber() got = %v, want %v", got, tt.want)
			}
		})
	}
}
