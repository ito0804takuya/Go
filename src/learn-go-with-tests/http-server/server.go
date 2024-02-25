package main

import (
	"fmt"
	"net/http"
	"strings"
)

// プレイヤーが勝ったゲームの数を追跡できるWebサーバー
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	fmt.Fprint(w, GetPlayerScore(player))
}

func GetPlayerScore(name string) string {
	if name == "Pepper" {
		return "20"
	} else if name == "Floyd" {
		return "10"
	}
	return ""
}
