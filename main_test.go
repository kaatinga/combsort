package main

import (
	"reflect"
	"sort"
	"testing"
)

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

func BenchmarkSortBySort(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sort.Sort(testSlice)
	}
}

func Test_strStr(t *testing.T) {
	tests := []struct {
		haystack string
		needle   string
		want     int
	}{
		{"hello", "lo", 3},
		{"aaaaa", "baa", -1},
		{"abc", "c", 2},
		{"aaaaaa", "aa", 0},
		{"", "a", -1},
		{"mississippi", "issipi", -1},
		{"mississippi", "sippia", -1},
	}
	for _, tt := range tests {
		t.Run(tt.haystack, func(t *testing.T) {
			if got := strStr(tt.haystack, tt.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSort(t *testing.T) {
	tests := []struct {
		name   string
		input  toWorkSLice
		output toWorkSLice
	}{
		{"1", testSlice, rightSlice},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Sort(tt.input)
			if !reflect.DeepEqual(tt.input, tt.output) {
				t.Errorf("incorrect order:\n%v\n%v", tt.input, tt.output)
			}
		})
	}
}
