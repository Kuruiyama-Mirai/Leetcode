package leetcode

/*
 * @lc app=leetcode.cn id=501 lang=golang
 *
 * [501] 二叉搜索树中的众数
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
func findMode(root *TreeNode) []int {
	res := make([]int, 0)
	count, max := 1, 1
	var prev *TreeNode
	var traverse501 func(root *TreeNode)
	traverse501 = func(root *TreeNode) {
		if root == nil {
			return
		}
		traverse501(root.Left)
		if prev != nil && root.Val == prev.Val {
			count++
		} else {
			count = 1
		}
		if count >= max {
			if count > max && len(res) > 0 {
				res = []int{root.Val}
			} else {
				res = append(res, root.Val)
			}
			max = count
		}
		prev = root
		traverse501(root.Right)
	}
	traverse501(root)
	return res
}

// @lc code=end
