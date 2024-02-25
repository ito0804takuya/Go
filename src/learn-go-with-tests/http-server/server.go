package main

import (
	"fmt"
	"net/http"
	"strings"
)

// プレイヤーが勝ったゲームの数を追跡できるWebサーバー
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	if player == "Pepper" {
		fmt.Fprint(w, "20")
		return
	} else if player == "Floyd" {
		fmt.Fprint(w, "10")
		return
	}
}
