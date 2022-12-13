package main

import (
	"github.com/leegeobuk/atm-controller/atm"
	"github.com/leegeobuk/atm-controller/bank"
	"github.com/leegeobuk/atm-controller/cashbin"
	"github.com/leegeobuk/atm-controller/db"
)

func main() {
	simpleDB := db.NewSimple[int]()
	simpleBank, cashBin := bank.NewSimple[int](simpleDB), cashbin.NewSimple[int](1_000_000_000)
	simpleATM := atm.New[int](simpleBank, cashBin)

	simpleATM.Start()
}
