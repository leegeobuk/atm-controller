package atm

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/leegeobuk/atm-controller/bank/account"
	"github.com/leegeobuk/atm-controller/bank/card"
)

func (atm *ATM[T]) promptBankAccounts(c *card.Card[T], iter int) {
	fmt.Print("User verified. ")
	bankAccount := c.BankAccount()

	for true {
		option, err := atm.selectBankAccount(bankAccount, os.Stdin, iter)
		if err != nil {
			fmt.Printf(inputFailedMsg, "bank account selection", iter)
			break
		} else if option == 1 {
			atm.promptBankActions(bankAccount, iter)
		} else if option == 2 {
			break
		} else if option == 3 {
			fmt.Println("Exit selected.")
			atm.exit()
		}
	}
}

func (atm *ATM[T]) selectBankAccount(bankAccount account.BankAccount[T], r io.Reader, iter int) (int, error) {
	scanner := bufio.NewScanner(r)

	var input string
	for i := 0; i < iter; i++ {
		fmt.Println("Select bank account.")
		fmt.Printf("1. %s\n", bankAccount.Type())
		fmt.Println("2. Back")
		fmt.Println("3. Exit")

		if scanner.Scan() {
			input = scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			fmt.Printf("Error accepting account to choose: %v\n", err)
			continue
		}

		option, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Please enter valid number. Try again.")
			continue
		}

		if isValid := 1 <= option && option <= 3; !isValid {
			fmt.Println("please enter between 1~3")
			continue
		}

		return option, nil
	}

	return -1, errInvalidInput
}
