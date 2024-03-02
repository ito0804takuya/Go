package main

import (
	"log"
	"net/http"
)

type InMemoryPlayerScore struct {}

func (i *InMemoryPlayerScore) GetPlayerScore(name string) int {
	return 123
}

func (i *InMemoryPlayerScore) RecordWin(name string) {}

func main() {
	// http.HandlerFunc : func(w http.ResponseWriter, r *http.Request)という形の普通の関数を、
	//                    http.HandlerFunc型に変換することでhandlerが得られる。
	// handler : ServeHTTP というメソッドを実装しているもの。
	//           （http.HandlerFunc というinterfaceを実装しているもの。）
	// handler := http.HandlerFunc(PlayerServer)

	server := &PlayerServer{&InMemoryPlayerScore{}}

	// http.ListenAndServe : サーバーを起動
	log.Fatal(http.ListenAndServe(":15000", server))
}
