// goos: darwin
// goarch: amd64
// pkg: github.com/exfly/cslab/Code/Go/lngo/bench/NumericConversions
// BenchmarkParseBool-4            500000000                3.89 ns/op            0 B/op          0 allocs/op
// BenchmarkParseInt-4             100000000               21.8 ns/op             0 B/op          0 allocs/op
// BenchmarkParseFloat-4           50000000                36.5 ns/op             0 B/op          0 allocs/op
// PASS
// ok      github.com/exfly/cslab/Code/Go/lngo/bench/NumericConversions    6.347s
package NumericConversions_test

import (
	"strconv"
	"testing"
)

func BenchmarkParseBool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := strconv.ParseBool("true")
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkParseInt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := strconv.ParseInt("7182818284", 10, 64)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkParseFloat(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := strconv.ParseFloat("3.1415926535", 64)
		if err != nil {
			panic(err)
		}
	}
}
