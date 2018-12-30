// goos: darwin
// goarch: amd64pkg: github.com/exfly/cslab/Code/Go/lngo/bench/RegularExpressions
// BenchmarkMatchString-4                    200000             11320 ns/op           42752 B/op           70 allocs/op
// BenchmarkMatchStringCompiled-4           2000000               663 ns/op               0 B/op            0 allocs/op
// PASS
// ok      github.com/exfly/cslab/Code/Go/lngo/bench/RegularExpressions    4.428s
package RegularExpressions_test

import (
	"regexp"
	"testing"
)

var testRegexp string = `^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]+$`

func BenchmarkMatchString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := regexp.MatchString(testRegexp, "jsmith@example.com")
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkMatchStringCompiled(b *testing.B) {
	r, err := regexp.Compile(testRegexp)
	if err != nil {
		panic(err)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r.MatchString("jsmith@example.com")
	}
}
