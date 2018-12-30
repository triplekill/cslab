// goos: darwin
// goarch: amd64
// pkg: github.com/exfly/cslab/Code/Go/lngo/bench/Compression
// BenchmarkWrite-4        10000000               134 ns/op               0 B/op          0 allocs/op
// BenchmarkRead-4          1000000              1076 ns/op             624 B/op          2 allocs/op
// PASS
// ok      github.com/exfly/cslab/Code/Go/lngo/bench/Compression   2.580s
package Compression_test

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"testing"
)

func BenchmarkWrite(b *testing.B) {
	data, err := ioutil.ReadFile("test.json")
	if err != nil {
		panic(err)
	}

	zw := gzip.NewWriter(ioutil.Discard)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err = zw.Write(data)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkRead(b *testing.B) {
	data, err := ioutil.ReadFile("test.json")
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)
	_, err = zw.Write(data)
	if err != nil {
		panic(err)
	}

	err = zw.Close()
	if err != nil {
		panic(err)
	}

	r := bytes.NewReader(buf.Bytes())
	zr, _ := gzip.NewReader(r)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r.Reset(buf.Bytes())
		zr.Reset(r)
		_, err := ioutil.ReadAll(zr)
		if err != nil {
			panic(err)
		}
	}
}
