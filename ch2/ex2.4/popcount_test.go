package main

import (
	"../popcount"
	"testing"
)

func BenchmarkPopcount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(uint64(i))
	}
}

func BenchmarkPopcount24(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount24(uint64(i))
	}
}
