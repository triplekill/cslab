package t2simplefactory

import "fmt"

type Api interface {
	Opt(s string) string
}

type ImplA struct {
}

func (api *ImplA) Opt(s string) string {
	s = "OptA:" + s
	fmt.Printf("%v\n", s)
	return s
}

type ImplB struct {
}

func (api *ImplB) Opt(s string) string {
	s = "OptB:" + s
	fmt.Printf("%v\n", s)
	return s
}

func CreateApiFactory(condition int) Api {
	switch condition {
	case 0:
		return &ImplA{}
	case 1:
		return &ImplB{}
	default:
		panic("unknown condition")
	}
}
