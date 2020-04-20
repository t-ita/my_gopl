package main

import (
	"testing"
)

func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo1([]string{"A", "B", "C", "D", "E", "F"})
	}
}

func BenchmarkEcho3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo3([]string{"A", "B", "C", "D", "E", "F"})
	}
}
