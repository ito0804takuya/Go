package main

import (
	"sync"
	"testing"
)

func TestCounters(t *testing.T) {

	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("it runs save concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		// ゴルーチンをまとめて管理するためのもの
		var wg sync.WaitGroup
		// 1000個のゴルーチン（1000回のインクリメント）を作成
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func(w *sync.WaitGroup) {
				counter.Inc()
				w.Done()
			}(&wg)
		}
		// 1000個のゴルーチン（1000回のインクリメント）が終わるまで待つ
		wg.Wait()

		assertCounter(t, counter, wantedCount)
	})

}

func assertCounter(t *testing.T, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d want %d", got.Value(), want)
	}
}
