// goos: darwin
// goarch: amd64
// pkg: github.com/exfly/cslab/Code/Go/lngo/bench/RandomNumbers
// BenchmarkMathRand-4             100000000               21.1 ns/op             0 B/op          0 allocs/op
// BenchmarkCryptoRand-4            5000000               388 ns/op             161 B/op          5 allocs/op
// PASS
// ok      github.com/exfly/cslab/Code/Go/lngo/bench/RandomNumbers 4.424s
package RandomNumbers_test

import (
	crand "crypto/rand"
	"math/big"
	"math/rand"
	"testing"
)

func BenchmarkMathRand(b *testing.B) {
	for n := 0; n < b.N; n++ {
		rand.Int63()
	}
}

func BenchmarkCryptoRand(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := crand.Int(crand.Reader, big.NewInt(27))
		if err != nil {
			panic(err)
		}
	}
}
