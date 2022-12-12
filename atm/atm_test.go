package atm

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/leegeobuk/atm-controller/bank"
	"github.com/leegeobuk/atm-controller/cashbin"
	"github.com/leegeobuk/atm-controller/typeutil"
)

func TestATM_selectMainAction(t *testing.T) {
	// given
	testATM, largeInput := setup[int]()

	tests := []struct {
		name    string
		input   string
		r       io.Reader
		iter    int
		want    int
		wantErr error
	}{
		{
			name:    "fail case: input=-1",
			input:   "-1\n",
			r:       nil,
			iter:    3,
			want:    -1,
			wantErr: errInvalidInput,
		},
		{
			name:    "fail case: large input",
			input:   largeInput,
			r:       nil,
			iter:    3,
			want:    -1,
			wantErr: errInvalidInput,
		},
		{
			name:    "fail case: input=s",
			input:   "s\n",
			r:       nil,
			iter:    3,
			want:    -1,
			wantErr: errInvalidInput,
		},
		{
			name:    "success case: input=1",
			input:   "1\n",
			r:       nil,
			iter:    3,
			want:    1,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// simulate failures for iter times
			tt.r = strings.NewReader(strings.Repeat(tt.input, tt.iter))

			got, err := testATM.selectMainAction(tt.r, tt.iter)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("selectMainAction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("selectMainAction() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func setup[T typeutil.Number]() (testATM *ATM[T], largeInput string) {
	testBank, cashBin := bank.NewSimple[T](), cashbin.NewSimple()
	testATM = New[T](testBank, cashBin)
	largeInput = fmt.Sprintf("%s\n", strings.Repeat("1", bufio.MaxScanTokenSize))

	return
}
