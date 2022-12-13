package db

import (
	"github.com/leegeobuk/atm-controller/bank/account"
	_card "github.com/leegeobuk/atm-controller/bank/card"
	"github.com/leegeobuk/atm-controller/typeutil"
)

// SimpleDB implements DB.
// It uses memory to store data.
// No persistence is provided.
type SimpleDB[T typeutil.Number] struct {
	store map[string]*_card.Card[T]
}

// NewSimple returns new NewSimpleSavings account.
func NewSimple[T typeutil.Number]() *SimpleDB[T] {
	db := &SimpleDB[T]{}
	db.init()

	return db
}

func (db *SimpleDB[T]) init() {
	db.store = map[string]*_card.Card[T]{
		"1234123412341234": _card.New[T]("1234123412341234", "1234", account.NewSimpleChecking[T](12_341_234)),
		"1111111111111111": _card.New[T]("1111111111111111", "1111", account.NewSimpleChecking[T](11_111_111)),
		"1234567812345678": _card.New[T]("1234567812345678", "1234", account.NewSimpleChecking[T](12_345_678)),
		"1000100010001000": _card.New[T]("1000100010001000", "1000", account.NewSimpleSavings[T](10_001_000)),
		"4321432143214321": _card.New[T]("4321432143214321", "4321", account.NewSimpleSavings[T](43_214_321)),
	}
}

// GetCard returns card obtained from store.
func (db *SimpleDB[T]) GetCard(carNumber string) (*_card.Card[T], bool) {
	c, ok := db.store[carNumber]
	return c, ok
}

// GetAccount returns bank account
// linked to the card with given card number.
func (db *SimpleDB[T]) GetAccount(cardNumber string) (account.BankAccount[T], bool) {
	card, ok := db.store[cardNumber]
	return card.BankAccount(), ok
}

// UpdateAccount updates bank account.
func (db *SimpleDB[T]) UpdateAccount(bankAccount account.BankAccount[T]) {
	//TODO implement me
	panic("implement me")
}
