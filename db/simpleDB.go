package db

import (
	"github.com/leegeobuk/atm-controller/bank/account"
	"github.com/leegeobuk/atm-controller/typeutil"
)

// SimpleDB implements DB.
// It uses memory to store data.
// No persistence is provided.
type SimpleDB[T typeutil.Number] struct {
	store map[string]account.BankAccount[T]
}

func NewSimple[T typeutil.Number]() *SimpleDB[T] {
	db := &SimpleDB[T]{}
	db.init()

	return db
}

func (db *SimpleDB[T]) init() {
	db.store = map[string]account.BankAccount[T]{
		"1234123412341234": account.NewSimpleChecking[T](12_341_234),
		"1111111111111111": account.NewSimpleChecking[T](11_111_111),
		"1234567812345678": account.NewSimpleChecking[T](12_345_678),
		"1000100010001000": account.NewSimpleChecking[T](10_001_000),
		"4321432143214321": account.NewSimpleChecking[T](43_214_321),
	}
}

func (db *SimpleDB[T]) GetAccount(cardNumber string) account.BankAccount[T] {
	return db.store[cardNumber]
}

func (db *SimpleDB[T]) UpdateAccount(bankAccount account.BankAccount[T]) {
	//TODO implement me
	panic("implement me")
}
