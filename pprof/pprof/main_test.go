package pproftest

import (
	"container/list"
	"math/rand"
	"testing"
)

const size int = 20000

func BenchmarkList(b *testing.B) {
	for i := 0; i <= size; i++ {
		lt := list.New()
		for j := 0; j < size; j++ {
			lt.PushBack(rand.Intn(30))
		}
	}
}

func BenchmarkSlice(b *testing.B) {
	for i := 0; i <= size; i++ {
		var s []int
		for j := 0; j < size; j++ {
			s = append(s, rand.Intn(30))
		}
	}
}
