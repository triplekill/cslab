package et

import (
	"reflect"
	"testing"
)

func TestGoSynax(t *testing.T) {

	// make slice
	a := make([]int, 3)
	t.Error(reflect.TypeOf(a))
	sp := make([]int, 10)
	t.Error(sp)

	// reflect
	t.Error(reflect.TypeOf(sp))
	sp = append(sp, 10)
	t.Error(sp)
	sp[0] = 1
	t.Error(sp)

	// int div
	t.Error(3 / 2)
	t.Error(3 / 2.0)
	t.Error(3.0 / 2)

}
