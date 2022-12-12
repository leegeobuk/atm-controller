package atm

import (
	"bufio"
	"io"
	"strings"
	"testing"

	"github.com/leegeobuk/atm-controller/bank"
	"github.com/leegeobuk/atm-controller/bank/account"
	"github.com/leegeobuk/atm-controller/cashbin"
)

func TestATM_selectBankAccounts(t *testing.T) {
	//given
	newBank, cashBin := bank.NewSimple[int](), cashbin.NewSimple()
	newATM := New[int](newBank, cashBin)
	accounts := []account.BankAccount[int]{&account.SimpleCheckingAccount[int]{}}
	sb := strings.Builder{}
	for i := 0; i <= bufio.MaxScanTokenSize; i++ {
		sb.WriteByte('a')
	}
	scannerErrInput := sb.String()
	sb.Reset()

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
			name:        "scanner error case: input=a*bufio.MaxScanTokenSize",
			accounts:    accounts,
			input:       scannerErrInput,
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
			for i := 0; i < tt.iter; i++ {
				sb.WriteString(tt.input)
			}
			tt.r = strings.NewReader(sb.String())
			sb.Reset()

			option, isValid := newATM.selectBankAccounts(tt.accounts, tt.r, tt.iter)
			if option != tt.wantOption {
				t.Errorf("selectBankAccounts() option = %v, wantOption %v", option, tt.wantOption)
			}
			if isValid != tt.wantIsValid {
				t.Errorf("selectBankAccounts() isValid = %v, wantIsValid %v", isValid, tt.wantIsValid)
			}
		})
	}
}
