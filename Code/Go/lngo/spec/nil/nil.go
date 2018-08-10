package main

import "fmt"

type S struct {
	s string
}

func iS() S {
	return S{}
}

func arrayS() []S {
	return nil
}

func array() []int {
	return nil
}

func main() {
	fmt.Println(array())
	fmt.Println(nil == []int{})
	fmt.Println(len([]int{}))

	fmt.Println(len(array()))
	s := ""
	var sn string
	fmt.Println(len(s))
	fmt.Println(sn == "")
}
