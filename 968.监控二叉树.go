package leetcode

/*
 * @lc app=leetcode.cn id=968 lang=golang
 *
 * [968] 监控二叉树
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
func minCameraCover(root *TreeNode) int {
	var traveral func(root *TreeNode) int
	//0：该节点无覆盖
	//1：本节点有摄像头
	//2：本节点有覆盖
	res := 0
	traveral = func(root *TreeNode) int {
		if root == nil {
			return 2
		}

		left := traveral(root.Left)
		right := traveral(root.Right)
		// 情况1
		// 左右节点都有覆盖
		if left == 2 && right == 2 {
			return 0
		} else if left == 0 || right == 0 {
			res++
			return 1
		} else {
			return 2
		}
	}
	//root几点未覆盖
	if traveral(root) == 0 {
		res++
	}
	return res
}

// @lc code=end
