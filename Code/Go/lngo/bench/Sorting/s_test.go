// goos: darwin
// goarch: amd64
// pkg: github.com/exfly/cslab/Code/Go/lngo/bench/Sorting
// BenchmarkSort1000-4                10000            141856 ns/op              32 B/op          1 allocs/op
// BenchmarkSort10000-4                1000           1827074 ns/op              32 B/op          1 allocs/op
// BenchmarkSort100000-4                 50          23026091 ns/op              32 B/op          1 allocs/op
// BenchmarkSort1000000-4                 5         256581906 ns/op              32 B/op          1 allocs/op
// PASS
// ok      github.com/exfly/cslab/Code/Go/lngo/bench/Sorting       8.540s
package Sorting_test

import (
	"math/rand"
	"sort"
	"testing"
)

func generateSlice(n int) []int {
	s := make([]int, 0, n)
	for i := 0; i < n; i++ {
		s = append(s, rand.Intn(1e9))
	}
	return s
}

func BenchmarkSort1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		s := generateSlice(1000)
		b.StartTimer()
		sort.Ints(s)
	}
}

func BenchmarkSort10000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		s := generateSlice(10000)
		b.StartTimer()
		sort.Ints(s)
	}
}

func BenchmarkSort100000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		s := generateSlice(100000)
		b.StartTimer()
		sort.Ints(s)
	}
}

func BenchmarkSort1000000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		s := generateSlice(1000000)
		b.StartTimer()
		sort.Ints(s)
	}
}
