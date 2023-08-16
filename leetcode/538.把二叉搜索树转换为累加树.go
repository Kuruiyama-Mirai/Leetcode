package leetcode

/*
 * @lc app=leetcode.cn id=538 lang=golang
 *
 * [538] 把二叉搜索树转换为累加树
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
func convertBST(root *TreeNode) *TreeNode {
	// 记录累加和
	var sum int
	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverse(node.Right)
		// 维护累加和
		sum += node.Val
		// 将 BST 转化成累加树
		node.Val = sum
		traverse(node.Left)
	}
	traverse(root)
	return root
}

// @lc code=end
