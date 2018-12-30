// goos: darwin
// goarch: amd64
// pkg: github.com/exfly/cslab/Code/Go/lngo/bench/Base64
// BenchmarkEncode-4        1000000              1618 ns/op            2816 B/op          2 allocs/op
// BenchmarkDecode-4         500000              2398 ns/op            2560 B/op          2 allocs/op
// PASS
// ok      github.com/exfly/cslab/Code/Go/lngo/bench/Base64        2.871s
package Base64_test

import (
	"encoding/base64"
	"math/rand"
	"testing"
)

func BenchmarkEncode(b *testing.B) {
	data := make([]byte, 1024)
	rand.Read(data)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		base64.StdEncoding.EncodeToString([]byte(data))
	}
}

func BenchmarkDecode(b *testing.B) {
	data := make([]byte, 1024)
	rand.Read(data)
	encoded := base64.StdEncoding.EncodeToString([]byte(data))

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := base64.StdEncoding.DecodeString(encoded)
		if err != nil {
			panic(err)
		}
	}
}
