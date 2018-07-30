package syntax

import (
	"reflect"
	"testing"
)

func Test_fib(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"a", args{4}, 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := factorial(tt.args.n); got != tt.want {
				t.Errorf("factorial() = %v, want %v, args %v", got, tt.want, tt.args)
			}
		})
	}
}

func Test_sliceAppend(t *testing.T) {
	type args struct {
		n int
	}
	t.Run("a", func(t *testing.T) { sliceAppend() })
}

func Test_maxSubsequnceSum(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"a", args{[]int{1, 2, 3}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubsequnceSum(tt.args.A); got != tt.want {
				t.Errorf("maxSubsequnceSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_insertSort(t *testing.T) {
	type args struct {
		A []int
		n int
	}
	tests := []struct {
		name string
		args args
		want args
	}{
		// TODO: Add test cases.
		{"a", args{[]int{8, 34, 64, 51, 32, 21}, 6}, args{[]int{8, 21, 32, 34, 51, 64}, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			insertSort(tt.args.A, tt.args.n)
		})
		t.Run(tt.name, func(t *testing.T) {
			insertSort(tt.args.A, tt.args.n)
			if !reflect.DeepEqual(tt.args.A, tt.want.A) {
				t.Errorf("insertSort() = %v, want %v", tt.args, tt.want)
			}
		})
	}
}
