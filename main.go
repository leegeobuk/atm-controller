package main

import (
	"github.com/leegeobuk/atm-controller/atm"
	"github.com/leegeobuk/atm-controller/bank"
	"github.com/leegeobuk/atm-controller/cashbin"
)

func main() {
	simpleBank, cashBin := bank.NewSimple[int](), cashbin.NewSimple()
	simpleATM := atm.New[int](simpleBank, cashBin)

	simpleATM.Start()
}
