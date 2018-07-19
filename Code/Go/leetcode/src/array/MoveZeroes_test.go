package array

import (
	"testing"
)

func TestMoveZeros(t *testing.T) {
	arr := []int{0,1,0,3,12}
	t.Error(arr)
	moveZeroes(arr)
	t.Error(arr)
}