package main

import "fmt"

// 口座
type Wallet struct {
	balance int // 残高
}

// 入金
// func (w Wallet) だと、w はWallet構造体の値のコピーなので、元の値が更新されない。
// func (w *Wallet) として、Wallet構造体のポインターとして読み取る
func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

// 残高
func (w *Wallet) Balance() int {
	// return (*w).balance と等価 (*演算子 : ポインターを参照して、（ポインターが示す）格納先のオブジェクトへアクセスする。)
	return w.balance
}
