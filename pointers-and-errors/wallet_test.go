package wallet

import (
	"testing"
)

func TestWall(t *testing.T) {

	wallet := Wallet{}

	wallet.Deposit(Bitcoin(10))

	got := wallet.Balance()

	// confirm that wallet is passed by address and not copied
	// fmt.Printf("address of balance in test is %v \n", &wallet.balance)

	want := Bitcoin(10)

	if got != want {
		t.Errorf("got %s want %s", got.String(), want.String())
	}

}
