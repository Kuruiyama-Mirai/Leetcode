package leetcode

import "math"

/*
 * @lc app=leetcode.cn id=530 lang=golang
 *
 * [530] 二叉搜索树的最小绝对差
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
func getMinimumDifference(root *TreeNode) int {
	// 保留前一个节点的指针
	var prev *TreeNode
	// 定义一个比较大的值
	min := math.MaxInt64

	var traverse530 func(root *TreeNode)
	traverse530 = func(root *TreeNode) {
		if root == nil {
			return
		}
		traverse530(root.Left)
		if prev != nil && root.Val-prev.Val < min {
			min = root.Val - prev.Val
		}
		prev = root
		traverse530(root.Right)
	}

	traverse530(root)
	return min
}

// @lc code=end
