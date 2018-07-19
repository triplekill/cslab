package array
import (
	"testing"
)

func TestMatrixReshap(t * testing.T) {
	var row, col int= 1, 4
	var arr = [][]int {{1, 2}, {3, 4}}
	ret := matrixReshape( arr, row, col)
	if len(ret)!=row || len(ret[0])!=col {
		t.Error("reshape error")
	}
	t.Log(len(arr), arr)
	t.Log(len(arr[0]), arr[0])
	t.Log(ret)
	t.Log(len([][]int {{}}))
	t.Log(len([][]int {{}}[0]))
}
