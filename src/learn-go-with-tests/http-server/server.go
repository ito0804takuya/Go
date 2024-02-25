package main

import (
	"fmt"
	"net/http"
)

// プレイヤーが勝ったゲームの数を追跡できるWebサーバー
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "20")
}
