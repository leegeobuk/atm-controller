package atm

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func (atm *ATM[T]) promptVerification(iter int) {
	cardNumber, err := atm.promptCardNumber(os.Stdin, iter)
	if err != nil {
		fmt.Printf(inputFailedMsg, "card number verification", iter)
		return
	}

	pin, err := atm.promptPIN(os.Stdin, iter)
	if err != nil {
		fmt.Printf(inputFailedMsg, "PIN verification", iter)
		return
	}

	//userAccount := account.NewUser(wantPIN, pin)
	atm.promptBankAccounts(cardNumber, pin, iter)
}

func (atm *ATM[T]) promptCardNumber(r io.Reader, iter int) (string, error) {
	scanner := bufio.NewScanner(r)

	var cardNumber string
	for i := 0; i < iter; i++ {
		fmt.Print("Enter your card number (16 digits): ")
		if scanner.Scan() {
			cardNumber = scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			fmt.Printf("Error accepting card number: %v\n", err)
			continue
		}

		if isValid := atm.bank.VerifyCardNumber(cardNumber); !isValid {
			fmt.Println("Card number doesn't exist. Please try again.")
			continue
		}

		return cardNumber, nil
	}

	return "", errInvalidInput
}

func (atm *ATM[T]) promptPIN(r io.Reader, iter int) (string, error) {
	scanner := bufio.NewScanner(r)

	var pin string
	for i := 0; i < iter; i++ {
		fmt.Print("Enter your PIN (4 digits): ")
		if scanner.Scan() {
			pin = scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			fmt.Printf("Error accepting PIN: %v\n", err)
			continue
		}

		if isValid := atm.bank.VerifyPIN(pin); !isValid {
			fmt.Println("PIN is not correct. Please try again.")
			continue
		}

		return pin, nil
	}

	return "", errInvalidInput
}
