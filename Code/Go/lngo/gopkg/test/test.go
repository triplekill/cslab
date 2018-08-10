package test

import "fmt"

var Gvar = 10

func init() {
	Gvar = 20
	fmt.Println("init")
}

func Change(val int) {
	Gvar = val
}
