package array

import (
	"reflect"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{"c", args{&ListNode{1, &ListNode{2, nil}}, 2}, &ListNode{2, nil}},
		{"b", args{&ListNode{1, &ListNode{2, nil}}, 1}, &ListNode{1, nil}},
		{"a", args{&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, 2}, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{5, nil}}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
