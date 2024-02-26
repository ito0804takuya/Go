package main

import (
	"fmt"
	"net/http"
	"strings"
)

// プレイヤーが勝ったゲームの数を追跡できるWebサーバー
type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	fmt.Fprint(w, p.store.GetPlayerScore(player))
}

type PlayerStore interface {
	GetPlayerScore(name string) int
}

func GetPlayerScore(name string) string {
	if name == "Pepper" {
		return "20"
	} else if name == "Floyd" {
		return "10"
	}
	return ""
}
