package main

import (
	"log"
	"os"
	"reflect"
	"runtime"
)

func init() {
	log.Println("init main")
	log.SetOutput(os.Stderr)
}

func main() {
	log.Println("main")
	log.Println(reflect.TypeOf(10))
	runtime.GOMAXPROCS(1)

}
