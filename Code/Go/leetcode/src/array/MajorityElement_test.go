package array

import (
	"testing"
)

func TestMajorityElement(t *testing.T) {
	arr := []int {3,2,3}
	res := majorityElement(arr)
	if res!=3 {
		t.Error("majority element error")
	}
}