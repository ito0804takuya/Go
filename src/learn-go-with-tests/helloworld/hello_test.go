package main

import "testing"

// ファイル名は _test.go である必要がある
// テスト関数名は Test で始まる必要がある
func TestHello(t *testing.T) {
	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("John")
		want := "Hello, John"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("say hello world when empty string is supplied", func(t *testing.T) {
		got := Hello("") // 空文字
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

}
