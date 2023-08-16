package leetcode

/*
 * @lc app=leetcode.cn id=1382 lang=golang
 *
 * [1382] 将二叉搜索树变平衡
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
func balanceBST(root *TreeNode) *TreeNode {
	res := make([]int, 0)
	var travesal1382 func(root *TreeNode)
	travesal1382 = func(root *TreeNode) {
		if root == nil {
			return
		}
		travesal1382(root.Left)
		res = append(res, root.Val)
		travesal1382(root.Right)
	}

	travesal1382(root)
	//将有序数组变为平衡二叉树
	var changeBalanceTree func(nums []int) *TreeNode
	changeBalanceTree = func(nums []int) *TreeNode {
		if len(nums) == 0 {
			return nil
		}
		left, right := 0, len(nums)
		index := left + (right-left)/2
		root := &TreeNode{Val: nums[index]}

		root.Left = changeBalanceTree(nums[:index])
		root.Right = changeBalanceTree(nums[index+1:])

		return root
	}

	return changeBalanceTree(res)
}

// @lc code=end
