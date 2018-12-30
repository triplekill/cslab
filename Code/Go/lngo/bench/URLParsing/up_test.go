// goos: darwin
// goarch: amd64
// pkg: github.com/exfly/cslab/Code/Go/lngo/bench/URLParsing
// BenchmarkParse-4         5000000               324 ns/op             128 B/op          1 allocs/op
// PASS
// ok      github.com/exfly/cslab/Code/Go/lngo/bench/URLParsing    1.951s
package URLParsing_test

import (
	"net/url"
	"testing"
)

func BenchmarkParse(b *testing.B) {
	testUrl := "https://www.example.com/path/file.html?param1=value1&param2=123"

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := url.Parse(testUrl)
		if err != nil {
			panic(err)
		}
	}
}
