package leetcode

/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
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

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftTree := maxDepth(root.Left)
	rightTree := maxDepth(root.Right)

	res := max(leftTree, rightTree) + 1
	return res
}

// @lc code=end
