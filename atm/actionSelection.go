package atm

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/leegeobuk/atm-controller/bank/account"
)

func (atm *ATM[T]) promptBankActions(account account.BankAccount[T], iter int) {
	fmt.Printf("%s selected. ", account.Name())

	if option, isValid := atm.selectBankActions(os.Stdin, iter); !isValid {
		fmt.Printf(wrongInputMsg, "option", iter)
	} else if option == 1 {
		fmt.Printf("%s balance: %v\n", account.Name(), atm.bank.Balance(account))
	} else if option == 2 || option == 3 {
		atm.depositOrWithdraw(account, strconv.Itoa(option))
	} else if option == 4 {
		return
	} else if option == 5 {
		atm.exit()
	}
}

func (atm *ATM[T]) selectBankActions(r io.Reader, iter int) (int, bool) {
	scanner := bufio.NewScanner(r)
	count := 0
	var input string
	for true {
		fmt.Println("Select what to do.")
		fmt.Println("1. See balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Previous")
		fmt.Println("5. Exit")
		if scanner.Scan() {
			input = scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			return -1, false
		}

		option, err := strconv.Atoi(input)
		if isValid := 1 <= option && option <= 5; err != nil || !isValid {
			fmt.Println("Please enter from 1~5.")
			count++
			if count == iter {
				break
			}
		} else if isValid {
			return option, true
		}
	}

	return -1, false
}

func (atm *ATM[T]) depositOrWithdraw(account account.BankAccount[T], input string) {
	var amount string
	scanner := bufio.NewScanner(os.Stdin)
	m := map[string]string{"2": "deposit", "3": "withdraw"}
	fmt.Printf("Enter amount to %s: ", m[input])
	if scanner.Scan() {
		amount = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error accepting amount to %s: %v", m[input], err)
		return
	}

	value, err := atm.stringToNumber(amount)
	if err != nil {
		fmt.Println("Please enter numeric value.")
	}

	switch input {
	case "2":
		atm.bank.Deposit(account, value)
		fmt.Printf("%s balance: %v\n", account.Name(), atm.bank.Balance(account))
	case "3":
		if err = atm.bank.Withdraw(account, value); err != nil {
			fmt.Printf("Error while withdrawing: %v\n", err)
			return
		}
		fmt.Printf("%s balance: %v\n", account.Name(), atm.bank.Balance(account))
	}
}

func (atm *ATM[T]) stringToNumber(amount string) (T, error) {
	intVal, err := strconv.Atoi(amount)
	if err != nil {
		floatVal, err := strconv.ParseFloat(amount, 64)
		if err != nil {
			fmt.Println("Please enter numeric value.")
			return -1, errors.New("amount not numeric")
		}

		return T(floatVal), nil
	}

	return T(intVal), nil
}
