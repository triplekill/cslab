package array

import "testing"

func Test_findMaxConsecutiveOnes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"none", args{[]int{}}, 0},
		{"one", args{[]int{1}}, 1},
		{"none", args{[]int{0, 1}}, 1},
		{"simple", args{[]int{1, 1, 0, 1, 1, 1}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxConsecutiveOnes(tt.args.nums); got != tt.want {
				t.Errorf("findMaxConsecutiveOnes() = %v, want %v, args %v", got, tt.want, tt.args)
			}
		})
	}
}
