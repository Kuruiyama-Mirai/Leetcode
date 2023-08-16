package leetcode

/*
 * @lc app=leetcode.cn id=108 lang=golang
 *
 * [108] 将有序数组转换为二叉搜索树
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
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	left, right := 0, len(nums)
	index := left + (right-left)/2
	root := &TreeNode{Val: nums[index]}

	root.Left = sortedArrayToBST(nums[:index])
	root.Right = sortedArrayToBST(nums[index+1:])

	return root
}

// @lc code=end
