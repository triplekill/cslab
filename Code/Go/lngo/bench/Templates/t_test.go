// goos: darwin
// goarch: amd64
// pkg: github.com/exfly/cslab/Code/Go/lngo/bench/Templates
// BenchmarkExecute-4        500000              2710 ns/op             168 B/op         12 allocs/op
// PASS
// ok      github.com/exfly/cslab/Code/Go/lngo/bench/Templates     1.390s
package Templates_test

import (
	"io/ioutil"
	"testing"
	"text/template"
)

var bookTemplate string = `
Title: {{.Title}},
Author: {{.Author}}
{{ if .Pages}}
Number of pages: {{ .Pages }}.
{{ end }}
{{ range .Chapters }}
{{ . }}, 
{{ end }}
`

type Book struct {
	Title    string   `json:"title"`
	Author   string   `json:"author"`
	Pages    int      `json:"num_pages"`
	Chapters []string `json:"chapters"`
}

var book *Book = &Book{
	Title:    "The Art of Computer Programming, Vol. 3",
	Author:   "Donald E. Knuth",
	Pages:    800,
	Chapters: []string{"Sorting", "Searching"},
}

func BenchmarkExecute(b *testing.B) {
	t := template.Must(template.New("book").Parse(bookTemplate))

	for n := 0; n < b.N; n++ {
		err := t.Execute(ioutil.Discard, book)
		if err != nil {
			panic(err)
		}
	}
}
