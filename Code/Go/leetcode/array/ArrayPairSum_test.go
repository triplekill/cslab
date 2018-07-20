package array
import (
	"testing"
)

func TestArrayPairSum2(t * testing.T){
	var array = []int{1, 4, 3, 2}
	res := arrayPairSum2(array)
	if res != 4 {
		t.Error("ttttttt")
	}
}
