package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		exp := Bitcoin(10)

		assertBallance(t, wallet, exp)
	})

	t.Run("Withdraw with funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		exp := Bitcoin(10)

		assertNoError(t, err)
		assertBallance(t, wallet, exp)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBallance(t, wallet, startingBalance)
	})
}

func assertBallance(t testing.TB, wallet Wallet, exp Bitcoin) {
	t.Helper()
	got := wallet.balance

	if got != exp {
		t.Errorf("\ngot: %s\nexpected: %s", got, exp)
	}
}

func assertError(t testing.TB, got, exp error) {
	t.Helper()

	if got == nil {
		t.Fatal("Wanted an error but did not get one")
	}

	if got != exp {
		t.Errorf("\ngot: %q\nexp: %q", got, exp)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()

	if got != nil {
		t.Fatal("Did not get error but wanted one")
	}
}
