package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(10)) // 入金

	got := wallet.Balance() // 残高
	want := Bitcoin(10)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}