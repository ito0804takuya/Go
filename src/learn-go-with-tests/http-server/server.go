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

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

// ストア
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

func GetPlayerScore(name string) string {
	if name == "Pepper" {
		return "20"
	} else if name == "Floyd" {
		return "10"
	}
	return ""
}
