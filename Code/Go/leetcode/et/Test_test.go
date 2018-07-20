package et

import (
	"reflect"
	"testing"
)

func TestGoSynax(t *testing.T) {
	a := make([]int, 3)
	t.Error(reflect.TypeOf(a))
	sp := make([]int, 10)
	t.Error(sp)
	t.Error(reflect.TypeOf(sp))
	sp = append(sp, 10)
	t.Error(sp)
	sp[0] = 1
	t.Error(sp)
}
