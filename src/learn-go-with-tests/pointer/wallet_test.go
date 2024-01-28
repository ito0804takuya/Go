package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10)) // 入金
	
		got := wallet.Balance() // 残高
		want := Bitcoin(10)
	
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(15))

		got := wallet.Balance()
		want := Bitcoin(5)

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}