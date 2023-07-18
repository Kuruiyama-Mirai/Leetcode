package leetcode

/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	node1, node2 := head, head
	for i := 0; i < k; i++ {
		if node2 == nil {
			return head
		}
		node2 = node2.Next
	}
	newHead := reverseTwoNode(node1, node2)
	node1.Next = reverseKGroup(node2, k)

	return newHead
}
func reverseTwoNode(head1, head2 *ListNode) *ListNode {
	pre, cur, next := &ListNode{-1, nil}, head1, head1
	for cur != head2 {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// @lc code=end
