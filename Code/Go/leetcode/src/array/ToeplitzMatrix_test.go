package array

import (
	"testing"
)

func TestToeplitzMatri(t *testing.T) {
	var matrix = [][]int {{1,2,3,4},{5,1,2,3},{9,5,1,2}}
	if !isToeplitzMatrix(matrix) {
		t.Error("toeplitz matrix error")
	}
	matrix = [][]int {{1, 2}, {2, 2}}
	if isToeplitzMatrix(matrix) {
		t.Error("toeplitz matrix error")
	}
}