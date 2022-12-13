package atm

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/leegeobuk/atm-controller/bank/account"
	"github.com/leegeobuk/atm-controller/typeutil"
)

func (atm *ATM[T]) promptBankActions(bankAccount account.BankAccount[T], iter int) {
	fmt.Printf("%s selected. ", bankAccount.Type())

	for true {
		if option, err := atm.selectBankActions(os.Stdin, iter); err != nil {
			fmt.Printf(inputFailedMsg, "bank action selection", iter)
			break
		} else if option == 1 {
			fmt.Printf("%s balance: %v\n", bankAccount.Type(), atm.bank.Balance(bankAccount))
		} else if option == 2 || option == 3 {
			amount, err := atm.promptAmount(os.Stdin, option, iter)
			if err != nil {
				fmt.Printf(inputFailedMsg, "entering amount", iter)
				continue
			}

			switch option {
			case 2:
				atm.bank.Deposit(bankAccount, amount)
				fmt.Printf("%s balance: %v\n", bankAccount.Type(), atm.bank.Balance(bankAccount))
				atm.cashBin.Deposit(amount)
				fmt.Printf("Cash bin cash: %v\n", atm.cashBin.Cash())
			case 3:
				if err = atm.bank.Withdraw(bankAccount, amount); errors.Is(err, account.ErrWithdrawAmount) {
					fmt.Println("Cannot withdraw more than balance.")
					continue
				} else if errors.Is(err, account.ErrWithdrawLimit) {
					fmt.Println("Cannot withdraw more than withdrawal limit.")
					continue
				}
				fmt.Printf("%s balance: %v\n", bankAccount.Type(), atm.bank.Balance(bankAccount))

				// withdraw from cash bin
				if err = atm.cashBin.Withdraw(amount); err != nil {
					fmt.Println("Cannot withdraw more than cash pool.")
				}
				fmt.Printf("Cash bin cash: %v\n", atm.cashBin.Cash())
			}
		} else if option == 4 {
			break
		} else if option == 5 {
			fmt.Println("Exit selected.")
			atm.exit()
		}
	}
}

func (atm *ATM[T]) selectBankActions(r io.Reader, iter int) (int, error) {
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
			fmt.Printf("Error accepting action to choose: %v\n", err)
			break
		}

		option, err := strconv.Atoi(input)
		if isValid := 1 <= option && option <= 5; err != nil || !isValid {
			fmt.Println("Please enter from 1~5.")
			continue
		}

		return option, nil
	}

	return -1, errInvalidInput
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
