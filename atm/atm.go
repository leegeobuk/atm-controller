package atm

import (
    "fmt"
    "github.com/leegeobuk/atm-controller/bank"
    "github.com/leegeobuk/atm-controller/cashbin"
)

// ATM is where ATM controller starts from.
// It communicates with Bank and CashBin.
type ATM struct {
    bank bank.Bank
    cashBin cashbin.CashBin
}

// New returns ATM
func New(bank bank.Bank, cashBin cashbin.CashBin) *ATM {
    return &ATM{
        bank:    bank,
        cashBin: cashBin,
    }
}

func (atm *ATM) Start() {
    fmt.Println("ATM controller started")
    fmt.Println("Welcome! Please insert your card")

}

func (atm *ATM) printActions() {

}
