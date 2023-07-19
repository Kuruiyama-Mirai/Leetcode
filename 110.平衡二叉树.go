package leetcode

/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] 平衡二叉树
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
func isBalanced(root *TreeNode) bool {
	h := getHeight(root)

	return h != -1
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := getHeight(root.Left), getHeight(root.Right)

	if left == -1 || right == -1 {
		return -1
	}
	if left-right > 1 || right-left > 1 {
		return -1
	}
	return Max110(left, right) + 1
}

func Max110(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
