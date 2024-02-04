package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	
	// テストの実行に指定秒以上かかってしまう、遅いテストになってしまっているため、
	// 実際に指定秒数を待っているかどうかはテストの責任対象外にする という意思決定をし、
	// 何回Sleepが呼ばれたかをカウントする
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if spySleeper.Calls != 4 {
		t.Errorf("not enough calls to sleeper, want 4 got %d", spySleeper.Calls)
	}
}