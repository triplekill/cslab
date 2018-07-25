package array

// ListNode listnode
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//https://leetcode-cn.com/problems/add-two-numbers/description/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	h := &ListNode{0, nil} // h 链表头
	n := h                 // n 游标
	remain := 0            // remain 余数
	// tn 新创建的节点， 待插入
	for l1 != nil && l2 != nil {
		tn := &ListNode{0, nil}
		tn.Val = (l1.Val + l2.Val + remain) % 10
		remain = (l1.Val + l2.Val + remain) / 10
		n.Next = tn
		n = tn
		l1 = l1.Next
		l2 = l2.Next
	}
	// 其中会有一个没有遍历完全，需要继续遍历
	var othern *ListNode
	if l1 != nil {
		othern = l1
	} else {
		othern = l2
	}
	for othern != nil {
		tn := &ListNode{0, nil}
		tn.Val = (othern.Val + remain) % 10
		remain = (othern.Val + remain) / 10
		n.Next = tn
		n = tn
		othern = othern.Next
	}

	for remain != 0 {
		tn := &ListNode{0, nil}
		tn.Val = remain % 10
		remain = remain / 10
		n.Next = tn
		n = tn
	}
	return h.Next
}
