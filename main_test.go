package main

import "testing"

func BenchmarkSort(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Sort(testSlice)
	}
}

func BenchmarkSortByTree(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		SortByTree(testSlice)
	}
}
