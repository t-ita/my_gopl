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

func BenchmarkPopcount23(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount23(uint64(i))
	}
}
