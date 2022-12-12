package atm

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/leegeobuk/atm-controller/bank"
	"github.com/leegeobuk/atm-controller/bank/account"
	"github.com/leegeobuk/atm-controller/cashbin"
	"github.com/leegeobuk/atm-controller/typeutil"
)

func TestATM_selectBankAccounts(t *testing.T) {
	//given
	testATM, bankAccount, largeInput := setup[int]()

	type tc[T typeutil.Number] struct {
		name       string
		account    account.BankAccount[T]
		input      string
		r          io.Reader
		iter       int
		wantOption int
		wantErr    error
	}

	tests := []tc[int]{
		{
			name:       "fail case: input=-1",
			account:    bankAccount,
			input:      "-1\n",
			r:          nil,
			iter:       3,
			wantOption: -1,
			wantErr:    errInvalidInput,
		},
		{
			name:       "fail case: large input",
			account:    bankAccount,
			input:      largeInput,
			r:          nil,
			iter:       3,
			wantOption: -1,
			wantErr:    errInvalidInput,
		},
		{
			name:       "fail case: input=s",
			account:    bankAccount,
			input:      "s\n",
			r:          nil,
			iter:       3,
			wantOption: -1,
			wantErr:    errInvalidInput,
		},
		{
			name:       "success case: input=1",
			account:    bankAccount,
			input:      "1\n",
			r:          nil,
			iter:       3,
			wantOption: 1,
			wantErr:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// simulate failures for iter times
			tt.r = strings.NewReader(strings.Repeat(tt.input, tt.iter))

			option, err := testATM.selectBankAccount(tt.account, tt.r, tt.iter)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("selectBankAccount() err = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if option != tt.wantOption {
				t.Errorf("selectBankAccount() option = %v, wantOption %v", option, tt.wantOption)
			}
		})
	}
}

func setup[T typeutil.Number]() (testATM *ATM[T], bankAccount account.BankAccount[T], largeInput string) {
	testBank, cashBin := bank.NewSimple[T](), cashbin.NewSimple()
	testATM = New[T](testBank, cashBin)
	bankAccount = &account.SimpleCheckingAccount[T]{}
	largeInput = fmt.Sprintf("%s\n", strings.Repeat("1", bufio.MaxScanTokenSize))

	return
}
