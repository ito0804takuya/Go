package bank

import (
	"errors"
	"fmt"
)

// 顧客情報
type Customer struct {
	Name    string
	Address string
	Phone   string
}

// アカウント
type Account struct {
	Customer
	Number  int32   // 口座番号
	Balance float64 // 残高
}

// 入金
func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("入金額は0より大きい必要があります")
	}

	// 残高に入金額を足す
	a.Balance += amount
	return nil
}

// 引き出し
func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("残高は0より大きい必要があります")
	}

	if a.Balance < amount {
		return errors.New("残高より大きい額を引き出せません")
	}

	// 残高から引き出し額を引く
	a.Balance -= amount
	return nil
}

// 明細表示
func (a *Account) Statement() string {
	return fmt.Sprintf("%v - %v - %v", a.Number, a.Name, a.Balance)
	// data, _ := json.Marshal(a)
	// return data
}

// 送金
// 第一引数はポインタ(アドレス)
func (send_account *Account) Send(receive_account *Account, amount float64) error {
	if amount <= 0 {
		return errors.New("入金額は0より大きい必要があります")
	}
	if send_account.Balance < amount {
		return errors.New("残高より大きい額を送金できません")
	}
	send_account.Withdraw(amount)
	receive_account.Deposit(amount)
	return nil
}
