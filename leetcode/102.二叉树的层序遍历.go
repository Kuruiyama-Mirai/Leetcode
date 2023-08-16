package leetcode

/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder102(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	//定义一个数组根节点
	q := []*TreeNode{root}
	// while 循环控制从上向下一层层遍历
	for len(q) > 0 {
		sz := len(q)
		level := []int{}
		// for 循环控制每一层从左向右遍历
		for i := 0; i < sz; i++ {
			cur := q[0]
			//弹出第一个元素
			q = q[1:]
			level = append(level, cur.Val)
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		res = append(res, level)
	}
	return res
}

// @lc code=end
