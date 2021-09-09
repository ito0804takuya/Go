package bank

import "testing"

// テスト
func TestAccount(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles",
			Phone:   "123 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}

	// 名前の入力が無ければアカウントを作れない
	if account.Name == "" {
		t.Error("アカウントを作れません")
	}
}

func TestDeposit(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles",
			Phone:   "123 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}

	// 入金
	account.Deposit(10)

	// 残高が入金額と等しくない場合、エラーを起こす
	if account.Balance != 10 {
		t.Error("入金したのに残高が更新されていません")
	}
}

func TestDepositInvalid(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles",
			Phone:   "123 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}

	// マイナスの金額を入金したのに、エラーがnilである場合、エラーを起こす
	if err := account.Deposit(-10); err == nil {
		t.Error("正の数値だけが入金額として許可されるべき")
	}
}

func TestWithdraw(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles",
			Phone:   "123 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}

	// 入金して、同じ額を引き出す
	account.Deposit(10)
	account.Withdraw(10)

	// 残高は0のはず
	if account.Balance != 0 {
		t.Error("引き出したのに残高が更新されていません")
	}
}

func TestStatement(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles",
			Phone:   "123 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}

	account.Deposit(100)
	statement := account.Statement()

	// 明細表示が正しいか
	if statement != "1001 - John - 100" {
		t.Error("明細表示が正しくありません")
	}
}

func TestSend(t *testing.T) {
	send_account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles",
			Phone:   "123 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}
	receive_account := Account{
		Customer: Customer{
			Name:    "Alice",
			Address: "Los Angeles",
			Phone:   "987 555 0147",
		},
		Number:  1002,
		Balance: 0,
	}

	send_account.Deposit(100)

	// 100を別アカウントへ送金
	send_account.Send(&receive_account, 100) // 重要: & でアドレスを渡す(参照渡し)
	
  // 明細表示が正しいか
	send_statement := send_account.Statement()
	if send_statement != "1001 - John - 0" {
		t.Error("明細表示が正しくありません")
	}

	receive_statement := receive_account.Statement()
	if receive_statement != "1002 - Alice - 100" {
		t.Error("明細表示が正しくありません")
	}
}