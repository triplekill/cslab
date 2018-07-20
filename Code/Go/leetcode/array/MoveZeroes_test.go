package array

import (
	"testing"
)

func TestMoveZeros(t *testing.T) {
	arr := []int{0,1,0,3,12}

	moveZeroes(arr)
	if arr[2]!=12 {
		t.Error("move zero error")
	}
}