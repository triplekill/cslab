package array

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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	h := &ListNode{0, nil} // h 链表头
	n := h                 // n 游标
	remain := 0            // remain 余数
	// tn 新创建的节点， 待插入
	for l1 != nil && l2 != nil {
		tn := &ListNode{0, nil}
		tn.Val = (l1.Val + l2.Val + remain) % 10
		remain = (l1.Val + l2.Val) / 10
		n.Next = tn
		n = tn
		l1 = l1.Next
		l2 = l2.Next
	}
	// 其中会有一个没有遍历完全，需要继续遍历
	for l1 != nil {
		tn := &ListNode{0, nil}
		tn.Val = (l1.Val + remain) % 10
		remain = l1.Val / 10
		n.Next = tn
		n = tn
	}
	for l2 != nil {
		tn := &ListNode{0, nil}
		tn.Val = (l1.Val + remain) % 10
		remain = l1.Val / 10
		n.Next = tn
		n = tn
	}
	if remain != 0 {
		tn := &ListNode{0, nil}
		tn.Val = remain
		n.Next = tn
		n = tn
	}
	return h.Next
}
