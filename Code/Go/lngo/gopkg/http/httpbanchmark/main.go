package main

import (
	"io"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	// Hello world, the web server

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		for i := 0; i < 100000000000000; i++ {
			io.WriteString(w, "Hello, world!\n")
		}
	}

	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
