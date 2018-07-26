package array

// 删除链表的倒数第N个节点 https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/description/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p, q := head, head
	if head.Next == nil {
		return nil
	}

	for i := 0; i < n && q != nil; i++ {
		q = q.Next
	}
	if q == nil {
		return head.Next
	}
	for q != nil && q.Next != nil {
		p = p.Next
		q = q.Next
	}
	p.Next = p.Next.Next
	return head
}
