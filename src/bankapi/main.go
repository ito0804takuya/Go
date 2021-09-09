// 注: なぜか、このディレクトリに移動してgo run main.goしないと動かない
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/msft/bank"
)

var accounts = map[float64]*bank.Account{}

func main() {
	accounts[1001] = &bank.Account{
		Customer: bank.Customer{
			Name:    "John",
			Address: "Los Angeles",
			Phone:   "123 555 0147",
		},
		Number: 1001,
	}

	// ルーティング設定
	http.HandleFunc("/statement", statement)
	http.HandleFunc("/deposit", deposit)
	http.HandleFunc("/withdraw", withdraw)
	// webサーバ起動
	log.Fatal(http.ListenAndServe("localhost:9000", nil))
}

// 明細表示
// w http.ResponseWriter : レスポンスを書き戻すためのオブジェクト
// req *http.Request : httpリクエストオブジェクト
func statement(w http.ResponseWriter, req *http.Request) {
	// クエリ文字列から口座番号を取得
	numberqs := req.URL.Query().Get("number")

	if numberqs == "" {
		fmt.Fprintf(w, "口座番号が見つかりません")
		return
	}

	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "不正な口座番号です")
	} else {
		// 口座番号を使ってアカウントを取得
		account, ok := accounts[number]
		if !ok {
			fmt.Fprintf(w, "口座番号 %v は見つかりません", number)
		} else {
			fmt.Fprintf(w, account.Statement())
		}
	}
}

// 入金
func deposit(w http.ResponseWriter, req *http.Request) {
	// クエリ文字列から口座番号、入金額を取得
	numberqs := req.URL.Query().Get("number")
	amountqs := req.URL.Query().Get("amount")

	if numberqs == "" {
		fmt.Fprintf(w, "口座番号が見つかりません")
		return
	}

	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "不正な口座番号です")
	} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
		fmt.Fprintf(w, "不正な入金額です")
	} else {
		// 口座番号を使ってアカウントを取得
		account, ok := accounts[number]
		if !ok {
			fmt.Fprintf(w, "口座番号 %v は見つかりません", number)
		} else {
			err := account.Deposit(amount)
			if err != nil {
				fmt.Fprintf(w, "%v", err)
			} else {
				fmt.Fprintf(w, account.Statement())
			}
		}
	}
}

// 引き出し
func withdraw(w http.ResponseWriter, req *http.Request) {
	numberqs := req.URL.Query().Get("number")
	amountqs := req.URL.Query().Get("amount")

	if numberqs == "" {
		fmt.Fprintf(w, "口座番号が見つかりません")
		return
	}

	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "不正な口座番号です")
	} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
		fmt.Fprintf(w, "不正な入金額です")
	} else {
		// 口座番号を使ってアカウントを取得
		account, ok := accounts[number]
		if !ok {
			fmt.Fprintf(w, "口座番号 %v は見つかりません", number)
		} else {
			err := account.Withdraw(amount)
			if err != nil {
				fmt.Fprintf(w, "%v", err)
			} else {
				fmt.Fprintf(w, account.Statement())
			}
		}
	}
}