package array

import "testing"

func Test_searchInsert(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"a", args{[]int{1, 3, 5, 6}, 5}, 2},
		{"a", args{[]int{1, 3, 5, 6}, 2}, 1},
		{"a", args{[]int{1, 3, 5, 6}, 7}, 4},
		{"a", args{[]int{1, 3, 5, 6}, 0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchInsert() = %v, want %v, args %v", got, tt.want, tt.args)
			}
		})
	}
}
