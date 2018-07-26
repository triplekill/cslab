package array

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"a", args{[]string{"flower", "flow", "flight"}}, "fl"},
		{"b", args{[]string{"dog", "race", "car"}}, ""},
		{"b", args{[]string{}}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v, args %v", got, tt.want, tt.args)
			}
		})
	}
}
