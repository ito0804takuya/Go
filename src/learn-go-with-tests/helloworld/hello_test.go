package main

import "testing"

// ファイル名は _test.go である必要がある
// テスト関数名は Test で始まる必要がある
func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello World"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
