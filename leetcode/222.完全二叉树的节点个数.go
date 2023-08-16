package leetcode

import "math"

/*
 * @lc app=leetcode.cn id=222 lang=golang
 *
 * [222] 完全二叉树的节点个数
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
func countNodes(root *TreeNode) int {
	l, r := root, root
	hl, hr := 0, 0

	for l != nil {
		l = l.Left
		hl++
	}
	for r != nil {
		r = r.Right
		hr++
	}
	if hl == hr {
		return int(math.Pow(2, float64(hl))) - 1
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

// @lc code=end
