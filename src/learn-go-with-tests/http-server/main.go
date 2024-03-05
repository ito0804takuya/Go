package main

import (
	"log"
	"net/http"
)

func main() {
	server := NewPlayerServer(NewInMemoryPlayerStore())

	// http.ListenAndServe : サーバーを起動
	log.Fatal(http.ListenAndServe(":15000", server))
}
