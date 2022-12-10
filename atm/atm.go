package atm

import (
	"fmt"
	"github.com/leegeobuk/atm-controller/bank"
	"github.com/leegeobuk/atm-controller/cashbin"
	"os"
)

// ATM is where ATM controller starts from.
// It communicates with Bank and CashBin.
type ATM struct {
	bank    bank.Bank
	cashBin cashbin.CashBin
}

// New returns ATM
func New(bank bank.Bank, cashBin cashbin.CashBin) *ATM {
	return &ATM{
		bank:    bank,
		cashBin: cashBin,
	}
}

// Start starts atm
func (atm *ATM) Start() error {
	fmt.Println("ATM controller started")

	return atm.showActions()
}

func (atm *ATM) showActions() error {
	var input int
	for input != 1 || input != 2 {
		atm.showFirstScreen()
		_, err := fmt.Scanln(&input)
		if err != nil {
			return fmt.Errorf("start atm: %w", err)
		}

		if input == 1 {
			// verify card number
			if isValid := atm.verifyCardNumber(); !isValid {
				fmt.Println("Invalid card number entered for 3 times. Moving to previous screen.")
				continue
			}

			fmt.Print("Enter your PIN number (4 digits): ")
			var pin string
			fmt.Scanln(&pin)
			fmt.Println(pin)
			// verify PIN number

		} else if input == 2 {
			atm.exit()
		} else {
			fmt.Println("Please select 1 or 2")
		}
	}

	return nil
}

func (atm *ATM) showFirstScreen() {
	fmt.Println("How may I help you? ")
	fmt.Println("1) Insert card")
	fmt.Println("2) Exit")
}

func (atm *ATM) verifyCardNumber() bool {
	var cardNumber string
	for i := 0; i < 3; i++ {
		fmt.Print("Enter your card number (16 digits): ")
		_, err := fmt.Scanln(&cardNumber)
		if err != nil {
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

func (atm *ATM) exit() {
	fmt.Println("Exit selected. Bye bye.")
	os.Exit(0)
}

// Terminate prints the error and terminates the program.
func (atm *ATM) Terminate(err error) {
	fmt.Println(err)
	os.Exit(1)
}
