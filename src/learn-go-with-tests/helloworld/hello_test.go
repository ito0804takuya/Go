package main

import "testing"

// ファイル名は _test.go である必要がある
// テスト関数名は Test で始まる必要がある
func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		// この関数が（テスト本体でなく）テストのヘルパーであることを宣言
		t.Helper()
		
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("John", "")
		want := "Hello, John"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world when empty string is supplied", func(t *testing.T) {
		got := Hello("", "") // 空文字
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Japanese", func(t *testing.T) {
		got := Hello("太郎", "Japanese")
		want := "こんにちは, 太郎"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("", "French")
		want := "Bonjour, World"
		assertCorrectMessage(t, got, want)
	})

}
