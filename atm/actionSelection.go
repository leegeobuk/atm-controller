package atm

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/leegeobuk/atm-controller/bank/account"
	"github.com/leegeobuk/atm-controller/typeutil"
)

func (atm *ATM[T]) promptBankActions(account account.BankAccount[T], iter int) {
	fmt.Printf("%s selected. ", account.Name())

	for true {
		if option, isValid := atm.selectBankActions(os.Stdin, iter); !isValid {
			fmt.Printf(wrongInputMsg, "option", iter)
			break
		} else if option == 1 {
			fmt.Printf("%s balance: %v\n", account.Name(), atm.bank.Balance(account))
		} else if option == 2 || option == 3 {
			amount, err := atm.promptAmount(os.Stdin, option, iter)
			if err != nil {
				fmt.Printf(wrongInputMsg, "amount", iter)
				continue
			}

			switch option {
			case 2:
				atm.bank.Deposit(account, amount)
				fmt.Printf("%s balance: %v\n", account.Name(), atm.bank.Balance(account))
			case 3:
				if err = atm.bank.Withdraw(account, amount); err != nil {
					fmt.Println("Withdrawal amount cannot be greater than the balance.")
				}
				fmt.Printf("%s balance: %v\n", account.Name(), atm.bank.Balance(account))
			}
		} else if option == 4 {
			break
		} else if option == 5 {
			atm.exit()
		}
	}
}

func (atm *ATM[T]) selectBankActions(r io.Reader, iter int) (int, bool) {
	scanner := bufio.NewScanner(r)
	var input string
	for i := 0; i < iter; i++ {
		fmt.Println("Select what to do.")
		fmt.Println("1. See balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Back")
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
		} else if isValid {
			return option, true
		}
	}

	return -1, false
}

// promptAmount prompts user to enter amount to deposit or withdraw.
func (atm *ATM[T]) promptAmount(r io.Reader, option, iter int) (T, error) {
	scanner := bufio.NewScanner(r)
	m := map[int]string{2: "deposit", 3: "withdraw"}

	var input string
	for i := 0; i < iter; i++ {
		fmt.Printf("Enter amount to %s: ", m[option])
		if scanner.Scan() {
			input = scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			fmt.Printf("Error accepting amount to %s: %v\n", m[option], err)
			continue
		}

		amount, err := stringToNumber[T](input)
		if err != nil {
			fmt.Println("Please enter numeric value.")
			continue
		}

		if amount < 0 {
			fmt.Println("Please enter positive value.")
			continue
		}

		return amount, nil
	}

	return -1, errInvalidInput
}

func stringToNumber[T typeutil.Number](input string) (T, error) {
	var amount T

	numType := fmt.Sprintf("%T", *new(T))
	switch numType {
	case "int":
		intVal, err := strconv.Atoi(input)
		if err != nil {
			return -1, err
		}

		amount = T(intVal)
	case "float64":
		floatVal, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return -1, err
		}

		amount = T(floatVal)
	}

	return amount, nil
}
