package et

import (
	"reflect"
	"testing"
)

func TestGoSynax(t *testing.T) {
	a := make([]int, 3)
	t.Error(reflect.TypeOf(a))
}
