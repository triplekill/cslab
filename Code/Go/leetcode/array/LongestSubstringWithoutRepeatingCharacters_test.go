package array

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"f", args{"dvdf"}, 3},
		{"e", args{"au"}, 2},
		{"d", args{"c"}, 1},
		{"a", args{"abcabcbb"}, 3},
		{"b", args{"bbbbbb"}, 1},
		{"c", args{"pwwkew"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v, args %v", got, tt.want, tt.args)
			}
		})
	}
}
