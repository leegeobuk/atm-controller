package atm

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func (atm *ATM) verificationStage(iter int) {
	// verify card number
	if isValid := atm.verifyCardNumber(os.Stdin, iter); !isValid {
		fmt.Printf("Invalid card number entered for %d times. Moving to previous screen.\n", iter)
		return
	}

	// verify PIN number
	if isValid := atm.verifyPIN(os.Stdin, iter); !isValid {
		fmt.Printf("Invalid PIN entered for %d times. Moving to previous screen.\n", iter)
	} else if isValid {
		atm.accountSelectionStage()
	}
}

func (atm *ATM) verifyCardNumber(r io.Reader, iter int) bool {
	scanner := bufio.NewScanner(r)
	var cardNumber string
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

func (atm *ATM) verifyPIN(r io.Reader, iter int) bool {
	scanner := bufio.NewScanner(r)
	var pin string
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
