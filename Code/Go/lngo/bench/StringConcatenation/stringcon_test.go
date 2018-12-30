// goos: darwin
// goarch: amd64
// pkg: github.com/exfly/cslab/Code/Go/lngo/bench/StringConcatenation
// BenchmarkConcatString-4         20000000               144 ns/op             530 B/op          0 allocs/op
// BenchmarkConcatBuffer-4         200000000                8.39 ns/op            2 B/op          0 allocs/op
// BenchmarkConcatBuilder-4        1000000000               2.38 ns/op            2 B/op          0 allocs/op
// PASS
// ok      github.com/exfly/cslab/Code/Go/lngo/bench/StringConcatenation   8.165s
package StringConcatenation_test

import (
	"bytes"
	"strings"
	"testing"
)

var strLen int = 1000

func BenchmarkConcatString(b *testing.B) {
	var str string

	i := 0

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		str += "x"

		i++
		if i >= strLen {
			i = 0
			str = ""
		}
	}
}

func BenchmarkConcatBuffer(b *testing.B) {
	var buffer bytes.Buffer

	i := 0

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		buffer.WriteString("x")

		i++
		if i >= strLen {
			i = 0
			buffer = bytes.Buffer{}
		}
	}
}

func BenchmarkConcatBuilder(b *testing.B) {
	var builder strings.Builder

	i := 0

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		builder.WriteString("x")

		i++
		if i >= strLen {
			i = 0
			builder = strings.Builder{}
		}
	}
}
