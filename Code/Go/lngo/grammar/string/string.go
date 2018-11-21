package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func memlayout() {
	s := "hello, world世界"
	hello := s[:5]
	world := s[7:]
	fmt.Println(hello, world)

	fmt.Printf("%#v\n", []byte(s))                                          // 打印 []byte
	fmt.Println("len(s):", (*reflect.StringHeader)(unsafe.Pointer(&s)).Len) // 12
}

func utf8range() {
	fmt.Printf("%#v\n", []rune("a世界"))                 // []int32{97,19990, 30028}
	fmt.Printf("%#v\n", string([]rune{'a', '世', '界'})) // 世界
}
func main() {
	memlayout()
	utf8range()
}
