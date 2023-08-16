package leetcode

/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//链表问题 虚拟头节点必要的
	dummy := &ListNode{Val: -1, Next: head}
	x := findNextNthNode(dummy, n+1)
	x.Next = x.Next.Next
	return dummy.Next
}
func findNextNthNode(head *ListNode, n int) *ListNode {
	fast := head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	slow := head
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

// @lc code=end
