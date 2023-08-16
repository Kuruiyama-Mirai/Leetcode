package leetcode

import "strconv"

/*
 * @lc app=leetcode.cn id=257 lang=golang
 *
 * [257] 二叉树的所有路径
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
func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)
	var travel func(root *TreeNode, s string)
	travel = func(root *TreeNode, s string) {
		if root.Left == nil && root.Right == nil {
			v := s + strconv.Itoa(root.Val)
			res = append(res, v)
			return
		}
		s = s + strconv.Itoa(root.Val) + "->"
		if root.Left != nil {
			travel(root.Left, s)
		}
		if root.Right != nil {
			travel(root.Right, s)
		}
	}
	travel(root, "")
	return res
}

// @lc code=end
