package main

import (
	"../popcount"
	"testing"
)

var results = make(map[int]int)

func BenchmarkPopcount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		results[i] = popcount.PopCount(123456789012345)
	}
}

func BenchmarkPopcount23(b *testing.B) {
	for i := 0; i < b.N; i++ {
		results[i] = popcount.PopCount23(123456789012345)
	}
}

func BenchmarkPopcount24(b *testing.B) {
	for i := 0; i < b.N; i++ {
		results[i] = popcount.PopCount23(123456789012345)
	}
}

func BenchmarkPopcount25(b *testing.B) {
	for i := 0; i < b.N; i++ {
		results[i] = popcount.PopCount23(123456789012345)
	}
}
