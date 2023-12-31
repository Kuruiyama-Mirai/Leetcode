package leetcode

/*
 * @lc app=leetcode.cn id=404 lang=golang
 *
 * [404] 左叶子之和
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
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := sumOfLeftLeaves(root.Left)
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		left = root.Left.Val
	}
	right := sumOfLeftLeaves(root.Right)

	return left + right
}

// @lc code=end
