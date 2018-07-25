package array

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"a", args{"42"}, 42},
		{"e", args{"4193 with words"}, 4193},
		{"e", args{"-4193 with words"}, -4193},
		{"c", args{"-42"}, -42},
		{"b", args{"    -42"}, -42},
		{"e", args{"4193 with words"}, 4193},
		{"f", args{"words and 987"}, 0},
		{"g", args{""}, 0},
		{"g", args{"  "}, 0},
		{"h", args{"-91283472332"}, -2147483648},
		{"i", args{"+4"}, 4},
		{"j", args{"-+1"}, 0},
		{"k", args{"+-1"}, 0},
		{"l", args{"9223372036854775808"}, 2147483647},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.str); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
