package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("Expected %q, but got %q", expected, repeated)
	}
}

// go test -bench=. でベンチマークを実行して、この処理にかかる時間を計測できる
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
