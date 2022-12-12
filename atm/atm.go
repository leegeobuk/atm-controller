package atm

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/leegeobuk/atm-controller/bank"
	"github.com/leegeobuk/atm-controller/cashbin"
	"github.com/leegeobuk/atm-controller/typeutil"
)

var (
	errInvalidInput = errors.New("invalid input")
	inputFailedMsg  = "Failed %s for %d times. Moving to previous screen.\n"
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

// Start starts atm and prompts main screen
func (atm *ATM[T]) Start() {
	fmt.Println("ATM controller started")

	const iter = 3
	for true {
		if option, err := atm.selectMainAction(os.Stdin, iter); err != nil {
			fmt.Printf("Invalid option entered for %d times. Program terminates.\n", iter)
			break
		} else if option == 1 {
			atm.promptVerification(iter)
		} else if option == 2 {
			fmt.Println("Exit selected.")
			atm.exit()
		}
	}
}

func (atm *ATM[T]) selectMainAction(r io.Reader, iter int) (int, error) {
	scanner := bufio.NewScanner(r)

	var input string
	for i := 0; i < iter; i++ {
		fmt.Println("How may I help you?")
		fmt.Println("1) Insert card")
		fmt.Println("2) Exit")

		if scanner.Scan() {
			input = scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			fmt.Printf("Error accepting action to choose: %v\n", err)
			continue
		}

		option, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Please enter valid number. Try again.")
			continue
		}

		if isValid := 1 <= option && option <= 2; !isValid {
			fmt.Println("please enter 1 or 2.")
			continue
		}

		return option, nil
	}

	return -1, errInvalidInput
}

func (atm *ATM[T]) exit() {
	fmt.Println("Bye bye.")
	os.Exit(0)
}

// Terminate prints the error and terminates the program.
func (atm *ATM[T]) Terminate(err error) {
	fmt.Println(err)
	os.Exit(1)
}
