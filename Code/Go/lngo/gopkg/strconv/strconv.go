package main

import (
	"fmt"
	"strconv"
)

func main() {
	v := "100000000000000000000"
	if s, err := strconv.Atoi(v); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
}
