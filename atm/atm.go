package atm

import (
	"fmt"
	"os"

	"github.com/leegeobuk/atm-controller/bank"
	"github.com/leegeobuk/atm-controller/cashbin"
	"github.com/leegeobuk/atm-controller/typeutil"
)

// ATM is where ATM controller starts from.
// It communicates with Bank and CashBin.
type ATM[T typeutil.Number] struct {
	bank    bank.Bank[T]
	cashBin cashbin.CashBin
}

// New returns ATM
func New[T typeutil.Number](bank bank.Bank[T], cashBin cashbin.CashBin) *ATM[T] {
	return &ATM[T]{
		bank:    bank,
		cashBin: cashBin,
	}
}

// Start starts atm
func (atm *ATM[T]) Start() error {
	fmt.Println("ATM controller started")

	return atm.showActions()
}

func (atm *ATM[T]) showActions() error {
	const iter = 3
	var input int
	for true {
		atm.showMainScreen()
		_, err := fmt.Scanln(&input)
		if err != nil {
			return fmt.Errorf("start atm: %w", err)
		}

		if input == 1 {
			atm.verifyUser(iter)
		} else if input == 2 {
			atm.exit()
		} else {
			fmt.Println("Please select 1 or 2")
		}
	}

	return nil
}

func (atm *ATM[T]) showMainScreen() {
	fmt.Println("How may I help you? ")
	fmt.Println("1) Insert card")
	fmt.Println("2) Exit")
}

func (atm *ATM[T]) exit() {
	fmt.Println("Exit selected. Bye bye.")
	os.Exit(0)
}

// Terminate prints the error and terminates the program.
func (atm *ATM[T]) Terminate(err error) {
	fmt.Println(err)
	os.Exit(1)
}
