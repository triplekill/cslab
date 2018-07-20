package array

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"a", args{[]int{7, 1, 5, 3, 6, 4}}, 7},
		{"b", args{[]int{1, 2, 3, 4, 5}}, 4},
		{"c", args{[]int{7, 6, 4, 3, 1}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v, args %v", got, tt.want, tt.args)
			}
		})
	}
}
