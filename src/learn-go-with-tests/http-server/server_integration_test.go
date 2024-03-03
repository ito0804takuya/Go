package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// 統合テスト
func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryPlayerStore()
	server := PlayerServer{store}
	player := "Pepper"

	// 3回 勝ちを記録
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))
	
	assertStatus(t, response.Code, http.StatusOK)
	assertResponseBody(t, response.Body.String(), "3")
}
