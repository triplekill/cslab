package main

import (
	"flag"
	"fmt"
	"net/http"
)

//	Usage:
// 		go run staticfile.go -path /your/path
//	eq:
//		go run staticfile.go -path .
func main() {
	path := flag.String("path", "./", "root path")
	flag.Parse()
	fmt.Println(*path)
	panic(http.ListenAndServe(":8080", http.FileServer(http.Dir(*path))))
}
