package wallet

import "fmt"

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	// confirm that wallet is passed by address and not copied
	// fmt.Printf("address of balance in test is %v \n", &(w.balance))
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
