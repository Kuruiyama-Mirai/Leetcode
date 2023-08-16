package leetcode

/*
 * @lc app=leetcode.cn id=129 lang=golang
 *
 * [129] 求根节点到叶节点数字之和
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
func sumNumbers(root *TreeNode) int {
	sum := 0
	dfs(root, root.Val, &sum)
	return sum
}

func dfs(root *TreeNode, tmpSum int, sum *int) {
	if root.Left == nil && root.Right == nil {
		*sum += tmpSum
	} else {
		if root.Left != nil {
			dfs(root.Left, tmpSum*10+root.Left.Val, sum)
		}
		if root.Right != nil {
			dfs(root.Right, tmpSum*10+root.Right.Val, sum)
		}
	}
}

// @lc code=end
