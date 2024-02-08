package main

import (
	"net/http"
)

func Racer(a, b string) (winner string) {
	// チャネルへの送信が早かったほうのURLを返す
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		// チャネルを閉じることで、"もう受信しない"ことを示す。また、リソースも解法できるらしい。
		close(ch)
	}()
	return ch
}
