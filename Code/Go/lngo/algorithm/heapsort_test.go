package algorithm

import (
	"reflect"
	"testing"
)

func TestHeapSort(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name   string
		args   args
		target []int
	}{
		// TODO: Add test cases.
		{"one", args{[]int{3, 2, 5, 2, 7, 5, 4}}, []int{2, 2, 3, 4, 5, 5, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HeapSort(tt.args.data)
			if !reflect.DeepEqual(tt.args.data, tt.target) {
				t.Errorf("%v %v", tt.args.data, tt.target)
			}
		})
	}
}
