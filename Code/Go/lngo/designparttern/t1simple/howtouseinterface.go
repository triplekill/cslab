package t1simple

import "fmt"

type Api interface {
	Opt(s string) string
}

type Impl struct {
}

func (api *Impl) Opt(s string) string {
	fmt.Printf("Opt %v\n", s)
	return "Opt" + s
}
