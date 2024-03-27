package main

import "testing"

// Benchmark_find-8   	238494856	         4.536 ns/op	       0 B/op	       0 allocs/op
func Benchmark_find(b *testing.B) {
	list := []int{1, 2, 3, 4, 5, 7, 9, 23, 34, 43, 65, 76, 90, 656}
	for i := 0; i < b.N; i++ {
		find(1, list)
	}
}
