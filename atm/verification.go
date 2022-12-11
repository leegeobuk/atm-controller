package atm

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func (atm *ATM) verifyUser(iter int) {
	failedMsg := "Invalid %s entered for %d times. Moving to previous screen.\n"

	// verify card number
	var cardNumber string
	if isValid := atm.verifyCardNumber(cardNumber, os.Stdin, iter); !isValid {
		fmt.Printf(failedMsg, "card number", iter)
		return
	}

	// verify PIN number
	var pin string
	if isValid := atm.verifyPIN(pin, os.Stdin, iter); !isValid {
		fmt.Printf(failedMsg, "PIN", iter)
	} else if isValid {
		//userAccount := account.NewUser(cardNumber, pin)
		atm.selectAccounts(cardNumber, pin)
	}
}

func (atm *ATM) verifyCardNumber(cardNumber string, r io.Reader, iter int) bool {
	scanner := bufio.NewScanner(r)
	for i := 0; i < iter; i++ {
		fmt.Print("Enter your card number (16 digits): ")
		if scanner.Scan() {
			cardNumber = scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			return false
		}

		if isValid := atm.bank.VerifyCardNumber(cardNumber); !isValid {
			fmt.Println("Card number is not valid. Please try again.")
		} else if isValid {
			return true
		}
	}

	return false
}

func (atm *ATM) verifyPIN(pin string, r io.Reader, iter int) bool {
	scanner := bufio.NewScanner(r)
	for i := 0; i < iter; i++ {
		fmt.Print("Enter your PIN number (4 digits): ")
		if scanner.Scan() {
			pin = scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			return false
		}

		if isValid := atm.bank.VerifyPIN(pin); !isValid {
			fmt.Println("PIN is not valid. Please try again.")
		} else if isValid {
			return true
		}
	}

	return false
}
