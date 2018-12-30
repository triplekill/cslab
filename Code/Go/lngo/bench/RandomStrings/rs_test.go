// goos: darwin
// goarch: amd64
// pkg: github.com/exfly/cslab/Code/Go/lngo/bench/RandomStrings
// BenchmarkMathRandString-4       20000000               106 ns/op              32 B/op          2 allocs/op
// BenchmarkCryptoRandString-4      3000000               471 ns/op              32 B/op          2 allocs/op
// PASS
// ok      github.com/exfly/cslab/Code/Go/lngo/bench/RandomStrings 4.106s
package RandomStrings_test

import (
	crand "crypto/rand"
	"math/rand"
	"testing"
)

// 64 letters
const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz+/"

func randomBytes(n int) []byte {
	bytes := make([]byte, n)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}

	return bytes
}

func cryptoRandomBytes(n int) []byte {
	bytes := make([]byte, n)
	_, err := crand.Read(bytes)
	if err != nil {
		panic(err)
	}

	return bytes
}

func randomString(bytes []byte) string {
	for i, b := range bytes {
		bytes[i] = letters[b%64]
	}

	return string(bytes)
}

func BenchmarkMathRandString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		randomString(randomBytes(16))
	}
}

func BenchmarkCryptoRandString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		randomString(cryptoRandomBytes(16))
	}
}
