package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	t.Run("Pepperのスコアを取得", func(t *testing.T) {
		// リクエストを生成
		request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
		// レスポンスを書き込みできるレコーダーを生成
		response := httptest.NewRecorder()

		// テスト対象のサーバー
		PlayerServer(response, request)

		got := response.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}