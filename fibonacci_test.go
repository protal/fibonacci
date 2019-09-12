package fibonacci

import "testing"

func TestFIb(t *testing.T) {
	n := 10
	want := 55
	got := Fib(n)
	if got != want {
		t.Errorf(`failed, want %d, got %d`, want, got)
	}
}

func benchmarkFib(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(i)
	}
}

func BenchmarkFib1(b *testing.B) {
	benchmarkFib(1, b)
}
func BenchmarkFib2(b *testing.B) {
	benchmarkFib(2, b)
}
func BenchmarkFib3(b *testing.B) {
	benchmarkFib(3, b)
}
func BenchmarkFib4(b *testing.B) {
	benchmarkFib(4, b)
}
