package array

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	res1 := &ListNode{7, &ListNode{0, &ListNode{8, nil}}}

	l21 := &ListNode{1, &ListNode{8, nil}}
	l22 := &ListNode{0, nil}
	res2 := &ListNode{1, &ListNode{8, nil}}

	l31 := &ListNode{1, nil}
	l32 := &ListNode{9, &ListNode{9, nil}}
	res3 := &ListNode{0, &ListNode{0, &ListNode{1, nil}}}

	l41 := &ListNode{3, &ListNode{7, nil}}
	l42 := &ListNode{9, &ListNode{2, nil}}
	res4 := &ListNode{2, &ListNode{0, &ListNode{1, nil}}}

	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{"a", args{l1, l2}, res1},
		{"b", args{l21, l22}, res2},
		{"c", args{l31, l32}, res3},
		{"d", args{l41, l42}, res4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
