package main

import "testing"

func BenchmarkIntToHex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntToHex(int64(i))
	}
}
