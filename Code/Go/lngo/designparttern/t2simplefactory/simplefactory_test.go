package t2simplefactory

import (
	"fmt"
	"testing"
)

func TestImplA_Opt(t *testing.T) {
	var cli Api
	cli = CreateApiFactory(0)
	fmt.Printf("%T %v\n", cli, cli)
	cli.Opt("factory")
}
