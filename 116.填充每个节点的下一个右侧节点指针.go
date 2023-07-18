package leetcode

/*
 * @lc app=leetcode.cn id=116 lang=golang
 *
 * [116] 填充每个节点的下一个右侧节点指针
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	//定义一个数组根节点
	q := []*Node{root}
	// while 循环控制从上向下一层层遍历
	for len(q) > 0 {
		sz := len(q)
		// for 循环控制每一层从左向右遍历
		for i := 0; i < sz; i++ {
			cur := q[0]
			//弹出第一个元素
			q = q[1:]
			if i == sz-1 {
				cur.Next = nil
			} else {
				cur.Next = q[0]
			}
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
	}
	return root
}

// @lc code=end
