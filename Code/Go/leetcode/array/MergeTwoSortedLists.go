package array

/**
 * https://leetcode-cn.com/problems/merge-two-sorted-lists/description/
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	ret := &ListNode{}
	cur := ret
	for l1 != nil && l2 != nil {
		if l2.Val > l1.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	l := l1
	if l1 == nil {
		l = l2
	}
	if l != nil {
		cur.Next = l
	}
	return ret.Next
}

// https://leetcode-cn.com/problems/merge-k-sorted-lists/description/
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	n := len(lists)
	for n > 1 {
		k := (n + 1) / 2
		for i := 0; i < n/2; i++ {
			lists[i] = mergeTwoLists(lists[i], lists[i+k])
		}
		n = k
	}
	return lists[0]
}
