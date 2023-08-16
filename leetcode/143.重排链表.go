package leetcode

/*
 * @lc app=leetcode.cn id=143 lang=golang
 *
 * [143] 重排链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	queue := make([]*ListNode, 0)
	cur := head
	for cur != nil {
		queue = append(queue, cur)
		cur = cur.Next
	}

	if len(queue) <= 2 {
		return
	}

	flag := true
	//head要固定为原链表首元素
	queue = queue[1:]
	cur = head
	for len(queue) > 0 {
		if flag {
			cur.Next = queue[len(queue)-1]
			queue = queue[:len(queue)-1]
			//因为在原链表修改，不置空就会内存超出
			cur.Next.Next = nil

		} else {
			cur.Next = queue[0]
			queue = queue[1:]
			cur.Next.Next = nil
		}
		cur = cur.Next
		flag = !flag
	}
}

// @lc code=end
