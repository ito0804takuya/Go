package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// テスト用のスタブstore
type StubStore struct {
	response string
}
func (s *StubStore) Fetch() string {
	return s.response
}

func TestHandler(t *testing.T) {
	// 期待するレスポンス
	data := "hello, world"
	svr := Server(&StubStore{response: data})

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	// レスポンスを受け取るレコーダー
	response := httptest.NewRecorder()

	svr.ServeHTTP(response, request)

	if response.Body.String() != data {
		t.Errorf(`got "%s", want "%s`, response.Body.String(), data)
	}
}
