package atm

import (
	"fmt"
)

func (atm *ATM) selectAccounts(cardNumber, pin string) {
	fmt.Println("User verified. Select accounts")
	accounts := atm.bank.GetBankAccounts(cardNumber, pin)
	for i, bankAccount := range accounts {
		fmt.Printf("%d. %s\n", i+1, bankAccount.Name())
	}

}
