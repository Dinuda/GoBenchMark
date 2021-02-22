package runecount

import "testing"

func BenchmarkRuneCount(b *testing.B) {
	s := "Gophers are amazing 😁"
	for i := 0; i < b.N; i++ {
		RuneCount(s)
	}
}
func BenchmarkRuneCount2(b *testing.B) {
	s := "Gophers are amazing 😁"
	for i := 0; i < b.N; i++ {
		RuneCount2(s)
	}
}
