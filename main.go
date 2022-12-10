package main

import (
	"github.com/leegeobuk/atm-controller/atm"
	"github.com/leegeobuk/atm-controller/bank"
	"github.com/leegeobuk/atm-controller/cashbin"
)

func main() {
	simpleBank, cashBin := bank.NewSimple(), cashbin.NewSimple()
	simpleATM := atm.New(simpleBank, cashBin)
	if err := simpleATM.Start(); err != nil {
		simpleATM.Terminate(err)
	}
}
