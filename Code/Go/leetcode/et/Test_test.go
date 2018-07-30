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

	l := []int{4, 5, 6}
	for _, i := range l {
		t.Error(i)
	}

	mp := make(map[int]bool)
	mp[1] = true
	mp[2] = false
	mp[3] = true
	for i, j := range mp {
		t.Error(i, j)
	}
	t.Error(-3 % 10)

	for 
}
