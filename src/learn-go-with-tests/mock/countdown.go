package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

// 3から1秒ごとにカウントダウンして、最後はGo!と出力
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprintf(out, finalWord)
}

type DefaultSleeper struct {}

func (d DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func main() {
	sleeper := DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}

// timeもSpySleeperも、用途に応じてDIしたいため、インターフェースを定義
type Sleeper interface {
	Sleep()
}

// テスト用に使うSleeper
type SpySleeper struct {
	// 何回Sleep()が呼ばれたかカウント
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}
