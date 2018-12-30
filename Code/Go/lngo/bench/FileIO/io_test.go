// goos: darwin
// goarch: amd64
// pkg: github.com/exfly/cslab/Code/Go/lngo/bench/FileIO
// BenchmarkWriteFile-4                   3         533807710 ns/op             120 B/op          3 allocs/op
// BenchmarkWriteFileBuffered-4         500           3559292 ns/op            4216 B/op          4 allocs/op
// BenchmarkReadFile-4                   10         159982265 ns/op             120 B/op          3 allocs/op
// BenchmarkReadFileBuffered-4          200           6722392 ns/op         3204226 B/op     200004 allocs/op
// PASS
// ok      github.com/exfly/cslab/Code/Go/lngo/bench/FileIO        8.796s
package FileIO_test

import (
	"bufio"
	"io"
	"os"
	"testing"
)

func BenchmarkWriteFile(b *testing.B) {
	for n := 0; n < b.N; n++ {
		f, err := os.Create("/tmp/test.txt")
		if err != nil {
			panic(err)
		}

		for i := 0; i < 100000; i++ {
			f.WriteString("some text!\n")
		}

		f.Close()
	}
}

func BenchmarkWriteFileBuffered(b *testing.B) {
	for n := 0; n < b.N; n++ {
		f, err := os.Create("/tmp/test.txt")
		if err != nil {
			panic(err)
		}

		w := bufio.NewWriter(f)

		for i := 0; i < 100000; i++ {
			w.WriteString("some text!\n")
		}

		w.Flush()
		f.Close()
	}
}

func BenchmarkReadFile(b *testing.B) {
	for n := 0; n < b.N; n++ {
		f, err := os.Open("/tmp/test.txt")
		if err != nil {
			panic(err)
		}

		b := make([]byte, 10)

		_, err = f.Read(b)
		for err == nil {
			_, err = f.Read(b)
		}
		if err != io.EOF {
			panic(err)
		}

		f.Close()
	}
}

func BenchmarkReadFileBuffered(b *testing.B) {
	for n := 0; n < b.N; n++ {
		f, err := os.Open("/tmp/test.txt")
		if err != nil {
			panic(err)
		}

		r := bufio.NewReader(f)

		_, err = r.ReadString('\n')
		for err == nil {
			_, err = r.ReadString('\n')
		}
		if err != io.EOF {
			panic(err)
		}

		f.Close()
	}
}
