package main

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(10) // 入金

	got := wallet.Balance() // 残高
	want := 10

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}