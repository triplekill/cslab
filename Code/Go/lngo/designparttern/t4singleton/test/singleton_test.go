package test

import (
	"fmt"
	"testing"

	"github.com/exfly/cslab/Code/Go/lngo/designparttern/t4singleton"
)

func Test_singleton_New(t *testing.T) {
	s := t4singleton.New()
	fmt.Println(s.A)
}
