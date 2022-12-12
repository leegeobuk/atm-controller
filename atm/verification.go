package atm

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

var (
	errInvalidInput = errors.New("invalid input")
	wrongInputMsg   = "Invalid %s entered for %d times. Moving to previous screen.\n"
)

func (atm *ATM[T]) verifyUser(iter int) {
	// verify card number
	cardNumber, isValid := atm.verifyCardNumber(os.Stdin, iter)
	if !isValid {
		fmt.Printf(wrongInputMsg, "card number", iter)
		return
	}

	// verify PIN number
	pin, isValid := atm.verifyPIN(os.Stdin, iter)
	if !isValid {
		fmt.Printf(wrongInputMsg, "PIN", iter)
	} else if isValid {
		//userAccount := account.NewUser(wantPIN, pin)
		atm.promptBankAccounts(cardNumber, pin, iter)
	}
}

func (atm *ATM[T]) verifyCardNumber(r io.Reader, iter int) (string, bool) {
	scanner := bufio.NewScanner(r)
	var cardNumber string
	for i := 0; i < iter; i++ {
		fmt.Print("Enter your card number (16 digits): ")
		if scanner.Scan() {
			cardNumber = scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			return "", false
		}

		if isValid := atm.bank.VerifyCardNumber(cardNumber); !isValid {
			fmt.Println("Card number is not valid. Please try again.")
		} else if isValid {
			return cardNumber, true
		}
	}

	return "", false
}

func (atm *ATM[T]) verifyPIN(r io.Reader, iter int) (string, bool) {
	scanner := bufio.NewScanner(r)
	var pin string
	for i := 0; i < iter; i++ {
		fmt.Print("Enter your PIN number (4 digits): ")
		if scanner.Scan() {
			pin = scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			return "", false
		}

		if isValid := atm.bank.VerifyPIN(pin); !isValid {
			fmt.Println("PIN is not valid. Please try again.")
		} else if isValid {
			return pin, true
		}
	}

	return "", false
}
