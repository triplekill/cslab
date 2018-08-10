package main

import (
	"fmt"

	"github.com/exfly/lngo/gopkg/test"
)

func main() {
	fmt.Println("main")
	fmt.Println(test.Gvar)
	test.Change(30)
	fmt.Println(test.Gvar)
}
