package main

import "testing"

func BenchmarkMu(b *testing.B) {
	example_3()
}
func BenchmarkRWMu(b *testing.B) {
	example_3RW()
}
