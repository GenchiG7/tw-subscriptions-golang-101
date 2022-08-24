package pproftest

import (
	"testing"
)

const size int = 500000

// func BenchmarkList(b *testing.B) {
// 	for i := 0; i <= size; i++ {
// 		lt := list.New()
// 		for j := 0; j < size; j++ {
// 			lt.PushBack(rand.Intn(30))
// 		}
// 	}
// }

// func BenchmarkSlice(b *testing.B) {
// 	for i := 0; i <= size; i++ {
// 		var s []int
// 		for j := 0; j < size; j++ {
// 			s = append(s, rand.Intn(30))
// 		}
// 	}
// }

type Polygon1 struct {
	IsActive     bool
	ID           int64
	HasSchedule  bool
	Hash         string
	DeliveryCost float32
}

func BenchmarkPolygon1(b *testing.B) {

	hash := "1111111111111111"
	ps := make([]Polygon1, size)
	for i := 0; i < size; i++ {
		ps[i] = Polygon1{true, 1, true, hash, 1.0}
	}
}

type Polygon2 struct {
	Hash         string
	ID           int64
	DeliveryCost float32
	IsActive     bool
	HasSchedule  bool
}

func BenchmarkPolygon2(b *testing.B) {

	hash := "1111111111111111"
	ps := make([]Polygon2, size)
	for i := 0; i < size; i++ {
		ps[i] = Polygon2{hash, 1, 1.0, true, true}
	}
}
