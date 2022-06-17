package task3

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

type Set struct {
	sync.Map
	m map[int]struct{}
}

func NewSet() *Set {
	return &Set{
		m: map[int]struct{}{},
	}
}

func (s *Set) Add(i int) {
	s.Map.Store(i, struct{}{})
}

func (s *Set) Has(i int) bool {
	_, ok := s.Map.Load(i)
	return ok
}

func BenchmarkSet(b *testing.B) {
	var set = NewSet()
	rand.Seed(time.Now().UnixNano())
	dataSet := make([]int, 1e8) //make data set
	for i := 0; i < 1e8; i++ {
		dataSet[i] = rand.Intn(100) // add random values [0, n)
	}

	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			i := 0
			for pb.Next() {
				if dataSet[i] < 10 { // we can change the value depending on the percentage of the record. Now 10% of wrtie and 90% of read
					set.Add(1)
				} else {
					set.Has(1)
				}
			}
			i++
		})
	})
}
