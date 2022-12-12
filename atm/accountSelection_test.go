package atm

import (
	"bufio"
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
	testATM, accounts, largeInput := setup[int]()

	tests := []struct {
		name        string
		accounts    []account.BankAccount[int]
		input       string
		r           io.Reader
		iter        int
		wantOption  int
		wantIsValid bool
	}{
		{
			name:        "fail case: input=-1, len(accounts)=1",
			accounts:    accounts,
			input:       "-1\n",
			r:           nil,
			iter:        3,
			wantOption:  -1,
			wantIsValid: false,
		},
		{
			name:        "fail case: input=-1, len(accounts)>1",
			accounts:    append(accounts, &account.SimpleCheckingAccount[int]{}),
			input:       "-1\n",
			r:           nil,
			iter:        3,
			wantOption:  -1,
			wantIsValid: false,
		},
		{
			name:        "scanner error case: large input",
			accounts:    accounts,
			input:       largeInput,
			r:           nil,
			iter:        3,
			wantOption:  -1,
			wantIsValid: false,
		},
		{
			name:        "strconv error case: input=\"\"",
			accounts:    accounts,
			input:       "",
			r:           nil,
			iter:        3,
			wantOption:  -1,
			wantIsValid: false,
		},
		{
			name:        "success case: input=1",
			accounts:    accounts,
			input:       "1\n",
			r:           nil,
			iter:        3,
			wantOption:  0,
			wantIsValid: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// simulate failures for iter times
			tt.r = strings.NewReader(strings.Repeat(tt.input, tt.iter))

			option, isValid := testATM.selectBankAccounts(tt.accounts, tt.r, tt.iter)
			if option != tt.wantOption {
				t.Errorf("selectBankAccounts() option = %v, wantOption %v", option, tt.wantOption)
			}
			if isValid != tt.wantIsValid {
				t.Errorf("selectBankAccounts() isValid = %v, wantIsValid %v", isValid, tt.wantIsValid)
			}
		})
	}
}

func setup[T typeutil.Number]() (*ATM[T], []account.BankAccount[T], string) {
	testBank, cashBin := bank.NewSimple[T](), cashbin.NewSimple()
	testATM := New[T](testBank, cashBin)
	accounts := []account.BankAccount[T]{&account.SimpleCheckingAccount[T]{}}
	largeInput := strings.Repeat("a", bufio.MaxScanTokenSize)

	return testATM, accounts, largeInput
}
