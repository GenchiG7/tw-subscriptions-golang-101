package unittest

import (
	"testing"
)

var nums = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

func BenchmarkLinerSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		linerSearch(7, nums)
	}
}

func BenchmarkQuickSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		quickSearch(7, nums)
	}
}
