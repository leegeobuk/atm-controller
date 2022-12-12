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

	for true {
		option, isValid := atm.selectBankAccount(accounts, os.Stdin, iter)
		if !isValid {
			fmt.Printf(wrongInputMsg, "option", iter)
			break
		} else if isValid {
			if option < len(accounts) {
				atm.promptBankActions(accounts[option], iter)
			} else if option == len(accounts) {
				break
			} else if option == len(accounts)+1 {
				atm.exit()
			}
		}
	}
}

func (atm *ATM[T]) selectBankAccount(accounts []account.BankAccount[T], r io.Reader, iter int) (int, bool) {
	scanner := bufio.NewScanner(r)
	actions := make([]string, 0, len(accounts)+2)
	for _, account := range accounts {
		actions = append(actions, account.Name())
	}
	actions = append(actions, "Back", "Exit")

	var input string
	for i := 0; i < iter; i++ {
		fmt.Println("Select bank account.")
		for idx, action := range actions {
			fmt.Printf("%d. %s\n", idx+1, action)
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

		if isValid := 1 <= option && option <= len(actions); !isValid {
			fmt.Printf("please enter between %d~%d\n", 1, len(actions))
		} else if isValid {
			return option - 1, true
		}
	}

	return -1, false
}
