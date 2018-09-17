package test

import (
	"log"
	"os"
	"runtime/pprof"
	"testing"
)

func Test_unmarshalt(t *testing.T) {
	f, err := os.Create("unmarshalt.prof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	for i := 0; i < 100000; i++ {
		unmarshalt()
	}
}

func BenchmarkFormat(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		unmarshalt()
	}
}

// 测试并发效率
func BenchmarkLoopsParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			unmarshalt()
		}
	})
}
