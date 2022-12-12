package atm

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/leegeobuk/atm-controller/bank/account"
)

func (atm *ATM[T]) promptBankAccounts(cardNumber, pin string, iter int) {
	fmt.Print("User verified. ")
	accounts := atm.bank.GetBankAccounts(cardNumber, pin)

	option, isValid := atm.selectBankAccounts(accounts, os.Stdin, iter)
	if !isValid {
		fmt.Printf(wrongInputMsg, "option", iter)
	} else if isValid {
		atm.promptBankActions(accounts[option], iter)
	}
}

func (atm *ATM[T]) selectBankAccounts(accounts []account.BankAccount[T], r io.Reader, iter int) (int, bool) {
	scanner := bufio.NewScanner(r)
	var input string
	for i := 0; i < iter; i++ {
		fmt.Println("Select bank account.")
		for i, bankAccount := range accounts {
			fmt.Printf("%d. %s\n", i+1, bankAccount.Name())
		}

		if scanner.Scan() {
			input = scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			return -1, false
		}

		option, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Please enter valid number. Try again.")
			continue
		}

		if isValid := 1 <= option && option <= len(accounts); !isValid {
			if len(accounts) == 1 {
				fmt.Println("Please enter 1.")
			} else {
				fmt.Printf("please enter between %d~%d\n", 1, len(accounts))
			}
		} else if isValid {
			return option - 1, true
		}
	}

	return -1, false
}
