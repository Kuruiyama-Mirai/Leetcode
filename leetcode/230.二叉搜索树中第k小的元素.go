package leetcode

/*
 * @lc app=leetcode.cn id=230 lang=golang
 *
 * [230] 二叉搜索树中第K小的元素
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
func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	res := midTraverse(root)

	return res[k-1]
}

func midTraverse(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	res = append(res, midTraverse(root.Left)...)
	res = append(res, root.Val)
	res = append(res, midTraverse(root.Right)...)

	return res
}

// @lc code=end
