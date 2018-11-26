package main

import (
	"fmt"
	"reflect"
)

// # [go doc reflect](https://golang.org/pkg/reflect/)
// # [laws of reflection](https://blog.golang.org/laws-of-reflection)
// # [research interface in go](https://research.swtch.com/interfaces)
func ref() {
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	fmt.Println("value:", reflect.ValueOf(x))
	fmt.Println("value:", reflect.ValueOf(x).String())

	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())
}
func main() {
	ref()
}
