package array

import (
	"testing"
)
func TestPascalTriangle(t *testing.T) {
	arr := generate(5)
	t.Log(arr)
	if arr[4][2] != 6 {
		t.Error("pascal`s tringle error")
	}
}
